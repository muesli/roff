# roff

[![Latest Release](https://img.shields.io/github/release/muesli/roff.svg?style=for-the-badge)](https://github.com/muesli/roff/releases)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/muesli/roff/build.yml?style=for-the-badge&branch=main)](https://github.com/muesli/roff/actions)
[![Coverage Status](https://img.shields.io/coverallsCoverage/github/muesli/roff.svg?branch=main&style=for-the-badge)](https://coveralls.io/github/muesli/roff?branch=main)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://pkg.go.dev/github.com/muesli/roff)

roff lets you write roff documents in Go

## Tutorial

Import the library:

```go
import "github.com/muesli/roff"
```

Then start a new roff document and write to it:

```go
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
doc.Text("... left-aligned again!")

// lists
doc.Section("Lists")
doc.Text("A list:")
doc.Paragraph()
doc.Indent(4)
doc.List("First list item")
doc.List("Second list item")
```

Fetch the roff document as a string and you're done:

```go
fmt.Println(doc.String())
```

## Feedback

Got some feedback or suggestions? Please open an issue or drop me a note!

* [Twitter](https://twitter.com/mueslix)
* [The Fediverse](https://mastodon.social/@fribbledom)
