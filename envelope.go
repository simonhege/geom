// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

//Envelope represents a two-dimensional rectangular region
type Envelope struct {
	Min, Max Point
}

//NewEnvelope returns a new empty envelope
func NewEnvelope() *Envelope {
	return &Envelope{
		Min: Point{X: math.Inf(1), Y: math.Inf(1)},
		Max: Point{X: math.Inf(-1), Y: math.Inf(-1)},
	}
}

//NewEnvelopeFromPoint returns a new envelope containing a single point
func NewEnvelopeFromPoint(pt Point) *Envelope {
	return &Envelope{
		Min: pt,
		Max: pt,
	}
}

//ExtendPoint extends the envelope with the givent point and returns the modified envelope
func (e *Envelope) ExtendPoint(pt Point) *Envelope {
	e.Min.X = math.Min(e.Min.X, pt.X)
	e.Min.Y = math.Min(e.Min.Y, pt.Y)
	e.Max.X = math.Max(e.Max.X, pt.X)
	e.Max.Y = math.Max(e.Max.Y, pt.Y)
	return e
}

//Extend extends the envelope with the givent envelope and returns the modified envelope
func (e *Envelope) Extend(other *Envelope) *Envelope {
	return e.ExtendPoint(other.Min).ExtendPoint(other.Max)
}

//EnvelopeZ represents a three-dimensional rectangular region
type EnvelopeZ struct {
	Min, Max PointZ
}

//NewEnvelopeZ returns a new empty envelope
func NewEnvelopeZ() *EnvelopeZ {
	return &EnvelopeZ{
		Min: PointZ{Point: Point{X: math.Inf(1), Y: math.Inf(1)}, Z: math.Inf(1)},
		Max: PointZ{Point: Point{X: math.Inf(-1), Y: math.Inf(-1)}, Z: math.Inf(-1)},
	}
}

//NewEnvelopeZFromPoint returns a new envelope containing a single point
func NewEnvelopeZFromPoint(pt PointZ) *EnvelopeZ {
	return &EnvelopeZ{
		Min: pt,
		Max: pt,
	}
}

//ExtendPoint extends the envelope with the givent point and returns the modified envelope
func (e *EnvelopeZ) ExtendPoint(pt PointZ) *EnvelopeZ {
	e.Min.X = math.Min(e.Min.X, pt.X)
	e.Min.Y = math.Min(e.Min.Y, pt.Y)
	e.Min.Z = math.Min(e.Min.Z, pt.Z)
	e.Max.X = math.Max(e.Max.X, pt.X)
	e.Max.Y = math.Max(e.Max.Y, pt.Y)
	e.Max.Z = math.Max(e.Max.Z, pt.Z)
	return e
}

//Extend extends the envelope with the givent envelope and returns the modified envelope
func (e *EnvelopeZ) Extend(other *EnvelopeZ) *EnvelopeZ {
	return e.ExtendPoint(other.Min).ExtendPoint(other.Max)
}

//EnvelopeM represents a two-dimensional rectangular region, with M values
type EnvelopeM struct {
	Min, Max PointM
}

//NewEnvelopeM returns a new empty envelope
func NewEnvelopeM() *EnvelopeM {
	return &EnvelopeM{
		Min: PointM{Point: Point{X: math.Inf(1), Y: math.Inf(1)}, M: math.Inf(1)},
		Max: PointM{Point: Point{X: math.Inf(-1), Y: math.Inf(-1)}, M: math.Inf(-1)},
	}
}

//NewEnvelopeMFromPoint returns a new envelope containing a single point
func NewEnvelopeMFromPoint(pt PointM) *EnvelopeM {
	return &EnvelopeM{
		Min: pt,
		Max: pt,
	}
}

//ExtendPoint extends the envelope with the givent point and returns the modified envelope
func (e *EnvelopeM) ExtendPoint(pt PointM) *EnvelopeM {
	e.Min.X = math.Min(e.Min.X, pt.X)
	e.Min.Y = math.Min(e.Min.Y, pt.Y)
	e.Min.M = math.Min(e.Min.M, pt.M)
	e.Max.X = math.Max(e.Max.X, pt.X)
	e.Max.Y = math.Max(e.Max.Y, pt.Y)
	e.Max.M = math.Max(e.Max.M, pt.M)
	return e
}

//Extend extends the envelope with the givent envelope and returns the modified envelope
func (e *EnvelopeM) Extend(other *EnvelopeM) *EnvelopeM {
	return e.ExtendPoint(other.Min).ExtendPoint(other.Max)
}

//EnvelopeZM represents a three-dimensional rectangular region, with M values
type EnvelopeZM struct {
	Min, Max PointZM
}

//NewEnvelopeZM returns a new empty envelope
func NewEnvelopeZM() *EnvelopeZM {
	return &EnvelopeZM{
		Min: PointZM{PointZ: PointZ{Point: Point{X: math.Inf(1), Y: math.Inf(1)}, Z: math.Inf(1)}, M: math.Inf(1)},
		Max: PointZM{PointZ: PointZ{Point: Point{X: math.Inf(-1), Y: math.Inf(-1)}, Z: math.Inf(-1)}, M: math.Inf(-1)},
	}
}

//NewEnvelopeZMFromPoint returns a new envelope containing a single point
func NewEnvelopeZMFromPoint(pt PointZM) *EnvelopeZM {
	return &EnvelopeZM{
		Min: pt,
		Max: pt,
	}
}

//ExtendPoint extends the envelope with the givent point and returns the modified envelope
func (e *EnvelopeZM) ExtendPoint(pt PointZM) *EnvelopeZM {
	e.Min.X = math.Min(e.Min.X, pt.X)
	e.Min.Y = math.Min(e.Min.Y, pt.Y)
	e.Min.Z = math.Min(e.Min.Z, pt.Z)
	e.Min.M = math.Min(e.Min.M, pt.M)
	e.Max.X = math.Max(e.Max.X, pt.X)
	e.Max.Y = math.Max(e.Max.Y, pt.Y)
	e.Max.Z = math.Max(e.Max.Z, pt.Z)
	e.Max.M = math.Max(e.Max.M, pt.M)
	return e
}

//Extend extends the envelope with the givent envelope and returns the modified envelope
func (e *EnvelopeZM) Extend(other *EnvelopeZM) *EnvelopeZM {
	return e.ExtendPoint(other.Min).ExtendPoint(other.Max)
}
