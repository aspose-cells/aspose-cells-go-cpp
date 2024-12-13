package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp"
)

func SearchCellValue() {

	workbook, _ := NewWorkbook_String("Data/Input/Book1.xlsx")
	println("NewWorkbook_String")
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	findOptions, _ := NewFindOptions()
	cell, _ := cells.GetFirstCell()
	findOptions.SetLookAtType(LookAtType_EntireContent)
	obj, _ := NewObject_String("Hello")
	println(obj)
	cell1, _ := cells.Find_Object_Cell_FindOptions(obj, cell, findOptions)
	cellValue, _ := cell1.GetStringValue()
	println(cellValue)
}
