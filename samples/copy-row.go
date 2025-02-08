package main

func CopyRow() {
	workbook, _, cells := create_table_data()
	cells.CopyRow(cells, 0, 6)
	workbook.Save_String("Data/Output/Book1_CopyRow.xlsx")
}
