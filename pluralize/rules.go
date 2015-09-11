package pluralize

var plural_rules = map[string]string{
	`$`:                         "s",
	`s$`:                        "s",
	`^(ax|test)is$`:             "es",
	`(octop|vir)us$`:            "i",
	`(octop|vir)i$`:             "i",
	`(alias|status)$`:           "es",
	`(bu)s$`:                    "ses",
	`(buffal|tomat)o$`:          "oes",
	`([ti])um$`:                 "a",
	`([ti])a$`:                  "a",
	`(.+)sis$`:                  "ses",
	`(?:([^f])fe|([lr])f)|f$`:   "ves",
	`(hive)$`:                   "s",
	`([^aeiouy]|qu)y$`:          "ies",
	`([a-z]{2,}x|ch|ss|sh)$`:    "es",
	`(matr|vert|ind)(?:ix|ex)$`: "ices",
	`^(m|l)ouse$`:               "ice",
	`^(m|l)ice$`:                "ice",
	`^(ox)$`:                    "en",
	`^(oxen)$`:                  "",
	`^(quiz)$`:                  "zes",
}

var irregular_rules = map[string]string{
	"person": "people",
	"man":    "men",
	"child":  "children",
	"sex":    "sexes",
	"move":   "moves",
	"zombie": "zombies",
}

var uncountable_list = []string{"equipment", "information", "rice", "money", "species", "series", "fish", "sheep", "jeans", "police"}
