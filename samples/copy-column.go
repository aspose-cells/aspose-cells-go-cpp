package main

func CopyColumn() {
	workbook, _, cells := create_table_data()
	cells.CopyColumn(cells, 0, 6)
	workbook.Save_String("Data/Output/Book1_CopyColumn.xlsx")
}
