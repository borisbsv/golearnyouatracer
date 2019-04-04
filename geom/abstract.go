package geom

type Material interface {
	Scatter(in Ray, p, n Vec) (out Ray, attenuation Vec, ok bool)
}
