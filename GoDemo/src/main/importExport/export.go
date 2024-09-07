package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

const SHEET1 = "sheet1"
const PATH = "D://test.xlsx"

func main() {
	file, err := excelize.OpenFile(PATH)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	cell, _ := excelize.CoordinatesToCellName(1, 2)
	index, _ := file.NewSheet(SHEET1)
	file.SetActiveSheet(index)
	file.SetCellValue(SHEET1, cell, "Hello Excel!")

	for i := 2; i <= 10; i++ {
		cellTitle, _ := excelize.CoordinatesToCellName(1, i)
		file.SetCellValue(SHEET1, cellTitle, "Title"+strconv.Itoa(i))
		cellContent, _ := excelize.CoordinatesToCellName(2, i)
		file.SetCellValue(SHEET1, cellContent, "Content"+strconv.Itoa(i))
		cellInfluence, _ := excelize.CoordinatesToCellName(3, i)
		file.SetCellValue(SHEET1, cellInfluence, "Influence"+strconv.Itoa(i))
	}

	file.SaveAs(PATH)
	fmt.Println("导出完成")
}
