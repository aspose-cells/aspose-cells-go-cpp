package main

func HideRowsInPivotTable() {
	workbook, worksheet, cells := create_table_data()
	pivotTables, _ := worksheet.GetPivotTables()
	pivotTableIndex, _ := pivotTables.Add_String_String_String("=Sheet1!A1:F6", "A12", "TestPivotTable")
	pivotTable, _ := pivotTables.Get_Int(pivotTableIndex)
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
