package stringfy

import (
	"math"
	"strings"
)

const (
	wordWrapDefaultLineWidth = 80
)

// NewWordWrap creates a new instance of the WordWraper struct if its defaults.
func NewWordWrap() *WordWraper {
	return &WordWraper{
		lineWidth: wordWrapDefaultLineWidth,
	}
}

// WordWraper struct
type WordWraper struct {
	lineWidth int
}

// Options sets LineWidthOption
func (wwr *WordWraper) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case LineWidthOption:
			opt.(LineWidthOption)(wwr)
		}
	}
}

// Perform the wraping of a given text.
func (wwr *WordWraper) Perform(text string) string {
	tl := len(text)

	// If the length of the text is less than or equal to the lineWidth, the whole
	// text will be transformed into a line.
	if tl <= wwr.lineWidth {
		return text
	}

	// Gets the number of time that the loop will iterate over a text in order to
	// break it into lines.
	times := int(math.Ceil(float64(len(text)) / float64(wwr.lineWidth)))

	lineH := make([]string, 0) // A slice that holds the lines
	for i := 0; i < times; i++ {
		if text == "" {
			continue
		}

		// Gets n number (wwr.lineWidth) of characters from text, starting from 0.
		// It needs to test if the next character related to the last char from 0
		// to wwr.lineWidth is a whitspace. If true, the whole slice will transformed
		// into a line. If not, the program will get the closest whitespace that
		// does not exceed the lineWidth and break the line at its index.
		processedTl := len(text)

		to := wwr.lineWidth
		if to > processedTl {
			to = processedTl
		}

		if processedTl <= wwr.lineWidth {
			lineH = append(lineH, text)
			text = ""
			break
		}

		// Gets the slice from the the text in order to test against the the width
		// that each line must have and for that, it gets the next char after the
		// end of the slice.
		textSlice := text[0:to]
		nextChar := string(text[to-1])

		if nextChar == " " {
			lineH = append(lineH, strings.Trim(textSlice, " "))
			text = text[to:]
			continue
		}

		// If any last index of a whitespace char is found, it means that a slice
		// has stopped within a word. It can't break a line within a word.
		lastWsIndex := strings.LastIndex(textSlice, " ")
		if lastWsIndex != -1 {
			lineH = append(lineH, text[0:lastWsIndex])
			text = text[lastWsIndex+1:]
			continue
		}

		// It means that the word is bigger than the lineWidth. The whole word
		// will be in one line and the program will get the first ocurrence of a
		// white space.
		firstWsIndex := strings.Index(text, " ")
		if firstWsIndex < 0 {
			lineH = append(lineH, text)
			text = ""
			continue
		}

		lineH = append(lineH, text[0:firstWsIndex])
		text = text[firstWsIndex+1:]
	}

	lineH = append(lineH, text)
	return strings.Trim(strings.Join(lineH, "\n"), "\n")
}

// Set fields
func (wwr *WordWraper) setLineWidth(lw int) { wwr.lineWidth = lw }
