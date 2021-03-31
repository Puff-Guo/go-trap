package feature

import (
	"fmt"
	"time"
)

/*
2 3
2 3
2 3
*/
func RangeTmp() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()

		// 正确写法
		// go func(i,v int) {
		//     fmt.Println(i, v)
		// }(i,v)
	}

	time.Sleep(time.Second * 3)
}
