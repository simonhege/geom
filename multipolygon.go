// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//MultiPolygon is a collection of two-dimensional geometries representing polygons
type MultiPolygon []Polygon

//MultiPolygonZ is a collection of three-dimensional geometries representing polygons
type MultiPolygonZ []PolygonZ

//MultiPolygonM is a collection of two-dimensional geometries representing polygons, with an additional value defined on each vertex
type MultiPolygonM []PolygonM

//MultiPolygonZM is a collection of three-dimensional geometries representing polygons, with an additional value defined on each vertex
type MultiPolygonZM []PolygonZM

//Envelope returns an envelope around the GeometryCollection
func (c MultiPolygon) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiPolygonZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (c MultiPolygonZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiPolygonM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (c MultiPolygonM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiPolygonZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (c MultiPolygonZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (c MultiPolygonZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//EnvelopeZM returns an envelope around the GeometryCollection
func (c MultiPolygonZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, g := range c {
		e.Extend(g.EnvelopeZM())
	}
	return e
}
