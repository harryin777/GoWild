package utils

import (
	"math/rand"
	"time"
)

// min > 0 max > 0
func GetRandomNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
