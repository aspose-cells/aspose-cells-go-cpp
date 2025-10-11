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
 
 	
	"errors"	
	"runtime"
	"unsafe" 
)

// Class DbfLoadOptions 

// Represents the options of loading .dbf file.
type DbfLoadOptions struct {
	ptr unsafe.Pointer
}

// The options.
func NewDbfLoadOptions() ( *DbfLoadOptions, error) {
	dbfloadoptions := &DbfLoadOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_DbfLoadOptions"),)
	if CGoReturnPtr.error_no == 0 {
		dbfloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(dbfloadoptions, DeleteDbfLoadOptions)
		return dbfloadoptions, nil
	} else {
		dbfloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return dbfloadoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - LoadOptions 
func NewDbfLoadOptions_LoadOptions(src *LoadOptions) ( *DbfLoadOptions, error) {
	dbfloadoptions := &DbfLoadOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_DbfLoadOptions_LoadOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		dbfloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(dbfloadoptions, DeleteDbfLoadOptions)
		return dbfloadoptions, nil
	} else {
		dbfloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return dbfloadoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DbfLoadOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the load format.
// Returns:
//   int32  
func (instance *DbfLoadOptions) GetLoadFormat()  (LoadFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DbfLoadOptions_GetLoadFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLoadFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the password of the workbook.
// Returns:
//   string  
func (instance *DbfLoadOptions) GetPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DbfLoadOptions_GetPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the password of the workbook.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetPassword(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DbfLoadOptions_SetPassword"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether parsing the formula when reading the file.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetParsingFormulaOnOpen()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetParsingFormulaOnOpen"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether parsing the formula when reading the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetParsingFormulaOnOpen(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetParsingFormulaOnOpen"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether parsing pivot cached records when loading the file.
// The default value is false.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetParsingPivotCachedRecords()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetParsingPivotCachedRecords"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether parsing pivot cached records when loading the file.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetParsingPivotCachedRecords(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetParsingPivotCachedRecords"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets the default print paper size from default printer's setting.
// Parameters:
//   type - int32 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetPaperSize(type_ PaperSizeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DbfLoadOptions_SetPaperSize"), instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
// Returns:
//   int32  
func (instance *DbfLoadOptions) GetLanguageCode()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DbfLoadOptions_GetLanguageCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCountryCode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetLanguageCode(value CountryCode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DbfLoadOptions_SetLanguageCode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the regional settings used for the Workbook that will be loaded.
// Returns:
//   int32  
func (instance *DbfLoadOptions) GetRegion()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DbfLoadOptions_GetRegion"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCountryCode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the regional settings used for the Workbook that will be loaded.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetRegion(value CountryCode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DbfLoadOptions_SetRegion"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the default style settings for initializing styles of the workbook
// Returns:
//   DefaultStyleSettings  
func (instance *DbfLoadOptions) GetDefaultStyleSettings()  (*DefaultStyleSettings,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DbfLoadOptions_GetDefaultStyleSettings"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DefaultStyleSettings{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDefaultStyleSettings) 

	return result, nil 
}
// Gets and sets the interrupt monitor.
// Returns:
//   AbstractInterruptMonitor  
func (instance *DbfLoadOptions) GetInterruptMonitor()  (*AbstractInterruptMonitor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DbfLoadOptions_GetInterruptMonitor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AbstractInterruptMonitor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAbstractInterruptMonitor) 

	return result, nil 
}
// Gets and sets the interrupt monitor.
// Parameters:
//   value - AbstractInterruptMonitor 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetInterruptMonitor(value *AbstractInterruptMonitor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DbfLoadOptions_SetInterruptMonitor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Ignore the data which are not printed if directly printing the file
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetIgnoreNotPrinted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetIgnoreNotPrinted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Ignore the data which are not printed if directly printing the file
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetIgnoreNotPrinted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetIgnoreNotPrinted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Check whether data is valid in the template file.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetCheckDataValid()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetCheckDataValid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Check whether data is valid in the template file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetCheckDataValid(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetCheckDataValid"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K such as by Cell.PutValue(string), if this property is true, you will get an Exception.
// If this property is false, we will accept your input string value as the cell's value so that later
// you can output the complete string value for other file formats such as CSV.
// However, if you have set such kind of value that is invalid for excel file format,
// you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetCheckExcelRestriction()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetCheckExcelRestriction"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K such as by Cell.PutValue(string), if this property is true, you will get an Exception.
// If this property is false, we will accept your input string value as the cell's value so that later
// you can output the complete string value for other file formats such as CSV.
// However, if you have set such kind of value that is invalid for excel file format,
// you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetCheckExcelRestriction(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetCheckExcelRestriction"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether keep the unparsed data in memory for the Workbook when it is loaded from template file. Default is true.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetKeepUnparsedData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetKeepUnparsedData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether keep the unparsed data in memory for the Workbook when it is loaded from template file. Default is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetKeepUnparsedData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetKeepUnparsedData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The filter to denote how to load data.
// Returns:
//   LoadFilter  
func (instance *DbfLoadOptions) GetLoadFilter()  (*LoadFilter,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DbfLoadOptions_GetLoadFilter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LoadFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLoadFilter) 

	return result, nil 
}
// The filter to denote how to load data.
// Parameters:
//   value - LoadFilter 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetLoadFilter(value *LoadFilter)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DbfLoadOptions_SetLoadFilter"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the memory mode for loaded workbook.
// Returns:
//   int32  
func (instance *DbfLoadOptions) GetMemorySetting()  (MemorySetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DbfLoadOptions_GetMemorySetting"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMemorySetting(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the memory mode for loaded workbook.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetMemorySetting(value MemorySetting)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DbfLoadOptions_SetMemorySetting"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the auto fitter options
// Returns:
//   AutoFitterOptions  
func (instance *DbfLoadOptions) GetAutoFitterOptions()  (*AutoFitterOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DbfLoadOptions_GetAutoFitterOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AutoFitterOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAutoFitterOptions) 

	return result, nil 
}
// Gets and sets the auto fitter options
// Parameters:
//   value - AutoFitterOptions 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetAutoFitterOptions(value *AutoFitterOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DbfLoadOptions_SetAutoFitterOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether auto filtering the data when loading the files.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetAutoFilter()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetAutoFilter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether auto filtering the data when loading the files.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetAutoFilter(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetAutoFilter"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets individual font configs.
// Only works for the <see cref="Workbook"/> which uses this <see cref="LoadOptions"/> to load.
// Returns:
//   IndividualFontConfigs  
func (instance *DbfLoadOptions) GetFontConfigs()  (*IndividualFontConfigs,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DbfLoadOptions_GetFontConfigs"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &IndividualFontConfigs{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteIndividualFontConfigs) 

	return result, nil 
}
// Gets and sets individual font configs.
// Only works for the <see cref="Workbook"/> which uses this <see cref="LoadOptions"/> to load.
// Parameters:
//   value - IndividualFontConfigs 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetFontConfigs(value *IndividualFontConfigs)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DbfLoadOptions_SetFontConfigs"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether ignoring useless shapes.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetIgnoreUselessShapes()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetIgnoreUselessShapes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether ignoring useless shapes.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetIgnoreUselessShapes(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetIgnoreUselessShapes"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether preserve those spaces and line breaks that are padded between formula tokens
// while getting and setting formulas.
// Default value is false.
// Returns:
//   bool  
func (instance *DbfLoadOptions) GetPreservePaddingSpacesInFormula()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfLoadOptions_GetPreservePaddingSpacesInFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether preserve those spaces and line breaks that are padded between formula tokens
// while getting and setting formulas.
// Default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfLoadOptions) SetPreservePaddingSpacesInFormula(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfLoadOptions_SetPreservePaddingSpacesInFormula"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *DbfLoadOptions) ToLoadOptions() *LoadOptions {
	parentClass := &LoadOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDbfLoadOptions(dbfloadoptions *DbfLoadOptions){
	runtime.SetFinalizer(dbfloadoptions, nil)
	C.Delete_CObject(C.CString("Delete_DbfLoadOptions"),dbfloadoptions.ptr)
	dbfloadoptions.ptr = nil
}

// Class DifLoadOptions 

// Represents the options of loading .dif file.
type DifLoadOptions struct {
	ptr unsafe.Pointer
}

// The options.
func NewDifLoadOptions() ( *DifLoadOptions, error) {
	difloadoptions := &DifLoadOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_DifLoadOptions"),)
	if CGoReturnPtr.error_no == 0 {
		difloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(difloadoptions, DeleteDifLoadOptions)
		return difloadoptions, nil
	} else {
		difloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return difloadoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - LoadOptions 
func NewDifLoadOptions_LoadOptions(src *LoadOptions) ( *DifLoadOptions, error) {
	difloadoptions := &DifLoadOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_DifLoadOptions_LoadOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		difloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(difloadoptions, DeleteDifLoadOptions)
		return difloadoptions, nil
	} else {
		difloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return difloadoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DifLoadOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the load format.
// Returns:
//   int32  
func (instance *DifLoadOptions) GetLoadFormat()  (LoadFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DifLoadOptions_GetLoadFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLoadFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the password of the workbook.
// Returns:
//   string  
func (instance *DifLoadOptions) GetPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DifLoadOptions_GetPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the password of the workbook.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DifLoadOptions) SetPassword(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DifLoadOptions_SetPassword"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether parsing the formula when reading the file.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetParsingFormulaOnOpen()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetParsingFormulaOnOpen"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether parsing the formula when reading the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetParsingFormulaOnOpen(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetParsingFormulaOnOpen"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether parsing pivot cached records when loading the file.
// The default value is false.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetParsingPivotCachedRecords()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetParsingPivotCachedRecords"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether parsing pivot cached records when loading the file.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetParsingPivotCachedRecords(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetParsingPivotCachedRecords"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets the default print paper size from default printer's setting.
// Parameters:
//   type - int32 
// Returns:
//   void  
func (instance *DifLoadOptions) SetPaperSize(type_ PaperSizeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DifLoadOptions_SetPaperSize"), instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
// Returns:
//   int32  
func (instance *DifLoadOptions) GetLanguageCode()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DifLoadOptions_GetLanguageCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCountryCode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DifLoadOptions) SetLanguageCode(value CountryCode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DifLoadOptions_SetLanguageCode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the regional settings used for the Workbook that will be loaded.
// Returns:
//   int32  
func (instance *DifLoadOptions) GetRegion()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DifLoadOptions_GetRegion"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCountryCode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the regional settings used for the Workbook that will be loaded.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DifLoadOptions) SetRegion(value CountryCode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DifLoadOptions_SetRegion"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the default style settings for initializing styles of the workbook
// Returns:
//   DefaultStyleSettings  
func (instance *DifLoadOptions) GetDefaultStyleSettings()  (*DefaultStyleSettings,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DifLoadOptions_GetDefaultStyleSettings"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DefaultStyleSettings{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDefaultStyleSettings) 

	return result, nil 
}
// Gets and sets the interrupt monitor.
// Returns:
//   AbstractInterruptMonitor  
func (instance *DifLoadOptions) GetInterruptMonitor()  (*AbstractInterruptMonitor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DifLoadOptions_GetInterruptMonitor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AbstractInterruptMonitor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAbstractInterruptMonitor) 

	return result, nil 
}
// Gets and sets the interrupt monitor.
// Parameters:
//   value - AbstractInterruptMonitor 
// Returns:
//   void  
func (instance *DifLoadOptions) SetInterruptMonitor(value *AbstractInterruptMonitor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DifLoadOptions_SetInterruptMonitor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Ignore the data which are not printed if directly printing the file
// Returns:
//   bool  
func (instance *DifLoadOptions) GetIgnoreNotPrinted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetIgnoreNotPrinted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Ignore the data which are not printed if directly printing the file
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetIgnoreNotPrinted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetIgnoreNotPrinted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Check whether data is valid in the template file.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetCheckDataValid()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetCheckDataValid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Check whether data is valid in the template file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetCheckDataValid(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetCheckDataValid"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K such as by Cell.PutValue(string), if this property is true, you will get an Exception.
// If this property is false, we will accept your input string value as the cell's value so that later
// you can output the complete string value for other file formats such as CSV.
// However, if you have set such kind of value that is invalid for excel file format,
// you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetCheckExcelRestriction()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetCheckExcelRestriction"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K such as by Cell.PutValue(string), if this property is true, you will get an Exception.
// If this property is false, we will accept your input string value as the cell's value so that later
// you can output the complete string value for other file formats such as CSV.
// However, if you have set such kind of value that is invalid for excel file format,
// you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetCheckExcelRestriction(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetCheckExcelRestriction"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether keep the unparsed data in memory for the Workbook when it is loaded from template file. Default is true.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetKeepUnparsedData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetKeepUnparsedData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether keep the unparsed data in memory for the Workbook when it is loaded from template file. Default is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetKeepUnparsedData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetKeepUnparsedData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The filter to denote how to load data.
// Returns:
//   LoadFilter  
func (instance *DifLoadOptions) GetLoadFilter()  (*LoadFilter,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DifLoadOptions_GetLoadFilter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LoadFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLoadFilter) 

	return result, nil 
}
// The filter to denote how to load data.
// Parameters:
//   value - LoadFilter 
// Returns:
//   void  
func (instance *DifLoadOptions) SetLoadFilter(value *LoadFilter)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DifLoadOptions_SetLoadFilter"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the memory mode for loaded workbook.
// Returns:
//   int32  
func (instance *DifLoadOptions) GetMemorySetting()  (MemorySetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DifLoadOptions_GetMemorySetting"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMemorySetting(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the memory mode for loaded workbook.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DifLoadOptions) SetMemorySetting(value MemorySetting)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DifLoadOptions_SetMemorySetting"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the auto fitter options
// Returns:
//   AutoFitterOptions  
func (instance *DifLoadOptions) GetAutoFitterOptions()  (*AutoFitterOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DifLoadOptions_GetAutoFitterOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AutoFitterOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAutoFitterOptions) 

	return result, nil 
}
// Gets and sets the auto fitter options
// Parameters:
//   value - AutoFitterOptions 
// Returns:
//   void  
func (instance *DifLoadOptions) SetAutoFitterOptions(value *AutoFitterOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DifLoadOptions_SetAutoFitterOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether auto filtering the data when loading the files.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetAutoFilter()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetAutoFilter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether auto filtering the data when loading the files.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetAutoFilter(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetAutoFilter"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets individual font configs.
// Only works for the <see cref="Workbook"/> which uses this <see cref="LoadOptions"/> to load.
// Returns:
//   IndividualFontConfigs  
func (instance *DifLoadOptions) GetFontConfigs()  (*IndividualFontConfigs,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("DifLoadOptions_GetFontConfigs"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &IndividualFontConfigs{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteIndividualFontConfigs) 

	return result, nil 
}
// Gets and sets individual font configs.
// Only works for the <see cref="Workbook"/> which uses this <see cref="LoadOptions"/> to load.
// Parameters:
//   value - IndividualFontConfigs 
// Returns:
//   void  
func (instance *DifLoadOptions) SetFontConfigs(value *IndividualFontConfigs)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DifLoadOptions_SetFontConfigs"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether ignoring useless shapes.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetIgnoreUselessShapes()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetIgnoreUselessShapes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether ignoring useless shapes.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetIgnoreUselessShapes(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetIgnoreUselessShapes"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether preserve those spaces and line breaks that are padded between formula tokens
// while getting and setting formulas.
// Default value is false.
// Returns:
//   bool  
func (instance *DifLoadOptions) GetPreservePaddingSpacesInFormula()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DifLoadOptions_GetPreservePaddingSpacesInFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether preserve those spaces and line breaks that are padded between formula tokens
// while getting and setting formulas.
// Default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DifLoadOptions) SetPreservePaddingSpacesInFormula(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DifLoadOptions_SetPreservePaddingSpacesInFormula"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *DifLoadOptions) ToLoadOptions() *LoadOptions {
	parentClass := &LoadOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDifLoadOptions(difloadoptions *DifLoadOptions){
	runtime.SetFinalizer(difloadoptions, nil)
	C.Delete_CObject(C.CString("Delete_DifLoadOptions"),difloadoptions.ptr)
	difloadoptions.ptr = nil
}
