package main

import (
	. "asposecells"
)

func SetStyle() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cell, _ := cells.Get_String("A1")
	cell.PutValue_String("Hello, World!")
	cell, _ = cells.Get_String("A2")
	cell.PutValue_String_Bool("ok", true)
	style, _ := cell.GetStyle()
	style.SetHorizontalAlignment(TextAlignmentType_Right)
	cell.SetStyle_Style(style)
	workbook.Save_String("Data/Output/HELLO_Style.xlsx")
	println("Finish to set style, check HELLO.xlsx file in output folder.")
}
