package utils

import (
	"math/rand"
	"time"
)

func GenerateSixDigitNumber() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000)
	return otp
}
