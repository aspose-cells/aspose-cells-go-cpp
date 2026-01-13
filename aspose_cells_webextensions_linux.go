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

/**************Enum WebExtensionStoreType *****************/

// Represents the store type of web extension.
type WebExtensionStoreType int32

const(
// Specifies that the store type is Office.com.
WebExtensionStoreType_OMEX WebExtensionStoreType = 0 

// Specifies that the store type is SharePoint corporate catalog.
WebExtensionStoreType_SPCatalog WebExtensionStoreType = 1 

// Specifies that the store type is a SharePoint web application.
WebExtensionStoreType_SPApp WebExtensionStoreType = 2 

// Specifies that the store type is an Exchange server.
WebExtensionStoreType_Exchange WebExtensionStoreType = 3 

// Specifies that the store type is a file system share.
WebExtensionStoreType_FileSystem WebExtensionStoreType = 4 

// Specifies that the store type is the system registry.
WebExtensionStoreType_Registry WebExtensionStoreType = 5 

// Specifies that the store type is Centralized Deployment via Exchange.
WebExtensionStoreType_ExCatalog WebExtensionStoreType = 6 
)

func Int32ToWebExtensionStoreType(value int32)(WebExtensionStoreType ,error){
	switch value {
		case 0:  return WebExtensionStoreType_OMEX, nil  
		case 1:  return WebExtensionStoreType_SPCatalog, nil  
		case 2:  return WebExtensionStoreType_SPApp, nil  
		case 3:  return WebExtensionStoreType_Exchange, nil  
		case 4:  return WebExtensionStoreType_FileSystem, nil  
		case 5:  return WebExtensionStoreType_Registry, nil  
		case 6:  return WebExtensionStoreType_ExCatalog, nil  
		default:
			return 0 ,fmt.Errorf("invalid WebExtensionStoreType value: %d", value)
	}
}
// Class WebExtension 

// Represents an Office Add-in instance.
type WebExtension struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtension) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtension_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the uniquely identifies the Office Add-in instance in the current document.
// Returns:
//   string  
func (instance *WebExtension) GetId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtension_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the uniquely identifies the Office Add-in instance in the current document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtension) SetId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtension_SetId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the user can interact with the Office Add-in or not.
// Returns:
//   bool  
func (instance *WebExtension) IsFrozen()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtension_IsFrozen"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the user can interact with the Office Add-in or not.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebExtension) SetIsFrozen(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebExtension_SetIsFrozen"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get the primary reference to an Office Add-in.
// Returns:
//   WebExtensionReference  
func (instance *WebExtension) GetReference()  (*WebExtensionReference,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebExtension_GetReference"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionReference{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionReference) 

	return result, nil 
}
// Gets a list of alter references.
// Returns:
//   WebExtensionReferenceCollection  
func (instance *WebExtension) GetAlterReferences()  (*WebExtensionReferenceCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebExtension_GetAlterReferences"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionReferenceCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionReferenceCollection) 

	return result, nil 
}
// Gets all properties of web extension.
// Returns:
//   WebExtensionPropertyCollection  
func (instance *WebExtension) GetProperties()  (*WebExtensionPropertyCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebExtension_GetProperties"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionPropertyCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionPropertyCollection) 

	return result, nil 
}
// Gets all bindings relationship between an Office Add-in and the data in the document.
// Returns:
//   WebExtensionBindingCollection  
func (instance *WebExtension) GetBindings()  (*WebExtensionBindingCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebExtension_GetBindings"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionBindingCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionBindingCollection) 

	return result, nil 
}



func DeleteWebExtension(webextension *WebExtension){
	runtime.SetFinalizer(webextension, nil)
	C.Delete_CObject(C.CString("Delete_WebExtension"),webextension.ptr)
	webextension.ptr = nil
}

// Class WebExtensionBinding 

// Represents a binding relationship between an Office Add-in and the data in the document.
type WebExtensionBinding struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewWebExtensionBinding() ( *WebExtensionBinding, error) {
	webextensionbinding := &WebExtensionBinding{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_WebExtensionBinding"),)
	if CGoReturnPtr.error_no == 0 {
		webextensionbinding.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(webextensionbinding, DeleteWebExtensionBinding)
		return webextensionbinding, nil
	} else {
		webextensionbinding.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return webextensionbinding, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionBinding) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionBinding_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the binding identifier.
// Returns:
//   string  
func (instance *WebExtensionBinding) GetId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionBinding_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the binding identifier.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionBinding) SetId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionBinding_SetId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binding type.
// Returns:
//   string  
func (instance *WebExtensionBinding) GetType()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionBinding_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the binding type.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionBinding) SetType(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionBinding_SetType"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binding key used to map the binding entry in this list with the bound data in the document.
// Returns:
//   string  
func (instance *WebExtensionBinding) GetAppref()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionBinding_GetAppref"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the binding key used to map the binding entry in this list with the bound data in the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionBinding) SetAppref(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionBinding_SetAppref"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteWebExtensionBinding(webextensionbinding *WebExtensionBinding){
	runtime.SetFinalizer(webextensionbinding, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionBinding"),webextensionbinding.ptr)
	webextensionbinding.ptr = nil
}

// Class WebExtensionBindingCollection 

// Represents the list of binding relationships between an Office Add-in and the data in the document.
type WebExtensionBindingCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewWebExtensionBindingCollection() ( *WebExtensionBindingCollection, error) {
	webextensionbindingcollection := &WebExtensionBindingCollection{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_WebExtensionBindingCollection"),)
	if CGoReturnPtr.error_no == 0 {
		webextensionbindingcollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(webextensionbindingcollection, DeleteWebExtensionBindingCollection)
		return webextensionbindingcollection, nil
	} else {
		webextensionbindingcollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return webextensionbindingcollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionBindingCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionBindingCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets web extension binding relationship by the specific index.
// Parameters:
//   index - int32 
// Returns:
//   WebExtensionBinding  
func (instance *WebExtensionBindingCollection) Get(index int32)  (*WebExtensionBinding,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("WebExtensionBindingCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionBinding{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionBinding) 

	return result, nil 
}
// Adds an a binding relationship between an Office Add-in and the data in the document.
// Returns:
//   int32  
func (instance *WebExtensionBindingCollection) Add()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionBindingCollection_Add"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *WebExtensionBindingCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionBindingCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteWebExtensionBindingCollection(webextensionbindingcollection *WebExtensionBindingCollection){
	runtime.SetFinalizer(webextensionbindingcollection, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionBindingCollection"),webextensionbindingcollection.ptr)
	webextensionbindingcollection.ptr = nil
}

// Class WebExtensionCollection 

// Represents the list of web extension.
type WebExtensionCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets web extension by the specific index.
// Parameters:
//   index - int32 
// Returns:
//   WebExtension  
func (instance *WebExtensionCollection) Get(index int32)  (*WebExtension,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("WebExtensionCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtension{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtension) 

	return result, nil 
}
// Adds a web extension.
// Returns:
//   int32  
func (instance *WebExtensionCollection) Add()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionCollection_Add"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a web video player into exel.
// Parameters:
//   url - string 
//   autoPlay - bool 
//   startTime - int32 
//   endTime - int32 
// Returns:
//   int32  
func (instance *WebExtensionCollection) AddWebVideoPlayer(url string, autoplay bool, starttime int32, endtime int32)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPT(C.CString("WebExtensionCollection_AddWebVideoPlayer"), instance.ptr, C.CString(url), C.bool(autoplay), C.int(starttime), C.int(endtime))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Remove web extension by the index.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *WebExtensionCollection) RemoveAt(index int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("WebExtensionCollection_RemoveAt"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *WebExtensionCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteWebExtensionCollection(webextensioncollection *WebExtensionCollection){
	runtime.SetFinalizer(webextensioncollection, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionCollection"),webextensioncollection.ptr)
	webextensioncollection.ptr = nil
}

// Class WebExtensionProperty 

// Represents an Office Add-in custom property.
type WebExtensionProperty struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionProperty) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionProperty_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set a custom property name.
// Returns:
//   string  
func (instance *WebExtensionProperty) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionProperty_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set a custom property name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionProperty) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionProperty_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a custom property value.
// Returns:
//   string  
func (instance *WebExtensionProperty) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionProperty_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a custom property value.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionProperty) SetValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionProperty_SetValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteWebExtensionProperty(webextensionproperty *WebExtensionProperty){
	runtime.SetFinalizer(webextensionproperty, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionProperty"),webextensionproperty.ptr)
	webextensionproperty.ptr = nil
}

// Class WebExtensionPropertyCollection 

// Represents the list of web extension properties.
type WebExtensionPropertyCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewWebExtensionPropertyCollection() ( *WebExtensionPropertyCollection, error) {
	webextensionpropertycollection := &WebExtensionPropertyCollection{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_WebExtensionPropertyCollection"),)
	if CGoReturnPtr.error_no == 0 {
		webextensionpropertycollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(webextensionpropertycollection, DeleteWebExtensionPropertyCollection)
		return webextensionpropertycollection, nil
	} else {
		webextensionpropertycollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return webextensionpropertycollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionPropertyCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionPropertyCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the property of web extension by the index.
// Parameters:
//   index - int32 
// Returns:
//   WebExtensionProperty  
func (instance *WebExtensionPropertyCollection) Get_Int(index int32)  (*WebExtensionProperty,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("WebExtensionPropertyCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionProperty) 

	return result, nil 
}
// Gets the property of web extension.
// Parameters:
//   name - string 
// Returns:
//   WebExtensionProperty  
func (instance *WebExtensionPropertyCollection) Get_String(name string)  (*WebExtensionProperty,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("WebExtensionPropertyCollection_Get_String"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionProperty) 

	return result, nil 
}
// Adds web extension property.
// Parameters:
//   name - string 
//   value - string 
// Returns:
//   int32  
func (instance *WebExtensionPropertyCollection) Add(name string, value string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJG(C.CString("WebExtensionPropertyCollection_Add"), instance.ptr, C.CString(name), C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Remove the property by the name.
// Parameters:
//   name - string 
// Returns:
//   void  
func (instance *WebExtensionPropertyCollection) RemoveAt(name string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionPropertyCollection_RemoveAt"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *WebExtensionPropertyCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionPropertyCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteWebExtensionPropertyCollection(webextensionpropertycollection *WebExtensionPropertyCollection){
	runtime.SetFinalizer(webextensionpropertycollection, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionPropertyCollection"),webextensionpropertycollection.ptr)
	webextensionpropertycollection.ptr = nil
}

// Class WebExtensionReference 

// Represents identify the provider location and version of the extension.
type WebExtensionReference struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionReference) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionReference_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the identifier associated with the Office Add-in within a catalog provider.
// The identifier MUST be unique within a catalog provider.
// Returns:
//   string  
func (instance *WebExtensionReference) GetId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionReference_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the identifier associated with the Office Add-in within a catalog provider.
// The identifier MUST be unique within a catalog provider.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionReference) SetId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionReference_SetId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the version.
// Returns:
//   string  
func (instance *WebExtensionReference) GetVersion()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionReference_GetVersion"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the version.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionReference) SetVersion(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionReference_SetVersion"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the instance of the marketplace where the Office Add-in is stored. .
// Returns:
//   string  
func (instance *WebExtensionReference) GetStoreName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionReference_GetStoreName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the instance of the marketplace where the Office Add-in is stored. .
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionReference) SetStoreName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionReference_SetStoreName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of marketplace that the store attribute identifies.
// Returns:
//   int32  
func (instance *WebExtensionReference) GetStoreType()  (WebExtensionStoreType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebExtensionReference_GetStoreType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToWebExtensionStoreType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of marketplace that the store attribute identifies.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebExtensionReference) SetStoreType(value WebExtensionStoreType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("WebExtensionReference_SetStoreType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteWebExtensionReference(webextensionreference *WebExtensionReference){
	runtime.SetFinalizer(webextensionreference, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionReference"),webextensionreference.ptr)
	webextensionreference.ptr = nil
}

// Class WebExtensionReferenceCollection 

// Represents the list of web extension reference.
type WebExtensionReferenceCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewWebExtensionReferenceCollection() ( *WebExtensionReferenceCollection, error) {
	webextensionreferencecollection := &WebExtensionReferenceCollection{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_WebExtensionReferenceCollection"),)
	if CGoReturnPtr.error_no == 0 {
		webextensionreferencecollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(webextensionreferencecollection, DeleteWebExtensionReferenceCollection)
		return webextensionreferencecollection, nil
	} else {
		webextensionreferencecollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return webextensionreferencecollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionReferenceCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionReferenceCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets web extension by the specific index.
// Parameters:
//   index - int32 
// Returns:
//   WebExtensionReference  
func (instance *WebExtensionReferenceCollection) Get(index int32)  (*WebExtensionReference,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("WebExtensionReferenceCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionReference{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionReference) 

	return result, nil 
}
// Adds an empty reference of web extension.
// Returns:
//   int32  
func (instance *WebExtensionReferenceCollection) Add()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionReferenceCollection_Add"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *WebExtensionReferenceCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionReferenceCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteWebExtensionReferenceCollection(webextensionreferencecollection *WebExtensionReferenceCollection){
	runtime.SetFinalizer(webextensionreferencecollection, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionReferenceCollection"),webextensionreferencecollection.ptr)
	webextensionreferencecollection.ptr = nil
}

// Class WebExtensionTaskPane 

// Represents a persisted taskpane object.
type WebExtensionTaskPane struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionTaskPane) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionTaskPane_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the web extension part associated with the taskpane instance
// Returns:
//   WebExtension  
func (instance *WebExtensionTaskPane) GetWebExtension()  (*WebExtension,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebExtensionTaskPane_GetWebExtension"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtension{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtension) 

	return result, nil 
}
// Gets and sets the web extension part associated with the taskpane instance
// Parameters:
//   value - WebExtension 
// Returns:
//   void  
func (instance *WebExtensionTaskPane) SetWebExtension(value *WebExtension)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("WebExtensionTaskPane_SetWebExtension"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the last-docked location of this taskpane object.
// Returns:
//   string  
func (instance *WebExtensionTaskPane) GetDockState()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebExtensionTaskPane_GetDockState"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the last-docked location of this taskpane object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebExtensionTaskPane) SetDockState(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebExtensionTaskPane_SetDockState"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the Task Pane shows as visible by default when the document opens.
// Returns:
//   bool  
func (instance *WebExtensionTaskPane) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionTaskPane_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the Task Pane shows as visible by default when the document opens.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebExtensionTaskPane) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebExtensionTaskPane_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the taskpane is locked to the document in the UI and cannot be closed by the user.
// Returns:
//   bool  
func (instance *WebExtensionTaskPane) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionTaskPane_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the taskpane is locked to the document in the UI and cannot be closed by the user.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebExtensionTaskPane) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebExtensionTaskPane_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default width value for this taskpane instance.
// Returns:
//   float64  
func (instance *WebExtensionTaskPane) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("WebExtensionTaskPane_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the default width value for this taskpane instance.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *WebExtensionTaskPane) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("WebExtensionTaskPane_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the index, enumerating from the outside to the inside, of this taskpane among other persisted taskpanes docked in the same default location.
// Returns:
//   int32  
func (instance *WebExtensionTaskPane) GetRow()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionTaskPane_GetRow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the index, enumerating from the outside to the inside, of this taskpane among other persisted taskpanes docked in the same default location.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebExtensionTaskPane) SetRow(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("WebExtensionTaskPane_SetRow"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteWebExtensionTaskPane(webextensiontaskpane *WebExtensionTaskPane){
	runtime.SetFinalizer(webextensiontaskpane, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionTaskPane"),webextensiontaskpane.ptr)
	webextensiontaskpane.ptr = nil
}

// Class WebExtensionTaskPaneCollection 

// Represents the list of task pane.
type WebExtensionTaskPaneCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebExtensionTaskPaneCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebExtensionTaskPaneCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets task pane by the specific index.
// Parameters:
//   index - int32 
// Returns:
//   WebExtensionTaskPane  
func (instance *WebExtensionTaskPaneCollection) Get(index int32)  (*WebExtensionTaskPane,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("WebExtensionTaskPaneCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &WebExtensionTaskPane{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWebExtensionTaskPane) 

	return result, nil 
}
// Adds task pane.
// Returns:
//   int32  
func (instance *WebExtensionTaskPaneCollection) Add()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionTaskPaneCollection_Add"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *WebExtensionTaskPaneCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebExtensionTaskPaneCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteWebExtensionTaskPaneCollection(webextensiontaskpanecollection *WebExtensionTaskPaneCollection){
	runtime.SetFinalizer(webextensiontaskpanecollection, nil)
	C.Delete_CObject(C.CString("Delete_WebExtensionTaskPaneCollection"),webextensiontaskpanecollection.ptr)
	webextensiontaskpanecollection.ptr = nil
}
