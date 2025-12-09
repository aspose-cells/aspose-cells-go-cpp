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

/**************Enum SmartTagShowType *****************/

// Represents the show type of the smart tag.
type SmartTagShowType int32

const(
// Indicates that smart tags are enabled and shown
SmartTagShowType_All SmartTagShowType = 0 

// Indicates that the smart tags are enabled but the indicator not be shown.
SmartTagShowType_NoSmartTagIndicator SmartTagShowType = 1 

// Indicates that smart tags are disabled and not displayed.
SmartTagShowType_None SmartTagShowType = 2 
)

func Int32ToSmartTagShowType(value int32)(SmartTagShowType ,error){
	switch value {
		case 0:  return SmartTagShowType_All, nil  
		case 1:  return SmartTagShowType_NoSmartTagIndicator, nil  
		case 2:  return SmartTagShowType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid SmartTagShowType value: %d", value)
	}
}
// Class CustomXmlPart 

// Represents a Custom XML Data Storage Part (custom XML data within a package).
type CustomXmlPart struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CustomXmlPart) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CustomXmlPart_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the XML content of this Custom XML Data Storage Part.
// Returns:
//   []byte  
func (instance *CustomXmlPart) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CustomXmlPart_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets or sets the XML content of this Custom XML Data Storage Part.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *CustomXmlPart) SetData(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("CustomXmlPart_SetData"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the XML content of this Custom XML Schema Data Storage Part.
// Returns:
//   []byte  
func (instance *CustomXmlPart) GetSchemaData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CustomXmlPart_GetSchemaData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets or sets the XML content of this Custom XML Schema Data Storage Part.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *CustomXmlPart) SetSchemaData(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("CustomXmlPart_SetSchemaData"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the id of the custom xml part.
// Returns:
//   string  
func (instance *CustomXmlPart) GetID()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CustomXmlPart_GetID"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the id of the custom xml part.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CustomXmlPart) SetID(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("CustomXmlPart_SetID"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteCustomXmlPart(customxmlpart *CustomXmlPart){
	runtime.SetFinalizer(customxmlpart, nil)
	C.Delete_CObject(C.CString("Delete_CustomXmlPart"),customxmlpart.ptr)
	customxmlpart.ptr = nil
}

// Class CustomXmlPartCollection 

// Represents a Custom XML Data Storage Part (custom XML data within a package).
type CustomXmlPartCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CustomXmlPartCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CustomXmlPartCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds an item to the collection.
// Parameters:
//   data - []byte 
//   shemaData - []byte 
// Returns:
//   int32  
func (instance *CustomXmlPartCollection) Add(data []byte, shemadata []byte)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZOV(C.CString("CustomXmlPartCollection_Add"), instance.ptr, unsafe.Pointer(&data[0]), C.int( len(data)), unsafe.Pointer(&shemadata[0]), C.int( len(shemadata)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets an item at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   CustomXmlPart  
func (instance *CustomXmlPartCollection) Get(index int32)  (*CustomXmlPart,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("CustomXmlPartCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CustomXmlPart{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCustomXmlPart) 

	return result, nil 
}
// Gets an item by id.
// Parameters:
//   id - string 
// Returns:
//   CustomXmlPart  
func (instance *CustomXmlPartCollection) SelectByID(id string)  (*CustomXmlPart,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("CustomXmlPartCollection_SelectByID"), instance.ptr, C.CString(id))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CustomXmlPart{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCustomXmlPart) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *CustomXmlPartCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("CustomXmlPartCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteCustomXmlPartCollection(customxmlpartcollection *CustomXmlPartCollection){
	runtime.SetFinalizer(customxmlpartcollection, nil)
	C.Delete_CObject(C.CString("Delete_CustomXmlPartCollection"),customxmlpartcollection.ptr)
	customxmlpartcollection.ptr = nil
}

// Class SmartTag 

// Represents a smart tag.
type SmartTag struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SmartTag) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTag_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the smart tag is deleted.
// Returns:
//   bool  
func (instance *SmartTag) GetDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTag_GetDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the smart tag is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SmartTag) SetDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SmartTag_SetDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the properties of the smart tag.
// Returns:
//   SmartTagPropertyCollection  
func (instance *SmartTag) GetProperties()  (*SmartTagPropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SmartTag_GetProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTagPropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTagPropertyCollection) 

	return result, nil 
}
// Gets and set the properties of the smart tag.
// Parameters:
//   value - SmartTagPropertyCollection 
// Returns:
//   void  
func (instance *SmartTag) SetProperties(value *SmartTagPropertyCollection)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SmartTag_SetProperties"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the namespace URI of the smart tag.
// Returns:
//   string  
func (instance *SmartTag) GetUri()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SmartTag_GetUri"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the smart tag.
// Returns:
//   string  
func (instance *SmartTag) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SmartTag_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Change the name and  the namespace URI of the smart tag.
// Parameters:
//   uri - string 
//   name - string 
// Returns:
//   void  
func (instance *SmartTag) SetLink(uri string, name string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZFN(C.CString("SmartTag_SetLink"), instance.ptr, C.CString(uri), C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSmartTag(smarttag *SmartTag){
	runtime.SetFinalizer(smarttag, nil)
	C.Delete_CObject(C.CString("Delete_SmartTag"),smarttag.ptr)
	smarttag.ptr = nil
}

// Class SmartTagCollection 

// Represents all smart tags in the cell.
type SmartTagCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SmartTagCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTagCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the row of the cell smart tags.
// Returns:
//   int32  
func (instance *SmartTagCollection) GetRow()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SmartTagCollection_GetRow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the column of the cell smart tags.
// Returns:
//   int32  
func (instance *SmartTagCollection) GetColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SmartTagCollection_GetColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets a <see cref="SmartTag"/> object at the specific index
// Parameters:
//   index - int32 
// Returns:
//   SmartTag  
func (instance *SmartTagCollection) Get(index int32)  (*SmartTag,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SmartTagCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTag{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTag) 

	return result, nil 
}
// Adds a smart tag.
// Parameters:
//   uri - string 
//   name - string 
// Returns:
//   int32  
func (instance *SmartTagCollection) Add(uri string, name string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJG(C.CString("SmartTagCollection_Add"), instance.ptr, C.CString(uri), C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *SmartTagCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SmartTagCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSmartTagCollection(smarttagcollection *SmartTagCollection){
	runtime.SetFinalizer(smarttagcollection, nil)
	C.Delete_CObject(C.CString("Delete_SmartTagCollection"),smarttagcollection.ptr)
	smarttagcollection.ptr = nil
}

// Class SmartTagOptions 

// Represents the options of the smart tag.
type SmartTagOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewSmartTagOptions() ( *SmartTagOptions, error) {
	smarttagoptions := &SmartTagOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_SmartTagOptions"),)
	if CGoReturnPtr.error_no == 0 {
		smarttagoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(smarttagoptions, DeleteSmartTagOptions)
		return smarttagoptions, nil
	} else {
		smarttagoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return smarttagoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SmartTagOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTagOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether saving smart tags with the workbook.
// Returns:
//   bool  
func (instance *SmartTagOptions) GetEmbedSmartTags()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTagOptions_GetEmbedSmartTags"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether saving smart tags with the workbook.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SmartTagOptions) SetEmbedSmartTags(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SmartTagOptions_SetEmbedSmartTags"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the show type of smart tag.
// Returns:
//   int32  
func (instance *SmartTagOptions) GetShowType()  (SmartTagShowType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SmartTagOptions_GetShowType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSmartTagShowType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the show type of smart tag.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SmartTagOptions) SetShowType(value SmartTagShowType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SmartTagOptions_SetShowType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSmartTagOptions(smarttagoptions *SmartTagOptions){
	runtime.SetFinalizer(smarttagoptions, nil)
	C.Delete_CObject(C.CString("Delete_SmartTagOptions"),smarttagoptions.ptr)
	smarttagoptions.ptr = nil
}

// Class SmartTagProperty 

// Represents the property of the cell smart tag.
type SmartTagProperty struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SmartTagProperty) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTagProperty_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the property.
// Returns:
//   string  
func (instance *SmartTagProperty) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SmartTagProperty_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SmartTagProperty) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("SmartTagProperty_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the value of the property.
// Returns:
//   string  
func (instance *SmartTagProperty) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SmartTagProperty_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the value of the property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SmartTagProperty) SetValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("SmartTagProperty_SetValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSmartTagProperty(smarttagproperty *SmartTagProperty){
	runtime.SetFinalizer(smarttagproperty, nil)
	C.Delete_CObject(C.CString("Delete_SmartTagProperty"),smarttagproperty.ptr)
	smarttagproperty.ptr = nil
}

// Class SmartTagPropertyCollection 

// Represents all properties of cell smart tag.
type SmartTagPropertyCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewSmartTagPropertyCollection() ( *SmartTagPropertyCollection, error) {
	smarttagpropertycollection := &SmartTagPropertyCollection{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_SmartTagPropertyCollection"),)
	if CGoReturnPtr.error_no == 0 {
		smarttagpropertycollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(smarttagpropertycollection, DeleteSmartTagPropertyCollection)
		return smarttagpropertycollection, nil
	} else {
		smarttagpropertycollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return smarttagpropertycollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SmartTagPropertyCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTagPropertyCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets a <see cref="SmartTagProperty"/> object.
// Parameters:
//   index - int32 
// Returns:
//   SmartTagProperty  
func (instance *SmartTagPropertyCollection) Get_Int(index int32)  (*SmartTagProperty,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SmartTagPropertyCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTagProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTagProperty) 

	return result, nil 
}
// Gets a <see cref="SmartTagProperty"/> object by the name of the property.
// Parameters:
//   name - string 
// Returns:
//   SmartTagProperty  
func (instance *SmartTagPropertyCollection) Get_String(name string)  (*SmartTagProperty,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("SmartTagPropertyCollection_Get_String"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTagProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTagProperty) 

	return result, nil 
}
// Adds a property of cell's smart tag.
// Parameters:
//   name - string 
//   value - string 
// Returns:
//   int32  
func (instance *SmartTagPropertyCollection) Add(name string, value string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJG(C.CString("SmartTagPropertyCollection_Add"), instance.ptr, C.CString(name), C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *SmartTagPropertyCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SmartTagPropertyCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSmartTagPropertyCollection(smarttagpropertycollection *SmartTagPropertyCollection){
	runtime.SetFinalizer(smarttagpropertycollection, nil)
	C.Delete_CObject(C.CString("Delete_SmartTagPropertyCollection"),smarttagpropertycollection.ptr)
	smarttagpropertycollection.ptr = nil
}

// Class SmartTagSetting 

// Represents all <see cref="SmartTagCollection"/> object in the worksheet.
type SmartTagSetting struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SmartTagSetting) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SmartTagSetting_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a <see cref="SmartTagCollection"/> object to a cell.
// Parameters:
//   row - int32 
//   column - int32 
// Returns:
//   int32  
func (instance *SmartTagSetting) Add_Int_Int(row int32, column int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZEL(C.CString("SmartTagSetting_Add_Integer_Integer"), instance.ptr, C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a cell smart tags.
// Parameters:
//   cellName - string 
// Returns:
//   int32  
func (instance *SmartTagSetting) Add_String(cellname string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZEM(C.CString("SmartTagSetting_Add_String"), instance.ptr, C.CString(cellname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets a <see cref="SmartTagCollection"/> object by the index.
// Parameters:
//   index - int32 
// Returns:
//   SmartTagCollection  
func (instance *SmartTagSetting) Get_Int(index int32)  (*SmartTagCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("SmartTagSetting_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTagCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTagCollection) 

	return result, nil 
}
// Gets the <see cref="SmartTagCollection"/> object of the cell.
// Parameters:
//   row - int32 
//   column - int32 
// Returns:
//   SmartTagCollection  
func (instance *SmartTagSetting) Get_Int_Int(row int32, column int32)  (*SmartTagCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBG(C.CString("SmartTagSetting_Get_Integer_Integer"), instance.ptr, C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTagCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTagCollection) 

	return result, nil 
}
// Gets the <see cref="SmartTagCollection"/> object of the cell.
// Parameters:
//   cellName - string 
// Returns:
//   SmartTagCollection  
func (instance *SmartTagSetting) Get_String(cellname string)  (*SmartTagCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("SmartTagSetting_Get_String"), instance.ptr, C.CString(cellname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SmartTagCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSmartTagCollection) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *SmartTagSetting) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SmartTagSetting_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSmartTagSetting(smarttagsetting *SmartTagSetting){
	runtime.SetFinalizer(smarttagsetting, nil)
	C.Delete_CObject(C.CString("Delete_SmartTagSetting"),smarttagsetting.ptr)
	smarttagsetting.ptr = nil
}
