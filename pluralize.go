package stringfy

import (
	"fmt"
	"regexp"
	"strings"
)

// NewPlural initializes the PluralSet struct
func NewPlural() *PluralSet {
	return &PluralSet{}
}

// AddPlural adds custom plural to the word
func AddPlural(pl string) OptPlural {
	return func(ps *PluralSet) {
		ps.plural = pl
	}
}

// PluralSet struct
type PluralSet struct {
	plural string
}

// OptPlural is a type func(*PluralSet)
type OptPlural func(*PluralSet)

// Options set the options passed to plural
func (ps *PluralSet) Options(options ...OptPlural) {
	for _, opt := range options {
		opt(ps)
	}
}

// Perform the pluralization. It accepts a connt, a singular version of the word
// and, optionaly, the pluralized version of the word.
// E.g pluralize.AddPlural("boatys")
func (ps *PluralSet) Perform(count int, singular string) string {
	singlDown := strings.ToLower(singular)

	if count == 1 {
		return fmt.Sprintf("%d %s", count, singlDown)
	}

	if ps.plural != "" {
		return fmt.Sprintf("%d %s", count, ps.plural)
	}

	pluralWord := pluralizeWord(singlDown)
	return fmt.Sprintf("%d %s", count, pluralWord)
}

func pluralizeWord(w string) string {
	// Checks if it is uncountable
	for _, unc := range uncountableList {
		if w == unc {
			return unc
		}
	}

	// Check if it is in the irregularRules
	for k, v := range irregularRules {
		if w == k {
			return v
		}
	}

	var word string

	// Iterates over pluralRules map and look for the word's plural in it. If it
	// find the plural, it appends to the regexMatchesSlice [][]string.
	regexMatchesSlice := make([][]string, 0)
	for reg, replacer := range pluralRules {
		re := regexp.MustCompile(reg)
		out := re.FindStringSubmatch(w)
		if len(out) > 1 && out[0] != "" {
			out = append(out, reg, replacer)
			regexMatchesSlice = append(regexMatchesSlice, out)
		}
	}

	// Filters if more than one regex has matched
	switch len(regexMatchesSlice) {
	case 1:
		re := regexp.MustCompile(regexMatchesSlice[0][len(regexMatchesSlice[0])-2])
		word = re.ReplaceAllString(w, regexMatchesSlice[0][len(regexMatchesSlice[0])-1])
	case 2:
		for _, v := range regexMatchesSlice {
			if v[len(v)-3] != w {
				re := regexp.MustCompile(v[len(v)-2])
				word = re.ReplaceAllString(w, v[len(v)-1])
			}
		}
	}

	// If word return empty, append the 's'
	if word == "" {
		word = fmt.Sprintf("%ss", w)
	}

	return word
}
