package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Greet(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(msg)
	}
	wg.Done()
}

func DisPlayPrimeNumber(primry chan int) {
	for v := range primry {
		fmt.Println(v)
	}
	wg.Done()
}

func CountPrimeNumber(primry chan int) {
	for i := 1; i < 100; i++ {
		if IsPrimeNumber(i) {
			primry <- i
		}
	}
	close(primry)
	wg.Done()
}
func IsPrimeNumber(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	/*wg.Add(1)
	go Greet("你好", 5)
	wg.Add(1)
	go Greet("早上好", 5)

	wg.Wait()*/
	wg.Add(1)
	start := time.Now()
	countCh := make(chan int, 1000)
	go CountPrimeNumber(countCh)
	wg.Add(1)
	go DisPlayPrimeNumber(countCh)

	wg.Wait()
	fmt.Println("打印花费", time.Now().Sub(start))
}
