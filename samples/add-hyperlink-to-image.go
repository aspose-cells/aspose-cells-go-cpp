package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func AddHyperlinkToImage() {

	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cell, _ := cells.Get_String("C2")
	cell.PutValue_String("Image Hyperlink")
	cells.SetColumnWidth(2, 21)
	cells.SetRowHeight(3, 100)
	pictures, _ := worksheet.GetPictures()
	picture_index, err := pictures.Add_Int_Int_Int_Int_String(3, 2, 4, 3, "asposecells.png")
	if err != nil {
		println("err.Error()")
		println(err.Error())
	}
	count, _ := pictures.GetCount()
	println("Getting picture count...", count)
	picture, _ := pictures.Get(picture_index)
	picture.SetPlacement(PlacementType_FreeFloating)
	hyperlink, _ := picture.AddHyperlink("https://www.aspose.com/")
	hyperlink.SetScreenTip("Click to go to Aspose site")
	workbook.Save_String("Data/Output/HELLO_Image_HyperLink.xlsx")
	println("Finish to operate, check HELLO_Image_HyperLink.xlsx file in output folder.")
}
