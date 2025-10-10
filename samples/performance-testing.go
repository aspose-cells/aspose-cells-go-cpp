package main

import (
	"fmt"
	"time"

	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func PerformanceTesting1() {
	now := time.Now()
	fmt.Println("Begin Time:", now.Format(time.RFC3339))
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	var i int32
	for i = 0; i < 10; i++ {
		worksheetIndex, _ := worksheets.Add()
		worksheet, _ := worksheets.Get_Int(worksheetIndex)
		cells, _ := worksheet.GetCells()
		rows, _ := cells.GetRows()
		// rowEnumerator ,_ := rows.GetEnumerator()
		// rowEnumerator.MoveNext()

		var rowIndex, columnIndex int32
		for rowIndex = 0; rowIndex < 100000; rowIndex++ {
			row, _ := rows.Get(rowIndex)
			for columnIndex = 0; columnIndex < 255; columnIndex++ {
				cell, _ := row.Get(columnIndex)
				cell.PutValue_Int(i*columnIndex + rowIndex)
			}
		}
	}
	workbook.Save_String("Data/Output/PerformanceTesting.xlsx")
	now = time.Now()
	fmt.Println("End Time:", now.Format(time.RFC3339))
	println("Finish to set value , check .xlsx file in output folder.")
}

func PerformanceTesting() {
	now := time.Now()
	fmt.Println("Begin Time:", now.Format(time.RFC3339))
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	var i int32
	for i = 0; i < 10; i++ {
		worksheetIndex, _ := worksheets.Add()
		worksheet, _ := worksheets.Get_Int(worksheetIndex)
		cells, _ := worksheet.GetCells()
		var row, column int32
		for row = 0; row < 100000; row++ {
			for column = 0; column < 200; column++ {
				cell, _ := cells.Get_Int_Int(row, column)
				cell.PutValue_Int(i*column + row)
			}
		}
	}
	workbook.Save_String("Data/Output/PerformanceTesting.xlsx")
	now = time.Now()
	fmt.Println("End Time:", now.Format(time.RFC3339))
	println("Finish to set value , check .xlsx file in output folder.")
}
