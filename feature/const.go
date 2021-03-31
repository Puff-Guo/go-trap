package feature

import "fmt"

const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func ConstAssign() {
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}
