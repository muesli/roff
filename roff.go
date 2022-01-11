package roff

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const (
	// Title heading (Document structure macro)
	TitleHeading = `.TH %[1]s %[2]d "%[4]s" "%[3]s" "%[5]s"`
	// Paragraph macro
	Paragraph = "\n.PP"
	// Relative-indent start (Document structure macro)
	Indent = "\n.RS"
	// Relative-indent end (Document structure macro)
	IndentEnd = "\n.RE"
	// Indented paragraph
	IndentedParagraph = "\n.IP"
	// Section heading (Document structure macro)
	SectionHeading = "\n.SH %s"
	// Tagged paragraph
	TaggedParagraph = "\n.TP"

	// Bold escape
	Bold = `\fB`
	// Italic escape
	Italic = `\fI`
	// Return to previous font setting
	PreviousFont = `\fP`
)

// Document is a roff document.
type Document struct {
	buffer bytes.Buffer
}

// NewDocument returns a new roff Document.
func NewDocument() *Document {
	return &Document{}
}

// write writes the given text to the internal buffer. Following the roff docs,
// we prevent empty lines in its ouput, as that may mysteriously break some roff
// renderers.
func (tr *Document) write(format string, args ...interface{}) {
	if bytes.HasSuffix(tr.buffer.Bytes(), []byte("\n")) &&
		strings.HasPrefix(format, "\n") {
		// prevent empty lines in output
		format = strings.TrimPrefix(format, "\n")
	}

	fmt.Fprintf(&tr.buffer, format, args...)
}

func (tr *Document) writeln(format string, args ...interface{}) {
	tr.write(format+"\n", args...)
}

// Heading writes the title heading of the document.
func (tr *Document) Heading(section uint, title, description string, ts time.Time) {
	tr.write(TitleHeading, strings.ToUpper(title), section, title, ts.Format("2006-01-02"), description)
}

// Paragraph starts a new paragraph.
func (tr *Document) Paragraph() {
	tr.writeln(Paragraph)
}

// Indent increases the indentation level.
func (tr *Document) Indent(n int) {
	if n >= 0 {
		tr.writeln(Indent+" %d", n)
	} else {
		tr.writeln(Indent)
	}
}

// IndentEnd decreases the indentation level.
func (tr *Document) IndentEnd() {
	tr.writeln(IndentEnd)
}

// TaggedParagraph starts a new tagged paragraph.
func (tr *Document) TaggedParagraph(indentation int) {
	if indentation >= 0 {
		tr.writeln(TaggedParagraph+" %d", indentation)
	} else {
		tr.writeln(TaggedParagraph)
	}
}

// List writes a list item.
func (tr *Document) List(text string) {
	tr.writeln(IndentedParagraph+" \\(bu 3\n%s", escapeText(strings.TrimSpace(text)))
}

// Section writes a section heading.
func (tr *Document) Section(text string) {
	tr.writeln(SectionHeading, strings.ToUpper(text))
}

// EndSection ends the current section.
func (tr *Document) EndSection() {
	tr.writeln("")
}

// Text writes text.
func (tr *Document) Text(text string) {
	inList := false
	for i, s := range strings.Split(text, "\n") {
		if i > 0 && !inList {
			// start a new paragraph if we're not in a list
			tr.Paragraph()
		}

		if strings.HasPrefix(s, "*") {
			// list item
			if !inList {
				// start a new indented list if we're not in one
				tr.Indent(-1)
				inList = true
			}

			tr.List(s[1:])
		} else {
			// regular text
			if inList {
				// end the list if we're in one
				tr.IndentEnd()
				inList = false
			}

			tr.write(escapeText(s))
		}
	}
}

// TextBold writes text in bold.
func (tr *Document) TextBold(text string) {
	tr.write(Bold)
	tr.Text(text)
	tr.write(PreviousFont)
}

// TextItalic writes text in italic.
func (tr *Document) TextItalic(text string) {
	tr.write(Italic)
	tr.Text(text)
	tr.write(PreviousFont)
}

// String returns the roff document as a string.
func (tr Document) String() string {
	return tr.buffer.String()
}

func escapeText(s string) string {
	s = strings.ReplaceAll(s, `\`, `\e`)
	s = strings.ReplaceAll(s, ".", "\\&.")
	return s
}
