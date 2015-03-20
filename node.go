package treap

import (
	"math/rand"
)

const (
	left direction = iota
	right
)

type direction uint8

type node struct {
	x int
	y int
	l [2]*node
}

func newNode(x int) *node {
	return &node{x: x, y: rand.Int(), l: [2]*node{nil, nil}}
}

func (n *node) rotate(d direction) *node {
	f := ^d & 1
	x := n.l[d]
	z := x.l[f]
	x.l[f] = n
	n.l[d] = z
	return x
}

func (n *node) insert(x int) (*node, bool) {
	s := new(stack)
	z := n
	for z != nil {
		switch {
		case z.x > x:
			s.push(z, left)
			z = z.l[left]
		case z.x < x:
			s.push(z, right)
			z = z.l[right]
		default:
			return n, false
		}
	}
	p, pd := s.pop()
	p.l[pd] = newNode(x)
	for p.l[pd].y > p.y {
		switch {
		case p != n:
			g, gd := s.pop()
			g.l[gd] = p.rotate(pd)
			p, pd = g, gd
		default:
			return p.rotate(pd), true
		}
	}
	return n, true
}

func (n *node) search(x int) *node {
	for n != nil {
		switch {
		case n.x > x:
			n = n.l[left]
		case n.x < x:
			n = n.l[right]
		default:
			return n
		}
	}
	return nil
}
