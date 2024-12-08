package main

import (
	. "asposecells"
	"os"
)

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
	lic, _ := NewLicense()
	lic.SetLicense_String(os.Getenv("LicensePath"))

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
