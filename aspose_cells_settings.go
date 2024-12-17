// +build windows

// Copyright (c) 2001-2024 Aspose Pty Ltd. All Rights Reserved.
// Powered by Aspose.Cells.
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

// Class PivotGlobalizationSettings 

// Represents the globalization settings for pivot tables.
type PivotGlobalizationSettings struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPivotGlobalizationSettings() ( *PivotGlobalizationSettings, error) {
	pivotglobalizationsettings := &PivotGlobalizationSettings{}
	CGoReturnPtr := C.New_PivotGlobalizationSettings()
	if CGoReturnPtr.error_no == 0 {
		pivotglobalizationsettings.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotglobalizationsettings, DeletePivotGlobalizationSettings)
		return pivotglobalizationsettings, nil
	} else {
		pivotglobalizationsettings.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotglobalizationsettings, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotGlobalizationSettings) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the text of "Total" label in the PivotTable.
// You need to override this method when the PivotTable contains two or more PivotFields in the data area.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfTotal()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfTotal( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of "Grand Total" label in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfGrandTotal()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfGrandTotal( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of "(Multiple Items)" label in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfMultipleItems()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfMultipleItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of "(All)" label in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfAll()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfAll( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text for specified protected name.
// Parameters:
//   protectedName - string 
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfProtectedName(protectedname string)  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfProtectedName( instance.ptr, C.CString(protectedname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of "Column Labels" label in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfColumnLabels()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfColumnLabels( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of "Row Labels" label in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfRowLabels()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfRowLabels( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of "(blank)" label in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfEmptyData()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfEmptyData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the the text of the value area field header in the PivotTable.
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfDataFieldHeader()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfDataFieldHeader( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Years".
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfYears()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfYears( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Get the local text of "Quarters".
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfQuarters()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfQuarters( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Months".
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfMonths()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfMonths( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Days".
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfDays()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfDays( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Hours".
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfHours()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfHours( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Minutes".
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfMinutes()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfMinutes( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Seconds"
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfSeconds()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfSeconds( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the local text of "Range"
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfRange()  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the text of <see cref="PivotFieldSubtotalType"/> type in the PivotTable.
// Parameters:
//   subTotalType - int32 
// Returns:
//   string  
func (instance *PivotGlobalizationSettings) GetTextOfSubTotal(subtotaltype PivotFieldSubtotalType)  (string,  error)  {
	CGoReturnPtr := C.PivotGlobalizationSettings_GetTextOfSubTotal( instance.ptr, C.int( int32(subtotaltype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotGlobalizationSettings(pivotglobalizationsettings *PivotGlobalizationSettings){
	runtime.SetFinalizer(pivotglobalizationsettings, nil)
	C.Delete_PivotGlobalizationSettings(pivotglobalizationsettings.ptr)
	pivotglobalizationsettings.ptr = nil
}
