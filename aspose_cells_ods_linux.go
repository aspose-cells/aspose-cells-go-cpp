// +build linux

// Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
// Powered by Aspose.Cells.
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

/**************Enum OdsCellFieldType *****************/

// Represents the cell field type of ods.
type OdsCellFieldType int32

const(
// Current date.
OdsCellFieldType_Date OdsCellFieldType = 0 

// The name of the sheet.
OdsCellFieldType_SheetName OdsCellFieldType = 1 

// The name of the file.
OdsCellFieldType_Title OdsCellFieldType = 2 
)

func Int32ToOdsCellFieldType(value int32)(OdsCellFieldType ,error){
	switch value {
		case 0:  return OdsCellFieldType_Date, nil  
		case 1:  return OdsCellFieldType_SheetName, nil  
		case 2:  return OdsCellFieldType_Title, nil  
		default:
			return 0 ,fmt.Errorf("invalid OdsCellFieldType value: %d", value)
	}
}

/**************Enum OdsGeneratorType *****************/

// Represents the type of ODS generator.
type OdsGeneratorType int32

const(
// Libre Office
OdsGeneratorType_LibreOffice OdsGeneratorType = 0 

// Open Office
OdsGeneratorType_OpenOffice OdsGeneratorType = 1 
)

func Int32ToOdsGeneratorType(value int32)(OdsGeneratorType ,error){
	switch value {
		case 0:  return OdsGeneratorType_LibreOffice, nil  
		case 1:  return OdsGeneratorType_OpenOffice, nil  
		default:
			return 0 ,fmt.Errorf("invalid OdsGeneratorType value: %d", value)
	}
}

/**************Enum OdsPageBackgroundGraphicPositionType *****************/

// Represents the position.
type OdsPageBackgroundGraphicPositionType int32

const(
// Top left.
OdsPageBackgroundGraphicPositionType_TopLeft OdsPageBackgroundGraphicPositionType = 0 

// Top center.
OdsPageBackgroundGraphicPositionType_TopCenter OdsPageBackgroundGraphicPositionType = 1 

// Top right.
OdsPageBackgroundGraphicPositionType_TopRight OdsPageBackgroundGraphicPositionType = 2 

// Center left.
OdsPageBackgroundGraphicPositionType_CenterLeft OdsPageBackgroundGraphicPositionType = 3 

// Center.
OdsPageBackgroundGraphicPositionType_CenterCenter OdsPageBackgroundGraphicPositionType = 4 

// Center right.
OdsPageBackgroundGraphicPositionType_CenterRight OdsPageBackgroundGraphicPositionType = 5 

// Bottom left.
OdsPageBackgroundGraphicPositionType_BottomLeft OdsPageBackgroundGraphicPositionType = 6 

// Bottom center.
OdsPageBackgroundGraphicPositionType_BottomCenter OdsPageBackgroundGraphicPositionType = 7 

// Bottom right.
OdsPageBackgroundGraphicPositionType_BottomRight OdsPageBackgroundGraphicPositionType = 8 
)

func Int32ToOdsPageBackgroundGraphicPositionType(value int32)(OdsPageBackgroundGraphicPositionType ,error){
	switch value {
		case 0:  return OdsPageBackgroundGraphicPositionType_TopLeft, nil  
		case 1:  return OdsPageBackgroundGraphicPositionType_TopCenter, nil  
		case 2:  return OdsPageBackgroundGraphicPositionType_TopRight, nil  
		case 3:  return OdsPageBackgroundGraphicPositionType_CenterLeft, nil  
		case 4:  return OdsPageBackgroundGraphicPositionType_CenterCenter, nil  
		case 5:  return OdsPageBackgroundGraphicPositionType_CenterRight, nil  
		case 6:  return OdsPageBackgroundGraphicPositionType_BottomLeft, nil  
		case 7:  return OdsPageBackgroundGraphicPositionType_BottomCenter, nil  
		case 8:  return OdsPageBackgroundGraphicPositionType_BottomRight, nil  
		default:
			return 0 ,fmt.Errorf("invalid OdsPageBackgroundGraphicPositionType value: %d", value)
	}
}

/**************Enum OdsPageBackgroundGraphicType *****************/

// Represents the type of formatting page background with image.
type OdsPageBackgroundGraphicType int32

const(
// Set the image at specific position.
OdsPageBackgroundGraphicType_Position OdsPageBackgroundGraphicType = 0 

// Stretch the image.
OdsPageBackgroundGraphicType_Area OdsPageBackgroundGraphicType = 1 

// Repeat and repeat the image.
OdsPageBackgroundGraphicType_Tile OdsPageBackgroundGraphicType = 2 
)

func Int32ToOdsPageBackgroundGraphicType(value int32)(OdsPageBackgroundGraphicType ,error){
	switch value {
		case 0:  return OdsPageBackgroundGraphicType_Position, nil  
		case 1:  return OdsPageBackgroundGraphicType_Area, nil  
		case 2:  return OdsPageBackgroundGraphicType_Tile, nil  
		default:
			return 0 ,fmt.Errorf("invalid OdsPageBackgroundGraphicType value: %d", value)
	}
}

/**************Enum OdsPageBackgroundType *****************/

// Represents the page background type of ods.
type OdsPageBackgroundType int32

const(
// No background.
OdsPageBackgroundType_None OdsPageBackgroundType = 0 

// Formats the background with color.
OdsPageBackgroundType_Color OdsPageBackgroundType = 1 

// Formats the background with image.
OdsPageBackgroundType_Graphic OdsPageBackgroundType = 2 
)

func Int32ToOdsPageBackgroundType(value int32)(OdsPageBackgroundType ,error){
	switch value {
		case 0:  return OdsPageBackgroundType_None, nil  
		case 1:  return OdsPageBackgroundType_Color, nil  
		case 2:  return OdsPageBackgroundType_Graphic, nil  
		default:
			return 0 ,fmt.Errorf("invalid OdsPageBackgroundType value: %d", value)
	}
}

/**************Enum OpenDocumentFormatVersionType *****************/

// Open Document Format version type.
type OpenDocumentFormatVersionType int32

const(
// None strict.
OpenDocumentFormatVersionType_None OpenDocumentFormatVersionType = 0 

// ODF Version 1.1
OpenDocumentFormatVersionType_Odf11 OpenDocumentFormatVersionType = 1 

// ODF Version 1.2
OpenDocumentFormatVersionType_Odf12 OpenDocumentFormatVersionType = 2 

// ODF Version 1.3
OpenDocumentFormatVersionType_Odf13 OpenDocumentFormatVersionType = 3 
)

func Int32ToOpenDocumentFormatVersionType(value int32)(OpenDocumentFormatVersionType ,error){
	switch value {
		case 0:  return OpenDocumentFormatVersionType_None, nil  
		case 1:  return OpenDocumentFormatVersionType_Odf11, nil  
		case 2:  return OpenDocumentFormatVersionType_Odf12, nil  
		case 3:  return OpenDocumentFormatVersionType_Odf13, nil  
		default:
			return 0 ,fmt.Errorf("invalid OpenDocumentFormatVersionType value: %d", value)
	}
}
// Class OdsCellField 

// Represents the cell field of ods.
type OdsCellField struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *OdsCellField) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.OdsCellField_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the custom format of the field's value.
// Returns:
//   string  
func (instance *OdsCellField) GetCustomFormat()  (string,  error)  {
	
	CGoReturnPtr := C.OdsCellField_GetCustomFormat( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the custom format of the field's value.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *OdsCellField) SetCustomFormat(value string)  error {
	
	CGoReturnPtr := C.OdsCellField_SetCustomFormat( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of the field.
// Returns:
//   int32  
func (instance *OdsCellField) GetFieldType()  (OdsCellFieldType,  error)  {
	
	CGoReturnPtr := C.OdsCellField_GetFieldType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOdsCellFieldType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of the field.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *OdsCellField) SetFieldType(value OdsCellFieldType)  error {
	
	CGoReturnPtr := C.OdsCellField_SetFieldType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get and sets the row index of the cell.
// Returns:
//   int32  
func (instance *OdsCellField) GetRow()  (int32,  error)  {
	
	CGoReturnPtr := C.OdsCellField_GetRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Get and sets the row index of the cell.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *OdsCellField) SetRow(value int32)  error {
	
	CGoReturnPtr := C.OdsCellField_SetRow( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get and sets the column index of the cell.
// Returns:
//   int32  
func (instance *OdsCellField) GetColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.OdsCellField_GetColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Get and sets the column index of the cell.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *OdsCellField) SetColumn(value int32)  error {
	
	CGoReturnPtr := C.OdsCellField_SetColumn( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteOdsCellField(odscellfield *OdsCellField){
	runtime.SetFinalizer(odscellfield, nil)
	C.Delete_OdsCellField(odscellfield.ptr)
	odscellfield.ptr = nil
}

// Class OdsCellFieldCollection 

// Represents the fields of ODS.
type OdsCellFieldCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *OdsCellFieldCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.OdsCellFieldCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the field by the index.
// Parameters:
//   index - int32 
// Returns:
//   OdsCellField  
func (instance *OdsCellFieldCollection) Get_Int(index int32)  (*OdsCellField,  error)  {
	
	CGoReturnPtr := C.OdsCellFieldCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &OdsCellField{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteOdsCellField) 

	return result, nil 
}
// Gets the field by row and column index.
// Parameters:
//   row - int32 
//   column - int32 
// Returns:
//   OdsCellField  
func (instance *OdsCellFieldCollection) Get_Int_Int(row int32, column int32)  (*OdsCellField,  error)  {
	
	CGoReturnPtr := C.OdsCellFieldCollection_Get_Integer_Integer( instance.ptr, C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &OdsCellField{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteOdsCellField) 

	return result, nil 
}
// Adds a field.
// Parameters:
//   row - int32 
//   column - int32 
//   fieldType - int32 
//   format - string 
// Returns:
//   int32  
func (instance *OdsCellFieldCollection) Add(row int32, column int32, fieldtype OdsCellFieldType, format string)  (int32,  error)  {
	
	CGoReturnPtr := C.OdsCellFieldCollection_Add( instance.ptr, C.int(row), C.int(column), C.int( int32(fieldtype)), C.CString(format))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Update fields value to the cells.
// Returns:
//   void  
func (instance *OdsCellFieldCollection) UpdateFieldsValue()  error {
	
	CGoReturnPtr := C.OdsCellFieldCollection_UpdateFieldsValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *OdsCellFieldCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.OdsCellFieldCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteOdsCellFieldCollection(odscellfieldcollection *OdsCellFieldCollection){
	runtime.SetFinalizer(odscellfieldcollection, nil)
	C.Delete_OdsCellFieldCollection(odscellfieldcollection.ptr)
	odscellfieldcollection.ptr = nil
}

// Class OdsPageBackground 

// Represents the page background of ods.
type OdsPageBackground struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewOdsPageBackground() ( *OdsPageBackground, error) {
	odspagebackground := &OdsPageBackground{}
	CGoReturnPtr := C.New_OdsPageBackground()
	if CGoReturnPtr.error_no == 0 {
		odspagebackground.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(odspagebackground, DeleteOdsPageBackground)
		return odspagebackground, nil
	} else {
		odspagebackground.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return odspagebackground, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *OdsPageBackground) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the page background type.
// Returns:
//   int32  
func (instance *OdsPageBackground) GetType()  (OdsPageBackgroundType,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOdsPageBackgroundType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the page background type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *OdsPageBackground) SetType(value OdsPageBackgroundType)  error {
	
	CGoReturnPtr := C.OdsPageBackground_SetType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color of background.
// Returns:
//   Color  
func (instance *OdsPageBackground) GetColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_GetColor( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets and sets the color of background.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *OdsPageBackground) SetColor(value *Color)  error {
	
	CGoReturnPtr := C.OdsPageBackground_SetColor( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the page background graphic type.
// Returns:
//   int32  
func (instance *OdsPageBackground) GetGraphicType()  (OdsPageBackgroundGraphicType,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_GetGraphicType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOdsPageBackgroundGraphicType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the page background graphic type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *OdsPageBackground) SetGraphicType(value OdsPageBackgroundGraphicType)  error {
	
	CGoReturnPtr := C.OdsPageBackground_SetGraphicType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the background graphic position.
// Returns:
//   int32  
func (instance *OdsPageBackground) GetGraphicPositionType()  (OdsPageBackgroundGraphicPositionType,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_GetGraphicPositionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOdsPageBackgroundGraphicPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the background graphic position.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *OdsPageBackground) SetGraphicPositionType(value OdsPageBackgroundGraphicPositionType)  error {
	
	CGoReturnPtr := C.OdsPageBackground_SetGraphicPositionType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether it's a linked graphic.
// Returns:
//   bool  
func (instance *OdsPageBackground) IsLink()  (bool,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_IsLink( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the linked graphic path.
// Returns:
//   string  
func (instance *OdsPageBackground) GetLinkedGraphic()  (string,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_GetLinkedGraphic( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked graphic path.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *OdsPageBackground) SetLinkedGraphic(value string)  error {
	
	CGoReturnPtr := C.OdsPageBackground_SetLinkedGraphic( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the graphic data.
// Returns:
//   []byte  
func (instance *OdsPageBackground) GetGraphicData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.OdsPageBackground_GetGraphicData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the graphic data.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *OdsPageBackground) SetGraphicData(value []byte)  error {
	
	CGoReturnPtr := C.OdsPageBackground_SetGraphicData( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteOdsPageBackground(odspagebackground *OdsPageBackground){
	runtime.SetFinalizer(odspagebackground, nil)
	C.Delete_OdsPageBackground(odspagebackground.ptr)
	odspagebackground.ptr = nil
}
