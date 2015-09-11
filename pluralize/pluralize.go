package pluralize

import (
	"fmt"
	"regexp"
	"strings"
)

type PluralSet struct {
	isPluralSet bool
	plural      string
}

type option func(*PluralSet)

func (ps *PluralSet) Options(options ...option) {
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

//It initializes the PluralSet struct
func New() *PluralSet {
	return &PluralSet{}
}

func AddPlural(pl string) option {
	return func(ps *PluralSet) {
		ps.isPluralSet = true
		ps.plural = pl
	}
}

func pluralize_word(w string) string {

	var word string

	//Checks if it is uncountable
	for _, unc := range uncountable_list {
		if w == unc {
			return unc
		}
	}

	//Check if it is in the irregular_rules
	for k, v := range irregular_rules {
		if w == k {
			return v
		}
	}

	//Iterates through plural_rules
	for k, v := range plural_rules {
		re := regexp.MustCompile(k)
		out := re.FindStringSubmatch(w)
		if len(out) > 1 {
			word = re.ReplaceAllString(w, "$1${2}"+v)
			break
		}
	}

	// If word return empty, append the 's'
	if word == "" {
		word = fmt.Sprintf("%ss", w)
	}

	return word
}
