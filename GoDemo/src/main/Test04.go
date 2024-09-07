package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*a := A{
		BookName: "Go",
	}
	a.f()*/
	p := People{
		Name:   "张三",
		Age:    18,
		Gender: 1,
	}
	p.DisplayInfo()

}

type A struct {
	BookName string
}
type B interface {
	f()
}

func (a A) f() {
	fmt.Println(a.BookName)
}

type People struct {
	Name   string
	Age    int
	Gender int
}

type OtherPeople struct {
	People
}

func (p People) DisplayInfo() {
	fmt.Println(p)
}
func (p People) String() string {
	return fmt.Sprintf("Name=%s,Age=%d,Gender=%s", p.Name,
		p.Age, strconv.Itoa(p.Gender))
}
