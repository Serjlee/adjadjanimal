package adjadjanimal

import (
	"math/rand"
	"sync"
)

// syncedSource provide a thread-safe source for random numbers generation
type syncedSource struct {
	lock   sync.Mutex
	source rand.Source
}

func (r *syncedSource) Int63() int64 {
	r.lock.Lock()
	n := r.source.Int63()
	r.lock.Unlock()

	return n
}

func (r *syncedSource) Seed(seed int64) {
	r.lock.Lock()
	r.source.Seed(seed)
	r.lock.Unlock()
}
