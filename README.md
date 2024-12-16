# Aspose.Cells for Go via C++

Aspose.Cells for Go via C++ is a native Go library designed for Go developers to programmatically create, manipulate, and convert spreadsheets without needing Office Automation or Microsoft Excel. It supports a variety of spreadsheet formats, including MS Excel 97-2003 (XLS), MS Excel 2007-2016 (XLSX, XLSM, XLSB), Open Office XML, and more. With Aspose.Cells for Go via C++, you can also extract images from worksheets and convert Excel files to PDF. The library further enables the creation and manipulation of charts and shapes. Additionally, it offers robust formula calculation capabilities, providing you with a comprehensive solution for spreadsheet management.

## Supported platforms

- *Windows x64*  
- *Linux x64*

## Environments and versions

- Go 1.13 or greater

## Support file format

|**Format**|**Description**|**Load**|**Save**|
| :- | :- | :- | :- |
|[XLS](https://docs.fileformat.com/spreadsheet/xls/)|Excel 95/5.0 - 2003 Workbook.|&radic;|&radic;|
|[XLSX](https://docs.fileformat.com/spreadsheet/xlsx/)|Office Open XML SpreadsheetML Workbook or template file, with or without macros.|&radic;|&radic;|
|[XLSB](https://docs.fileformat.com/spreadsheet/xlsb/)|Excel Binary Workbook.|&radic;|&radic;|
|[XLSM](https://docs.fileformat.com/spreadsheet/xlsm/)|Excel Macro-Enabled Workbook.|&radic;|&radic;|
|[XLT](https://docs.fileformat.com/spreadsheet/xlt/)|Excel 97 - Excel 2003 Template.|&radic;|&radic;|
|[XLTX](https://docs.fileformat.com/spreadsheet/xltx/)|Excel Template.|&radic;|&radic;|
|[XLTM](https://docs.fileformat.com/spreadsheet/xltm/)|Excel Macro-Enabled Template.|&radic;|&radic;|
|[XLAM](https://docs.fileformat.com/spreadsheet/xlam/)|An Excel Macro-Enabled Add-In file that's used to add new functions to Excel.| |&radic;|
|[CSV](https://docs.fileformat.com/spreadsheet/csv/)|CSV (Comma Separated Value) file.|&radic;|&radic;|
|[TSV](https://docs.fileformat.com/spreadsheet/tsv/)|TSV (Tab-separated values) file.|&radic;|&radic;|
|[TXT](https://docs.fileformat.com/word-processing/txt/)|Delimited plain text file.|&radic;|&radic;|
|[HTML](https://docs.fileformat.com/web/html/)|HTML format.|&radic;|&radic;|
|[MHTML](https://docs.fileformat.com/web/mhtml/)|MHTML file.|&radic;|&radic;|
|[ODS](https://docs.fileformat.com/spreadsheet/ods/)|ODS (OpenDocument Spreadsheet).|&radic;|&radic;|
|[JSON](https://docs.fileformat.com/web/json/)|JavaScript Object Notation|&radic;|&radic;|
|[DIF](https://docs.fileformat.com/spreadsheet/dif/)|Data Interchange Format.| |&radic;|
|[PDF](https://docs.fileformat.com/pdf/)|Adobe Portable Document Format.| |&radic;|
|[XPS](https://docs.fileformat.com/page-description-language/xps/)|XML Paper Specification Format.| |&radic;|
|[SVG](https://docs.fileformat.com/page-description-language/svg/)|Scalable Vector Graphics Format.| |&radic;|
|[TIFF](https://docs.fileformat.com/image/tiff/)|Tagged Image File Format| |&radic;|
|[PNG](https://docs.fileformat.com/image/png/)|Portable Network Graphics Format| |&radic;|
|[BMP](https://docs.fileformat.com/image/bmp/)|Bitmap Image Format| |&radic;|
|[EMF](https://docs.fileformat.com/image/emf/)|Enhanced metafile Format| |&radic;|
|[JPEG](https://docs.fileformat.com/image/jpeg/)|JPEG is a type of image format that is saved using the method of lossy compression.| |&radic;|
|[GIF](https://docs.fileformat.com/image/gif/)|Graphical Interchange Format| |&radic;|
|[MARKDOWN](https://docs.fileformat.com/word-processing/md/)|Represents a markdown document.| |&radic;|
|[SXC](https://docs.fileformat.com/spreadsheet/sxc/)|An XML based format used by OpenOffice and StarOffice|&radic;|&radic;|
|[FODS](https://docs.fileformat.com/spreadsheet/fods/)|This is an Open Document format stored as flat XML.|&radic;|&radic;|
|[DOCX](https://docs.fileformat.com/word-processing/docx/)|A well-known format for Microsoft Word documents that is a combination of XML and binary files.||&radic;|
|[PPTX](https://docs.fileformat.com/presentation/pptx/)|The PPTX format is based on the Microsoft PowerPoint open XML presentation file format.||&radic;|

# Run Aspose.Cells for Go via C++ in production

A commercial license key is required for use in a production environment. Please <a href="https://purchase.aspose.com/buy">contact us to purchase a commercial license</a> if you do not have a valid license key.

## Quick Start Guide

### Running Aspose.Cells for Go via C++ in your project

1. Import `github.com/aspose-cells/aspose-cells-go-cpp/v24` into your project
   a. On **Windows**, you will have to locate the DLLs for running the project and append them to your path.

   ```
       $env:PATH = $env:Path + ";$env:GOPATH\github.com\aspose-cells\aspose-cells-go-cpp\v24@v24.12.0\lib\win_x86_64" 
   ```

   b. On **Linux**, you will have to locate the DLLs for running the project and append them to your path.

   ```
   set PATH=%GOPATH%/github.com/aspose-cells/aspose-cells-go-cpp/v24@your_version/libs/win/Lib/linux_x86_64
   ```

   You may also copy these directly to your project directory.

2. Create a main.go in your project directory

```go

package main

import (
 . github.com/aspose-cells/aspose-cells-go-cpp/v24
)

func main() {
 lic, _ := NewLicense()
 lic.SetLicense_String(os.Getenv("LicensePath"))
    workbook, _ := NewWorkbook()
    worksheets, _ := workbook.GetWorksheets()
    worksheet, _ := worksheets.Get_Int(0)
    cells, _ := worksheet.GetCells()
    cell, _ := cells.Get_String("A1")
    cell.PutValue_Int(5)
    cell, _ = cells.Get_String("A2")
    cell.PutValue_String("Hollo World")
    workbook.Save_String("HolloWorld.xlsx")
}

```

### Create a table in excel

```go

package main

import (
 . github.com/aspose-cells/aspose-cells-go-cpp/v24
)

func main() {
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

 listObjects, _ := worksheet.GetListObjects()
 index, _ := listObjects.Add_String_String_Bool("A1", "F6", true)
 listObject, _ := listObjects.Get_Int(index)
 listObject.SetShowHeaderRow(true)
 listObject.SetTableStyleType(TableStyleType_TableStyleMedium10)
 listObject.SetShowTotals(true)

 workbook.Save_String("CreateTable.xlsx")
}
```

### Create bubble chart in excel

```go

package main

import (
 . github.com/aspose-cells/aspose-cells-go-cpp/v24
)

func main() {
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
   nSeries, _ := chart.GetNSeries()
   nSeries.Add_String_Bool("B1:D1", true)
   series, _ := nSeries.Get(0)
   series.SetBubbleSizes("B2:D2")
   series.SetXValues("B3:D3")
   series.SetValues("B1:D1")
   workbook.Save_String("CreateBubbleChart.xlsx")

}

```
