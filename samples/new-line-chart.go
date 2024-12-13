package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v2"
)

func NewLineChart() {
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
	charts.AddFloatingChart(ChartType_Line, 5, 0, 20, 8)
	chart, _ := charts.Get_Int(0)
	nseries, _ := chart.GetNSeries()
	nseries.Add_String_Bool("A1:B3", true)
	workbook.Save_String("Data/Output/NewLineChart.xlsx")
	println("New Line Chart, and save as NewLineChart.xlsx.")
}
