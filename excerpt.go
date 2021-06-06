package stringfy

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	defaultExcerptRadius = 100
)

// NewExcerpt creates a new instace of the Excerpter struct with its defaults.
func NewExcerpt() *Excerpter {
	return &Excerpter{
		radius:   defaultExcerptRadius,
		omission: "...",
	}
}

// Excerpter struct
type Excerpter struct {
	radius    int
	separator string
	omission  string
}

// Options sets RadiusOption, OmissionOption and SeparatorOption
func (ex *Excerpter) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case RadiusOption:
			opt.(RadiusOption)(ex)
		case OmissionOption:
			opt.(OmissionOption)(ex)
		case SeparatorOption:
			opt.(SeparatorOption)(ex)
		}
	}
}

// Set fields
func (ex *Excerpter) setRadius(r int) {
	ex.radius = r
}

func (ex *Excerpter) setOmission(om string) {
	ex.omission = om
}

func (ex *Excerpter) setSeparator(sep string) {
	ex.separator = sep
}

// Test text boundaries in order to add the omission
func (ex *Excerpter) addOmissionToText(text string, max int, textBnd []int) string {
	subText := text

	if textBnd[0] > 0 {
		subText = fmt.Sprintf("%s%s", ex.omission, subText)
	}

	if textBnd[1] != max {
		subText = fmt.Sprintf("%s%s", subText, ex.omission)
	}

	return subText
}

func (ex *Excerpter) emptySpaceSeparator(text, phrase string) (string, error) {
	// Removes all the returns from the text, splits it by by the separator and
	// retrieves the number of words in the text.
	rawText := strings.Replace(text, "\n", "", -1)
	sText := strings.Split(rawText, ex.separator)
	sTextLen := len(sText)

	// Checks if the text contains the phrase and setup the range of words based
	// on the radius.
	rangeSlice := make([]int, 0)
	for i, word := range sText {
		if word != phrase {
			continue
		}

		leftRange := i - ex.radius
		if leftRange < 0 {
			leftRange = 0
		}

		rightRange := (i + ex.radius) + 1
		if rightRange >= sTextLen {
			rightRange = sTextLen
		}

		rangeSlice = append(rangeSlice, leftRange, rightRange)

		break
	}

	// Tests if phrase was found in the text.
	if len(rangeSlice) > 0 {
		radText := strings.Join(sText[rangeSlice[0]:rangeSlice[1]], " ")
		return ex.addOmissionToText(
			radText,
			sTextLen,
			[]int{rangeSlice[0], rangeSlice[1]},
		), nil
	}

	return "", fmt.Errorf("phrase (%s) not found", phrase)
}

func (ex *Excerpter) notEmptySpaceSeparator(text, phrase string) (string, error) {
	tl := len(text)
	ol := len(ex.omission)

	// Looks for the first occurrence of 'phrase' in the text and return its
	// index boundaries
	reg := regexp.MustCompile(phrase)
	pIndex := reg.FindAllStringIndex(text, 1)
	if len(pIndex) == 0 {
		return "", fmt.Errorf("phrase (%s) not found", phrase)
	}

	// Calculates the left radios index. If its less of equal to the length of
	// the omission, the left radios index will equal to 0.
	leftRad := pIndex[0][0] - ex.radius
	if leftRad <= ol {
		leftRad = 0
	}

	// Calculates the right radios index. If its greater than the subtraction of
	// the text length by the omission length, the radius index will equal to 0.
	rightRad := pIndex[0][1] + ex.radius
	if rightRad > (tl - ol) {
		rightRad = tl
	}

	// Extracts the excerpt based on the left and the right radius and then, adds
	// the omission symbol.
	radText := text[leftRad:rightRad]

	return ex.addOmissionToText(radText, tl, []int{leftRad, rightRad}), nil
}

// Perform excerpt of a given text. Receices a text and a phrase.
func (ex *Excerpter) Perform(text, phrase string) (string, error) {
	if strings.ContainsAny(phrase, " ") {
		return "", errors.New("when composing the excerpt, your phrase should not contain more than one word")
	}

	// Checks if the separator is an empty space, which means that it will compose
	// the excerpt based on words.
	if ex.separator == " " {
		radText, err := ex.emptySpaceSeparator(text, phrase)
		if err != nil {
			return "", err
		}

		return radText, nil
	}

	// The separator is different than empty space.
	radText, err := ex.notEmptySpaceSeparator(text, phrase)
	if err != nil {
		return "", err
	}

	return radText, nil
}
