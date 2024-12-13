package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v2"
)

func AddCustomProperty() {
	workbook, _ := NewWorkbook()
	customDocumentProperties, _ := workbook.GetCustomDocumentProperties()
	customDocumentProperties.Add_String_String("MyCustom5", "This is my custom five.")
	workbook.Save_String("Data/Output/HELLO_CustomDocumentPropertie.xlsx")
	println("Finish to add custom properties, check HELLO_CustomDocumentPropertie.xlsx file in output folder.")
}
