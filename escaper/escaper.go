package escaper

import "regexp"

//It performs the escaping
func PerformOn(s string) string {
	text := s
	for k, v := range rules {
		re := regexp.MustCompile(k)
		text = re.ReplaceAllString(text, v)
	}

	return text
}
