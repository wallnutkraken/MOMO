package MOMO

import (
	"strings"
)

func (d *dictionary) Generate(maxWords int) string {
	strs := make([]string, 0)

	startingIndex := RandomExclusive(len(d.keys))
	currentWord := d.content[d.keys[startingIndex]]

	for iterations := 0; iterations < maxWords; iterations++ {
		if currentWord == nil {
			break
		}
		strs = append(strs, currentWord.String())

		/* Decide on next word */
		choices := currentWord.Next()
		currentWord = choices[RandomExclusive(len(choices))]
	}

	return strings.Join(strs, " ") + "."
}
