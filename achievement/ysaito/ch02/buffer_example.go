package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	// byte 列に変換した文字列を Write() に渡している
	buffer.Write([]byte("bytes.Buffer example\n"))

	// 書き換え 文字列を受け取れる WriteString() メソッド
	// buffer.WriteString("bytes.Buffer example\n")

	// 書き換え キャスト不要 io.WriteString()
	// io.WriteString("bytes.Buffer example\n")
	fmt.Println(buffer.String())
}
