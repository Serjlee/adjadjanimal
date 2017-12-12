package adjadjanimal

import (
	"bytes"
	"math/rand"
	"time"
	"unicode"

	"github.com/serjlee/adjadjanimal/static"
)

// rng is a thread-safe source of random numbers
// see https://github.com/golang/go/issues/3611
// the global rand source could also be used, but it could be very busy for other concurrent uses
var rng *rand.Rand

func init() {
	rng = rand.New(&syncedSource{source: rand.NewSource(time.Now().UnixNano())})
}

// New creates a new randomized ID that is formatted as an AdjectiveAdjectiveAnimal
func New() string {

	// Pick first adjective
	adj1 := static.Adjectives[rng.Intn(len(static.Adjectives))]
	adj1[0] = capitalizeByte(adj1[0])

	// Pick second
	adj2 := static.Adjectives[rng.Intn(len(static.Adjectives))]
	adj2[0] = capitalizeByte(adj2[0])

	// Pick an animal
	animal := static.Animals[rng.Intn(len(static.Animals))]
	animal[0] = capitalizeByte(animal[0])

	return string(bytes.Join([][]byte{adj1, adj2, animal}, []byte{}))
}

func capitalizeByte(letter byte) byte {
	return byte(unicode.ToUpper(rune(letter)))
}
