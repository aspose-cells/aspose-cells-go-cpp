// +build windows

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/win_x86_64" -L"${SRCDIR}/lib/win_x86_64" -lAspose.Cells.CWrapper
// #include <AsposeCellsCWrapper.h>
import "C"
import (
 
 	
	"errors"	
	"runtime"
	"unsafe" 
)

// Class Timeline 

// Summary description of Timeline View
// Due to MS Excel, Excel 2003 does not support Timeline
type Timeline struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Timeline) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.Timeline_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the caption of the specified Timeline.
// Returns:
//   string  
func (instance *Timeline) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.Timeline_GetCaption( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the caption of the specified Timeline.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Timeline) SetCaption(value string)  error {
	
	CGoReturnPtr := C.Timeline_SetCaption( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the <see cref="TimelineShape"/> object associated with this Timeline. Read-only.
// Returns:
//   TimelineShape  
func (instance *Timeline) GetShape()  (*TimelineShape,  error)  {
	
	CGoReturnPtr := C.Timeline_GetShape( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TimelineShape{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTimelineShape) 

	return result, nil 
}
// Returns or sets the name of the specified Timeline
// Returns:
//   string  
func (instance *Timeline) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.Timeline_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the name of the specified Timeline
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Timeline) SetName(value string)  error {
	
	CGoReturnPtr := C.Timeline_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteTimeline(timeline *Timeline){
	runtime.SetFinalizer(timeline, nil)
	C.Delete_Timeline(timeline.ptr)
	timeline.ptr = nil
}

// Class TimelineCollection 

// Specifies the collection of all the Timeline objects on the specified worksheet.
// Due to MS Excel, Excel 2003 does not support Timeline.
type TimelineCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TimelineCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.TimelineCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the Timeline by index.
// Parameters:
//   index - int32 
// Returns:
//   Timeline  
func (instance *TimelineCollection) Get_Int(index int32)  (*Timeline,  error)  {
	
	CGoReturnPtr := C.TimelineCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Timeline{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTimeline) 

	return result, nil 
}
// Gets the Timeline  by Timeline's name.
// Parameters:
//   name - string 
// Returns:
//   Timeline  
func (instance *TimelineCollection) Get_String(name string)  (*Timeline,  error)  {
	
	CGoReturnPtr := C.TimelineCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Timeline{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTimeline) 

	return result, nil 
}
// Add a new Timeline using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   row - int32 
//   column - int32 
//   baseFieldName - string 
// Returns:
//   int32  
func (instance *TimelineCollection) Add_PivotTable_Int_Int_String(pivot *PivotTable, row int32, column int32, basefieldname string)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.TimelineCollection_Add_PivotTable_Integer_Integer_String( instance.ptr, pivot_ptr, C.int(row), C.int(column), C.CString(basefieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Timeline using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   destCellName - string 
//   baseFieldName - string 
// Returns:
//   int32  
func (instance *TimelineCollection) Add_PivotTable_String_String(pivot *PivotTable, destcellname string, basefieldname string)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.TimelineCollection_Add_PivotTable_String_String( instance.ptr, pivot_ptr, C.CString(destcellname), C.CString(basefieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Timeline using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   row - int32 
//   column - int32 
//   baseFieldIndex - int32 
// Returns:
//   int32  
func (instance *TimelineCollection) Add_PivotTable_Int_Int_Int(pivot *PivotTable, row int32, column int32, basefieldindex int32)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.TimelineCollection_Add_PivotTable_Integer_Integer_Integer( instance.ptr, pivot_ptr, C.int(row), C.int(column), C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Timeline using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   destCellName - string 
//   baseFieldIndex - int32 
// Returns:
//   int32  
func (instance *TimelineCollection) Add_PivotTable_String_Int(pivot *PivotTable, destcellname string, basefieldindex int32)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}

	CGoReturnPtr := C.TimelineCollection_Add_PivotTable_String_Integer( instance.ptr, pivot_ptr, C.CString(destcellname), C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Timeline using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   row - int32 
//   column - int32 
//   baseField - PivotField 
// Returns:
//   int32  
func (instance *TimelineCollection) Add_PivotTable_Int_Int_PivotField(pivot *PivotTable, row int32, column int32, basefield *PivotField)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}
	var basefield_ptr unsafe.Pointer = nil
	if basefield != nil {
	  basefield_ptr =basefield.ptr
	}

	CGoReturnPtr := C.TimelineCollection_Add_PivotTable_Integer_Integer_PivotField( instance.ptr, pivot_ptr, C.int(row), C.int(column), basefield_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a new Timeline using PivotTable as data source
// Parameters:
//   pivot - PivotTable 
//   destCellName - string 
//   baseField - PivotField 
// Returns:
//   int32  
func (instance *TimelineCollection) Add_PivotTable_String_PivotField(pivot *PivotTable, destcellname string, basefield *PivotField)  (int32,  error)  {
	
	var pivot_ptr unsafe.Pointer = nil
	if pivot != nil {
	  pivot_ptr =pivot.ptr
	}
	var basefield_ptr unsafe.Pointer = nil
	if basefield != nil {
	  basefield_ptr =basefield.ptr
	}

	CGoReturnPtr := C.TimelineCollection_Add_PivotTable_String_PivotField( instance.ptr, pivot_ptr, C.CString(destcellname), basefield_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *TimelineCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.TimelineCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteTimelineCollection(timelinecollection *TimelineCollection){
	runtime.SetFinalizer(timelinecollection, nil)
	C.Delete_TimelineCollection(timelinecollection.ptr)
	timelinecollection.ptr = nil
}
