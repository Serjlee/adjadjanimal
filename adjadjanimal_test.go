package adjadjanimal

import (
	"testing"
)

func TestCollisions(t *testing.T) {
	// super un-scientific test to ensure we have a manageable amount of collisions
	if avg := countAverageCollisions(1000000, 25); avg > 150 {
		t.Errorf("Average collisions after %d inserts: %d\n", 1000000, avg)
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New()
	}
}

func BenchmarkNewParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			New()
		}
	})
}

func countAverageCollisions(goal, loops int) int {
	var generated map[string]bool
	collisions := 0
	for j := 0; j < loops; j++ {
		generated = map[string]bool{}
		for i := 0; i < goal; i++ {
			id := New()
			if _, collision := generated[id]; collision {
				collisions += 1
				continue
			}
			generated[id] = true
		}
	}
	return collisions / loops
}
