package feature

var (
	n int
)

const (
	N int = 5
)

/*
	报错:non-constant array bound n
	指定数组长度需要常量
*/
// func GetMax(a [n]int) int {
// 	return n
// }

func GetMax(a [N]int) int {
	return n
}
