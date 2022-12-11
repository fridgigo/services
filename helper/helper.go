package helper

import (
	"math/rand"
)

// generate a random number
func RandomNumber() int {
	min := 10000
	max := 99999
	return rand.Intn(max-min) + min
}
