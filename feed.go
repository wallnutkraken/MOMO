package MOMO

import (
	"strings"
)

func (d *dictionary) Feed(text string) {
	words := strings.Split(sanitizeString(strings.ToLower(text)), " ")

	if len(words) < 2 {
		a := NewWord(words[0])
		d.AddWord(a)
		d.AddWord(a)
	}

	var previous Word
	for x := 1; x < len(words); x++ {
		currentWord := NewWord(words[x])
		currentWord.addPrev(previous)
		d.AddWord(currentWord)
		if previous != nil {
			previous.addNext(currentWord)
		}
		previous = currentWord
	}
}
