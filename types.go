package stringfy

// Interfaces

type omissioner interface {
	setOmission(string)
}

type separatorer interface {
	setSeparator(string)
}

type radiouser interface {
	setRadious(int)
}

type lengther interface {
	setLength(int)
}

// Types

type omissionOption func(omissioner)
type separatorOption func(separatorer)
type radiousOption func(radiouser)
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
func AddRadious(rad int) radiousOption {
	return func(obj radiouser) {
		obj.setRadious(rad)
	}
}

// Adds a custom length
func AddLength(l int) lengthOption {
	return func(obj lengther) {
		obj.setLength(l)
	}
}
