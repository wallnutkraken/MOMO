package MOMO

import "math/rand"

func randomExclusive(maxExclusive int) int {
	return rand.Intn(maxExclusive)
}

func randomBool() bool {
	return rand.Float64() > 0.5
}
