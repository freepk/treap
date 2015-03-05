package treap

import (
	"fmt"
)

const (
	stackSize = 64
)

type treap struct {
	x uint32
	y uint32
	l [2]*treap
}

type pair struct {
	t *treap
	d int
}

func newTreap(x, y uint32, l, r *treap) *treap {
	return &treap{x: x, y: y, l: [2]*treap{l, r}}
}

func (t *treap) rotate(d int) *treap {
	var n, z *treap
	var f int

	f = ^d & 1
	n = t.l[d]
	z = n.l[f]
	n.l[f] = t
	t.l[d] = z

	return n
}

func (t *treap) add(x, y uint32) *treap {
	var n, p *treap
	var s []pair
	var i, d, f int
	var r *pair

	n = t
	s = make([]pair, stackSize)
	i = 0
	r = &s[0]
	for n != nil {
		r = &s[i]
		r.t = n
		if n.x > x {
			r.d = 0
			n = n.l[0]
		} else if n.x < x {
			r.d = 1
			n = n.l[1]
		} else {
			return t
		}
		i++
	}
	if n != nil {
		panic("n != nil")
	}
	if i < 1 {
		panic("i < 1")
	}
	n = newTreap(x, y, nil, nil)
	i--
	r = &s[i]
	p = r.t
	d = r.d
	p.l[d] = n
	for (i > 0) && (p.l[d].y > p.y) {
		n = p
		f = d
		i--
		r = &s[i]
		p = r.t
		d = r.d
		p.l[d] = n.rotate(f)
		if p.l[d].x != x {
			panic("p.l[d].x != x")
		}
	}

	if (i == 0) && (p.l[d].y > p.y) {
		if p.l[d].x != x {
			panic("p.l[d].x != x")
		}
		return p.rotate(d)
	}

	return t
}

func traverse(t *treap) {
	if t.l[0] != nil {
		traverse(t.l[0])
	}
	fmt.Printf("ptr: %p treap: %v\n", t, t)
	if t.l[1] != nil {
		traverse(t.l[1])
	}
}
