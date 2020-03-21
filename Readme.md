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
だめだった。色々見ると、go mod initの次にモジュール名を付けるべきらしい。なんだかgithubのURLみたいなの書いてるのが多いのでそうしてみる。
```shell
go mod init https://github.com/hkj/monkey
go: malformed import path "https://github.com/hkj/monkey": double slash

go mod init github.com/hkj/monkey
go: creating new go.mod: module github.com/hkj/monkey
```
URLそのものでは怒られたが、httpsなどを取って、githubから初めると作られた。

満を持してのtest実行
```shell
go test ./lexer
$GOPATH/go.mod exists but should not
```
あら?$GOPATHとやらはいらないのかな?良く分からんが外してみよう。
```shell
unset GOPATH
go test ./lexer
package github.com/hkj/monkey/lexer (test): package monkey/token is not in GOROOT (/usr/local/Cellar/go/1.14/libexec/src/monkey/token)
FAIL	github.com/hkj/monkey/lexer [setup failed]
FAIL
```
正しいのか、まだ分からないけれど、テストが失敗するところまではこれた。でもpackegeで怒られてるな。さてどうしたものやら。