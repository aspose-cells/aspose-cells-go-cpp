package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func GetStyle() {
	workbook, err1 := NewWorkbook_String("Data/Input/Book1.xlsx")
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
	style, _ := cell.GetStyle()
	isnull, _ := style.IsNull()
	if isnull {
		print("Get Style address: ")
		print(style)

	}

}
