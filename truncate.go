package stringfy

import (
	"fmt"
	"strings"
)

const (
	truncDefaultValue = 30
)

type Truncater struct {
	length    int
	omission  string
	separator string
}

func (t *Truncater) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case lengthOption:
			opt.(lengthOption)(t)
		case omissionOption:
			opt.(omissionOption)(t)
		case separatorOption:
			opt.(separatorOption)(t)
		}
	}
}

// Set fields
func (t *Truncater) setLength(l int)         { t.length = l }
func (t *Truncater) setOmission(om string)   { t.omission = om }
func (t *Truncater) setSeparator(sep string) { t.separator = sep }

// Performs the truncation of a given text.
func (t *Truncater) Perform(text string) string {
	// Returns the full text if its size is less
	// or equal to the length passed.
	if len(text) <= t.length {
		return text
	}

	ol := len(t.omission)

	// If there is no separator, it will perform the truncation
	// by getting the text truncated and appending the
	// the omission passed.
	if t.separator == "" {
		return fmt.Sprintf("%s%s", text[:t.length-ol], t.omission)
	}

	// If the user passes a separator e.g. (" "), the
	// function will trim the truncated text and gets the
	// the index of the last occurrence of the separator
	// in the truncated text.
	trimT := strings.Trim(text[:t.length], t.separator)
	lIndex := strings.LastIndex(trimT, t.separator)

	// If the trimT length is less than or equal to
	// the length of the omission passed, it will
	// return the text truncated by the default value (30).
	// Before returning, it tests if the length of the raw text
	// is not less than or equal to the default value (30). If so,
	// returns the raw text.
	if len(trimT) <= len(t.omission) {
		if len(text) <= truncDefaultValue {
			return text
		}
		return fmt.Sprintf("%s%s", text[:truncDefaultValue-ol], t.omission)
	}

	return fmt.Sprintf("%s%s", trimT[:lIndex], t.omission)

}

// Creates a new instance of the Truncater struct
// if its defaults.
func NewTruncate() *Truncater {
	return &Truncater{truncDefaultValue, "...", ""}
}
