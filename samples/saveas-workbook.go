package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func SaveasWorkbook() {
	workbook, err1 := NewWorkbook_String("Data/Input/Book1.xlsx")
	if err1 != nil {
		println(err1)
	}

	workbook.Save_String("Data/Output/Book1.pdf")

	println("Open Book1.xlsx, and save as Book.pdf.")
}
