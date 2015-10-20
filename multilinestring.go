// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//MultiLineString is a collection of two-dimensional geometries representing multi-vertex lines
type MultiLineString []LineString

//MultiLineStringZ is a collection of three-dimensional geometries representing multi-vertex lines
type MultiLineStringZ []LineStringZ

//MultiLineStringM is a collection of two-dimensional geometries representing multi-vertex lines, with an additional value defined on each vertex
type MultiLineStringM []LineStringM

//MultiLineStringZM is a collection of three-dimensional geometries representing multi-vertex lines, with an additional value defined on each vertex
type MultiLineStringZM []LineStringZM

//Envelope returns an envelope around the GeometryCollection
func (c MultiLineString) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiLineStringZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (c MultiLineStringZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiLineStringM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (c MultiLineStringM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiLineStringZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (c MultiLineStringZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (c MultiLineStringZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//EnvelopeZM returns an envelope around the GeometryCollection
func (c MultiLineStringZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, g := range c {
		e.Extend(g.EnvelopeZM())
	}
	return e
}
