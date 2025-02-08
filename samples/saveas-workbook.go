package main

func SaveasWorkbook() {
	workbook, _, _ := create_table_data()
	workbook.Save_String("Data/Output/Book1.pdf")
	println("Open Book1.xlsx, and save as Book.pdf.")
}
