package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func ChartToImage() {

	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	set_cell_int_value(cells, "A1", 50)
	set_cell_int_value(cells, "A2", 100)
	set_cell_int_value(cells, "A3", 150)

	set_cell_int_value(cells, "B1", 4)
	set_cell_int_value(cells, "B2", 20)
	set_cell_int_value(cells, "B3", 50)

	charts, _ := worksheet.GetCharts()
	charts.AddFloatingChart(ChartType_Pyramid, 5, 0, 20, 8)
	chart, _ := charts.Get_Int(0)
	chart.ToImage_String_ImageType("Data/Output/ChartToImage.png", ImageType_Png)
	println("Chart to png, and save as ChartToImage.png.")
}
