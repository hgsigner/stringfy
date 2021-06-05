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

type lineWidther interface {
	setLineWidth(int)
}

// OmissionOption function that accepts interface omissioner
type OmissionOption func(omissioner)

// SeparatorOption is a function that accepts interface separatorer
type SeparatorOption func(separatorer)

// RadiusOption is a function that accepts interface radiuser
type RadiusOption func(radiuser)

// LengthOption is a function that accepts interface lengther
type LengthOption func(lengther)

// LineWidthOption is a function that accepts interface lineWidther
type LineWidthOption func(lineWidther)

// AddOmission adds a custom omission
func AddOmission(om string) OmissionOption {
	return func(obj omissioner) {
		obj.setOmission(om)
	}
}

// AddSeparator adds a custom separator
func AddSeparator(sep string) SeparatorOption {
	return func(obj separatorer) {
		obj.setSeparator(sep)
	}
}

// AddRadius adds a custom separator
func AddRadius(rad int) RadiusOption {
	return func(obj radiuser) {
		obj.setRadius(rad)
	}
}

// AddLength adds a custom length
func AddLength(l int) LengthOption {
	return func(obj lengther) {
		obj.setLength(l)
	}
}

// AddLineWidth Adds a custom line width
func AddLineWidth(lw int) LineWidthOption {
	return func(obj lineWidther) {
		obj.setLineWidth(lw)
	}
}
