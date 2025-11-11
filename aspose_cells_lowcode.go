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

// Class AbstractLowCodeLoadOptionsProvider 

// Implementation to provide multiple load options for processes
// that use multiple inputs(such as template files).
type AbstractLowCodeLoadOptionsProvider struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *AbstractLowCodeLoadOptionsProvider) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AbstractLowCodeLoadOptionsProvider_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Checks whether there is more input.
// Returns:
//   bool  
func (instance *AbstractLowCodeLoadOptionsProvider) MoveNext()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AbstractLowCodeLoadOptionsProvider_MoveNext"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the load options from which to load data of currently processed part.
// Returns:
//   LowCodeLoadOptions  
func (instance *AbstractLowCodeLoadOptionsProvider) GetCurrent()  (*LowCodeLoadOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("AbstractLowCodeLoadOptionsProvider_GetCurrent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeLoadOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeLoadOptions) 

	return result, nil 
}
// Releases resources after processing currently part of input.
// Parameters:
//   part - LowCodeLoadOptions 
// Returns:
//   void  
func (instance *AbstractLowCodeLoadOptionsProvider) Finish(part *LowCodeLoadOptions)  error {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("AbstractLowCodeLoadOptionsProvider_Finish"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteAbstractLowCodeLoadOptionsProvider(abstractlowcodeloadoptionsprovider *AbstractLowCodeLoadOptionsProvider){
	runtime.SetFinalizer(abstractlowcodeloadoptionsprovider, nil)
	C.Delete_CObject(C.CString("Delete_AbstractLowCodeLoadOptionsProvider"),abstractlowcodeloadoptionsprovider.ptr)
	abstractlowcodeloadoptionsprovider.ptr = nil
}

// Class AbstractLowCodeProtectionProvider 

// Implementation to provide protection settings
type AbstractLowCodeProtectionProvider struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewAbstractLowCodeProtectionProvider() ( *AbstractLowCodeProtectionProvider, error) {
	abstractlowcodeprotectionprovider := &AbstractLowCodeProtectionProvider{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_AbstractLowCodeProtectionProvider"),)
	if CGoReturnPtr.error_no == 0 {
		abstractlowcodeprotectionprovider.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(abstractlowcodeprotectionprovider, DeleteAbstractLowCodeProtectionProvider)
		return abstractlowcodeprotectionprovider, nil
	} else {
		abstractlowcodeprotectionprovider.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return abstractlowcodeprotectionprovider, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *AbstractLowCodeProtectionProvider) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AbstractLowCodeProtectionProvider_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the password to open spread sheet file.
// Returns:
//   string  
func (instance *AbstractLowCodeProtectionProvider) GetOpenPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("AbstractLowCodeProtectionProvider_GetOpenPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the password to modify spread sheet file.
// Returns:
//   string  
func (instance *AbstractLowCodeProtectionProvider) GetWritePassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("AbstractLowCodeProtectionProvider_GetWritePassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the password to protect the workbook with specified protection type.
// Returns:
//   string  
func (instance *AbstractLowCodeProtectionProvider) GetWorkbookPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("AbstractLowCodeProtectionProvider_GetWorkbookPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the protection type to protect the workbook.
// Returns:
//   int32  
func (instance *AbstractLowCodeProtectionProvider) GetWorkbookProtectionType()  (ProtectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("AbstractLowCodeProtectionProvider_GetWorkbookProtectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToProtectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the password to protect the specified worksheet.
// Parameters:
//   sheetName - string 
// Returns:
//   string  
func (instance *AbstractLowCodeProtectionProvider) GetWorksheetPassword(sheetname string)  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGN(C.CString("AbstractLowCodeProtectionProvider_GetWorksheetPassword"), instance.ptr, C.CString(sheetname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the protection type to protect the specified worksheet.
// Parameters:
//   sheetName - string 
// Returns:
//   int32  
func (instance *AbstractLowCodeProtectionProvider) GetWorksheetProtectionType(sheetname string)  (ProtectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKM(C.CString("AbstractLowCodeProtectionProvider_GetWorksheetProtectionType"), instance.ptr, C.CString(sheetname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToProtectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}



func DeleteAbstractLowCodeProtectionProvider(abstractlowcodeprotectionprovider *AbstractLowCodeProtectionProvider){
	runtime.SetFinalizer(abstractlowcodeprotectionprovider, nil)
	C.Delete_CObject(C.CString("Delete_AbstractLowCodeProtectionProvider"),abstractlowcodeprotectionprovider.ptr)
	abstractlowcodeprotectionprovider.ptr = nil
}

// Class AbstractLowCodeSaveOptionsProvider 

// Implementation to provide multiple save options for processes
// that require multiple outputs. For example,
// <see cref="SpreadsheetSplitter"/> feature requires multiple destinations
// to save the split files.
type AbstractLowCodeSaveOptionsProvider struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *AbstractLowCodeSaveOptionsProvider) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AbstractLowCodeSaveOptionsProvider_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the save options from which to get the output settings for currently split part.
// Returning null denotes to skip given part.
// Parameters:
//   part - SplitPartInfo 
// Returns:
//   LowCodeSaveOptions  
func (instance *AbstractLowCodeSaveOptionsProvider) GetSaveOptions(part *SplitPartInfo)  (*LowCodeSaveOptions,  error)  {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZDK(C.CString("AbstractLowCodeSaveOptionsProvider_GetSaveOptions"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// Releases resources after processing currently split part.
// Parameters:
//   part - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *AbstractLowCodeSaveOptionsProvider) Finish(part *LowCodeSaveOptions)  error {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("AbstractLowCodeSaveOptionsProvider_Finish"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteAbstractLowCodeSaveOptionsProvider(abstractlowcodesaveoptionsprovider *AbstractLowCodeSaveOptionsProvider){
	runtime.SetFinalizer(abstractlowcodesaveoptionsprovider, nil)
	C.Delete_CObject(C.CString("Delete_AbstractLowCodeSaveOptionsProvider"),abstractlowcodesaveoptionsprovider.ptr)
	abstractlowcodesaveoptionsprovider.ptr = nil
}

// Class HtmlConverter 

// Converter for conversion between html files(html or mht) and other spreadsheet file formats.
type HtmlConverter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *HtmlConverter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("HtmlConverter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts given template file between html and other formats.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func HtmlConverter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("HtmlConverter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts file between html and other spreadsheet file formats.
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
// Returns:
//   void  
func HtmlConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("HtmlConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions"),loadoptions_ptr, saveoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteHtmlConverter(htmlconverter *HtmlConverter){
	runtime.SetFinalizer(htmlconverter, nil)
	C.Delete_CObject(C.CString("Delete_HtmlConverter"),htmlconverter.ptr)
	htmlconverter.ptr = nil
}

// Class ImageConverter 

// Converter for converting template file to images.
type ImageConverter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ImageConverter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageConverter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts template file to images.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func ImageConverter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("ImageConverter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts template file to images
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
// Returns:
//   void  
func ImageConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ImageConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions"),loadoptions_ptr, saveoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts template file to images
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
//   provider - AbstractLowCodeSaveOptionsProvider 
// Returns:
//   void  
func ImageConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions_AbstractLowCodeSaveOptionsProvider(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions, provider *AbstractLowCodeSaveOptionsProvider)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}
	var provider_ptr unsafe.Pointer = nil
	if provider != nil {
	  provider_ptr =provider.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZAP(C.CString("ImageConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions_AbstractLowCodeSaveOptionsProvider"),loadoptions_ptr, saveoptions_ptr, provider_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteImageConverter(imageconverter *ImageConverter){
	runtime.SetFinalizer(imageconverter, nil)
	C.Delete_CObject(C.CString("Delete_ImageConverter"),imageconverter.ptr)
	imageconverter.ptr = nil
}

// Class JsonConverter 

// Converter for conversion between json data structure and other spreadsheet file formats.
type JsonConverter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *JsonConverter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("JsonConverter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts given template file between json and other formats.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func JsonConverter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("JsonConverter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts between json data and other spreadsheet file formats.
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
// Returns:
//   void  
func JsonConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("JsonConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions"),loadoptions_ptr, saveoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteJsonConverter(jsonconverter *JsonConverter){
	runtime.SetFinalizer(jsonconverter, nil)
	C.Delete_CObject(C.CString("Delete_JsonConverter"),jsonconverter.ptr)
	jsonconverter.ptr = nil
}

// Class LowCodeHtmlSaveOptions 

// Options for saving html in low code way.
type LowCodeHtmlSaveOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeHtmlSaveOptions() ( *LowCodeHtmlSaveOptions, error) {
	lowcodehtmlsaveoptions := &LowCodeHtmlSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeHtmlSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodehtmlsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodehtmlsaveoptions, DeleteLowCodeHtmlSaveOptions)
		return lowcodehtmlsaveoptions, nil
	} else {
		lowcodehtmlsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodehtmlsaveoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - LowCodeSaveOptions 
func NewLowCodeHtmlSaveOptions_LowCodeSaveOptions(src *LowCodeSaveOptions) ( *LowCodeHtmlSaveOptions, error) {
	lowcodehtmlsaveoptions := &LowCodeHtmlSaveOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_LowCodeHtmlSaveOptions_LowCodeSaveOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		lowcodehtmlsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodehtmlsaveoptions, DeleteLowCodeHtmlSaveOptions)
		return lowcodehtmlsaveoptions, nil
	} else {
		lowcodehtmlsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodehtmlsaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeHtmlSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeHtmlSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the format of spreadsheet.
// Returns:
//   int32  
func (instance *LowCodeHtmlSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LowCodeHtmlSaveOptions_GetSaveFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSaveFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the format of spreadsheet.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeHtmlSaveOptions) SetSaveFormat(value SaveFormat)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LowCodeHtmlSaveOptions_SetSaveFormat"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The general options for saving html.
// Returns:
//   HtmlSaveOptions  
func (instance *LowCodeHtmlSaveOptions) GetHtmlOptions()  (*HtmlSaveOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeHtmlSaveOptions_GetHtmlOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &HtmlSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteHtmlSaveOptions) 

	return result, nil 
}
// The general options for saving html.
// Parameters:
//   value - HtmlSaveOptions 
// Returns:
//   void  
func (instance *LowCodeHtmlSaveOptions) SetHtmlOptions(value *HtmlSaveOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeHtmlSaveOptions_SetHtmlOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Returns:
//   string  
func (instance *LowCodeHtmlSaveOptions) GetOutputFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeHtmlSaveOptions_GetOutputFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeHtmlSaveOptions) SetOutputFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeHtmlSaveOptions_SetOutputFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Returns:
//   []byte  
func (instance *LowCodeHtmlSaveOptions) GetOutputStream()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LowCodeHtmlSaveOptions_GetOutputStream"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LowCodeHtmlSaveOptions) SetOutputStream(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LowCodeHtmlSaveOptions_SetOutputStream"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *LowCodeHtmlSaveOptions) ToLowCodeSaveOptions() *LowCodeSaveOptions {
	parentClass := &LowCodeSaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLowCodeHtmlSaveOptions(lowcodehtmlsaveoptions *LowCodeHtmlSaveOptions){
	runtime.SetFinalizer(lowcodehtmlsaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeHtmlSaveOptions"),lowcodehtmlsaveoptions.ptr)
	lowcodehtmlsaveoptions.ptr = nil
}

// Class LowCodeImageSaveOptions 

// Options for saving image in low code way.
type LowCodeImageSaveOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeImageSaveOptions() ( *LowCodeImageSaveOptions, error) {
	lowcodeimagesaveoptions := &LowCodeImageSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeImageSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodeimagesaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodeimagesaveoptions, DeleteLowCodeImageSaveOptions)
		return lowcodeimagesaveoptions, nil
	} else {
		lowcodeimagesaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodeimagesaveoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - LowCodeSaveOptions 
func NewLowCodeImageSaveOptions_LowCodeSaveOptions(src *LowCodeSaveOptions) ( *LowCodeImageSaveOptions, error) {
	lowcodeimagesaveoptions := &LowCodeImageSaveOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_LowCodeImageSaveOptions_LowCodeSaveOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		lowcodeimagesaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodeimagesaveoptions, DeleteLowCodeImageSaveOptions)
		return lowcodeimagesaveoptions, nil
	} else {
		lowcodeimagesaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodeimagesaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeImageSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeImageSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the save format.
// Returns:
//   int32  
func (instance *LowCodeImageSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LowCodeImageSaveOptions_GetSaveFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSaveFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the save format.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeImageSaveOptions) SetSaveFormat(value SaveFormat)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LowCodeImageSaveOptions_SetSaveFormat"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The options for rendering images.
// Returns:
//   ImageOrPrintOptions  
func (instance *LowCodeImageSaveOptions) GetImageOptions()  (*ImageOrPrintOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeImageSaveOptions_GetImageOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ImageOrPrintOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteImageOrPrintOptions) 

	return result, nil 
}
// The options for rendering images.
// Parameters:
//   value - ImageOrPrintOptions 
// Returns:
//   void  
func (instance *LowCodeImageSaveOptions) SetImageOptions(value *ImageOrPrintOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeImageSaveOptions_SetImageOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Provider of save options for saving generated images.
// Returns:
//   AbstractLowCodeSaveOptionsProvider  
func (instance *LowCodeImageSaveOptions) GetSaveOptionsProvider()  (*AbstractLowCodeSaveOptionsProvider,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeImageSaveOptions_GetSaveOptionsProvider"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AbstractLowCodeSaveOptionsProvider{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAbstractLowCodeSaveOptionsProvider) 

	return result, nil 
}
// Provider of save options for saving generated images.
// Parameters:
//   value - AbstractLowCodeSaveOptionsProvider 
// Returns:
//   void  
func (instance *LowCodeImageSaveOptions) SetSaveOptionsProvider(value *AbstractLowCodeSaveOptionsProvider)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeImageSaveOptions_SetSaveOptionsProvider"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Returns:
//   string  
func (instance *LowCodeImageSaveOptions) GetOutputFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeImageSaveOptions_GetOutputFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeImageSaveOptions) SetOutputFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeImageSaveOptions_SetOutputFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Returns:
//   []byte  
func (instance *LowCodeImageSaveOptions) GetOutputStream()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LowCodeImageSaveOptions_GetOutputStream"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LowCodeImageSaveOptions) SetOutputStream(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LowCodeImageSaveOptions_SetOutputStream"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *LowCodeImageSaveOptions) ToLowCodeSaveOptions() *LowCodeSaveOptions {
	parentClass := &LowCodeSaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLowCodeImageSaveOptions(lowcodeimagesaveoptions *LowCodeImageSaveOptions){
	runtime.SetFinalizer(lowcodeimagesaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeImageSaveOptions"),lowcodeimagesaveoptions.ptr)
	lowcodeimagesaveoptions.ptr = nil
}

// Class LowCodeLoadOptions 

// Options for loading template file.
type LowCodeLoadOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeLoadOptions() ( *LowCodeLoadOptions, error) {
	lowcodeloadoptions := &LowCodeLoadOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeLoadOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodeloadoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodeloadoptions, DeleteLowCodeLoadOptions)
		return lowcodeloadoptions, nil
	} else {
		lowcodeloadoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodeloadoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeLoadOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeLoadOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) of the template.
// Returns:
//   string  
func (instance *LowCodeLoadOptions) GetInputFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeLoadOptions_GetInputFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) of the template.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeLoadOptions) SetInputFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeLoadOptions_SetInputFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the Stream of the template.
// Returns:
//   []byte  
func (instance *LowCodeLoadOptions) GetInputStream()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LowCodeLoadOptions_GetInputStream"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the Stream of the template.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LowCodeLoadOptions) SetInputStream(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LowCodeLoadOptions_SetInputStream"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteLowCodeLoadOptions(lowcodeloadoptions *LowCodeLoadOptions){
	runtime.SetFinalizer(lowcodeloadoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeLoadOptions"),lowcodeloadoptions.ptr)
	lowcodeloadoptions.ptr = nil
}

// Class LowCodeMergeOptions 

// Options for merging multiple template files into one.
type LowCodeMergeOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeMergeOptions() ( *LowCodeMergeOptions, error) {
	lowcodemergeoptions := &LowCodeMergeOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeMergeOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodemergeoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodemergeoptions, DeleteLowCodeMergeOptions)
		return lowcodemergeoptions, nil
	} else {
		lowcodemergeoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodemergeoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeMergeOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeMergeOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Save options for saving the split parts.
// Returns:
//   LowCodeSaveOptions  
func (instance *LowCodeMergeOptions) GetSaveOptions()  (*LowCodeSaveOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeMergeOptions_GetSaveOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// Save options for saving the split parts.
// Parameters:
//   value - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *LowCodeMergeOptions) SetSaveOptions(value *LowCodeSaveOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeMergeOptions_SetSaveOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Provider of save options for saving the split parts.
// Returns:
//   AbstractLowCodeLoadOptionsProvider  
func (instance *LowCodeMergeOptions) GetLoadOptionsProvider()  (*AbstractLowCodeLoadOptionsProvider,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeMergeOptions_GetLoadOptionsProvider"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AbstractLowCodeLoadOptionsProvider{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAbstractLowCodeLoadOptionsProvider) 

	return result, nil 
}
// Provider of save options for saving the split parts.
// Parameters:
//   value - AbstractLowCodeLoadOptionsProvider 
// Returns:
//   void  
func (instance *LowCodeMergeOptions) SetLoadOptionsProvider(value *AbstractLowCodeLoadOptionsProvider)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeMergeOptions_SetLoadOptionsProvider"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteLowCodeMergeOptions(lowcodemergeoptions *LowCodeMergeOptions){
	runtime.SetFinalizer(lowcodemergeoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeMergeOptions"),lowcodemergeoptions.ptr)
	lowcodemergeoptions.ptr = nil
}

// Class LowCodePdfSaveOptions 

// Options for saving pdf in low code way.
type LowCodePdfSaveOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodePdfSaveOptions() ( *LowCodePdfSaveOptions, error) {
	lowcodepdfsaveoptions := &LowCodePdfSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodePdfSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodepdfsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodepdfsaveoptions, DeleteLowCodePdfSaveOptions)
		return lowcodepdfsaveoptions, nil
	} else {
		lowcodepdfsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodepdfsaveoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - LowCodeSaveOptions 
func NewLowCodePdfSaveOptions_LowCodeSaveOptions(src *LowCodeSaveOptions) ( *LowCodePdfSaveOptions, error) {
	lowcodepdfsaveoptions := &LowCodePdfSaveOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_LowCodePdfSaveOptions_LowCodeSaveOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		lowcodepdfsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodepdfsaveoptions, DeleteLowCodePdfSaveOptions)
		return lowcodepdfsaveoptions, nil
	} else {
		lowcodepdfsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodepdfsaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodePdfSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodePdfSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// The save format for the output.
// For converting to pdf, it can only be <see cref="SaveFormat.Pdf"/>.
// Returns:
//   int32  
func (instance *LowCodePdfSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LowCodePdfSaveOptions_GetSaveFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSaveFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// The save format for the output.
// For converting to pdf, it can only be <see cref="SaveFormat.Pdf"/>.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodePdfSaveOptions) SetSaveFormat(value SaveFormat)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LowCodePdfSaveOptions_SetSaveFormat"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The options for saving Pdf file.
// Returns:
//   PdfSaveOptions  
func (instance *LowCodePdfSaveOptions) GetPdfOptions()  (*PdfSaveOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodePdfSaveOptions_GetPdfOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PdfSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePdfSaveOptions) 

	return result, nil 
}
// The options for saving Pdf file.
// Parameters:
//   value - PdfSaveOptions 
// Returns:
//   void  
func (instance *LowCodePdfSaveOptions) SetPdfOptions(value *PdfSaveOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodePdfSaveOptions_SetPdfOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Returns:
//   string  
func (instance *LowCodePdfSaveOptions) GetOutputFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodePdfSaveOptions_GetOutputFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodePdfSaveOptions) SetOutputFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodePdfSaveOptions_SetOutputFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Returns:
//   []byte  
func (instance *LowCodePdfSaveOptions) GetOutputStream()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LowCodePdfSaveOptions_GetOutputStream"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LowCodePdfSaveOptions) SetOutputStream(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LowCodePdfSaveOptions_SetOutputStream"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *LowCodePdfSaveOptions) ToLowCodeSaveOptions() *LowCodeSaveOptions {
	parentClass := &LowCodeSaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLowCodePdfSaveOptions(lowcodepdfsaveoptions *LowCodePdfSaveOptions){
	runtime.SetFinalizer(lowcodepdfsaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodePdfSaveOptions"),lowcodepdfsaveoptions.ptr)
	lowcodepdfsaveoptions.ptr = nil
}

// Class LowCodeSaveOptions 

// Options for saving generated results in low code way.
type LowCodeSaveOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeSaveOptions() ( *LowCodeSaveOptions, error) {
	lowcodesaveoptions := &LowCodeSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodesaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodesaveoptions, DeleteLowCodeSaveOptions)
		return lowcodesaveoptions, nil
	} else {
		lowcodesaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodesaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Returns:
//   string  
func (instance *LowCodeSaveOptions) GetOutputFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptions_GetOutputFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the file(with path if needed) for saving the generated data.
// When setting this property with value other than null or empty string, <see cref="OutputStream"/> will be ignored.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptions) SetOutputFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptions_SetOutputFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Returns:
//   []byte  
func (instance *LowCodeSaveOptions) GetOutputStream()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LowCodeSaveOptions_GetOutputStream"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the Stream for writing the generated data to.
// When setting this property with value other than null, <see cref="OutputFile"/> will be ignored.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LowCodeSaveOptions) SetOutputStream(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LowCodeSaveOptions_SetOutputStream"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the save format for the output.
// Generally, for specific process in low code way, only some specific formats are allowed.
// Please specify the correct format for corresponding process, otherwise unexpected result
// or even exception may be caused.
// Returns:
//   int32  
func (instance *LowCodeSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LowCodeSaveOptions_GetSaveFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSaveFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the save format for the output.
// Generally, for specific process in low code way, only some specific formats are allowed.
// Please specify the correct format for corresponding process, otherwise unexpected result
// or even exception may be caused.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeSaveOptions) SetSaveFormat(value SaveFormat)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LowCodeSaveOptions_SetSaveFormat"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteLowCodeSaveOptions(lowcodesaveoptions *LowCodeSaveOptions){
	runtime.SetFinalizer(lowcodesaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeSaveOptions"),lowcodesaveoptions.ptr)
	lowcodesaveoptions.ptr = nil
}

// Class LowCodeSaveOptionsProviderOfAssembling 

// Implementation to provide save options which save split parts to files
// and the path of resultant file are named as(it may contains directories):
// <see cref="PathHeader"/>+<see cref="SheetPrefix"/>+SheetIndex(or SheetName)
// +<see cref="SplitPartPrefix"/>+SplitPartIndex+<see cref="PathTail"/>.
type LowCodeSaveOptionsProviderOfAssembling struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeSaveOptionsProviderOfAssembling() ( *LowCodeSaveOptionsProviderOfAssembling, error) {
	lowcodesaveoptionsproviderofassembling := &LowCodeSaveOptionsProviderOfAssembling{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeSaveOptionsProviderOfAssembling"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodesaveoptionsproviderofassembling.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodesaveoptionsproviderofassembling, DeleteLowCodeSaveOptionsProviderOfAssembling)
		return lowcodesaveoptionsproviderofassembling, nil
	} else {
		lowcodesaveoptionsproviderofassembling.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodesaveoptionsproviderofassembling, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - AbstractLowCodeSaveOptionsProvider 
func NewLowCodeSaveOptionsProviderOfAssembling_AbstractLowCodeSaveOptionsProvider(src *AbstractLowCodeSaveOptionsProvider) ( *LowCodeSaveOptionsProviderOfAssembling, error) {
	lowcodesaveoptionsproviderofassembling := &LowCodeSaveOptionsProviderOfAssembling{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_LowCodeSaveOptionsProviderOfAssembling_AbstractLowCodeSaveOptionsProvider"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		lowcodesaveoptionsproviderofassembling.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodesaveoptionsproviderofassembling, DeleteLowCodeSaveOptionsProviderOfAssembling)
		return lowcodesaveoptionsproviderofassembling, nil
	} else {
		lowcodesaveoptionsproviderofassembling.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodesaveoptionsproviderofassembling, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfAssembling) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfAssembling_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Header part(before added content of sheet and split part) of file path.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetPathHeader()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetPathHeader"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Header part(before added content of sheet and split part) of file path.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetPathHeader(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetPathHeader"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Tailing part(after sequence numbers) of file path.
// It should include extension of file name.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetPathTail()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetPathTail"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Tailing part(after sequence numbers) of file path.
// It should include extension of file name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetPathTail(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetPathTail"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether builds the file path with sheet name instead of sheet index. Default value is false.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetUseSheetName()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetUseSheetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether builds the file path with sheet name instead of sheet index. Default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetUseSheetName(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetUseSheetName"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Prefix for the index of worksheet.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetSheetPrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetSheetPrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Prefix for the index of worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetSheetPrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetSheetPrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Prefix for the index of split part.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetSplitPartPrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetSplitPartPrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Prefix for the index of split part.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetSplitPartPrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetSplitPartPrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Offset of sheet's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.SheetIndex"/>).
// Returns:
//   int32  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetSheetIndexOffset()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetSheetIndexOffset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Offset of sheet's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.SheetIndex"/>).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetSheetIndexOffset(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetSheetIndexOffset"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Offset of split part's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.PartIndex"/>).
// Returns:
//   int32  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetSplitPartIndexOffset()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetSplitPartIndexOffset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Offset of split part's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.PartIndex"/>).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetSplitPartIndexOffset(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetSplitPartIndexOffset"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether add sheet index or name to file path always.
// Default value is false, that is, when there is only one sheet,
// the sheet index(or name) and corresponding prefix will not be added to the file path.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetBuildPathWithSheetAlways()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetBuildPathWithSheetAlways"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether add sheet index or name to file path always.
// Default value is false, that is, when there is only one sheet,
// the sheet index(or name) and corresponding prefix will not be added to the file path.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetBuildPathWithSheetAlways(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetBuildPathWithSheetAlways"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether add split part index to file path always.
// Default value is false, that is, when there is only one split part,
// the split part index and corresponding prefix will not be added to the file path.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetBuildPathWithSplitPartAlways()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetBuildPathWithSplitPartAlways"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether add split part index to file path always.
// Default value is false, that is, when there is only one split part,
// the split part index and corresponding prefix will not be added to the file path.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetBuildPathWithSplitPartAlways(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetBuildPathWithSplitPartAlways"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The template for creating instance of save options in <see cref="GetSaveOptions(SplitPartInfo)"/>.
// Returns:
//   LowCodeSaveOptions  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetSaveOptionsTemplate()  (*LowCodeSaveOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetSaveOptionsTemplate"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// The template for creating instance of save options in <see cref="GetSaveOptions(SplitPartInfo)"/>.
// Parameters:
//   value - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) SetSaveOptionsTemplate(value *LowCodeSaveOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSaveOptionsProviderOfAssembling_SetSaveOptionsTemplate"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the save options from which to get the output settings for currently split part.
// Parameters:
//   part - SplitPartInfo 
// Returns:
//   LowCodeSaveOptions  
func (instance *LowCodeSaveOptionsProviderOfAssembling) GetSaveOptions(part *SplitPartInfo)  (*LowCodeSaveOptions,  error)  {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZDK(C.CString("LowCodeSaveOptionsProviderOfAssembling_GetSaveOptions"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// Releases resources after processing currently split part.
// Parameters:
//   part - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfAssembling) Finish(part *LowCodeSaveOptions)  error {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSaveOptionsProviderOfAssembling_Finish"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *LowCodeSaveOptionsProviderOfAssembling) ToAbstractLowCodeSaveOptionsProvider() *AbstractLowCodeSaveOptionsProvider {
	parentClass := &AbstractLowCodeSaveOptionsProvider{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLowCodeSaveOptionsProviderOfAssembling(lowcodesaveoptionsproviderofassembling *LowCodeSaveOptionsProviderOfAssembling){
	runtime.SetFinalizer(lowcodesaveoptionsproviderofassembling, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeSaveOptionsProviderOfAssembling"),lowcodesaveoptionsproviderofassembling.ptr)
	lowcodesaveoptionsproviderofassembling.ptr = nil
}

// Class LowCodeSaveOptionsProviderOfPlaceHolders 

// Implementation to provide save options which save split parts to files
// and the path of resultant file are defined with placeholders.
type LowCodeSaveOptionsProviderOfPlaceHolders struct {
	ptr unsafe.Pointer
}

// Instantiates an instance to provide save options according to specified templates.
// Parameters:
//   pathTemplate - string 
func NewLowCodeSaveOptionsProviderOfPlaceHolders_String(pathtemplate string) ( *LowCodeSaveOptionsProviderOfPlaceHolders, error) {
	lowcodesaveoptionsproviderofplaceholders := &LowCodeSaveOptionsProviderOfPlaceHolders{}
	CGoReturnPtr := C.CellsGoFunctoinZZFQ(C.CString("New_LowCodeSaveOptionsProviderOfPlaceHolders_String"),C.CString(pathtemplate))
	if CGoReturnPtr.error_no == 0 {
		lowcodesaveoptionsproviderofplaceholders.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodesaveoptionsproviderofplaceholders, DeleteLowCodeSaveOptionsProviderOfPlaceHolders)
		return lowcodesaveoptionsproviderofplaceholders, nil
	} else {
		lowcodesaveoptionsproviderofplaceholders.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodesaveoptionsproviderofplaceholders, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - AbstractLowCodeSaveOptionsProvider 
func NewLowCodeSaveOptionsProviderOfPlaceHolders_AbstractLowCodeSaveOptionsProvider(src *AbstractLowCodeSaveOptionsProvider) ( *LowCodeSaveOptionsProviderOfPlaceHolders, error) {
	lowcodesaveoptionsproviderofplaceholders := &LowCodeSaveOptionsProviderOfPlaceHolders{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_LowCodeSaveOptionsProviderOfPlaceHolders_AbstractLowCodeSaveOptionsProvider"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		lowcodesaveoptionsproviderofplaceholders.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodesaveoptionsproviderofplaceholders, DeleteLowCodeSaveOptionsProviderOfPlaceHolders)
		return lowcodesaveoptionsproviderofplaceholders, nil
	} else {
		lowcodesaveoptionsproviderofplaceholders.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodesaveoptionsproviderofplaceholders, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Offset of sheet's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.SheetIndex"/>).
// Returns:
//   int32  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSheetIndexOffset()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSheetIndexOffset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Offset of sheet's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.SheetIndex"/>).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetSheetIndexOffset(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetSheetIndexOffset"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Offset of split part's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.PartIndex"/>).
// Returns:
//   int32  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSplitPartIndexOffset()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSplitPartIndexOffset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Offset of split part's index between what used in file path
// and its actual value(<see cref="SplitPartInfo.PartIndex"/>).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetSplitPartIndexOffset(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetSplitPartIndexOffset"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether add sheet index or name to file path always.
// Default value is false, that is, when there is only one sheet,
// the sheet index and name and corresponding prefix(<see cref="SheetNamePrefix"/>)
// will not be added to the file path.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetBuildPathWithSheetAlways()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetBuildPathWithSheetAlways"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether add sheet index or name to file path always.
// Default value is false, that is, when there is only one sheet,
// the sheet index and name and corresponding prefix(<see cref="SheetNamePrefix"/>)
// will not be added to the file path.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetBuildPathWithSheetAlways(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetBuildPathWithSheetAlways"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether add split part index to file path always.
// Default value is false, that is, when there is only one split part,
// the split part index and corresponding prefix(<see cref="SplitPartPrefix"/>)
// will not be added to the file path.
// Returns:
//   bool  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetBuildPathWithSplitPartAlways()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetBuildPathWithSplitPartAlways"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether add split part index to file path always.
// Default value is false, that is, when there is only one split part,
// the split part index and corresponding prefix(<see cref="SplitPartPrefix"/>)
// will not be added to the file path.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetBuildPathWithSplitPartAlways(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetBuildPathWithSplitPartAlways"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Prefix for the index of worksheet.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSheetNamePrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSheetNamePrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Prefix for the index of worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetSheetNamePrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetSheetNamePrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Prefix for the index of worksheet.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSheetIndexPrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSheetIndexPrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Prefix for the index of worksheet.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetSheetIndexPrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetSheetIndexPrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Prefix for the index of split part.
// Returns:
//   string  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSplitPartPrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSplitPartPrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Prefix for the index of split part.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetSplitPartPrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetSplitPartPrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The template for creating instance of save options in <see cref="GetSaveOptions(SplitPartInfo)"/>.
// Returns:
//   LowCodeSaveOptions  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSaveOptionsTemplate()  (*LowCodeSaveOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSaveOptionsTemplate"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// The template for creating instance of save options in <see cref="GetSaveOptions(SplitPartInfo)"/>.
// Parameters:
//   value - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) SetSaveOptionsTemplate(value *LowCodeSaveOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_SetSaveOptionsTemplate"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the save options from which to get the output settings for currently split part.
// Parameters:
//   part - SplitPartInfo 
// Returns:
//   LowCodeSaveOptions  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) GetSaveOptions(part *SplitPartInfo)  (*LowCodeSaveOptions,  error)  {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZDK(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_GetSaveOptions"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// Releases resources after processing currently split part.
// Parameters:
//   part - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) Finish(part *LowCodeSaveOptions)  error {
	
	var part_ptr unsafe.Pointer = nil
	if part != nil {
	  part_ptr =part.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSaveOptionsProviderOfPlaceHolders_Finish"), instance.ptr, part_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *LowCodeSaveOptionsProviderOfPlaceHolders) ToAbstractLowCodeSaveOptionsProvider() *AbstractLowCodeSaveOptionsProvider {
	parentClass := &AbstractLowCodeSaveOptionsProvider{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLowCodeSaveOptionsProviderOfPlaceHolders(lowcodesaveoptionsproviderofplaceholders *LowCodeSaveOptionsProviderOfPlaceHolders){
	runtime.SetFinalizer(lowcodesaveoptionsproviderofplaceholders, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeSaveOptionsProviderOfPlaceHolders"),lowcodesaveoptionsproviderofplaceholders.ptr)
	lowcodesaveoptionsproviderofplaceholders.ptr = nil
}

// Class LowCodeSplitOptions 

// Options for splitting spreadsheet.
type LowCodeSplitOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewLowCodeSplitOptions() ( *LowCodeSplitOptions, error) {
	lowcodesplitoptions := &LowCodeSplitOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_LowCodeSplitOptions"),)
	if CGoReturnPtr.error_no == 0 {
		lowcodesplitoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(lowcodesplitoptions, DeleteLowCodeSplitOptions)
		return lowcodesplitoptions, nil
	} else {
		lowcodesplitoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return lowcodesplitoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LowCodeSplitOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LowCodeSplitOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Load options for loading the spreadsheet that will be split.
// Returns:
//   LowCodeLoadOptions  
func (instance *LowCodeSplitOptions) GetLoadOptions()  (*LowCodeLoadOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeSplitOptions_GetLoadOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeLoadOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeLoadOptions) 

	return result, nil 
}
// Load options for loading the spreadsheet that will be split.
// Parameters:
//   value - LowCodeLoadOptions 
// Returns:
//   void  
func (instance *LowCodeSplitOptions) SetLoadOptions(value *LowCodeLoadOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSplitOptions_SetLoadOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Save options for saving the split parts.
// Returns:
//   LowCodeSaveOptions  
func (instance *LowCodeSplitOptions) GetSaveOptions()  (*LowCodeSaveOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeSplitOptions_GetSaveOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LowCodeSaveOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLowCodeSaveOptions) 

	return result, nil 
}
// Save options for saving the split parts.
// Parameters:
//   value - LowCodeSaveOptions 
// Returns:
//   void  
func (instance *LowCodeSplitOptions) SetSaveOptions(value *LowCodeSaveOptions)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSplitOptions_SetSaveOptions"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Provider of save options for saving the split parts.
// Returns:
//   AbstractLowCodeSaveOptionsProvider  
func (instance *LowCodeSplitOptions) GetSaveOptionsProvider()  (*AbstractLowCodeSaveOptionsProvider,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LowCodeSplitOptions_GetSaveOptionsProvider"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AbstractLowCodeSaveOptionsProvider{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAbstractLowCodeSaveOptionsProvider) 

	return result, nil 
}
// Provider of save options for saving the split parts.
// Parameters:
//   value - AbstractLowCodeSaveOptionsProvider 
// Returns:
//   void  
func (instance *LowCodeSplitOptions) SetSaveOptionsProvider(value *AbstractLowCodeSaveOptionsProvider)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LowCodeSplitOptions_SetSaveOptionsProvider"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteLowCodeSplitOptions(lowcodesplitoptions *LowCodeSplitOptions){
	runtime.SetFinalizer(lowcodesplitoptions, nil)
	C.Delete_CObject(C.CString("Delete_LowCodeSplitOptions"),lowcodesplitoptions.ptr)
	lowcodesplitoptions.ptr = nil
}

// Class PdfConverter 

// Converter for converting template file to pdf.
type PdfConverter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PdfConverter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfConverter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts given template file to pdf.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func PdfConverter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("PdfConverter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts template file to pdf
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
// Returns:
//   void  
func PdfConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("PdfConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions"),loadoptions_ptr, saveoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeletePdfConverter(pdfconverter *PdfConverter){
	runtime.SetFinalizer(pdfconverter, nil)
	C.Delete_CObject(C.CString("Delete_PdfConverter"),pdfconverter.ptr)
	pdfconverter.ptr = nil
}

// Class SplitPartInfo 

// Represents the information of one input/output for multiple inputs/outputs,
// such as current page to be rendered when converting spreadsheet to image.
type SplitPartInfo struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SplitPartInfo) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SplitPartInfo_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Index of current part in sequence(0 based).
// -1 means there are no multiple parts so the result is single.
// Returns:
//   int32  
func (instance *SplitPartInfo) GetPartIndex()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SplitPartInfo_GetPartIndex"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Index of the sheet where current part is in. -1 denotes there is only one sheet.
// Returns:
//   int32  
func (instance *SplitPartInfo) GetSheetIndex()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SplitPartInfo_GetSheetIndex"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Name of the sheet where current part is in.
// Returns:
//   string  
func (instance *SplitPartInfo) GetSheetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SplitPartInfo_GetSheetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSplitPartInfo(splitpartinfo *SplitPartInfo){
	runtime.SetFinalizer(splitpartinfo, nil)
	C.Delete_CObject(C.CString("Delete_SplitPartInfo"),splitpartinfo.ptr)
	splitpartinfo.ptr = nil
}

// Class SpreadsheetConverter 

// Converter for conversion between different spreadsheet file formats, such as xls, xlsx, xlsb, spreadsheet ml...
type SpreadsheetConverter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SpreadsheetConverter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpreadsheetConverter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts given template file between spreadsheet file formats.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func SpreadsheetConverter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("SpreadsheetConverter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts between different spreadsheet file formats.
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
// Returns:
//   void  
func SpreadsheetConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SpreadsheetConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions"),loadoptions_ptr, saveoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSpreadsheetConverter(spreadsheetconverter *SpreadsheetConverter){
	runtime.SetFinalizer(spreadsheetconverter, nil)
	C.Delete_CObject(C.CString("Delete_SpreadsheetConverter"),spreadsheetconverter.ptr)
	spreadsheetconverter.ptr = nil
}

// Class SpreadsheetLocker 

// Low code api to lock spreadsheet file.
type SpreadsheetLocker struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SpreadsheetLocker) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpreadsheetLocker_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Locks spreadsheet file with specified settings.
// Parameters:
//   templateFile - string 
//   resultFile - string 
//   openPassword - string 
//   writePassword - string 
// Returns:
//   void  
func SpreadsheetLocker_Process_String_String_String_String(templatefile string, resultfile string, openpassword string, writepassword string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKO(C.CString("SpreadsheetLocker_Process_String_String_String_String"),C.CString(templatefile), C.CString(resultfile), C.CString(openpassword), C.CString(writepassword))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Locks spreadsheet file with specified settings.
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
//   openPassword - string 
//   writePassword - string 
// Returns:
//   void  
func SpreadsheetLocker_Process_LowCodeLoadOptions_LowCodeSaveOptions_String_String(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions, openpassword string, writepassword string)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZKP(C.CString("SpreadsheetLocker_Process_LowCodeLoadOptions_LowCodeSaveOptions_String_String"),loadoptions_ptr, saveoptions_ptr, C.CString(openpassword), C.CString(writepassword))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Locks spreadsheet file with specified settings.
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
//   openPassword - string 
//   writePassword - string 
//   workbookPassword - string 
//   workbookType - int32 
// Returns:
//   void  
func SpreadsheetLocker_Process_LowCodeLoadOptions_LowCodeSaveOptions_String_String_String_ProtectionType(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions, openpassword string, writepassword string, workbookpassword string, workbooktype ProtectionType)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZKQ(C.CString("SpreadsheetLocker_Process_LowCodeLoadOptions_LowCodeSaveOptions_String_String_String_ProtectionType"),loadoptions_ptr, saveoptions_ptr, C.CString(openpassword), C.CString(writepassword), C.CString(workbookpassword), C.int( int32(workbooktype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Locks spreadsheet file with specified settings.
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
//   provider - AbstractLowCodeProtectionProvider 
// Returns:
//   void  
func SpreadsheetLocker_Process_LowCodeLoadOptions_LowCodeSaveOptions_AbstractLowCodeProtectionProvider(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions, provider *AbstractLowCodeProtectionProvider)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}
	var provider_ptr unsafe.Pointer = nil
	if provider != nil {
	  provider_ptr =provider.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZAP(C.CString("SpreadsheetLocker_Process_LowCodeLoadOptions_LowCodeSaveOptions_AbstractLowCodeProtectionProvider"),loadoptions_ptr, saveoptions_ptr, provider_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSpreadsheetLocker(spreadsheetlocker *SpreadsheetLocker){
	runtime.SetFinalizer(spreadsheetlocker, nil)
	C.Delete_CObject(C.CString("Delete_SpreadsheetLocker"),spreadsheetlocker.ptr)
	spreadsheetlocker.ptr = nil
}

// Class SpreadsheetMerger 

// Merges multiple template files into one.
type SpreadsheetMerger struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SpreadsheetMerger) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpreadsheetMerger_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Merge given template files.
// Parameters:
//   templateFiles - []string 
//   resultFile - string 
// Returns:
//   void  
func SpreadsheetMerger_Process_stringArray_String(templatefiles []string, resultfile string)  error {
	
	vector_templatefiles := make([]*C.char, len(templatefiles))
	for i, str := range templatefiles {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	vector_templatefiles[i] = cStr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZKR(C.CString("SpreadsheetMerger_Process_U16StringArray_String"),unsafe.Pointer(&vector_templatefiles[0]), C.int( len(templatefiles)), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Merges multiple template files into one.
// Parameters:
//   options - LowCodeMergeOptions 
// Returns:
//   void  
func SpreadsheetMerger_Process_LowCodeMergeOptions(options *LowCodeMergeOptions)  error {
	
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("SpreadsheetMerger_Process_LowCodeMergeOptions"),options_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSpreadsheetMerger(spreadsheetmerger *SpreadsheetMerger){
	runtime.SetFinalizer(spreadsheetmerger, nil)
	C.Delete_CObject(C.CString("Delete_SpreadsheetMerger"),spreadsheetmerger.ptr)
	spreadsheetmerger.ptr = nil
}

// Class SpreadsheetSplitter 

// Splits spreadsheet file into multiple parts.
type SpreadsheetSplitter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SpreadsheetSplitter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpreadsheetSplitter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Splits given template file into multiple parts.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func SpreadsheetSplitter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("SpreadsheetSplitter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Splits spreadsheet file into multiple parts.
// Parameters:
//   options - LowCodeSplitOptions 
// Returns:
//   void  
func SpreadsheetSplitter_Process_LowCodeSplitOptions(options *LowCodeSplitOptions)  error {
	
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("SpreadsheetSplitter_Process_LowCodeSplitOptions"),options_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSpreadsheetSplitter(spreadsheetsplitter *SpreadsheetSplitter){
	runtime.SetFinalizer(spreadsheetsplitter, nil)
	C.Delete_CObject(C.CString("Delete_SpreadsheetSplitter"),spreadsheetsplitter.ptr)
	spreadsheetsplitter.ptr = nil
}

// Class TextConverter 

// Converter for conversion between text based formats(csv, tsv, dif...) and other spreadsheet file formats.
type TextConverter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextConverter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextConverter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Converts given template file between text based files and other formats.
// Parameters:
//   templateFile - string 
//   resultFile - string 
// Returns:
//   void  
func TextConverter_Process_String_String(templatefile string, resultfile string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZKN(C.CString("TextConverter_Process_String_String"),C.CString(templatefile), C.CString(resultfile))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Converts file format between text based formats and other spreadsheet file formats
// Parameters:
//   loadOptions - LowCodeLoadOptions 
//   saveOptions - LowCodeSaveOptions 
// Returns:
//   void  
func TextConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions(loadoptions *LowCodeLoadOptions, saveoptions *LowCodeSaveOptions)  error {
	
	var loadoptions_ptr unsafe.Pointer = nil
	if loadoptions != nil {
	  loadoptions_ptr =loadoptions.ptr
	}
	var saveoptions_ptr unsafe.Pointer = nil
	if saveoptions != nil {
	  saveoptions_ptr =saveoptions.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("TextConverter_Process_LowCodeLoadOptions_LowCodeSaveOptions"),loadoptions_ptr, saveoptions_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteTextConverter(textconverter *TextConverter){
	runtime.SetFinalizer(textconverter, nil)
	C.Delete_CObject(C.CString("Delete_TextConverter"),textconverter.ptr)
	textconverter.ptr = nil
}
