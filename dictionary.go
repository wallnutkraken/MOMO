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
	}
}

type Dictionary interface {
	AddWord()
}
