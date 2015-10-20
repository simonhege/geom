// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package geojson implements decoding of GeoJSON objects as defined at http://geojson.org/. 

GeoJSON is a format for encoding a variety of geographic data structures.
*/
package geojson

import (
	"fmt"
	"reflect"

	"github.com/xeonx/geom"
)

//Geometry represents a GeoJSON geometry object
type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates,omitempty"`
	Geometries  []*Geometry `json:"geometries,omitempty"`
}

//ToGeoJSON creates a new GeoJSON geometry object based on the given geometry.ToGeoJSON
//It returns an error if the geometry type is not supported
func ToGeoJSON(g geom.Geometry) (*Geometry, error) {
	switch g := g.(type) {
	/* Point */
	case geom.Point:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y},
		}, nil
	case geom.PointZ:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y, g.Z},
		}, nil
	case geom.PointM:
		return &Geometry{
			Type:        "Point",
			Coordinates: []float64{g.X, g.Y, g.M},
		}, nil
	case geom.PointZM:
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
