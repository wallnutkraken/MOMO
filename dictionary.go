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
		d.content[theWord.String()].addNexts(theWord.next()...)
		d.content[theWord.String()].addPrevs(theWord.prev()...)
	}
}

func NewDictionary() Dictionary {
	d := new(dictionary)
	d.content = make(map[string]Word)
	d.keys = make([]string, 0)

	return d
}

type Dictionary interface {
	AddWord(Word)
	Feed(string)
	Generate(int) string
}
