package main

import "fmt"

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func intergers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// // 戻り値の破棄
	// q, _ := div(19, 7)

	// // 関数とエラー処理
	// result, err := doSomething()
	// if (err != nil) {
	// 	// エラー処理
	// }

	// // クロージャとしての無名関数
	// f := later()

	// fmt.Println(f("Golang"))
	// fmt.Println(f("is"))
	// fmt.Println(f("awesome!"))

	// クロージャによるジェネレータの実装
	ints := intergers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := intergers()
	fmt.Println(otherInts())
}

// // 戻り値を表す変数
// func doSomething() (x, y int) {
// 	y = 5
// 	return // x == 0, y == 5
// }
