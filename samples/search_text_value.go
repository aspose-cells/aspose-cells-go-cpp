package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func SearchCellValue() {
	_, _, cells := create_table_data()
	findOptions, _ := NewFindOptions()
	cell, _ := cells.GetFirstCell()
	findOptions.SetLookAtType(LookAtType_EntireContent)
	obj, _ := NewObject_String("Product")
	println(obj)
	cell1, _ := cells.Find_Object_Cell_FindOptions(obj, cell, findOptions)
	cellValue, _ := cell1.GetStringValue()
	println(cellValue)
}
