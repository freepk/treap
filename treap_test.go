package treap

import (
	"math/rand"
	"testing"
)

func TestTreap(t *testing.T) {
	x := newNode(0)
	for i := 0; i < 1000; i++ {
		x, _ = x.insert(rand.Int())
	}
	if isValid(x) == 0 {
		t.Fail()
	}
}

func BenchmarkInsSeq(b *testing.B) {
	x := newNode(0)
	for i := 1; i <= b.N; i++ {
		x, _ = x.insert(i)
	}
}

func BenchmarkInsRnd(b *testing.B) {
	x := newNode(0)
	for i := 1; i <= b.N; i++ {
		x, _ = x.insert(rand.Int())
	}
}
