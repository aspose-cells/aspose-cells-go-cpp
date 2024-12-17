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
	CGoReturnPtr := C.CustomXmlPart_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the id of the custom xml part.
// Returns:
//   string  
func (instance *CustomXmlPart) GetID()  (string,  error)  {
	CGoReturnPtr := C.CustomXmlPart_GetID( instance.ptr)
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
	CGoReturnPtr := C.CustomXmlPart_SetID( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteCustomXmlPart(customxmlpart *CustomXmlPart){
	runtime.SetFinalizer(customxmlpart, nil)
	C.Delete_CustomXmlPart(customxmlpart.ptr)
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
	CGoReturnPtr := C.CustomXmlPartCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets an item at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   CustomXmlPart  
func (instance *CustomXmlPartCollection) Get(index int32)  (*CustomXmlPart,  error)  {
	CGoReturnPtr := C.CustomXmlPartCollection_Get( instance.ptr, C.int(index))
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
	CGoReturnPtr := C.CustomXmlPartCollection_SelectByID( instance.ptr, C.CString(id))
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
	CGoReturnPtr := C.CustomXmlPartCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteCustomXmlPartCollection(customxmlpartcollection *CustomXmlPartCollection){
	runtime.SetFinalizer(customxmlpartcollection, nil)
	C.Delete_CustomXmlPartCollection(customxmlpartcollection.ptr)
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
	CGoReturnPtr := C.SmartTag_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the smart tag is deleted.
// Returns:
//   bool  
func (instance *SmartTag) GetDeleted()  (bool,  error)  {
	CGoReturnPtr := C.SmartTag_GetDeleted( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the smart tag is deleted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SmartTag) SetDeleted(value bool)  error {
	CGoReturnPtr := C.SmartTag_SetDeleted( instance.ptr, C.bool(value))
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
	CGoReturnPtr := C.SmartTag_GetProperties( instance.ptr)
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
	CGoReturnPtr := C.SmartTag_SetProperties( instance.ptr, value.ptr)
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
	CGoReturnPtr := C.SmartTag_GetUri( instance.ptr)
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
	CGoReturnPtr := C.SmartTag_GetName( instance.ptr)
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
	CGoReturnPtr := C.SmartTag_SetLink( instance.ptr, C.CString(uri), C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteSmartTag(smarttag *SmartTag){
	runtime.SetFinalizer(smarttag, nil)
	C.Delete_SmartTag(smarttag.ptr)
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
	CGoReturnPtr := C.SmartTagCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the row of the cell smart tags.
// Returns:
//   int32  
func (instance *SmartTagCollection) GetRow()  (int32,  error)  {
	CGoReturnPtr := C.SmartTagCollection_GetRow( instance.ptr)
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
	CGoReturnPtr := C.SmartTagCollection_GetColumn( instance.ptr)
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
	CGoReturnPtr := C.SmartTagCollection_Get( instance.ptr, C.int(index))
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
	CGoReturnPtr := C.SmartTagCollection_Add( instance.ptr, C.CString(uri), C.CString(name))
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
	CGoReturnPtr := C.SmartTagCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteSmartTagCollection(smarttagcollection *SmartTagCollection){
	runtime.SetFinalizer(smarttagcollection, nil)
	C.Delete_SmartTagCollection(smarttagcollection.ptr)
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
	CGoReturnPtr := C.New_SmartTagOptions()
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
	CGoReturnPtr := C.SmartTagOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether saving smart tags with the workbook.
// Returns:
//   bool  
func (instance *SmartTagOptions) GetEmbedSmartTags()  (bool,  error)  {
	CGoReturnPtr := C.SmartTagOptions_GetEmbedSmartTags( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether saving smart tags with the workbook.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SmartTagOptions) SetEmbedSmartTags(value bool)  error {
	CGoReturnPtr := C.SmartTagOptions_SetEmbedSmartTags( instance.ptr, C.bool(value))
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
	CGoReturnPtr := C.SmartTagOptions_GetShowType( instance.ptr)
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
	CGoReturnPtr := C.SmartTagOptions_SetShowType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteSmartTagOptions(smarttagoptions *SmartTagOptions){
	runtime.SetFinalizer(smarttagoptions, nil)
	C.Delete_SmartTagOptions(smarttagoptions.ptr)
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
	CGoReturnPtr := C.SmartTagProperty_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the name of the property.
// Returns:
//   string  
func (instance *SmartTagProperty) GetName()  (string,  error)  {
	CGoReturnPtr := C.SmartTagProperty_GetName( instance.ptr)
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
	CGoReturnPtr := C.SmartTagProperty_SetName( instance.ptr, C.CString(value))
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
	CGoReturnPtr := C.SmartTagProperty_GetValue( instance.ptr)
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
	CGoReturnPtr := C.SmartTagProperty_SetValue( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteSmartTagProperty(smarttagproperty *SmartTagProperty){
	runtime.SetFinalizer(smarttagproperty, nil)
	C.Delete_SmartTagProperty(smarttagproperty.ptr)
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
	CGoReturnPtr := C.New_SmartTagPropertyCollection()
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
	CGoReturnPtr := C.SmartTagPropertyCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets a <see cref="SmartTagProperty"/> object.
// Parameters:
//   index - int32 
// Returns:
//   SmartTagProperty  
func (instance *SmartTagPropertyCollection) Get_Int(index int32)  (*SmartTagProperty,  error)  {
	CGoReturnPtr := C.SmartTagPropertyCollection_Get_Integer( instance.ptr, C.int(index))
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
	CGoReturnPtr := C.SmartTagPropertyCollection_Get_String( instance.ptr, C.CString(name))
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
	CGoReturnPtr := C.SmartTagPropertyCollection_Add( instance.ptr, C.CString(name), C.CString(value))
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
	CGoReturnPtr := C.SmartTagPropertyCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteSmartTagPropertyCollection(smarttagpropertycollection *SmartTagPropertyCollection){
	runtime.SetFinalizer(smarttagpropertycollection, nil)
	C.Delete_SmartTagPropertyCollection(smarttagpropertycollection.ptr)
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
	CGoReturnPtr := C.SmartTagSetting_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Adds a <see cref="SmartTagCollection"/> object to a cell.
// Parameters:
//   row - int32 
//   column - int32 
// Returns:
//   int32  
func (instance *SmartTagSetting) Add_Int_Int(row int32, column int32)  (int32,  error)  {
	CGoReturnPtr := C.SmartTagSetting_Add_Integer_Integer( instance.ptr, C.int(row), C.int(column))
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
	CGoReturnPtr := C.SmartTagSetting_Add_String( instance.ptr, C.CString(cellname))
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
	CGoReturnPtr := C.SmartTagSetting_Get_Integer( instance.ptr, C.int(index))
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
	CGoReturnPtr := C.SmartTagSetting_Get_Integer_Integer( instance.ptr, C.int(row), C.int(column))
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
	CGoReturnPtr := C.SmartTagSetting_Get_String( instance.ptr, C.CString(cellname))
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
	CGoReturnPtr := C.SmartTagSetting_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteSmartTagSetting(smarttagsetting *SmartTagSetting){
	runtime.SetFinalizer(smarttagsetting, nil)
	C.Delete_SmartTagSetting(smarttagsetting.ptr)
	smarttagsetting.ptr = nil
}
