package main

func Unlock() {
	workbook, worksheet, _ := create_table_data()
	worksheet.Unprotect()
	workbook.Save_String("Data/Output/Book1-Unlock.xlsx")
}
