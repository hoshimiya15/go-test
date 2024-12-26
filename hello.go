package main

import "fmt"

func main() {
	// // マップ
	// m := make(map[int]string)
	// m[1] = "US"
	// m[81] = "Japan"
	// m[86] = "China"
	// m[86] = "Korea" // キーの値が重複する代入を行うと要素の値は上書きされる
	// fmt.Println(m)
	
	// // マップのリテラル
	// m := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	// fmt.Println(m)

	// m := map[int]string{
	// 	1: "Taro", 
	// 	2: "Hanako", 
	// 	3: "Jiro",
	// }
	// fmt.Println(m)

	// m := map[int][]int{
	// 	1: []int{1},
	// 	2: []int{1, 2}, 
	// 	3: []int{1, 2, 3},
	// }
	// fmt.Println(m)

	// m := map[int][]int{
	// 	1: {1},
	// 	2: {1, 2}, 
	// 	3: {1, 2, 3},
	// }
	// fmt.Println(m)

	// 要素の参照
	m := map[int]string{
		1: "A",
		2: "B", 
		3: "C",
	}

	// s := m[1]
	// t := m[9]
	// fmt.Println(s, t)

	s, ok := m[1]
	t, ok := m[9]
	_, ok := m[3]
	fmt.Println(s, t)
}