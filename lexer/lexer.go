// Package lexer lexer/lexer.go .
package lexer

import "monkey/token"

// Lexer 解析に使う構造体。文字関連のもが入っている.
type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字を指し示す)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字(ASCIIのみ。UTF-8にするにはbyteではなく rune)
}

// New 関数、入力を覗き見して、現在の文字の次に何が来るか考慮する.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// reachChar ：次の一文字を読んで、input文字列の現在位置を進めること.
func (l *Lexer) readChar() {
	// 入力が終端に到達したかのチェック
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// 終端に来てない場合
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	// l.readPosionは次の位置を示す
	l.readPosition++
}

// newToken: 新しいトークン .
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// NextToken : 次のトークン .
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

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
	}
	l.readChar()
	return tok
}
