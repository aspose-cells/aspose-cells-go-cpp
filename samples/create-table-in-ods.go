package main

import (
	. "asposecells"
)

func CreateTableInOds() {

	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	set_cell_string_value(cells, "A1", "Employee")
	set_cell_string_value(cells, "B1", "Quarter")
	set_cell_string_value(cells, "C1", "Product")
	set_cell_string_value(cells, "D1", "Continent")
	set_cell_string_value(cells, "E1", "Country")
	set_cell_string_value(cells, "F1", "Sale")

	set_cell_string_value(cells, "A2", "David")
	set_cell_string_value(cells, "A3", "David")
	set_cell_string_value(cells, "A4", "David")
	set_cell_string_value(cells, "A5", "David")
	set_cell_string_value(cells, "A6", "James")

	set_cell_int_value(cells, "B2", 1)
	set_cell_int_value(cells, "B3", 2)
	set_cell_int_value(cells, "B4", 3)
	set_cell_int_value(cells, "B5", 4)
	set_cell_int_value(cells, "B6", 1)

	set_cell_string_value(cells, "C2", "Maxilaku")
	set_cell_string_value(cells, "C3", "Maxilaku")
	set_cell_string_value(cells, "C4", "Chai")
	set_cell_string_value(cells, "C5", "Maxilaku")
	set_cell_string_value(cells, "C6", "Chang")

	set_cell_string_value(cells, "D2", "Asia")
	set_cell_string_value(cells, "D3", "Asia")
	set_cell_string_value(cells, "D4", "Asia")
	set_cell_string_value(cells, "D5", "Asia")
	set_cell_string_value(cells, "D6", "Europe")

	set_cell_string_value(cells, "E2", "China")
	set_cell_string_value(cells, "E3", "India")
	set_cell_string_value(cells, "E4", "Korea")
	set_cell_string_value(cells, "E5", "India")
	set_cell_string_value(cells, "E6", "France")

	set_cell_int_value(cells, "F2", 2000)
	set_cell_int_value(cells, "F3", 500)
	set_cell_int_value(cells, "F4", 1200)
	set_cell_int_value(cells, "F5", 1500)
	set_cell_int_value(cells, "F6", 500)

	listObjects, _ := worksheet.GetListObjects()
	index, _ := listObjects.Add_String_String_Bool("A1", "F6", true)
	listObject, _ := listObjects.Get_Int(index)
	listObject.SetShowHeaderRow(true)
	listObject.SetTableStyleType(TableStyleType_TableStyleMedium10)
	listObject.SetShowTotals(true)
	listObject.SetTableStyleType(TableStyleType_TableStyleMedium10)
	workbook.Save_String("../Data/Output/CreateTable.ods")
}
