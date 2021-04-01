package feature

import (
	"fmt"
	"reflect"
)

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

/*
a len: 0 ,cap: 3
b len: 2 ,cap: 3
c len: 1 ,cap: 2
截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
*/
func SlcLen() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Println("a len:", len(a), ",cap:", cap(a))
	fmt.Println("b len:", len(b), ",cap:", cap(b))
	fmt.Println("c len:", len(c), ",cap:", cap(c))
}

/*
invalid operation: [1]int literal == [2]int literal (mismatched types [1]int and [2]int)
invalid operation: []int literal == []int literal (slice can only be compared to nil)
go 中不同类型是不能比较的，而数组长度是数组类型的一部分
*/
func SlcCmp() {
	//fmt.Println([...]int{1} == [2]int{1})
	//fmt.Println([]int{1} == []int{1})
}

func SlcRange() {
	fmt.Println("SlcRange slice")
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		fmt.Println("i:", i, " v:", v, " v point:", &v, " v type:", reflect.TypeOf(v).Name())
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func ArgRange() {
	fmt.Println("ArgRange array")
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		fmt.Println("i:", i, " v:", v, " v point:", &v, " v type:", reflect.TypeOf(v).Name())
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

//循环切片追加
//不会出现死循环，能正常结束。循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。
func SlcRangeAdd() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Println(v)
}

//切片共享底层数组及扩容
func AppendShareArrary() {
	fmt.Println("AppendShareArrary")
	a := [3]int{0, 1, 2}
	s := a[1:2] // {1}

	s[0] = 11         // 由于共享底层数组，所以a变为{0, 11, 2} s变为{11}
	s = append(s, 12) // a变为{0, 11, 12} s变为{11, 12}
	s = append(s, 13) // 超过cap，s发生扩容，不再共享a的底层数组 s变为{11, 12, 13} a仍为{0, 11, 12}
	s[0] = 21         // s变为{21, 12, 13} 底层数组不共享a，a仍然为{0, 11, 12}

	fmt.Println(a)
	fmt.Println(s)
}

//切片由切片初始化
func AppendShareSlc() {
	fmt.Println("AppendShareSlc")
	a := []int{0, 1, 2}
	s := a[1:2] // {1}

	s[0] = 11         // 由于共享底层数组，所以a变为{0, 11, 2} s变为{11}
	s = append(s, 12) // a变为{0, 11, 12} s变为{11, 12}
	s = append(s, 13) // 超过cap，s发生扩容，不再共享a的底层数组 s变为{11, 12, 13} a仍为{0, 11, 12}
	s[0] = 21         // s变为{21, 12, 13} 底层数组不共享a，a仍然为{0, 11, 12}

	fmt.Println(a)
	fmt.Println(s)
}

func Append(s []int) {
	s = append(s, 5)
}

func Add(s []int) {
	for i := range s {
		s[i] = s[i] + 5
	}
}

//由于go都是值传递,外部s值指定位置没有边,传入Append由于扩容后底层数据换了,所以外部s是没有变的,Add是操作引用还是会影响实际值
func SlcAppendPoint() {
	s := []int{1, 2, 3, 4}
	Append(s)
	fmt.Println(s)
	Add(s)
	fmt.Println(s)
}
