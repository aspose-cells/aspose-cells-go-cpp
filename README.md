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
|TabDelimited|Tab-delimited text file, same with TSV file.|&radic;|&radic;|
|[TXT](https://docs.fileformat.com/word-processing/txt/)|Delimited plain text file.|&radic;|&radic;|
|[HTML](https://docs.fileformat.com/web/html/)|HTML format.|&radic;|&radic;|
|[MHTML](https://docs.fileformat.com/web/mhtml/)|MHTML file.|&radic;|&radic;|
|[ODS](https://docs.fileformat.com/spreadsheet/ods/)|ODS (OpenDocument Spreadsheet).|&radic;|&radic;|
|SpreadsheetML|Excel 2003 XML file.|&radic;|&radic;|
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

## Running Aspose.Cells for Go via C++ in your project

1. Import `github.com/aspose-cells/aspose-cells-go-cpp/v2` into your project
   a. On **Windows**, you will have to locate the DLLs for running the project and append them to your path.
   ```
   set PATH=%GOPATH%/pkg/mod/github.com/aspose-cells/aspose-cells-go-cpp/v2@your_version/libs/win/Lib/win_x86_64
   ```
   You may also copy these directly to your project directory.

2. Create a main.go in your project directory
```go

package main

import (
 . github.com/aspose-cells/aspose-cells-go-cpp/v2
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
