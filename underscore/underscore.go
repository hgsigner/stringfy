package underscore

import (
	"strings"

	"github.com/hgsigner/stringfy/escaper"
)

//Performs the the underscore function on a string.
func PerformOn(s string) string {

	strim := strings.Trim(s, " ")
	ssplit := strings.Split(strings.ToLower(strim), " ")

	if len(ssplit) == 1 {
		return ssplit[0]
	}

	joined := escaper.PerformOn(strings.Join(ssplit, "_"))
	return joined

}
