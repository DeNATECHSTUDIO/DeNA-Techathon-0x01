package main

import "fmt"

// interface Talker 定義
type Talker interface {
	Talk()
}

// 上記の interface を満たす構造体 Greeter の定義
type Greeter struct {
	name string
}

// レシーバ(g Greeter) 部分を置くと構造体 Greeter にメソッドを定義できる
// 副作用のあるメソッドの場合は、レシーバの型をポインタ型に(g *Greeter)
func (g Greeter) Talk() {
	fmt.Println("Hello, my name is %s\n", g.name)
}

func main() {
	// 「初期化パラメータを与えてGreeter型の構造体のインスタンスを作成し、そのポインタを代入
	// interface 型の変数宣言
	var talker Talker
	// interface を満たす構造体のポインタに代入
	talker = &Greeter{"wozozo"}
	// Go 言語のコンパイラが、
	// 構造体のインスタンスのポインタを
	// インタフェース型の変数に代入するコードを見つけると、
	// そこで構造体とインタフエースの互換性をチェック
	talker.Talk()
}
