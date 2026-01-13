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

/**************Enum VbaProjectReferenceType *****************/

// Represents the type of VBA project reference.
type VbaProjectReferenceType int32

const(
// Specifies a reference to an Automation type library.
VbaProjectReferenceType_Registered VbaProjectReferenceType = 0 

// Specifies a reference to a twiddled type library and its extended type library.
VbaProjectReferenceType_Control VbaProjectReferenceType = 1 

// Specifies a reference to an external VBA project.
VbaProjectReferenceType_Project VbaProjectReferenceType = 2 
)

func Int32ToVbaProjectReferenceType(value int32)(VbaProjectReferenceType ,error){
	switch value {
		case 0:  return VbaProjectReferenceType_Registered, nil  
		case 1:  return VbaProjectReferenceType_Control, nil  
		case 2:  return VbaProjectReferenceType_Project, nil  
		default:
			return 0 ,fmt.Errorf("invalid VbaProjectReferenceType value: %d", value)
	}
}
// Class VbaModule 

// Represents the module in VBA project.
type VbaModule struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *VbaModule) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaModule_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of Module.
// Returns:
//   string  
func (instance *VbaModule) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaModule_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of Module.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaModule) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaModule_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the type of module.
// Returns:
//   VbaModuleType  
func (instance *VbaModule) GetType()  (*VbaModuleType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("VbaModule_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &VbaModuleType{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteVbaModuleType) 

	return result, nil 
}
// Gets and sets the binary codes of module.
// Returns:
//   []byte  
func (instance *VbaModule) GetBinaryCodes()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("VbaModule_GetBinaryCodes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the codes of module.
// Returns:
//   string  
func (instance *VbaModule) GetCodes()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaModule_GetCodes"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the codes of module.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaModule) SetCodes(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaModule_SetCodes"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteVbaModule(vbamodule *VbaModule){
	runtime.SetFinalizer(vbamodule, nil)
	C.Delete_CObject(C.CString("Delete_VbaModule"),vbamodule.ptr)
	vbamodule.ptr = nil
}

// Class VbaModuleCollection 

// Represents the list of <see cref="VbaModule"/>
type VbaModuleCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *VbaModuleCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaModuleCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Parameters:
//   name - string 
//   data - []byte 
// Returns:
//   void  
func (instance *VbaModuleCollection) AddDesignerStorage(name string, data []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPP(C.CString("VbaModuleCollection_AddDesignerStorage"), instance.ptr, C.CString(name), unsafe.Pointer(&data[0]), C.int( len(data)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the data of Designer.
// Parameters:
//   name - string 
// Returns:
//   []byte  
func (instance *VbaModuleCollection) GetDesignerStorage(name string)  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJT(C.CString("VbaModuleCollection_GetDesignerStorage"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Adds module for a worksheet.
// Parameters:
//   sheet - Worksheet 
// Returns:
//   int32  
func (instance *VbaModuleCollection) Add_Worksheet(sheet *Worksheet)  (int32,  error)  {
	
	var sheet_ptr unsafe.Pointer = nil
	if sheet != nil {
	  sheet_ptr =sheet.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZFM(C.CString("VbaModuleCollection_Add_Worksheet"), instance.ptr, sheet_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds module.
// Parameters:
//   type - VbaModuleType 
//   name - string 
// Returns:
//   int32  
func (instance *VbaModuleCollection) Add_VbaModuleType_String(type_ *VbaModuleType, name string)  (int32,  error)  {
	
	var type__ptr unsafe.Pointer = nil
	if type_ != nil {
	  type__ptr =type_.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZPQ(C.CString("VbaModuleCollection_Add_VbaModuleType_String"), instance.ptr, type__ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Inser user form into VBA Project.
// Parameters:
//   name - string 
//   codes - string 
//   designerStorage - []byte 
// Returns:
//   int32  
func (instance *VbaModuleCollection) AddUserForm(name string, codes string, designerstorage []byte)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPR(C.CString("VbaModuleCollection_AddUserForm"), instance.ptr, C.CString(name), C.CString(codes), unsafe.Pointer(&designerstorage[0]), C.int( len(designerstorage)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets <see cref="VbaModule"/> in the list by the index.
// Parameters:
//   index - int32 
// Returns:
//   VbaModule  
func (instance *VbaModuleCollection) Get_Int(index int32)  (*VbaModule,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("VbaModuleCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &VbaModule{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteVbaModule) 

	return result, nil 
}
// Removes module for a worksheet.
// Parameters:
//   sheet - Worksheet 
// Returns:
//   void  
func (instance *VbaModuleCollection) Remove_Worksheet(sheet *Worksheet)  error {
	
	var sheet_ptr unsafe.Pointer = nil
	if sheet != nil {
	  sheet_ptr =sheet.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("VbaModuleCollection_Remove_Worksheet"), instance.ptr, sheet_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Remove the module by the name
// Parameters:
//   name - string 
// Returns:
//   void  
func (instance *VbaModuleCollection) Remove_String(name string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaModuleCollection_Remove_String"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets <see cref="VbaModule"/> in the list by the name.
// Parameters:
//   name - string 
// Returns:
//   VbaModule  
func (instance *VbaModuleCollection) Get_String(name string)  (*VbaModule,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("VbaModuleCollection_Get_String"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &VbaModule{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteVbaModule) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *VbaModuleCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("VbaModuleCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteVbaModuleCollection(vbamodulecollection *VbaModuleCollection){
	runtime.SetFinalizer(vbamodulecollection, nil)
	C.Delete_CObject(C.CString("Delete_VbaModuleCollection"),vbamodulecollection.ptr)
	vbamodulecollection.ptr = nil
}

// Class VbaModuleType 

// Represents a class module.
type VbaModuleType struct {
	ptr unsafe.Pointer
}





func DeleteVbaModuleType(vbamoduletype *VbaModuleType){
	runtime.SetFinalizer(vbamoduletype, nil)
	C.Delete_CObject(C.CString("Delete_VbaModuleType"),vbamoduletype.ptr)
	vbamoduletype.ptr = nil
}

// Class VbaProject 

// Represents the VBA project.
type VbaProject struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *VbaProject) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProject_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Sign this VBA project by a DigitalSignature
// Parameters:
//   digitalSignature - DigitalSignature 
// Returns:
//   void  
func (instance *VbaProject) Sign(digitalsignature *DigitalSignature)  error {
	
	var digitalsignature_ptr unsafe.Pointer = nil
	if digitalsignature != nil {
	  digitalsignature_ptr =digitalsignature.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("VbaProject_Sign"), instance.ptr, digitalsignature_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the signature of VBA project is valid or not.
// Returns:
//   bool  
func (instance *VbaProject) IsValidSigned()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProject_IsValidSigned"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets certificate raw data if this VBA project is signed.
// Returns:
//   []byte  
func (instance *VbaProject) GetCertRawData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("VbaProject_GetCertRawData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the encoding of VBA project.
// Returns:
//   int32  
func (instance *VbaProject) GetEncoding()  (EncodingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("VbaProject_GetEncoding"), instance.ptr)
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
// Gets and sets the encoding of VBA project.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *VbaProject) SetEncoding(value EncodingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("VbaProject_SetEncoding"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the name of the VBA project.
// Returns:
//   string  
func (instance *VbaProject) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaProject_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the VBA project.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaProject) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaProject_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether VBAcode is signed or not.
// Returns:
//   bool  
func (instance *VbaProject) IsSigned()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProject_IsSigned"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Protects or unprotects this VBA project.
// Parameters:
//   islockedForViewing - bool 
//   password - string 
// Returns:
//   void  
func (instance *VbaProject) Protect(islockedforviewing bool, password string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZIQ(C.CString("VbaProject_Protect"), instance.ptr, C.bool(islockedforviewing), C.CString(password))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this VBA project is protected.
// Returns:
//   bool  
func (instance *VbaProject) IsProtected()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProject_IsProtected"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this VBA project is locked for viewing.
// Returns:
//   bool  
func (instance *VbaProject) GetIslockedForViewing()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProject_GetIslockedForViewing"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Copy VBA project from other file.
// Parameters:
//   source - VbaProject 
// Returns:
//   void  
func (instance *VbaProject) Copy(source *VbaProject)  error {
	
	var source_ptr unsafe.Pointer = nil
	if source != nil {
	  source_ptr =source.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("VbaProject_Copy"), instance.ptr, source_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets all <see cref="VbaModule"/> objects.
// Returns:
//   VbaModuleCollection  
func (instance *VbaProject) GetModules()  (*VbaModuleCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("VbaProject_GetModules"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &VbaModuleCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteVbaModuleCollection) 

	return result, nil 
}
// Gets all references of VBA project.
// Returns:
//   VbaProjectReferenceCollection  
func (instance *VbaProject) GetReferences()  (*VbaProjectReferenceCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("VbaProject_GetReferences"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &VbaProjectReferenceCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteVbaProjectReferenceCollection) 

	return result, nil 
}
// Validates protection password.
// Parameters:
//   password - string 
// Returns:
//   bool  
func (instance *VbaProject) ValidatePassword(password string)  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZG(C.CString("VbaProject_ValidatePassword"), instance.ptr, C.CString(password))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteVbaProject(vbaproject *VbaProject){
	runtime.SetFinalizer(vbaproject, nil)
	C.Delete_CObject(C.CString("Delete_VbaProject"),vbaproject.ptr)
	vbaproject.ptr = nil
}

// Class VbaProjectReference 

// Represents the reference of VBA project.
type VbaProjectReference struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *VbaProjectReference) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProjectReference_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Parameters:
//   source - VbaProjectReference 
// Returns:
//   void  
func (instance *VbaProjectReference) Copy(source *VbaProjectReference)  error {
	
	var source_ptr unsafe.Pointer = nil
	if source != nil {
	  source_ptr =source.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("VbaProjectReference_Copy"), instance.ptr, source_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the type of this reference.
// Returns:
//   int32  
func (instance *VbaProjectReference) GetType()  (VbaProjectReferenceType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("VbaProjectReference_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToVbaProjectReferenceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the name of the reference.
// Returns:
//   string  
func (instance *VbaProjectReference) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaProjectReference_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the reference.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaProjectReference) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaProjectReference_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the Libid of the reference.
// Returns:
//   string  
func (instance *VbaProjectReference) GetLibid()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaProjectReference_GetLibid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the Libid of the reference.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaProjectReference) SetLibid(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaProjectReference_SetLibid"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the twiddled Libid of the reference.
// Returns:
//   string  
func (instance *VbaProjectReference) GetTwiddledlibid()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaProjectReference_GetTwiddledlibid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the twiddled Libid of the reference.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaProjectReference) SetTwiddledlibid(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaProjectReference_SetTwiddledlibid"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the extended Libid of the reference.
// Returns:
//   string  
func (instance *VbaProjectReference) GetExtendedLibid()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaProjectReference_GetExtendedLibid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the extended Libid of the reference.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaProjectReference) SetExtendedLibid(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaProjectReference_SetExtendedLibid"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the referenced VBA project's identifier with an relative path.
// Returns:
//   string  
func (instance *VbaProjectReference) GetRelativeLibid()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("VbaProjectReference_GetRelativeLibid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the referenced VBA project's identifier with an relative path.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *VbaProjectReference) SetRelativeLibid(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("VbaProjectReference_SetRelativeLibid"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteVbaProjectReference(vbaprojectreference *VbaProjectReference){
	runtime.SetFinalizer(vbaprojectreference, nil)
	C.Delete_CObject(C.CString("Delete_VbaProjectReference"),vbaprojectreference.ptr)
	vbaprojectreference.ptr = nil
}

// Class VbaProjectReferenceCollection 

// Represents all references of VBA project.
type VbaProjectReferenceCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *VbaProjectReferenceCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("VbaProjectReferenceCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Get the reference in the list by the index.
// Parameters:
//   i - int32 
// Returns:
//   VbaProjectReference  
func (instance *VbaProjectReferenceCollection) Get(i int32)  (*VbaProjectReference,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("VbaProjectReferenceCollection_Get"), instance.ptr, C.int(i))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &VbaProjectReference{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteVbaProjectReference) 

	return result, nil 
}
// Add a reference to an Automation type library.
// Parameters:
//   name - string 
//   libid - string 
// Returns:
//   int32  
func (instance *VbaProjectReferenceCollection) AddRegisteredReference(name string, libid string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJG(C.CString("VbaProjectReferenceCollection_AddRegisteredReference"), instance.ptr, C.CString(name), C.CString(libid))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a reference to a twiddled type library and its extended type library.
// Parameters:
//   name - string 
//   libid - string 
//   twiddledlibid - string 
//   extendedLibid - string 
// Returns:
//   int32  
func (instance *VbaProjectReferenceCollection) AddControlRefrernce(name string, libid string, twiddledlibid string, extendedlibid string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZPS(C.CString("VbaProjectReferenceCollection_AddControlRefrernce"), instance.ptr, C.CString(name), C.CString(libid), C.CString(twiddledlibid), C.CString(extendedlibid))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a reference to an external VBA project.
// Parameters:
//   name - string 
//   absoluteLibid - string 
//   relativeLibid - string 
// Returns:
//   int32  
func (instance *VbaProjectReferenceCollection) AddProjectRefrernce(name string, absolutelibid string, relativelibid string)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZIX(C.CString("VbaProjectReferenceCollection_AddProjectRefrernce"), instance.ptr, C.CString(name), C.CString(absolutelibid), C.CString(relativelibid))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Copies references from other VBA project.
// Parameters:
//   source - VbaProjectReferenceCollection 
// Returns:
//   void  
func (instance *VbaProjectReferenceCollection) Copy(source *VbaProjectReferenceCollection)  error {
	
	var source_ptr unsafe.Pointer = nil
	if source != nil {
	  source_ptr =source.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("VbaProjectReferenceCollection_Copy"), instance.ptr, source_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *VbaProjectReferenceCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("VbaProjectReferenceCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteVbaProjectReferenceCollection(vbaprojectreferencecollection *VbaProjectReferenceCollection){
	runtime.SetFinalizer(vbaprojectreferencecollection, nil)
	C.Delete_CObject(C.CString("Delete_VbaProjectReferenceCollection"),vbaprojectreferencecollection.ptr)
	vbaprojectreferencecollection.ptr = nil
}
