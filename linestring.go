// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//LineString is a two-dimensional geometry representing a multi-vertex line
type LineString []Point

//LineStringZ is a three-dimensional geometry representing a multi-vertex line
type LineStringZ []PointZ

//LineStringM is a two-dimensional geometry representing a multi-vertex line, with an additional value defined on each vertex
type LineStringM []PointM

//LineStringZM is a three-dimensional geometry representing a multi-vertex line, with an additional value defined on each vertex
type LineStringZM []PointZM

//Envelope returns an envelope around the GeometryCollection
func (l LineString) Envelope() *Envelope {
	e := NewEnvelope()
	for _, pt := range l {
		e.ExtendPoint(pt)
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (l LineStringZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, pt := range l {
		e.ExtendPoint(pt.Point)
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (l LineStringZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, pt := range l {
		e.ExtendPoint(pt)
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (l LineStringM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, pt := range l {
		e.ExtendPoint(pt.Point)
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (l LineStringM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, pt := range l {
		e.ExtendPoint(pt)
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (l LineStringZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, pt := range l {
		e.ExtendPoint(pt.PointZ.Point)
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (l LineStringZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, pt := range l {
		e.ExtendPoint(pt.PointZ)
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (l LineStringZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, pt := range l {
		e.ExtendPoint(PointM{Point: pt.Point, M: pt.M})
	}
	return e
}

//EnvelopeZM returns an envelope around the GeometryCollection
func (l LineStringZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, pt := range l {
		e.ExtendPoint(pt)
	}
	return e
}
