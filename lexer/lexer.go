// lexer/lexer.go
package lexer

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
