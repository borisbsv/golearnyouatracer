package draw

import (
	"github.com/templarrei/golearnyouatracer/geom"
)

type HittableList struct {
	s []Hittable
}

func NewList(args ...Hittable) HittableList {
	return HittableList{s: args}
}

func (l *HittableList) Append(h Hittable) {
	l.s = append(l.s, h)
}

func (l HittableList) Hit(r geom.Ray, tMin, tMax float64) (float64, geom.Vec, geom.Vec, geom.Material) {
	closest := tMax
	var t float64
	var p, n geom.Vec
	var mat geom.Material
	for _, s := range l.s {
		if st, sp, sn, m := s.Hit(r, tMin, closest); st > 0 {
			closest, t = st, st
			p = sp
			n = sn
			mat = m
		}
	}
	return t, p, n, mat
}
