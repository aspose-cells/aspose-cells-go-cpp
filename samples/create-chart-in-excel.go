package main

import (
	. "github.com/aspose-cells/aspose-cells-go-cpp"
)

func CreateChartInExcel() {
	workbook, _ := NewWorkbook()
	worksheets, _ := workbook.GetWorksheets()
	worksheet, _ := worksheets.Get_Int(0)
	cells, _ := worksheet.GetCells()
	cell, _ := cells.Get_String("A1")
	cell.PutValue_Int(50)
	cell, _ = cells.Get_String("A2")
	cell.PutValue_Int(100)
	cell, _ = cells.Get_String("A3")
	cell.PutValue_Int(150)
	cell, _ = cells.Get_String("B1")
	cell.PutValue_Int(4)
	cell, _ = cells.Get_String("B2")
	cell.PutValue_Int(20)
	cell, _ = cells.Get_String("B3")
	cell.PutValue_Int(50)
	charts, _ := worksheet.GetCharts()
	chartIndex, _ := charts.Add_ChartType_Int_Int_Int_Int(ChartType_Pyramid, 5, 0, 20, 8)
	chart, _ := charts.Get_Int(chartIndex)
	series, _ := chart.GetNSeries()
	series.Add_String_Bool("A1:B3", true)
	workbook.Save_String("Data/Output/Hello_CreateChart.xlsx")
}
