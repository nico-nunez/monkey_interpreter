package lexer

import (
	"monkey/interpreter/token"
)

type Lexer struct {
	input        string // source code
	position     int    // current index read within input string
	readPosition int    // current reading position (after read char)
	ch           byte   // current char being examined
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		if lex.peakChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tokLiteral := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.EQ, Literal: tokLiteral}
		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '!':
		if lex.peakChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tokLiteral := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: tokLiteral}
		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(lex.ch) {
			tokLiteral := lex.readIdentifer()
			tokType := token.LookupIdent(tok.Literal)
			return token.Token{Type: tokType, Literal: tokLiteral}
		} else if isDigit(lex.ch) {
			tokLiteral := lex.readNumber()
			return token.Token{Type: token.INT, Literal: tokLiteral}
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}

	lex.readChar()
	return tok
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

func (lex *Lexer) peakChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readPosition]
	}
}

func (lex *Lexer) readIdentifer() string {
	startPosition := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[startPosition:lex.position]
}

func (lex *Lexer) readNumber() string {
	startPosition := lex.position
	for isDigit(lex.ch) {
		lex.readChar()
	}
	return lex.input[startPosition:lex.position]
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
