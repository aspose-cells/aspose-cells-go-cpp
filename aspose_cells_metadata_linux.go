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

/**************Enum MetadataType *****************/

// Represents the type of metadata.
type MetadataType int32

const(
// Encrypts the file.
MetadataType_Encryption MetadataType = 1 

// Decrypts the file.
MetadataType_Decryption MetadataType = 2 

// Load the properties of the file.
MetadataType_Document_Properties MetadataType = 4 
)

func Int32ToMetadataType(value int32)(MetadataType ,error){
	switch value {
		case 1:  return MetadataType_Encryption, nil  
		case 2:  return MetadataType_Decryption, nil  
		case 4:  return MetadataType_Document_Properties, nil  
		default:
			return 0 ,fmt.Errorf("invalid MetadataType value: %d", value)
	}
}
// Class MetadataOptions 

// Represents the options of loading metadata of the file.
type MetadataOptions struct {
	ptr unsafe.Pointer
}

// Creates an options of loading the metadata.
// Parameters:
//   metadataType - int32 
func NewMetadataOptions(metadatatype MetadataType) ( *MetadataOptions, error) {
	metadataoptions := &MetadataOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZFK(C.CString("New_MetadataOptions"),C.int( int32(metadatatype)))
	if CGoReturnPtr.error_no == 0 {
		metadataoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(metadataoptions, DeleteMetadataOptions)
		return metadataoptions, nil
	} else {
		metadataoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return metadataoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *MetadataOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("MetadataOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the type of the metadata which is loading.
// Returns:
//   int32  
func (instance *MetadataOptions) GetMetadataType()  (MetadataType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("MetadataOptions_GetMetadataType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToMetadataType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents Workbook file encryption password.
// Returns:
//   string  
func (instance *MetadataOptions) GetPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("MetadataOptions_GetPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents Workbook file encryption password.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *MetadataOptions) SetPassword(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("MetadataOptions_SetPassword"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The key length.
// Returns:
//   int32  
func (instance *MetadataOptions) GetKeyLength()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("MetadataOptions_GetKeyLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// The key length.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *MetadataOptions) SetKeyLength(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("MetadataOptions_SetKeyLength"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteMetadataOptions(metadataoptions *MetadataOptions){
	runtime.SetFinalizer(metadataoptions, nil)
	C.Delete_CObject(C.CString("Delete_MetadataOptions"),metadataoptions.ptr)
	metadataoptions.ptr = nil
}

// Class WorkbookMetadata 

// Represents the meta data.
type WorkbookMetadata struct {
	ptr unsafe.Pointer
}

// Create the meta data object.
// Parameters:
//   fileName - string 
//   options - MetadataOptions 
func NewWorkbookMetadata_String_MetadataOptions(filename string, options *MetadataOptions) ( *WorkbookMetadata, error) {
	workbookmetadata := &WorkbookMetadata{}
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZJC(C.CString("New_WorkbookMetadata_String_MetadataOptions"),C.CString(filename), options_ptr)
	if CGoReturnPtr.error_no == 0 {
		workbookmetadata.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(workbookmetadata, DeleteWorkbookMetadata)
		return workbookmetadata, nil
	} else {
		workbookmetadata.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return workbookmetadata, err
	}	
}
// Create the meta data object.
// Parameters:
//   stream - []byte 
//   options - MetadataOptions 
func NewWorkbookMetadata_Stream_MetadataOptions(stream []byte, options *MetadataOptions) ( *WorkbookMetadata, error) {
	workbookmetadata := &WorkbookMetadata{}
	var options_ptr unsafe.Pointer = nil
	if options != nil {
	  options_ptr =options.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZJD(C.CString("New_WorkbookMetadata_Stream_MetadataOptions"),unsafe.Pointer(&stream[0]), C.int( len(stream)), options_ptr)
	if CGoReturnPtr.error_no == 0 {
		workbookmetadata.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(workbookmetadata, DeleteWorkbookMetadata)
		return workbookmetadata, nil
	} else {
		workbookmetadata.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return workbookmetadata, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WorkbookMetadata) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WorkbookMetadata_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the options of the metadata.
// Returns:
//   MetadataOptions  
func (instance *WorkbookMetadata) GetOptions()  (*MetadataOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WorkbookMetadata_GetOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &MetadataOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteMetadataOptions) 

	return result, nil 
}
// Returns a <see cref="DocumentProperty"/> collection that represents all the  built-in document properties of the spreadsheet.
// Returns:
//   BuiltInDocumentPropertyCollection  
func (instance *WorkbookMetadata) GetBuiltInDocumentProperties()  (*BuiltInDocumentPropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WorkbookMetadata_GetBuiltInDocumentProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &BuiltInDocumentPropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteBuiltInDocumentPropertyCollection) 

	return result, nil 
}
// Returns a <see cref="DocumentProperty"/> collection that represents all the custom document properties of the spreadsheet.
// Returns:
//   CustomDocumentPropertyCollection  
func (instance *WorkbookMetadata) GetCustomDocumentProperties()  (*CustomDocumentPropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WorkbookMetadata_GetCustomDocumentProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CustomDocumentPropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCustomDocumentPropertyCollection) 

	return result, nil 
}
// Save the modified metadata to the file.
// Parameters:
//   fileName - string 
// Returns:
//   void  
func (instance *WorkbookMetadata) Save_String(filename string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WorkbookMetadata_Save_String"), instance.ptr, C.CString(filename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Save the modified metadata to the stream.
// Parameters:
//   stream - []byte 
// Returns:
//   void  
func (instance *WorkbookMetadata) Save_Stream(stream []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("WorkbookMetadata_Save_Stream"), instance.ptr, unsafe.Pointer(&stream[0]), C.int( len(stream)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteWorkbookMetadata(workbookmetadata *WorkbookMetadata){
	runtime.SetFinalizer(workbookmetadata, nil)
	C.Delete_CObject(C.CString("Delete_WorkbookMetadata"),workbookmetadata.ptr)
	workbookmetadata.ptr = nil
}
