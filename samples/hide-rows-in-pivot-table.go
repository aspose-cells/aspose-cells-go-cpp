package main

import (
	. "asposecells"
)

func HideRowsInPivotTable() {
	println("LicensePath.")
	workbook, _ := NewWorkbook_String("Data/Input/TestCase.xlsx")
	println("NewWorkbook_String.")
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(1)
	cells, _ := worksheet.GetCells()
	println("worksheets 0.")
	pivotTables, _ := worksheet.GetPivotTables()
	pivotTable, _ := pivotTables.Get_Int(0)
	println("pivotTable 0.")
	dataBodyRange, _ := pivotTable.GetDataBodyRange()
	println("dataBodyRange.", dataBodyRange)
	currentRow := int32(5)
	rowsUsed, _ := dataBodyRange.Get_EndRow()
	println("rowsUsed.", rowsUsed)
	for i := currentRow; i <= rowsUsed; i++ {
		cell, _ := cells.Get_Int_Int(i, 4)
		value, _ := cell.GetStringValue()
		if value == "Total" {
			cells.HideRow(i)
		}
	}
	println(rowsUsed)
	workbook.Save_String("Data/Output/HELLO_HideRow.xlsx")
	println("Finish to set value , check HELLO_HideRow.xlsx file in output folder.")
}
