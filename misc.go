package treap

import (
	"fmt"
)

func traverse(n *node) {
	if n.l[left] != nil {
		traverse(n.l[left])
	}
	fmt.Printf("ptr: %p node: %v\n", n, n)
	if n.l[right] != nil {
		traverse(n.l[right])
	}
}

func isValid(n *node) int {
	c := 1
	l := n.l[left]
	if l != nil {
		if l.y > n.y {
			fmt.Printf("Heap violation l.y: %d, n.y: %d\n", l.y, n.y)
			return 0
		}
		if l.x > n.x {
			fmt.Printf("Binary tree violation l.x: %d, n.x: %d\n", l.x, n.x)
			return 0
		}
		v := isValid(l)
		if v == 0 {
			return 0
		}
		c += v
	}
	r := n.l[right]
	if r != nil {
		if r.y > n.y {
			fmt.Printf("Heap violation r.y: %d, n.y: %d\n", r.y, n.y)
			return 0
		}
		if r.x < n.x {
			fmt.Printf("Binary tree violation r.x: %d, n.x: %d\n", r.x, n.x)
			return 0
		}
		v := isValid(r)
		if v == 0 {
			return 0
		}
		c += v
	}
	return c
}
