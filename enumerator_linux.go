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
"errors"
"runtime"
"unsafe" 
"time"
)


/**************Enum ReferredAreaEnumerator *****************/ 
type ReferredAreaEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *ReferredAreaEnumerator) GetCurrent() (*ReferredArea, error) {
	CGoReturnPtr := C.ReferredAreaEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &ReferredArea{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteReferredArea)
	return result, nil
}

func (instance *ReferredAreaEnumerator) Reset() error {
	CGoReturnPtr := C.ReferredAreaEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *ReferredAreaEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.ReferredAreaEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteReferredAreaEnumerator(referredareaEnumerator *ReferredAreaEnumerator){
	runtime.SetFinalizer(referredareaEnumerator, nil)
	C.Delete_ReferredAreaEnumerator(referredareaEnumerator.ptr)
	referredareaEnumerator.ptr = nil
}


/**************Enum CellEnumerator *****************/ 
type CellEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *CellEnumerator) GetCurrent() (*Cell, error) {
	CGoReturnPtr := C.CellEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &Cell{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteCell)
	return result, nil
}

func (instance *CellEnumerator) Reset() error {
	CGoReturnPtr := C.CellEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *CellEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.CellEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteCellEnumerator(cellEnumerator *CellEnumerator){
	runtime.SetFinalizer(cellEnumerator, nil)
	C.Delete_CellEnumerator(cellEnumerator.ptr)
	cellEnumerator.ptr = nil
}


/**************Enum ExternalLinkEnumerator *****************/ 
type ExternalLinkEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *ExternalLinkEnumerator) GetCurrent() (*ExternalLink, error) {
	CGoReturnPtr := C.ExternalLinkEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &ExternalLink{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteExternalLink)
	return result, nil
}

func (instance *ExternalLinkEnumerator) Reset() error {
	CGoReturnPtr := C.ExternalLinkEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *ExternalLinkEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.ExternalLinkEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteExternalLinkEnumerator(externallinkEnumerator *ExternalLinkEnumerator){
	runtime.SetFinalizer(externallinkEnumerator, nil)
	C.Delete_ExternalLinkEnumerator(externallinkEnumerator.ptr)
	externallinkEnumerator.ptr = nil
}


/**************Enum RowEnumerator *****************/ 
type RowEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *RowEnumerator) GetCurrent() (*Row, error) {
	CGoReturnPtr := C.RowEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &Row{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteRow)
	return result, nil
}

func (instance *RowEnumerator) Reset() error {
	CGoReturnPtr := C.RowEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *RowEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.RowEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteRowEnumerator(rowEnumerator *RowEnumerator){
	runtime.SetFinalizer(rowEnumerator, nil)
	C.Delete_RowEnumerator(rowEnumerator.ptr)
	rowEnumerator.ptr = nil
}


/**************Enum TextParagraphEnumerator *****************/ 
type TextParagraphEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *TextParagraphEnumerator) GetCurrent() (*TextParagraph, error) {
	CGoReturnPtr := C.TextParagraphEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &TextParagraph{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteTextParagraph)
	return result, nil
}

func (instance *TextParagraphEnumerator) Reset() error {
	CGoReturnPtr := C.TextParagraphEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *TextParagraphEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.TextParagraphEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteTextParagraphEnumerator(textparagraphEnumerator *TextParagraphEnumerator){
	runtime.SetFinalizer(textparagraphEnumerator, nil)
	C.Delete_TextParagraphEnumerator(textparagraphEnumerator.ptr)
	textparagraphEnumerator.ptr = nil
}


/**************Enum ChartPointEnumerator *****************/ 
type ChartPointEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *ChartPointEnumerator) GetCurrent() (*ChartPoint, error) {
	CGoReturnPtr := C.ChartPointEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &ChartPoint{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteChartPoint)
	return result, nil
}

func (instance *ChartPointEnumerator) Reset() error {
	CGoReturnPtr := C.ChartPointEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *ChartPointEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.ChartPointEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteChartPointEnumerator(chartpointEnumerator *ChartPointEnumerator){
	runtime.SetFinalizer(chartpointEnumerator, nil)
	C.Delete_ChartPointEnumerator(chartpointEnumerator.ptr)
	chartpointEnumerator.ptr = nil
}


/**************Enum PivotFieldEnumerator *****************/ 
type PivotFieldEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *PivotFieldEnumerator) GetCurrent() (*PivotField, error) {
	CGoReturnPtr := C.PivotFieldEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &PivotField{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeletePivotField)
	return result, nil
}

func (instance *PivotFieldEnumerator) Reset() error {
	CGoReturnPtr := C.PivotFieldEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *PivotFieldEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.PivotFieldEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeletePivotFieldEnumerator(pivotfieldEnumerator *PivotFieldEnumerator){
	runtime.SetFinalizer(pivotfieldEnumerator, nil)
	C.Delete_PivotFieldEnumerator(pivotfieldEnumerator.ptr)
	pivotfieldEnumerator.ptr = nil
}


/**************Enum PivotItemEnumerator *****************/ 
type PivotItemEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *PivotItemEnumerator) GetCurrent() (*PivotItem, error) {
	CGoReturnPtr := C.PivotItemEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &PivotItem{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeletePivotItem)
	return result, nil
}

func (instance *PivotItemEnumerator) Reset() error {
	CGoReturnPtr := C.PivotItemEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *PivotItemEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.PivotItemEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeletePivotItemEnumerator(pivotitemEnumerator *PivotItemEnumerator){
	runtime.SetFinalizer(pivotitemEnumerator, nil)
	C.Delete_PivotItemEnumerator(pivotitemEnumerator.ptr)
	pivotitemEnumerator.ptr = nil
}


/**************Enum DigitalSignatureEnumerator *****************/ 
type DigitalSignatureEnumerator struct {
	ptr unsafe.Pointer
}

func (instance *DigitalSignatureEnumerator) GetCurrent() (*DigitalSignature, error) {
	CGoReturnPtr := C.DigitalSignatureEnumerator_GetCurrent(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nil, err
	}
	result := &DigitalSignature{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteDigitalSignature)
	return result, nil
}

func (instance *DigitalSignatureEnumerator) Reset() error {
	CGoReturnPtr := C.DigitalSignatureEnumerator_Reset(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return err
	}
	return nil
}

func (instance *DigitalSignatureEnumerator) MoveNext() (bool, error) {
	CGoReturnPtr := C.DigitalSignatureEnumerator_MoveNext(instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return false, err
	}
	result := bool(CGoReturnPtr.return_value)

	return result, nil
}

func DeleteDigitalSignatureEnumerator(digitalsignatureEnumerator *DigitalSignatureEnumerator){
	runtime.SetFinalizer(digitalsignatureEnumerator, nil)
	C.Delete_DigitalSignatureEnumerator(digitalsignatureEnumerator.ptr)
	digitalsignatureEnumerator.ptr = nil
}


// Puts a boolean value into the cell.
// Parameters:
//   boolValue - bool 
// Returns:
//   void  
func (instance *Cell) PutValue_Null()  error {

	CGoReturnPtr := C.Cell_PutValue_Null( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}

// Constructs from nil.
// Parameters:
//
//	value - bool
func NewObject_Null() (*Object, error) {
	object := &Object{}
	CGoReturnPtr := C.New_Object_Null()
	if CGoReturnPtr.error_no == 0 {
		object.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(object, DeleteObject)
		return object, nil
	} else {
		object.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return object, err
	}
}

func getArrayMaxColumns(matrix [][]interface{}) int {
	maxCols := 0
	for _, row := range matrix {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}
	return maxCols
}

func toObjectArray(mixedArray []interface{}) ([]unsafe.Pointer, int) {
	array_row_length := len(mixedArray)
	vector_values := make([]unsafe.Pointer, array_row_length)
	for i, cell := range mixedArray {
		vector_values[i] = toObject(cell)
	}
	return vector_values, array_row_length
}

func toObject2Array(mixedArray [][]interface{}) ([]unsafe.Pointer, int, int) {
	array_row_length := len(mixedArray)
	array_column_length := getArrayMaxColumns(mixedArray)
	vector_values := make([]unsafe.Pointer, array_row_length*array_column_length)
	for i, row := range mixedArray {
		column_position := 0
		for j, cell := range row {
			vector_values[i*array_column_length+j] = toObject(cell)
			column_position = j
		}
		for ; column_position >= array_column_length; column_position++ {
			vector_values[i*array_column_length+column_position] = toObject(nil)
		}
	}
	return vector_values, array_row_length, array_column_length
}

func toObject(value interface{}) unsafe.Pointer {
	result, _ := NewObject_Null()
	switch v := value.(type) {
	case string:
		result, _ = NewObject_String(v)
		return result.ptr
	case bool:
		result, _ = NewObject_Bool(v)
		return result.ptr
	case byte:
		result, _ = NewObject_Byte(v)
		return result.ptr
	case int32:
		result, _ = NewObject_Int(v)
		return result.ptr
	case int:
		result, _ = NewObject_Int(int32(v))
		return result.ptr
	case float32:
		result, _ = NewObject_Float(v)
		return result.ptr
	case float64:
		result, _ = NewObject_Double(v)
		return result.ptr
	case Color:
		result, _ = NewObject_Color(&v)
		return result.ptr
	case *Color:
		result, _ = NewObject_Color(v)
		return result.ptr
	case time.Time:
		result, _ = NewObject_Date(v)
		return result.ptr
	case Range:
		result, _ = NewObject_Range(&v)
		return result.ptr
	case *Range:
		result, _ = NewObject_Range(v)
		return result.ptr
	default:
		return result.ptr
	}
}
