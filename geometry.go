// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package geom is a geometry library in Go.

It provides definition of basic geometry structures (Point, LineString, Polygon), including Z, M and ZM variants.
MultiGeometry and Geometry collections are also provided.

*/
package geom

//Geometry represents a geometry object
type Geometry interface {
	Envelope() *Envelope

	Clone() Geometry

	Iterate(f func([]Point) error) error //The iterate function can modify Point in place
}

//GeometryZ represents a three-dimensional geometry object
type GeometryZ interface {
	Geometry

	EnvelopeZ() *EnvelopeZ
}

//GeometryM represents a geometry object, with an additional value defined on each vertex
type GeometryM interface {
	Geometry

	EnvelopeM() *EnvelopeM
}

//GeometryZM represents a three-dimensional geometry object, with an additional value defined on each vertex
type GeometryZM interface {
	GeometryZ

	//Duplication of methods of GeometryM because we can not embed it directly (kind of diamond inheritance problem)
	EnvelopeM() *EnvelopeM

	EnvelopeZM() *EnvelopeZM
}

//Ensure that geometry structs implements the geometry interfaces
var _ Geometry = &Point{}
var _ GeometryZ = &PointZ{}
var _ GeometryM = &PointM{}
var _ GeometryZM = &PointZM{}

var _ Geometry = &LineString{}
var _ GeometryZ = &LineStringZ{}
var _ GeometryM = &LineStringM{}
var _ GeometryZM = &LineStringZM{}

var _ Geometry = &Polygon{}
var _ GeometryZ = &PolygonZ{}
var _ GeometryM = &PolygonM{}
var _ GeometryZM = &PolygonZM{}

var _ Geometry = &MultiPoint{}
var _ GeometryZ = &MultiPointZ{}
var _ GeometryM = &MultiPointM{}
var _ GeometryZM = &MultiPointZM{}

var _ Geometry = &MultiLineString{}
var _ GeometryZ = &MultiLineStringZ{}
var _ GeometryM = &MultiLineStringM{}
var _ GeometryZM = &MultiLineStringZM{}

var _ Geometry = &MultiPolygon{}
var _ GeometryZ = &MultiPolygonZ{}
var _ GeometryM = &MultiPolygonM{}
var _ GeometryZM = &MultiPolygonZM{}

var _ Geometry = &GeometryCollection{}
var _ GeometryZ = &GeometryCollectionZ{}
var _ GeometryM = &GeometryCollectionM{}
var _ GeometryZM = &GeometryCollectionZM{}
