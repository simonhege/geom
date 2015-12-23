// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

//GeometryCollection is a collection of two-dimensional geometries
type GeometryCollection []Geometry

//GeometryCollectionZ is a collection of three-dimensional geometries
type GeometryCollectionZ []GeometryZ

//GeometryCollectionM is a collection of two-dimensional geometries, with an additional value defined on each vertex
type GeometryCollectionM []GeometryM

//GeometryCollectionZM is a collection of three-dimensional geometries, with an additional value defined on each vertex
type GeometryCollectionZM []GeometryZM

//Envelope returns an envelope around the GeometryCollection
func (c GeometryCollection) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//Envelope returns an envelope around the geometry collection
func (c GeometryCollectionZ) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the geometry collection
func (c GeometryCollectionZ) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//Envelope returns an envelope around the geometry collection
func (c GeometryCollectionM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeM returns an envelope around the geometry collection
func (c GeometryCollectionM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//Envelope returns an envelope around the geometry collection
func (c GeometryCollectionZM) Envelope() *Envelope {
	e := NewEnvelope()
	for _, g := range c {
		e.Extend(g.Envelope())
	}
	return e
}

//EnvelopeZ returns an envelope around the geometry collection
func (c GeometryCollectionZM) EnvelopeZ() *EnvelopeZ {
	e := NewEnvelopeZ()
	for _, g := range c {
		e.Extend(g.EnvelopeZ())
	}
	return e
}

//EnvelopeM returns an envelope around the geometry collection
func (c GeometryCollectionZM) EnvelopeM() *EnvelopeM {
	e := NewEnvelopeM()
	for _, g := range c {
		e.Extend(g.EnvelopeM())
	}
	return e
}

//EnvelopeZM returns an envelope around the GeometryCollection
func (c GeometryCollectionZM) EnvelopeZM() *EnvelopeZM {
	e := NewEnvelopeZM()
	for _, g := range c {
		e.Extend(g.EnvelopeZM())
	}
	return e
}

//Clone returns a deep copy of the geometry collection
func (c GeometryCollection) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the geometry collection
func (c GeometryCollectionZ) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the geometry collection
func (c GeometryCollectionM) Clone() Geometry {
	return &c
}

//Clone returns a deep copy of the geometry collection
func (c GeometryCollectionZM) Clone() Geometry {
	return &c
}

//Iterate walks over the points (and can modify in situ) the geometry collection
func (c GeometryCollection) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the geometry collection
func (c GeometryCollectionZ) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the geometry collection
func (c GeometryCollectionM) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}

//Iterate walks over the points (and can modify in situ) the geometry collection
func (c GeometryCollectionZM) Iterate(f func([]Point) error) error {
	for i := range c {
		if err := c[i].Iterate(f); err != nil {
			return err
		}
	}
	return nil
}
