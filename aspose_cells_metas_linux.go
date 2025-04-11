// +build linux

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


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

/**************Enum SensitivityLabelAssignmentType *****************/

// Represents the assignment method for the sensitivity label.
type SensitivityLabelAssignmentType int32

const(
// Use for any sensitivity label that was not directly applied by the user.
SensitivityLabelAssignmentType_Standard SensitivityLabelAssignmentType = 0 

// Use for any sensitivity label that was directly applied by the user.
SensitivityLabelAssignmentType_Privileged SensitivityLabelAssignmentType = 1 
)

func Int32ToSensitivityLabelAssignmentType(value int32)(SensitivityLabelAssignmentType ,error){
	switch value {
		case 0:  return SensitivityLabelAssignmentType_Standard, nil  
		case 1:  return SensitivityLabelAssignmentType_Privileged, nil  
		default:
			return 0 ,fmt.Errorf("invalid SensitivityLabelAssignmentType value: %d", value)
	}
}

/**************Enum SensitivityLabelMarkType *****************/

// Represents the types of content marking that ought to be applied to a file.
type SensitivityLabelMarkType int32

const(
// None
SensitivityLabelMarkType_None SensitivityLabelMarkType = 0 

// Header
SensitivityLabelMarkType_Header SensitivityLabelMarkType = 1 

// Footer
SensitivityLabelMarkType_Footer SensitivityLabelMarkType = 2 

// Watermark
SensitivityLabelMarkType_Watermark SensitivityLabelMarkType = 4 

// Encryption
SensitivityLabelMarkType_Encryption SensitivityLabelMarkType = 8 
)

func Int32ToSensitivityLabelMarkType(value int32)(SensitivityLabelMarkType ,error){
	switch value {
		case 0:  return SensitivityLabelMarkType_None, nil  
		case 1:  return SensitivityLabelMarkType_Header, nil  
		case 2:  return SensitivityLabelMarkType_Footer, nil  
		case 4:  return SensitivityLabelMarkType_Watermark, nil  
		case 8:  return SensitivityLabelMarkType_Encryption, nil  
		default:
			return 0 ,fmt.Errorf("invalid SensitivityLabelMarkType value: %d", value)
	}
}
// Class SensitivityLabel 

// Represents the sensitivity label.
type SensitivityLabel struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SensitivityLabel) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the id of sensitivity label.
// Returns:
//   string  
func (instance *SensitivityLabel) GetId()  (string,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the id of sensitivity label.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SensitivityLabel) SetId(value string)  error {
	
	CGoReturnPtr := C.SensitivityLabel_SetId( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the sensitivity label is enabled
// Returns:
//   bool  
func (instance *SensitivityLabel) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the sensitivity label is enabled
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SensitivityLabel) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.SensitivityLabel_SetIsEnabled( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the assignment method for the sensitivity label.
// Returns:
//   int32  
func (instance *SensitivityLabel) GetAssignmentType()  (SensitivityLabelAssignmentType,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_GetAssignmentType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSensitivityLabelAssignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the assignment method for the sensitivity label.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SensitivityLabel) SetAssignmentType(value SensitivityLabelAssignmentType)  error {
	
	CGoReturnPtr := C.SensitivityLabel_SetAssignmentType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the Azure Active Directory (Azure AD) site identifier corresponding to the sensitivity label policy which describes the sensitivity label.
// Returns:
//   string  
func (instance *SensitivityLabel) GetSiteId()  (string,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_GetSiteId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the Azure Active Directory (Azure AD) site identifier corresponding to the sensitivity label policy which describes the sensitivity label.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SensitivityLabel) SetSiteId(value string)  error {
	
	CGoReturnPtr := C.SensitivityLabel_SetSiteId( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the types of content marking that ought to be applied to a file.
// Returns:
//   int32  
func (instance *SensitivityLabel) GetContentMarkType()  (SensitivityLabelMarkType,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_GetContentMarkType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSensitivityLabelMarkType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the types of content marking that ought to be applied to a file.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SensitivityLabel) SetContentMarkType(value SensitivityLabelMarkType)  error {
	
	CGoReturnPtr := C.SensitivityLabel_SetContentMarkType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the sensitivity label was removed.
// Returns:
//   bool  
func (instance *SensitivityLabel) IsRemoved()  (bool,  error)  {
	
	CGoReturnPtr := C.SensitivityLabel_IsRemoved( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the sensitivity label was removed.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SensitivityLabel) SetIsRemoved(value bool)  error {
	
	CGoReturnPtr := C.SensitivityLabel_SetIsRemoved( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteSensitivityLabel(sensitivitylabel *SensitivityLabel){
	runtime.SetFinalizer(sensitivitylabel, nil)
	C.Delete_SensitivityLabel(sensitivitylabel.ptr)
	sensitivitylabel.ptr = nil
}

// Class SensitivityLabelCollection 

// Represents the list of sensitivity labels.
type SensitivityLabelCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewSensitivityLabelCollection() ( *SensitivityLabelCollection, error) {
	sensitivitylabelcollection := &SensitivityLabelCollection{}
	CGoReturnPtr := C.New_SensitivityLabelCollection()
	if CGoReturnPtr.error_no == 0 {
		sensitivitylabelcollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(sensitivitylabelcollection, DeleteSensitivityLabelCollection)
		return sensitivitylabelcollection, nil
	} else {
		sensitivitylabelcollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return sensitivitylabelcollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SensitivityLabelCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.SensitivityLabelCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a sensitivity label.
// Parameters:
//   id - string 
//   isEnabled - bool 
//   methodType - int32 
//   siteId - string 
//   markType - int32 
// Returns:
//   int32  
func (instance *SensitivityLabelCollection) Add(id string, isenabled bool, methodtype SensitivityLabelAssignmentType, siteid string, marktype SensitivityLabelMarkType)  (int32,  error)  {
	
	CGoReturnPtr := C.SensitivityLabelCollection_Add( instance.ptr, C.CString(id), C.bool(isenabled), C.int( int32(methodtype)), C.CString(siteid), C.int( int32(marktype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *SensitivityLabelCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.SensitivityLabelCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteSensitivityLabelCollection(sensitivitylabelcollection *SensitivityLabelCollection){
	runtime.SetFinalizer(sensitivitylabelcollection, nil)
	C.Delete_SensitivityLabelCollection(sensitivitylabelcollection.ptr)
	sensitivitylabelcollection.ptr = nil
}
