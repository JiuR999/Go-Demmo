package main

import "fmt"

func main() {
	ch := make(chan int, 20)
	go rec(ch)
	func() {
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Println(i, "-Send Successful!")
		}
	}()

}

func rec(ch chan int) {
	ret := <-ch
	fmt.Println(ret)
}
