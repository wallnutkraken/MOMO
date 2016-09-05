package MOMO

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	/* Initialize and seed the RNG */
	rand.Seed(time.Now().UTC().UnixNano())
	replacer = strings.NewReplacer(".", "", "-", "", ",", "", "@", "", "'", "", "\"", "", "!", "", "?", "", "\n", "")
}
