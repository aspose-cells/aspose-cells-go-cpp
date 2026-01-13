// +build windows

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/win_x86_64" -L"${SRCDIR}/lib/win_x86_64" -lAspose.Cells.CWrapper
// #include <CellsFunctionMap.h>
import "C"
import (
	"fmt"  
 	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum ItemsWithNoDataShowMode *****************/

// Represent the type how to show items with no data for slicer.
type ItemsWithNoDataShowMode int32

const(
// Hide items with no data.
ItemsWithNoDataShowMode_None ItemsWithNoDataShowMode = 0 

// Show items with no data last.
ItemsWithNoDataShowMode_Last ItemsWithNoDataShowMode = 1 

// Show items with no data with natural order.
ItemsWithNoDataShowMode_Natural ItemsWithNoDataShowMode = 2 
)

func Int32ToItemsWithNoDataShowMode(value int32)(ItemsWithNoDataShowMode ,error){
	switch value {
		case 0:  return ItemsWithNoDataShowMode_None, nil  
		case 1:  return ItemsWithNoDataShowMode_Last, nil  
		case 2:  return ItemsWithNoDataShowMode_Natural, nil  
		default:
			return 0 ,fmt.Errorf("invalid ItemsWithNoDataShowMode value: %d", value)
	}
}

/**************Enum SlicerCacheCrossFilterType *****************/

// Represent the type how to show items with no data for slicer.
type SlicerCacheCrossFilterType int32

const(
// Hide items with no data.
SlicerCacheCrossFilterType_None SlicerCacheCrossFilterType = 0 

// Show items with data at top.
SlicerCacheCrossFilterType_ShowItemsWithDataAtTop SlicerCacheCrossFilterType = 1 

// Show items with no data with natural order.
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Slicer_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the type of sorting items.
// Returns:
//   int32  
func (instance *Slicer) GetSortOrderType()  (SortOrder,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Slicer_GetSortOrderType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSortOrder(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates the type of sorting items.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetSortOrderType(value SortOrder)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Slicer_SetSortOrderType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show items deteleted from the data source.
// Returns:
//   bool  
func (instance *Slicer) GetShowMissing()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Slicer_GetShowMissing"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show items deteleted from the data source.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetShowMissing(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Slicer_SetShowMissing"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show items deteleted from the data source.
// Returns:
//   int32  
func (instance *Slicer) GetShowTypeOfItemsWithNoData()  (ItemsWithNoDataShowMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Slicer_GetShowTypeOfItemsWithNoData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToItemsWithNoDataShowMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether to show items deteleted from the data source.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetShowTypeOfItemsWithNoData(value ItemsWithNoDataShowMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Slicer_SetShowTypeOfItemsWithNoData"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show all items even if there are no data or they are deleted.
// Default value is true;
// Returns:
//   bool  
func (instance *Slicer) GetShowAllItems()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Slicer_GetShowAllItems"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show all items even if there are no data or they are deleted.
// Default value is true;
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetShowAllItems(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Slicer_SetShowAllItems"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
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

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Slicer_AddPivotConnection"), instance.ptr, pivot_ptr)
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

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("Slicer_RemovePivotConnection"), instance.ptr, pivot_ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Slicer_GetLockedPosition"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Slicer_SetLockedPosition"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Refreshing the slicer.
// Meanwhile, Refreshing and Calculating PivotTables which this slicer based on.
// Returns:
//   void  
func (instance *Slicer) Refresh()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("Slicer_Refresh"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Slicer_GetShape"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Slicer_GetSlicerCache"), instance.ptr)
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
func (instance *Slicer) GetWorksheet()  (*Worksheet,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Slicer_GetWorksheet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Specify the type of Built-in slicer style.
// The default type is SlicerStyleLight1.
// Returns:
//   int32  
func (instance *Slicer) GetStyleType()  (SlicerStyleType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Slicer_GetStyleType"), instance.ptr)
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
// Specify the type of Built-in slicer style.
// The default type is SlicerStyleLight1.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetStyleType(value SlicerStyleType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Slicer_SetStyleType"), instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Slicer_GetName"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Slicer_SetName"), instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Slicer_GetCaption"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Slicer_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the zero-based index of the first slicer item.
// Returns:
//   int32  
func (instance *Slicer) GetFirstItemIndex()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Slicer_GetFirstItemIndex"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the zero-based index of the first slicer item.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetFirstItemIndex(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Slicer_SetFirstItemIndex"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the header of the slicer is visible.
// The default value is true
// Returns:
//   bool  
func (instance *Slicer) GetShowCaption()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Slicer_GetShowCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the header of the slicer is visible.
// The default value is true
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Slicer) SetShowCaption(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Slicer_SetShowCaption"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the number of columns in the specified slicer.
// The default value is 1.
// Returns:
//   int32  
func (instance *Slicer) GetNumberOfColumns()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Slicer_GetNumberOfColumns"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the number of columns in the specified slicer.
// The default value is 1.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetNumberOfColumns(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Slicer_SetNumberOfColumns"), instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Slicer_GetColumnWidthPixel"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Slicer_SetColumnWidthPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the width of each column in the slicer in unit of points.
// Returns:
//   float64  
func (instance *Slicer) GetColumnWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Slicer_GetColumnWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the width of each column in the slicer in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Slicer) SetColumnWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Slicer_SetColumnWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the height of each row in the specified slicer, in unit of pixels.
// Returns:
//   int32  
func (instance *Slicer) GetRowHeightPixel()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("Slicer_GetRowHeightPixel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the height of each row in the specified slicer, in unit of pixels.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Slicer) SetRowHeightPixel(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("Slicer_SetRowHeightPixel"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the height of each row in the specified slicer in unit of points.
// Returns:
//   float64  
func (instance *Slicer) GetRowHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("Slicer_GetRowHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the height of each row in the specified slicer in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *Slicer) SetRowHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("Slicer_SetRowHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSlicer(slicer *Slicer){
	runtime.SetFinalizer(slicer, nil)
	C.Delete_CObject(C.CString("Delete_Slicer"),slicer.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SlicerCache_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the slicer associated with the specified slicer cache is based on an Non-OLAP data source.
// Returns:
//   bool  
func (instance *SlicerCache) GetList()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SlicerCache_GetList"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SlicerCache_GetSlicerCacheItems"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SlicerCache_GetName"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SlicerCache_GetSourceName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCache(slicercache *SlicerCache){
	runtime.SetFinalizer(slicercache, nil)
	C.Delete_CObject(C.CString("Delete_SlicerCache"),slicercache.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SlicerCacheItem_IsNull"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SlicerCacheItem_GetSelected"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SlicerCacheItem_SetSelected"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the label text for the slicer item.
// Returns:
//   string  
func (instance *SlicerCacheItem) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SlicerCacheItem_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCacheItem(slicercacheitem *SlicerCacheItem){
	runtime.SetFinalizer(slicercacheitem, nil)
	C.Delete_CObject(C.CString("Delete_SlicerCacheItem"),slicercacheitem.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SlicerCacheItemCollection_IsNull"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SlicerCacheItemCollection_Get"), instance.ptr, C.int(index))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SlicerCacheItemCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCacheItemCollection(slicercacheitemcollection *SlicerCacheItemCollection){
	runtime.SetFinalizer(slicercacheitemcollection, nil)
	C.Delete_CObject(C.CString("Delete_SlicerCacheItemCollection"),slicercacheitemcollection.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SlicerCollection_IsNull"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SlicerCollection_Get_Integer"), instance.ptr, C.int(index))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("SlicerCollection_Get_String"), instance.ptr, C.CString(name))
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

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SlicerCollection_Remove"), instance.ptr, slicer_ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SlicerCollection_RemoveAt"), instance.ptr, C.int(index))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("SlicerCollection_Clear"), instance.ptr)
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

	CGoReturnPtr := C.CellsGoFunctoinZZON(C.CString("SlicerCollection_Add_PivotTable_String_String"), instance.ptr, pivot_ptr, C.CString(destcellname), C.CString(basefieldname))
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

	CGoReturnPtr := C.CellsGoFunctoinZZOO(C.CString("SlicerCollection_Add_PivotTable_Integer_Integer_String"), instance.ptr, pivot_ptr, C.int(row), C.int(column), C.CString(basefieldname))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPI(C.CString("SlicerCollection_Add_PivotTable_Integer_Integer_Integer"), instance.ptr, pivot_ptr, C.int(row), C.int(column), C.int(basefieldindex))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPJ(C.CString("SlicerCollection_Add_PivotTable_String_Integer"), instance.ptr, pivot_ptr, C.CString(destcellname), C.int(basefieldindex))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPK(C.CString("SlicerCollection_Add_PivotTable_Integer_Integer_PivotField"), instance.ptr, pivot_ptr, C.int(row), C.int(column), basefield_ptr)
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

	CGoReturnPtr := C.CellsGoFunctoinZZPL(C.CString("SlicerCollection_Add_PivotTable_String_PivotField"), instance.ptr, pivot_ptr, C.CString(destcellname), basefield_ptr)
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

	CGoReturnPtr := C.CellsGoFunctoinZZPM(C.CString("SlicerCollection_Add_ListObject_Integer_String"), instance.ptr, table_ptr, C.int(index), C.CString(destcellname))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPN(C.CString("SlicerCollection_Add_ListObject_ListColumn_String"), instance.ptr, table_ptr, listcolumn_ptr, C.CString(destcellname))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPO(C.CString("SlicerCollection_Add_ListObject_ListColumn_Integer_Integer"), instance.ptr, table_ptr, listcolumn_ptr, C.int(row), C.int(column))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SlicerCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSlicerCollection(slicercollection *SlicerCollection){
	runtime.SetFinalizer(slicercollection, nil)
	C.Delete_CObject(C.CString("Delete_SlicerCollection"),slicercollection.ptr)
	slicercollection.ptr = nil
}
