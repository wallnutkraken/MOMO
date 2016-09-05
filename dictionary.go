package MOMO

type dictionary struct {
	keys    []string
	content map[string]Word
}

func (d *dictionary) addWord(theWord Word) {
	d.keys = append(d.keys, theWord.String())
	d.content[theWord.String()] = theWord
}

func (d *dictionary) AddWord(theWord Word) {
	if d.content[theWord.String()] == nil {
		d.addWord(theWord)
	} else {
		d.content[theWord.String()].AddNexts(theWord.Next()...)
		d.content[theWord.String()].AddPrevs(theWord.Prev()...)
	}
}

type Dictionary interface {
	AddWord()
	Feed(string)
}
