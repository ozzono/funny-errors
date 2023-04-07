package utils

import (
	"math/rand"
	"time"
)

// random integer with given max value
func RInt(i int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
}
