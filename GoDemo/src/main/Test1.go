package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	/*for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "*", j, "=", i*j, "	")
		}
		fmt.Println()
	}*/
	doGet("https://api.pearktrue.cn/api/xfai/?message=锄禾日当午下一句")

	i, _ := strconv.ParseInt("19", 10, 64)
	fmt.Printf("类型%T,i=%v\n", i, i)
	fmt.Println(getSum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	fmt.Println(getJie(10))

	r1 := oper(3, 4, add)
	fmt.Println(r1)

	r2 := oper(3, 4, sub)
	fmt.Println(r2)

	a, b := 2, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func getJie(n int) int {
	if n == 1 {
		return 1
	}
	return n * getJie(n-1)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func getSum(nums ...int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

func oper(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

func doGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	fmt.Printf("%s", result)
}
