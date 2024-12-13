package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v2"
)

func CopyRow() {
	workbook, _ := NewWorkbook_String("Data/Input/Book1.xlsx")
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cells.CopyRow(cells, 0, 6)
	workbook.Save_String("Data/Output/Book1_CopyRow.xlsx")
}
