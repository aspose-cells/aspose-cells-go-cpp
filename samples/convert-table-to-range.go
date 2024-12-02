package main

import (
	. "asposecells"
)

func ConvertTableToRange() {
	workbook, err1 := NewWorkbook_String("../Data/Input/BookPivotTable.xlsx")
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

	listobjects, err4 := worksheet.GetListObjects()
	if err4 != nil {
		println(err4)
	}

	listobject, err5 := listobjects.Get_Int(0)
	if err5 != nil {
		println(err5)
	}

	listobject.ConvertToRange()
	workbook.Save_String("../Data/Output/ConvertToRange.xlsx")

	println("Convert To Range, and save as ConvertToRange.xlsx.")
}