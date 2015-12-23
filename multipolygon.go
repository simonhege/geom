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

//Clone returns a deep copy of the multi-polygon
func (c MultiPolygon) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the multi-polygon
func (c MultiPolygonZ) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the multi-polygon
func (c MultiPolygonM) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the multi-polygon
func (c MultiPolygonZM) Clone() Geometry {
	return &c
}

//Iterate walks over the points (and can modify in situ) the multi-polygon
func (c MultiPolygon) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the multi-polygon
func (c MultiPolygonZ) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the multi-polygon
func (c MultiPolygonM) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the multi-polygon
func (c MultiPolygonZM) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}
