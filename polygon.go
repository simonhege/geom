// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//Polygon is a two-dimensional geometry representing a polygon
type Polygon []LineString

//PolygonZ is a three-dimensional geometry representing a polygon
type PolygonZ []LineStringZ

//PolygonM is a two-dimensional geometry representing a polygon, with an additional value defined on each vertex
type PolygonM []LineStringM

//PolygonZM is a three-dimensional geometry representing a polygon, with an additional value defined on each vertex
type PolygonZM []LineStringZM

//Envelope returns an envelope around the polygon
func (p Polygon) Envelope() *Envelope {
	e := NewEnvelope()
	for _, ls := range p {
		e.Extend(ls.Envelope())
	}
	return e
}

//Envelope returns an envelope around the polygon
func (p PolygonZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, ls := range p {
		e.Extend(ls.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the polygon
func (p PolygonZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, ls := range p {
		e.Extend(ls.EnvelopeZ())
	}
	return e
}

//Envelope returns an envelope around the polygon
func (p PolygonM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, ls := range p {
		e.Extend(ls.Envelope())
	}
	return e
}

//EnvelopeM returns an envelope around the polygon
func (p PolygonM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, ls := range p {
		e.Extend(ls.EnvelopeM())
	}
	return e
}

//Envelope returns an envelope around the polygon
func (p PolygonZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, ls := range p {
		e.Extend(ls.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the polygon
func (p PolygonZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, ls := range p {
		e.Extend(ls.EnvelopeZ())
	}
	return e
}

//EnvelopeM returns an envelope around the polygon
func (p PolygonZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, ls := range p {
		e.Extend(ls.EnvelopeM())
	}
	return e
}

//EnvelopeZM returns an envelope around the polygon
func (p PolygonZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, ls := range p {
		e.Extend(ls.EnvelopeZM())
	}
	return e
}

//Clone returns a deep copy of the polygon
func (p Polygon) Clone() Geometry {
	return &p
}

//Clone returns a deep copy of the polygon
func (p PolygonZ) Clone() Geometry {
	return &p
}

//Clone returns a deep copy of the polygon
func (p PolygonM) Clone() Geometry {
	return &p
}

//Clone returns a deep copy of the polygon
func (p PolygonZM) Clone() Geometry {
	return &p
}

//Iterate walks over the points (and can modify in situ) the polygon
func (p Polygon) Iterate(f func([]Point) error) error {
	for i := range p {
		if err := p[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the polygon
func (p PolygonZ) Iterate(f func([]Point) error) error {
	for i := range p {
		if err := p[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the polygon
func (p PolygonM) Iterate(f func([]Point) error) error {
	for i := range p {
		if err := p[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the polygon
func (p PolygonZM) Iterate(f func([]Point) error) error {
	for i := range p {
		if err := p[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}
