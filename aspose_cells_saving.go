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
	"fmt"  
 	
	"errors"	
	"runtime"
	"unsafe" 
)

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
	CGoReturnPtr := C.New_DbfSaveOptions()
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
	CGoReturnPtr := C.New_DbfSaveOptions_SaveOptions(src.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_IsNull( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetExportAsString( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetExportAsString( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetSaveFormat( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetClearData( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetClearData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The cached file folder is used to store some large data.
// Returns:
//   string  
func (instance *DbfSaveOptions) GetCachedFileFolder()  (string,  error)  {
	
	CGoReturnPtr := C.DbfSaveOptions_GetCachedFileFolder( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The cached file folder is used to store some large data.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DbfSaveOptions) SetCachedFileFolder(value string)  error {
	
	CGoReturnPtr := C.DbfSaveOptions_SetCachedFileFolder( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetValidateMergedAreas( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetValidateMergedAreas( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetMergeAreas( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetMergeAreas( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetCreateDirectory( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetCreateDirectory( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetSortNames( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetSortNames( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetSortExternalNames( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetSortExternalNames( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetRefreshChartCache( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetRefreshChartCache( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetCheckExcelRestriction( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetCheckExcelRestriction( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetUpdateSmartArt( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetUpdateSmartArt( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.DbfSaveOptions_GetEncryptDocumentProperties( instance.ptr)
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
	
	CGoReturnPtr := C.DbfSaveOptions_SetEncryptDocumentProperties( instance.ptr, C.bool(value))
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
	C.Delete_DbfSaveOptions(dbfsaveoptions.ptr)
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
	CGoReturnPtr := C.New_EbookSaveOptions()
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
	CGoReturnPtr := C.New_EbookSaveOptions_SaveFormat(C.int( int32(saveformat)))
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
	CGoReturnPtr := C.New_EbookSaveOptions_HtmlSaveOptions(src.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsNull( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetIgnoreInvisibleShapes( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIgnoreInvisibleShapes( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetPageTitle( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetPageTitle( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetAttachedFilesDirectory( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetAttachedFilesDirectory( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetAttachedFilesUrlPrefix( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetAttachedFilesUrlPrefix( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetDefaultFontName( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetDefaultFontName( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetAddGenericFont( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetAddGenericFont( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetWorksheetScalable( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetWorksheetScalable( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsExportComments( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsExportComments( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportCommentsType( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportCommentsType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetDisableDownlevelRevealedComments( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetDisableDownlevelRevealedComments( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsExpImageToTempDir( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsExpImageToTempDir( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetImageScalable( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetImageScalable( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetWidthScalable( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetWidthScalable( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportSingleTab( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportSingleTab( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportImagesAsBase64( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportImagesAsBase64( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportActiveWorksheetOnly( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportActiveWorksheetOnly( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportPrintAreaOnly( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportPrintAreaOnly( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportArea( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportArea( instance.ptr, value.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetParseHtmlTagInCell( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetParseHtmlTagInCell( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetHtmlCrossStringType( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetHtmlCrossStringType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetHiddenColDisplayType( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetHiddenColDisplayType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetHiddenRowDisplayType( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetHiddenRowDisplayType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetEncoding( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetEncoding( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetImageOptions( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetSaveAsSingleFile( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetSaveAsSingleFile( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetShowAllSheets( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetShowAllSheets( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportPageHeaders( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportPageHeaders( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportPageFooters( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportPageFooters( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportHiddenWorksheet( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportHiddenWorksheet( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetPresentationPreference( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetPresentationPreference( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetCellCssPrefix( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetCellCssPrefix( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetTableCssId( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetTableCssId( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsFullPathLink( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsFullPathLink( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportWorksheetCSSSeparately( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportWorksheetCSSSeparately( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportSimilarBorderStyle( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportSimilarBorderStyle( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetMergeEmptyTdType( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetMergeEmptyTdType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportCellCoordinate( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportCellCoordinate( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportExtraHeadings( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportExtraHeadings( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportRowColumnHeadings( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportRowColumnHeadings( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportFormula( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportFormula( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetAddTooltipText( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetAddTooltipText( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportGridLines( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportGridLines( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportBogusRowData( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportBogusRowData( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExcludeUnusedStyles( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExcludeUnusedStyles( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportDocumentProperties( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportDocumentProperties( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportWorksheetProperties( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportWorksheetProperties( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportWorkbookProperties( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportWorkbookProperties( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportFrameScriptsAndProperties( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportFrameScriptsAndProperties( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetExportDataOptions( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetExportDataOptions( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetLinkTargetType( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetLinkTargetType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsIECompatible( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsIECompatible( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetFormatDataIgnoreColumnWidth( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetFormatDataIgnoreColumnWidth( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetCalculateFormula( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetCalculateFormula( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsJsBrowserCompatible( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsJsBrowserCompatible( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsMobileCompatible( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsMobileCompatible( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetCssStyles( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetCssStyles( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetHideOverflowWrappedText( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetHideOverflowWrappedText( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_IsBorderCollapsed( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetIsBorderCollapsed( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetEncodeEntityAsCode( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetEncodeEntityAsCode( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetOfficeMathOutputMode( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetOfficeMathOutputMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetCellNameAttribute( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetCellNameAttribute( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetDisableCss( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetDisableCss( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetEnableCssCustomProperties( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetEnableCssCustomProperties( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetHtmlVersion( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetHtmlVersion( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetSaveFormat( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetClearData( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetClearData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The cached file folder is used to store some large data.
// Returns:
//   string  
func (instance *EbookSaveOptions) GetCachedFileFolder()  (string,  error)  {
	
	CGoReturnPtr := C.EbookSaveOptions_GetCachedFileFolder( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The cached file folder is used to store some large data.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *EbookSaveOptions) SetCachedFileFolder(value string)  error {
	
	CGoReturnPtr := C.EbookSaveOptions_SetCachedFileFolder( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetValidateMergedAreas( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetValidateMergedAreas( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetMergeAreas( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetMergeAreas( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetCreateDirectory( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetCreateDirectory( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetSortNames( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetSortNames( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetSortExternalNames( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetSortExternalNames( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetRefreshChartCache( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetRefreshChartCache( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetCheckExcelRestriction( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetCheckExcelRestriction( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetUpdateSmartArt( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetUpdateSmartArt( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.EbookSaveOptions_GetEncryptDocumentProperties( instance.ptr)
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
	
	CGoReturnPtr := C.EbookSaveOptions_SetEncryptDocumentProperties( instance.ptr, C.bool(value))
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
	C.Delete_EbookSaveOptions(ebooksaveoptions.ptr)
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
	CGoReturnPtr := C.New_SqlScriptColumnTypeMap()
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
	
	CGoReturnPtr := C.SqlScriptColumnTypeMap_IsNull( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptColumnTypeMap_GetStringType( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptColumnTypeMap_GetNumbericType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSqlScriptColumnTypeMap(sqlscriptcolumntypemap *SqlScriptColumnTypeMap){
	runtime.SetFinalizer(sqlscriptcolumntypemap, nil)
	C.Delete_SqlScriptColumnTypeMap(sqlscriptcolumntypemap.ptr)
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
	CGoReturnPtr := C.New_SqlScriptSaveOptions()
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
	CGoReturnPtr := C.New_SqlScriptSaveOptions_SaveOptions(src.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_IsNull( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetCheckIfTableExists( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetCheckIfTableExists( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetColumnTypeMap( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetColumnTypeMap( instance.ptr, value.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetCheckAllDataForColumnType( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetCheckAllDataForColumnType( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetAddBlankLineBetweenRows( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetAddBlankLineBetweenRows( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetSeparator( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetSeparator( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetOperatorType( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetOperatorType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetPrimaryKey( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetPrimaryKey( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetCreateTable( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetCreateTable( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetIdName( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetIdName( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetStartId( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetStartId( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetTableName( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetTableName( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetExportAsString( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetExportAsString( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetSheetIndexes( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetSheetIndexes( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetExportArea( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetExportArea( instance.ptr, value.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetHasHeaderRow( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetHasHeaderRow( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetSaveFormat( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetClearData( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetClearData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The cached file folder is used to store some large data.
// Returns:
//   string  
func (instance *SqlScriptSaveOptions) GetCachedFileFolder()  (string,  error)  {
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetCachedFileFolder( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The cached file folder is used to store some large data.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SqlScriptSaveOptions) SetCachedFileFolder(value string)  error {
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetCachedFileFolder( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetValidateMergedAreas( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetValidateMergedAreas( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetMergeAreas( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetMergeAreas( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetCreateDirectory( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetCreateDirectory( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetSortNames( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetSortNames( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetSortExternalNames( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetSortExternalNames( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetRefreshChartCache( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetRefreshChartCache( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetCheckExcelRestriction( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetCheckExcelRestriction( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetUpdateSmartArt( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetUpdateSmartArt( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_GetEncryptDocumentProperties( instance.ptr)
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
	
	CGoReturnPtr := C.SqlScriptSaveOptions_SetEncryptDocumentProperties( instance.ptr, C.bool(value))
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
	C.Delete_SqlScriptSaveOptions(sqlscriptsaveoptions.ptr)
	sqlscriptsaveoptions.ptr = nil
}
