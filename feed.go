package MOMO

import (
	"strings"
)

func (d *dictionary) Feed(text string) {
	words := strings.Split(strings.ToLower(text), " ")

	if len(words) < 2 {
		d.AddWord(NewWord(words[0]))
	}

	var previous Word
	for x := 1; x < len(words); x++ {
		currentWord := NewWord(words[x])
		currentWord.AddPrev(previous)
		if previous != nil {
			previous.AddNext(currentWord)
		}
		previous = currentWord
	}
}
