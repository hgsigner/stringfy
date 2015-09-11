package pluralize

// inflect.plural(/sis$/i, 'ses')
// inflect.plural(/(?:([^f])fe|([lr])f)$/i, '\1\2ves')
// inflect.plural(/(hive)$/i, '\1s')
// inflect.plural(/([^aeiouy]|qu)y$/i, '\1ies')
// inflect.plural(/(x|ch|ss|sh)$/i, '\1es')
// inflect.plural(/(matr|vert|ind)(?:ix|ex)$/i, '\1ices')
// inflect.plural(/^(m|l)ouse$/i, '\1ice')
// inflect.plural(/^(m|l)ice$/i, '\1ice')
// inflect.plural(/^(ox)$/i, '\1en')
// inflect.plural(/^(oxen)$/i, '\1')
// inflect.plural(/(quiz)$/i, '\1zes')

var plural_rules = map[string]string{
	`$`:                "s",
	`s$`:               "s",
	`^(ax|test)is$`:    "es",
	`(octop|vir)us$`:   "i",
	`(octop|vir)i$`:    "i",
	`(alias|status)$`:  "es",
	`(bu)s$`:           "ses",
	`(buffal|tomat)o$`: "oes",
	`([ti])um$`:        "a",
	`([ti])a$`:         "a",
	`(.+)sis$`:         "ses",
}
