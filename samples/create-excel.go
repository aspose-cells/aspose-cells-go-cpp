package main

import (
	. "asposecells"
)

func CreateExcel() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cell, _ := cells.Get_String("A1")
	cell.PutValue_Int(5)
	cell, _ = cells.Get_String("A2")
	cell.PutValue_Int(15)
	cell, _ = cells.Get_String("A3")
	cell.PutValue_Int(25)
	workbook.Save_String("Data/Output/HELLO_Create.xlsx")
	println("Finish to set value , check .xlsx file in output folder.")
}
