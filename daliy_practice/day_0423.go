package main

import (
	"fmt"
)

type Student struct {
	name   string
	age    int
	weight float32
	score  []int
}

func main() {
	pp := new(Student) //使用 new 关键字创建一个指针
	*pp = Student{"qishuangming", 23, 65.0, []int{2, 3, 6}}
	fmt.Printf("stu pp have %d subjects\n", len((*pp).score))
	*pp = Student{name: "zc"}
	fmt.Printf("stu name %s\n", pp.name) //Go语言自带隐式解引用
	aa := new(Student)
	*aa = Student{name: "123"}
	*pp = Student{name: "zc"}
	fmt.Printf("stu name %s\n", aa.name)

	a := &Student{name: "456"}
	a.name = "999"
	fmt.Printf("%v", a)
}
