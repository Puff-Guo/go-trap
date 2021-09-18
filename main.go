package main

import (
	"fmt"

	"github.com/Puff-Guo/gotrap/feature"
)

func main() {
	//testDefer()
	//feature.AppendSlc()
	//feature.SlcLen()
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
	//feature.InterfaceCmp()
	//feature.TestSelect()
	//feature.TestSliceAsOption()
	// 	在一起 starting_today 20170126  	#在一起
	// 领证用 nice_to_marry_you 20200930  	#领证
	// 办婚礼用 pick_me_up 20211003  		#举办婚礼
	outAscii("starting_today", "20170126")
	outAscii("nice_to_marry_you", "20200930")
	outAscii("pick_me_up", "20211003")
	// c := make(chan bool)
	// <-c
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

/*
input:starting_today    	out:20170126  		#在一起
input:nice_to_marry_you 	out:20200930  		#领证
input:pick_me_up        	out:20211003  		#举办婚礼
*/
func OutStr() {

}

func outAscii(in, out string) {
	inByte := []byte(in)
	outByte := []byte(out)

	fmt.Print("in:", in, ";len:", len(in), ";ascii:")
	for _, inb := range inByte {
		fmt.Print(inb)
		fmt.Print(",")
	}

	fmt.Print("\n")
	fmt.Print("out:", out, ";len:", len(out), ";ascii:")
	for _, outb := range outByte {
		fmt.Print(outb)
		fmt.Print(",")
	}
	fmt.Print("\n")
	fmt.Print("\n")
}

func dealAscii() {
	input := []byte("nicetomeetyou")
	for index, inb := range input {
		s := inb * (inb - 97)
		if s%2 == 0 {

		} else {
			input[index] = inb
		}
	}

	str := string(input)
	fmt.Println(str)
}
