package main

func SetListobjectComment() {
	workbook, worksheet, _ := create_table_data()
	listObjects, _ := worksheet.GetListObjects()
	index, _ := listObjects.Add_String_String_Bool("A1", "F6", true)
	listObject, _ := listObjects.Get_Int(index)
	listObject.SetComment("This is Aspose.Cells comment.")
	workbook.Save_String("Data/Output/BookPivotTableSetComment.xlsx")

	println("Set comment, and save as BookPivotTableSetComment.xlsx.")
}
