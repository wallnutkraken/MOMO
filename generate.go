package MOMO

import (
	"strings"
)

func (d *dictionary) Generate(maxWords int) string {
	strs := make([]string, 0)

	startingIndex := randomExclusive(len(d.keys))
	currentWord := d.content[d.keys[startingIndex]]

	for iterations := 0; iterations < maxWords; iterations++ {
		if currentWord == nil {
			break
		}
		strs = append(strs, currentWord.String())

		/* Decide on next word */
		currentWord = currentWord.pickNext()
	}

	return strings.Join(strs, " ") + "."
}

func (w *word) pickNext() Word {
	choices := w.next()
	var wrd Word
	keepGoing := true

	for keepGoing {
		wrd = choices[randomExclusive(len(choices))]
		if wrd != nil {
			/* Break if we have a word */
			break
		}
		if wrd == nil && !randomBool() {
			/* Break if word is nil and we're unlucky */
			break
		}
	}

	return wrd
}
