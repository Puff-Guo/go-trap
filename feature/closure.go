package feature

import (
	"errors"
	"fmt"
)

type Person struct {
	age int
}

// 1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
// 2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
// 3.闭包引用，相当与指针,输出 29；
func FuncClosure() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

/*
	return defer 返回值 执行顺序:1 return 把返回值放到堆栈 2 defer执行收尾 3 调用者获取堆栈值(返回值)
	func()int 这种因为没有提前声明,第三步返回值由于defer没有办法引用,所以不能改变
	func()(i int) 这种由于有提前声明,所以defer可以修改堆栈值,可以改变
*/
func foo() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()

	defer func(e error) {
		fmt.Println(e)
		e = errors.New("b")
	}(err)
	return errors.New("c")
}

/*
<nil>
c
a
*/
func DeferTest() {
	fmt.Println(foo())
}
