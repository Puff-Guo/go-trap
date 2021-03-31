package main

import (
	"fmt"

	"github.com/Puff-Guo/gotrap/feature"
)

func main() {
	//testDefer()
	//feature.AppendSlc()
	//feature.CompStruct()
	//feature.CmpStruct()
	//feature.FuncClosure()
	feature.SlcLen()
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
