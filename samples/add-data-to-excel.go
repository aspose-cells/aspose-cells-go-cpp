package main

import (
	. "asposecells"
)

func AddDataToExcel() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cell, _ := cells.Get_String("A1")
	cell.PutValue_String("嘿,你好!")
	cell, _ = cells.Get_String("A2")
	cell.PutValue_Double(20.5)
	cell, _ = cells.Get_String("A3")
	cell.PutValue_Int(25)
	cell, _ = cells.Get_String("A4")
	cell.PutValue_Bool(true)
	cell, _ = cells.Get_String("A5")
	cell.PutValue_Int(45739)
	style, _ := cell.GetStyle()
	style.SetNumber(11)
	cell.SetStyle_Style(style)
	workbook.Save_String("../Data/Output/HELLO_AddData.xlsx")
	println("Finish to set value , check HELLO_AddData.xlsx file in output folder.")
}
