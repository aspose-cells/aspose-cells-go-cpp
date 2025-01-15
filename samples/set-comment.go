package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func SetComment() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	comments, _ := worksheet.GetComments()
	comment_index, _ := comments.Add_String("A1")
	println("Get comment index:")
	comment, _ := comments.Get_Int(comment_index)
	comment.SetNote("Hello, World!")
	workbook.Save_String("Data/Output/HELLO_Comment.xlsx")
	println("Finish to set comment, check HELLO_Comment.xlsx file in output folder.")
}
