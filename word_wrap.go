package stringfy

const (
	wordWrapDefaultLineWidth = 80
)

type WordWraper struct {
	lineWidth int
}

func (wwr *WordWraper) Options(options ...interface{}) {
	for _, opt := range options {
		switch opt.(type) {
		case lineWidthOption:
			opt.(lineWidthOption)(wwr)
		}
	}
}

// Set fields
func (wwr *WordWraper) setLineWidth(lw int) { wwr.lineWidth = lw }

// Performs the wraping of a given text.
func (t *WordWraper) Perform(text string) string {
	return ""
}

// Creates a new instace of the WordWraper struct
// if its defaults.
func NewWordWrap() *WordWraper {
	return &WordWraper{wordWrapDefaultLineWidth}
}
