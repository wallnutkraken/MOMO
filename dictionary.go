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

// NewDictionary constructs a new Dictionary interface and returns it. Completely empty.
func NewDictionary() Dictionary {
	d := new(dictionary)
	d.content = make(map[string]Word)
	d.keys = make([]string, 0)

	return d
}

// Dictionary is the way one interfaces with the underlying text generation. This includes feeding text and also
// generating text from it.
type Dictionary interface {
	addWord(Word)
	// Feed will take a string that's your text, and process it into the dictionary for further use with Generate()
	Feed(string)
	// Generate uses the data collected in the dictionary to create a pseudorandom sentence based on the length.
	Generate(int) string
}
