package internal

import (
	"math/rand"
	"time"
)

/**
 * GenerateRandomNumber generates a random number.
 * @return {uint64} random number
 */
func RandomNum() uint64 {
	src := rand.NewSource(time.Now().UnixNano())
	return rand.New(src).Uint64()
}
