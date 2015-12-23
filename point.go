// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//Point is a two-dimensional geometry representing a point
type Point struct {
	X float64
	Y float64
}

//PointZ is a three-dimensional geometry representing a point
type PointZ struct {
	Point
	Z float64
}

//PointM is a two-dimensional geometry representing a point, with an additionnal value
type PointM struct {
	Point
	M float64
}

//PointZM is a three-dimensional geometry representing a point, with an additionnal value
type PointZM struct {
	PointZ
	M float64
}

//Envelope returns an envelope around the point
func (pt Point) Envelope() *Envelope {
	return NewEnvelopeFromPoint(pt)
}

//EnvelopeZ returns an envelope around the point
func (pt PointZ) EnvelopeZ() *EnvelopeZ {
	return NewEnvelopeZFromPoint(pt)
}

//EnvelopeM returns an envelope around the point
func (pt PointM) EnvelopeM() *EnvelopeM {
	return NewEnvelopeMFromPoint(pt)
}

//EnvelopeM returns an envelope around the point
func (pt PointZM) EnvelopeM() *EnvelopeM {
	return NewEnvelopeMFromPoint(PointM{Point: pt.PointZ.Point, M: pt.M})
}

//EnvelopeZM returns an envelope around the point
func (pt PointZM) EnvelopeZM() *EnvelopeZM {
	return NewEnvelopeZMFromPoint(pt)
}

//Clone returns a deep copy of the point
func (pt Point) Clone() Geometry {
	return &pt
}

//Iterate walks over the points (and can modify in situ) the point
func (pt *Point) Iterate(f func([]Point) error) error {
	points := []Point{*pt}
	err := f(points)
	*pt = points[0]
	return err
}
