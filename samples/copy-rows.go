package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func CopyRows() {
	workbook, _ := NewWorkbook_String("Data/Input/Book1.xlsx")
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cells.CopyRows_Cells_Int_Int_Int(cells, 0, 6, 2)
	workbook.Save_String("Data/Output/Book1_CopyRows.xlsx")
}
