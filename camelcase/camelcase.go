package camelcase

import (
	"fmt"
	"regexp"
	"strings"
)

func PerformOn(st string) string {

	reg := regexp.MustCompile("[_-]")
	subWord := reg.ReplaceAllString(st, " ")
	splitWord := strings.Split(subWord, " ")

	casedSilice := make([]string, 0)
	for _, w := range splitWord {
		if w != "" {
			nw := fmt.Sprintf("%s%s", strings.ToUpper(string(w[0])), string(w[1:len(w)]))
			casedSilice = append(casedSilice, nw)
		}
	}
	sj := strings.Join(casedSilice, "")

	return string(
		sj)
}
