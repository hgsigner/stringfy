package stringfy

type (
	// OmissionOption function that accepts interface omissioner
	OmissionOption func(omissioner)

	// SeparatorOption is a function that accepts interface separatorer
	SeparatorOption func(separatorer)

	// RadiusOption is a function that accepts interface radiuser
	RadiusOption func(radiuser)

	// LengthOption is a function that accepts interface lengther
	LengthOption func(lengther)

	// LineWidthOption is a function that accepts interface lineWidther
	LineWidthOption func(lineWidther)
)

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

type target int

const (
	targetRadius target = iota
	targetSeparator
)

// Option -
type Option struct {
	Target    target
	Integer   int64
	String    string
	Interface interface{}
}

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

// AddRadius adds a custom radius
func AddRadius(rad int) RadiusOption {
	return func(obj radiuser) {
		obj.setRadius(rad)
	}
}

// AddRadiusV2 adds a custom separator
func AddRadiusV2(rad int) Option {
	return Option{
		Target:    targetRadius,
		Integer:   int64(rad),
		Interface: rad,
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
