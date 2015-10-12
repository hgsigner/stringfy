package stringfy

import (
	"fmt"
	"strings"
)

const (
	truncDefaultValue = 30
)

type Truncater struct {
	length    int
	omission  string
	separator string
}

type optTruncate func(*Truncater)

func (t *Truncater) Options(options ...optTruncate) {
	for _, opt := range options {
		opt(t)
	}
}

// ApplicationController.helpers.truncate("Once upon a time in a world far far away", length: 17, separator: ' ')
// # => "Once upon a..."

func (t *Truncater) Perform(text string) string {
	if len(text) <= t.length {
		return text
	}

	ol := len(t.omission)

	if t.separator == "" {
		return fmt.Sprintf("%s%s", text[:t.length-ol], t.omission)
	}

	trimT := strings.Trim(text[:t.length], t.separator)
	lIndex := strings.LastIndex(trimT, t.separator)

	if len(trimT) <= len(t.omission) {
		if len(text) <= truncDefaultValue {
			return text
		}

		return fmt.Sprintf("%s%s", text[:truncDefaultValue-ol], t.omission)
	}

	return fmt.Sprintf("%s%s", trimT[:lIndex], t.omission)

}

// Creates a new instace of the Truncater struct
// if its defaults.
func NewTruncate() *Truncater {
	return &Truncater{truncDefaultValue, "...", ""}
}

func AddLength(l int) optTruncate {
	return func(t *Truncater) {
		t.length = l
	}
}

func AddOmission(om string) optTruncate {
	return func(t *Truncater) {
		t.omission = om
	}
}

func AddSeparator(sep string) optTruncate {
	return func(t *Truncater) {
		t.separator = sep
	}
}
