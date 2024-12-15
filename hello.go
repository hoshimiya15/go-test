package main

import "fmt"

func main() {
	// 配列型
	// a := [5]int{1, 2, 3, 4, 5}
	// a := [5]int{1, 2, 3}
	// a := [5]int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("%v", a)

	// 要素数の省略
	// a1 := [...]int{1, 2, 3, 4, 5}
	// a1[2] = 0
	// fmt.Printf("%v", a1)

	// 配列型の互換性
	// var (
	// 	a1 [3]int
	// 	a2 [5]int
	// )
	// a1 = a2 // エラー

	// a1の要素を書き換えてもa2の要素には影響がない
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}
	a1 = a2
	fmt.Printf("%v\n", a1)

	a1[0] = 0
	a1[2] = 0
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
}
