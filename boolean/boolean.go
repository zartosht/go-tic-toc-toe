package boolean

import (
	"math/rand"
	"time"
)

// Boolgen This variable holds information about randome generator
// src is Source from rand library
// cache hold string returned by Source to use them one at a time and
// reduce need to generate new ones and just using the first char
// remaining holds how many character left in cache
type Boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

// Bool check if bool already generatad, if so, use the first char as random value
// if not create a new one ans store it in Boolgen
func (b *Boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

// New seed and generate new Boolgen for first time
// usage:
// first add `r := boolean.New()``
// then use it like: `r.Bool()`
func New() *Boolgen {
	rand.Seed(time.Now().UnixNano())
	return &Boolgen{src: rand.NewSource(time.Now().UnixNano())}
}
