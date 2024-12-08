// +build linux

package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./lib/linux_x86_64 -lAspose.Cells.CWrapper
// #include <AsposeCellsCWrapper.h>
import "C"
import (
	"fmt"  
	"errors"
	"runtime"
	"unsafe" 
)

/**************Enum XAdESType *****************/
// Type of XML Advanced Electronic Signature (XAdES).
type XAdESType int32

const(
// XAdES is off.
XAdESType_None XAdESType = 0 

// Basic XAdES.
XAdESType_XAdES XAdESType = 1 
)

func Int32ToXAdESType(value int32)(XAdESType ,error){
	switch value {
		case 0:  return XAdESType_None, nil  
		case 1:  return XAdESType_XAdES, nil  
		default:
			return 0 ,fmt.Errorf("invalid XAdESType value: %d", value)
	}
}
// Class DigitalSignature 

// Signature in file.
type DigitalSignature struct {
	ptr unsafe.Pointer
}

// Constructor of DigitalSignature.
// Parameters:
//   fileName - string 
//   password - string 
//   comments - string 
//   signTime - Date 
func NewDigitalSignature(filename string, password string, comments string, signtime *Date) ( *DigitalSignature, error) {
	digitalsignature := &DigitalSignature{}
	CGoReturnPtr := C.New_DigitalSignature(C.CString(filename), C.CString(password), C.CString(comments), signtime.ptr)
	if CGoReturnPtr.error_no == 0 {
		digitalsignature.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(digitalsignature, DeleteDigitalSignature)
		return digitalsignature, nil
	} else {
		digitalsignature.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return digitalsignature, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DigitalSignature) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.DigitalSignature_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// The purpose to signature.
// Returns:
//   string  
func (instance *DigitalSignature) GetComments()  (string,  error)  {
	CGoReturnPtr := C.DigitalSignature_GetComments( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The purpose to signature.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DigitalSignature) SetComments(value string)  error {
	CGoReturnPtr := C.DigitalSignature_SetComments( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The time when the document was signed.
// Returns:
//   Date  
func (instance *DigitalSignature) GetSignTime()  (*Date,  error)  {
	CGoReturnPtr := C.DigitalSignature_GetSignTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// The time when the document was signed.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *DigitalSignature) SetSignTime(value *Date)  error {
	CGoReturnPtr := C.DigitalSignature_SetSignTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a GUID which can be cross-referenced with the GUID of the signature line stored in the document content.
// Default value is Empty (all zeroes) Guid.
// Parameters:
//   uuid - UUID 
// Returns:
//   void  
func (instance *DigitalSignature) GetId(uuid *UUID)  error {
	CGoReturnPtr := C.DigitalSignature_GetId( instance.ptr, uuid.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a GUID which can be cross-referenced with the GUID of the signature line stored in the document content.
// Default value is Empty (all zeroes) Guid.
// Parameters:
//   value - UUID 
// Returns:
//   void  
func (instance *DigitalSignature) SetId(value *UUID)  error {
	CGoReturnPtr := C.DigitalSignature_SetId( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the text of actual signature in the digital signature.
// Default value is Empty.
// Returns:
//   string  
func (instance *DigitalSignature) GetText()  (string,  error)  {
	CGoReturnPtr := C.DigitalSignature_GetText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the text of actual signature in the digital signature.
// Default value is Empty.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DigitalSignature) SetText(value string)  error {
	CGoReturnPtr := C.DigitalSignature_SetText( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the class ID of the signature provider.
// Default value is Empty (all zeroes) Guid.
// Parameters:
//   uuid - UUID 
// Returns:
//   void  
func (instance *DigitalSignature) GetProviderId(uuid *UUID)  error {
	CGoReturnPtr := C.DigitalSignature_GetProviderId( instance.ptr, uuid.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the class ID of the signature provider.
// Default value is Empty (all zeroes) Guid.
// Parameters:
//   value - UUID 
// Returns:
//   void  
func (instance *DigitalSignature) SetProviderId(value *UUID)  error {
	CGoReturnPtr := C.DigitalSignature_SetProviderId( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If this digital signature is valid and the document has not been tampered with,
// this value will be true.
// Returns:
//   bool  
func (instance *DigitalSignature) IsValid()  (bool,  error)  {
	CGoReturnPtr := C.DigitalSignature_IsValid( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// XAdES type.
// Default value is None(XAdES is off).
// Returns:
//   int32  
func (instance *DigitalSignature) GetXAdESType()  (XAdESType,  error)  {
	CGoReturnPtr := C.DigitalSignature_GetXAdESType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToXAdESType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// XAdES type.
// Default value is None(XAdES is off).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DigitalSignature) SetXAdESType(value XAdESType)  error {
	CGoReturnPtr := C.DigitalSignature_SetXAdESType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteDigitalSignature(digitalsignature *DigitalSignature){
	runtime.SetFinalizer(digitalsignature, nil)
	C.Delete_DigitalSignature(digitalsignature.ptr)
	digitalsignature.ptr = nil
}

// Class DigitalSignatureCollection 

// Provides a collection of digital signatures attached to a document.
type DigitalSignatureCollection struct {
	ptr unsafe.Pointer
}

// The constructor of DigitalSignatureCollection.
func NewDigitalSignatureCollection() ( *DigitalSignatureCollection, error) {
	digitalsignaturecollection := &DigitalSignatureCollection{}
	CGoReturnPtr := C.New_DigitalSignatureCollection()
	if CGoReturnPtr.error_no == 0 {
		digitalsignaturecollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(digitalsignaturecollection, DeleteDigitalSignatureCollection)
		return digitalsignaturecollection, nil
	} else {
		digitalsignaturecollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return digitalsignaturecollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DigitalSignatureCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.DigitalSignatureCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Add one signature to DigitalSignatureCollection.
// Parameters:
//   digitalSignature - DigitalSignature 
// Returns:
//   void  
func (instance *DigitalSignatureCollection) Add(digitalsignature *DigitalSignature)  error {
	CGoReturnPtr := C.DigitalSignatureCollection_Add( instance.ptr, digitalsignature.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteDigitalSignatureCollection(digitalsignaturecollection *DigitalSignatureCollection){
	runtime.SetFinalizer(digitalsignaturecollection, nil)
	C.Delete_DigitalSignatureCollection(digitalsignaturecollection.ptr)
	digitalsignaturecollection.ptr = nil
}
