package lexer

type Lexer struct {
	input        string // source code
	position     int    // position of most recently read char
	readPosition int    // current reading position (after read char)
	ch           byte   // current char being examined
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}
