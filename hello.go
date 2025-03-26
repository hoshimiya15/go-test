package main

import (
	"fmt"
	"os"
)

// インターフェース
type Stringer interface {
	String() string
}

// 任意の値を受け取ってString型にして返す、エラーであればMyError型のエラーメッセージ
func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s,nil
	}
	return nil, MyError("CastError")
}

type MyError string

// MyError型にErrorメソッドの実装
func (e MyError) Error() string {
	return string(e)
}

type S string

// S型にStringメソッドの実装
func (s S) String() string {
	return string(s)
}

func main() {
	v := S("hoge")

	// sとerrにToStringer()の返り値が入る
	if s, err := ToStringer(v); err != nil {
		// エラーあり
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}
}	
