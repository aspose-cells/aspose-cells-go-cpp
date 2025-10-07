package main

import (
	"fmt"
	"time"

	. "github.com/aspose-cells/aspose-cells-go-cpp/v25"
)

type ReportGenerator struct {
	workbook  *Workbook
	worksheet *Worksheet
}

type SaleRecord struct {
	Date     time.Time
	Product  string
	Quantity int32
	Price    float64
}

func (reportGenerator *ReportGenerator) initialization() {
	reportGenerator.workbook, _ = NewWorkbook()
	worksheets, _ := reportGenerator.workbook.GetWorksheets()
	reportGenerator.worksheet, _ = worksheets.Get_Int(0)
}
func NewReportGenerator() *ReportGenerator {
	reportGenerator := &ReportGenerator{}
	reportGenerator.initialization()
	return reportGenerator
}

func (reportGenerator *ReportGenerator) GenerateSaleReport(records []SaleRecord) {
	reportGenerator.setupReportHeader()
	reportGenerator.fillSalesData(records)
	reportGenerator.addSummaryStatistics(records)
	reportGenerator.createCharts(records)
	reportGenerator.formatReport()
}
func (reportGenerator *ReportGenerator) createHeaderStyle() *Style {
	style, _ := reportGenerator.workbook.CreateStyle()
	font, _ := style.GetFont()
	font.SetIsBold(true)
	font.SetSize(16)
	style.SetHorizontalAlignment(TextAlignmentType_Center)
	return style
}
func (reportGenerator *ReportGenerator) setupReportHeader() {
	cells, _ := reportGenerator.worksheet.GetCells()
	cell, _ := cells.Get_String("A1")
	cell.PutValue_String("Sales Report")
	cell.SetStyle_Style(reportGenerator.createHeaderStyle())
	cells.Merge_Int_Int_Int_Int(0, 0, 1, 5)
}

func (reportGenerator *ReportGenerator) fillSalesData(saleRecords []SaleRecord) {
	headers := []string{"date", "product", "Quantity", "Unit Price", "Total Amount"}
	cells, _ := reportGenerator.worksheet.GetCells()
	rows, _ := cells.GetRows()
	row, _ := rows.Get(2)
	for i := 0; i < len(headers); i++ {
		cell, _ := row.Get(int32(i))
		cell.PutValue_String(headers[i])
	}

	var rowIndex int32 = 3
	for i, record := range saleRecords {
		row, _ := rows.Get(rowIndex + int32(i))
		cell, _ := row.Get(0)
		cell.PutValue_Date(record.Date)
		cell, _ = row.Get(1)
		cell.PutValue_String(record.Product)
		cell, _ = row.Get(2)
		cell.PutValue_Int(record.Quantity)
		cell, _ = row.Get(3)
		cell.PutValue_Double(record.Price)
		cell, _ = row.Get(4)
		print(rowIndex + int32(i))
		cell.SetFormula_String(fmt.Sprintf("=C%d*D%d", rowIndex+int32(i)+1, rowIndex+int32(i)+1))
	}

}
func (reportGenerator *ReportGenerator) addSummaryStatistics(records []SaleRecord) {
	dataRowCount := len(records)
	summaryRow := dataRowCount + 4

	cells, _ := reportGenerator.worksheet.GetCells()
	cell, _ := cells.Get_Int_Int(int32(summaryRow), 0)
	cell.PutValue_String("Summary Statistics:")
	cell, _ = cells.Get_Int_Int(int32(summaryRow)+1, 0)
	cell.PutValue_String("Total Sales:")
	cell, _ = cells.Get_Int_Int(int32(summaryRow)+1, 1)
	cell.SetFormula_String(fmt.Sprintf("=SUM(E4:E%d)", dataRowCount+3))
	cell, _ = cells.Get_Int_Int(int32(summaryRow)+2, 0)
	cell.PutValue_String("Average Sales:")
	cell, _ = cells.Get_Int_Int(int32(summaryRow)+2, 1)
	cell.SetFormula_String(fmt.Sprintf("=AVERAGE(E4:E%d)", dataRowCount+3))
}

func (reportGenerator *ReportGenerator) createCharts(records []SaleRecord) {
	chartRow := int32(len(records) + 8)
	charts, _ := reportGenerator.worksheet.GetCharts()
	chartIndex, _ := charts.Add_ChartType_Int_Int_Int_Int(ChartType_Column, chartRow, 0, chartRow+10, 5)
	chart, _ := charts.Get_Int(chartIndex)
	nseries, _ := chart.GetNSeries()
	nseries.Add_String_Bool(fmt.Sprintf("=Sheet1!$E$4:$E$%d", len(records)+3), true)
	nseries.SetCategoryData(fmt.Sprintf("=Sheet1!$B$4:$B$%d", len(records)+3))
	chart_title, _ := chart.GetTitle()
	chart_title.SetText("Product sales comparison")
}

func (reportGenerator *ReportGenerator) formatReport() {
	// Set  column width
	cells, _ := reportGenerator.worksheet.GetCells()
	cells.SetColumnWidth(0, 15)
	cells.SetColumnWidth(1, 20)
	cells.SetColumnWidth(2, 10)
	cells.SetColumnWidth(3, 10)
	cells.SetColumnWidth(4, 12)

	// Set number Format
	numberStyle, _ := reportGenerator.workbook.CreateStyle()
	numberStyle.SetCustom_String("#,##0.00")
	columns, _ := cells.GetColumns()
	column, _ := columns.Get(3)
	column.SetStyle(numberStyle)
	column, _ = columns.Get(4)
	column.SetStyle(numberStyle)
}
func (reportGenerator *ReportGenerator) SaveReport(filePath string) {
	reportGenerator.workbook.Save_String(filePath)
}

func GenerateReport() {
	t1, _ := time.ParseInLocation("2006-01-02", "2024-01-01", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", "2024-01-02", time.Local)
	t3, _ := time.ParseInLocation("2006-01-02", "2024-01-03", time.Local)
	records := []SaleRecord{
		{t1, "Product A", 100, 25.5},
		{t2, "Product B", 150, 30.0},
		{t3, "Product A", 80, 25.5},
	}
	generator := NewReportGenerator()
	generator.GenerateSaleReport(records)
	generator.SaveReport("Data/Output/SalesAnalysisReport.xlsx")
}
