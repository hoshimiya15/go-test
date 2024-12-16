package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // ポインタpを通してiから値を読みだす
	*p = 21         // ポインタpを通してiへ値を代入する
	fmt.Println(i)  // see the new value of i

	p = &j          // point to j
	*p = *p / 37    // divide j through the pointer
	fmt.Println(j)  // see the new value of j
}