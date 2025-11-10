package lexer

import (
	"github.com/Jihyun3478/logi-lang/token"
)

// 소스 코드를 토큰으로 변환하는 렉서
type Lexer struct {
	input string
	position int     // 입력에서 현재 위치(현재 문자를 가리킴)
	readPosition int // 입력에서 현재 읽는 위치(현재 문자의 다음을 가리킴)
	ch byte          // 현재 조사하고 있는 문자
}

// 주어진 입력으로 새 렉서를 생성하고 첫 문자를 읽음
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 현재 문자를 검사하여 해당하는 토큰을 반환하고 다음 문자로 이동
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// 문자가 알파벳이거나 언더바(_)인지 확인
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 주어진 타입과 문자로 새 토큰을 생성
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 연속된 문자들을 읽어 식별자(변수명, 키워드 등)를 반환
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 다음 문자를 읽고 입력에서 위치를 전진시킨 후 EOF이면 0을 설정
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 백, 탭, 개행 등 whitespace 문자들을 건너뜀
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 연속된 숫자들을 읽어 숫자 리터럴 문자열을 반환
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 문자가 숫자(0-9)인지 확인
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
