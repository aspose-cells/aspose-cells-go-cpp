package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func CalculateFormulaInExcel() {

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
	cell, _ = cells.Get_String("A4")
	cell.SetFormula_String("=SUM(A1:A3)")
	workbook.CalculateFormula()
	println(cell.GetIntValue())
}
