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

// Class ConversionUtility 

// Represents utility to convert files to other formats.
type ConversionUtility struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ConversionUtility) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.ConversionUtility_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts Excel files to other formats.
// Parameters:
//   source - string 
//   saveAs - string 
// Returns:
//   void  
func ConversionUtility_Convert_String_String(source string, saveas string)  error {
	
	CGoReturnPtr := C.ConversionUtility_Convert_String_String(C.CString(source), C.CString(saveas))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts Excel files to other formats.
// Parameters:
//   source - string 
//   loadOptions - LoadOptions 
//   saveAs - string 
//   saveOptions - SaveOptions 
// Returns:
//   void  
func ConversionUtility_Convert_String_LoadOptions_String_SaveOptions(source string, loadoptions *LoadOptions, saveas string, saveoptions *SaveOptions)  error {
	
	CGoReturnPtr := C.ConversionUtility_Convert_String_LoadOptions_String_SaveOptions(C.CString(source), loadoptions.ptr, C.CString(saveas), saveoptions.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteConversionUtility(conversionutility *ConversionUtility){
	runtime.SetFinalizer(conversionutility, nil)
	C.Delete_ConversionUtility(conversionutility.ptr)
	conversionutility.ptr = nil
}

// Class ExportRangeToJsonOptions 

// Indicates the options that exporting range to json.
type ExportRangeToJsonOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewExportRangeToJsonOptions() ( *ExportRangeToJsonOptions, error) {
	exportrangetojsonoptions := &ExportRangeToJsonOptions{}
	CGoReturnPtr := C.New_ExportRangeToJsonOptions()
	if CGoReturnPtr.error_no == 0 {
		exportrangetojsonoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(exportrangetojsonoptions, DeleteExportRangeToJsonOptions)
		return exportrangetojsonoptions, nil
	} else {
		exportrangetojsonoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return exportrangetojsonoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ExportRangeToJsonOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the range contains header row.
// Returns:
//   bool  
func (instance *ExportRangeToJsonOptions) GetHasHeaderRow()  (bool,  error)  {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_GetHasHeaderRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the range contains header row.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExportRangeToJsonOptions) SetHasHeaderRow(value bool)  error {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_SetHasHeaderRow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Exports the string value of the cells to json.
// Returns:
//   bool  
func (instance *ExportRangeToJsonOptions) GetExportAsString()  (bool,  error)  {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_GetExportAsString( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Exports the string value of the cells to json.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExportRangeToJsonOptions) SetExportAsString(value bool)  error {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_SetExportAsString( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting empty cells as null.
// Returns:
//   bool  
func (instance *ExportRangeToJsonOptions) GetExportEmptyCells()  (bool,  error)  {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_GetExportEmptyCells( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting empty cells as null.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExportRangeToJsonOptions) SetExportEmptyCells(value bool)  error {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_SetExportEmptyCells( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates the indent.
// Returns:
//   string  
func (instance *ExportRangeToJsonOptions) GetIndent()  (string,  error)  {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_GetIndent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates the indent.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExportRangeToJsonOptions) SetIndent(value string)  error {
	
	CGoReturnPtr := C.ExportRangeToJsonOptions_SetIndent( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteExportRangeToJsonOptions(exportrangetojsonoptions *ExportRangeToJsonOptions){
	runtime.SetFinalizer(exportrangetojsonoptions, nil)
	C.Delete_ExportRangeToJsonOptions(exportrangetojsonoptions.ptr)
	exportrangetojsonoptions.ptr = nil
}

// Class JsonLayoutOptions 

// Represents the options of json layout type.
type JsonLayoutOptions struct {
	ptr unsafe.Pointer
}

// Constructor of loading JSON layout options.
func NewJsonLayoutOptions() ( *JsonLayoutOptions, error) {
	jsonlayoutoptions := &JsonLayoutOptions{}
	CGoReturnPtr := C.New_JsonLayoutOptions()
	if CGoReturnPtr.error_no == 0 {
		jsonlayoutoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(jsonlayoutoptions, DeleteJsonLayoutOptions)
		return jsonlayoutoptions, nil
	} else {
		jsonlayoutoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return jsonlayoutoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *JsonLayoutOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Processes Array as table.
// Returns:
//   bool  
func (instance *JsonLayoutOptions) GetArrayAsTable()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_GetArrayAsTable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Processes Array as table.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetArrayAsTable(value bool)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetArrayAsTable( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether ignoring null value.
// Returns:
//   bool  
func (instance *JsonLayoutOptions) GetIgnoreNull()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_GetIgnoreNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether ignoring null value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetIgnoreNull(value bool)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetIgnoreNull( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Ingores titles of attributes
// Returns:
//   bool  
func (instance *JsonLayoutOptions) GetIgnoreTitle()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_GetIgnoreTitle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Ingores titles of attributes
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetIgnoreTitle(value bool)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetIgnoreTitle( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether converting the string in json to numeric or date value.
// Returns:
//   bool  
func (instance *JsonLayoutOptions) GetConvertNumericOrDate()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_GetConvertNumericOrDate( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether converting the string in json to numeric or date value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetConvertNumericOrDate(value bool)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetConvertNumericOrDate( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the format of numeric value.
// Returns:
//   string  
func (instance *JsonLayoutOptions) Get_NumberFormat()  (string,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_Get_NumberFormat( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the format of numeric value.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetNumberFormat(value string)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetNumberFormat( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the format of date value.
// Returns:
//   string  
func (instance *JsonLayoutOptions) Get_DateFormat()  (string,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_Get_DateFormat( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the format of date value.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetDateFormat(value string)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetDateFormat( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the style of the title.
// Returns:
//   Style  
func (instance *JsonLayoutOptions) GetTitleStyle()  (*Style,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_GetTitleStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Gets and sets the style of the title.
// Parameters:
//   value - Style 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetTitleStyle(value *Style)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetTitleStyle( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether keeping schema of this json.
// Returns:
//   bool  
func (instance *JsonLayoutOptions) GetKeptSchema()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonLayoutOptions_GetKeptSchema( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether keeping schema of this json.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *JsonLayoutOptions) SetKeptSchema(value bool)  error {
	
	CGoReturnPtr := C.JsonLayoutOptions_SetKeptSchema( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteJsonLayoutOptions(jsonlayoutoptions *JsonLayoutOptions){
	runtime.SetFinalizer(jsonlayoutoptions, nil)
	C.Delete_JsonLayoutOptions(jsonlayoutoptions.ptr)
	jsonlayoutoptions.ptr = nil
}

// Class JsonUtility 

// Represents the utility class of processing json.
type JsonUtility struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewJsonUtility() ( *JsonUtility, error) {
	jsonutility := &JsonUtility{}
	CGoReturnPtr := C.New_JsonUtility()
	if CGoReturnPtr.error_no == 0 {
		jsonutility.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(jsonutility, DeleteJsonUtility)
		return jsonutility, nil
	} else {
		jsonutility.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return jsonutility, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *JsonUtility) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.JsonUtility_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Import the json string.
// Parameters:
//   json - string 
//   cells - Cells 
//   row - int32 
//   column - int32 
//   option - JsonLayoutOptions 
// Returns:
//   []int32_t  
func JsonUtility_ImportData(json string, cells *Cells, row int32, column int32, option *JsonLayoutOptions)  ([]int32,  error)  {
	
	CGoReturnPtr := C.JsonUtility_ImportData(C.CString(json), cells.ptr, C.int(row), C.int(column), option.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]int32, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   cObject := *(*C.int)(unsafe.Pointer( uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset))
	   goObject :=int32(cObject)
	   result[i] = goObject
	}
	 

	return result, nil 
}
// Exporting the range to json file.
// Parameters:
//   range - Range 
//   options - JsonSaveOptions 
// Returns:
//   string  
func JsonUtility_ExportRangeToJson(range_ *Range, options *JsonSaveOptions)  (string,  error)  {
	
	CGoReturnPtr := C.JsonUtility_ExportRangeToJson(range_.ptr, options.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteJsonUtility(jsonutility *JsonUtility){
	runtime.SetFinalizer(jsonutility, nil)
	C.Delete_JsonUtility(jsonutility.ptr)
	jsonutility.ptr = nil
}
