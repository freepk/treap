package treap

import (
	"math/rand"
	"testing"
)

func TestTreap(t *testing.T) {
	x := newTreap(4, 6, nil, nil)
	x = x.add(6, 2)
	x = x.add(13, 8)
	x = x.add(7, 10)
	x = x.add(11, 3)

	if x.x != 7 {
		t.Fail()
	}
}

func BenchmarkTreap(b *testing.B) {
	x := newTreap(0, rand.Uint32(), nil, nil)
	for i := 1; i <= b.N; i++ {
		x = x.add(uint32(i), rand.Uint32())
	}
}

