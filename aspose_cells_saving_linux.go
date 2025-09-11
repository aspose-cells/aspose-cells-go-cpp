// +build linux

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/linux_x86_64" -L"${SRCDIR}/lib/linux_x86_64" -lAspose.Cells.CWrapper
// #include <CellsFunctionMap.h>
import "C"
import (
	"fmt"  
 	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum SaveElementType *****************/

// Represents what kind of elements should be saved.
type SaveElementType int32

const(
// All data.
SaveElementType_All SaveElementType = 251658239 

// Only charts.
SaveElementType_Chart SaveElementType = 256 
)

func Int32ToSaveElementType(value int32)(SaveElementType ,error){
	switch value {
		case 251658239:  return SaveElementType_All, nil  
		case 256:  return SaveElementType_Chart, nil  
		default:
			return 0 ,fmt.Errorf("invalid SaveElementType value: %d", value)
	}
}

/**************Enum SqlScriptOperatorType *****************/

// Represents the type of operating data.
type SqlScriptOperatorType int32

const(
// Insert data.
SqlScriptOperatorType_Insert SqlScriptOperatorType = 0 

// Update data.
SqlScriptOperatorType_Update SqlScriptOperatorType = 1 

// Delete data.
SqlScriptOperatorType_Delete SqlScriptOperatorType = 2 
)

func Int32ToSqlScriptOperatorType(value int32)(SqlScriptOperatorType ,error){
	switch value {
		case 0:  return SqlScriptOperatorType_Insert, nil  
		case 1:  return SqlScriptOperatorType_Update, nil  
		case 2:  return SqlScriptOperatorType_Delete, nil  
		default:
			return 0 ,fmt.Errorf("invalid SqlScriptOperatorType value: %d", value)
	}
}
// Class DbfSaveOptions 

// Represents the options of saving dbf file
type DbfSaveOptions struct {
	ptr unsafe.Pointer
}

// The options of saving .dbf file.
func NewDbfSaveOptions() ( *DbfSaveOptions, error) {
	dbfsaveoptions := &DbfSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_DbfSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		dbfsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(dbfsaveoptions, DeleteDbfSaveOptions)
		return dbfsaveoptions, nil
	} else {
		dbfsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return dbfsaveoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - SaveOptions 
func NewDbfSaveOptions_SaveOptions(src *SaveOptions) ( *DbfSaveOptions, error) {
	dbfsaveoptions := &DbfSaveOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZLB(C.CString("New_DbfSaveOptions_SaveOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		dbfsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(dbfsaveoptions, DeleteDbfSaveOptions)
		return dbfsaveoptions, nil
	} else {
		dbfsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return dbfsaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DbfSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting as string value
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetExportAsString()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetExportAsString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting as string value
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetExportAsString(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetExportAsString"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the save file format.
// Returns:
//   int32  
func (instance *DbfSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZLC(C.CString("DbfSaveOptions_GetSaveFormat"), instance.ptr)
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
// Make the workbook empty after saving the file.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetClearData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetClearData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Make the workbook empty after saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetClearData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetClearData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The folder for temporary files that may be used as data cache.
// Returns:
//   string  
func (instance *DbfSaveOptions) GetCachedFileFolder()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("DbfSaveOptions_GetCachedFileFolder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The folder for temporary files that may be used as data cache.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetCachedFileFolder(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("DbfSaveOptions_SetCachedFileFolder"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether validate merged cells before saving the file.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetValidateMergedAreas()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetValidateMergedAreas"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether validate merged cells before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetValidateMergedAreas(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetValidateMergedAreas"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether merge the areas of conditional formatting and validation before saving the file.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetMergeAreas()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetMergeAreas"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether merge the areas of conditional formatting and validation before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetMergeAreas(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetMergeAreas"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If true and the directory does not exist, the directory will be automatically created before saving the file.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetCreateDirectory()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetCreateDirectory"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// If true and the directory does not exist, the directory will be automatically created before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetCreateDirectory(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetCreateDirectory"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether sorting defined names before saving file.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetSortNames()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetSortNames"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether sorting defined names before saving file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetSortNames(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetSortNames"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether sorting external defined names before saving file.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetSortExternalNames()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetSortExternalNames"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether sorting external defined names before saving file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetSortExternalNames(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetSortExternalNames"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether refreshing chart cache data
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetRefreshChartCache()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetRefreshChartCache"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether refreshing chart cache data
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetRefreshChartCache(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetRefreshChartCache"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K, it will be truncated.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetCheckExcelRestriction()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetCheckExcelRestriction"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K, it will be truncated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetCheckExcelRestriction(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetCheckExcelRestriction"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether updating smart art setting.
// The default value is false.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetUpdateSmartArt()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetUpdateSmartArt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether updating smart art setting.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetUpdateSmartArt(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetUpdateSmartArt"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether encrypt document properties when saving as .xls file.
// The default value is true.
// Returns:
//   bool  
func (instance *DbfSaveOptions) GetEncryptDocumentProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DbfSaveOptions_GetEncryptDocumentProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether encrypt document properties when saving as .xls file.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetEncryptDocumentProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DbfSaveOptions_SetEncryptDocumentProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *DbfSaveOptions) ToSaveOptions() *SaveOptions {
	parentClass := &SaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDbfSaveOptions(dbfsaveoptions *DbfSaveOptions){
	runtime.SetFinalizer(dbfsaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_DbfSaveOptions"),dbfsaveoptions.ptr)
	dbfsaveoptions.ptr = nil
}

// Class EbookSaveOptions 

// Represents the options for saving ebook file.
type EbookSaveOptions struct {
	ptr unsafe.Pointer
}

// Creates options for saving ebook file.
func NewEbookSaveOptions() ( *EbookSaveOptions, error) {
	ebooksaveoptions := &EbookSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_EbookSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		ebooksaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(ebooksaveoptions, DeleteEbookSaveOptions)
		return ebooksaveoptions, nil
	} else {
		ebooksaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return ebooksaveoptions, err
	}	
}
// Creates options for saving ebook file.
// Parameters:
//   saveFormat - int32 
func NewEbookSaveOptions_SaveFormat(saveformat SaveFormat) ( *EbookSaveOptions, error) {
	ebooksaveoptions := &EbookSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZPQ(C.CString("New_EbookSaveOptions_SaveFormat"),C.int( int32(saveformat)))
	if CGoReturnPtr.error_no == 0 {
		ebooksaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(ebooksaveoptions, DeleteEbookSaveOptions)
		return ebooksaveoptions, nil
	} else {
		ebooksaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return ebooksaveoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - HtmlSaveOptions 
func NewEbookSaveOptions_HtmlSaveOptions(src *HtmlSaveOptions) ( *EbookSaveOptions, error) {
	ebooksaveoptions := &EbookSaveOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZBGN(C.CString("New_EbookSaveOptions_HtmlSaveOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		ebooksaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(ebooksaveoptions, DeleteEbookSaveOptions)
		return ebooksaveoptions, nil
	} else {
		ebooksaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return ebooksaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicate whether exporting those not visible shapes
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetIgnoreInvisibleShapes()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetIgnoreInvisibleShapes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicate whether exporting those not visible shapes
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIgnoreInvisibleShapes(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIgnoreInvisibleShapes"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The title of the html page.
// Only for saving to html stream.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetPageTitle()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetPageTitle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The title of the html page.
// Only for saving to html stream.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetPageTitle(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetPageTitle"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The directory that the attached files will be saved to.
// Only for saving to html stream.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetAttachedFilesDirectory()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetAttachedFilesDirectory"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The directory that the attached files will be saved to.
// Only for saving to html stream.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetAttachedFilesDirectory(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetAttachedFilesDirectory"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specify the Url prefix of attached files such as image in the html file.
// Only for saving to html stream.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetAttachedFilesUrlPrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetAttachedFilesUrlPrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specify the Url prefix of attached files such as image in the html file.
// Only for saving to html stream.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetAttachedFilesUrlPrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetAttachedFilesUrlPrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specify the default font name for exporting html, the default font will be used  when the font of style is not existing,
// If this property is null, Aspose.Cells will use universal font which have the same family with the original font,
// the default value is null.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetDefaultFontName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetDefaultFontName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specify the default font name for exporting html, the default font will be used  when the font of style is not existing,
// If this property is null, Aspose.Cells will use universal font which have the same family with the original font,
// the default value is null.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetDefaultFontName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetDefaultFontName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to add a generic font to CSS font-family.
// The default value is true
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetAddGenericFont()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetAddGenericFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to add a generic font to CSS font-family.
// The default value is true
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetAddGenericFont(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetAddGenericFont"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if zooming in or out the html via worksheet zoom level when saving file to html, the default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetWorksheetScalable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetWorksheetScalable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if zooming in or out the html via worksheet zoom level when saving file to html, the default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetWorksheetScalable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetWorksheetScalable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if exporting comments when saving file to html, the default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsExportComments()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsExportComments"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if exporting comments when saving file to html, the default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsExportComments(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsExportComments"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents type of exporting comments to html files.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetExportCommentsType()  (PrintCommentsType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPR(C.CString("EbookSaveOptions_GetExportCommentsType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPrintCommentsType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents type of exporting comments to html files.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportCommentsType(value PrintCommentsType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPS(C.CString("EbookSaveOptions_SetExportCommentsType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if disable Downlevel-revealed conditional comments when exporting file to html, the default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetDisableDownlevelRevealedComments()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetDisableDownlevelRevealedComments"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if disable Downlevel-revealed conditional comments when exporting file to html, the default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetDisableDownlevelRevealedComments(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetDisableDownlevelRevealedComments"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting image files to temp directory.
// Only for saving to html stream.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsExpImageToTempDir()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsExpImageToTempDir"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting image files to temp directory.
// Only for saving to html stream.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsExpImageToTempDir(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsExpImageToTempDir"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether using scalable unit to describe the image width
// when using scalable unit to describe the column width.
// The default value is true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetImageScalable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetImageScalable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether using scalable unit to describe the image width
// when using scalable unit to describe the column width.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetImageScalable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetImageScalable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting column width in unit of scale to html.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetWidthScalable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetWidthScalable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting column width in unit of scale to html.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetWidthScalable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetWidthScalable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting the single tab when the file only has one worksheet.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportSingleTab()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportSingleTab"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting the single tab when the file only has one worksheet.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportSingleTab(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportSingleTab"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether images are saved in Base64 format to HTML, MHTML or EPUB.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportImagesAsBase64()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportImagesAsBase64"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether images are saved in Base64 format to HTML, MHTML or EPUB.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportImagesAsBase64(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportImagesAsBase64"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if only exporting the active worksheet to html file.
// If true then only the active worksheet will be exported to html file;
// If false then the whole workbook will be exported to html file.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportActiveWorksheetOnly()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportActiveWorksheetOnly"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if only exporting the active worksheet to html file.
// If true then only the active worksheet will be exported to html file;
// If false then the whole workbook will be exported to html file.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportActiveWorksheetOnly(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportActiveWorksheetOnly"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if only exporting the print area to html file. The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportPrintAreaOnly()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportPrintAreaOnly"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if only exporting the print area to html file. The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportPrintAreaOnly(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportPrintAreaOnly"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or Sets the exporting CellArea of current active Worksheet.
// If you set this attribute, the print area of current active Worksheet will be omitted.
// Only the specified area will be exported when saving the file to html.
// Returns:
//   CellArea  
func (instance *EbookSaveOptions) GetExportArea()  (*CellArea,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("EbookSaveOptions_GetExportArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets or Sets the exporting CellArea of current active Worksheet.
// If you set this attribute, the print area of current active Worksheet will be omitted.
// Only the specified area will be exported when saving the file to html.
// Parameters:
//   value - CellArea 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportArea(value *CellArea)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZGH(C.CString("EbookSaveOptions_SetExportArea"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether html tag(such as <c>&lt;div&gt;&lt;/div&gt;</c>) in cell should be parsed as cell value or preserved as it is.
// The default value is true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetParseHtmlTagInCell()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetParseHtmlTagInCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether html tag(such as <c>&lt;div&gt;&lt;/div&gt;</c>) in cell should be parsed as cell value or preserved as it is.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetParseHtmlTagInCell(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetParseHtmlTagInCell"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if a cross-cell string will be displayed in the same way as MS Excel when saving an Excel file in html format.
// By default the value is Default, so, for cross-cell strings, there is little difference between the html files created by Aspose.Cells and MS Excel.
// But the performance for creating large html files,setting the value to Cross would be several times faster than setting it to Default or Fit2Cell.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetHtmlCrossStringType()  (HtmlCrossType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPT(C.CString("EbookSaveOptions_GetHtmlCrossStringType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlCrossType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates if a cross-cell string will be displayed in the same way as MS Excel when saving an Excel file in html format.
// By default the value is Default, so, for cross-cell strings, there is little difference between the html files created by Aspose.Cells and MS Excel.
// But the performance for creating large html files,setting the value to Cross would be several times faster than setting it to Default or Fit2Cell.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetHtmlCrossStringType(value HtmlCrossType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPU(C.CString("EbookSaveOptions_SetHtmlCrossStringType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Hidden column(the width of this column is 0) in excel,before save this into html format,
// if HtmlHiddenColDisplayType is "Remove",the hidden column would not been output,
// if the value is "Hidden", the column would been output,but was hidden,the default value is "Hidden"
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetHiddenColDisplayType()  (HtmlHiddenColDisplayType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPV(C.CString("EbookSaveOptions_GetHiddenColDisplayType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlHiddenColDisplayType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Hidden column(the width of this column is 0) in excel,before save this into html format,
// if HtmlHiddenColDisplayType is "Remove",the hidden column would not been output,
// if the value is "Hidden", the column would been output,but was hidden,the default value is "Hidden"
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetHiddenColDisplayType(value HtmlHiddenColDisplayType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPW(C.CString("EbookSaveOptions_SetHiddenColDisplayType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Hidden row(the height of this row is 0) in excel,before save this into html format,
// if HtmlHiddenRowDisplayType is "Remove",the hidden row would not been output,
// if the value is "Hidden", the row would been output,but was hidden,the default value is "Hidden"
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetHiddenRowDisplayType()  (HtmlHiddenRowDisplayType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPX(C.CString("EbookSaveOptions_GetHiddenRowDisplayType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlHiddenRowDisplayType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Hidden row(the height of this row is 0) in excel,before save this into html format,
// if HtmlHiddenRowDisplayType is "Remove",the hidden row would not been output,
// if the value is "Hidden", the row would been output,but was hidden,the default value is "Hidden"
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetHiddenRowDisplayType(value HtmlHiddenRowDisplayType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPY(C.CString("EbookSaveOptions_SetHiddenRowDisplayType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If not set,use Encoding.UTF8 as default enconding type.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetEncoding()  (EncodingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("EbookSaveOptions_GetEncoding"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEncodingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// If not set,use Encoding.UTF8 as default enconding type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetEncoding(value EncodingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("EbookSaveOptions_SetEncoding"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get the ImageOrPrintOptions object before exporting
// Returns:
//   ImageOrPrintOptions  
func (instance *EbookSaveOptions) GetImageOptions()  (*ImageOrPrintOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQA(C.CString("EbookSaveOptions_GetImageOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ImageOrPrintOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteImageOrPrintOptions) 

	return result, nil 
}
// Indicates whether save the html as single file.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetSaveAsSingleFile()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetSaveAsSingleFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether save the html as single file.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetSaveAsSingleFile(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetSaveAsSingleFile"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing all sheets when saving  as a single html file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetShowAllSheets()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetShowAllSheets"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether showing all sheets when saving  as a single html file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetShowAllSheets(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetShowAllSheets"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting page headers.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportPageHeaders()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportPageHeaders"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting page headers.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportPageHeaders(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportPageHeaders"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting page headers.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportPageFooters()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportPageFooters"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting page headers.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportPageFooters(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportPageFooters"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating if exporting the hidden worksheet content.The default value is true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportHiddenWorksheet()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportHiddenWorksheet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating if exporting the hidden worksheet content.The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportHiddenWorksheet(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportHiddenWorksheet"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating if html or mht file is presentation preference.
// The default value is false.
// if you want to get more beautiful presentation,please set the value to true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetPresentationPreference()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetPresentationPreference"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating if html or mht file is presentation preference.
// The default value is false.
// if you want to get more beautiful presentation,please set the value to true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetPresentationPreference(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetPresentationPreference"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the prefix of the css name,the default value is "".
// Returns:
//   string  
func (instance *EbookSaveOptions) GetCellCssPrefix()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetCellCssPrefix"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the prefix of the css name,the default value is "".
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCellCssPrefix(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetCellCssPrefix"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the prefix of the type css name such as tr,col,td and so on, they are contained in the table element
// which has the specific TableCssId attribute. The default value is "".
// Returns:
//   string  
func (instance *EbookSaveOptions) GetTableCssId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetTableCssId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the prefix of the type css name such as tr,col,td and so on, they are contained in the table element
// which has the specific TableCssId attribute. The default value is "".
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetTableCssId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetTableCssId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether using full path link in sheet00x.htm,filelist.xml and tabstrip.htm.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsFullPathLink()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsFullPathLink"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether using full path link in sheet00x.htm,filelist.xml and tabstrip.htm.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsFullPathLink(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsFullPathLink"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether export the worksheet css separately.The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportWorksheetCSSSeparately()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportWorksheetCSSSeparately"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether export the worksheet css separately.The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportWorksheetCSSSeparately(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportWorksheetCSSSeparately"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting the similar border style when the border style is not supported by browsers.
// If you want to import the html or mht file to excel, please keep the default value.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportSimilarBorderStyle()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportSimilarBorderStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting the similar border style when the border style is not supported by browsers.
// If you want to import the html or mht file to excel, please keep the default value.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportSimilarBorderStyle(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportSimilarBorderStyle"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The option to merge contiguous empty cells(empty td elements)
// The default value is MergeEmptyTdType.Default.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetMergeEmptyTdType()  (MergeEmptyTdType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQB(C.CString("EbookSaveOptions_GetMergeEmptyTdType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMergeEmptyTdType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// The option to merge contiguous empty cells(empty td elements)
// The default value is MergeEmptyTdType.Default.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetMergeEmptyTdType(value MergeEmptyTdType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQC(C.CString("EbookSaveOptions_SetMergeEmptyTdType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting excel coordinate of nonblank cells when saving file to html. The default value is false.
// If you want to import the output html to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportCellCoordinate()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportCellCoordinate"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting excel coordinate of nonblank cells when saving file to html. The default value is false.
// If you want to import the output html to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportCellCoordinate(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportCellCoordinate"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting extra headings when the length of text is longer than max display column.
// The default value is false. If you want to import the html file to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportExtraHeadings()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportExtraHeadings"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting extra headings when the length of text is longer than max display column.
// The default value is false. If you want to import the html file to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportExtraHeadings(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportExtraHeadings"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exports sheet's row and column headings when saving to HTML files.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportRowColumnHeadings()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportRowColumnHeadings"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exports sheet's row and column headings when saving to HTML files.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportRowColumnHeadings(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportRowColumnHeadings"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting formula when saving file to html. The default value is true.
// If you want to import the output html to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportFormula()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting formula when saving file to html. The default value is true.
// If you want to import the output html to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportFormula(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportFormula"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether adding tooltip text when the data can't be fully displayed.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetAddTooltipText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetAddTooltipText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether adding tooltip text when the data can't be fully displayed.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetAddTooltipText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetAddTooltipText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting the gridlines.The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportGridLines()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportGridLines"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting the gridlines.The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportGridLines(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportGridLines"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting bogus bottom row data. The default value is true.If you want to import the html or mht file
// to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportBogusRowData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportBogusRowData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting bogus bottom row data. The default value is true.If you want to import the html or mht file
// to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportBogusRowData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportBogusRowData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether excludes unused styles.
// For the generated html files, excluding unused styles can make the file size smaller
// without affecting the visual effects. So the default value of this property is true.
// If user needs to keep all styles in the workbook for the generated html(such as the scenario that user
// needs to restore the workbook from the generated html later), please set this property as false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExcludeUnusedStyles()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExcludeUnusedStyles"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether excludes unused styles.
// For the generated html files, excluding unused styles can make the file size smaller
// without affecting the visual effects. So the default value of this property is true.
// If user needs to keep all styles in the workbook for the generated html(such as the scenario that user
// needs to restore the workbook from the generated html later), please set this property as false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExcludeUnusedStyles(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExcludeUnusedStyles"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting document properties.The default value is true.If you want to import
// the html or mht file to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportDocumentProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportDocumentProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting document properties.The default value is true.If you want to import
// the html or mht file to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportDocumentProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportDocumentProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting worksheet properties.The default value is true.If you want to import
// the html or mht file to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportWorksheetProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportWorksheetProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting worksheet properties.The default value is true.If you want to import
// the html or mht file to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportWorksheetProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportWorksheetProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting workbook properties.The default value is true.If you want to import
// the html or mht file to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportWorkbookProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportWorkbookProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting workbook properties.The default value is true.If you want to import
// the html or mht file to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportWorkbookProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportWorkbookProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether exporting frame scripts and document properties. The default value is true.If you want to import the html or mht file
// to excel, please keep the default value.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetExportFrameScriptsAndProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetExportFrameScriptsAndProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether exporting frame scripts and document properties. The default value is true.If you want to import the html or mht file
// to excel, please keep the default value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportFrameScriptsAndProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetExportFrameScriptsAndProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating the rule of exporting html file data.The default value is All.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetExportDataOptions()  (HtmlExportDataOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQD(C.CString("EbookSaveOptions_GetExportDataOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlExportDataOptions(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicating the rule of exporting html file data.The default value is All.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetExportDataOptions(value HtmlExportDataOptions)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQE(C.CString("EbookSaveOptions_SetExportDataOptions"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating the type of target attribute in <c>&lt;a&gt;</c> link. The default value is HtmlLinkTargetType.Parent.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetLinkTargetType()  (HtmlLinkTargetType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQF(C.CString("EbookSaveOptions_GetLinkTargetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlLinkTargetType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicating the type of target attribute in <c>&lt;a&gt;</c> link. The default value is HtmlLinkTargetType.Parent.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetLinkTargetType(value HtmlLinkTargetType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQG(C.CString("EbookSaveOptions_SetLinkTargetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether the output HTML is compatible with IE browser.
// The defalut value is false
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsIECompatible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsIECompatible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether the output HTML is compatible with IE browser.
// The defalut value is false
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsIECompatible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsIECompatible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicating whether show the whole formatted data of cell when overflowing the column.
// If true then ignore the column width and the whole data of cell will be exported.
// If false then the data will be exported same as Excel.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetFormatDataIgnoreColumnWidth()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetFormatDataIgnoreColumnWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicating whether show the whole formatted data of cell when overflowing the column.
// If true then ignore the column width and the whole data of cell will be exported.
// If false then the data will be exported same as Excel.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetFormatDataIgnoreColumnWidth(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetFormatDataIgnoreColumnWidth"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to calculate formulas before saving html file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetCalculateFormula()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetCalculateFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to calculate formulas before saving html file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCalculateFormula(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetCalculateFormula"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether JavaScript is compatible with browsers that do not support JavaScript.
// The default value is true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsJsBrowserCompatible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsJsBrowserCompatible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether JavaScript is compatible with browsers that do not support JavaScript.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsJsBrowserCompatible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsJsBrowserCompatible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the output HTML is compatible with mobile devices.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsMobileCompatible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsMobileCompatible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the output HTML is compatible with mobile devices.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsMobileCompatible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsMobileCompatible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the additional css styles for the formatter.
// Only works when <see cref="SaveAsSingleFile"/> is True.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetCssStyles()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetCssStyles"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the additional css styles for the formatter.
// Only works when <see cref="SaveAsSingleFile"/> is True.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCssStyles(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetCssStyles"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to hide overflow text when the cell format is set to wrap text.
// The default value is false
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetHideOverflowWrappedText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetHideOverflowWrappedText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to hide overflow text when the cell format is set to wrap text.
// The default value is false
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetHideOverflowWrappedText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetHideOverflowWrappedText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the table borders are collapsed.
// The default value is true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) IsBorderCollapsed()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_IsBorderCollapsed"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the table borders are collapsed.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetIsBorderCollapsed(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetIsBorderCollapsed"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the html character entities are replaced with decimal code.
// (e.g. "&amp;nbsp;" is replaced with "&amp;#160;").
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetEncodeEntityAsCode()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetEncodeEntityAsCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the html character entities are replaced with decimal code.
// (e.g. "&amp;nbsp;" is replaced with "&amp;#160;").
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetEncodeEntityAsCode(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetEncodeEntityAsCode"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how export OfficeMath objects to HTML, Default value is Image.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetOfficeMathOutputMode()  (HtmlOfficeMathOutputType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQH(C.CString("EbookSaveOptions_GetOfficeMathOutputMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlOfficeMathOutputType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates how export OfficeMath objects to HTML, Default value is Image.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetOfficeMathOutputMode(value HtmlOfficeMathOutputType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQI(C.CString("EbookSaveOptions_SetOfficeMathOutputMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the attribute that indicates the CellName to be written.
// (e.g. If the value is "id", then for cell "A1", the output will be:&lt;td id='A1'&gt;).
// The default value is null.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetCellNameAttribute()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetCellNameAttribute"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the attribute that indicates the CellName to be written.
// (e.g. If the value is "id", then for cell "A1", the output will be:&lt;td id='A1'&gt;).
// The default value is null.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCellNameAttribute(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetCellNameAttribute"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether only inline styles are applied, without relying on CSS.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetDisableCss()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetDisableCss"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether only inline styles are applied, without relying on CSS.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetDisableCss(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetDisableCss"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Optimize the output of html by using CSS custom properties. For example, for the scenario that there are multiple occurences for one base64 image, with custom property the image data only needs to be saved once so the performance of the resultant html can be improved.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetEnableCssCustomProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetEnableCssCustomProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Optimize the output of html by using CSS custom properties. For example, for the scenario that there are multiple occurences for one base64 image, with custom property the image data only needs to be saved once so the performance of the resultant html can be improved.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetEnableCssCustomProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetEnableCssCustomProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies version of HTML standard that should be used when saving the HTML format.
// Default value is HtmlVersion.Default.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetHtmlVersion()  (HtmlVersion,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQJ(C.CString("EbookSaveOptions_GetHtmlVersion"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlVersion(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies version of HTML standard that should be used when saving the HTML format.
// Default value is HtmlVersion.Default.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetHtmlVersion(value HtmlVersion)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQK(C.CString("EbookSaveOptions_SetHtmlVersion"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the sheets to render. Default is all visible sheets in the workbook: <see cref="Aspose.Cells.Rendering.SheetSet.Visible"/>.
// Returns:
//   SheetSet  
func (instance *EbookSaveOptions) GetSheetSet()  (*SheetSet,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZLN(C.CString("EbookSaveOptions_GetSheetSet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SheetSet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSheetSet) 

	return result, nil 
}
// Gets or sets the sheets to render. Default is all visible sheets in the workbook: <see cref="Aspose.Cells.Rendering.SheetSet.Visible"/>.
// Parameters:
//   value - SheetSet 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetSheetSet(value *SheetSet)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZLO(C.CString("EbookSaveOptions_SetSheetSet"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the type of font that embedded in html.
// Default value is <see cref="HtmlEmbeddedFontType.None"/> which indicates that it will not embed font in html.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetEmbeddedFontType()  (HtmlEmbeddedFontType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQL(C.CString("EbookSaveOptions_GetEmbeddedFontType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlEmbeddedFontType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the type of font that embedded in html.
// Default value is <see cref="HtmlEmbeddedFontType.None"/> which indicates that it will not embed font in html.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetEmbeddedFontType(value HtmlEmbeddedFontType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZQM(C.CString("EbookSaveOptions_SetEmbeddedFontType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the save file format.
// Returns:
//   int32  
func (instance *EbookSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZLC(C.CString("EbookSaveOptions_GetSaveFormat"), instance.ptr)
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
// Make the workbook empty after saving the file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetClearData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetClearData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Make the workbook empty after saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetClearData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetClearData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The folder for temporary files that may be used as data cache.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetCachedFileFolder()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("EbookSaveOptions_GetCachedFileFolder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The folder for temporary files that may be used as data cache.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCachedFileFolder(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("EbookSaveOptions_SetCachedFileFolder"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether validate merged cells before saving the file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetValidateMergedAreas()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetValidateMergedAreas"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether validate merged cells before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetValidateMergedAreas(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetValidateMergedAreas"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether merge the areas of conditional formatting and validation before saving the file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetMergeAreas()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetMergeAreas"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether merge the areas of conditional formatting and validation before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetMergeAreas(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetMergeAreas"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If true and the directory does not exist, the directory will be automatically created before saving the file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetCreateDirectory()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetCreateDirectory"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// If true and the directory does not exist, the directory will be automatically created before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCreateDirectory(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetCreateDirectory"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether sorting defined names before saving file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetSortNames()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetSortNames"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether sorting defined names before saving file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetSortNames(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetSortNames"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether sorting external defined names before saving file.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetSortExternalNames()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetSortExternalNames"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether sorting external defined names before saving file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetSortExternalNames(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetSortExternalNames"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether refreshing chart cache data
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetRefreshChartCache()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetRefreshChartCache"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether refreshing chart cache data
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetRefreshChartCache(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetRefreshChartCache"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K, it will be truncated.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetCheckExcelRestriction()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetCheckExcelRestriction"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K, it will be truncated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCheckExcelRestriction(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetCheckExcelRestriction"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether updating smart art setting.
// The default value is false.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetUpdateSmartArt()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetUpdateSmartArt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether updating smart art setting.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetUpdateSmartArt(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetUpdateSmartArt"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether encrypt document properties when saving as .xls file.
// The default value is true.
// Returns:
//   bool  
func (instance *EbookSaveOptions) GetEncryptDocumentProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("EbookSaveOptions_GetEncryptDocumentProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether encrypt document properties when saving as .xls file.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetEncryptDocumentProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("EbookSaveOptions_SetEncryptDocumentProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *EbookSaveOptions) ToHtmlSaveOptions() *HtmlSaveOptions {
	parentClass := &HtmlSaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *EbookSaveOptions) ToSaveOptions() *SaveOptions {
	parentClass := &SaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteEbookSaveOptions(ebooksaveoptions *EbookSaveOptions){
	runtime.SetFinalizer(ebooksaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_EbookSaveOptions"),ebooksaveoptions.ptr)
	ebooksaveoptions.ptr = nil
}

// Class SqlScriptColumnTypeMap 

// Represents column type map.
type SqlScriptColumnTypeMap struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewSqlScriptColumnTypeMap() ( *SqlScriptColumnTypeMap, error) {
	sqlscriptcolumntypemap := &SqlScriptColumnTypeMap{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_SqlScriptColumnTypeMap"),)
	if CGoReturnPtr.error_no == 0 {
		sqlscriptcolumntypemap.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(sqlscriptcolumntypemap, DeleteSqlScriptColumnTypeMap)
		return sqlscriptcolumntypemap, nil
	} else {
		sqlscriptcolumntypemap.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return sqlscriptcolumntypemap, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SqlScriptColumnTypeMap) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptColumnTypeMap_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets string type in the database.
// Returns:
//   string  
func (instance *SqlScriptColumnTypeMap) GetStringType()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("SqlScriptColumnTypeMap_GetStringType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets numeric type in the database.
// Returns:
//   string  
func (instance *SqlScriptColumnTypeMap) GetNumbericType()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("SqlScriptColumnTypeMap_GetNumbericType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSqlScriptColumnTypeMap(sqlscriptcolumntypemap *SqlScriptColumnTypeMap){
	runtime.SetFinalizer(sqlscriptcolumntypemap, nil)
	C.Delete_CObject(C.CString("Delete_SqlScriptColumnTypeMap"),sqlscriptcolumntypemap.ptr)
	sqlscriptcolumntypemap.ptr = nil
}

// Class SqlScriptSaveOptions 

// Represents the options of saving sql.
type SqlScriptSaveOptions struct {
	ptr unsafe.Pointer
}

// Creates options for saving sql file.
func NewSqlScriptSaveOptions() ( *SqlScriptSaveOptions, error) {
	sqlscriptsaveoptions := &SqlScriptSaveOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_SqlScriptSaveOptions"),)
	if CGoReturnPtr.error_no == 0 {
		sqlscriptsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(sqlscriptsaveoptions, DeleteSqlScriptSaveOptions)
		return sqlscriptsaveoptions, nil
	} else {
		sqlscriptsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return sqlscriptsaveoptions, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - SaveOptions 
func NewSqlScriptSaveOptions_SaveOptions(src *SaveOptions) ( *SqlScriptSaveOptions, error) {
	sqlscriptsaveoptions := &SqlScriptSaveOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZLB(C.CString("New_SqlScriptSaveOptions_SaveOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		sqlscriptsaveoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(sqlscriptsaveoptions, DeleteSqlScriptSaveOptions)
		return sqlscriptsaveoptions, nil
	} else {
		sqlscriptsaveoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return sqlscriptsaveoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Check if the table name exists before creating
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetCheckIfTableExists()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetCheckIfTableExists"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Check if the table name exists before creating
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCheckIfTableExists(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetCheckIfTableExists"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the map of column type for different database.
// Returns:
//   SqlScriptColumnTypeMap  
func (instance *SqlScriptSaveOptions) GetColumnTypeMap()  (*SqlScriptColumnTypeMap,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZBGO(C.CString("SqlScriptSaveOptions_GetColumnTypeMap"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SqlScriptColumnTypeMap{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSqlScriptColumnTypeMap) 

	return result, nil 
}
// Gets and sets the map of column type for different database.
// Parameters:
//   value - SqlScriptColumnTypeMap 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetColumnTypeMap(value *SqlScriptColumnTypeMap)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZBGP(C.CString("SqlScriptSaveOptions_SetColumnTypeMap"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Check all data to find columns' data type.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetCheckAllDataForColumnType()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetCheckAllDataForColumnType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Check all data to find columns' data type.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCheckAllDataForColumnType(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetCheckAllDataForColumnType"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Insert blank line between each data.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetAddBlankLineBetweenRows()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetAddBlankLineBetweenRows"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert blank line between each data.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetAddBlankLineBetweenRows(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetAddBlankLineBetweenRows"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets character separator of sql script.
// Returns:
//   byte  
func (instance *SqlScriptSaveOptions) GetSeparator()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPI(C.CString("SqlScriptSaveOptions_GetSeparator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets character separator of sql script.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetSeparator(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZSA(C.CString("SqlScriptSaveOptions_SetSeparator"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the operator type of sql.
// Returns:
//   int32  
func (instance *SqlScriptSaveOptions) GetOperatorType()  (SqlScriptOperatorType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZBGQ(C.CString("SqlScriptSaveOptions_GetOperatorType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSqlScriptOperatorType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the operator type of sql.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetOperatorType(value SqlScriptOperatorType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZBGR(C.CString("SqlScriptSaveOptions_SetOperatorType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents which column is primary key of the data table.
// Returns:
//   int32  
func (instance *SqlScriptSaveOptions) GetPrimaryKey()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SqlScriptSaveOptions_GetPrimaryKey"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents which column is primary key of the data table.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetPrimaryKey(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SqlScriptSaveOptions_SetPrimaryKey"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting sql of creating table.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetCreateTable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetCreateTable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting sql of creating table.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCreateTable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetCreateTable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the name of id column.
// Returns:
//   string  
func (instance *SqlScriptSaveOptions) GetIdName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("SqlScriptSaveOptions_GetIdName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of id column.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetIdName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("SqlScriptSaveOptions_SetIdName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the start id.
// Returns:
//   int32  
func (instance *SqlScriptSaveOptions) GetStartId()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SqlScriptSaveOptions_GetStartId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the start id.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetStartId(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SqlScriptSaveOptions_SetStartId"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the table name.
// Returns:
//   string  
func (instance *SqlScriptSaveOptions) GetTableName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("SqlScriptSaveOptions_GetTableName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the table name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetTableName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("SqlScriptSaveOptions_SetTableName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether exporting all data as string value.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetExportAsString()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetExportAsString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether exporting all data as string value.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetExportAsString(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetExportAsString"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the indexes of exported sheets.
// Returns:
//   []int32_t  
func (instance *SqlScriptSaveOptions) GetSheetIndexes()  ([]int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAS(C.CString("SqlScriptSaveOptions_GetSheetIndexes"), instance.ptr)
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
// Represents the indexes of exported sheets.
// Parameters:
//   value - []int32_t 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetSheetIndexes(value []int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZRD(C.CString("SqlScriptSaveOptions_SetSheetIndexes"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the exporting range.
// Returns:
//   CellArea  
func (instance *SqlScriptSaveOptions) GetExportArea()  (*CellArea,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SqlScriptSaveOptions_GetExportArea"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets or sets the exporting range.
// Parameters:
//   value - CellArea 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetExportArea(value *CellArea)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZGH(C.CString("SqlScriptSaveOptions_SetExportArea"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the range contains header row.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetHasHeaderRow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetHasHeaderRow"), instance.ptr)
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
func (instance *SqlScriptSaveOptions) SetHasHeaderRow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetHasHeaderRow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the save file format.
// Returns:
//   int32  
func (instance *SqlScriptSaveOptions) GetSaveFormat()  (SaveFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZLC(C.CString("SqlScriptSaveOptions_GetSaveFormat"), instance.ptr)
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
// Make the workbook empty after saving the file.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetClearData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetClearData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Make the workbook empty after saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetClearData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetClearData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The folder for temporary files that may be used as data cache.
// Returns:
//   string  
func (instance *SqlScriptSaveOptions) GetCachedFileFolder()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("SqlScriptSaveOptions_GetCachedFileFolder"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The folder for temporary files that may be used as data cache.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCachedFileFolder(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("SqlScriptSaveOptions_SetCachedFileFolder"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether validate merged cells before saving the file.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetValidateMergedAreas()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetValidateMergedAreas"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether validate merged cells before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetValidateMergedAreas(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetValidateMergedAreas"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether merge the areas of conditional formatting and validation before saving the file.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetMergeAreas()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetMergeAreas"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether merge the areas of conditional formatting and validation before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetMergeAreas(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetMergeAreas"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If true and the directory does not exist, the directory will be automatically created before saving the file.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetCreateDirectory()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetCreateDirectory"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// If true and the directory does not exist, the directory will be automatically created before saving the file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCreateDirectory(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetCreateDirectory"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether sorting defined names before saving file.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetSortNames()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetSortNames"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether sorting defined names before saving file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetSortNames(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetSortNames"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether sorting external defined names before saving file.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetSortExternalNames()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetSortExternalNames"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether sorting external defined names before saving file.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetSortExternalNames(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetSortExternalNames"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether refreshing chart cache data
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetRefreshChartCache()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetRefreshChartCache"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether refreshing chart cache data
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetRefreshChartCache(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetRefreshChartCache"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K, it will be truncated.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetCheckExcelRestriction()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetCheckExcelRestriction"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether check restriction of excel file when user modify cells related objects.
// For example, excel does not allow inputting string value longer than 32K.
// When you input a value longer than 32K, it will be truncated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCheckExcelRestriction(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetCheckExcelRestriction"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether updating smart art setting.
// The default value is false.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetUpdateSmartArt()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetUpdateSmartArt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether updating smart art setting.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetUpdateSmartArt(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetUpdateSmartArt"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether encrypt document properties when saving as .xls file.
// The default value is true.
// Returns:
//   bool  
func (instance *SqlScriptSaveOptions) GetEncryptDocumentProperties()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SqlScriptSaveOptions_GetEncryptDocumentProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether encrypt document properties when saving as .xls file.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetEncryptDocumentProperties(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SqlScriptSaveOptions_SetEncryptDocumentProperties"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *SqlScriptSaveOptions) ToSaveOptions() *SaveOptions {
	parentClass := &SaveOptions{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteSqlScriptSaveOptions(sqlscriptsaveoptions *SqlScriptSaveOptions){
	runtime.SetFinalizer(sqlscriptsaveoptions, nil)
	C.Delete_CObject(C.CString("Delete_SqlScriptSaveOptions"),sqlscriptsaveoptions.ptr)
	sqlscriptsaveoptions.ptr = nil
}
