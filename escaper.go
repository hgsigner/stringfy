package stringfy

import "regexp"

// Escape special characters from a text.
// Reads through the rules' table and
// mateches if a given character is present
// in the passed string. If so, it will replace it.
func Escape(s string) string {
	text := s
	for k, v := range rules {
		re := regexp.MustCompile(k)
		text = re.ReplaceAllString(text, v)
	}
	return text
}
