package stringfy

const (
	defaultExcerptRadious = 100
)

type Excerpter struct {
	radious   int
	separator string
	omission  string
}

func (ex *Excerpter) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case radiousOption:
			opt.(radiousOption)(ex)
		case omissionOption:
			opt.(omissionOption)(ex)
		case separatorOption:
			opt.(separatorOption)(ex)
		}
	}
}

// Set fields
func (ex *Excerpter) setRadious(r int)        { ex.radious = r }
func (ex *Excerpter) setOmission(om string)   { ex.omission = om }
func (ex *Excerpter) setSeparator(sep string) { ex.separator = sep }

// Performs excerpt of a given text.
// Receices a text and a phrase.
func (ex *Excerpter) Perform(text, phrase string) string {
	return ""
}

// Creates a new instace of the Excerpter struct
// if its defaults.
func NewExcerpt() *Excerpter {
	return &Excerpter{defaultExcerptRadious, "", "..."}
}
