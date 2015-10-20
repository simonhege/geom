// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//MultiPoint is a collection of two-dimensional geometries representing points
type MultiPoint []Point

//MultiPointZ is a collection of three-dimensional geometries representing points
type MultiPointZ []PointZ

//MultiPointM is a collection of two-dimensional geometries representing points, with an additional value defined on each point
type MultiPointM []PointM

//MultiPointZM is a collection of three-dimensional geometries representing points, with an additional value defined on each point
type MultiPointZM []PointZM

//Envelope returns an envelope around the GeometryCollection
func (c MultiPoint) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiPointZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (c MultiPointZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiPointM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (c MultiPointM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//Envelope returns an envelope around the GeometryCollection
func (c MultiPointZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the GeometryCollection
func (c MultiPointZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//EnvelopeM returns an envelope around the GeometryCollection
func (c MultiPointZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//EnvelopeZM returns an envelope around the GeometryCollection
func (c MultiPointZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, g := range c {
		e.Extend(g.EnvelopeZM())
	}
	return e
}
