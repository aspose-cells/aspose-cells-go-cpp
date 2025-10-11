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
	"time"  	
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
//   rawData - []byte 
//   password - string 
//   comments - string 
//   signTime - Date 
func NewDigitalSignature_Stream_String_String_Date(rawdata []byte, password string, comments string, signtime time.Time) ( *DigitalSignature, error) {
	digitalsignature := &DigitalSignature{}
	time_signtime := C.Get_Date( C.int(signtime.Year()), C.int(signtime.Month()) , C.int(signtime.Day()) , C.int(signtime.Hour()) , C.int(signtime.Minute()) , C.int(signtime.Second())  )

	CGoReturnPtr := C.CellsGoFunctoinZZOT(C.CString("New_DigitalSignature_Stream_String_String_Date"),unsafe.Pointer(&rawdata[0]), C.int( len(rawdata)), C.CString(password), C.CString(comments), time_signtime)
	C.Delete_GetDate( time_signtime)

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
// Constructor of DigitalSignature.
// Parameters:
//   fileName - string 
//   password - string 
//   comments - string 
//   signTime - Date 
func NewDigitalSignature_String_String_String_Date(filename string, password string, comments string, signtime time.Time) ( *DigitalSignature, error) {
	digitalsignature := &DigitalSignature{}
	time_signtime := C.Get_Date( C.int(signtime.Year()), C.int(signtime.Month()) , C.int(signtime.Day()) , C.int(signtime.Hour()) , C.int(signtime.Minute()) , C.int(signtime.Second())  )

	CGoReturnPtr := C.CellsGoFunctoinZZOU(C.CString("New_DigitalSignature_String_String_String_Date"),C.CString(filename), C.CString(password), C.CString(comments), time_signtime)
	C.Delete_GetDate( time_signtime)

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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DigitalSignature_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// The purpose to signature.
// Returns:
//   string  
func (instance *DigitalSignature) GetComments()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DigitalSignature_GetComments"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DigitalSignature_SetComments"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The time when the document was signed.
// Returns:
//   Date  
func (instance *DigitalSignature) GetSignTime()  (time.Time,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("DigitalSignature_GetSignTime"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  time.Unix(0, 0), err
	}
	result := time.Date(int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_year"), CGoReturnPtr.return_value).return_value ),time.Month(int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_month"), CGoReturnPtr.return_value).return_value )),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_day"), CGoReturnPtr.return_value).return_value ),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_hour"), CGoReturnPtr.return_value).return_value ),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_minute"), CGoReturnPtr.return_value).return_value ),int( C.CellsGoFunctoinZZBO(C.CString("Date_Get_second"), CGoReturnPtr.return_value).return_value ), 0, time.UTC) 

	return result, nil 
}
// The time when the document was signed.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *DigitalSignature) SetSignTime(value time.Time)  error {
	
	time_value := C.Get_Date( C.int(value.Year()), C.int(value.Month()) , C.int(value.Day()) , C.int(value.Hour()) , C.int(value.Minute()) , C.int(value.Second())  )

	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("DigitalSignature_SetSignTime"), instance.ptr, time_value)
	C.Delete_GetDate( time_value)

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
	
	var uuid_ptr unsafe.Pointer = nil
	if uuid != nil {
	  uuid_ptr =uuid.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DigitalSignature_GetId"), instance.ptr, uuid_ptr)
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
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DigitalSignature_SetId"), instance.ptr, value_ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DigitalSignature_GetText"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DigitalSignature_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies an image for the digital signature.
// Default value is null.
// Returns:
//   []byte  
func (instance *DigitalSignature) GetImage()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("DigitalSignature_GetImage"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Specifies an image for the digital signature.
// Default value is null.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *DigitalSignature) SetImage(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("DigitalSignature_SetImage"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	var uuid_ptr unsafe.Pointer = nil
	if uuid != nil {
	  uuid_ptr =uuid.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DigitalSignature_GetProviderId"), instance.ptr, uuid_ptr)
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
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DigitalSignature_SetProviderId"), instance.ptr, value_ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DigitalSignature_IsValid"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// XAdES type.
// Default value is None(XAdES is off).
// Returns:
//   int32  
func (instance *DigitalSignature) GetXAdESType()  (XAdESType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DigitalSignature_GetXAdESType"), instance.ptr)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DigitalSignature_SetXAdESType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteDigitalSignature(digitalsignature *DigitalSignature){
	runtime.SetFinalizer(digitalsignature, nil)
	C.Delete_CObject(C.CString("Delete_DigitalSignature"),digitalsignature.ptr)
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
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_DigitalSignatureCollection"),)
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
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DigitalSignatureCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add one signature to DigitalSignatureCollection.
// Parameters:
//   digitalSignature - DigitalSignature 
// Returns:
//   void  
func (instance *DigitalSignatureCollection) Add(digitalsignature *DigitalSignature)  error {
	
	var digitalsignature_ptr unsafe.Pointer = nil
	if digitalsignature != nil {
	  digitalsignature_ptr =digitalsignature.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZH(C.CString("DigitalSignatureCollection_Add"), instance.ptr, digitalsignature_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get the enumerator for DigitalSignatureCollection,
// this enumerator allows iteration over the collection
// Returns:
//   unsafe.Pointer  
func (instance *DigitalSignatureCollection) GetEnumerator()  (*DigitalSignatureEnumerator,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAW(C.CString("DigitalSignatureCollection_GetEnumerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DigitalSignatureEnumerator{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteDigitalSignatureEnumerator)
	 

	return result, nil 
}



func DeleteDigitalSignatureCollection(digitalsignaturecollection *DigitalSignatureCollection){
	runtime.SetFinalizer(digitalsignaturecollection, nil)
	C.Delete_CObject(C.CString("Delete_DigitalSignatureCollection"),digitalsignaturecollection.ptr)
	digitalsignaturecollection.ptr = nil
}
