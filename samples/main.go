package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
	"os"
)

func create_folder(folder string) {
	_, err := os.Stat(folder)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		os.MkdirAll(folder, 0755)
		return
	}
}
func create_table_data() (*Workbook, *Worksheet, *Cells) {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	set_cell_string_value(cells, "A1", "Employee")
	set_cell_string_value(cells, "B1", "Quarter")
	set_cell_string_value(cells, "C1", "Product")
	set_cell_string_value(cells, "D1", "Continent")
	set_cell_string_value(cells, "E1", "Country")
	set_cell_string_value(cells, "F1", "Sale")

	set_cell_string_value(cells, "A2", "David")
	set_cell_string_value(cells, "A3", "David")
	set_cell_string_value(cells, "A4", "David")
	set_cell_string_value(cells, "A5", "David")
	set_cell_string_value(cells, "A6", "James")

	set_cell_int_value(cells, "B2", 1)
	set_cell_int_value(cells, "B3", 2)
	set_cell_int_value(cells, "B4", 3)
	set_cell_int_value(cells, "B5", 4)
	set_cell_int_value(cells, "B6", 1)

	set_cell_string_value(cells, "C2", "Maxilaku")
	set_cell_string_value(cells, "C3", "Maxilaku")
	set_cell_string_value(cells, "C4", "Chai")
	set_cell_string_value(cells, "C5", "Maxilaku")
	set_cell_string_value(cells, "C6", "Chang")

	set_cell_string_value(cells, "D2", "Asia")
	set_cell_string_value(cells, "D3", "Asia")
	set_cell_string_value(cells, "D4", "Asia")
	set_cell_string_value(cells, "D5", "Asia")
	set_cell_string_value(cells, "D6", "Europe")

	set_cell_string_value(cells, "E2", "China")
	set_cell_string_value(cells, "E3", "India")
	set_cell_string_value(cells, "E4", "Korea")
	set_cell_string_value(cells, "E5", "India")
	set_cell_string_value(cells, "E6", "France")

	set_cell_int_value(cells, "F2", 2000)
	set_cell_int_value(cells, "F3", 500)
	set_cell_int_value(cells, "F4", 1200)
	set_cell_int_value(cells, "F5", 1500)
	set_cell_int_value(cells, "F6", 500)
	return workbook, worksheet, cells
}

func set_cell_int_value(cells *Cells, cell_name string, value int32) {
	cell, _ := cells.Get_String(cell_name)
	cell.PutValue_Int(value)
}

func set_cell_string_value(cells *Cells, cell_name string, value string) {
	cell, _ := cells.Get_String(cell_name)
	cell.PutValue_String(value)
}

func set_cell_int_value_2(cells *Cells, row int32, column int32, value int32) {
	cell, _ := cells.Get_Int_Int(row, column)
	cell.PutValue_Int(value)
}

func set_cell_string_value_2(cells *Cells, row int32, column int32, value string) {
	cell, _ := cells.Get_Int_Int(row, column)
	cell.PutValue_String(value)
}

func main() {
	println("Hello World!")
	create_folder("Data")
	create_folder("Data/Output")
	lic, _ := NewLicense()
	err := lic.SetLicense_String(os.Getenv("LicensePath"))
	if err != nil {
		println(err)
	}
	Version()
	AddCustomProperty()
	AddDataToExcel()
	AddHyperlinkToImage()
	CalculateFormulaInExcel()
	ChartToImage()
	ConvertExcelToOds()
	ConvertExcelToPdf()
	ConvertTableToRange()
	CopyColumn()
	CopyColumns()
	CopyRow()
	CopyRows()
	CreateChartInExcel()
	CreateExcel()
	CreateTableInExcel()
	CreateTableInOds()
	GetCellValue()
	GetStyle()
	HideRowsInPivotTable()
	NewBubbleChart()
	NewLineChartInOds()
	NewLineChart()
	NewPyramidChart()
	New_Workbook()
	SaveasWorkbook()
	SetBackgroundColor()
	SetCellValue()
	SetComment()
	SetListobjectComment()
	SetStyle()
	SearchCellValue()
	ShowCellsHelper()
	ReplaceCellValue()
	Unlock()
}
