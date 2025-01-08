package main

import "fmt"

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
}