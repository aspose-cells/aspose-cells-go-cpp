// +build linux

// Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
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

/**************Enum AdjustFontSizeForRowType *****************/

// Represents which kind of rows should be ajusted.
type AdjustFontSizeForRowType int32

const(
// No adjsut.
AdjustFontSizeForRowType_None AdjustFontSizeForRowType = 0 

// If the row is empty, change font size to fit row height.
AdjustFontSizeForRowType_EmptyRows AdjustFontSizeForRowType = 1 
)

func Int32ToAdjustFontSizeForRowType(value int32)(AdjustFontSizeForRowType ,error){
	switch value {
		case 0:  return AdjustFontSizeForRowType_None, nil  
		case 1:  return AdjustFontSizeForRowType_EmptyRows, nil  
		default:
			return 0 ,fmt.Errorf("invalid AdjustFontSizeForRowType value: %d", value)
	}
}

/**************Enum SlideViewType *****************/

// Represents the type when exporting to slides.
type SlideViewType int32

const(
// Exporting as view in MS Excel.
SlideViewType_View SlideViewType = 0 

// Exporting as printing.
SlideViewType_Print SlideViewType = 1 
)

func Int32ToSlideViewType(value int32)(SlideViewType ,error){
	switch value {
		case 0:  return SlideViewType_View, nil  
		case 1:  return SlideViewType_Print, nil  
		default:
			return 0 ,fmt.Errorf("invalid SlideViewType value: %d", value)
	}
}
