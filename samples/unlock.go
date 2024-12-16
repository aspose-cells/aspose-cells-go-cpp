package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func Unlock() {
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

	worksheet.Unprotect()
	workbook.Save_String("Data/Output/Book1-Unlock.xlsx")
}
