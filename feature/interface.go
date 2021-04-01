package feature

import (
	"fmt"
)

type T interface{}

/*
	空接口比较:比较type和value,类型和value都相等时true,如果type是不可比较类型(map,slice)程序会panic
				1> 先比较type,type不同返回false
				2> 再比较value,如果type是指针,只需要比较指针值,如果value相同,返回true,否则返回false

	接口之间的赋值:
				1> type改为源接口的type
				2> value改为源接口值的副本

	p *T 不能算接口间的赋值,*T是未命名类型,i2.type未*T
*/
func InterfaceCmp() {
	var (
		t T
		p *T

		//i1 interface{} = T		//这里会报:type T is not an expression
		i1 interface{} = t
		i2 interface{} = p
	)

	fmt.Println(i1 == t, i1 == nil)
	fmt.Println(i2 == p, i2 == nil)
}
