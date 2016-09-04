package MOMO

type word struct {
	content string
	prev    []Word
	next    []Word
}

func (w *word) String() string {
	return w.content
}

func (w *word) Prev() []Word {
	return w.prev
}

func (w *word) Next() []Word {
	return w.next
}

func hasWord(words []Word, word Word) bool {
	for _, wrd := range words {
		if wrd == word {
			return true
		}
	}
	return false
}

func (w *word) AddPrev(aWord Word) {
	if !hasWord(w.prev, aWord) {
		w.prev = append(w.prev, aWord)
	}
}

func (w *word) AddNext(aWord Word) {
	w.next = append(w.next, aWord)
}

type Word interface {
	String() string
	Prev() []Word
	AddPrev(Word)
	Next() []Word
	AddNext(Word)
}
