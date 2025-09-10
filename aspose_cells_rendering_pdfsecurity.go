// +build windows

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/win_x86_64" -L"${SRCDIR}/lib/win_x86_64" -lAspose.Cells.CWrapper
// #include <CellsFunctionMap.h>
import "C"
import (
 
 	
	"errors"	
	"runtime"
	"unsafe" 
)

// Class PdfSecurityOptions 

// Options for encrypting and access permissions for a PDF document.
// PDF/A does not allow security setting.
type PdfSecurityOptions struct {
	ptr unsafe.Pointer
}

// The constructor of PdfSecurityOptions
func NewPdfSecurityOptions() ( *PdfSecurityOptions, error) {
	pdfsecurityoptions := &PdfSecurityOptions{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_PdfSecurityOptions"),)
	if CGoReturnPtr.error_no == 0 {
		pdfsecurityoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pdfsecurityoptions, DeletePdfSecurityOptions)
		return pdfsecurityoptions, nil
	} else {
		pdfsecurityoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pdfsecurityoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the user password required for opening the encrypted PDF document.
// Returns:
//   string  
func (instance *PdfSecurityOptions) GetUserPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("PdfSecurityOptions_GetUserPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the user password required for opening the encrypted PDF document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetUserPassword(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("PdfSecurityOptions_SetUserPassword"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the owner password for the encrypted PDF document.
// Returns:
//   string  
func (instance *PdfSecurityOptions) GetOwnerPassword()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZO(C.CString("PdfSecurityOptions_GetOwnerPassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the owner password for the encrypted PDF document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetOwnerPassword(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("PdfSecurityOptions_SetOwnerPassword"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to print the document.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetPrintPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetPrintPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to print the document.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetPrintPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetPrintPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to modify the contents of the document by operations other than those controlled
// by <see cref="AnnotationsPermission"/>, <see cref="FillFormsPermission"/> and <see cref="AssembleDocumentPermission"/>.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetModifyDocumentPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetModifyDocumentPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to modify the contents of the document by operations other than those controlled
// by <see cref="AnnotationsPermission"/>, <see cref="FillFormsPermission"/> and <see cref="AssembleDocumentPermission"/>.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetModifyDocumentPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetModifyDocumentPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to add or modify text annotations, fill in interactive form fields.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetAnnotationsPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetAnnotationsPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to add or modify text annotations, fill in interactive form fields.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetAnnotationsPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetAnnotationsPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to fill in existing interactive form fields (including signature fields),
// even if <see cref="ModifyDocumentPermission"/> is clear.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetFillFormsPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetFillFormsPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to fill in existing interactive form fields (including signature fields),
// even if <see cref="ModifyDocumentPermission"/> is clear.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetFillFormsPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetFillFormsPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to copy or otherwise extract text and graphics from the document
// by operations other than that controlled by <see cref="AccessibilityExtractContent"/>.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetExtractContentPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetExtractContentPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to copy or otherwise extract text and graphics from the document
// by operations other than that controlled by <see cref="AccessibilityExtractContent"/>.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetExtractContentPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetExtractContentPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to extract text and graphics (in support of accessibility to users with disabilities or for other purposes).
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetAccessibilityExtractContent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetAccessibilityExtractContent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to extract text and graphics (in support of accessibility to users with disabilities or for other purposes).
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetAccessibilityExtractContent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetAccessibilityExtractContent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images),
// even if <see cref="ModifyDocumentPermission"/> is clear.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetAssembleDocumentPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetAssembleDocumentPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images),
// even if <see cref="ModifyDocumentPermission"/> is clear.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetAssembleDocumentPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetAssembleDocumentPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to allow to print the document to a representation from
// which a faithful digital copy of the PDF content could be generated.
// Returns:
//   bool  
func (instance *PdfSecurityOptions) GetFullQualityPrintPermission()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PdfSecurityOptions_GetFullQualityPrintPermission"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to allow to print the document to a representation from
// which a faithful digital copy of the PDF content could be generated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfSecurityOptions) SetFullQualityPrintPermission(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("PdfSecurityOptions_SetFullQualityPrintPermission"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeletePdfSecurityOptions(pdfsecurityoptions *PdfSecurityOptions){
	runtime.SetFinalizer(pdfsecurityoptions, nil)
	C.Delete_CObject(C.CString("Delete_PdfSecurityOptions"),pdfsecurityoptions.ptr)
	pdfsecurityoptions.ptr = nil
}
