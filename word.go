package MOMO

type word struct {
	content string
	prevW   []Word
	nextW   []Word
}

func (w *word) String() string {
	return w.content
}

func (w *word) prev() []Word {
	return w.prevW
}

func (w *word) next() []Word {
	return w.nextW
}

func hasWord(words []Word, word Word) bool {
	for _, wrd := range words {
		if wrd == word {
			return true
		}
	}
	return false
}

func (w *word) addPrev(aWord Word) {
	if !hasWord(w.prevW, aWord) {
		w.prevW = append(w.prevW, aWord)
	}
}

func (w *word) addNext(aWord Word) {
	w.nextW = append(w.nextW, aWord)
}

func (w *word) addNexts(words ...Word) {
	for _, wrd := range words {
		w.addNext(wrd)
	}
}

func (w *word) addPrevs(words ...Word) {
	for _, wrd := range words {
		w.addPrev(wrd)
	}
}

func NewWord(text string) Word {
	w := new(word)
	w.content = text
	w.nextW = make([]Word, 1)
	w.prevW = make([]Word, 0)

	return w
}

type Word interface {
	String() string
	prev() []Word
	addPrev(Word)
	next() []Word
	addNext(Word)
	addNexts(...Word)
	addPrevs(...Word)
	pickNext() Word
}
