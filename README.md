![](https://img.shields.io/badge/aspose.cells%20for%20Go%20via%20C++-v26.1.0-green?style=for-the-badge&logo=go) [![License](https://img.shields.io/github/license/aspose-cells/aspose-cells-go-cpp?style=for-the-badge&logo=rocket&logoColor=white)](https://github.com/aspose-cells/aspose-cells-go-cpp/blob/master/LICENSE?style=for-the-badge)  [![Product Page](https://img.shields.io/badge/Product-0288d1?style=for-the-badge&logo=Google-Chrome&logoColor=white)](https://products.aspose.com/cells/go-cpp/) [![Documentation](https://img.shields.io/badge/Documentation-388e3c?style=for-the-badge&logo=Hugo&logoColor=white)](https://docs.aspose.com/cells/go-cpp/) [![API Ref](https://img.shields.io/badge/Reference-f39c12?style=for-the-badge&logo=html5&logoColor=white)](https://reference.aspose.com/cells/go-cpp/) [![Examples](https://img.shields.io/badge/Examples-1565c0?style=for-the-badge&logo=Github&logoColor=white)](https://github.com/aspose-cells/aspose-cells-go-cpp/tree/main/samples) [![Blog](https://img.shields.io/badge/Blog-d32f2f?style=for-the-badge&logo=WordPress&logoColor=white)](https://blog.aspose.com/categories/aspose.cells-product-family/) [![Support](https://img.shields.io/badge/Support-7b1fa2?style=for-the-badge&logo=Discourse&logoColor=white)](https://forum.aspose.com/c/cells/9) ![GitHub commits since latest release (by date)](https://img.shields.io/github/commits-since/aspose-cells/aspose-cells-go-cpp/v26.1.0?style=for-the-badge)

# Aspose.Cells for Go via C++

Aspose.Cells for Go via C++ is a native Go library that creates, edits, calculates and converts Excel, PDF and more—no Office install, one import.

## Quick Start Guide

<a id="installationinyourproject"></a>

### Installation Aspose.Cells for Go via C++ package and running your code in your project

1. Create a directory for your project and a main.go file within. Add the following code to your main.go.

```Go

package main

import (
 . "github.com/aspose-cells/aspose-cells-go-cpp/v25"
 "fmt"
)

func main() {
 lic, _ := NewLicense()
 lic.SetLicense_String("YOUR_LICENSE_File_PATH")
 workbook, _ := NewWorkbook()
 worksheets, _ := workbook.GetWorksheets()
 worksheet, _ := worksheets.Get_Int(0)
 cells, _ := worksheet.GetCells()
 cell, _ := cells.Get_String("A1")
 cell.PutValue_String_Bool("Hello World!", true)
 style, _ := cell.GetStyle()
 style.SetPattern(BackgroundType_Solid)
 color, _ := NewColor()
 color.Set_Color_R(uint8(255))
 color.Set_Color_G(uint8(128))
 style.SetForegroundColor(color)
 cell.SetStyle_Style(style)
 workbook.Save_String("HELLO.pdf")

}

```

1. Initialize project go.mod

```bash

go mod init main

```

1. Fetch the dependencies for your project.

```bash

go mod tidy

```

If Aspose.Cells for Go via C++ is not installed in the development environment, Go will automatically install the package in the `$GOPATH` subdirectory.

1. Set your PATH to point to the shared libraries in Aspose.Cells for Go via C++ in your current command shell. Replace your_version with the version of Aspose.Cells for Go via C++ you are running.

```cmd

set PATH=%PATH%;%GOPATH%\github.com\aspose-cells\aspose-cells-go-cpp\v25@v26.1.0\lib\win_x86_64\

```

Or in your powershell

```powershell

$env:Path = $env:Path+ ";${env:GOPATH}\github.com\aspose-cells\aspose-cells-go-cpp\v25@v26.1.0\lib\win_x86_64\"

```

Or in your linux bash

```bash
export PATH=$PATH:$GOPATH/github.com/aspose-cells/aspose-cells-go-cpp/v25@v26.1.0/lib/linux_x86_64/

```

You may also copy the DLL files from the above path to the same place as your project executable.

1. Run your created application.

```bash

go run main.go

```

## Introduction

Aspose.Cells for Go via C++ is a native Go library designed for Go developers to programmatically create, manipulate, and convert spreadsheets without needing Office Automation or Microsoft Excel. It supports a variety of spreadsheet formats, including MS Excel 97-2003 (XLS), MS Excel 2007-2016 (XLSX, XLSM, XLSB), Open Office XML, and more. With Aspose.Cells for Go via C++, you can also extract images from worksheets and convert Excel files to PDF. The library further enables the creation and manipulation of charts and shapes. Additionally, it offers robust formula calculation capabilities, providing you with a comprehensive solution for spreadsheet management.

## Supported platforms

- *Windows x64*  
- *Linux x64*

## Environments and versions

- Go 1.16 or greater

## **Rich Features**

Aspose.Cells for Go via C++ offers a wide arrange of features for creating, converting and manipulating spreadsheets:

- Built-In Properties
- Custom Properties
- Themes
- Styles and Formatting
- Data Validation
- Conditional Formatting
- Hyperlink
- AutoFilter
- PageSetup
- Reading, Writing and Calculating Formulas
- Group Rows and Columns
- PivotTable
- Table

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

# Evaluate Aspose.Cells for Go via C++

You can use Aspose.Cells for Go via C++ free of cost for evaluation.The evaluation version provides almost all functionality of the product with certain limitations. The same evaluation version becomes licensed when you purchase a license and add a couple of lines of code to apply the license.
If you want to test Aspose.Cells for Go via C++ without evaluation version limitations, you can also try a 30-Day Temporary License. Please refer to  <a href="https://purchase.aspose.com/temporary-license/"> How to get a Temporary License</a>?

# Limitations of Evaluation version

The evaluation version of Aspose. Cells for Go via C++ provides complete product functionality. An evaluation watermark will be inserted when saving the file. And the evaluation version can open up to 100 files.

# Run Aspose.Cells for Go via C++ in production

A commercial license key is required in a production environment. Please contact us to <a href="https://purchase.aspose.com/buy">purchase a commercial license</a> if you want to publish application to the product server.

## How to Install Aspose.Cells for Go via C++

### **How to Install Aspose.Cells for Go via C++ Using the Go Command**

- From Go 1.16, you use `go install` command directly to install the Aspose.Cells for Go via C++ package. The command allows for the global installation of command-line tools without worrying about affecting existing project dependencies.

```go

go install github.com/aspose-cells/aspose-cells-go-cpp/v25@latest

```

- When you use `go get` command to install the Aspose.Cells for Go via C++ package, need exist the go.mod file in the current directory or any parent directory. see to [installation Aspose.Cells for Go via C++ package and running your code in your project](#installationinyourproject)

```go

go get github.com/aspose-cells/aspose-cells-go-cpp/v25@v26.1.0

```

### **How to build Aspose.Cells for Go via C++ from the Source Code Package and run samples**

1. Create a work directory for your project.

1. Get the source code about Aspose.Cells for Go via C++ source package

- Clone the source code from GitHub

```shell
git clone https://github.com/aspose-cells/aspose-cells-go-cpp.git cells-go-cpp
```

Or

- Download the source code package from the [Aspose.Cells for Go via C++ download page](https://downloads.aspose.com/cells/go-cpp/) and extract the ZIP file to cells-go-cpp folder in your work directory.

1. Navigate to cells-go-cpp folder and build source code

```bash
cd "your work directory"/cells-go-cpp/
go build .
```

1. **Set your PATH to point to the shared libraries in Aspose.Cells for Go via C++ in your current command shell.**

```cmd

set PATH=%PATH%;%YourProjectPath%\cells-go-cpp\lib\win_x86_64\

```

Or in your powershell

```powershell

$env:Path = $env:Path+ ";${YourProjectPath}\cells-go-cpp\lib\win_x86_64\"

```

Or in your bash

```bash
export PATH=$PATH:${YourProjectPath}/cells-go-cpp/lib/linux_x86_64/

```

1. Navigate to the `./samples` directory and run all sample.

```shell
go run .

```
