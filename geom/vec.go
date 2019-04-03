package geom

import (
	"fmt"
	"io"
	"math"
)

type Vec [3]float64

func NewVec(x, y, z float64) Vec {
	return Vec{x, y, z}
}

func (v Vec) X() float64 {
	return v[0]
}
func (v Vec) Y() float64 {
	return v[1]
}
func (v Vec) Z() float64 {
	return v[2]
}

func (v Vec) R() float64 {
	return v[0]
}
func (v Vec) G() float64 {
	return v[1]
}
func (v Vec) B() float64 {
	return v[2]
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{
		v[0] + v2[0],
		v[1] + v2[1],
		v[2] + v2[2],
	}
}
func (v Vec) Sub(v2 Vec) Vec {
	return Vec{
		v[0] - v2[0],
		v[1] - v2[1],
		v[2] - v2[2],
	}
}
func (v Vec) Mul(v2 Vec) Vec {
	return Vec{
		v[0] * v2[0],
		v[1] * v2[1],
		v[2] * v2[2],
	}
}
func (v Vec) Div(v2 Vec) Vec {
	return Vec{
		v[0] / v2[0],
		v[1] / v2[1],
		v[2] / v2[2],
	}
}

// IStream streams in space-separated vector values from a Reader
func (v Vec) IStream(r io.Reader) error {
	_, err := fmt.Fscan(r, v[0], v[1], v[2])
	return err
}

// OStream writes space-separated vector values to a Writer
func (v Vec) OStream(w io.Writer) error {
	_, err := fmt.Fprint(w, v[0], v[1], v[2])
	return err
}

// Inv returns this vector's inverse as a new vector
func (v Vec) Inv() Vec {
	return Vec{-v[0], -v[1], -v[2]}
}

// Len returns the vector's length
func (v Vec) Len() float64 {
	return math.Sqrt(v.LenSq())
}

// LenSq returns the square of the vector's length
func (v Vec) LenSq() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Scaled returns a vector scaled by a scalar
func (v Vec) Scaled(n float64) Vec {
	return Vec{v[0] * n, v[1] * n, v[2] * n}
}

// Dot returns the dot product of two vectors
func (v Vec) Dot(v2 Vec) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}

// Cross returns the cross product of two vectors
func (v Vec) Cross(v2 Vec) Vec {
	return Vec{
		v[1]*v2[2] - v[2]*v2[1],
		v[2]*v2[0] - v[0]*v2[2],
		v[0]*v2[1] - v[1]*v2[0],
	}
}

// Unit represents a unit vector (length 1)
type Unit Vec

// ToUnit converts this vector to a unit vector
func (v Vec) ToUnit() (u Unit) {
	k := 1.0 / v.Len()
	u[0] = v[0] * k
	u[1] = v[1] * k
	u[2] = v[2] * k
	return
}
