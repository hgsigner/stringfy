package camelcase

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hgsigner/stringfy/escaper"
)

// It performs the camelcase on a given word
func PerformOn(st string) string {

	lstring := strings.ToLower(st)

	reg := regexp.MustCompile("[_-]")
	subWord := reg.ReplaceAllString(lstring, " ")
	splitWord := strings.Split(subWord, " ")

	casedSlice := make([]string, 0)
	for _, w := range splitWord {
		if w != "" {
			rword := []rune(w)
			nw := fmt.Sprintf("%s%s", strings.ToUpper(string(rword[0])), string(rword[1:len(rword)]))
			enw := escaper.PerformOn(string(nw))
			casedSlice = append(casedSlice, enw)
		}
	}
	sj := strings.Join(casedSlice, "")

	return string(sj)
}
