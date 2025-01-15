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
	"errors"
	"runtime"
	"unsafe" 
)

/**************Enum LoadNumbersTableType *****************/

// Indicates type of loading tables when some tables are in a sheet.
type LoadNumbersTableType int32

const(
LoadNumbersTableType_OneTablePerSheet LoadNumbersTableType = 1 


LoadNumbersTableType_OverrideOtherTables LoadNumbersTableType = 2 


LoadNumbersTableType_TileTables LoadNumbersTableType = 4 
)

func Int32ToLoadNumbersTableType(value int32)(LoadNumbersTableType ,error){
	switch value {
		case 1:  return LoadNumbersTableType_OneTablePerSheet, nil  
		case 2:  return LoadNumbersTableType_OverrideOtherTables, nil  
		case 4:  return LoadNumbersTableType_TileTables, nil  
		default:
			return 0 ,fmt.Errorf("invalid LoadNumbersTableType value: %d", value)
	}
}
// Class NumbersLoadOptions 

// Represents the options of loading Apple Numbers files.
type NumbersLoadOptions struct {
	ptr unsafe.Pointer
}

// Constructor.
func NewNumbersLoadOptions() ( *NumbersLoadOptions, error) {
	numbersloadoptions := &NumbersLoadOptions{}
	CGoReturnPtr := C.New_NumbersLoadOptions()
	if CGoReturnPtr.error_no == 0 {
		numbersloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(numbersloadoptions, DeleteNumbersLoadOptions)
		return numbersloadoptions, nil
	} else {
		numbersloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return numbersloadoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - LoadOptions 
func NewNumbersLoadOptions_LoadOptions(src *LoadOptions) ( *NumbersLoadOptions, error) {
	numbersloadoptions := &NumbersLoadOptions{}
	CGoReturnPtr := C.New_NumbersLoadOptions_LoadOptions(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		numbersloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(numbersloadoptions, DeleteNumbersLoadOptions)
		return numbersloadoptions, nil
	} else {
		numbersloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return numbersloadoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *NumbersLoadOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the type of loading multiple tables in one worksheet.
// Returns:
//   int32  
func (instance *NumbersLoadOptions) GetLoadTableType()  (LoadNumbersTableType,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetLoadTableType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLoadNumbersTableType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of loading multiple tables in one worksheet.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetLoadTableType(value LoadNumbersTableType)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetLoadTableType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the load format.
// Returns:
//   int32  
func (instance *NumbersLoadOptions) GetLoadFormat()  (LoadFormat,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetLoadFormat( instance.ptr)
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
func (instance *NumbersLoadOptions) GetPassword()  (string,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetPassword( instance.ptr)
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
func (instance *NumbersLoadOptions) SetPassword(value string)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetPassword( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether parsing the formula when reading the file.
// Returns:
//   bool  
func (instance *NumbersLoadOptions) GetParsingFormulaOnOpen()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetParsingFormulaOnOpen( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether parsing the formula when reading the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetParsingFormulaOnOpen(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetParsingFormulaOnOpen( instance.ptr, C.bool(value))
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
func (instance *NumbersLoadOptions) GetParsingPivotCachedRecords()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetParsingPivotCachedRecords( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether parsing pivot cached records when loading the file.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetParsingPivotCachedRecords(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetParsingPivotCachedRecords( instance.ptr, C.bool(value))
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
func (instance *NumbersLoadOptions) SetPaperSize(type_ PaperSizeType)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetPaperSize( instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
// Returns:
//   int32  
func (instance *NumbersLoadOptions) GetLanguageCode()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetLanguageCode( instance.ptr)
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
func (instance *NumbersLoadOptions) SetLanguageCode(value CountryCode)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetLanguageCode( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the system regional settings based on CountryCode at the time the file was loaded.
// Returns:
//   int32  
func (instance *NumbersLoadOptions) GetRegion()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetRegion( instance.ptr)
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
// Gets or sets the system regional settings based on CountryCode at the time the file was loaded.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetRegion(value CountryCode)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetRegion( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the default style settings for initializing styles of the workbook
// Returns:
//   DefaultStyleSettings  
func (instance *NumbersLoadOptions) GetDefaultStyleSettings()  (*DefaultStyleSettings,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetDefaultStyleSettings( instance.ptr)
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
func (instance *NumbersLoadOptions) GetInterruptMonitor()  (*AbstractInterruptMonitor,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetInterruptMonitor( instance.ptr)
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
func (instance *NumbersLoadOptions) SetInterruptMonitor(value *AbstractInterruptMonitor)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetInterruptMonitor( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Ignore the data which are not printed if directly printing the file
// Returns:
//   bool  
func (instance *NumbersLoadOptions) GetIgnoreNotPrinted()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetIgnoreNotPrinted( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Ignore the data which are not printed if directly printing the file
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetIgnoreNotPrinted(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetIgnoreNotPrinted( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Check whether data is valid in the template file.
// Returns:
//   bool  
func (instance *NumbersLoadOptions) GetCheckDataValid()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetCheckDataValid( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Check whether data is valid in the template file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetCheckDataValid(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetCheckDataValid( instance.ptr, C.bool(value))
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
func (instance *NumbersLoadOptions) GetCheckExcelRestriction()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetCheckExcelRestriction( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

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
func (instance *NumbersLoadOptions) SetCheckExcelRestriction(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetCheckExcelRestriction( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether keep the unparsed data in memory for the Workbook when it is loaded from template file. Default is true.
// Returns:
//   bool  
func (instance *NumbersLoadOptions) GetKeepUnparsedData()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetKeepUnparsedData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Whether keep the unparsed data in memory for the Workbook when it is loaded from template file. Default is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetKeepUnparsedData(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetKeepUnparsedData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The filter to denote how to load data.
// Returns:
//   LoadFilter  
func (instance *NumbersLoadOptions) GetLoadFilter()  (*LoadFilter,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetLoadFilter( instance.ptr)
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
func (instance *NumbersLoadOptions) SetLoadFilter(value *LoadFilter)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetLoadFilter( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the memory usage options.
// Returns:
//   int32  
func (instance *NumbersLoadOptions) GetMemorySetting()  (MemorySetting,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetMemorySetting( instance.ptr)
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
// Gets or sets the memory usage options.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetMemorySetting(value MemorySetting)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetMemorySetting( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the auto fitter options
// Returns:
//   AutoFitterOptions  
func (instance *NumbersLoadOptions) GetAutoFitterOptions()  (*AutoFitterOptions,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetAutoFitterOptions( instance.ptr)
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
func (instance *NumbersLoadOptions) SetAutoFitterOptions(value *AutoFitterOptions)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetAutoFitterOptions( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether auto filtering the data when loading the files.
// Returns:
//   bool  
func (instance *NumbersLoadOptions) GetAutoFilter()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetAutoFilter( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether auto filtering the data when loading the files.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetAutoFilter(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetAutoFilter( instance.ptr, C.bool(value))
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
func (instance *NumbersLoadOptions) GetFontConfigs()  (*IndividualFontConfigs,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetFontConfigs( instance.ptr)
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
func (instance *NumbersLoadOptions) SetFontConfigs(value *IndividualFontConfigs)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetFontConfigs( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether ignoring useless shapes.
// Returns:
//   bool  
func (instance *NumbersLoadOptions) GetIgnoreUselessShapes()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetIgnoreUselessShapes( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether ignoring useless shapes.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetIgnoreUselessShapes(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetIgnoreUselessShapes( instance.ptr, C.bool(value))
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
func (instance *NumbersLoadOptions) GetPreservePaddingSpacesInFormula()  (bool,  error)  {
	
	CGoReturnPtr := C.NumbersLoadOptions_GetPreservePaddingSpacesInFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether preserve those spaces and line breaks that are padded between formula tokens
// while getting and setting formulas.
// Default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NumbersLoadOptions) SetPreservePaddingSpacesInFormula(value bool)  error {
	
	CGoReturnPtr := C.NumbersLoadOptions_SetPreservePaddingSpacesInFormula( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteNumbersLoadOptions(numbersloadoptions *NumbersLoadOptions){
	runtime.SetFinalizer(numbersloadoptions, nil)
	C.Delete_NumbersLoadOptions(numbersloadoptions.ptr)
	numbersloadoptions.ptr = nil
}
