package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v2"
)

func ReplaceCellValue() {

	workbook, _ := NewWorkbook_String("Data/Input/Book1.xlsx")
	replaceOption, _ := NewReplaceOptions()
	replaceOption.SetCaseSensitive(false)
	replaceOption.SetRegexKey(true)
	replaceOption.SetMatchEntireCellContents(false)
	workbook.Replace_String_String_ReplaceOptions("\bKIM\b", "^^^^^^^^", replaceOption)
	workbook.Save_String("Data/Output/Book1.pdf")
}
