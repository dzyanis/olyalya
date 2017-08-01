package cmd

import (
	"testing"
)

var (
	lex = NewLexer()
)

func TestIsArrayBegin(t *testing.T) {
	if !lex.IsArrayBegin('[') {
		t.Error("Unexpected behavior")
	}
	if lex.IsArrayBegin('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsArrayEnd(t *testing.T) {
	if !lex.IsArrayEnd(']') {
		t.Error("Unexpected behavior")
	}
	if lex.IsArrayEnd('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsHashBegin(t *testing.T) {
	if !lex.IsHashBegin('{') {
		t.Error("Unexpected behavior")
	}
	if lex.IsHashBegin('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsHashEnd(t *testing.T) {
	if !lex.IsHashEnd('}') {
		t.Error("Unexpected behavior")
	}
	if lex.IsHashEnd('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsStringBegin(t *testing.T) {
	if !lex.IsStringBegin('"') {
		t.Error("Unexpected behavior")
	}
	if lex.IsStringBegin('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsStringEnd(t *testing.T) {
	if !lex.IsStringEnd('"') {
		t.Error("Unexpected behavior")
	}
	if lex.IsStringEnd('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsSpace(t *testing.T) {
	if !lex.IsSpace(' ') {
		t.Error("Unexpected behavior")
	}

	if !lex.IsSpace('\t') {
		t.Error("Unexpected behavior")
	}

	if !lex.IsSpace('\n') {
		t.Error("Unexpected behavior")
	}

	if lex.IsSpace('a') {
		t.Error("Unexpected behavior")
	}
}

func TestIsPreviusCharEscape(t *testing.T) {
	lex.termSet("")
	if lex.IsPreviusCharEscape() {
		t.Error("Unexpected behavior")
	}

	lex.termSet("TEST")
	if lex.IsPreviusCharEscape() {
		t.Error("Unexpected behavior")
	}

	lex.termAdd("\\")
	if !lex.IsPreviusCharEscape() {
		t.Error("Unexpected behavior")
	}
}

func TestIsStatusOpen(t *testing.T) {
	lex.Close()
	if lex.IsStatusOpen() {
		t.Error("Unexpected behavior")
	}

	lex.Open(TERM_NAME)
	if !lex.IsStatusOpen() {
		t.Error("Unexpected behavior")
	}

	lex.Close()
	if lex.IsStatusOpen() {
		t.Error("Unexpected behavior")
	}
}

func TestClose(t *testing.T) {
	ls := NewLexer()
	ls.termSet("TEST1")
	ls.Close()
	ls.Close()
	ls.termSet("TEST2")
	ls.Close()
	terms := ls.GetTerms()
	if terms[0] != "TEST1" {
		t.Error("Unexpected behavior")
	}
	if terms[1] != "" {
		t.Error("Unexpected behavior")
	}
	if terms[2] != "TEST2" {
		t.Error("Unexpected behavior")
	}
}

func TestParseCommandName(t *testing.T) {
	ls := NewLexer()
	lexems, err := ls.Parse(` SET/INDEX/ARRAY value `)
	if err != nil {
		t.Error(err)
	}
	if lexems[0] != "SET/INDEX/ARRAY" {
		t.Error("Unexpected behavior")
	}
}

func TestParse(t *testing.T) {
	ls := NewLexer()
	terms, err := ls.Parse(`   TEST hello " World!" 42 ["1", "2", "3"] {"zero":"0", "two":"1"}   `)
	if err != nil {
		t.Error(err)
	}
	if len(terms) < 6 {
		t.Error("Unexpected behavior")
	}
	if terms[0] != "TEST" {
		t.Error("Unexpected behavior")
	}
	if terms[1] != "hello" {
		t.Error("Unexpected behavior")
	}
	if terms[2] != "\" World!\"" {
		t.Error("Unexpected behavior")
	}
	if terms[3] != "42" {
		t.Error("Unexpected behavior")
	}
	if terms[4] != "[\"1\", \"2\", \"3\"]" {
		t.Error("Unexpected behavior")
	}
	if terms[5] != "{\"zero\":\"0\", \"two\":\"1\"}" {
		t.Error("Unexpected behavior")
	}
}
