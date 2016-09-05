package MOMO

import (
	"strings"
)

func (d *dictionary) Generate(maxWords int) string {
	var strs = make([]string, 0)

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

	/* ToUpper first letter of first word */
	if len(strs) > 0 {
		strs[0] = strings.Title(strs[0])
	}

	return strings.Join(strs, " ") + "."
}

func (w *word) pickNext() Word {
	choices := w.next()
	if len(choices) == 0 {
		return nil
	}
	return choices[randomExclusive(len(choices))]
}
