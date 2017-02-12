// @todo Add Number type
// @todo Fixed a problem \\" for strings
package cmd

import(
	"errors"
	"regexp"
)

var (
	ErrTermIsWrong = errors.New("Term is wrong")
)

const (
	STATUS_NONE  = iota
	STATUS_OPEN  = iota
	STATUS_CLOSE = iota
)

const (
	TERM_NONE   = iota
	TERM_ARRAY  = iota
	TERM_HASH   = iota
	TERM_STRING = iota
	TERM_NAME   = iota
)

type Lexer struct {
	typeTerm int
	status int
	curTerm string
	terms []string
}

func NewLexer() *Lexer {
	return &Lexer{}
}

func (lx *Lexer) termAdd(str string) {
	lx.curTerm = lx.curTerm + str
}

func (lx *Lexer) termSet(str string) {
	lx.curTerm = str
}

func (lx *Lexer) IsStatusOpen() bool {
	return lx.status == STATUS_OPEN
}

func (lx *Lexer) IsArrayBegin(ch byte) bool {
	return ch == '['
}

func (lx *Lexer) IsArrayEnd(ch byte) bool {
	return ch == ']'
}

func (lx *Lexer) IsHashBegin(ch byte) bool {
	return ch == '{'
}

func (lx *Lexer) IsHashEnd(ch byte) bool {
	return ch == '}'
}

func (lx *Lexer) IsStringBegin(ch byte) bool {
	return ch == '"'
}

func (lx *Lexer) IsStringEnd(ch byte) bool {
	return ch == '"'
}

func (lx *Lexer) TermType() int {
	return lx.typeTerm
}

func (lx *Lexer) IsPreviusCharEscape() bool {
	ln := len(lx.curTerm)
	if ln > 1 {
		return lx.curTerm[ln-1] == '\\'
	}
	return false
}

func (lx *Lexer) Open(termType int) {
	lx.status = STATUS_OPEN
	lx.typeTerm = termType
}

func (lx *Lexer) Close() {
	lx.status = STATUS_NONE
	lx.typeTerm = TERM_NONE
	lx.terms = append(lx.terms, lx.curTerm)
	lx.curTerm = ""
}

func (lx *Lexer) IsSpace(b byte) bool {
	r := regexp.MustCompile(`[ \t\r\n\f]`)
	return r.Match([]byte{b})
}

func (lx *Lexer) Parse(cmd string) ([]string, error) {
	for _, symbol := range cmd {
		err := lx.Step(byte(symbol))
		if err!=nil {
			return []string{}, err
		}
	}

	if lx.IsStatusOpen() {
		lx.Close()
	}

	return lx.terms, nil
}

func (lx *Lexer) Step(ch byte) error {
	if !lx.IsStatusOpen() {
		if lx.IsArrayBegin(ch) {
			lx.Open(TERM_ARRAY)
			lx.termSet(`[`)
		} else if lx.IsHashBegin(ch) {
			lx.Open(TERM_HASH)
			lx.termSet(`{`)
		} else if lx.IsStringBegin(ch) {
			lx.Open(TERM_STRING)
			lx.termSet(`"`)
		} else if !lx.IsSpace(ch) {
			lx.Open(TERM_NAME)
			lx.termSet(string(ch))
		} else {
			//return ErrTermIsWrong
		}
		return nil
	}

	if lx.IsStatusOpen() {
		switch lx.TermType() {
		case TERM_ARRAY:
			if lx.IsArrayEnd(ch) {
				lx.termAdd(`]`)
				lx.Close()
			} else {
				lx.termAdd(string(ch))
			}
		case TERM_HASH:
			if lx.IsHashEnd(ch) {
				lx.termAdd(`}`)
				lx.Close()
			} else {
				lx.termAdd(string(ch))
			}
		case TERM_STRING:
			if lx.IsStringEnd(ch) && !lx.IsPreviusCharEscape() {
				lx.termAdd(`"`)
				lx.Close()
			} else {
				lx.termAdd(string(ch))
			}
		case TERM_NAME:
			if lx.IsSpace(ch) {
				lx.Close()
			} else {
				lx.termAdd(string(ch))
			}
		}
	}

	return nil
}

func (lx *Lexer) GetTerms() []string {
	return lx.terms
}