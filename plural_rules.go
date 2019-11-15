package stringfy

var pluralRules = map[string]string{
	`$`:                              "s",
	`s$`:                             "s",
	`^(ax|test)is$`:                  "${1}es",
	`(octop|vir)us$`:                 "${1}i",
	`(octop|vir)i$`:                  "${1}i",
	`(alias|status)$`:                "${1}es",
	`(bu)s$`:                         "${1}ses",
	`(buffal|tomat)o$`:               "${1}oes",
	`([ti])um$`:                      "${1}a",
	`([ti])a$`:                       "${1}a",
	`(.+)sis$`:                       "${1}ses",
	`(?:([^f])fe|([lr])f|([aei])f)$`: "$1$2${3}ves",
	`(hive)$`:                        "${1}s",
	`([^aeiouy]|qu)y$`:               "${1}ies",
	`([a-z]{2,}x|ch|ss|sh)$`:         "${1}es",
	`(matr|vert|ind)(?:ix|ex)$`:      "${1}ices",
	`^(m|l)ouse$`:                    "${1}ice",
	`^(m|l)ice$`:                     "${1}ice",
	`^(ox)$`:                         "${1}en",
	`^(oxen)$`:                       "$1",
	`^(quiz)$`:                       "${1}zes",
}

var irregularRules = map[string]string{
	"person": "people",
	"man":    "men",
	"child":  "children",
	"sex":    "sexes",
	"move":   "moves",
	"zombie": "zombies",
}

var uncountableList = []string{
	"equipment",
	"information",
	"rice",
	"money",
	"species",
	"series",
	"fish",
	"sheep",
	"jeans",
	"police",
}
