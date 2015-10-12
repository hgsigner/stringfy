package stringfy

import (
	"fmt"
	"regexp"
	"strings"
)

type PluralSet struct {
	isPluralSet bool
	plural      string
}

type optPlural func(*PluralSet)

func (ps *PluralSet) Options(options ...optPlural) {
	for _, opt := range options {
		opt(ps)
	}
}

//It permorfs the pluralization. It accepts a connt, a singular version of the word and, optionaly, the pluralized version of the word. E.g pluralize.AddPlural("boatys")
func (ps *PluralSet) Perform(count int, singular string) string {

	singlDown := strings.ToLower(singular)

	if count == 1 {
		return fmt.Sprintf("%d %s", count, singlDown)
	}

	var plural_word string

	if ps.isPluralSet {
		return fmt.Sprintf("%d %s", count, ps.plural)
	} else {
		plural_word = pluralize_word(singlDown)
	}

	return fmt.Sprintf("%d %s", count, plural_word)
}

// Initializes the PluralSet struct
func NewPlural() *PluralSet {
	return &PluralSet{}
}

// Adds custom plural to the word
func AddPlural(pl string) optPlural {
	return func(ps *PluralSet) {
		ps.isPluralSet = true
		ps.plural = pl
	}
}

func pluralize_word(w string) string {

	var word string

	// Checks if it is uncountable
	for _, unc := range uncountable_list {
		if w == unc {
			return unc
		}
	}

	// Check if it is in the irregular_rules
	for k, v := range irregular_rules {
		if w == k {
			return v
		}
	}

	// Iterates over plural_rules map
	// and look for the word's plural in it.
	// If it find the plural, it appends to the
	// regex_matches_slice [][]string
	regex_matches_slice := make([][]string, 0)
	for reg, replacer := range plural_rules {
		re := regexp.MustCompile(reg)
		out := re.FindStringSubmatch(w)
		if len(out) > 1 && out[0] != "" {
			out = append(out, reg, replacer)
			regex_matches_slice = append(regex_matches_slice, out)
		}
	}

	// Filters if more than one regex has matched
	switch len(regex_matches_slice) {
	case 1:
		re := regexp.MustCompile(regex_matches_slice[0][len(regex_matches_slice[0])-2])
		word = re.ReplaceAllString(w, regex_matches_slice[0][len(regex_matches_slice[0])-1])
	case 2:
		for _, v := range regex_matches_slice {
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
