// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package geojson implements decoding of GeoJSON objects as defined at http://geojson.org/.

GeoJSON is a format for encoding a variety of geographic data structures.
*/
package geojson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/xeonx/geom"
)

//FeatureCollection represents a feature collection
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

//Feature represents a object having a geometry and mutliple properties
type Feature struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

//A Decoder reads and decodes GeoJSON objects from an input stream
type Decoder struct {
	d *json.Decoder
}

//NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		d: json.NewDecoder(r),
	}
}

//DecodeCollection decodes the next JSON value as a FeatureCollection
func (dec *Decoder) DecodeCollection(c *FeatureCollection) error {
	return dec.d.Decode(c)
}

//DecodeFeature decodes the next JSON value as a Feature
func (dec *Decoder) DecodeFeature(f *Feature) error {
	return dec.d.Decode(f)
}

//Geometry represents a GeoJSON geometry object
type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates,omitempty"`
	Geometries  []*Geometry `json:"geometries,omitempty"`
}

func pointFromCoordinates(coord []float64) (geom.Geometry, error) {
	switch len(coord) {
	case 2:
		return &geom.Point{X: coord[0], Y: coord[1]}, nil
	case 3:
		return &geom.PointZ{Point: geom.Point{X: coord[0], Y: coord[1]}, Z: coord[2]}, nil
	case 4:
		return &geom.PointZM{PointZ: geom.PointZ{Point: geom.Point{X: coord[0], Y: coord[1]}, Z: coord[2]}, M: coord[3]}, nil
	}

	return nil, errors.New("Unsupported GeoJSON coordinates dimension")
}

func linestringFromCoordinates(coord [][]float64) (geom.Geometry, error) {
	if len(coord) == 0 {
		return geom.LineString{}, nil
	}

	switch len(coord[0]) {
	case 2:
		ret := geom.LineString{}

		for _, c := range coord {

			g, err := pointFromCoordinates(c)
			if err != nil {
				return nil, err
			}

			ret = append(ret, *g.(*geom.Point))
		}

		return ret, nil
	case 3:
		ret := geom.LineStringZ{}

		for _, c := range coord {

			g, err := pointFromCoordinates(c)
			if err != nil {
				return nil, err
			}

			ret = append(ret, *g.(*geom.PointZ))
		}

		return ret, nil
	case 4:
		ret := geom.LineStringZM{}

		for _, c := range coord {

			g, err := pointFromCoordinates(c)
			if err != nil {
				return nil, err
			}

			ret = append(ret, *g.(*geom.PointZM))
		}

		return ret, nil
	}

	return nil, errors.New("Unsupported GeoJSON coordinates dimension")
}

func polygonFromInterface(icoord []interface{}) (geom.Geometry, error) {

	var coord [][][]float64

	for _, icring := range icoord {

		iring := icring.([]interface{})
		var ring [][]float64

		for _, irpoint := range iring {

			ipoint := irpoint.([]interface{})

			var point []float64
			for _, ipcoordinates := range ipoint {
				point = append(point, ipcoordinates.(float64))
			}

			ring = append(ring, point)
		}

		coord = append(coord, ring)
	}

	return polygonFromCoordinates(coord)
}
func polygonFromCoordinates(coord [][][]float64) (geom.Geometry, error) {

	if len(coord) == 0 {
		return geom.Polygon{}, nil
	}

	dim := 0
	for _, ring := range coord {
		if len(ring) > 0 {
			dim = len(ring[0])
			break
		}
	}

	switch dim {
	case 2:
		ret := geom.Polygon{}

		for _, ring := range coord {

			g, err := linestringFromCoordinates(ring)
			if err != nil {
				return nil, err
			}

			ret = append(ret, g.(geom.LineString))
		}

		return ret, nil
	case 3:
		ret := geom.PolygonZ{}

		for _, ring := range coord {

			g, err := linestringFromCoordinates(ring)
			if err != nil {
				return nil, err
			}

			ret = append(ret, g.(geom.LineStringZ))
		}

		return ret, nil
	case 4:
		ret := geom.PolygonZM{}

		for _, ring := range coord {

			g, err := linestringFromCoordinates(ring)
			if err != nil {
				return nil, err
			}

			ret = append(ret, g.(geom.LineStringZM))
		}

		return ret, nil

	}

	return nil, errors.New("Unsupported GeoJSON coordinates dimension")

}

func multiPolygonFromInterface(icoord []interface{}) (geom.Geometry, error) {

	var coord [][][][]float64

	for _, icpoly := range icoord {

		ipoly := icpoly.([]interface{})
		var poly [][][]float64

		for _, icring := range ipoly {

			iring := icring.([]interface{})
			var ring [][]float64

			for _, irpoint := range iring {

				ipoint := irpoint.([]interface{})

				var point []float64
				for _, ipcoordinates := range ipoint {
					point = append(point, ipcoordinates.(float64))
				}

				ring = append(ring, point)
			}

			poly = append(poly, ring)
		}

		coord = append(coord, poly)
	}

	return multiPolygonFromCoordinates(coord)
}
func multiPolygonFromCoordinates(coord [][][][]float64) (geom.Geometry, error) {

	if len(coord) == 0 {
		return geom.MultiPolygon{}, nil
	}

	dim := 0
	for _, p := range coord {
		if len(p) > 0 {
			if len(p[0]) > 0 {
				dim = len(p[0][0])
				break
			}
		}
	}

	switch dim {
	case 2:
		ret := geom.MultiPolygon{}

		for _, p := range coord {

			g, err := polygonFromCoordinates(p)
			if err != nil {
				return nil, err
			}

			ret = append(ret, g.(geom.Polygon))
		}

		return ret, nil
	case 3:
		ret := geom.MultiPolygonZ{}

		for _, p := range coord {

			g, err := polygonFromCoordinates(p)
			if err != nil {
				return nil, err
			}

			ret = append(ret, g.(geom.PolygonZ))
		}

		return ret, nil
	case 4:
		ret := geom.MultiPolygonZM{}

		for _, p := range coord {

			g, err := polygonFromCoordinates(p)
			if err != nil {
				return nil, err
			}

			ret = append(ret, g.(geom.PolygonZM))
		}

		return ret, nil

	}

	return nil, errors.New("Unsupported GeoJSON coordinates dimension")

}

//FromGeoJSON creates a new Geometry object based on the GeoJSON geometry.
//
//It returns an error if the geometry type is not supported.
//Coordinates array with 3 values are interpreted as X/Y/Z.
func FromGeoJSON(g Geometry) (geom.Geometry, error) {

	switch g.Type {
	case "Point":
		coord := g.Coordinates.([]float64)

		return pointFromCoordinates(coord)
	case "LineString":
		coord := g.Coordinates.([][]float64)

		return linestringFromCoordinates(coord)
	case "Polygon":
		coord := g.Coordinates.([]interface{})

		return polygonFromInterface(coord)

	case "MultiPolygon":
		coord := g.Coordinates.([]interface{})

		return multiPolygonFromInterface(coord)
	}

	return nil, errors.New("Unsupported geometry type: " + g.Type)
}

//ToGeoJSON creates a new GeoJSON geometry object based on the given geometry.
//It returns an error if the geometry type is not supported
func ToGeoJSON(g geom.Geometry) (*Geometry, error) {
	switch g := g.(type) {
	/* Point */
	case *geom.Point:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y},
		}, nil
	case *geom.PointZ:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y, g.Z},
		}, nil
	case *geom.PointM:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y, g.M},
		}, nil
	case *geom.PointZM:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y, g.Z, g.M},
		}, nil
	/* LineString */
	case geom.LineString:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y}
		}
		return &Geometry{
			Type:        "LineString",
			Coordinates: coordinates,
		}, nil
	case geom.LineStringZ:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y, pt.Z}
		}
		return &Geometry{
			Type:        "LineString",
			Coordinates: coordinates,
		}, nil
	case geom.LineStringM:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y, pt.M}
		}
		return &Geometry{
			Type:        "LineString",
			Coordinates: coordinates,
		}, nil
	case geom.LineStringZM:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y, pt.Z, pt.M}
		}
		return &Geometry{
			Type:        "LineString",
			Coordinates: coordinates,
		}, nil
	/* Polygon */
	case geom.Polygon:
		coordinates := make([][][]float64, len(g))
		for i, ring := range g {
			coordinates[i] = make([][]float64, len(ring))
			for j, pt := range ring {
				coordinates[i][j] = []float64{pt.X, pt.Y}
			}
		}
		return &Geometry{
			Type:        "Polygon",
			Coordinates: coordinates,
		}, nil
	case geom.PolygonZ:
		coordinates := make([][][]float64, len(g))
		for i, ring := range g {
			coordinates[i] = make([][]float64, len(ring))
			for j, pt := range ring {
				coordinates[i][j] = []float64{pt.X, pt.Y, pt.Z}
			}
		}
		return &Geometry{
			Type:        "Polygon",
			Coordinates: coordinates,
		}, nil
	case geom.PolygonM:
		coordinates := make([][][]float64, len(g))
		for i, ring := range g {
			coordinates[i] = make([][]float64, len(ring))
			for j, pt := range ring {
				coordinates[i][j] = []float64{pt.X, pt.Y, pt.M}
			}
		}
		return &Geometry{
			Type:        "Polygon",
			Coordinates: coordinates,
		}, nil
	case geom.PolygonZM:
		coordinates := make([][][]float64, len(g))
		for i, ring := range g {
			coordinates[i] = make([][]float64, len(ring))
			for j, pt := range ring {
				coordinates[i][j] = []float64{pt.X, pt.Y, pt.Z, pt.M}
			}
		}
		return &Geometry{
			Type:        "Polygon",
			Coordinates: coordinates,
		}, nil
	/* MultiPoint */
	case geom.MultiPoint:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y}
		}
		return &Geometry{
			Type:        "MultiPoint",
			Coordinates: coordinates,
		}, nil
	case geom.MultiPointZ:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y, pt.Z}
		}
		return &Geometry{
			Type:        "MultiPoint",
			Coordinates: coordinates,
		}, nil
	case geom.MultiPointM:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y, pt.M}
		}
		return &Geometry{
			Type:        "MultiPoint",
			Coordinates: coordinates,
		}, nil
	case geom.MultiPointZM:
		coordinates := make([][]float64, len(g))
		for i, pt := range g {
			coordinates[i] = []float64{pt.X, pt.Y, pt.Z, pt.M}
		}
		return &Geometry{
			Type:        "MultiPoint",
			Coordinates: coordinates,
		}, nil
	/* MultiLineString */
	case geom.MultiLineString:
		coordinates := make([][][]float64, len(g))
		for i, linestring := range g {
			coordinates[i] = make([][]float64, len(linestring))
			for j, pt := range linestring {
				coordinates[i][j] = []float64{pt.X, pt.Y}
			}
		}
		return &Geometry{
			Type:        "MultiLineString",
			Coordinates: coordinates,
		}, nil
	case geom.MultiLineStringZ:
		coordinates := make([][][]float64, len(g))
		for i, linestring := range g {
			coordinates[i] = make([][]float64, len(linestring))
			for j, pt := range linestring {
				coordinates[i][j] = []float64{pt.X, pt.Y, pt.Z}
			}
		}
		return &Geometry{
			Type:        "MultiLineString",
			Coordinates: coordinates,
		}, nil
	case geom.MultiLineStringM:
		coordinates := make([][][]float64, len(g))
		for i, linestring := range g {
			coordinates[i] = make([][]float64, len(linestring))
			for j, pt := range linestring {
				coordinates[i][j] = []float64{pt.X, pt.Y, pt.M}
			}
		}
		return &Geometry{
			Type:        "MultiLineString",
			Coordinates: coordinates,
		}, nil
	case geom.MultiLineStringZM:
		coordinates := make([][][]float64, len(g))
		for i, linestring := range g {
			coordinates[i] = make([][]float64, len(linestring))
			for j, pt := range linestring {
				coordinates[i][j] = []float64{pt.X, pt.Y, pt.Z, pt.M}
			}
		}
		return &Geometry{
			Type:        "MultiLineString",
			Coordinates: coordinates,
		}, nil
	/* MultiPolygon */
	case geom.MultiPolygon:
		coordinates := make([][][][]float64, len(g))
		for p, polygon := range g {
			coordinates[p] = make([][][]float64, len(polygon))
			for r, ring := range polygon {
				coordinates[p][r] = make([][]float64, len(ring))
				for i, pt := range ring {
					coordinates[p][r][i] = []float64{pt.X, pt.Y}
				}
			}
		}
		return &Geometry{
			Type:        "MultiPolygon",
			Coordinates: coordinates,
		}, nil
	case geom.MultiPolygonZ:
		coordinates := make([][][][]float64, len(g))
		for p, polygon := range g {
			coordinates[p] = make([][][]float64, len(polygon))
			for r, ring := range polygon {
				coordinates[p][r] = make([][]float64, len(ring))
				for i, pt := range ring {
					coordinates[p][r][i] = []float64{pt.X, pt.Y, pt.Z}
				}
			}
		}
		return &Geometry{
			Type:        "MultiPolygon",
			Coordinates: coordinates,
		}, nil
	case geom.MultiPolygonM:
		coordinates := make([][][][]float64, len(g))
		for p, polygon := range g {
			coordinates[p] = make([][][]float64, len(polygon))
			for r, ring := range polygon {
				coordinates[p][r] = make([][]float64, len(ring))
				for i, pt := range ring {
					coordinates[p][r][i] = []float64{pt.X, pt.Y, pt.M}
				}
			}
		}
		return &Geometry{
			Type:        "MultiPolygon",
			Coordinates: coordinates,
		}, nil
	case geom.MultiPolygonZM:
		coordinates := make([][][][]float64, len(g))
		for p, polygon := range g {
			coordinates[p] = make([][][]float64, len(polygon))
			for r, ring := range polygon {
				coordinates[p][r] = make([][]float64, len(ring))
				for i, pt := range ring {
					coordinates[p][r][i] = []float64{pt.X, pt.Y, pt.Z, pt.M}
				}
			}
		}
		return &Geometry{
			Type:        "MultiPolygon",
			Coordinates: coordinates,
		}, nil
	/* GeometryCollection */
	case geom.GeometryCollection:
		geometries := make([]*Geometry, len(g))
		var err error
		for i, child := range g {
			geometries[i], err = ToGeoJSON(child)
			if err != nil {
				return nil, err
			}
		}
		return &Geometry{
			Type:       "GeometryCollection",
			Geometries: geometries,
		}, nil
	case geom.GeometryCollectionZ:
		geometries := make([]*Geometry, len(g))
		var err error
		for i, child := range g {
			geometries[i], err = ToGeoJSON(child)
			if err != nil {
				return nil, err
			}
		}
		return &Geometry{
			Type:       "GeometryCollection",
			Geometries: geometries,
		}, nil
	case geom.GeometryCollectionM:
		geometries := make([]*Geometry, len(g))
		var err error
		for i, child := range g {
			geometries[i], err = ToGeoJSON(child)
			if err != nil {
				return nil, err
			}
		}
		return &Geometry{
			Type:       "GeometryCollection",
			Geometries: geometries,
		}, nil
	case geom.GeometryCollectionZM:
		geometries := make([]*Geometry, len(g))
		var err error
		for i, child := range g {
			geometries[i], err = ToGeoJSON(child)
			if err != nil {
				return nil, err
			}
		}
		return &Geometry{
			Type:       "GeometryCollection",
			Geometries: geometries,
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported geometry type: %s", reflect.TypeOf(g).String())
	}
}
