package main

import (
	"fmt"
	"time"

	"github.com/muesli/roff"
)

func main() {
	doc := roff.NewDocument()
	doc.Heading(1, "Title", "A short description", time.Now())

	// a section of text
	doc.Section("Introduction")
	doc.Text("Here is a quick introduction to writing roff documents with Go!")

	// fonts
	doc.Section("Fonts")
	doc.Text("This is a text with a bold font: ")
	doc.TextBold("I am bold!")
	doc.Paragraph()
	doc.Text("This is a text with an italic font: ")
	doc.TextItalic("I am italic!")

	// indentation
	doc.Section("Indentation")
	doc.Text("This block of text is left-aligned to the section.")
	doc.Indent(4)
	doc.Text("This block of text is totally indented.")
	doc.IndentEnd()
	doc.Text("... left-aligned & happy again!")

	// lists
	doc.Section("Lists")
	doc.Text("A list:")
	doc.Paragraph()
	doc.Indent(4)
	doc.List("First list item")
	doc.List("Second list item")

	fmt.Println(doc.String())
}
