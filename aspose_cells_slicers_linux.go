// +build linux

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/linux_x86_64" -L"${SRCDIR}/lib/linux_x86_64" -lAspose.Cells.CWrapper
// #include <AsposeCellsCWrapper.h>
import "C"
import (
	"fmt"  
 	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum SlicerCacheCrossFilterType *****************/

// Represent the type of SlicerCacheCrossFilterType
type SlicerCacheCrossFilterType int32

const(
// The table style element of the slicer style for slicer items
// with no data is not applied to slicer items with no data, and slicer items
// with no data are not sorted separately in the list of slicer items in the slicer view
SlicerCacheCrossFilterType_None SlicerCacheCrossFilterType = 0 

// The table style element of the slicer style for slicer items with
// no data is applied to slicer items with no data, and slicer items
// with no data are sorted at the bottom in the list of slicer items in the slicer view
SlicerCacheCrossFilterType_ShowItemsWithDataAtTop SlicerCacheCrossFilterType = 1 

// The table style element of the slicer style for slicer items with no data
// is applied to slicer items with no data, and slicer items with no data
// are not sorted separately in the list of slicer items in the slicer view.
SlicerCacheCrossFilterType_ShowItemsWithNoData SlicerCacheCrossFilterType = 2 
)

func Int32ToSlicerCacheCrossFilterType(value int32)(SlicerCacheCrossFilterType ,error){
	switch value {
		case 0:  return SlicerCacheCrossFilterType_None, nil  
		case 1:  return SlicerCacheCrossFilterType_ShowItemsWithDataAtTop, nil  
		case 2:  return SlicerCacheCrossFilterType_ShowItemsWithNoData, nil  
		default:
			return 0 ,fmt.Errorf("invalid SlicerCacheCrossFilterType value: %d", value)
	}
}

/**************Enum SlicerCacheItemSortType *****************/

// Specify the sort type of SlicerCacheItem
type SlicerCacheItemSortType int32

const(
// Original data order.
SlicerCacheItemSortType_Natural SlicerCacheItemSortType = 0 

// Ascending sort type
SlicerCacheItemSortType_Ascending SlicerCacheItemSortType = 1 

// Descending sort type
SlicerCacheItemSortType_Descending SlicerCacheItemSortType = 2 
)

func Int32ToSlicerCacheItemSortType(value int32)(SlicerCacheItemSortType ,error){
	switch value {
		case 0:  return SlicerCacheItemSortType_Natural, nil  
		case 1:  return SlicerCacheItemSortType_Ascending, nil  
		case 2:  return SlicerCacheItemSortType_Descending, nil  
		default:
			return 0 ,fmt.Errorf("invalid SlicerCacheItemSortType value: %d", value)
	}
}

/**************Enum SlicerStyleType *****************/

// Specify the style of slicer view
type SlicerStyleType int32

const(
// built-in light style one
SlicerStyleType_SlicerStyleLight1 SlicerStyleType = 0 

// built-in light style two
SlicerStyleType_SlicerStyleLight2 SlicerStyleType = 1 

// built-in light style three
SlicerStyleType_SlicerStyleLight3 SlicerStyleType = 2 

// built-in light style four
SlicerStyleType_SlicerStyleLight4 SlicerStyleType = 3 

// built-in light style five
SlicerStyleType_SlicerStyleLight5 SlicerStyleType = 4 

// built-in light style six
SlicerStyleType_SlicerStyleLight6 SlicerStyleType = 5 

// built-in style other one
SlicerStyleType_SlicerStyleOther1 SlicerStyleType = 6 

// built-in style other two
SlicerStyleType_SlicerStyleOther2 SlicerStyleType = 7 

// built-in dark style one
SlicerStyleType_SlicerStyleDark1 SlicerStyleType = 8 

// built-in dark style tow
SlicerStyleType_SlicerStyleDark2 SlicerStyleType = 9 

// built-in dark style three
SlicerStyleType_SlicerStyleDark3 SlicerStyleType = 10 

// built-in dark style four
SlicerStyleType_SlicerStyleDark4 SlicerStyleType = 11 

// built-in dark style five
SlicerStyleType_SlicerStyleDark5 SlicerStyleType = 12 

// built-in dark style six
SlicerStyleType_SlicerStyleDark6 SlicerStyleType = 13 

// user-defined style, unsupported for now
SlicerStyleType_Custom SlicerStyleType = 14 
)

func Int32ToSlicerStyleType(value int32)(SlicerStyleType ,error){
	switch value {
		case 0:  return SlicerStyleType_SlicerStyleLight1, nil  
		case 1:  return SlicerStyleType_SlicerStyleLight2, nil  
		case 2:  return SlicerStyleType_SlicerStyleLight3, nil  
		case 3:  return SlicerStyleType_SlicerStyleLight4, nil  
		case 4:  return SlicerStyleType_SlicerStyleLight5, nil  
		case 5:  return SlicerStyleType_SlicerStyleLight6, nil  
		case 6:  return SlicerStyleType_SlicerStyleOther1, nil  
		case 7:  return SlicerStyleType_SlicerStyleOther2, nil  
		case 8:  return SlicerStyleType_SlicerStyleDark1, nil  
		case 9:  return SlicerStyleType_SlicerStyleDark2, nil  
		case 10:  return SlicerStyleType_SlicerStyleDark3, nil  
		case 11:  return SlicerStyleType_SlicerStyleDark4, nil  
		case 12:  return SlicerStyleType_SlicerStyleDark5, nil  
		case 13:  return SlicerStyleType_SlicerStyleDark6, nil  
		case 14:  return SlicerStyleType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid SlicerStyleType value: %d", value)
	}
}
// Class Slicer 

// summary description of Slicer View
type Slicer struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Slicer) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.Slicer_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds PivotTable connection.
// Parameters:
//   pivot - PivotTable 
// Returns:
//   void  
func (instance *Slicer) AddPivotConnection(pivot *PivotTable)  error {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.Slicer_AddPivotConnection( instance.ptr, pivot_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes PivotTable connection.
// Parameters:
//   pivot - PivotTable 
// Returns:
//   void  
func (instance *Slicer) RemovePivotConnection(pivot *PivotTable)  error {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.Slicer_RemovePivotConnection( instance.ptr, pivot_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the title of the current Slicer object.
// Returns:
//   string  
func (instance *Slicer) GetTitle()  (string,  error)  {
	
	CGoReturnPtr := C.Slicer_GetTitle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the title of the current Slicer object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Slicer) SetTitle(value string)  error {
	
	CGoReturnPtr := C.Slicer_SetTitle( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the descriptive (alternative) text string of the Slicer object.
// Returns:
//   string  
func (instance *Slicer) GetAlternativeText()  (string,  error)  {
	
	CGoReturnPtr := C.Slicer_GetAlternativeText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the descriptive (alternative) text string of the Slicer object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Slicer) SetAlternativeText(value string)  error {
	
	CGoReturnPtr := C.Slicer_SetAlternativeText( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the slicer object is printable.
// Returns:
//   bool  
func (instance *Slicer) IsPrintable()  (bool,  error)  {
	
	CGoReturnPtr := C.Slicer_IsPrintable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the slicer object is printable.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetIsPrintable(value bool)  error {
	
	CGoReturnPtr := C.Slicer_SetIsPrintable( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the slicer shape is locked.
// Returns:
//   bool  
func (instance *Slicer) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.Slicer_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the slicer shape is locked.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.Slicer_SetIsLocked( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the way the drawing object is attached to the cells below it.
// The property controls the placement of an object on a worksheet.
// Returns:
//   int32  
func (instance *Slicer) GetPlacement()  (PlacementType,  error)  {
	
	CGoReturnPtr := C.Slicer_GetPlacement( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPlacementType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the way the drawing object is attached to the cells below it.
// The property controls the placement of an object on a worksheet.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetPlacement(value PlacementType)  error {
	
	CGoReturnPtr := C.Slicer_SetPlacement( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether locking aspect ratio.
// Returns:
//   bool  
func (instance *Slicer) GetLockedAspectRatio()  (bool,  error)  {
	
	CGoReturnPtr := C.Slicer_GetLockedAspectRatio( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether locking aspect ratio.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetLockedAspectRatio(value bool)  error {
	
	CGoReturnPtr := C.Slicer_SetLockedAspectRatio( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified slicer can be moved or resized by using the user interface.
// Returns:
//   bool  
func (instance *Slicer) GetLockedPosition()  (bool,  error)  {
	
	CGoReturnPtr := C.Slicer_GetLockedPosition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the specified slicer can be moved or resized by using the user interface.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetLockedPosition(value bool)  error {
	
	CGoReturnPtr := C.Slicer_SetLockedPosition( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Refreshing the slicer.Meanwhile, Refreshing and Calculating  relative PivotTables.
// Returns:
//   void  
func (instance *Slicer) Refresh()  error {
	
	CGoReturnPtr := C.Slicer_Refresh( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the Shape object associated with the specified slicer. Read-only.
// Returns:
//   SlicerShape  
func (instance *Slicer) GetShape()  (*SlicerShape,  error)  {
	
	CGoReturnPtr := C.Slicer_GetShape( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SlicerShape{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSlicerShape) 

	return result, nil 
}
// Returns the SlicerCache object associated with the slicer. Read-only.
// Returns:
//   SlicerCache  
func (instance *Slicer) GetSlicerCache()  (*SlicerCache,  error)  {
	
	CGoReturnPtr := C.Slicer_GetSlicerCache( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SlicerCache{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSlicerCache) 

	return result, nil 
}
// Returns the <see cref="Worksheet"/> object which contains this slicer. Read-only.
// Returns:
//   Worksheet  
func (instance *Slicer) GetParent()  (*Worksheet,  error)  {
	
	CGoReturnPtr := C.Slicer_GetParent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Specify the type of Built-in slicer style
// the default type is SlicerStyleLight1
// Returns:
//   int32  
func (instance *Slicer) GetStyleType()  (SlicerStyleType,  error)  {
	
	CGoReturnPtr := C.Slicer_GetStyleType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSlicerStyleType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specify the type of Built-in slicer style
// the default type is SlicerStyleLight1
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetStyleType(value SlicerStyleType)  error {
	
	CGoReturnPtr := C.Slicer_SetStyleType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the name of the specified slicer
// Returns:
//   string  
func (instance *Slicer) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.Slicer_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the name of the specified slicer
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Slicer) SetName(value string)  error {
	
	CGoReturnPtr := C.Slicer_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the caption of the specified slicer.
// Returns:
//   string  
func (instance *Slicer) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.Slicer_GetCaption( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the caption of the specified slicer.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Slicer) SetCaption(value string)  error {
	
	CGoReturnPtr := C.Slicer_SetCaption( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets whether the header that displays the slicer Caption is visible
// the default value is true
// Returns:
//   bool  
func (instance *Slicer) GetCaptionVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.Slicer_GetCaptionVisible( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets whether the header that displays the slicer Caption is visible
// the default value is true
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetCaptionVisible(value bool)  error {
	
	CGoReturnPtr := C.Slicer_SetCaptionVisible( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the number of columns in the specified slicer.
// Returns:
//   int32  
func (instance *Slicer) GetNumberOfColumns()  (int32,  error)  {
	
	CGoReturnPtr := C.Slicer_GetNumberOfColumns( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the number of columns in the specified slicer.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetNumberOfColumns(value int32)  error {
	
	CGoReturnPtr := C.Slicer_SetNumberOfColumns( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the width of each column in the slicer, in unit of pixels.
// Returns:
//   int32  
func (instance *Slicer) GetColumnWidthPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.Slicer_GetColumnWidthPixel( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the width of each column in the slicer, in unit of pixels.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetColumnWidthPixel(value int32)  error {
	
	CGoReturnPtr := C.Slicer_SetColumnWidthPixel( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the width, in points, of each column in the slicer.
// Returns:
//   float64  
func (instance *Slicer) GetColumnWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.Slicer_GetColumnWidth( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the width, in points, of each column in the slicer.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Slicer) SetColumnWidth(value float64)  error {
	
	CGoReturnPtr := C.Slicer_SetColumnWidth( instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the height, in pixels, of each row in the specified slicer.
// Returns:
//   int32  
func (instance *Slicer) GetRowHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.Slicer_GetRowHeightPixel( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the height, in pixels, of each row in the specified slicer.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetRowHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.Slicer_SetRowHeightPixel( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the height, in points, of each row in the specified slicer.
// Returns:
//   float64  
func (instance *Slicer) GetRowHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.Slicer_GetRowHeight( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the height, in points, of each row in the specified slicer.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Slicer) SetRowHeight(value float64)  error {
	
	CGoReturnPtr := C.Slicer_SetRowHeight( instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSlicer(slicer *Slicer){
	runtime.SetFinalizer(slicer, nil)
	C.Delete_Slicer(slicer.ptr)
	slicer.ptr = nil
}

// Class SlicerCache 

// Represent summary description of slicer cache
type SlicerCache struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SlicerCache) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.SlicerCache_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets whether a slicer is participating in cross filtering with other slicers
// that share the same slicer cache, and how cross filtering is displayed. Read/write
// Returns:
//   int32  
func (instance *SlicerCache) GetCrossFilterType()  (SlicerCacheCrossFilterType,  error)  {
	
	CGoReturnPtr := C.SlicerCache_GetCrossFilterType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSlicerCacheCrossFilterType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Returns or sets whether a slicer is participating in cross filtering with other slicers
// that share the same slicer cache, and how cross filtering is displayed. Read/write
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SlicerCache) SetCrossFilterType(value SlicerCacheCrossFilterType)  error {
	
	CGoReturnPtr := C.SlicerCache_SetCrossFilterType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns whether the slicer associated with the specified slicer cache is based on an Non-OLAP data source. Read-only
// Returns:
//   bool  
func (instance *SlicerCache) GetList()  (bool,  error)  {
	
	CGoReturnPtr := C.SlicerCache_GetList( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns a SlicerCacheItem collection that contains the collection of all items in the slicer cache. Read-only
// Returns:
//   SlicerCacheItemCollection  
func (instance *SlicerCache) GetSlicerCacheItems()  (*SlicerCacheItemCollection,  error)  {
	
	CGoReturnPtr := C.SlicerCache_GetSlicerCacheItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SlicerCacheItemCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSlicerCacheItemCollection) 

	return result, nil 
}
// Returns the name of the slicer cache.
// Returns:
//   string  
func (instance *SlicerCache) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.SlicerCache_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the name of this slicer cache.
// Returns:
//   string  
func (instance *SlicerCache) GetSourceName()  (string,  error)  {
	
	CGoReturnPtr := C.SlicerCache_GetSourceName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCache(slicercache *SlicerCache){
	runtime.SetFinalizer(slicercache, nil)
	C.Delete_SlicerCache(slicercache.ptr)
	slicercache.ptr = nil
}

// Class SlicerCacheItem 

// Represent slicer data source item
type SlicerCacheItem struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SlicerCacheItem) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.SlicerCacheItem_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether the SlicerItem is selected or not.
// Returns:
//   bool  
func (instance *SlicerCacheItem) GetSelected()  (bool,  error)  {
	
	CGoReturnPtr := C.SlicerCacheItem_GetSelected( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether the SlicerItem is selected or not.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SlicerCacheItem) SetSelected(value bool)  error {
	
	CGoReturnPtr := C.SlicerCacheItem_SetSelected( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the label text for the slicer item. Read-only.
// Returns:
//   string  
func (instance *SlicerCacheItem) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.SlicerCacheItem_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCacheItem(slicercacheitem *SlicerCacheItem){
	runtime.SetFinalizer(slicercacheitem, nil)
	C.Delete_SlicerCacheItem(slicercacheitem.ptr)
	slicercacheitem.ptr = nil
}

// Class SlicerCacheItemCollection 

// Represent the collection of SlicerCacheItem
type SlicerCacheItemCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SlicerCacheItemCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.SlicerCacheItemCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the SlicerCacheItem object by index.
// Parameters:
//   index - int32 
// Returns:
//   SlicerCacheItem  
func (instance *SlicerCacheItemCollection) Get(index int32)  (*SlicerCacheItem,  error)  {
	
	CGoReturnPtr := C.SlicerCacheItemCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SlicerCacheItem{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSlicerCacheItem) 

	return result, nil 
}
// Gets the count of the SlicerCacheItem.
// Returns:
//   int32  
func (instance *SlicerCacheItemCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.SlicerCacheItemCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCacheItemCollection(slicercacheitemcollection *SlicerCacheItemCollection){
	runtime.SetFinalizer(slicercacheitemcollection, nil)
	C.Delete_SlicerCacheItemCollection(slicercacheitemcollection.ptr)
	slicercacheitemcollection.ptr = nil
}

// Class SlicerCollection 

// Specifies the collection of all the Slicer objects on the specified worksheet.
type SlicerCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SlicerCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.SlicerCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the Slicer by index.
// Parameters:
//   index - int32 
// Returns:
//   Slicer  
func (instance *SlicerCollection) Get_Int(index int32)  (*Slicer,  error)  {
	
	CGoReturnPtr := C.SlicerCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Slicer{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSlicer) 

	return result, nil 
}
// Gets the Slicer  by slicer's name.
// Parameters:
//   name - string 
// Returns:
//   Slicer  
func (instance *SlicerCollection) Get_String(name string)  (*Slicer,  error)  {
	
	CGoReturnPtr := C.SlicerCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Slicer{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSlicer) 

	return result, nil 
}
// Remove the specified Slicer
// Parameters:
//   slicer - Slicer 
// Returns:
//   void  
func (instance *SlicerCollection) Remove(slicer *Slicer)  error {
	
	var slicer_ptr unsafe.Pointer = nil
	if slicer != nil {
	  slicer_ptr =slicer.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Remove( instance.ptr, slicer_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Deletes the Slicer at the specified index
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *SlicerCollection) RemoveAt(index int32)  error {
	
	CGoReturnPtr := C.SlicerCollection_RemoveAt( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Clear all Slicers.
// Returns:
//   void  
func (instance *SlicerCollection) Clear()  error {
	
	CGoReturnPtr := C.SlicerCollection_Clear( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Add a new Slicer using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   destCellName - string 
//   baseFieldName - string 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_PivotTable_String_String(pivot *PivotTable, destcellname string, basefieldname string)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_PivotTable_String_String( instance.ptr, pivot_ptr, C.CString(destcellname), C.CString(basefieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   row - int32 
//   column - int32 
//   baseFieldName - string 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_PivotTable_Int_Int_String(pivot *PivotTable, row int32, column int32, basefieldname string)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_PivotTable_Integer_Integer_String( instance.ptr, pivot_ptr, C.int(row), C.int(column), C.CString(basefieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   row - int32 
//   column - int32 
//   baseFieldIndex - int32 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_PivotTable_Int_Int_Int(pivot *PivotTable, row int32, column int32, basefieldindex int32)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_PivotTable_Integer_Integer_Integer( instance.ptr, pivot_ptr, C.int(row), C.int(column), C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   destCellName - string 
//   baseFieldIndex - int32 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_PivotTable_String_Int(pivot *PivotTable, destcellname string, basefieldindex int32)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_PivotTable_String_Integer( instance.ptr, pivot_ptr, C.CString(destcellname), C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   row - int32 
//   column - int32 
//   baseField - PivotField 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_PivotTable_Int_Int_PivotField(pivot *PivotTable, row int32, column int32, basefield *PivotField)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}
	var basefield_ptr unsafe.Pointer = nil
	if basefield != nil {
	  basefield_ptr =basefield.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_PivotTable_Integer_Integer_PivotField( instance.ptr, pivot_ptr, C.int(row), C.int(column), basefield_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   destCellName - string 
//   baseField - PivotField 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_PivotTable_String_PivotField(pivot *PivotTable, destcellname string, basefield *PivotField)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}
	var basefield_ptr unsafe.Pointer = nil
	if basefield != nil {
	  basefield_ptr =basefield.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_PivotTable_String_PivotField( instance.ptr, pivot_ptr, C.CString(destcellname), basefield_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using ListObjet as data source
// Parameters:
//   table - ListObject 
//   index - int32 
//   destCellName - string 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_ListObject_Int_String(table *ListObject, index int32, destcellname string)  (int32,  error)  {
	
	var table_ptr unsafe.Pointer = nil
	if table != nil {
	  table_ptr =table.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_ListObject_Integer_String( instance.ptr, table_ptr, C.int(index), C.CString(destcellname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using ListObjet as data source
// Parameters:
//   table - ListObject 
//   listColumn - ListColumn 
//   destCellName - string 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_ListObject_ListColumn_String(table *ListObject, listcolumn *ListColumn, destcellname string)  (int32,  error)  {
	
	var table_ptr unsafe.Pointer = nil
	if table != nil {
	  table_ptr =table.ptr
	}
	var listcolumn_ptr unsafe.Pointer = nil
	if listcolumn != nil {
	  listcolumn_ptr =listcolumn.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_ListObject_ListColumn_String( instance.ptr, table_ptr, listcolumn_ptr, C.CString(destcellname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Slicer using ListObjet as data source
// Parameters:
//   table - ListObject 
//   listColumn - ListColumn 
//   row - int32 
//   column - int32 
// Returns:
//   int32  
func (instance *SlicerCollection) Add_ListObject_ListColumn_Int_Int(table *ListObject, listcolumn *ListColumn, row int32, column int32)  (int32,  error)  {
	
	var table_ptr unsafe.Pointer = nil
	if table != nil {
	  table_ptr =table.ptr
	}
	var listcolumn_ptr unsafe.Pointer = nil
	if listcolumn != nil {
	  listcolumn_ptr =listcolumn.ptr
	}

	CGoReturnPtr := C.SlicerCollection_Add_ListObject_ListColumn_Integer_Integer( instance.ptr, table_ptr, listcolumn_ptr, C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *SlicerCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.SlicerCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCollection(slicercollection *SlicerCollection){
	runtime.SetFinalizer(slicercollection, nil)
	C.Delete_SlicerCollection(slicercollection.ptr)
	slicercollection.ptr = nil
}
