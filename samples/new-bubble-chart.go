package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp/v24"
)

func NewBubbleChart() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	set_cell_string_value_2(cells, 0, 0, "Values")
	set_cell_int_value_2(cells, 0, 1, 2)
	set_cell_int_value_2(cells, 0, 2, 4)
	set_cell_int_value_2(cells, 0, 3, 6)

	set_cell_string_value_2(cells, 1, 0, "Bubble Size")
	set_cell_int_value_2(cells, 1, 1, 2)
	set_cell_int_value_2(cells, 1, 2, 3)
	set_cell_int_value_2(cells, 1, 3, 1)

	set_cell_string_value_2(cells, 1, 0, "X Values")
	set_cell_int_value_2(cells, 2, 1, 1)
	set_cell_int_value_2(cells, 2, 2, 2)
	set_cell_int_value_2(cells, 2, 3, 3)
	cells.SetColumnWidth(0, 12)

	charts, _ := worksheet.GetCharts()
	charts.AddFloatingChart(ChartType_Bubble, 5, 0, 20, 8)
	chart, _ := charts.Get_Int(0)
	nseries, _ := chart.GetNSeries()
	nseries.Add_String_Bool("B1:D1", true)
	series, _ := nseries.Get(0)
	series.SetBubbleSizes("B2:D2")
	series.SetXValues("B3:D3")
	series.SetValues("B1:D1")
	workbook.Save_String("Data/Output/NewBubbleChart.xlsx")
	println("New Bubble Chart, and save as NewBubbleChart.xlsx.")
}
