package feature

import "fmt"

func AppendSlc() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println("s:", s)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3)
	fmt.Println("s2:", s2)

	//未初始化
	var s3 []int
	s3 = append(s3, 1, 2, 3)
	fmt.Println("s3:", s3)

	//
	s4 := []int{1, 2, 3}
	s5 := []int{4, 5}
	//s4 = append(s4, s5)		build error need add ...
	s4 = append(s4, s5...)
	fmt.Println(s4)

}
