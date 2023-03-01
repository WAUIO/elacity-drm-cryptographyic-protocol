package internal

import (
	"math/rand"
	"time"
)

func RandomSrc() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

/**
 * GenerateRandomNumber generates a random number.
 * @return {uint64} random number
 */
func RandomNum() uint64 {
	return RandomSrc().Uint64()
}
