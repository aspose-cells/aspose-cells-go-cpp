package main

func CopyColumns() {
	workbook, _, cells := create_table_data()
	cells.CopyColumns_Cells_Int_Int_Int(cells, 0, 6, 2)
	workbook.Save_String("Data/Output/Book1_CopyColumns.xlsx")
}
