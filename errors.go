package MOMO

import (
	"errors"
)

func (d *dictionary) ErrNotEnoughWords() error {
	return errors.New("Not enough words in input text")
}
