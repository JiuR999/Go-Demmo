package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

func main() {
	file, err := excelize.OpenFile("C:\\Users\\Simon\\Desktop\\hard_ware.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := file.GetRows("hard_ware")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range rows {
		for i, cell := range v {
			fmt.Printf("%v-%v\t", i, cell)
		}
		fmt.Println()
	}

	fmt.Println(os.Getwd())
}
