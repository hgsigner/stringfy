package stringfy

// Interfaces

type omissioner interface {
	setOmission(string)
}

type separatorer interface {
	setSeparator(string)
}

type radiuser interface {
	setRadius(int)
}

type lengther interface {
	setLength(int)
}

// Types

type omissionOption func(omissioner)
type separatorOption func(separatorer)
type radiusOption func(radiuser)
type lengthOption func(lengther)

// Adds a custom omission
func AddOmission(om string) omissionOption {
	return func(obj omissioner) {
		obj.setOmission(om)
	}
}

// Adds a custom separator
func AddSeparator(sep string) separatorOption {
	return func(obj separatorer) {
		obj.setSeparator(sep)
	}
}

// Adds a custom separator
func AddRadius(rad int) radiusOption {
	return func(obj radiuser) {
		obj.setRadius(rad)
	}
}

// Adds a custom length
func AddLength(l int) lengthOption {
	return func(obj lengther) {
		obj.setLength(l)
	}
}
