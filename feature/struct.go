package feature

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

//可排序的数据类型:integer,floating point,string
//可比较的数据类型:boolean,Complex,Pointer,Channel,interface,array
//不可比较的数据类型:slice,map,function

//struct可比较条件:每个字段必须是可比较的

type Contact struct {
	Phone string
	Email string
}

type User struct {
	Name    string
	Age     int
	Contact *Contact
}

func CompStruct() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//结构体只能比较是否相等
	// if sn1 > sn2 {
	// 	fmt.Println("sn1 == sn2")
	// }

	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}
	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }

	// var comp1 Contact = Contact{Phone: 11, Email: "qq"}
	// var comp2 Contact = Contact{Phone: 11, Email: "qq"}

	// if comp1 == comp2 {
	// 	fmt.Println("comp1 == comp2")
	// }

	// if comp2 > comp1 {
	// 	fmt.Println("comp2 == comp1")
	// }
}

func CmpStruct() {
	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 18}

	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c1 := &Contact{Phone: "123456789", Email: "dj@example.com"}
	c2 := &Contact{Phone: "123456789", Email: "dj@example.com"}

	u1.Contact = c1
	u2.Contact = c1
	fmt.Println("u1 == u2 with same pointer?", u1 == u2)
	fmt.Println("u1 equals u2 with same pointer?", cmp.Equal(u1, u2))

	u2.Contact = c2
	fmt.Println("u1 == u2 with different pointer?", u1 == u2)
	fmt.Println("u1 equals u2 with different pointer?", cmp.Equal(u1, u2))
}
