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

/**************Enum PowerQueryFormulaType *****************/

// Represents the type of power query formula.
type PowerQueryFormulaType int32

const(
// Formula power query formula.
PowerQueryFormulaType_Formula PowerQueryFormulaType = 0 

// Function power query formula.
PowerQueryFormulaType_Function PowerQueryFormulaType = 1 

// Parameter power query formula.
PowerQueryFormulaType_Parameter PowerQueryFormulaType = 2 
)

func Int32ToPowerQueryFormulaType(value int32)(PowerQueryFormulaType ,error){
	switch value {
		case 0:  return PowerQueryFormulaType_Formula, nil  
		case 1:  return PowerQueryFormulaType_Function, nil  
		case 2:  return PowerQueryFormulaType_Parameter, nil  
		default:
			return 0 ,fmt.Errorf("invalid PowerQueryFormulaType value: %d", value)
	}
}
// Class DataMashup 

// Represents mashup data.
type DataMashup struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataMashup) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.DataMashup_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets all power query formulas.
// Returns:
//   PowerQueryFormulaCollection  
func (instance *DataMashup) GetPowerQueryFormulas()  (*PowerQueryFormulaCollection,  error)  {
	
	CGoReturnPtr := C.DataMashup_GetPowerQueryFormulas( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaCollection) 

	return result, nil 
}



func DeleteDataMashup(datamashup *DataMashup){
	runtime.SetFinalizer(datamashup, nil)
	C.Delete_DataMashup(datamashup.ptr)
	datamashup.ptr = nil
}

// Class PowerQueryFormula 

// Represents the definition of power query formula.
type PowerQueryFormula struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormula) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of group which contains this power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormula) GetGroupName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_GetGroupName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormula) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the power query formula.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormula) SetName(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormula_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the description of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormula) GetDescription()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_GetDescription( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the description of the power query formula.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormula) SetDescription(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormula_SetDescription( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets all items of power query formula.
// Returns:
//   PowerQueryFormulaItemCollection  
func (instance *PowerQueryFormula) GetPowerQueryFormulaItems()  (*PowerQueryFormulaItemCollection,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_GetPowerQueryFormulaItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaItemCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaItemCollection) 

	return result, nil 
}
// Gets the type of this power query formula.
// Returns:
//   int32  
func (instance *PowerQueryFormula) GetType()  (PowerQueryFormulaType,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPowerQueryFormulaType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the definition of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormula) GetFormulaDefinition()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormula_GetFormulaDefinition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeletePowerQueryFormula(powerqueryformula *PowerQueryFormula){
	runtime.SetFinalizer(powerqueryformula, nil)
	C.Delete_PowerQueryFormula(powerqueryformula.ptr)
	powerqueryformula.ptr = nil
}

// Class PowerQueryFormulaCollection 

// Represents all power query formulas in the mashup data.
type PowerQueryFormulaCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormulaCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets <see cref=" PowerQueryFormula"/> by the index in the list.
// Parameters:
//   index - int32 
// Returns:
//   PowerQueryFormula  
func (instance *PowerQueryFormulaCollection) Get_Int(index int32)  (*PowerQueryFormula,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormula{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormula) 

	return result, nil 
}
// Gets <see cref=" PowerQueryFormula"/> by the name of the power query formula.
// Parameters:
//   name - string 
// Returns:
//   PowerQueryFormula  
func (instance *PowerQueryFormulaCollection) Get_String(name string)  (*PowerQueryFormula,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormula{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormula) 

	return result, nil 
}
// Remove power query formula by name.
// Parameters:
//   name - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaCollection) RemoveBy(name string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaCollection_RemoveBy( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *PowerQueryFormulaCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeletePowerQueryFormulaCollection(powerqueryformulacollection *PowerQueryFormulaCollection){
	runtime.SetFinalizer(powerqueryformulacollection, nil)
	C.Delete_PowerQueryFormulaCollection(powerqueryformulacollection.ptr)
	powerqueryformulacollection.ptr = nil
}

// Class PowerQueryFormulaFunction 

// Represents the function of power query.
type PowerQueryFormulaFunction struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PowerQueryFormula 
func NewPowerQueryFormulaFunction(src *PowerQueryFormula) ( *PowerQueryFormulaFunction, error) {
	powerqueryformulafunction := &PowerQueryFormulaFunction{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.New_PowerQueryFormulaFunction(src_ptr)
	if CGoReturnPtr.error_no == 0 {
		powerqueryformulafunction.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(powerqueryformulafunction, DeletePowerQueryFormulaFunction)
		return powerqueryformulafunction, nil
	} else {
		powerqueryformulafunction.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return powerqueryformulafunction, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormulaFunction) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of power query formula.
// Returns:
//   int32  
func (instance *PowerQueryFormulaFunction) GetType()  (PowerQueryFormulaType,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPowerQueryFormulaType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the definition of function.
// Returns:
//   string  
func (instance *PowerQueryFormulaFunction) GetF()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetF( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the definition of function.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaFunction) SetF(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_SetF( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the name of group which contains this power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaFunction) GetGroupName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetGroupName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaFunction) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the power query formula.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaFunction) SetName(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the description of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaFunction) GetDescription()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetDescription( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the description of the power query formula.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaFunction) SetDescription(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_SetDescription( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets all items of power query formula.
// Returns:
//   PowerQueryFormulaItemCollection  
func (instance *PowerQueryFormulaFunction) GetPowerQueryFormulaItems()  (*PowerQueryFormulaItemCollection,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetPowerQueryFormulaItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaItemCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaItemCollection) 

	return result, nil 
}
// Gets the definition of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaFunction) GetFormulaDefinition()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaFunction_GetFormulaDefinition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}


func (instance *PowerQueryFormulaFunction) ToPowerQueryFormula() *PowerQueryFormula {
	parentClass := &PowerQueryFormula{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeletePowerQueryFormulaFunction(powerqueryformulafunction *PowerQueryFormulaFunction){
	runtime.SetFinalizer(powerqueryformulafunction, nil)
	C.Delete_PowerQueryFormulaFunction(powerqueryformulafunction.ptr)
	powerqueryformulafunction.ptr = nil
}

// Class PowerQueryFormulaItem 

// Represents the item of the power query formula.
type PowerQueryFormulaItem struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormulaItem) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItem_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the item.
// Returns:
//   string  
func (instance *PowerQueryFormulaItem) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItem_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the value of the item.
// Returns:
//   string  
func (instance *PowerQueryFormulaItem) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItem_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the value of the item.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaItem) SetValue(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaItem_SetValue( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeletePowerQueryFormulaItem(powerqueryformulaitem *PowerQueryFormulaItem){
	runtime.SetFinalizer(powerqueryformulaitem, nil)
	C.Delete_PowerQueryFormulaItem(powerqueryformulaitem.ptr)
	powerqueryformulaitem.ptr = nil
}

// Class PowerQueryFormulaItemCollection 

// Represents all item of the power query formula.
type PowerQueryFormulaItemCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormulaItemCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItemCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets <see cref=" PowerQueryFormulaItem"/> by the index in the list.
// Parameters:
//   index - int32 
// Returns:
//   PowerQueryFormulaItem  
func (instance *PowerQueryFormulaItemCollection) Get_Int(index int32)  (*PowerQueryFormulaItem,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItemCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaItem{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaItem) 

	return result, nil 
}
// Gets <see cref=" PowerQueryFormulaItem"/> by the name of the item.
// Parameters:
//   name - string 
// Returns:
//   PowerQueryFormulaItem  
func (instance *PowerQueryFormulaItemCollection) Get_String(name string)  (*PowerQueryFormulaItem,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItemCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaItem{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaItem) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *PowerQueryFormulaItemCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaItemCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeletePowerQueryFormulaItemCollection(powerqueryformulaitemcollection *PowerQueryFormulaItemCollection){
	runtime.SetFinalizer(powerqueryformulaitemcollection, nil)
	C.Delete_PowerQueryFormulaItemCollection(powerqueryformulaitemcollection.ptr)
	powerqueryformulaitemcollection.ptr = nil
}

// Class PowerQueryFormulaParameter 

// Represents the parameter of power query formula.
type PowerQueryFormulaParameter struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PowerQueryFormula 
func NewPowerQueryFormulaParameter(src *PowerQueryFormula) ( *PowerQueryFormulaParameter, error) {
	powerqueryformulaparameter := &PowerQueryFormulaParameter{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.New_PowerQueryFormulaParameter(src_ptr)
	if CGoReturnPtr.error_no == 0 {
		powerqueryformulaparameter.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(powerqueryformulaparameter, DeletePowerQueryFormulaParameter)
		return powerqueryformulaparameter, nil
	} else {
		powerqueryformulaparameter.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return powerqueryformulaparameter, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormulaParameter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of power query formula.
// Returns:
//   int32  
func (instance *PowerQueryFormulaParameter) GetType()  (PowerQueryFormulaType,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPowerQueryFormulaType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the value of parameter.
// Returns:
//   string  
func (instance *PowerQueryFormulaParameter) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the value of parameter.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaParameter) SetValue(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_SetValue( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the definition of the parameter.
// Returns:
//   string  
func (instance *PowerQueryFormulaParameter) GetFormulaDefinition()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetFormulaDefinition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of group which contains this power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaParameter) GetGroupName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetGroupName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaParameter) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the power query formula.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaParameter) SetName(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the description of the power query formula.
// Returns:
//   string  
func (instance *PowerQueryFormulaParameter) GetDescription()  (string,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetDescription( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the description of the power query formula.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PowerQueryFormulaParameter) SetDescription(value string)  error {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_SetDescription( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets all items of power query formula.
// Returns:
//   PowerQueryFormulaItemCollection  
func (instance *PowerQueryFormulaParameter) GetPowerQueryFormulaItems()  (*PowerQueryFormulaItemCollection,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameter_GetPowerQueryFormulaItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaItemCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaItemCollection) 

	return result, nil 
}


func (instance *PowerQueryFormulaParameter) ToPowerQueryFormula() *PowerQueryFormula {
	parentClass := &PowerQueryFormula{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeletePowerQueryFormulaParameter(powerqueryformulaparameter *PowerQueryFormulaParameter){
	runtime.SetFinalizer(powerqueryformulaparameter, nil)
	C.Delete_PowerQueryFormulaParameter(powerqueryformulaparameter.ptr)
	powerqueryformulaparameter.ptr = nil
}

// Class PowerQueryFormulaParameterCollection 

// Represents the parameters of power query formula.
type PowerQueryFormulaParameterCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPowerQueryFormulaParameterCollection() ( *PowerQueryFormulaParameterCollection, error) {
	powerqueryformulaparametercollection := &PowerQueryFormulaParameterCollection{}
	CGoReturnPtr := C.New_PowerQueryFormulaParameterCollection()
	if CGoReturnPtr.error_no == 0 {
		powerqueryformulaparametercollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(powerqueryformulaparametercollection, DeletePowerQueryFormulaParameterCollection)
		return powerqueryformulaparametercollection, nil
	} else {
		powerqueryformulaparametercollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return powerqueryformulaparametercollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PowerQueryFormulaParameterCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameterCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets <see cref=" PowerQueryFormulaParameter"/> by the index in the list.
// Parameters:
//   index - int32 
// Returns:
//   PowerQueryFormulaParameter  
func (instance *PowerQueryFormulaParameterCollection) Get_Int(index int32)  (*PowerQueryFormulaParameter,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameterCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaParameter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaParameter) 

	return result, nil 
}
// Gets <see cref=" PowerQueryFormulaParameter"/> by the name of the item.
// Parameters:
//   name - string 
// Returns:
//   PowerQueryFormulaParameter  
func (instance *PowerQueryFormulaParameterCollection) Get_String(name string)  (*PowerQueryFormulaParameter,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameterCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormulaParameter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormulaParameter) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *PowerQueryFormulaParameterCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.PowerQueryFormulaParameterCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeletePowerQueryFormulaParameterCollection(powerqueryformulaparametercollection *PowerQueryFormulaParameterCollection){
	runtime.SetFinalizer(powerqueryformulaparametercollection, nil)
	C.Delete_PowerQueryFormulaParameterCollection(powerqueryformulaparametercollection.ptr)
	powerqueryformulaparametercollection.ptr = nil
}
