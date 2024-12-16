package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func New_Workbook() {
	workbook, err1 := NewWorkbook()
	if err1 != nil {
		println(err1)
	}

	worksheets, err2 := workbook.GetWorksheets()
	if err2 != nil {
		println(err2)
	}
	worksheet, err3 := worksheets.Get_Int(0)
	if err3 != nil {
		println(err3)
	}

	cells, err4 := worksheet.GetCells()
	if err4 != nil {
		println(err4)
	}

	cell, err5 := cells.Get_String("A1")
	if err5 != nil {
		println(err5)
	}

	cell.PutValue_String_Bool("Hello", true)
	workbook.Save_String("Data/Output/HELLO.xlsx")

	println("New workbook, and save as HELLO.xlsx.")
}
