// +build linux

// Copyright (c) 2001-2024 Aspose Pty Ltd. All Rights Reserved.
// Powered by Aspose.Cells.
package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/linux_x86_64" -L"${SRCDIR}/lib/linux_x86_64" -lAspose.Cells.CWrapper
// #include <AsposeCellsCWrapper.h>
import "C"
import (
	"fmt"  
)

/**************Enum JsonExportHyperlinkType *****************/
// Represents type of exporting hyperlinks to json.
type JsonExportHyperlinkType int32

const(
// Export display string
JsonExportHyperlinkType_DisplayString JsonExportHyperlinkType = 0 

// Export url
JsonExportHyperlinkType_Address JsonExportHyperlinkType = 1 

// Export as html string.
JsonExportHyperlinkType_HtmlString JsonExportHyperlinkType = 2 
)

func Int32ToJsonExportHyperlinkType(value int32)(JsonExportHyperlinkType ,error){
	switch value {
		case 0:  return JsonExportHyperlinkType_DisplayString, nil  
		case 1:  return JsonExportHyperlinkType_Address, nil  
		case 2:  return JsonExportHyperlinkType_HtmlString, nil  
		default:
			return 0 ,fmt.Errorf("invalid JsonExportHyperlinkType value: %d", value)
	}
}
