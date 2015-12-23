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

//Envelope returns an envelope around the multi-point
func (c MultiPoint) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//Envelope returns an envelope around the multi-point
func (c MultiPointZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the multi-point
func (c MultiPointZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//Envelope returns an envelope around the multi-point
func (c MultiPointM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeM returns an envelope around the multi-point
func (c MultiPointM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//Envelope returns an envelope around the multi-point
func (c MultiPointZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the multi-point
func (c MultiPointZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//EnvelopeM returns an envelope around the multi-point
func (c MultiPointZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//EnvelopeZM returns an envelope around the multi-point
func (c MultiPointZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, g := range c {
		e.Extend(g.EnvelopeZM())
	}
	return e
}

//Clone returns a deep copy of the multi-point
func (c MultiPoint) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the multi-point
func (c MultiPointZ) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the multi-point
func (c MultiPointM) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the multi-point
func (c MultiPointZM) Clone() Geometry {
	return &c
}

//Iterate walks over the points (and can modify in situ) the multi-point
func (c MultiPoint) Iterate(f func([]Point) error) error {
	return f(c)
}

//Iterate walks over the points (and can modify in situ) the multi-point
func (c MultiPointZ) Iterate(f func([]Point) error) error {
	points := make([]Point, len(c))
	for i := range c {
		points[i] = c[i].Point
	}
	err := f(points)
	for i := range c {
		c[i].Point = points[i]
	}
	return err
}

//Iterate walks over the points (and can modify in situ) the multi-point
func (c MultiPointM) Iterate(f func([]Point) error) error {
	points := make([]Point, len(c))
	for i := range c {
		points[i] = c[i].Point
	}
	err := f(points)
	for i := range c {
		c[i].Point = points[i]
	}
	return err
}

//Iterate walks over the points (and can modify in situ) the multi-point
func (c MultiPointZM) Iterate(f func([]Point) error) error {
	points := make([]Point, len(c))
	for i := range c {
		points[i] = c[i].Point
	}
	err := f(points)
	for i := range c {
		c[i].Point = points[i]
	}
	return err
}
