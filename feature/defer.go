package feature

import (
	"fmt"
)

//
func DeferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	// defer func() {
	// 	err := recover()
	// 	fmt.Println(err)
	// }()
	panic("触发异常")
}
