package main

import (
	"GoDemo/src/utils"
	"fmt"
)

func main() {
	i := 10
	fmt.Println(&i)
	var ptr = &i
	*ptr = 20
	fmt.Printf("%p存放的数据是%d\n", ptr, i)
	utils.Get()
}
