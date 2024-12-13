package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v2"
)

func SetCellValue() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cell, _ := cells.Get_String("A1")
	cell.PutValue_String("嘿,你好!")
	cell, _ = cells.Get_String("A2")
	cell.PutValue_String_Bool("ok", true)
	style, _ := cell.GetStyle()
	style.SetHorizontalAlignment(TextAlignmentType_Right)
	cell.SetStyle_Style(style)
	workbook.Save_String("Data/Output/HELLO_Value.xlsx")
	println("Finish to set value , check HELLO_Value.xlsx file in output folder.")
}
