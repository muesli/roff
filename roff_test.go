package roff

import (
	"testing"
	"time"
)

func TestTitleHeading(t *testing.T) {
	doc := NewDocument()
	doc.Heading(1, "Title", "A short description", time.Now())

	if doc.String() != `.TH TITLE 1 "2022-01-11" "Title" "A short description"` {
		t.Error("Expected title heading, got:", doc.String())
	}
}

func TestSectionHeading(t *testing.T) {
	doc := NewDocument()
	doc.Section("Test")

	if doc.String() != "\n.SH TEST\n" {
		t.Error("Expected section heading, got:", []byte(doc.String()))
	}
}

func TestText(t *testing.T) {
	doc := NewDocument()
	doc.Text("Test")

	if doc.String() != "Test" {
		t.Error("Expected text, got:", []byte(doc.String()))
	}
}

func TestTextBold(t *testing.T) {
	doc := NewDocument()
	doc.TextBold("Test")

	if doc.String() != `\fBTest\fP` {
		t.Error("Expected bold text, got:", []byte(doc.String()))
	}
}

func TestTextItalic(t *testing.T) {
	doc := NewDocument()
	doc.TextItalic("Test")

	if doc.String() != `\fITest\fP` {
		t.Error("Expected italic text, got:", []byte(doc.String()))
	}
}

func TestParagraph(t *testing.T) {
	doc := NewDocument()
	doc.Paragraph()

	if doc.String() != "\n.PP\n" {
		t.Error("Expected italic text, got:", []byte(doc.String()))
	}
}

func TestIndentation(t *testing.T) {
	doc := NewDocument()
	doc.Indent(4)
	doc.Text("Test")
	doc.IndentEnd()

	if doc.String() != "\n.RS 4\nTest\n.RE\n" {
		t.Error("Expected indentation, got:", []byte(doc.String()))
	}
}

func TestList(t *testing.T) {
	doc := NewDocument()
	doc.Indent(4)
	doc.List("First")
	doc.List("Second")

	if doc.String() != "\n.RS 4\n.IP \\(bu 3\nFirst\n.IP \\(bu 3\nSecond\n" {
		t.Error("Expected list, got:", []byte(doc.String()))
	}
}
