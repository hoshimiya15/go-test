package main

import "fmt"

<<<<<<< HEAD
func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	// // # チャネルの型
	// // 変数chはint型のチャネル
	// var ch chan int
	// // 変数ch1はint型の受信専用チャネル
	// var ch1 <-chan int
	// // 変数chはint型の送信専用チャネル
	// var ch2 chan<- int

	// // # チャネルの生成
	// // 変数chはバッファサイズ0のチャネル
	// ch := make(chan int)
	// // 変数chはバッファサイズ8のチャネル
	// ch := make(chan int, 8)
	// // チャネルに整数5を送信
	// ch <- 5
	// // チャネルから整数値を受信
	// i := <-ch
	
	// # チャネルとゴルーチン
	ch := make(chan int)
	go receiver(ch)
	i := 0
	for i < 1000 {
		ch <- i
		i++
	}
=======
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
>>>>>>> a2d66c95702063ec97498a347a50e2ff87c33105
}