package main

import (
	"fmt"
)

// nはパッケージ変数、mainパッケージの中であればどこからでも参照できる
var n = 100

func main() {
	// 変数
	// 明示的な定義
	var n int
	var x, y, z int
	var (
		x, y int
		name string
	)
	n = 5
	x, y = 1, 2

	// 暗黙的な定義
	i := 1 // 型推論
	b := true
	f := 3.14
	s := "abc"
}

func one() int {
	return 1
}

n := one


// varと暗黙的な定義
var (
	n = 1
	s = "string"
	b = true
)

n := 1
s := "string"
b := true