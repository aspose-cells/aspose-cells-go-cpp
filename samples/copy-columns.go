package main

import (
	. "asposecells"
)

func CopyColumns() {
	workbook, _ := NewWorkbook_String("../Data/Input/Book1.xlsx")
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cells.CopyColumns_Cells_Int_Int_Int(cells, 0, 6, 2)
	workbook.Save_String("../Data/Output/Book1_CopyColumns.xlsx")
}
