// +build linux

package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./lib/linux_x86_64 -lAspose.Cells.CWrapper
// #include <AsposeCellsCWrapper.h>
import "C"
import (
	"fmt"  
)

/**************Enum AdjustFontSizeForRowType *****************/
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