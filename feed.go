package MOMO

import (
	"strings"
)

func (d *dictionary) Feed(text string) {
	words := strings.Split(strings.ToLower(text), " ")

	if len(words) < 2 {
        a := NewWord(words[0])
		d.AddWord(a)
        d.AddWord(a)
	}

	var previous Word
	for x := 1; x < len(words); x++ {
		currentWord := NewWord(words[x])
		currentWord.AddPrev(previous)
        d.AddWord(currentWord)
		if previous != nil {
			previous.AddNext(currentWord)
		}
		previous = currentWord
	}
}
