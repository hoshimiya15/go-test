package main

func main() {
	// interface{}とnil
	// すべての型と互換
	// var x interface{}
	// x = 1
	// x = 3.14
	// x = '山'
	// x = "文字列"
	// x = [...]uint8{1, 2, 3, 4, 5}
	// fmt.Printf("%v\n", x)

	// 演算の対象としては利用できない
	var x, y interface{}
	x, y = 1, 2
	z := x + y
}
