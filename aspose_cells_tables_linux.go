// +build linux

// Copyright (c) 2001-2024 Aspose Pty Ltd. All Rights Reserved.
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

/**************Enum TableDataSourceType *****************/
// Represents the table's data source type.
type TableDataSourceType int32

const(
// Excel Worksheet Table
TableDataSourceType_Worksheet TableDataSourceType = 0 

// Read-write SharePoint linked List
TableDataSourceType_SharePoint TableDataSourceType = 1 

// XML mapper Table
TableDataSourceType_XML TableDataSourceType = 2 

// Query Table
TableDataSourceType_QueryTable TableDataSourceType = 3 
)

func Int32ToTableDataSourceType(value int32)(TableDataSourceType ,error){
	switch value {
		case 0:  return TableDataSourceType_Worksheet, nil  
		case 1:  return TableDataSourceType_SharePoint, nil  
		case 2:  return TableDataSourceType_XML, nil  
		case 3:  return TableDataSourceType_QueryTable, nil  
		default:
			return 0 ,fmt.Errorf("invalid TableDataSourceType value: %d", value)
	}
}

/**************Enum TableStyleElementType *****************/
// Represents the Table or PivotTable style element type.
type TableStyleElementType int32

const(
// Table style element that applies to PivotTable's blank rows.
TableStyleElementType_BlankRow TableStyleElementType = 18 

// Table style element that applies to table's first column.
TableStyleElementType_FirstColumn TableStyleElementType = 8 

// Table style element that applies to table's first column stripes.
TableStyleElementType_FirstColumnStripe TableStyleElementType = 3 

// Table style element that applies to PivotTable's first column subheading.
TableStyleElementType_FirstColumnSubheading TableStyleElementType = 22 

// Table style element that applies to table's first header row cell.
TableStyleElementType_FirstHeaderCell TableStyleElementType = 11 

// Table style element that applies to table's first row stripes.
TableStyleElementType_FirstRowStripe TableStyleElementType = 5 

// Table style element that applies to PivotTable's first row subheading.
TableStyleElementType_FirstRowSubheading TableStyleElementType = 25 

// Table style element that applies to PivotTable's first subtotal column.
TableStyleElementType_FirstSubtotalColumn TableStyleElementType = 15 

// Table style element that applies to pivot table's first subtotal row.
TableStyleElementType_FirstSubtotalRow TableStyleElementType = 19 

// Table style element that applies to pivot table's grand total column.
TableStyleElementType_GrandTotalColumn TableStyleElementType = 28 

// Table style element that applies to pivot table's grand total row.
TableStyleElementType_GrandTotalRow TableStyleElementType = 29 

// Table style element that applies to table's first total row cell.
TableStyleElementType_FirstTotalCell TableStyleElementType = 13 

// Table style element that applies to table's header row.
TableStyleElementType_HeaderRow TableStyleElementType = 9 

// Table style element that applies to table's last column.
TableStyleElementType_LastColumn TableStyleElementType = 7 

// Table style element that applies to table's last header row cell.
TableStyleElementType_LastHeaderCell TableStyleElementType = 12 

// Table style element that applies to table's last total row cell.
TableStyleElementType_LastTotalCell TableStyleElementType = 14 

// Table style element that applies to pivot table's page field labels.
TableStyleElementType_PageFieldLabels TableStyleElementType = 1 

// Table style element that applies to pivot table's page field values.
TableStyleElementType_PageFieldValues TableStyleElementType = 2 

// Table style element that applies to table's second column stripes.
TableStyleElementType_SecondColumnStripe TableStyleElementType = 4 

// Table style element that applies to pivot table's second column subheading.
TableStyleElementType_SecondColumnSubheading TableStyleElementType = 23 

// Table style element that applies to table's second row stripes.
TableStyleElementType_SecondRowStripe TableStyleElementType = 6 

// Table style element that applies to pivot table's second row subheading.
TableStyleElementType_SecondRowSubheading TableStyleElementType = 26 

// Table style element that applies to PivotTable's second subtotal column.
TableStyleElementType_SecondSubtotalColumn TableStyleElementType = 16 

// Table style element that applies to PivotTable's second subtotal row.
TableStyleElementType_SecondSubtotalRow TableStyleElementType = 20 

// Table style element that applies to PivotTable's third column subheading.
TableStyleElementType_ThirdColumnSubheading TableStyleElementType = 24 

// Table style element that applies to PivotTable's third row subheading.
TableStyleElementType_ThirdRowSubheading TableStyleElementType = 27 

// Table style element that applies to pivot table's third subtotal column.
TableStyleElementType_ThirdSubtotalColumn TableStyleElementType = 17 

// Table style element that applies to PivotTable's third subtotal row.
TableStyleElementType_ThirdSubtotalRow TableStyleElementType = 21 

// Table style element that applies to table's total row.
TableStyleElementType_TotalRow TableStyleElementType = 10 

// Table style element that applies to table's entire content.
TableStyleElementType_WholeTable TableStyleElementType = 0 
)

func Int32ToTableStyleElementType(value int32)(TableStyleElementType ,error){
	switch value {
		case 18:  return TableStyleElementType_BlankRow, nil  
		case 8:  return TableStyleElementType_FirstColumn, nil  
		case 3:  return TableStyleElementType_FirstColumnStripe, nil  
		case 22:  return TableStyleElementType_FirstColumnSubheading, nil  
		case 11:  return TableStyleElementType_FirstHeaderCell, nil  
		case 5:  return TableStyleElementType_FirstRowStripe, nil  
		case 25:  return TableStyleElementType_FirstRowSubheading, nil  
		case 15:  return TableStyleElementType_FirstSubtotalColumn, nil  
		case 19:  return TableStyleElementType_FirstSubtotalRow, nil  
		case 28:  return TableStyleElementType_GrandTotalColumn, nil  
		case 29:  return TableStyleElementType_GrandTotalRow, nil  
		case 13:  return TableStyleElementType_FirstTotalCell, nil  
		case 9:  return TableStyleElementType_HeaderRow, nil  
		case 7:  return TableStyleElementType_LastColumn, nil  
		case 12:  return TableStyleElementType_LastHeaderCell, nil  
		case 14:  return TableStyleElementType_LastTotalCell, nil  
		case 1:  return TableStyleElementType_PageFieldLabels, nil  
		case 2:  return TableStyleElementType_PageFieldValues, nil  
		case 4:  return TableStyleElementType_SecondColumnStripe, nil  
		case 23:  return TableStyleElementType_SecondColumnSubheading, nil  
		case 6:  return TableStyleElementType_SecondRowStripe, nil  
		case 26:  return TableStyleElementType_SecondRowSubheading, nil  
		case 16:  return TableStyleElementType_SecondSubtotalColumn, nil  
		case 20:  return TableStyleElementType_SecondSubtotalRow, nil  
		case 24:  return TableStyleElementType_ThirdColumnSubheading, nil  
		case 27:  return TableStyleElementType_ThirdRowSubheading, nil  
		case 17:  return TableStyleElementType_ThirdSubtotalColumn, nil  
		case 21:  return TableStyleElementType_ThirdSubtotalRow, nil  
		case 10:  return TableStyleElementType_TotalRow, nil  
		case 0:  return TableStyleElementType_WholeTable, nil  
		default:
			return 0 ,fmt.Errorf("invalid TableStyleElementType value: %d", value)
	}
}

/**************Enum TableStyleType *****************/
// Represents the built-in table style type.
type TableStyleType int32

const(
TableStyleType_None TableStyleType = 0 


TableStyleType_TableStyleLight1 TableStyleType = 1 


TableStyleType_TableStyleLight2 TableStyleType = 2 


TableStyleType_TableStyleLight3 TableStyleType = 3 


TableStyleType_TableStyleLight4 TableStyleType = 4 


TableStyleType_TableStyleLight5 TableStyleType = 5 


TableStyleType_TableStyleLight6 TableStyleType = 6 


TableStyleType_TableStyleLight7 TableStyleType = 7 


TableStyleType_TableStyleLight8 TableStyleType = 8 


TableStyleType_TableStyleLight9 TableStyleType = 9 


TableStyleType_TableStyleLight10 TableStyleType = 10 


TableStyleType_TableStyleLight11 TableStyleType = 11 


TableStyleType_TableStyleLight12 TableStyleType = 12 


TableStyleType_TableStyleLight13 TableStyleType = 13 


TableStyleType_TableStyleLight14 TableStyleType = 14 


TableStyleType_TableStyleLight15 TableStyleType = 15 


TableStyleType_TableStyleLight16 TableStyleType = 16 


TableStyleType_TableStyleLight17 TableStyleType = 17 


TableStyleType_TableStyleLight18 TableStyleType = 18 


TableStyleType_TableStyleLight19 TableStyleType = 19 


TableStyleType_TableStyleLight20 TableStyleType = 20 


TableStyleType_TableStyleLight21 TableStyleType = 21 


TableStyleType_TableStyleMedium1 TableStyleType = 22 


TableStyleType_TableStyleMedium2 TableStyleType = 23 


TableStyleType_TableStyleMedium3 TableStyleType = 24 


TableStyleType_TableStyleMedium4 TableStyleType = 25 


TableStyleType_TableStyleMedium5 TableStyleType = 26 


TableStyleType_TableStyleMedium6 TableStyleType = 27 


TableStyleType_TableStyleMedium7 TableStyleType = 28 


TableStyleType_TableStyleMedium8 TableStyleType = 29 


TableStyleType_TableStyleMedium9 TableStyleType = 30 


TableStyleType_TableStyleMedium10 TableStyleType = 31 


TableStyleType_TableStyleMedium11 TableStyleType = 32 


TableStyleType_TableStyleMedium12 TableStyleType = 33 


TableStyleType_TableStyleMedium13 TableStyleType = 34 


TableStyleType_TableStyleMedium14 TableStyleType = 35 


TableStyleType_TableStyleMedium15 TableStyleType = 36 


TableStyleType_TableStyleMedium16 TableStyleType = 37 


TableStyleType_TableStyleMedium17 TableStyleType = 38 


TableStyleType_TableStyleMedium18 TableStyleType = 39 


TableStyleType_TableStyleMedium19 TableStyleType = 40 


TableStyleType_TableStyleMedium20 TableStyleType = 41 


TableStyleType_TableStyleMedium21 TableStyleType = 42 


TableStyleType_TableStyleMedium22 TableStyleType = 43 


TableStyleType_TableStyleMedium23 TableStyleType = 44 


TableStyleType_TableStyleMedium24 TableStyleType = 45 


TableStyleType_TableStyleMedium25 TableStyleType = 46 


TableStyleType_TableStyleMedium26 TableStyleType = 47 


TableStyleType_TableStyleMedium27 TableStyleType = 48 


TableStyleType_TableStyleMedium28 TableStyleType = 49 


TableStyleType_TableStyleDark1 TableStyleType = 50 


TableStyleType_TableStyleDark2 TableStyleType = 51 


TableStyleType_TableStyleDark3 TableStyleType = 52 


TableStyleType_TableStyleDark4 TableStyleType = 53 


TableStyleType_TableStyleDark5 TableStyleType = 54 


TableStyleType_TableStyleDark6 TableStyleType = 55 


TableStyleType_TableStyleDark7 TableStyleType = 56 


TableStyleType_TableStyleDark8 TableStyleType = 57 


TableStyleType_TableStyleDark9 TableStyleType = 58 


TableStyleType_TableStyleDark10 TableStyleType = 59 


TableStyleType_TableStyleDark11 TableStyleType = 60 


TableStyleType_Custom TableStyleType = 61 
)

func Int32ToTableStyleType(value int32)(TableStyleType ,error){
	switch value {
		case 0:  return TableStyleType_None, nil  
		case 1:  return TableStyleType_TableStyleLight1, nil  
		case 2:  return TableStyleType_TableStyleLight2, nil  
		case 3:  return TableStyleType_TableStyleLight3, nil  
		case 4:  return TableStyleType_TableStyleLight4, nil  
		case 5:  return TableStyleType_TableStyleLight5, nil  
		case 6:  return TableStyleType_TableStyleLight6, nil  
		case 7:  return TableStyleType_TableStyleLight7, nil  
		case 8:  return TableStyleType_TableStyleLight8, nil  
		case 9:  return TableStyleType_TableStyleLight9, nil  
		case 10:  return TableStyleType_TableStyleLight10, nil  
		case 11:  return TableStyleType_TableStyleLight11, nil  
		case 12:  return TableStyleType_TableStyleLight12, nil  
		case 13:  return TableStyleType_TableStyleLight13, nil  
		case 14:  return TableStyleType_TableStyleLight14, nil  
		case 15:  return TableStyleType_TableStyleLight15, nil  
		case 16:  return TableStyleType_TableStyleLight16, nil  
		case 17:  return TableStyleType_TableStyleLight17, nil  
		case 18:  return TableStyleType_TableStyleLight18, nil  
		case 19:  return TableStyleType_TableStyleLight19, nil  
		case 20:  return TableStyleType_TableStyleLight20, nil  
		case 21:  return TableStyleType_TableStyleLight21, nil  
		case 22:  return TableStyleType_TableStyleMedium1, nil  
		case 23:  return TableStyleType_TableStyleMedium2, nil  
		case 24:  return TableStyleType_TableStyleMedium3, nil  
		case 25:  return TableStyleType_TableStyleMedium4, nil  
		case 26:  return TableStyleType_TableStyleMedium5, nil  
		case 27:  return TableStyleType_TableStyleMedium6, nil  
		case 28:  return TableStyleType_TableStyleMedium7, nil  
		case 29:  return TableStyleType_TableStyleMedium8, nil  
		case 30:  return TableStyleType_TableStyleMedium9, nil  
		case 31:  return TableStyleType_TableStyleMedium10, nil  
		case 32:  return TableStyleType_TableStyleMedium11, nil  
		case 33:  return TableStyleType_TableStyleMedium12, nil  
		case 34:  return TableStyleType_TableStyleMedium13, nil  
		case 35:  return TableStyleType_TableStyleMedium14, nil  
		case 36:  return TableStyleType_TableStyleMedium15, nil  
		case 37:  return TableStyleType_TableStyleMedium16, nil  
		case 38:  return TableStyleType_TableStyleMedium17, nil  
		case 39:  return TableStyleType_TableStyleMedium18, nil  
		case 40:  return TableStyleType_TableStyleMedium19, nil  
		case 41:  return TableStyleType_TableStyleMedium20, nil  
		case 42:  return TableStyleType_TableStyleMedium21, nil  
		case 43:  return TableStyleType_TableStyleMedium22, nil  
		case 44:  return TableStyleType_TableStyleMedium23, nil  
		case 45:  return TableStyleType_TableStyleMedium24, nil  
		case 46:  return TableStyleType_TableStyleMedium25, nil  
		case 47:  return TableStyleType_TableStyleMedium26, nil  
		case 48:  return TableStyleType_TableStyleMedium27, nil  
		case 49:  return TableStyleType_TableStyleMedium28, nil  
		case 50:  return TableStyleType_TableStyleDark1, nil  
		case 51:  return TableStyleType_TableStyleDark2, nil  
		case 52:  return TableStyleType_TableStyleDark3, nil  
		case 53:  return TableStyleType_TableStyleDark4, nil  
		case 54:  return TableStyleType_TableStyleDark5, nil  
		case 55:  return TableStyleType_TableStyleDark6, nil  
		case 56:  return TableStyleType_TableStyleDark7, nil  
		case 57:  return TableStyleType_TableStyleDark8, nil  
		case 58:  return TableStyleType_TableStyleDark9, nil  
		case 59:  return TableStyleType_TableStyleDark10, nil  
		case 60:  return TableStyleType_TableStyleDark11, nil  
		case 61:  return TableStyleType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid TableStyleType value: %d", value)
	}
}

/**************Enum TotalsCalculation *****************/
// Determines the type of calculation in the Totals row of the list column.
type TotalsCalculation int32

const(
// Represents Sum totals calculation.
TotalsCalculation_Sum TotalsCalculation = 6 

// Represents Count totals calculation.
TotalsCalculation_Count TotalsCalculation = 2 

// Represents Average totals calculation.
TotalsCalculation_Average TotalsCalculation = 1 

// Represents Max totals calculation.
TotalsCalculation_Max TotalsCalculation = 4 

// Represents Min totals calculation.
TotalsCalculation_Min TotalsCalculation = 5 

// Represents Var totals calculation.
TotalsCalculation_Var TotalsCalculation = 8 

// Represents Count Nums totals calculation.
TotalsCalculation_CountNums TotalsCalculation = 3 

// Represents StdDev totals calculation.
TotalsCalculation_StdDev TotalsCalculation = 7 

// Represents No totals calculation.
TotalsCalculation_None TotalsCalculation = 0 

// Represents custom calculation.
TotalsCalculation_Custom TotalsCalculation = 9 
)

func Int32ToTotalsCalculation(value int32)(TotalsCalculation ,error){
	switch value {
		case 6:  return TotalsCalculation_Sum, nil  
		case 2:  return TotalsCalculation_Count, nil  
		case 1:  return TotalsCalculation_Average, nil  
		case 4:  return TotalsCalculation_Max, nil  
		case 5:  return TotalsCalculation_Min, nil  
		case 8:  return TotalsCalculation_Var, nil  
		case 3:  return TotalsCalculation_CountNums, nil  
		case 7:  return TotalsCalculation_StdDev, nil  
		case 0:  return TotalsCalculation_None, nil  
		case 9:  return TotalsCalculation_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid TotalsCalculation value: %d", value)
	}
}
// Class ListColumn 

// Represents a column in a Table.
type ListColumn struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ListColumn) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ListColumn_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the name of the column.
// Returns:
//   string  
func (instance *ListColumn) GetName()  (string,  error)  {
	CGoReturnPtr := C.ListColumn_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the column.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListColumn) SetName(value string)  error {
	CGoReturnPtr := C.ListColumn_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of calculation in the Totals row of the list column.
// Returns:
//   int32  
func (instance *ListColumn) GetTotalsCalculation()  (TotalsCalculation,  error)  {
	CGoReturnPtr := C.ListColumn_GetTotalsCalculation( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTotalsCalculation(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of calculation in the Totals row of the list column.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListColumn) SetTotalsCalculation(value TotalsCalculation)  error {
	CGoReturnPtr := C.ListColumn_SetTotalsCalculation( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the range of this list column.
// Returns:
//   Range  
func (instance *ListColumn) GetRange()  (*Range,  error)  {
	CGoReturnPtr := C.ListColumn_GetRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Range{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRange) 

	return result, nil 
}
// Gets the formula of totals row of this list column.
// Parameters:
//   isR1C1 - bool 
//   isLocal - bool 
// Returns:
//   string  
func (instance *ListColumn) GetCustomTotalsRowFormula(isr1c1 bool, islocal bool)  (string,  error)  {
	CGoReturnPtr := C.ListColumn_GetCustomTotalsRowFormula( instance.ptr, C.bool(isr1c1), C.bool(islocal))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the formula of totals row of this list column.
// Parameters:
//   formula - string 
//   isR1C1 - bool 
//   isLocal - bool 
// Returns:
//   void  
func (instance *ListColumn) SetCustomTotalsRowFormula(formula string, isr1c1 bool, islocal bool)  error {
	CGoReturnPtr := C.ListColumn_SetCustomTotalsRowFormula( instance.ptr, C.CString(formula), C.bool(isr1c1), C.bool(islocal))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the formula of the list column.
// Returns:
//   string  
func (instance *ListColumn) GetFormula()  (string,  error)  {
	CGoReturnPtr := C.ListColumn_GetFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the formula of the list column.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListColumn) SetFormula(value string)  error {
	CGoReturnPtr := C.ListColumn_SetFormula( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the formula of this list column.
// Parameters:
//   isR1C1 - bool 
//   isLocal - bool 
// Returns:
//   string  
func (instance *ListColumn) GetCustomCalculatedFormula(isr1c1 bool, islocal bool)  (string,  error)  {
	CGoReturnPtr := C.ListColumn_GetCustomCalculatedFormula( instance.ptr, C.bool(isr1c1), C.bool(islocal))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Sets the formula for this list column.
// Parameters:
//   formula - string 
//   isR1C1 - bool 
//   isLocal - bool 
// Returns:
//   void  
func (instance *ListColumn) SetCustomCalculatedFormula(formula string, isr1c1 bool, islocal bool)  error {
	CGoReturnPtr := C.ListColumn_SetCustomCalculatedFormula( instance.ptr, C.CString(formula), C.bool(isr1c1), C.bool(islocal))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the display labels of total row.
// Returns:
//   string  
func (instance *ListColumn) GetTotalsRowLabel()  (string,  error)  {
	CGoReturnPtr := C.ListColumn_GetTotalsRowLabel( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the display labels of total row.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListColumn) SetTotalsRowLabel(value string)  error {
	CGoReturnPtr := C.ListColumn_SetTotalsRowLabel( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the style of the data in this column of the table.
// Returns:
//   Style  
func (instance *ListColumn) GetDataStyle()  (*Style,  error)  {
	CGoReturnPtr := C.ListColumn_GetDataStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Sets the style of the data in this column of the table.
// Parameters:
//   style - Style 
// Returns:
//   void  
func (instance *ListColumn) SetDataStyle(style *Style)  error {
	CGoReturnPtr := C.ListColumn_SetDataStyle( instance.ptr, style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteListColumn(listcolumn *ListColumn){
	runtime.SetFinalizer(listcolumn, nil)
	C.Delete_ListColumn(listcolumn.ptr)
	listcolumn.ptr = nil
}

// Class ListColumnCollection 

// Represents A collection of all the <see cref="ListColumn"/> objects in the specified ListObject object.
type ListColumnCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ListColumnCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ListColumnCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the ListColumn by the index.
// Parameters:
//   index - int32 
// Returns:
//   ListColumn  
func (instance *ListColumnCollection) Get_Int(index int32)  (*ListColumn,  error)  {
	CGoReturnPtr := C.ListColumnCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ListColumn{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteListColumn) 

	return result, nil 
}
// Gets the ListColumn by the name.
// Parameters:
//   name - string 
// Returns:
//   ListColumn  
func (instance *ListColumnCollection) Get_String(name string)  (*ListColumn,  error)  {
	CGoReturnPtr := C.ListColumnCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ListColumn{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteListColumn) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *ListColumnCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.ListColumnCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteListColumnCollection(listcolumncollection *ListColumnCollection){
	runtime.SetFinalizer(listcolumncollection, nil)
	C.Delete_ListColumnCollection(listcolumncollection.ptr)
	listcolumncollection.ptr = nil
}

// Class ListObject 

// Represents a list object on a worksheet.
// The ListObject object is a member of the ListObjects collection.
// The ListObjects collection contains all the list objects on a worksheet.
type ListObject struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ListObject) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the start row of the range.
// Returns:
//   int32  
func (instance *ListObject) GetStartRow()  (int32,  error)  {
	CGoReturnPtr := C.ListObject_GetStartRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the start column of the range.
// Returns:
//   int32  
func (instance *ListObject) GetStartColumn()  (int32,  error)  {
	CGoReturnPtr := C.ListObject_GetStartColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the end  row of the range.
// Returns:
//   int32  
func (instance *ListObject) GetEndRow()  (int32,  error)  {
	CGoReturnPtr := C.ListObject_GetEndRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the end column of the range.
// Returns:
//   int32  
func (instance *ListObject) GetEndColumn()  (int32,  error)  {
	CGoReturnPtr := C.ListObject_GetEndColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets ListColumns of the ListObject.
// Returns:
//   ListColumnCollection  
func (instance *ListObject) GetListColumns()  (*ListColumnCollection,  error)  {
	CGoReturnPtr := C.ListObject_GetListColumns( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ListColumnCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteListColumnCollection) 

	return result, nil 
}
// Resize the range of the list object.
// Parameters:
//   startRow - int32 
//   startColumn - int32 
//   endRow - int32 
//   endColumn - int32 
//   hasHeaders - bool 
// Returns:
//   void  
func (instance *ListObject) Resize(startrow int32, startcolumn int32, endrow int32, endcolumn int32, hasheaders bool)  error {
	CGoReturnPtr := C.ListObject_Resize( instance.ptr, C.int(startrow), C.int(startcolumn), C.int(endrow), C.int(endcolumn), C.bool(hasheaders))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Put the value to the cell.
// Parameters:
//   rowOffset - int32 
//   columnOffset - int32 
//   value - Object 
// Returns:
//   void  
func (instance *ListObject) PutCellValue_Int_Int_Object(rowoffset int32, columnoffset int32, value *Object)  error {
	CGoReturnPtr := C.ListObject_PutCellValue_Integer_Integer_Object( instance.ptr, C.int(rowoffset), C.int(columnoffset), value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Put the value to the cell.
// Parameters:
//   rowOffset - int32 
//   columnOffset - int32 
//   value - Object 
//   isTotalsRowLabel - bool 
// Returns:
//   void  
func (instance *ListObject) PutCellValue_Int_Int_Object_Bool(rowoffset int32, columnoffset int32, value *Object, istotalsrowlabel bool)  error {
	CGoReturnPtr := C.ListObject_PutCellValue_Integer_Integer_Object_Boolean( instance.ptr, C.int(rowoffset), C.int(columnoffset), value.ptr, C.bool(istotalsrowlabel))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Put the formula to the cell in the table.
// Parameters:
//   rowOffset - int32 
//   columnOffset - int32 
//   formula - string 
// Returns:
//   void  
func (instance *ListObject) PutCellFormula_Int_Int_String(rowoffset int32, columnoffset int32, formula string)  error {
	CGoReturnPtr := C.ListObject_PutCellFormula_Integer_Integer_String( instance.ptr, C.int(rowoffset), C.int(columnoffset), C.CString(formula))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Put the formula to the cell in the table.
// Parameters:
//   rowOffset - int32 
//   columnOffset - int32 
//   formula - string 
//   isTotalsRowFormula - bool 
// Returns:
//   void  
func (instance *ListObject) PutCellFormula_Int_Int_String_Bool(rowoffset int32, columnoffset int32, formula string, istotalsrowformula bool)  error {
	CGoReturnPtr := C.ListObject_PutCellFormula_Integer_Integer_String_Boolean( instance.ptr, C.int(rowoffset), C.int(columnoffset), C.CString(formula), C.bool(istotalsrowformula))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets whether this ListObject show header row.
// Returns:
//   bool  
func (instance *ListObject) GetShowHeaderRow()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_GetShowHeaderRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets whether this ListObject show header row.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListObject) SetShowHeaderRow(value bool)  error {
	CGoReturnPtr := C.ListObject_SetShowHeaderRow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets whether this ListObject show total row.
// Returns:
//   bool  
func (instance *ListObject) GetShowTotals()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_GetShowTotals( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets whether this ListObject show total row.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListObject) SetShowTotals(value bool)  error {
	CGoReturnPtr := C.ListObject_SetShowTotals( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the data range of the ListObject.
// Returns:
//   Range  
func (instance *ListObject) GetDataRange()  (*Range,  error)  {
	CGoReturnPtr := C.ListObject_GetDataRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Range{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRange) 

	return result, nil 
}
// Updates all list columns' name from the worksheet.
// Returns:
//   void  
func (instance *ListObject) UpdateColumnName()  error {
	CGoReturnPtr := C.ListObject_UpdateColumnName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the linked QueryTable.
// Returns:
//   QueryTable  
func (instance *ListObject) GetQueryTable()  (*QueryTable,  error)  {
	CGoReturnPtr := C.ListObject_GetQueryTable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &QueryTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteQueryTable) 

	return result, nil 
}
// Gets the data source type of the table.
// Returns:
//   int32  
func (instance *ListObject) GetDataSourceType()  (TableDataSourceType,  error)  {
	CGoReturnPtr := C.ListObject_GetDataSourceType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTableDataSourceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Filter the table.
// Returns:
//   AutoFilter  
func (instance *ListObject) Filter()  (*AutoFilter,  error)  {
	CGoReturnPtr := C.ListObject_Filter( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AutoFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAutoFilter) 

	return result, nil 
}
// Gets auto filter.
// Returns:
//   AutoFilter  
func (instance *ListObject) GetAutoFilter()  (*AutoFilter,  error)  {
	CGoReturnPtr := C.ListObject_GetAutoFilter( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &AutoFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteAutoFilter) 

	return result, nil 
}
// Gets and sets the display name.
// Returns:
//   string  
func (instance *ListObject) GetDisplayName()  (string,  error)  {
	CGoReturnPtr := C.ListObject_GetDisplayName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the display name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListObject) SetDisplayName(value string)  error {
	CGoReturnPtr := C.ListObject_SetDisplayName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the comment of the table.
// Returns:
//   string  
func (instance *ListObject) GetComment()  (string,  error)  {
	CGoReturnPtr := C.ListObject_GetComment( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the comment of the table.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListObject) SetComment(value string)  error {
	CGoReturnPtr := C.ListObject_SetComment( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the first column in the table should have the style applied.
// Returns:
//   bool  
func (instance *ListObject) GetShowTableStyleFirstColumn()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_GetShowTableStyleFirstColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the first column in the table should have the style applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListObject) SetShowTableStyleFirstColumn(value bool)  error {
	CGoReturnPtr := C.ListObject_SetShowTableStyleFirstColumn( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the last column in the table should have the style applied.
// Returns:
//   bool  
func (instance *ListObject) GetShowTableStyleLastColumn()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_GetShowTableStyleLastColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the last column in the table should have the style applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListObject) SetShowTableStyleLastColumn(value bool)  error {
	CGoReturnPtr := C.ListObject_SetShowTableStyleLastColumn( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether row stripe formatting is applied.
// Returns:
//   bool  
func (instance *ListObject) GetShowTableStyleRowStripes()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_GetShowTableStyleRowStripes( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether row stripe formatting is applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListObject) SetShowTableStyleRowStripes(value bool)  error {
	CGoReturnPtr := C.ListObject_SetShowTableStyleRowStripes( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether column stripe formatting is applied.
// Returns:
//   bool  
func (instance *ListObject) GetShowTableStyleColumnStripes()  (bool,  error)  {
	CGoReturnPtr := C.ListObject_GetShowTableStyleColumnStripes( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether column stripe formatting is applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListObject) SetShowTableStyleColumnStripes(value bool)  error {
	CGoReturnPtr := C.ListObject_SetShowTableStyleColumnStripes( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Apply the table style to the range.
// Returns:
//   void  
func (instance *ListObject) ApplyStyleToRange()  error {
	CGoReturnPtr := C.ListObject_ApplyStyleToRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert the table to range.
// Returns:
//   void  
func (instance *ListObject) ConvertToRange()  error {
	CGoReturnPtr := C.ListObject_ConvertToRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert the table to range.
// Parameters:
//   options - TableToRangeOptions 
// Returns:
//   void  
func (instance *ListObject) ConvertToRange_TableToRangeOptions(options *TableToRangeOptions)  error {
	CGoReturnPtr := C.ListObject_ConvertToRange_TableToRangeOptions( instance.ptr, options.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and the built-in table style.
// Returns:
//   int32  
func (instance *ListObject) GetTableStyleType()  (TableStyleType,  error)  {
	CGoReturnPtr := C.ListObject_GetTableStyleType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTableStyleType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and the built-in table style.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListObject) SetTableStyleType(value TableStyleType)  error {
	CGoReturnPtr := C.ListObject_SetTableStyleType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the table style name.
// Returns:
//   string  
func (instance *ListObject) GetTableStyleName()  (string,  error)  {
	CGoReturnPtr := C.ListObject_GetTableStyleName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the table style name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListObject) SetTableStyleName(value string)  error {
	CGoReturnPtr := C.ListObject_SetTableStyleName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets an <see cref="XmlMap"/> used for this list.
// Returns:
//   XmlMap  
func (instance *ListObject) GetXmlMap()  (*XmlMap,  error)  {
	CGoReturnPtr := C.ListObject_GetXmlMap( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &XmlMap{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteXmlMap) 

	return result, nil 
}
// Gets and sets the alternative text.
// Returns:
//   string  
func (instance *ListObject) GetAlternativeText()  (string,  error)  {
	CGoReturnPtr := C.ListObject_GetAlternativeText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the alternative text.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListObject) SetAlternativeText(value string)  error {
	CGoReturnPtr := C.ListObject_SetAlternativeText( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the alternative description.
// Returns:
//   string  
func (instance *ListObject) GetAlternativeDescription()  (string,  error)  {
	CGoReturnPtr := C.ListObject_GetAlternativeDescription( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the alternative description.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListObject) SetAlternativeDescription(value string)  error {
	CGoReturnPtr := C.ListObject_SetAlternativeDescription( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteListObject(listobject *ListObject){
	runtime.SetFinalizer(listobject, nil)
	C.Delete_ListObject(listobject.ptr)
	listobject.ptr = nil
}

// Class ListObjectCollection 

// Represents a collection of <see cref="ListObject"/> objects in the worksheet.
type ListObjectCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ListObjectCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ListObjectCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the ListObject by index.
// Parameters:
//   index - int32 
// Returns:
//   ListObject  
func (instance *ListObjectCollection) Get_Int(index int32)  (*ListObject,  error)  {
	CGoReturnPtr := C.ListObjectCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ListObject{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteListObject) 

	return result, nil 
}
// Gets the ListObject by specified name.
// Parameters:
//   tableName - string 
// Returns:
//   ListObject  
func (instance *ListObjectCollection) Get_String(tablename string)  (*ListObject,  error)  {
	CGoReturnPtr := C.ListObjectCollection_Get_String( instance.ptr, C.CString(tablename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ListObject{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteListObject) 

	return result, nil 
}
// Adds a ListObject to the worksheet.
// Parameters:
//   startRow - int32 
//   startColumn - int32 
//   endRow - int32 
//   endColumn - int32 
//   hasHeaders - bool 
// Returns:
//   int32  
func (instance *ListObjectCollection) Add_Int_Int_Int_Int_Bool(startrow int32, startcolumn int32, endrow int32, endcolumn int32, hasheaders bool)  (int32,  error)  {
	CGoReturnPtr := C.ListObjectCollection_Add_Integer_Integer_Integer_Integer_Boolean( instance.ptr, C.int(startrow), C.int(startcolumn), C.int(endrow), C.int(endcolumn), C.bool(hasheaders))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a ListObject to the worksheet.
// Parameters:
//   startCell - string 
//   endCell - string 
//   hasHeaders - bool 
// Returns:
//   int32  
func (instance *ListObjectCollection) Add_String_String_Bool(startcell string, endcell string, hasheaders bool)  (int32,  error)  {
	CGoReturnPtr := C.ListObjectCollection_Add_String_String_Boolean( instance.ptr, C.CString(startcell), C.CString(endcell), C.bool(hasheaders))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Update all column name of the tables.
// Returns:
//   void  
func (instance *ListObjectCollection) UpdateColumnName()  error {
	CGoReturnPtr := C.ListObjectCollection_UpdateColumnName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *ListObjectCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.ListObjectCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteListObjectCollection(listobjectcollection *ListObjectCollection){
	runtime.SetFinalizer(listobjectcollection, nil)
	C.Delete_ListObjectCollection(listobjectcollection.ptr)
	listobjectcollection.ptr = nil
}

// Class TableStyle 

// Represents the table style.
type TableStyle struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TableStyle) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.TableStyle_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the name of table style.
// Returns:
//   string  
func (instance *TableStyle) GetName()  (string,  error)  {
	CGoReturnPtr := C.TableStyle_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets all elements of the table style.
// Returns:
//   TableStyleElementCollection  
func (instance *TableStyle) GetTableStyleElements()  (*TableStyleElementCollection,  error)  {
	CGoReturnPtr := C.TableStyle_GetTableStyleElements( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TableStyleElementCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTableStyleElementCollection) 

	return result, nil 
}


func DeleteTableStyle(tablestyle *TableStyle){
	runtime.SetFinalizer(tablestyle, nil)
	C.Delete_TableStyle(tablestyle.ptr)
	tablestyle.ptr = nil
}

// Class TableStyleCollection 

// Represents all custom table styles.
type TableStyleCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TableStyleCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.TableStyleCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the default style name of the table.
// Returns:
//   string  
func (instance *TableStyleCollection) GetDefaultTableStyleName()  (string,  error)  {
	CGoReturnPtr := C.TableStyleCollection_GetDefaultTableStyleName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the default style name of the table.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TableStyleCollection) SetDefaultTableStyleName(value string)  error {
	CGoReturnPtr := C.TableStyleCollection_SetDefaultTableStyleName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the  default style name of pivot table .
// Returns:
//   string  
func (instance *TableStyleCollection) GetDefaultPivotStyleName()  (string,  error)  {
	CGoReturnPtr := C.TableStyleCollection_GetDefaultPivotStyleName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the  default style name of pivot table .
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TableStyleCollection) SetDefaultPivotStyleName(value string)  error {
	CGoReturnPtr := C.TableStyleCollection_SetDefaultPivotStyleName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds a custom table style.
// Parameters:
//   name - string 
// Returns:
//   int32  
func (instance *TableStyleCollection) AddTableStyle(name string)  (int32,  error)  {
	CGoReturnPtr := C.TableStyleCollection_AddTableStyle( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a custom pivot table style.
// Parameters:
//   name - string 
// Returns:
//   int32  
func (instance *TableStyleCollection) AddPivotTableStyle(name string)  (int32,  error)  {
	CGoReturnPtr := C.TableStyleCollection_AddPivotTableStyle( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the table style by the index.
// Parameters:
//   index - int32 
// Returns:
//   TableStyle  
func (instance *TableStyleCollection) Get_Int(index int32)  (*TableStyle,  error)  {
	CGoReturnPtr := C.TableStyleCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TableStyle{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTableStyle) 

	return result, nil 
}
// Gets the table style by the name.
// Parameters:
//   name - string 
// Returns:
//   TableStyle  
func (instance *TableStyleCollection) Get_String(name string)  (*TableStyle,  error)  {
	CGoReturnPtr := C.TableStyleCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TableStyle{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTableStyle) 

	return result, nil 
}
// Gets the builtin table style
// Parameters:
//   type - int32 
// Returns:
//   TableStyle  
func (instance *TableStyleCollection) GetBuiltinTableStyle(type_ TableStyleType)  (*TableStyle,  error)  {
	CGoReturnPtr := C.TableStyleCollection_GetBuiltinTableStyle( instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TableStyle{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTableStyle) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *TableStyleCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.TableStyleCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteTableStyleCollection(tablestylecollection *TableStyleCollection){
	runtime.SetFinalizer(tablestylecollection, nil)
	C.Delete_TableStyleCollection(tablestylecollection.ptr)
	tablestylecollection.ptr = nil
}

// Class TableStyleElement 

// Represents the element of the table style.
type TableStyleElement struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TableStyleElement) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.TableStyleElement_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Number of rows or columns in a single band of striping.
// Applies only when type is firstRowStripe, secondRowStripe, firstColumnStripe, or secondColumnStripe.
// Returns:
//   int32  
func (instance *TableStyleElement) GetSize()  (int32,  error)  {
	CGoReturnPtr := C.TableStyleElement_GetSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Number of rows or columns in a single band of striping.
// Applies only when type is firstRowStripe, secondRowStripe, firstColumnStripe, or secondColumnStripe.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TableStyleElement) SetSize(value int32)  error {
	CGoReturnPtr := C.TableStyleElement_SetSize( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the element type.
// Returns:
//   int32  
func (instance *TableStyleElement) GetType()  (TableStyleElementType,  error)  {
	CGoReturnPtr := C.TableStyleElement_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTableStyleElementType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the element style.
// Returns:
//   Style  
func (instance *TableStyleElement) GetElementStyle()  (*Style,  error)  {
	CGoReturnPtr := C.TableStyleElement_GetElementStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Sets the element style.
// Parameters:
//   style - Style 
// Returns:
//   void  
func (instance *TableStyleElement) SetElementStyle(style *Style)  error {
	CGoReturnPtr := C.TableStyleElement_SetElementStyle( instance.ptr, style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteTableStyleElement(tablestyleelement *TableStyleElement){
	runtime.SetFinalizer(tablestyleelement, nil)
	C.Delete_TableStyleElement(tablestyleelement.ptr)
	tablestyleelement.ptr = nil
}

// Class TableStyleElementCollection 

// Represents all elements of the table style.
type TableStyleElementCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TableStyleElementCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.TableStyleElementCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets an element of the table style by the index.
// Parameters:
//   index - int32 
// Returns:
//   TableStyleElement  
func (instance *TableStyleElementCollection) Get_Int(index int32)  (*TableStyleElement,  error)  {
	CGoReturnPtr := C.TableStyleElementCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TableStyleElement{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTableStyleElement) 

	return result, nil 
}
// Gets the element of the table style by the element type.
// Parameters:
//   type - int32 
// Returns:
//   TableStyleElement  
func (instance *TableStyleElementCollection) Get_TableStyleElementType(type_ TableStyleElementType)  (*TableStyleElement,  error)  {
	CGoReturnPtr := C.TableStyleElementCollection_Get_TableStyleElementType( instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TableStyleElement{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTableStyleElement) 

	return result, nil 
}
// Adds an element.
// Parameters:
//   type - int32 
// Returns:
//   int32  
func (instance *TableStyleElementCollection) Add(type_ TableStyleElementType)  (int32,  error)  {
	CGoReturnPtr := C.TableStyleElementCollection_Add( instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *TableStyleElementCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.TableStyleElementCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteTableStyleElementCollection(tablestyleelementcollection *TableStyleElementCollection){
	runtime.SetFinalizer(tablestyleelementcollection, nil)
	C.Delete_TableStyleElementCollection(tablestyleelementcollection.ptr)
	tablestyleelementcollection.ptr = nil
}

// Class TableToRangeOptions 

// Represents the options when converting table to range.
type TableToRangeOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewTableToRangeOptions() ( *TableToRangeOptions, error) {
	tabletorangeoptions := &TableToRangeOptions{}
	CGoReturnPtr := C.New_TableToRangeOptions()
	if CGoReturnPtr.error_no == 0 {
		tabletorangeoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(tabletorangeoptions, DeleteTableToRangeOptions)
		return tabletorangeoptions, nil
	} else {
		tabletorangeoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return tabletorangeoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TableToRangeOptions) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.TableToRangeOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets the last row index of the table.
// Returns:
//   int32  
func (instance *TableToRangeOptions) GetLastRow()  (int32,  error)  {
	CGoReturnPtr := C.TableToRangeOptions_GetLastRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the last row index of the table.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TableToRangeOptions) SetLastRow(value int32)  error {
	CGoReturnPtr := C.TableToRangeOptions_SetLastRow( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteTableToRangeOptions(tabletorangeoptions *TableToRangeOptions){
	runtime.SetFinalizer(tabletorangeoptions, nil)
	C.Delete_TableToRangeOptions(tabletorangeoptions.ptr)
	tabletorangeoptions.ptr = nil
}
