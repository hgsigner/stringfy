package stringfy

import (
	"fmt"
	"regexp"
	"strings"
)

// CamelCase formats strings to CamalCase,
// escaping special characters from the text.
func CamelCase(st string) string {
	lstring := strings.ToLower(st)

	// Substitutes any _- character that might be
	// included in the string and splits the
	// the string by blank space.
	reg := regexp.MustCompile("[_-]")
	subWord := reg.ReplaceAllString(lstring, " ")
	splitWord := strings.Split(subWord, " ")

	// Iterates over the splitted string
	// ensuring the the word is not blank.
	// Converts the word to a slice of rune,
	// upercases the first letter and escapes
	// any spcial character. Finaly, it appends the word
	// to casedSlice, joins the casedSlice and returns it.
	casedSlice := make([]string, 0)
	for _, w := range splitWord {
		if w != "" {
			rword := []rune(w)
			nw := fmt.Sprintf("%s%s", strings.ToUpper(string(rword[0])), string(rword[1:]))
			enw := Escape(string(nw))
			casedSlice = append(casedSlice, enw)
		}
	}
	sj := strings.Join(casedSlice, "")

	return string(sj)
}
