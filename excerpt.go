package stringfy

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	defaultExcerptRadious = 100
)

type Excerpter struct {
	radious   int
	separator string
	omission  string
}

func (ex *Excerpter) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case radiousOption:
			opt.(radiousOption)(ex)
		case omissionOption:
			opt.(omissionOption)(ex)
		case separatorOption:
			opt.(separatorOption)(ex)
		}
	}
}

// Set fields
func (ex *Excerpter) setRadious(r int)        { ex.radious = r }
func (ex *Excerpter) setOmission(om string)   { ex.omission = om }
func (ex *Excerpter) setSeparator(sep string) { ex.separator = sep }

// Performs excerpt of a given text.
// Receices a text and a phrase.
func (ex *Excerpter) Perform(text, phrase string) string {

	tl := len(text)
	ol := len(ex.omission)

	var radText string

	if ex.separator == "" {

		// Looks for the first occurrence of 'phrase'
		// in the text and return its index boundaries
		reg := regexp.MustCompile(phrase)
		pIndex := reg.FindAllStringIndex(text, 1)
		if len(pIndex) == 0 {
			return text
		}

		leftRad := pIndex[0][0] - ex.radious
		if leftRad <= ol {
			leftRad = 0
		}

		rightRad := pIndex[0][1] + ex.radious
		if rightRad > (tl - ol) {
			rightRad = tl
		}

		radText = text[leftRad:rightRad]

		if leftRad > 0 {
			radText = fmt.Sprintf("%s%s", ex.omission, radText)
		}

		if rightRad != tl {
			radText = fmt.Sprintf("%s%s", radText, ex.omission)
		}
	} else {

		rawText := strings.Replace(text, "\n", "", -1)
		sText := strings.Split(rawText, ex.separator)
		sTextLen := len(sText)

		rangeSlice := make([]int, 0)
		for i, word := range sText {
			if word == phrase {
				leftRange := i - ex.radious
				if leftRange < 0 {
					leftRange = 0
				}
				rangeSlice = append(rangeSlice, leftRange)

				rightRange := (i + ex.radious) + 1
				if rightRange >= sTextLen {
					rightRange = sTextLen
				}
				rangeSlice = append(rangeSlice, rightRange)

				break
			}
		}

		if len(rangeSlice) > 0 {
			radText = strings.Join(sText[rangeSlice[0]:rangeSlice[1]], " ")
			if rangeSlice[0] > 0 {
				radText = fmt.Sprintf("%s%s", ex.omission, radText)
			}

			if rangeSlice[1] != sTextLen {
				radText = fmt.Sprintf("%s%s", radText, ex.omission)
			}
		} else {
			radText = text
		}

	}

	return radText
}

// Creates a new instace of the Excerpter struct
// if its defaults.
func NewExcerpt() *Excerpter {
	return &Excerpter{defaultExcerptRadious, "", "..."}
}
