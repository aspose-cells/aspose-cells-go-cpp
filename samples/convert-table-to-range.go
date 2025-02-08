package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func ConvertTableToRange() {
	workbook, worksheet, _ := create_table_data()
	listObjects, _ := worksheet.GetListObjects()
	index, _ := listObjects.Add_String_String_Bool("A1", "F6", true)
	listObject, _ := listObjects.Get_Int(index)
	listObject.SetShowHeaderRow(true)
	listObject.SetTableStyleType(TableStyleType_TableStyleMedium10)
	listObject.SetShowTotals(true)
	listObject.ConvertToRange()
	workbook.Save_String("Data/Output/ConvertToRange.xlsx")

	println("Convert To Range, and save as ConvertToRange.xlsx.")
}
