package feature

import (
	"fmt"
)

func TestSelect() {
	var intChan chan (int)
	for i := 0; i < 10; i++ {
		select {
		case <-intChan:
			fmt.Println("select read!")
		default:
			fmt.Println("select defalt!")
		}
	}
}
