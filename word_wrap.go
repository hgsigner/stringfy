package stringfy

import (
	"math"
	"strings"
)

const (
	wordWrapDefaultLineWidth = 80
)

type WordWraper struct {
	lineWidth int
}

func (wwr *WordWraper) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case lineWidthOption:
			opt.(lineWidthOption)(wwr)
		}
	}
}

// Set fields
func (wwr *WordWraper) setLineWidth(lw int) { wwr.lineWidth = lw }

// Performs the wraping of a given text.
func (wwr *WordWraper) Perform(text string) string {

	tl := len(text)

	if tl <= wwr.lineWidth {
		return text
	}

	times := int(math.Ceil(float64(len(text)) / float64(wwr.lineWidth)))
	//fmt.Println("times:", times)

	lineH := make([]string, 0)
	for i := 0; i < times; i++ {

		if text == "" {
			continue
		}

		// Gets n number (wwr.lineWidth) characters from
		// the text, starting from 0.
		// It needs to test if the next character
		// related to the last char from 0 to wwr.lineWidth
		// is a whitspace. If true, the whole slice will
		// transformed into a line. If not, the program will
		// get the closest whitespace that does not exceed the
		// lineWidth and break the line at its index.

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

		//fmt.Println("text:", text)

		textSlice := text[0:to]
		//fmt.Println("textSlice:", textSlice)
		nextChar := string(text[to-1])
		//fmt.Printf("nextChar: (%s)\n", nextChar)

		if nextChar != " " {
			// If its found any last index of a whitespace char
			lastWsIndex := strings.LastIndex(textSlice, " ")
			//fmt.Println("lastWsIndex:", lastWsIndex)
			if lastWsIndex != -1 {
				//fmt.Printf("Line: (%s)\n", text[0:lastWsIndex])
				lineH = append(lineH, text[0:lastWsIndex])
				text = text[lastWsIndex+1:]
			} else {
				// It means that the word it bigger than the lineWidth
				// The whole word will be line and the program will get
				// the first ocurrence of a white space.

				firstWsIndex := strings.Index(text, " ")
				//fmt.Println("firstWsIndex:", firstWsIndex)
				if firstWsIndex < 0 {
					//fmt.Printf("Line: (%s)\n", text)
					lineH = append(lineH, text)
					text = ""
				} else {
					//fmt.Printf("Line: (%s)\n", text[0:firstWsIndex])
					lineH = append(lineH, text[0:firstWsIndex])
					text = text[firstWsIndex+1:]
				}
			}
		} else {
			//fmt.Printf("Line: (%s)\n", textSlice)
			lineH = append(lineH, strings.Trim(textSlice, " "))
			text = text[to:]
		}
		//fmt.Println("-------------")
	}

	lineH = append(lineH, text)
	return strings.Trim(strings.Join(lineH, "\n"), "\n")

}

// Creates a new instance of the WordWraper struct
// if its defaults.
func NewWordWrap() *WordWraper {
	return &WordWraper{wordWrapDefaultLineWidth}
}
