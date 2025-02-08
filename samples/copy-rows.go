package main

func CopyRows() {
	workbook, _, cells := create_table_data()
	cells.CopyRows_Cells_Int_Int_Int(cells, 0, 6, 2)
	workbook.Save_String("Data/Output/Book1_CopyRows.xlsx")
}
