package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func CopyColumn() {
	workbook, _ := NewWorkbook_String("Data/Input/Book1.xlsx")
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cells.CopyColumn(cells, 0, 6)
	workbook.Save_String("Data/Output/Book1_CopyColumn.xlsx")
}
