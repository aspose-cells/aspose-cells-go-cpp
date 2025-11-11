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
 
 	
	"errors"	
	"runtime"
	"unsafe" 
)

// Class DataModel 

// Represents the data model.
type DataModel struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataModel) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModel_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets all relationships of the tables in the data model.
// Returns:
//   DataModelRelationshipCollection  
func (instance *DataModel) GetRelationships()  (*DataModelRelationshipCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataModel_GetRelationships"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataModelRelationshipCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataModelRelationshipCollection) 

	return result, nil 
}
// Gets all tables in the data model.
// Returns:
//   DataModelTableCollection  
func (instance *DataModel) GetTables()  (*DataModelTableCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataModel_GetTables"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataModelTableCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataModelTableCollection) 

	return result, nil 
}



func DeleteDataModel(datamodel *DataModel){
	runtime.SetFinalizer(datamodel, nil)
	C.Delete_CObject(C.CString("Delete_DataModel"),datamodel.ptr)
	datamodel.ptr = nil
}

// Class DataModelRelationship 

// Represents a single relationship in the spreadsheet data model.
type DataModelRelationship struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataModelRelationship) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelRelationship_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the foreign key table for this relationship.
// Returns:
//   string  
func (instance *DataModelRelationship) GetForeignKeyTable()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelRelationship_GetForeignKeyTable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the primary key table for this relationship.
// Returns:
//   string  
func (instance *DataModelRelationship) GetPrimaryKeyTable()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelRelationship_GetPrimaryKeyTable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the foreign key table column for this relationship.
// Returns:
//   string  
func (instance *DataModelRelationship) GetForeignKeyColumn()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelRelationship_GetForeignKeyColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the primary key table column for this relationship.
// Returns:
//   string  
func (instance *DataModelRelationship) GetPrimaryKeyColumn()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelRelationship_GetPrimaryKeyColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteDataModelRelationship(datamodelrelationship *DataModelRelationship){
	runtime.SetFinalizer(datamodelrelationship, nil)
	C.Delete_CObject(C.CString("Delete_DataModelRelationship"),datamodelrelationship.ptr)
	datamodelrelationship.ptr = nil
}

// Class DataModelRelationshipCollection 

// Represents the relationships.
type DataModelRelationshipCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataModelRelationshipCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelRelationshipCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the relationship.
// Parameters:
//   index - int32 
// Returns:
//   DataModelRelationship  
func (instance *DataModelRelationshipCollection) Get(index int32)  (*DataModelRelationship,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("DataModelRelationshipCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataModelRelationship{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataModelRelationship) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *DataModelRelationshipCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataModelRelationshipCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteDataModelRelationshipCollection(datamodelrelationshipcollection *DataModelRelationshipCollection){
	runtime.SetFinalizer(datamodelrelationshipcollection, nil)
	C.Delete_CObject(C.CString("Delete_DataModelRelationshipCollection"),datamodelrelationshipcollection.ptr)
	datamodelrelationshipcollection.ptr = nil
}

// Class DataModelTable 

// Represents properties of a single table in spreadsheet data model.
type DataModelTable struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataModelTable) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelTable_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the id of the data model table.
// Returns:
//   string  
func (instance *DataModelTable) GetId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelTable_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the data model table.
// Returns:
//   string  
func (instance *DataModelTable) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelTable_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the connection name of the data model table.
// Returns:
//   string  
func (instance *DataModelTable) GetConnectionName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelTable_GetConnectionName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteDataModelTable(datamodeltable *DataModelTable){
	runtime.SetFinalizer(datamodeltable, nil)
	C.Delete_CObject(C.CString("Delete_DataModelTable"),datamodeltable.ptr)
	datamodeltable.ptr = nil
}

// Class DataModelTableCollection 

// Represents the list of the data model table.
type DataModelTableCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataModelTableCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelTableCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the data model table by position of the collection.
// Parameters:
//   index - int32 
// Returns:
//   DataModelTable  
func (instance *DataModelTableCollection) Get_Int(index int32)  (*DataModelTable,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("DataModelTableCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataModelTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataModelTable) 

	return result, nil 
}
// Gets the data model table by the name.
// Parameters:
//   name - string 
// Returns:
//   DataModelTable  
func (instance *DataModelTableCollection) Get_String(name string)  (*DataModelTable,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("DataModelTableCollection_Get_String"), instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DataModelTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDataModelTable) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *DataModelTableCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataModelTableCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteDataModelTableCollection(datamodeltablecollection *DataModelTableCollection){
	runtime.SetFinalizer(datamodeltablecollection, nil)
	C.Delete_CObject(C.CString("Delete_DataModelTableCollection"),datamodeltablecollection.ptr)
	datamodeltablecollection.ptr = nil
}
