package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

func ReplaceCellValue() {

	workbook, _, _ := create_table_data()
	replaceOption, _ := NewReplaceOptions()
	replaceOption.SetCaseSensitive(false)
	replaceOption.SetRegexKey(true)
	replaceOption.SetMatchEntireCellContents(false)
	workbook.Replace_String_String_ReplaceOptions("China", "^^^^^^^^", replaceOption)
	workbook.Save_String("Data/Output/Book1.pdf")
}
