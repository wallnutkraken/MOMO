package MOMO

import (
    "time"
    "math/rand"
)

func init() {
    /* Initialize and seed the RNG */
    rand.Seed(time.Now().UTC().UnixNano())
}

func RandomExclusive(maxExclusive int) int {
    return rand.Intn(maxExclusive)
}