package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func SetBackgroundColor() {
	workbook, err1 := NewWorkbook()
	if err1 != nil {
		println(err1)
	}

	worksheets, err2 := workbook.GetWorksheets()
	if err2 != nil {
		println(err2)
	}
	worksheet, err3 := worksheets.Get_Int(0)
	if err3 != nil {
		println(err3)
	}

	cells, err4 := worksheet.GetCells()
	if err4 != nil {
		println(err4)
	}

	cell, err5 := cells.Get_String("A1")
	if err5 != nil {
		println(err5)
	}

	cell.PutValue_String_Bool("Color", true)
	style, _ := cell.GetStyle()
	style.SetPattern(BackgroundType_Solid)
	color, err6 := NewColor()
	if err6 != nil {
		println(err6)
	}
	color.Set_Color_R(uint8(255))
	color.Set_Color_G(uint8(128))
	err7 := style.SetForegroundColor(color)
	if err7 != nil {
		println(err7)
	}
	cell.SetStyle_Style(style)
	workbook.Save_String("Data/Output/Coloer.xlsx")

	println("New workbook, set color, and save as HELLO.xlsx.")
}
