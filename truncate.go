package stringfy

import "fmt"

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

func (t *Truncater) Perform(text string) string {
	if len(text) <= t.length {
		return text
	}

	var final string

	switch t.separator {
	case "":
		ol := len(t.omission)
		tt := text[:t.length-ol]
		final = fmt.Sprintf("%s%s", tt, t.omission)
	}

	return final
}

// Creates a new instace of the Truncater struct
// if its defaults.
func NewTruncate() *Truncater {
	return &Truncater{30, "...", ""}
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
