# Monkey言語の作成

* 「Go言語で作るインタプリタ」のコードを写経する
* https://www.oreilly.co.jp/books/9784873118222/
* Go言語はTour Goをほんの少しやっただけ。

## 作業メモ
### go test ./lexerでやったこと
```shell
go test ./lexer

go: cannot find main module, but found .git/config in ~/monkey
	to create a module there, run:
	go mod init
```
テストは動かない。なるほど、go mod initとすれば良いのかと
```shell
go mod init

go: cannot determine module path for source directory ~/monkey (outside GOPATH, module path must be specified)

Example usage:
	'go mod init example.com/m' to initialize a v0 or v1 module
	'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.
```
だめだった。色々見ると、go mod initの次にモジュール名を付けるべきらしい。とりあえず作業フォルダのトップレベルの名前を入れると良いっぽい。
```shell
cd monkey
go mod init monkey
go: creating new go.mod: module monkey
```

満を持してのtest実行
```shell
go test ./lexer
$GOPATH/go.mod exists but should not
```
あら?$GOPATHとやらはいらないのかな?良く分からんが外してみよう。
```shell
unset GOPATH
go test ./lexer
# monkey/lexer [monkey/lexer.test]
lexer/lexer.go:16:9: cannot use 1 (type int) as type *Lexer in return argument
lexer/lexer_test.go:31:11: l.NextToken undefined (type *Lexer has no field or method NextToken)
FAIL	monkey/lexer [build failed]
FAIL
```
正しいのか、まだ分からないけれど、テストが失敗するところまではこれた。