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
	//feature.RangeTmp()
	//feature.SlcRange()
	//feature.ArgRange()
	//feature.SlcRangeAdd()
	//feature.ConstAssign()
	//feature.MutexAssign()
	//feature.AppendShareArrary()
	//feature.AppendShareSlc()
	//feature.SlcAppendPoint()
	//feature.DeferTest()
	feature.InterfaceCmp()
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
