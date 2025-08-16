package random

import (
	"math/rand"
	"time"
)

func SixDigitsOTPGenerator() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return 100000 + rng.Intn(900000)
}
