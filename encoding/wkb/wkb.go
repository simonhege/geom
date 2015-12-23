// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package wkb implements decoding of WKB objects as defined in OGC 06-103r4.

WKB is a binary format for geometry encoding. It is described in OGC 06-103r4 OpenGIS®
Implementation Standard for Geographic information - Simple feature access - Part 1:
Common architecture Version: 1.2.1 2011-05-28
http://portal.opengeospatial.org/files/?artifact_id=25355 (also ISO/TC211 19125 Part 1)
*/
package wkb

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"

	"github.com/xeonx/geom"
)

const (
	//XDR is the flag used in WKB to represents Big Endian encoding
	XDR = 0x00
	//LDR is the flag used in WKB to represents Little Endian encoding
	LDR = 0x01
)

//Type represents a geometry type as defined in the WKB specification
type Type uint32

//WKB types as defined in the WKB specification
const (
	WKBPoint              Type = 1
	WKBLineString         Type = 2
	WKBPolygon            Type = 3
	WKBMultiPoint         Type = 4
	WKBMultiLineString    Type = 5
	WKBMultiPolygon       Type = 6
	WKBGeometryCollection Type = 7
)

//Flatten returns the flattened (ie. 2-dimensionnal) WKB type for the given type
func (Type Type) Flatten() Type {
	return Type % 1000
}

//HasZ returns true if the WKB type as a Z component
func (Type Type) HasZ() bool {
	return (Type >= 1000 && Type < 2000) || (Type >= 3000 && Type < 4000)
}

//HasM returns true if the WKB type as a M component
func (Type Type) HasM() bool {
	return (Type >= 2000 && Type < 4000)
}

//Dimensioner represents objects able indicate if they have a Z or M component.
type Dimensioner interface {
	HasZ() bool
	HasM() bool
}

//Read reads a WKB geometry
func Read(r io.Reader) (geom.Geometry, error) {

	//Read byte order
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.BigEndian, &wkbByteOrder); err != nil {
		return nil, err
	}
	var byteOrder binary.ByteOrder = binary.BigEndian
	if wkbByteOrder == LDR {
		byteOrder = binary.LittleEndian
	} else if wkbByteOrder != XDR {
		return nil, fmt.Errorf("Invalid WKB byte order. Expecting %d or %d, got '%d'", XDR, LDR, wkbByteOrder)
	}

	//Read geometry type
	var Type Type
	if err := binary.Read(r, byteOrder, &Type); err != nil {
		return nil, err
	}

	switch Type.Flatten() {
	case WKBPoint:
		return readWkbPoint(r, byteOrder, Type)
	case WKBLineString:
		return readWkbLineString(r, byteOrder, Type)
	case WKBPolygon:
		return readWkbPolygon(r, byteOrder, Type)
	case WKBMultiPoint:
		return readWkbMultiPoint(r, byteOrder, Type)
	case WKBMultiLineString:
		return readWkbMultiLineString(r, byteOrder, Type)
	case WKBMultiPolygon:
		return readWkbMultiPolygon(r, byteOrder, Type)
	case WKBGeometryCollection:
		return readWkbGeometryCollection(r, byteOrder, Type)
	}

	return nil, fmt.Errorf("Unsupported WKB geometry type. Got '%d'", Type)
}

func readWkbPoint(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var pt geom.Point
	var z, m float64

	if err := binary.Read(r, byteOrder, &pt.X); err != nil {
		return nil, err
	}
	if err := binary.Read(r, byteOrder, &pt.Y); err != nil {
		return nil, err
	}
	if dim.HasZ() {
		if err := binary.Read(r, byteOrder, &z); err != nil {
			return nil, err
		}
	}
	if dim.HasM() {
		if err := binary.Read(r, byteOrder, &m); err != nil {
			return nil, err
		}
	}

	if dim.HasZ() {

		ptz := geom.PointZ{
			Point: pt,
			Z:     z,
		}

		if dim.HasM() {
			return &geom.PointZM{
				PointZ: ptz,
				M:      m,
			}, nil
		}

		return &ptz, nil

	} else if dim.HasM() {
		return &geom.PointM{
			Point: pt,
			M:     m,
		}, nil
	}

	return &pt, nil
}
func readWkbLineString(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var numPoints uint32
	if err := binary.Read(r, byteOrder, &numPoints); err != nil {
		return nil, err
	}

	if dim.HasZ() {
		if dim.HasM() {
			line := make(geom.LineStringZM, numPoints)
			if err := binary.Read(r, byteOrder, &line); err != nil {
				return nil, err
			}
			return line, nil
		}

		line := make(geom.LineStringZ, numPoints)
		if err := binary.Read(r, byteOrder, &line); err != nil {
			return nil, err
		}
		return line, nil

	} else if dim.HasM() {
		line := make(geom.LineStringM, numPoints)
		if err := binary.Read(r, byteOrder, &line); err != nil {
			return nil, err
		}
		return line, nil
	}

	line := make(geom.LineString, numPoints)
	if err := binary.Read(r, byteOrder, &line); err != nil {
		return nil, err
	}
	return line, nil
}
func readWkbPolygon(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var numRings uint32
	if err := binary.Read(r, byteOrder, &numRings); err != nil {
		return nil, err
	}

	geoms := make([]geom.Geometry, numRings)

	for i := uint32(0); i < numRings; i++ {
		geom, err := readWkbLineString(r, byteOrder, dim)
		if err != nil {
			return nil, err
		}
		geoms[i] = geom
	}

	var ok bool
	if dim.HasZ() {
		if dim.HasM() {
			rings := make(geom.PolygonZM, len(geoms))
			for i, g := range geoms {
				rings[i], ok = g.(geom.LineStringZM)
				if !ok {
					return nil, fmt.Errorf("Unexpected child geometry type in PolygonZM: %s", reflect.TypeOf(g).String())
				}
			}
			return rings, nil
		}

		rings := make(geom.PolygonZ, len(geoms))
		for i, g := range geoms {
			rings[i], ok = g.(geom.LineStringZ)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in PolygonZ: %s", reflect.TypeOf(g).String())
			}
		}
		return rings, nil

	} else if dim.HasM() {
		rings := make(geom.PolygonM, len(geoms))
		for i, g := range geoms {
			rings[i], ok = g.(geom.LineStringM)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in PolygonM: %s", reflect.TypeOf(g).String())
			}
		}
		return rings, nil
	}

	rings := make(geom.Polygon, len(geoms))
	for i, g := range geoms {
		rings[i], ok = g.(geom.LineString)
		if !ok {
			return nil, fmt.Errorf("Unexpected child geometry type in Polygon: %s", reflect.TypeOf(g).String())
		}
	}
	return rings, nil
}
func readWkbMultiPoint(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var numPoints uint32
	if err := binary.Read(r, byteOrder, &numPoints); err != nil {
		return nil, err
	}

	geoms := make([]geom.Geometry, numPoints)

	for i := uint32(0); i < numPoints; i++ {
		geom, err := Read(r)
		if err != nil {
			return nil, err
		}
		geoms[i] = geom
	}

	if dim.HasZ() {
		if dim.HasM() {
			points := make(geom.MultiPointZM, len(geoms))
			for i, g := range geoms {
				pt, ok := g.(*geom.PointZM)
				if !ok {
					return nil, fmt.Errorf("Unexpected child geometry type in MultiPointZM: %s", reflect.TypeOf(g).String())
				}
				points[i] = *pt
			}
			return points, nil
		}

		points := make(geom.MultiPointZ, len(geoms))
		for i, g := range geoms {
			pt, ok := g.(*geom.PointZ)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in MultiPointZ: %s", reflect.TypeOf(g).String())
			}
			points[i] = *pt
		}
		return points, nil

	} else if dim.HasM() {
		points := make(geom.MultiPointM, len(geoms))
		for i, g := range geoms {
			pt, ok := g.(*geom.PointM)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in MultiPointM: %s", reflect.TypeOf(g).String())
			}
			points[i] = *pt
		}
		return points, nil
	}

	points := make(geom.MultiPoint, len(geoms))
	for i, g := range geoms {
		pt, ok := g.(*geom.Point)
		if !ok {
			return nil, fmt.Errorf("Unexpected child geometry type in MultiPoint: %s", reflect.TypeOf(g).String())
		}
		points[i] = *pt
	}
	return points, nil
}
func readWkbMultiLineString(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var numLineStrings uint32
	if err := binary.Read(r, byteOrder, &numLineStrings); err != nil {
		return nil, err
	}

	geoms := make([]geom.Geometry, numLineStrings)

	for i := uint32(0); i < numLineStrings; i++ {
		geom, err := Read(r)
		if err != nil {
			return nil, err
		}
		geoms[i] = geom
	}

	var ok bool
	if dim.HasZ() {
		if dim.HasM() {
			lines := make(geom.MultiLineStringZM, len(geoms))
			for i, g := range geoms {
				lines[i], ok = g.(geom.LineStringZM)
				if !ok {
					return nil, fmt.Errorf("Unexpected child geometry type in MultiLineStringZM: %s", reflect.TypeOf(g).String())
				}
			}
			return lines, nil
		}

		lines := make(geom.MultiLineStringZ, len(geoms))
		for i, g := range geoms {
			lines[i], ok = g.(geom.LineStringZ)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in MultiLineStringZ: %s", reflect.TypeOf(g).String())
			}
		}
		return lines, nil

	} else if dim.HasM() {
		lines := make(geom.MultiLineStringM, len(geoms))
		for i, g := range geoms {
			lines[i], ok = g.(geom.LineStringM)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in MultiLineStringM: %s", reflect.TypeOf(g).String())
			}
		}
		return lines, nil
	}

	lines := make(geom.MultiLineString, len(geoms))
	for i, g := range geoms {
		lines[i], ok = g.(geom.LineString)
		if !ok {
			return nil, fmt.Errorf("Unexpected child geometry type in MultiLineString: %s", reflect.TypeOf(g).String())
		}
	}
	return lines, nil
}
func readWkbMultiPolygon(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var numPolygons uint32
	if err := binary.Read(r, byteOrder, &numPolygons); err != nil {
		return nil, err
	}

	geoms := make([]geom.Geometry, numPolygons)

	for i := uint32(0); i < numPolygons; i++ {
		geom, err := Read(r)
		if err != nil {
			return nil, err
		}
		geoms[i] = geom
	}

	var ok bool
	if dim.HasZ() {
		if dim.HasM() {
			polygons := make(geom.MultiPolygonZM, len(geoms))
			for i, g := range geoms {
				polygons[i], ok = g.(geom.PolygonZM)
				if !ok {
					return nil, fmt.Errorf("Unexpected child geometry type in MultiPolygonZM: %s", reflect.TypeOf(g).String())
				}
			}
			return polygons, nil
		}

		polygons := make(geom.MultiPolygonZ, len(geoms))
		for i, g := range geoms {
			polygons[i], ok = g.(geom.PolygonZ)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in MultiPolygonZ: %s", reflect.TypeOf(g).String())
			}
		}
		return polygons, nil

	} else if dim.HasM() {
		polygons := make(geom.MultiPolygonM, len(geoms))
		for i, g := range geoms {
			polygons[i], ok = g.(geom.PolygonM)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in MultiPolygonM: %s", reflect.TypeOf(g).String())
			}
		}
		return polygons, nil
	}

	polygons := make(geom.MultiPolygon, len(geoms))
	for i, g := range geoms {
		polygons[i], ok = g.(geom.Polygon)
		if !ok {
			return nil, fmt.Errorf("Unexpected child geometry type in MultiPolygon: %s", reflect.TypeOf(g).String())
		}
	}
	return polygons, nil
}
func readWkbGeometryCollection(r io.Reader, byteOrder binary.ByteOrder, dim Dimensioner) (geom.Geometry, error) {

	var numGeoms uint32
	if err := binary.Read(r, byteOrder, &numGeoms); err != nil {
		return nil, err
	}

	geoms := make([]geom.Geometry, numGeoms)

	for i := uint32(0); i < numGeoms; i++ {
		geom, err := Read(r)
		if err != nil {
			return nil, err
		}
		geoms[i] = geom
	}

	var ok bool
	if dim.HasZ() {
		if dim.HasM() {
			collection := make(geom.GeometryCollectionZM, len(geoms))
			for i, g := range geoms {
				collection[i], ok = g.(geom.GeometryZM)
				if !ok {
					return nil, fmt.Errorf("Unexpected child geometry type in GeometryCollectionZM: %s", reflect.TypeOf(g).String())
				}
			}
			return collection, nil
		}

		collection := make(geom.GeometryCollectionZ, len(geoms))
		for i, g := range geoms {
			collection[i], ok = g.(geom.GeometryZ)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in GeometryCollectionZ: %s", reflect.TypeOf(g).String())
			}
		}
		return collection, nil

	} else if dim.HasM() {
		collection := make(geom.GeometryCollectionM, len(geoms))
		for i, g := range geoms {
			collection[i], ok = g.(geom.GeometryM)
			if !ok {
				return nil, fmt.Errorf("Unexpected child geometry type in GeometryCollectionM: %s", reflect.TypeOf(g).String())
			}
		}
		return collection, nil
	}

	collection := make(geom.GeometryCollection, len(geoms))
	for i, g := range geoms {
		collection[i], ok = g.(geom.Geometry)
		if !ok {
			return nil, fmt.Errorf("Unexpected child geometry type in GeometryCollection: %s", reflect.TypeOf(g).String())
		}
	}
	return collection, nil
}
