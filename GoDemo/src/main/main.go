package main

import (
	"bytes"
	"fmt"
)

func main() {
	a, b := 8, "Hello Go"
	s := ""
	buf := bytes.NewBufferString(s)
	fmt.Fprint(buf, a, b)
	fmt.Println(buf)
	fmt.Println("Hello, World!")

	var length int
	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	data := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&data[i])
	}
	/*array2 := sort(data)
	display(array2)*/
	selectSort(data)
}

// 冒泡排序
func sort(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1; j++ {
			if array[j] < array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
	return array
}

func selectSort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		k := i
		for j := i + 1; j < length-1; j++ {
			if array[j] < array[k] {
				k = j
			}
		}
		if k != i {
			tmp := array[i]
			array[i] = array[k]
			array[k] = tmp
		}
		fmt.Print("第", i, "次排序: ")
		display(array)
		fmt.Println()
	}
	return array
}

func display(array []int) {
	for i := 0; i < len(array); i++ {
		print("  ", array[i])
	}
}
