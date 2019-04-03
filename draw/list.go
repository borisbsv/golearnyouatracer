package draw

import "github.com/templarrei/golearnyouatracer/geom"

type HittableList struct {
	s []Hittable
}

func NewList(args ...Hittable) HittableList {
	return HittableList{s: args}
}

func (l HittableList) Hit(r geom.Ray, tMin, tMax float64) (float64, geom.Vec, geom.Vec) {
	closest := tMax
	var t float64
	var p, n geom.Vec
	for _, s := range l.s {
		if st, sp, sn := s.Hit(r, tMin, closest); st > 0 {
			closest, t = st, st
			p = sp
			n = sn
		}
	}
	return t, p, n
}
