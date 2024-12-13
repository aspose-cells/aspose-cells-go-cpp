package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp"
)

func NewLineChartInOds() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	set_cell_int_value(cells, "A1", 50)
	set_cell_int_value(cells, "A2", 100)
	set_cell_int_value(cells, "A3", 150)
	set_cell_int_value(cells, "A4", 110)
	set_cell_int_value(cells, "B1", 260)
	set_cell_int_value(cells, "B2", 12)
	set_cell_int_value(cells, "B3", 50)
	set_cell_int_value(cells, "B4", 100)
	charts, _ := worksheet.GetCharts()
	charts.AddFloatingChart(ChartType_Column, 5, 0, 20, 8)
	chart, _ := charts.Get_Int(0)
	nseries, _ := chart.GetNSeries()
	nseries.Add_String_Bool("A1:B4", true)
	series, _ := nseries.Get(0)
	series.SetType(ChartType_Line)
	workbook.Save_String("Data/Output/NewColumnChart.xlsx")
	println("New Column Chart, and save as NewColumnChart.xlsx.")

}
