package main

import (
	"fmt"

	"github.com/Puff-Guo/gotrap/feature"
)

func main() {
	//testDefer()
	//feature.AppendSlc()
	//feature.CompStruct()
	feature.CmpStruct()
}

func testDefer() {

	defer func() {
		if err := recover(); nil != err {
			fmt.Println("main test:", err)
		}
	}()

	feature.DeferCall()

	// defer func() {
	// 	if err := recover(); nil != err {
	// 		fmt.Println("main test:", err)
	// 	}
	// }()
}
