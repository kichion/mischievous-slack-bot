package test

import "fmt"

// Equal はassert判定のメソッドを提供します
// 第一引数と第二引数を比較して一致しない場合にエラーメッセージを返却します
func Equal(prop interface{}, expected interface{}, propName string) string {
	if prop == expected {
		return ""
	}
	return fmt.Sprintf("%s equal wrong. expected=%#v and got=%#v", propName, expected, prop)
}
