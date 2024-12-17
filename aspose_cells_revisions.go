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

/**************Enum RevisionActionType *****************/

// Represents the type of revision action.
type RevisionActionType int32

const(
// Add revision.
RevisionActionType_Add RevisionActionType = 0 

// Delete revision.
RevisionActionType_Delete RevisionActionType = 1 

// Column delete revision.
RevisionActionType_DeleteColumn RevisionActionType = 2 

// Row delete revision.
RevisionActionType_DeleteRow RevisionActionType = 3 

// Column insert revision.
RevisionActionType_InsertColumn RevisionActionType = 4 

// Row insert revision.
RevisionActionType_InsertRow RevisionActionType = 5 
)

func Int32ToRevisionActionType(value int32)(RevisionActionType ,error){
	switch value {
		case 0:  return RevisionActionType_Add, nil  
		case 1:  return RevisionActionType_Delete, nil  
		case 2:  return RevisionActionType_DeleteColumn, nil  
		case 3:  return RevisionActionType_DeleteRow, nil  
		case 4:  return RevisionActionType_InsertColumn, nil  
		case 5:  return RevisionActionType_InsertRow, nil  
		default:
			return 0 ,fmt.Errorf("invalid RevisionActionType value: %d", value)
	}
}

/**************Enum RevisionType *****************/

// Represents the revision type.
type RevisionType int32

const(
// Custom view.
RevisionType_CustomView RevisionType = 0 

// Defined name.
RevisionType_DefinedName RevisionType = 1 

// Cells change.
RevisionType_ChangeCells RevisionType = 2 

// Auto format.
RevisionType_AutoFormat RevisionType = 3 

// Merge conflict.
RevisionType_MergeConflict RevisionType = 4 

// Comment.
RevisionType_Comment RevisionType = 5 

// Format.
RevisionType_Format RevisionType = 6 

// Insert worksheet.
RevisionType_InsertSheet RevisionType = 7 

// Move cells.
RevisionType_MoveCells RevisionType = 8 

// Undo.
RevisionType_Undo RevisionType = 9 

// Query table.
RevisionType_QueryTable RevisionType = 10 

// Inserting or deleting.
RevisionType_InsertDelete RevisionType = 11 

// Rename worksheet.
RevisionType_RenameSheet RevisionType = 12 

// Unknown.
RevisionType_Unknown RevisionType = 13 
)

func Int32ToRevisionType(value int32)(RevisionType ,error){
	switch value {
		case 0:  return RevisionType_CustomView, nil  
		case 1:  return RevisionType_DefinedName, nil  
		case 2:  return RevisionType_ChangeCells, nil  
		case 3:  return RevisionType_AutoFormat, nil  
		case 4:  return RevisionType_MergeConflict, nil  
		case 5:  return RevisionType_Comment, nil  
		case 6:  return RevisionType_Format, nil  
		case 7:  return RevisionType_InsertSheet, nil  
		case 8:  return RevisionType_MoveCells, nil  
		case 9:  return RevisionType_Undo, nil  
		case 10:  return RevisionType_QueryTable, nil  
		case 11:  return RevisionType_InsertDelete, nil  
		case 12:  return RevisionType_RenameSheet, nil  
		case 13:  return RevisionType_Unknown, nil  
		default:
			return 0 ,fmt.Errorf("invalid RevisionType value: %d", value)
	}
}
// Class HighlightChangesOptions 

// Represents options of highlighting revsions or changes of shared Excel files.
type HighlightChangesOptions struct {
	ptr unsafe.Pointer
}

// Represents options of highlighting revsions or changes of shared Excel files.
// Parameters:
//   highlightOnScreen - bool 
//   listOnNewSheet - bool 
func NewHighlightChangesOptions(highlightonscreen bool, listonnewsheet bool) ( *HighlightChangesOptions, error) {
	highlightchangesoptions := &HighlightChangesOptions{}
	CGoReturnPtr := C.New_HighlightChangesOptions(C.bool(highlightonscreen), C.bool(listonnewsheet))
	if CGoReturnPtr.error_no == 0 {
		highlightchangesoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(highlightchangesoptions, DeleteHighlightChangesOptions)
		return highlightchangesoptions, nil
	} else {
		highlightchangesoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return highlightchangesoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *HighlightChangesOptions) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.HighlightChangesOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}


func DeleteHighlightChangesOptions(highlightchangesoptions *HighlightChangesOptions){
	runtime.SetFinalizer(highlightchangesoptions, nil)
	C.Delete_HighlightChangesOptions(highlightchangesoptions.ptr)
	highlightchangesoptions.ptr = nil
}

// Class Revision 

// Represents the revision.
type Revision struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Revision) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.Revision_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *Revision) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.Revision_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *Revision) GetId()  (int32,  error)  {
	CGoReturnPtr := C.Revision_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the type of revision.
// Returns:
//   int32  
func (instance *Revision) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.Revision_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}


func DeleteRevision(revision *Revision){
	runtime.SetFinalizer(revision, nil)
	C.Delete_Revision(revision.ptr)
	revision.ptr = nil
}

// Class RevisionAutoFormat 

// represents a revision record of information about a formatting change.
type RevisionAutoFormat struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionAutoFormat(src *Revision) ( *RevisionAutoFormat, error) {
	revisionautoformat := &RevisionAutoFormat{}
	CGoReturnPtr := C.New_RevisionAutoFormat(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisionautoformat.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisionautoformat, DeleteRevisionAutoFormat)
		return revisionautoformat, nil
	} else {
		revisionautoformat.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisionautoformat, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionAutoFormat) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionAutoFormat_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the revision.
// Returns:
//   int32  
func (instance *RevisionAutoFormat) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionAutoFormat_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the location where the formatting was applied.
// Returns:
//   CellArea  
func (instance *RevisionAutoFormat) GetCellArea()  (*CellArea,  error)  {
	CGoReturnPtr := C.RevisionAutoFormat_GetCellArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionAutoFormat) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionAutoFormat_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionAutoFormat) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionAutoFormat_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionAutoFormat(revisionautoformat *RevisionAutoFormat){
	runtime.SetFinalizer(revisionautoformat, nil)
	C.Delete_RevisionAutoFormat(revisionautoformat.ptr)
	revisionautoformat.ptr = nil
}

// Class RevisionCellChange 

// Represents the revision that changing cells.
type RevisionCellChange struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionCellChange(src *Revision) ( *RevisionCellChange, error) {
	revisioncellchange := &RevisionCellChange{}
	CGoReturnPtr := C.New_RevisionCellChange(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisioncellchange.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisioncellchange, DeleteRevisionCellChange)
		return revisioncellchange, nil
	} else {
		revisioncellchange.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisioncellchange, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionCellChange) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCellChange_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of revision.
// Returns:
//   int32  
func (instance *RevisionCellChange) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the name of the cell.
// Returns:
//   string  
func (instance *RevisionCellChange) GetCellName()  (string,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetCellName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the row index of the cell.
// Returns:
//   int32  
func (instance *RevisionCellChange) GetRow()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the column index of the cell.
// Returns:
//   int32  
func (instance *RevisionCellChange) GetColumn()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this cell is new formatted.
// Returns:
//   bool  
func (instance *RevisionCellChange) IsNewFormatted()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCellChange_IsNewFormatted( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether this cell is old formatted.
// Returns:
//   bool  
func (instance *RevisionCellChange) IsOldFormatted()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCellChange_IsOldFormatted( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the old formula.
// Returns:
//   string  
func (instance *RevisionCellChange) GetOldFormula()  (string,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetOldFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets old value of the cell.
// Returns:
//   Object  
func (instance *RevisionCellChange) GetOldValue()  (*Object,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetOldValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Gets new value of the cell.
// Returns:
//   Object  
func (instance *RevisionCellChange) GetNewValue()  (*Object,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetNewValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Gets the old formula.
// Returns:
//   string  
func (instance *RevisionCellChange) GetNewFormula()  (string,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetNewFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the new style of the cell.
// Returns:
//   Style  
func (instance *RevisionCellChange) GetNewStyle()  (*Style,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetNewStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Gets the old style of the cell.
// Returns:
//   Style  
func (instance *RevisionCellChange) GetOldStyle()  (*Style,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetOldStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionCellChange) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionCellChange) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellChange_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionCellChange(revisioncellchange *RevisionCellChange){
	runtime.SetFinalizer(revisioncellchange, nil)
	C.Delete_RevisionCellChange(revisioncellchange.ptr)
	revisioncellchange.ptr = nil
}

// Class RevisionCellComment 

// Represents a revision record of a cell comment change.
type RevisionCellComment struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionCellComment(src *Revision) ( *RevisionCellComment, error) {
	revisioncellcomment := &RevisionCellComment{}
	CGoReturnPtr := C.New_RevisionCellComment(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisioncellcomment.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisioncellcomment, DeleteRevisionCellComment)
		return revisioncellcomment, nil
	} else {
		revisioncellcomment.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisioncellcomment, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionCellComment) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCellComment_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of revision.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the row index of the which contains a comment.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetRow()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the column index of the which contains a comment.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetColumn()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the cell.
// Returns:
//   string  
func (instance *RevisionCellComment) GetCellName()  (string,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetCellName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *RevisionCellComment) SetCellName(value string)  error {
	CGoReturnPtr := C.RevisionCellComment_SetCellName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the action type of the revision.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetActionType()  (RevisionActionType,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetActionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionActionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether it's an  old comment.
// Returns:
//   bool  
func (instance *RevisionCellComment) IsOldComment()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCellComment_IsOldComment( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets Length of the comment text added in this revision.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetOldLength()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetOldLength( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets Length of the comment before this revision was made.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetNewLength()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetNewLength( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionCellComment) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionCellComment) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellComment_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionCellComment(revisioncellcomment *RevisionCellComment){
	runtime.SetFinalizer(revisioncellcomment, nil)
	C.Delete_RevisionCellComment(revisioncellcomment.ptr)
	revisioncellcomment.ptr = nil
}

// Class RevisionCellMove 

// Represents a revision record on a cell(s) that moved.
type RevisionCellMove struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionCellMove(src *Revision) ( *RevisionCellMove, error) {
	revisioncellmove := &RevisionCellMove{}
	CGoReturnPtr := C.New_RevisionCellMove(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisioncellmove.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisioncellmove, DeleteRevisionCellMove)
		return revisioncellmove, nil
	} else {
		revisioncellmove.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisioncellmove, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionCellMove) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCellMove_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of revision.
// Returns:
//   int32  
func (instance *RevisionCellMove) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionCellMove_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the source area.
// Returns:
//   CellArea  
func (instance *RevisionCellMove) GetSourceArea()  (*CellArea,  error)  {
	CGoReturnPtr := C.RevisionCellMove_GetSourceArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets the destination area.
// Returns:
//   CellArea  
func (instance *RevisionCellMove) GetDestinationArea()  (*CellArea,  error)  {
	CGoReturnPtr := C.RevisionCellMove_GetDestinationArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets the source worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionCellMove) GetSourceWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionCellMove_GetSourceWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionCellMove) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionCellMove_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionCellMove) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCellMove_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionCellMove(revisioncellmove *RevisionCellMove){
	runtime.SetFinalizer(revisioncellmove, nil)
	C.Delete_RevisionCellMove(revisioncellmove.ptr)
	revisioncellmove.ptr = nil
}

// Class RevisionCollection 

// Represents all revision logs.
type RevisionCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets <see cref="Revision"/> by the index.
// Parameters:
//   index - int32 
// Returns:
//   Revision  
func (instance *RevisionCollection) Get(index int32)  (*Revision,  error)  {
	CGoReturnPtr := C.RevisionCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Revision{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRevision) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *RevisionCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionCollection(revisioncollection *RevisionCollection){
	runtime.SetFinalizer(revisioncollection, nil)
	C.Delete_RevisionCollection(revisioncollection.ptr)
	revisioncollection.ptr = nil
}

// Class RevisionCustomView 

// Represents a revision record of adding or removing a custom view to the workbook
type RevisionCustomView struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionCustomView(src *Revision) ( *RevisionCustomView, error) {
	revisioncustomview := &RevisionCustomView{}
	CGoReturnPtr := C.New_RevisionCustomView(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisioncustomview.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisioncustomview, DeleteRevisionCustomView)
		return revisioncustomview, nil
	} else {
		revisioncustomview.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisioncustomview, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionCustomView) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionCustomView_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of revision.
// Returns:
//   int32  
func (instance *RevisionCustomView) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionCustomView_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the type of action.
// Returns:
//   int32  
func (instance *RevisionCustomView) GetActionType()  (RevisionActionType,  error)  {
	CGoReturnPtr := C.RevisionCustomView_GetActionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionActionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the globally unique identifier of the custom view.
// Parameters:
//   uuid - UUID 
// Returns:
//   void  
func (instance *RevisionCustomView) GetGuid(uuid *UUID)  error {
	CGoReturnPtr := C.RevisionCustomView_GetGuid( instance.ptr, uuid.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionCustomView) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionCustomView_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionCustomView) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionCustomView_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionCustomView(revisioncustomview *RevisionCustomView){
	runtime.SetFinalizer(revisioncustomview, nil)
	C.Delete_RevisionCustomView(revisioncustomview.ptr)
	revisioncustomview.ptr = nil
}

// Class RevisionDefinedName 

// Represents a revision record of a defined name change.
type RevisionDefinedName struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionDefinedName(src *Revision) ( *RevisionDefinedName, error) {
	revisiondefinedname := &RevisionDefinedName{}
	CGoReturnPtr := C.New_RevisionDefinedName(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisiondefinedname.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisiondefinedname, DeleteRevisionDefinedName)
		return revisiondefinedname, nil
	} else {
		revisiondefinedname.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisiondefinedname, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionDefinedName) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of revision.
// Returns:
//   int32  
func (instance *RevisionDefinedName) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the text of the defined name.
// Returns:
//   string  
func (instance *RevisionDefinedName) GetText()  (string,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_GetText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the old formula.
// Returns:
//   string  
func (instance *RevisionDefinedName) GetOldFormula()  (string,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_GetOldFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the formula.
// Returns:
//   string  
func (instance *RevisionDefinedName) GetNewFormula()  (string,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_GetNewFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionDefinedName) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionDefinedName) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionDefinedName_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionDefinedName(revisiondefinedname *RevisionDefinedName){
	runtime.SetFinalizer(revisiondefinedname, nil)
	C.Delete_RevisionDefinedName(revisiondefinedname.ptr)
	revisiondefinedname.ptr = nil
}

// Class RevisionFormat 

// Represents a revision record of information about a formatting change.
type RevisionFormat struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionFormat(src *Revision) ( *RevisionFormat, error) {
	revisionformat := &RevisionFormat{}
	CGoReturnPtr := C.New_RevisionFormat(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisionformat.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisionformat, DeleteRevisionFormat)
		return revisionformat, nil
	} else {
		revisionformat.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisionformat, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionFormat) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionFormat_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of revision.
// Returns:
//   int32  
func (instance *RevisionFormat) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionFormat_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the applied style.
// Returns:
//   Style  
func (instance *RevisionFormat) GetStyle()  (*Style,  error)  {
	CGoReturnPtr := C.RevisionFormat_GetStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionFormat) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionFormat_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionFormat) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionFormat_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionFormat(revisionformat *RevisionFormat){
	runtime.SetFinalizer(revisionformat, nil)
	C.Delete_RevisionFormat(revisionformat.ptr)
	revisionformat.ptr = nil
}

// Class RevisionHeader 

// Represents a list of specific changes that have taken place for this workbook.
type RevisionHeader struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewRevisionHeader() ( *RevisionHeader, error) {
	revisionheader := &RevisionHeader{}
	CGoReturnPtr := C.New_RevisionHeader()
	if CGoReturnPtr.error_no == 0 {
		revisionheader.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisionheader, DeleteRevisionHeader)
		return revisionheader, nil
	} else {
		revisionheader.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisionheader, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionHeader) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionHeader_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets rhe date and time when this set of revisions was saved.
// Returns:
//   Date  
func (instance *RevisionHeader) GetSavedTime()  (*Date,  error)  {
	CGoReturnPtr := C.RevisionHeader_GetSavedTime( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets and sets rhe date and time when this set of revisions was saved.
// Parameters:
//   value - Date 
// Returns:
//   void  
func (instance *RevisionHeader) SetSavedTime(value *Date)  error {
	CGoReturnPtr := C.RevisionHeader_SetSavedTime( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the name of the user making the revision.
// Returns:
//   string  
func (instance *RevisionHeader) Get_UserName()  (string,  error)  {
	CGoReturnPtr := C.RevisionHeader_Get_UserName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the user making the revision.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *RevisionHeader) SetUserName(value string)  error {
	CGoReturnPtr := C.RevisionHeader_SetUserName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteRevisionHeader(revisionheader *RevisionHeader){
	runtime.SetFinalizer(revisionheader, nil)
	C.Delete_RevisionHeader(revisionheader.ptr)
	revisionheader.ptr = nil
}

// Class RevisionInsertDelete 

// Represents a revision record of a row/column insert/delete action.
type RevisionInsertDelete struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionInsertDelete(src *Revision) ( *RevisionInsertDelete, error) {
	revisioninsertdelete := &RevisionInsertDelete{}
	CGoReturnPtr := C.New_RevisionInsertDelete(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisioninsertdelete.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisioninsertdelete, DeleteRevisionInsertDelete)
		return revisioninsertdelete, nil
	} else {
		revisioninsertdelete.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisioninsertdelete, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionInsertDelete) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of revision.
// Returns:
//   int32  
func (instance *RevisionInsertDelete) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the inserting/deleting range.
// Returns:
//   CellArea  
func (instance *RevisionInsertDelete) GetCellArea()  (*CellArea,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_GetCellArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets the action type of this revision.
// Returns:
//   int32  
func (instance *RevisionInsertDelete) GetActionType()  (RevisionActionType,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_GetActionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionActionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets revision list by this operation.
// Returns:
//   RevisionCollection  
func (instance *RevisionInsertDelete) GetRevisions()  (*RevisionCollection,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_GetRevisions( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &RevisionCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRevisionCollection) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionInsertDelete) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionInsertDelete) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionInsertDelete_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionInsertDelete(revisioninsertdelete *RevisionInsertDelete){
	runtime.SetFinalizer(revisioninsertdelete, nil)
	C.Delete_RevisionInsertDelete(revisioninsertdelete.ptr)
	revisioninsertdelete.ptr = nil
}

// Class RevisionInsertSheet 

// Represents a revision record of a sheet that was inserted.
type RevisionInsertSheet struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionInsertSheet(src *Revision) ( *RevisionInsertSheet, error) {
	revisioninsertsheet := &RevisionInsertSheet{}
	CGoReturnPtr := C.New_RevisionInsertSheet(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisioninsertsheet.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisioninsertsheet, DeleteRevisionInsertSheet)
		return revisioninsertsheet, nil
	} else {
		revisioninsertsheet.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisioninsertsheet, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionInsertSheet) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of revision.
// Returns:
//   int32  
func (instance *RevisionInsertSheet) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the action type of the revision.
// Returns:
//   int32  
func (instance *RevisionInsertSheet) GetActionType()  (RevisionActionType,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_GetActionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionActionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the name of the worksheet.
// Returns:
//   string  
func (instance *RevisionInsertSheet) GetName()  (string,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the zero based position of the new sheet in the sheet tab bar.
// Returns:
//   int32  
func (instance *RevisionInsertSheet) GetSheetPosition()  (int32,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_GetSheetPosition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionInsertSheet) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionInsertSheet) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionInsertSheet_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionInsertSheet(revisioninsertsheet *RevisionInsertSheet){
	runtime.SetFinalizer(revisioninsertsheet, nil)
	C.Delete_RevisionInsertSheet(revisioninsertsheet.ptr)
	revisioninsertsheet.ptr = nil
}

// Class RevisionLog 

// Represents the revision log.
type RevisionLog struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionLog) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionLog_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets table that contains metadata about a list of specific changes that have taken place
// for this workbook.
// Returns:
//   RevisionHeader  
func (instance *RevisionLog) GetMetadataTable()  (*RevisionHeader,  error)  {
	CGoReturnPtr := C.RevisionLog_GetMetadataTable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &RevisionHeader{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRevisionHeader) 

	return result, nil 
}
// Gets all revisions in this log.
// Returns:
//   RevisionCollection  
func (instance *RevisionLog) GetRevisions()  (*RevisionCollection,  error)  {
	CGoReturnPtr := C.RevisionLog_GetRevisions( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &RevisionCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRevisionCollection) 

	return result, nil 
}


func DeleteRevisionLog(revisionlog *RevisionLog){
	runtime.SetFinalizer(revisionlog, nil)
	C.Delete_RevisionLog(revisionlog.ptr)
	revisionlog.ptr = nil
}

// Class RevisionLogCollection 

// Represents all revision logs.
type RevisionLogCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionLogCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionLogCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the number of days the spreadsheet application will keep the change history for this workbook.
// Returns:
//   int32  
func (instance *RevisionLogCollection) GetDaysPreservingHistory()  (int32,  error)  {
	CGoReturnPtr := C.RevisionLogCollection_GetDaysPreservingHistory( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the number of days the spreadsheet application will keep the change history for this workbook.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RevisionLogCollection) SetDaysPreservingHistory(value int32)  error {
	CGoReturnPtr := C.RevisionLogCollection_SetDaysPreservingHistory( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets <see cref="RevisionLog"/> by index.
// Parameters:
//   index - int32 
// Returns:
//   RevisionLog  
func (instance *RevisionLogCollection) Get(index int32)  (*RevisionLog,  error)  {
	CGoReturnPtr := C.RevisionLogCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &RevisionLog{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRevisionLog) 

	return result, nil 
}
// Highlights changes of shared workbook.
// Parameters:
//   options - HighlightChangesOptions 
// Returns:
//   void  
func (instance *RevisionLogCollection) HighlightChanges(options *HighlightChangesOptions)  error {
	CGoReturnPtr := C.RevisionLogCollection_HighlightChanges( instance.ptr, options.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *RevisionLogCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.RevisionLogCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionLogCollection(revisionlogcollection *RevisionLogCollection){
	runtime.SetFinalizer(revisionlogcollection, nil)
	C.Delete_RevisionLogCollection(revisionlogcollection.ptr)
	revisionlogcollection.ptr = nil
}

// Class RevisionMergeConflict 

// Represents a revision record which indicates that there was a merge conflict.
type RevisionMergeConflict struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionMergeConflict(src *Revision) ( *RevisionMergeConflict, error) {
	revisionmergeconflict := &RevisionMergeConflict{}
	CGoReturnPtr := C.New_RevisionMergeConflict(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisionmergeconflict.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisionmergeconflict, DeleteRevisionMergeConflict)
		return revisionmergeconflict, nil
	} else {
		revisionmergeconflict.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisionmergeconflict, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionMergeConflict) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionMergeConflict_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of revision.
// Returns:
//   int32  
func (instance *RevisionMergeConflict) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionMergeConflict_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionMergeConflict) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionMergeConflict_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionMergeConflict) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionMergeConflict_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionMergeConflict(revisionmergeconflict *RevisionMergeConflict){
	runtime.SetFinalizer(revisionmergeconflict, nil)
	C.Delete_RevisionMergeConflict(revisionmergeconflict.ptr)
	revisionmergeconflict.ptr = nil
}

// Class RevisionQueryTable 

// Represents a revision of a query table field change.
type RevisionQueryTable struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionQueryTable(src *Revision) ( *RevisionQueryTable, error) {
	revisionquerytable := &RevisionQueryTable{}
	CGoReturnPtr := C.New_RevisionQueryTable(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisionquerytable.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisionquerytable, DeleteRevisionQueryTable)
		return revisionquerytable, nil
	} else {
		revisionquerytable.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisionquerytable, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionQueryTable) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionQueryTable_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of the revision.
// Returns:
//   int32  
func (instance *RevisionQueryTable) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionQueryTable_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the location of the affected query table.
// Returns:
//   CellArea  
func (instance *RevisionQueryTable) GetCellArea()  (*CellArea,  error)  {
	CGoReturnPtr := C.RevisionQueryTable_GetCellArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Gets ID of the specific query table field that was removed.
// Returns:
//   int32  
func (instance *RevisionQueryTable) GetFieldId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionQueryTable_GetFieldId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionQueryTable) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionQueryTable_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionQueryTable) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionQueryTable_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionQueryTable(revisionquerytable *RevisionQueryTable){
	runtime.SetFinalizer(revisionquerytable, nil)
	C.Delete_RevisionQueryTable(revisionquerytable.ptr)
	revisionquerytable.ptr = nil
}

// Class RevisionRenameSheet 

// Represents a revision of renaming sheet.
type RevisionRenameSheet struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Revision 
func NewRevisionRenameSheet(src *Revision) ( *RevisionRenameSheet, error) {
	revisionrenamesheet := &RevisionRenameSheet{}
	CGoReturnPtr := C.New_RevisionRenameSheet(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		revisionrenamesheet.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(revisionrenamesheet, DeleteRevisionRenameSheet)
		return revisionrenamesheet, nil
	} else {
		revisionrenamesheet.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return revisionrenamesheet, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RevisionRenameSheet) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RevisionRenameSheet_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of the revision.
// Returns:
//   int32  
func (instance *RevisionRenameSheet) GetType()  (RevisionType,  error)  {
	CGoReturnPtr := C.RevisionRenameSheet_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToRevisionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the old name of the worksheet.
// Returns:
//   string  
func (instance *RevisionRenameSheet) GetOldName()  (string,  error)  {
	CGoReturnPtr := C.RevisionRenameSheet_GetOldName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the new name of the worksheet.
// Returns:
//   string  
func (instance *RevisionRenameSheet) GetNewName()  (string,  error)  {
	CGoReturnPtr := C.RevisionRenameSheet_GetNewName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the worksheet.
// Returns:
//   Worksheet  
func (instance *RevisionRenameSheet) GetWorksheet()  (*Worksheet,  error)  {
	CGoReturnPtr := C.RevisionRenameSheet_GetWorksheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Worksheet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorksheet) 

	return result, nil 
}
// Gets the number of this revision.
// Returns:
//   int32  
func (instance *RevisionRenameSheet) GetId()  (int32,  error)  {
	CGoReturnPtr := C.RevisionRenameSheet_GetId( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteRevisionRenameSheet(revisionrenamesheet *RevisionRenameSheet){
	runtime.SetFinalizer(revisionrenamesheet, nil)
	C.Delete_RevisionRenameSheet(revisionrenamesheet.ptr)
	revisionrenamesheet.ptr = nil
}
