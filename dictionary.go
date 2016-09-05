package MOMO

type dictionary struct {
	keys    []string
	content map[string]Word
}

func (d *dictionary) addWord(theWord Word) {
	if d.content[theWord.String()] == nil {
		d.keys = append(d.keys, theWord.String())
		d.content[theWord.String()] = theWord
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
	addWord(Word)
	Feed(string)
	Generate(int) string
}
