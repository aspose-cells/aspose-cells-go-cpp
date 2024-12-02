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

/**************Enum PropertyType *****************/
type PropertyType int32

const(
// The property is a boolean value.
PropertyType_Boolean PropertyType = 0 

// The property is a date time value.
PropertyType_DateTime PropertyType = 1 

// The property is a floating number.
PropertyType_Double PropertyType = 2 

// The property is an integer number.
PropertyType_Number PropertyType = 3 

// The property is a string value.
PropertyType_String PropertyType = 4 

// The property is a byte array.
PropertyType_Blob PropertyType = 5 
)

func Int32ToPropertyType(value int32)(PropertyType ,error){
	switch value {
		case 0:  return PropertyType_Boolean, nil  
		case 1:  return PropertyType_DateTime, nil  
		case 2:  return PropertyType_Double, nil  
		case 3:  return PropertyType_Number, nil  
		case 4:  return PropertyType_String, nil  
		case 5:  return PropertyType_Blob, nil  
		default:
			return 0 ,fmt.Errorf("invalid PropertyType value: %d", value)
	}
}
// Class BuiltInDocumentPropertyCollection 

// A collection of built-in document properties.
type BuiltInDocumentPropertyCollection struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - DocumentPropertyCollection 
func NewBuiltInDocumentPropertyCollection(src *DocumentPropertyCollection) ( *BuiltInDocumentPropertyCollection, error) {
	builtindocumentpropertycollection := &BuiltInDocumentPropertyCollection{}
	CGoReturnPtr := C.New_BuiltInDocumentPropertyCollection(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		builtindocumentpropertycollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(builtindocumentpropertycollection, DeleteBuiltInDocumentPropertyCollection)
		return builtindocumentpropertycollection, nil
	} else {
		builtindocumentpropertycollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return builtindocumentpropertycollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *BuiltInDocumentPropertyCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Returns a <see cref="DocumentProperty"/> object by the name of the property.
// Parameters:
//   name - string 
// Returns:
//   DocumentProperty  
func (instance *BuiltInDocumentPropertyCollection) Get_String(name string)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Returns a <see cref="DocumentProperty"/> object by index.
// Parameters:
//   index - int32 
// Returns:
//   DocumentProperty  
func (instance *BuiltInDocumentPropertyCollection) Get_Int(index int32)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Gets or sets the document's language.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetLanguage()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLanguage( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the document's language.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLanguage(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLanguage( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the name of the document's author.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetAuthor()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetAuthor( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the name of the document's author.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetAuthor(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetAuthor( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the document comments.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetComments()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetComments( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the document comments.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetComments(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetComments( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the category of the document.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetCategory()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetCategory( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the category of the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetCategory(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetCategory( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the content type of the document.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetContentType()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetContentType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the content type of the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetContentType(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetContentType( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the content status of the document.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetContentStatus()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetContentStatus( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the content status of the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetContentStatus(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetContentStatus( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the company property.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetCompany()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetCompany( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the company property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetCompany(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetCompany( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the hyperlinkbase property.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetHyperlinkBase()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetHyperlinkBase( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the hyperlinkbase property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetHyperlinkBase(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetHyperlinkBase( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets date of the document creation in local timezone.
// Returns:
//   Date  
func (instance *BuiltInDocumentPropertyCollection) GetCreatedTime()  (*Date,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetCreatedTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets date of the document creation in local timezone.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetCreatedTime(value *Date)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetCreatedTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the Universal time of the document creation.
// Returns:
//   Date  
func (instance *BuiltInDocumentPropertyCollection) GetCreatedUniversalTime()  (*Date,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetCreatedUniversalTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the Universal time of the document creation.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetCreatedUniversalTime(value *Date)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetCreatedUniversalTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the document keywords.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetKeywords()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetKeywords( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the document keywords.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetKeywords(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetKeywords( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the date when the document was last printed in local timezone.
// Returns:
//   Date  
func (instance *BuiltInDocumentPropertyCollection) GetLastPrinted()  (*Date,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLastPrinted( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the date when the document was last printed in local timezone.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLastPrinted(value *Date)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLastPrinted( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the Universal time when the document was last printed.
// Returns:
//   Date  
func (instance *BuiltInDocumentPropertyCollection) GetLastPrintedUniversalTime()  (*Date,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLastPrintedUniversalTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the Universal time when the document was last printed.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLastPrintedUniversalTime(value *Date)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLastPrintedUniversalTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the name of the last author.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetLastSavedBy()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLastSavedBy( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the name of the last author.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLastSavedBy(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLastSavedBy( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the time of the last save in local timezone.
// Returns:
//   Date  
func (instance *BuiltInDocumentPropertyCollection) GetLastSavedTime()  (*Date,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLastSavedTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the time of the last save in local timezone.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLastSavedTime(value *Date)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLastSavedTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the universal time of the last save.
// Returns:
//   Date  
func (instance *BuiltInDocumentPropertyCollection) GetLastSavedUniversalTime()  (*Date,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLastSavedUniversalTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the universal time of the last save.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLastSavedUniversalTime(value *Date)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLastSavedUniversalTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the manager property.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetManager()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetManager( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the manager property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetManager(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetManager( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the name of the application.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetNameOfApplication()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetNameOfApplication( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the name of the application.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetNameOfApplication(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetNameOfApplication( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents an estimate of the number of pages in the document.
// Returns:
//   int32  
func (instance *BuiltInDocumentPropertyCollection) GetPages()  (int32,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetPages( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents an estimate of the number of pages in the document.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetPages(value int32)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetPages( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the document revision number.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetRevisionNumber()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetRevisionNumber( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the document revision number.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetRevisionNumber(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetRevisionNumber( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the subject of the document.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetSubject()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetSubject( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the subject of the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetSubject(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetSubject( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the informational name of the document template.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetTemplate()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetTemplate( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the informational name of the document template.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetTemplate(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetTemplate( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the title of the document.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetTitle()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetTitle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the title of the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetTitle(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetTitle( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the total editing time in minutes.
// Returns:
//   float64  
func (instance *BuiltInDocumentPropertyCollection) GetTotalEditingTime()  (float64,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetTotalEditingTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the total editing time in minutes.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetTotalEditingTime(value float64)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetTotalEditingTime( instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the version number of the application that created the document.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetVersion()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetVersion( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the version number of the application that created the document.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetVersion(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetVersion( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the version of the file.
// Returns:
//   string  
func (instance *BuiltInDocumentPropertyCollection) GetDocumentVersion()  (string,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetDocumentVersion( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the version of the file.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetDocumentVersion(value string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetDocumentVersion( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates the display mode of the document thumbnail.
// Returns:
//   bool  
func (instance *BuiltInDocumentPropertyCollection) GetScaleCrop()  (bool,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetScaleCrop( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates the display mode of the document thumbnail.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetScaleCrop(value bool)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetScaleCrop( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether hyperlinks in a document are up-to-date.
// Returns:
//   bool  
func (instance *BuiltInDocumentPropertyCollection) GetLinksUpToDate()  (bool,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetLinksUpToDate( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether hyperlinks in a document are up-to-date.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetLinksUpToDate(value bool)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetLinksUpToDate( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents an estimate of the number of words in the document.
// Returns:
//   int32  
func (instance *BuiltInDocumentPropertyCollection) GetWords()  (int32,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetWords( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents an estimate of the number of words in the document.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) SetWords(value int32)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_SetWords( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns true if a property with the specified name exists in the collection.
// Parameters:
//   name - string 
// Returns:
//   bool  
func (instance *BuiltInDocumentPropertyCollection) Contains(name string)  (bool,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_Contains( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the index of a property by name.
// Parameters:
//   name - string 
// Returns:
//   int32  
func (instance *BuiltInDocumentPropertyCollection) IndexOf(name string)  (int32,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_IndexOf( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Removes a property with the specified name from the collection.
// Parameters:
//   name - string 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) Remove(name string)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_Remove( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes a property at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *BuiltInDocumentPropertyCollection) RemoveAt(index int32)  error {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_RemoveAt( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *BuiltInDocumentPropertyCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.BuiltInDocumentPropertyCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteBuiltInDocumentPropertyCollection(builtindocumentpropertycollection *BuiltInDocumentPropertyCollection){
	runtime.SetFinalizer(builtindocumentpropertycollection, nil)
	C.Delete_BuiltInDocumentPropertyCollection(builtindocumentpropertycollection.ptr)
	builtindocumentpropertycollection.ptr = nil
}

// Class ContentTypeProperty 

// Represents identifier information.
type ContentTypeProperty struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ContentTypeProperty) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ContentTypeProperty_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Returns or sets the name of the object.
// Returns:
//   string  
func (instance *ContentTypeProperty) GetName()  (string,  error)  {
	CGoReturnPtr := C.ContentTypeProperty_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the name of the object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ContentTypeProperty) SetName(value string)  error {
	CGoReturnPtr := C.ContentTypeProperty_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the value of the content type property.
// Returns:
//   string  
func (instance *ContentTypeProperty) GetValue()  (string,  error)  {
	CGoReturnPtr := C.ContentTypeProperty_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the value of the content type property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ContentTypeProperty) SetValue(value string)  error {
	CGoReturnPtr := C.ContentTypeProperty_SetValue( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of the property.
// Returns:
//   string  
func (instance *ContentTypeProperty) GetType()  (string,  error)  {
	CGoReturnPtr := C.ContentTypeProperty_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the type of the property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ContentTypeProperty) SetType(value string)  error {
	CGoReturnPtr := C.ContentTypeProperty_SetType( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the value could be empty.
// Returns:
//   bool  
func (instance *ContentTypeProperty) IsNillable()  (bool,  error)  {
	CGoReturnPtr := C.ContentTypeProperty_IsNillable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the value could be empty.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ContentTypeProperty) SetIsNillable(value bool)  error {
	CGoReturnPtr := C.ContentTypeProperty_SetIsNillable( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteContentTypeProperty(contenttypeproperty *ContentTypeProperty){
	runtime.SetFinalizer(contenttypeproperty, nil)
	C.Delete_ContentTypeProperty(contenttypeproperty.ptr)
	contenttypeproperty.ptr = nil
}

// Class ContentTypePropertyCollection 

// A collection of <see cref="ContentTypeProperty"/> objects that represent additional information.
type ContentTypePropertyCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ContentTypePropertyCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ContentTypePropertyCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Adds content type property information.
// Parameters:
//   name - string 
//   value - string 
// Returns:
//   int32  
func (instance *ContentTypePropertyCollection) Add_String_String(name string, value string)  (int32,  error)  {
	CGoReturnPtr := C.ContentTypePropertyCollection_Add_String_String( instance.ptr, C.CString(name), C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds content type property information.
// Parameters:
//   name - string 
//   value - string 
//   type - string 
// Returns:
//   int32  
func (instance *ContentTypePropertyCollection) Add_String_String_String(name string, value string, type_ string)  (int32,  error)  {
	CGoReturnPtr := C.ContentTypePropertyCollection_Add_String_String_String( instance.ptr, C.CString(name), C.CString(value), C.CString(type_))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the content type property by the specific index.
// Parameters:
//   index - int32 
// Returns:
//   ContentTypeProperty  
func (instance *ContentTypePropertyCollection) Get_Int(index int32)  (*ContentTypeProperty,  error)  {
	CGoReturnPtr := C.ContentTypePropertyCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ContentTypeProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteContentTypeProperty) 

	return result, nil 
}
// Gets the content type property by the property name.
// Parameters:
//   name - string 
// Returns:
//   ContentTypeProperty  
func (instance *ContentTypePropertyCollection) Get_String(name string)  (*ContentTypeProperty,  error)  {
	CGoReturnPtr := C.ContentTypePropertyCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ContentTypeProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteContentTypeProperty) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *ContentTypePropertyCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.ContentTypePropertyCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteContentTypePropertyCollection(contenttypepropertycollection *ContentTypePropertyCollection){
	runtime.SetFinalizer(contenttypepropertycollection, nil)
	C.Delete_ContentTypePropertyCollection(contenttypepropertycollection.ptr)
	contenttypepropertycollection.ptr = nil
}

// Class CustomDocumentPropertyCollection 

// A collection of custom document properties.
type CustomDocumentPropertyCollection struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - DocumentPropertyCollection 
func NewCustomDocumentPropertyCollection(src *DocumentPropertyCollection) ( *CustomDocumentPropertyCollection, error) {
	customdocumentpropertycollection := &CustomDocumentPropertyCollection{}
	CGoReturnPtr := C.New_CustomDocumentPropertyCollection(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		customdocumentpropertycollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(customdocumentpropertycollection, DeleteCustomDocumentPropertyCollection)
		return customdocumentpropertycollection, nil
	} else {
		customdocumentpropertycollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return customdocumentpropertycollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CustomDocumentPropertyCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Creates a new custom document property of the <b>PropertyType.String</b> data type.
// Parameters:
//   name - string 
//   value - string 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Add_String_String(name string, value string)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Add_String_String( instance.ptr, C.CString(name), C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Creates a new custom document property of the <b>PropertyType.Number</b> data type.
// Parameters:
//   name - string 
//   value - int32 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Add_String_Int(name string, value int32)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Add_String_Integer( instance.ptr, C.CString(name), C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Creates a new custom document property of the <b>PropertyType.DateTime</b> data type.
// Parameters:
//   name - string 
//   value - Date 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Add_String_Date(name string, value *Date)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Add_String_Date( instance.ptr, C.CString(name), value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Creates a new custom document property of the <b>PropertyType.Boolean</b> data type.
// Parameters:
//   name - string 
//   value - bool 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Add_String_Bool(name string, value bool)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Add_String_Boolean( instance.ptr, C.CString(name), C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Creates a new custom document property of the <b>PropertyType.Float</b> data type.
// Parameters:
//   name - string 
//   value - float64 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Add_String_Double(name string, value float64)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Add_String_Double( instance.ptr, C.CString(name), C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Creates a new custom document property which links to content.
// Parameters:
//   name - string 
//   source - string 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) AddLinkToContent(name string, source string)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_AddLinkToContent( instance.ptr, C.CString(name), C.CString(source))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Update custom document property value which links to content.
// Returns:
//   void  
func (instance *CustomDocumentPropertyCollection) UpdateLinkedPropertyValue()  error {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_UpdateLinkedPropertyValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Update custom document property value to linked range.
// Returns:
//   void  
func (instance *CustomDocumentPropertyCollection) UpdateLinkedRange()  error {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_UpdateLinkedRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a <see cref="DocumentProperty"/> object by index.
// Parameters:
//   index - int32 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Get_Int(index int32)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Returns true if a property with the specified name exists in the collection.
// Parameters:
//   name - string 
// Returns:
//   bool  
func (instance *CustomDocumentPropertyCollection) Contains(name string)  (bool,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Contains( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the index of a property by name.
// Parameters:
//   name - string 
// Returns:
//   int32  
func (instance *CustomDocumentPropertyCollection) IndexOf(name string)  (int32,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_IndexOf( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Removes a property with the specified name from the collection.
// Parameters:
//   name - string 
// Returns:
//   void  
func (instance *CustomDocumentPropertyCollection) Remove(name string)  error {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Remove( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes a property at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *CustomDocumentPropertyCollection) RemoveAt(index int32)  error {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_RemoveAt( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a <see cref="DocumentProperty"/> object by the name of the property.
// Parameters:
//   name - string 
// Returns:
//   DocumentProperty  
func (instance *CustomDocumentPropertyCollection) Get_String(name string)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *CustomDocumentPropertyCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.CustomDocumentPropertyCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteCustomDocumentPropertyCollection(customdocumentpropertycollection *CustomDocumentPropertyCollection){
	runtime.SetFinalizer(customdocumentpropertycollection, nil)
	C.Delete_CustomDocumentPropertyCollection(customdocumentpropertycollection.ptr)
	customdocumentpropertycollection.ptr = nil
}

// Class CustomProperty 

// Represents identifier information.
type CustomProperty struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewCustomProperty() ( *CustomProperty, error) {
	customproperty := &CustomProperty{}
	CGoReturnPtr := C.New_CustomProperty()
	if CGoReturnPtr.error_no == 0 {
		customproperty.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(customproperty, DeleteCustomProperty)
		return customproperty, nil
	} else {
		customproperty.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return customproperty, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CustomProperty) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.CustomProperty_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Returns or sets the name of the object.
// Returns:
//   string  
func (instance *CustomProperty) GetName()  (string,  error)  {
	CGoReturnPtr := C.CustomProperty_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the name of the object.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CustomProperty) SetName(value string)  error {
	CGoReturnPtr := C.CustomProperty_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the value of the custom property.
// Returns:
//   string  
func (instance *CustomProperty) GetValue()  (string,  error)  {
	CGoReturnPtr := C.CustomProperty_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the value of the custom property.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CustomProperty) SetValue(value string)  error {
	CGoReturnPtr := C.CustomProperty_SetValue( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteCustomProperty(customproperty *CustomProperty){
	runtime.SetFinalizer(customproperty, nil)
	C.Delete_CustomProperty(customproperty.ptr)
	customproperty.ptr = nil
}

// Class CustomPropertyCollection 

// A collection of <see cref="CustomProperty"/> objects that represent additional information.
type CustomPropertyCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CustomPropertyCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.CustomPropertyCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Adds custom property information.
// Parameters:
//   name - string 
//   value - string 
// Returns:
//   int32  
func (instance *CustomPropertyCollection) Add(name string, value string)  (int32,  error)  {
	CGoReturnPtr := C.CustomPropertyCollection_Add( instance.ptr, C.CString(name), C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the custom property by the specific index.
// Parameters:
//   index - int32 
// Returns:
//   CustomProperty  
func (instance *CustomPropertyCollection) Get_Int(index int32)  (*CustomProperty,  error)  {
	CGoReturnPtr := C.CustomPropertyCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CustomProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCustomProperty) 

	return result, nil 
}
// Gets the custom property by the property name.
// Parameters:
//   name - string 
// Returns:
//   CustomProperty  
func (instance *CustomPropertyCollection) Get_String(name string)  (*CustomProperty,  error)  {
	CGoReturnPtr := C.CustomPropertyCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CustomProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCustomProperty) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *CustomPropertyCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.CustomPropertyCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteCustomPropertyCollection(custompropertycollection *CustomPropertyCollection){
	runtime.SetFinalizer(custompropertycollection, nil)
	C.Delete_CustomPropertyCollection(custompropertycollection.ptr)
	custompropertycollection.ptr = nil
}

// Class DocumentProperty 

// Represents a custom or built-in document property.
type DocumentProperty struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DocumentProperty) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.DocumentProperty_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Returns the name of the property.
// Returns:
//   string  
func (instance *DocumentProperty) GetName()  (string,  error)  {
	CGoReturnPtr := C.DocumentProperty_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the value of the property.
// Returns:
//   Object  
func (instance *DocumentProperty) GetValue()  (*Object,  error)  {
	CGoReturnPtr := C.DocumentProperty_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Gets or sets the value of the property.
// Parameters:
//   value - Object 
// Returns:
//   void  
func (instance *DocumentProperty) SetValue(value *Object)  error {
	CGoReturnPtr := C.DocumentProperty_SetValue( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this property is linked to content
// Returns:
//   bool  
func (instance *DocumentProperty) IsLinkedToContent()  (bool,  error)  {
	CGoReturnPtr := C.DocumentProperty_IsLinkedToContent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// The linked content source.
// Returns:
//   string  
func (instance *DocumentProperty) GetSource()  (string,  error)  {
	CGoReturnPtr := C.DocumentProperty_GetSource( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the data type of the property.
// Returns:
//   int32  
func (instance *DocumentProperty) GetType()  (PropertyType,  error)  {
	CGoReturnPtr := C.DocumentProperty_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPropertyType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Returns true if this property does not have a name in the OLE2 storage
// and a unique name was generated only for the public API.
// Returns:
//   bool  
func (instance *DocumentProperty) IsGeneratedName()  (bool,  error)  {
	CGoReturnPtr := C.DocumentProperty_IsGeneratedName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Returns the property value as a string.
// Returns:
//   string  
func (instance *DocumentProperty) ToString()  (string,  error)  {
	CGoReturnPtr := C.DocumentProperty_ToString( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the property value as integer.
// Returns:
//   int32  
func (instance *DocumentProperty) ToInt()  (int32,  error)  {
	CGoReturnPtr := C.DocumentProperty_ToInt( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the property value as double.
// Returns:
//   float64  
func (instance *DocumentProperty) ToDouble()  (float64,  error)  {
	CGoReturnPtr := C.DocumentProperty_ToDouble( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the property value as DateTime in local timezone.
// Returns:
//   Date  
func (instance *DocumentProperty) ToDateTime()  (*Date,  error)  {
	CGoReturnPtr := C.DocumentProperty_ToDateTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Returns the property value as bool.
// Returns:
//   bool  
func (instance *DocumentProperty) ToBool()  (bool,  error)  {
	CGoReturnPtr := C.DocumentProperty_ToBool( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}


func DeleteDocumentProperty(documentproperty *DocumentProperty){
	runtime.SetFinalizer(documentproperty, nil)
	C.Delete_DocumentProperty(documentproperty.ptr)
	documentproperty.ptr = nil
}

// Class DocumentPropertyCollection 

// Base class for <see cref="BuiltInDocumentPropertyCollection"/> and <see cref="CustomDocumentPropertyCollection"/> collections.
type DocumentPropertyCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DocumentPropertyCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.DocumentPropertyCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Returns a <see cref="DocumentProperty"/> object by index.
// Parameters:
//   index - int32 
// Returns:
//   DocumentProperty  
func (instance *DocumentPropertyCollection) Get_Int(index int32)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.DocumentPropertyCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Returns true if a property with the specified name exists in the collection.
// Parameters:
//   name - string 
// Returns:
//   bool  
func (instance *DocumentPropertyCollection) Contains(name string)  (bool,  error)  {
	CGoReturnPtr := C.DocumentPropertyCollection_Contains( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the index of a property by name.
// Parameters:
//   name - string 
// Returns:
//   int32  
func (instance *DocumentPropertyCollection) IndexOf(name string)  (int32,  error)  {
	CGoReturnPtr := C.DocumentPropertyCollection_IndexOf( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Removes a property with the specified name from the collection.
// Parameters:
//   name - string 
// Returns:
//   void  
func (instance *DocumentPropertyCollection) Remove(name string)  error {
	CGoReturnPtr := C.DocumentPropertyCollection_Remove( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes a property at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *DocumentPropertyCollection) RemoveAt(index int32)  error {
	CGoReturnPtr := C.DocumentPropertyCollection_RemoveAt( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a <see cref="DocumentProperty"/> object by the name of the property.
// Parameters:
//   name - string 
// Returns:
//   DocumentProperty  
func (instance *DocumentPropertyCollection) Get_String(name string)  (*DocumentProperty,  error)  {
	CGoReturnPtr := C.DocumentPropertyCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DocumentProperty{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDocumentProperty) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *DocumentPropertyCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.DocumentPropertyCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteDocumentPropertyCollection(documentpropertycollection *DocumentPropertyCollection){
	runtime.SetFinalizer(documentpropertycollection, nil)
	C.Delete_DocumentPropertyCollection(documentpropertycollection.ptr)
	documentpropertycollection.ptr = nil
}
