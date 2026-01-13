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
	"time"  	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum TimelineLevelType *****************/

// Represents the level type of <see cref="Timeline"/>
type TimelineLevelType int32

const(
// Year
TimelineLevelType_Year TimelineLevelType = 0 

// Quarter
TimelineLevelType_Quarter TimelineLevelType = 1 

// Month
TimelineLevelType_Month TimelineLevelType = 2 

// Day
TimelineLevelType_Day TimelineLevelType = 3 
)

func Int32ToTimelineLevelType(value int32)(TimelineLevelType ,error){
	switch value {
		case 0:  return TimelineLevelType_Year, nil  
		case 1:  return TimelineLevelType_Quarter, nil  
		case 2:  return TimelineLevelType_Month, nil  
		case 3:  return TimelineLevelType_Day, nil  
		default:
			return 0 ,fmt.Errorf("invalid TimelineLevelType value: %d", value)
	}
}
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Timeline_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to display the header.
// Returns:
//   bool  
func (instance *Timeline) GetShowHeader()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Timeline_GetShowHeader"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to display the header.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Timeline) SetShowHeader(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Timeline_SetShowHeader"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to display the selction label.
// Returns:
//   bool  
func (instance *Timeline) GetShowSelectionLabel()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Timeline_GetShowSelectionLabel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to display the selction label.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Timeline) SetShowSelectionLabel(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Timeline_SetShowSelectionLabel"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to display the time level.
// Returns:
//   bool  
func (instance *Timeline) GetShowTimeLevel()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Timeline_GetShowTimeLevel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to display the time level.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Timeline) SetShowTimeLevel(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Timeline_SetShowTimeLevel"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to display the horizontal ccroll bar.
// Returns:
//   bool  
func (instance *Timeline) GetShowHorizontalScrollbar()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Timeline_GetShowHorizontalScrollbar"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to display the horizontal ccroll bar.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *Timeline) SetShowHorizontalScrollbar(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("Timeline_SetShowHorizontalScrollbar"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the start date of the timespan scrolling position of this <see cref="Timeline"/>.
// Returns:
//   Date  
func (instance *Timeline) GetStartDate()  (time.Time,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("Timeline_GetStartDate"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  time.Unix(0, 0), err
	}
	result := time.Date(int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_year"), CGoReturnPtr.return_value).return_value ),time.Month(int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_month"), CGoReturnPtr.return_value).return_value )),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_day"), CGoReturnPtr.return_value).return_value ),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_hour"), CGoReturnPtr.return_value).return_value ),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_minute"), CGoReturnPtr.return_value).return_value ),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_second"), CGoReturnPtr.return_value).return_value ), 0, time.UTC) 

	return result, nil 
}
// Gets and sets the start date of the timespan scrolling position of this <see cref="Timeline"/>.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *Timeline) SetStartDate(value time.Time)  error {
	
	time_value := C.Get_Date( C.int(value.Year()), C.int(value.Month()) , C.int(value.Day()) , C.int(value.Hour()) , C.int(value.Minute()) , C.int(value.Second())  )

	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("Timeline_SetStartDate"), instance.ptr, time_value)
	C.Delete_GetDate( time_value)

	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The current time level of the Timeline.
// Returns:
//   int32  
func (instance *Timeline) GetCurrentLevel()  (TimelineLevelType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Timeline_GetCurrentLevel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTimelineLevelType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// The current time level of the Timeline.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Timeline) SetCurrentLevel(value TimelineLevelType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Timeline_SetCurrentLevel"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the time level at which the current selection was made for the Timeline.
// Returns:
//   int32  
func (instance *Timeline) GetSelectionLevel()  (TimelineLevelType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Timeline_GetSelectionLevel"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTimelineLevelType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the time level at which the current selection was made for the Timeline.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Timeline) SetSelectionLevel(value TimelineLevelType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Timeline_SetSelectionLevel"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the caption of this Timeline.
// Returns:
//   string  
func (instance *Timeline) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Timeline_GetCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the caption of this Timeline.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Timeline) SetCaption(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Timeline_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the <see cref="TimelineShape"/> object associated with this Timeline.
// Returns:
//   TimelineShape  
func (instance *Timeline) GetShape()  (*TimelineShape,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Timeline_GetShape"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Timeline_GetName"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Timeline_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteTimeline(timeline *Timeline){
	runtime.SetFinalizer(timeline, nil)
	C.Delete_CObject(C.CString("Delete_Timeline"),timeline.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TimelineCollection_IsNull"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("TimelineCollection_Get_Integer"), instance.ptr, C.int(index))
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("TimelineCollection_Get_String"), instance.ptr, C.CString(name))
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

	CGoReturnPtr := C.CellsGoFunctoinZZOO(C.CString("TimelineCollection_Add_PivotTable_Integer_Integer_String"), instance.ptr, pivot_ptr, C.int(row), C.int(column), C.CString(basefieldname))
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

	CGoReturnPtr := C.CellsGoFunctoinZZON(C.CString("TimelineCollection_Add_PivotTable_String_String"), instance.ptr, pivot_ptr, C.CString(destcellname), C.CString(basefieldname))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPI(C.CString("TimelineCollection_Add_PivotTable_Integer_Integer_Integer"), instance.ptr, pivot_ptr, C.int(row), C.int(column), C.int(basefieldindex))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPJ(C.CString("TimelineCollection_Add_PivotTable_String_Integer"), instance.ptr, pivot_ptr, C.CString(destcellname), C.int(basefieldindex))
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

	CGoReturnPtr := C.CellsGoFunctoinZZPK(C.CString("TimelineCollection_Add_PivotTable_Integer_Integer_PivotField"), instance.ptr, pivot_ptr, C.int(row), C.int(column), basefield_ptr)
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

	CGoReturnPtr := C.CellsGoFunctoinZZPL(C.CString("TimelineCollection_Add_PivotTable_String_PivotField"), instance.ptr, pivot_ptr, C.CString(destcellname), basefield_ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TimelineCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteTimelineCollection(timelinecollection *TimelineCollection){
	runtime.SetFinalizer(timelinecollection, nil)
	C.Delete_CObject(C.CString("Delete_TimelineCollection"),timelinecollection.ptr)
	timelinecollection.ptr = nil
}
