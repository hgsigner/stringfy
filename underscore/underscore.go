package underscore

import (
	"regexp"
	"strings"
)

//Performs the the underscore function on a string.
func PerformOn(s string) string {

	strim := strings.Trim(s, " ")
	ssplit := strings.Split(strings.ToLower(strim), " ")

	if len(ssplit) == 1 {
		return ssplit[0]
	}

	joined := strings.Join(ssplit, "_")

	for k, v := range rules {
		re := regexp.MustCompile(k)
		joined = re.ReplaceAllString(joined, v)
	}

	return joined

}
