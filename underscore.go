package stringfy

import "strings"

//Performs the the underscore function on a string.
func Underscore(s string) string {

	// Trims the string removing extra blank spcaces
	// and splits the lower case version of the trimmed string.
	strim := strings.Trim(s, " ")
	ssplit := strings.Split(strings.ToLower(strim), " ")

	// Tests if the lentgth of the trimmed lowered string
	// is equal to 1. If so, it returns the first item from
	// the slice. If not, it will join the splited lowered case
	// with _ and will return it
	if len(ssplit) == 1 {
		return ssplit[0]
	}
	joined := Escape(strings.Join(ssplit, "_"))
	return joined

}
