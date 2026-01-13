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
 	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum ConnectionDataSourceType *****************/

// Specifies external database source type
type ConnectionDataSourceType int32

const(
// ODBC-based source
ConnectionDataSourceType_ODBCBasedSource ConnectionDataSourceType = 1 

// DAO-based source
ConnectionDataSourceType_DAOBasedSource ConnectionDataSourceType = 2 

// File based database source
ConnectionDataSourceType_FileBasedDataBaseSource ConnectionDataSourceType = 3 

// Web query
ConnectionDataSourceType_WebQuery ConnectionDataSourceType = 4 

// OLE DB-based source
ConnectionDataSourceType_OLEDBBasedSource ConnectionDataSourceType = 5 

// Text-based source
ConnectionDataSourceType_TextBasedSource ConnectionDataSourceType = 6 

// ADO record set
ConnectionDataSourceType_ADORecordSet ConnectionDataSourceType = 7 

// DSP
ConnectionDataSourceType_DSP ConnectionDataSourceType = 8 

// OLE DB data source created by the Spreadsheet Data Model.
ConnectionDataSourceType_OLEDBDataModel ConnectionDataSourceType = 100 

// Data feed data source created by the Spreadsheet Data Model.
ConnectionDataSourceType_DataFeedDataModel ConnectionDataSourceType = 101 

// Worksheet data source created by the Spreadsheet Data Model.
ConnectionDataSourceType_WorksheetDataModel ConnectionDataSourceType = 102 

// Text data source created by the Spreadsheet Data Model.
ConnectionDataSourceType_TextDataModel ConnectionDataSourceType = 103 

// Text data source created by the Spreadsheet Data Model.
ConnectionDataSourceType_Unknown ConnectionDataSourceType = 255 
)

func Int32ToConnectionDataSourceType(value int32)(ConnectionDataSourceType ,error){
	switch value {
		case 1:  return ConnectionDataSourceType_ODBCBasedSource, nil  
		case 2:  return ConnectionDataSourceType_DAOBasedSource, nil  
		case 3:  return ConnectionDataSourceType_FileBasedDataBaseSource, nil  
		case 4:  return ConnectionDataSourceType_WebQuery, nil  
		case 5:  return ConnectionDataSourceType_OLEDBBasedSource, nil  
		case 6:  return ConnectionDataSourceType_TextBasedSource, nil  
		case 7:  return ConnectionDataSourceType_ADORecordSet, nil  
		case 8:  return ConnectionDataSourceType_DSP, nil  
		case 100:  return ConnectionDataSourceType_OLEDBDataModel, nil  
		case 101:  return ConnectionDataSourceType_DataFeedDataModel, nil  
		case 102:  return ConnectionDataSourceType_WorksheetDataModel, nil  
		case 103:  return ConnectionDataSourceType_TextDataModel, nil  
		case 255:  return ConnectionDataSourceType_Unknown, nil  
		default:
			return 0 ,fmt.Errorf("invalid ConnectionDataSourceType value: %d", value)
	}
}

/**************Enum ConnectionParameterType *****************/

// Specifies the parameter type of external connection
type ConnectionParameterType int32

const(
// Get the parameter value from a cell on each refresh.
ConnectionParameterType_Cell ConnectionParameterType = 0 

// Prompt the user on each refresh for a parameter value.
ConnectionParameterType_Prompt ConnectionParameterType = 1 

// Use a constant value on each refresh for the parameter value.
ConnectionParameterType_Value ConnectionParameterType = 2 
)

func Int32ToConnectionParameterType(value int32)(ConnectionParameterType ,error){
	switch value {
		case 0:  return ConnectionParameterType_Cell, nil  
		case 1:  return ConnectionParameterType_Prompt, nil  
		case 2:  return ConnectionParameterType_Value, nil  
		default:
			return 0 ,fmt.Errorf("invalid ConnectionParameterType value: %d", value)
	}
}

/**************Enum CredentialsMethodType *****************/

// Specifies Credentials method used for server access.
type CredentialsMethodType int32

const(
// Integrated Authentication
CredentialsMethodType_Integrated CredentialsMethodType = 0 

// No Credentials
CredentialsMethodType_None CredentialsMethodType = 1 

// Prompt Credentials
CredentialsMethodType_Prompt CredentialsMethodType = 2 

// Stored Credentials
CredentialsMethodType_Stored CredentialsMethodType = 3 
)

func Int32ToCredentialsMethodType(value int32)(CredentialsMethodType ,error){
	switch value {
		case 0:  return CredentialsMethodType_Integrated, nil  
		case 1:  return CredentialsMethodType_None, nil  
		case 2:  return CredentialsMethodType_Prompt, nil  
		case 3:  return CredentialsMethodType_Stored, nil  
		default:
			return 0 ,fmt.Errorf("invalid CredentialsMethodType value: %d", value)
	}
}

/**************Enum ExternalConnectionClassType *****************/

// Represents the type of connection
type ExternalConnectionClassType int32

const(
// ODBC or OLE DB
ExternalConnectionClassType_Database ExternalConnectionClassType = 0 

// Web query
ExternalConnectionClassType_WebQuery ExternalConnectionClassType = 1 

// Based on text
ExternalConnectionClassType_TextBased ExternalConnectionClassType = 2 

// Data model
ExternalConnectionClassType_DataModel ExternalConnectionClassType = 3 


ExternalConnectionClassType_Unkown ExternalConnectionClassType = 4 
)

func Int32ToExternalConnectionClassType(value int32)(ExternalConnectionClassType ,error){
	switch value {
		case 0:  return ExternalConnectionClassType_Database, nil  
		case 1:  return ExternalConnectionClassType_WebQuery, nil  
		case 2:  return ExternalConnectionClassType_TextBased, nil  
		case 3:  return ExternalConnectionClassType_DataModel, nil  
		case 4:  return ExternalConnectionClassType_Unkown, nil  
		default:
			return 0 ,fmt.Errorf("invalid ExternalConnectionClassType value: %d", value)
	}
}

/**************Enum HtmlFormatHandlingType *****************/

// Specifies how to handle formatting from the HTML source
type HtmlFormatHandlingType int32

const(
// Transfer all HTML formatting into the worksheet along with data.
HtmlFormatHandlingType_All HtmlFormatHandlingType = 0 

// Bring data in as unformatted text (setting data types still occurs).
HtmlFormatHandlingType_None HtmlFormatHandlingType = 1 

// Translate HTML formatting to rich text formatting on the data brought into the worksheet.
HtmlFormatHandlingType_Rtf HtmlFormatHandlingType = 2 
)

func Int32ToHtmlFormatHandlingType(value int32)(HtmlFormatHandlingType ,error){
	switch value {
		case 0:  return HtmlFormatHandlingType_All, nil  
		case 1:  return HtmlFormatHandlingType_None, nil  
		case 2:  return HtmlFormatHandlingType_Rtf, nil  
		default:
			return 0 ,fmt.Errorf("invalid HtmlFormatHandlingType value: %d", value)
	}
}

/**************Enum OLEDBCommandType *****************/

// Specifies the OLE DB command type.
type OLEDBCommandType int32

const(
// The command type is not specified.
OLEDBCommandType_None OLEDBCommandType = 0 

// Specifies a cube name
OLEDBCommandType_CubeName OLEDBCommandType = 1 

// Specifies a SQL statement
OLEDBCommandType_SqlStatement OLEDBCommandType = 2 

// Specifies a table name
OLEDBCommandType_TableName OLEDBCommandType = 3 

// Specifies that default information has been given, and it is up to the provider how to interpret.
OLEDBCommandType_DefaultInformation OLEDBCommandType = 4 

// Specifies a query which is against a web based List Data Provider.
OLEDBCommandType_WebBasedList OLEDBCommandType = 5 

// Specifies the table list.
OLEDBCommandType_TableCollection OLEDBCommandType = 6 
)

func Int32ToOLEDBCommandType(value int32)(OLEDBCommandType ,error){
	switch value {
		case 0:  return OLEDBCommandType_None, nil  
		case 1:  return OLEDBCommandType_CubeName, nil  
		case 2:  return OLEDBCommandType_SqlStatement, nil  
		case 3:  return OLEDBCommandType_TableName, nil  
		case 4:  return OLEDBCommandType_DefaultInformation, nil  
		case 5:  return OLEDBCommandType_WebBasedList, nil  
		case 6:  return OLEDBCommandType_TableCollection, nil  
		default:
			return 0 ,fmt.Errorf("invalid OLEDBCommandType value: %d", value)
	}
}

/**************Enum ReConnectionMethodType *****************/

// Specifies what the spreadsheet application should do when a connection fails.
type ReConnectionMethodType int32

const(
// On refresh use the existing connection information and if it ends up being invalid
// then get updated connection information, if available from the external connection file.
ReConnectionMethodType_Required ReConnectionMethodType = 1 

// On every refresh get updated connection information from the external connection file,
// if available, and use that instead of the existing connection information.
// In this case the data refresh will fail if the external connection file is unavailable.
ReConnectionMethodType_Always ReConnectionMethodType = 2 

// Never get updated connection information from the external connection file
// even if it is available and even if the existing connection information is invalid
ReConnectionMethodType_Never ReConnectionMethodType = 3 
)

func Int32ToReConnectionMethodType(value int32)(ReConnectionMethodType ,error){
	switch value {
		case 1:  return ReConnectionMethodType_Required, nil  
		case 2:  return ReConnectionMethodType_Always, nil  
		case 3:  return ReConnectionMethodType_Never, nil  
		default:
			return 0 ,fmt.Errorf("invalid ReConnectionMethodType value: %d", value)
	}
}

/**************Enum SqlDataType *****************/

// Specifies SQL data type of the parameter. Only valid for ODBC sources.
type SqlDataType int32

const(
// sql unsigned offset
SqlDataType_SqlUnsignedOffset SqlDataType = -22 

// sql signed offset
SqlDataType_SqlSignedOffset SqlDataType = -20 

// sql guid
SqlDataType_SqlGUID SqlDataType = -11 

// sql wide long variable char
SqlDataType_SqlWLongVarchar SqlDataType = -10 

// sql wide variable char
SqlDataType_SqlWVarchar SqlDataType = -9 

// sql wide char
SqlDataType_SqlWChar SqlDataType = -8 

// sql bit
SqlDataType_SqlBit SqlDataType = -7 

// sql tiny int
SqlDataType_SqlTinyInt SqlDataType = -6 

// sql big int
SqlDataType_SqlBigInt SqlDataType = -5 

// sql long variable binary
SqlDataType_SqlLongVarBinary SqlDataType = -4 

// sql variable binary
SqlDataType_SqlVarBinary SqlDataType = -3 

// sql binary
SqlDataType_SqlBinary SqlDataType = -2 

// sql long variable char
SqlDataType_SqlLongVarChar SqlDataType = -1 

// sql unknown type
SqlDataType_SqlUnknownType SqlDataType = 0 

// sql char
SqlDataType_SqlChar SqlDataType = 1 

// sql numeric
SqlDataType_SqlNumeric SqlDataType = 2 

// sql decimal
SqlDataType_SqlDecimal SqlDataType = 3 

// sql integer
SqlDataType_SqlInteger SqlDataType = 4 

// sql small int
SqlDataType_SqlSmallInt SqlDataType = 5 

// sql float
SqlDataType_SqlFloat SqlDataType = 6 

// sql real
SqlDataType_SqlReal SqlDataType = 7 

// sql double
SqlDataType_SqlDouble SqlDataType = 8 

// sql date type
SqlDataType_SqlTypeDate SqlDataType = 9 

// sql time type
SqlDataType_SqlTypeTime SqlDataType = 10 

// sql timestamp type
SqlDataType_SqlTypeTimestamp SqlDataType = 11 

// sql variable char
SqlDataType_SqlVarChar SqlDataType = 12 

// sql interval year
SqlDataType_SqlIntervalYear SqlDataType = 101 

// sql interval month
SqlDataType_SqlIntervalMonth SqlDataType = 102 

// sql interval day
SqlDataType_SqlIntervalDay SqlDataType = 103 

// sql interval hour
SqlDataType_SqlIntervalHour SqlDataType = 104 

// sql interval minute
SqlDataType_SqlIntervalMinute SqlDataType = 105 

// sql interval second
SqlDataType_SqlIntervalSecond SqlDataType = 106 

// sql interval year to month
SqlDataType_SqlIntervalYearToMonth SqlDataType = 107 

// sql interval day to hour
SqlDataType_SqlIntervalDayToHour SqlDataType = 108 

// sql interval day to minute
SqlDataType_SqlIntervalDayToMinute SqlDataType = 109 

// sql interval day to second
SqlDataType_SqlIntervalDayToSecond SqlDataType = 110 

// sql interval hour to minute
SqlDataType_SqlIntervalHourToMinute SqlDataType = 111 

// sql interval hour to second
SqlDataType_SqlIntervalHourToSecond SqlDataType = 112 

// sql interval minute to second
SqlDataType_SqlIntervalMinuteToSecond SqlDataType = 113 
)

func Int32ToSqlDataType(value int32)(SqlDataType ,error){
	switch value {
		case -22:  return SqlDataType_SqlUnsignedOffset, nil  
		case -20:  return SqlDataType_SqlSignedOffset, nil  
		case -11:  return SqlDataType_SqlGUID, nil  
		case -10:  return SqlDataType_SqlWLongVarchar, nil  
		case -9:  return SqlDataType_SqlWVarchar, nil  
		case -8:  return SqlDataType_SqlWChar, nil  
		case -7:  return SqlDataType_SqlBit, nil  
		case -6:  return SqlDataType_SqlTinyInt, nil  
		case -5:  return SqlDataType_SqlBigInt, nil  
		case -4:  return SqlDataType_SqlLongVarBinary, nil  
		case -3:  return SqlDataType_SqlVarBinary, nil  
		case -2:  return SqlDataType_SqlBinary, nil  
		case -1:  return SqlDataType_SqlLongVarChar, nil  
		case 0:  return SqlDataType_SqlUnknownType, nil  
		case 1:  return SqlDataType_SqlChar, nil  
		case 2:  return SqlDataType_SqlNumeric, nil  
		case 3:  return SqlDataType_SqlDecimal, nil  
		case 4:  return SqlDataType_SqlInteger, nil  
		case 5:  return SqlDataType_SqlSmallInt, nil  
		case 6:  return SqlDataType_SqlFloat, nil  
		case 7:  return SqlDataType_SqlReal, nil  
		case 8:  return SqlDataType_SqlDouble, nil  
		case 9:  return SqlDataType_SqlTypeDate, nil  
		case 10:  return SqlDataType_SqlTypeTime, nil  
		case 11:  return SqlDataType_SqlTypeTimestamp, nil  
		case 12:  return SqlDataType_SqlVarChar, nil  
		case 101:  return SqlDataType_SqlIntervalYear, nil  
		case 102:  return SqlDataType_SqlIntervalMonth, nil  
		case 103:  return SqlDataType_SqlIntervalDay, nil  
		case 104:  return SqlDataType_SqlIntervalHour, nil  
		case 105:  return SqlDataType_SqlIntervalMinute, nil  
		case 106:  return SqlDataType_SqlIntervalSecond, nil  
		case 107:  return SqlDataType_SqlIntervalYearToMonth, nil  
		case 108:  return SqlDataType_SqlIntervalDayToHour, nil  
		case 109:  return SqlDataType_SqlIntervalDayToMinute, nil  
		case 110:  return SqlDataType_SqlIntervalDayToSecond, nil  
		case 111:  return SqlDataType_SqlIntervalHourToMinute, nil  
		case 112:  return SqlDataType_SqlIntervalHourToSecond, nil  
		case 113:  return SqlDataType_SqlIntervalMinuteToSecond, nil  
		default:
			return 0 ,fmt.Errorf("invalid SqlDataType value: %d", value)
	}
}
// Class ConnectionParameter 

// Specifies properties about any parameters used with external data connections
// Parameters are valid for ODBC and web queries.
type ConnectionParameter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ConnectionParameter) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ConnectionParameter_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// SQL data type of the parameter. Only valid for ODBC sources.
// Returns:
//   int32  
func (instance *ConnectionParameter) GetSqlType()  (SqlDataType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ConnectionParameter_GetSqlType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSqlDataType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// SQL data type of the parameter. Only valid for ODBC sources.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ConnectionParameter) SetSqlType(value SqlDataType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ConnectionParameter_SetSqlType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Flag indicating whether the query should automatically refresh when the contents of a
// cell that provides the parameter value changes. If true, then external data is refreshed
// using the new parameter value every time there's a change. If false, then external data
// is only refreshed when requested by the user, or some other event triggers refresh (e.g., workbook opened).
// Returns:
//   bool  
func (instance *ConnectionParameter) GetRefreshOnChange()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ConnectionParameter_GetRefreshOnChange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether the query should automatically refresh when the contents of a
// cell that provides the parameter value changes. If true, then external data is refreshed
// using the new parameter value every time there's a change. If false, then external data
// is only refreshed when requested by the user, or some other event triggers refresh (e.g., workbook opened).
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ConnectionParameter) SetRefreshOnChange(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ConnectionParameter_SetRefreshOnChange"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Prompt string for the parameter. Presented to the spreadsheet user along with input UI
// to collect the parameter value before refreshing the external data. Used only when
// parameterType = prompt.
// Returns:
//   string  
func (instance *ConnectionParameter) GetPrompt()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ConnectionParameter_GetPrompt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Prompt string for the parameter. Presented to the spreadsheet user along with input UI
// to collect the parameter value before refreshing the external data. Used only when
// parameterType = prompt.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ConnectionParameter) SetPrompt(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ConnectionParameter_SetPrompt"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Type of parameter used.
// If the parameterType=value, then the value from boolean, double, integer,
// or string will be used.  In this case, it is expected that only one of
// {boolean, double, integer, or string} will be specified.
// Returns:
//   int32  
func (instance *ConnectionParameter) GetType()  (ConnectionParameterType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ConnectionParameter_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToConnectionParameterType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Type of parameter used.
// If the parameterType=value, then the value from boolean, double, integer,
// or string will be used.  In this case, it is expected that only one of
// {boolean, double, integer, or string} will be specified.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ConnectionParameter) SetType(value ConnectionParameterType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ConnectionParameter_SetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The name of the parameter.
// Returns:
//   string  
func (instance *ConnectionParameter) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ConnectionParameter_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The name of the parameter.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ConnectionParameter) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ConnectionParameter_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Cell reference indicating which cell's value to use for the query parameter. Used only when parameterType is cell.
// Returns:
//   string  
func (instance *ConnectionParameter) GetCellReference()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ConnectionParameter_GetCellReference"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Cell reference indicating which cell's value to use for the query parameter. Used only when parameterType is cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ConnectionParameter) SetCellReference(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ConnectionParameter_SetCellReference"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Non-integer numeric value,Integer value,String value or Boolean value
// to use as the query parameter. Used only when parameterType is value.
// Returns:
//   Object  
func (instance *ConnectionParameter) GetValue()  (*Object,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ConnectionParameter_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Non-integer numeric value,Integer value,String value or Boolean value
// to use as the query parameter. Used only when parameterType is value.
// Parameters:
//   value - Object 
// Returns:
//   void  
func (instance *ConnectionParameter) SetValue(value *Object)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ConnectionParameter_SetValue"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteConnectionParameter(connectionparameter *ConnectionParameter){
	runtime.SetFinalizer(connectionparameter, nil)
	C.Delete_CObject(C.CString("Delete_ConnectionParameter"),connectionparameter.ptr)
	connectionparameter.ptr = nil
}

// Class ConnectionParameterCollection 

// Specifies the <see cref="ConnectionParameter"/> collection
type ConnectionParameterCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ConnectionParameterCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ConnectionParameterCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="ConnectionParameter"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   ConnectionParameter  
func (instance *ConnectionParameterCollection) Get_Int(index int32)  (*ConnectionParameter,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("ConnectionParameterCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ConnectionParameter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteConnectionParameter) 

	return result, nil 
}
// Gets the <see cref="ConnectionParameter"/> element with the specified name.
// Parameters:
//   connParamName - string 
// Returns:
//   ConnectionParameter  
func (instance *ConnectionParameterCollection) Get_String(connparamname string)  (*ConnectionParameter,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("ConnectionParameterCollection_Get_String"), instance.ptr, C.CString(connparamname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ConnectionParameter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteConnectionParameter) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *ConnectionParameterCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ConnectionParameterCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteConnectionParameterCollection(connectionparametercollection *ConnectionParameterCollection){
	runtime.SetFinalizer(connectionparametercollection, nil)
	C.Delete_CObject(C.CString("Delete_ConnectionParameterCollection"),connectionparametercollection.ptr)
	connectionparametercollection.ptr = nil
}

// Class DataModelConnection 

// Specifies a data model connection
type DataModelConnection struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ExternalConnection 
func NewDataModelConnection(src *ExternalConnection) ( *DataModelConnection, error) {
	datamodelconnection := &DataModelConnection{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_DataModelConnection"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		datamodelconnection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(datamodelconnection, DeleteDataModelConnection)
		return datamodelconnection, nil
	} else {
		datamodelconnection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return datamodelconnection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DataModelConnection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of this <see cref="ExternalConnection"/> object.
// Returns:
//   int32  
func (instance *DataModelConnection) GetClassType()  (ExternalConnectionClassType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataModelConnection_GetClassType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToExternalConnectionClassType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Returns:
//   string  
func (instance *DataModelConnection) GetCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns <see cref="OLEDBCommandType"/> type.
// Returns:
//   int32  
func (instance *DataModelConnection) GetCommandType()  (OLEDBCommandType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataModelConnection_GetCommandType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOLEDBCommandType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Returns <see cref="OLEDBCommandType"/> type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataModelConnection) SetCommandType(value OLEDBCommandType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataModelConnection_SetCommandType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Returns:
//   string  
func (instance *DataModelConnection) GetConnectionString()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetConnectionString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetConnectionString(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetConnectionString"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the id of the connection.
// Returns:
//   int32  
func (instance *DataModelConnection) GetId()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataModelConnection_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Returns:
//   int32  
func (instance *DataModelConnection) GetSourceType()  (ConnectionDataSourceType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataModelConnection_GetSourceType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToConnectionDataSourceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataModelConnection) SetSourceType(value ConnectionDataSourceType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataModelConnection_SetSourceType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Returns:
//   string  
func (instance *DataModelConnection) GetSSOId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetSSOId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetSSOId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetSSOId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Returns:
//   bool  
func (instance *DataModelConnection) GetSavePassword()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_GetSavePassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetSavePassword(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetSavePassword"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Returns:
//   bool  
func (instance *DataModelConnection) GetSaveData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_GetSaveData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetSaveData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetSaveData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Returns:
//   bool  
func (instance *DataModelConnection) GetRefreshOnLoad()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_GetRefreshOnLoad"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetRefreshOnLoad(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetRefreshOnLoad"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Returns:
//   int32  
func (instance *DataModelConnection) GetReconnectionMethodType()  (ReConnectionMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataModelConnection_GetReconnectionMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToReConnectionMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataModelConnection) SetReconnectionMethodType(value ReConnectionMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataModelConnection_SetReconnectionMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Returns:
//   bool  
func (instance *DataModelConnection) GetOnlyUseConnectionFile()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_GetOnlyUseConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetOnlyUseConnectionFile(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetOnlyUseConnectionFile"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Returns:
//   string  
func (instance *DataModelConnection) GetOdcFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetOdcFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetOdcFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetOdcFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Returns:
//   string  
func (instance *DataModelConnection) GetSourceFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetSourceFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetSourceFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetSourceFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Returns:
//   bool  
func (instance *DataModelConnection) IsNew()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_IsNew"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetIsNew(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetIsNew"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Returns:
//   string  
func (instance *DataModelConnection) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Returns:
//   bool  
func (instance *DataModelConnection) GetKeepAlive()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_GetKeepAlive"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetKeepAlive(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetKeepAlive"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Returns:
//   int32  
func (instance *DataModelConnection) GetRefreshInternal()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DataModelConnection_GetRefreshInternal"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataModelConnection) SetRefreshInternal(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DataModelConnection_SetRefreshInternal"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the user description for this connection
// Returns:
//   string  
func (instance *DataModelConnection) GetConnectionDescription()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetConnectionDescription"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the user description for this connection
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetConnectionDescription(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetConnectionDescription"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Returns:
//   bool  
func (instance *DataModelConnection) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Returns:
//   int32  
func (instance *DataModelConnection) GetCredentialsMethodType()  (CredentialsMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DataModelConnection_GetCredentialsMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCredentialsMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DataModelConnection) SetCredentialsMethodType(value CredentialsMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DataModelConnection_SetCredentialsMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Returns:
//   bool  
func (instance *DataModelConnection) GetBackgroundRefresh()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DataModelConnection_GetBackgroundRefresh"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DataModelConnection) SetBackgroundRefresh(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DataModelConnection_SetBackgroundRefresh"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets <see cref="ConnectionParameterCollection"/> for an ODBC or web query.
// Returns:
//   ConnectionParameterCollection  
func (instance *DataModelConnection) GetParameters()  (*ConnectionParameterCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataModelConnection_GetParameters"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ConnectionParameterCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteConnectionParameterCollection) 

	return result, nil 
}
// Gets the definition of power query formula.
// Returns:
//   PowerQueryFormula  
func (instance *DataModelConnection) GetPowerQueryFormula()  (*PowerQueryFormula,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DataModelConnection_GetPowerQueryFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormula{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormula) 

	return result, nil 
}
// Gets the connection file.
// Returns:
//   string  
func (instance *DataModelConnection) GetConnectionFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Returns:
//   string  
func (instance *DataModelConnection) GetSecondCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DataModelConnection_GetSecondCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DataModelConnection) SetSecondCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DataModelConnection_SetSecondCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *DataModelConnection) ToExternalConnection() *ExternalConnection {
	parentClass := &ExternalConnection{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDataModelConnection(datamodelconnection *DataModelConnection){
	runtime.SetFinalizer(datamodelconnection, nil)
	C.Delete_CObject(C.CString("Delete_DataModelConnection"),datamodelconnection.ptr)
	datamodelconnection.ptr = nil
}

// Class DBConnection 

// Specifies all properties associated with an ODBC or OLE DB external data connection.
type DBConnection struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ExternalConnection 
func NewDBConnection(src *ExternalConnection) ( *DBConnection, error) {
	dbconnection := &DBConnection{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_DBConnection"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		dbconnection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(dbconnection, DeleteDBConnection)
		return dbconnection, nil
	} else {
		dbconnection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return dbconnection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DBConnection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of this <see cref="ExternalConnection"/> object.
// Returns:
//   int32  
func (instance *DBConnection) GetClassType()  (ExternalConnectionClassType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DBConnection_GetClassType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToExternalConnectionClassType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Returns:
//   string  
func (instance *DBConnection) GetConnectionString()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetConnectionString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetConnectionString(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetConnectionString"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the definition of power query formula.
// Returns:
//   PowerQueryFormula  
func (instance *DBConnection) GetPowerQueryFormula()  (*PowerQueryFormula,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DBConnection_GetPowerQueryFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormula{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormula) 

	return result, nil 
}
// Specifies the OLE DB command type.
// 1. Query specifies a cube name
// 2. Query specifies a SQL statement
// 3. Query specifies a table name
// 4. Query specifies that default information has been given, and it is up to the provider how to interpret.
// 5. Query is against a web based List Data Provider.
// Returns:
//   int32  
func (instance *DBConnection) GetCommandType()  (OLEDBCommandType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DBConnection_GetCommandType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOLEDBCommandType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the OLE DB command type.
// 1. Query specifies a cube name
// 2. Query specifies a SQL statement
// 3. Query specifies a table name
// 4. Query specifies that default information has been given, and it is up to the provider how to interpret.
// 5. Query is against a web based List Data Provider.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DBConnection) SetCommandType(value OLEDBCommandType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DBConnection_SetCommandType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Returns:
//   string  
func (instance *DBConnection) GetCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Returns:
//   string  
func (instance *DBConnection) GetSecondCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetSecondCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetSecondCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetSecondCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the id of the connection.
// Returns:
//   int32  
func (instance *DBConnection) GetId()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DBConnection_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Returns:
//   int32  
func (instance *DBConnection) GetSourceType()  (ConnectionDataSourceType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DBConnection_GetSourceType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToConnectionDataSourceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DBConnection) SetSourceType(value ConnectionDataSourceType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DBConnection_SetSourceType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Returns:
//   string  
func (instance *DBConnection) GetSSOId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetSSOId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetSSOId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetSSOId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Returns:
//   bool  
func (instance *DBConnection) GetSavePassword()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_GetSavePassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetSavePassword(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetSavePassword"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Returns:
//   bool  
func (instance *DBConnection) GetSaveData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_GetSaveData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetSaveData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetSaveData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Returns:
//   bool  
func (instance *DBConnection) GetRefreshOnLoad()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_GetRefreshOnLoad"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetRefreshOnLoad(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetRefreshOnLoad"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Returns:
//   int32  
func (instance *DBConnection) GetReconnectionMethodType()  (ReConnectionMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DBConnection_GetReconnectionMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToReConnectionMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DBConnection) SetReconnectionMethodType(value ReConnectionMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DBConnection_SetReconnectionMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Returns:
//   bool  
func (instance *DBConnection) GetOnlyUseConnectionFile()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_GetOnlyUseConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetOnlyUseConnectionFile(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetOnlyUseConnectionFile"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Returns:
//   string  
func (instance *DBConnection) GetOdcFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetOdcFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetOdcFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetOdcFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Returns:
//   string  
func (instance *DBConnection) GetSourceFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetSourceFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetSourceFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetSourceFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Returns:
//   bool  
func (instance *DBConnection) IsNew()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_IsNew"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetIsNew(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetIsNew"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Returns:
//   string  
func (instance *DBConnection) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Returns:
//   bool  
func (instance *DBConnection) GetKeepAlive()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_GetKeepAlive"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetKeepAlive(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetKeepAlive"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Returns:
//   int32  
func (instance *DBConnection) GetRefreshInternal()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("DBConnection_GetRefreshInternal"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DBConnection) SetRefreshInternal(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("DBConnection_SetRefreshInternal"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the user description for this connection
// Returns:
//   string  
func (instance *DBConnection) GetConnectionDescription()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetConnectionDescription"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the user description for this connection
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DBConnection) SetConnectionDescription(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("DBConnection_SetConnectionDescription"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Returns:
//   bool  
func (instance *DBConnection) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Returns:
//   int32  
func (instance *DBConnection) GetCredentialsMethodType()  (CredentialsMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("DBConnection_GetCredentialsMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCredentialsMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DBConnection) SetCredentialsMethodType(value CredentialsMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("DBConnection_SetCredentialsMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Returns:
//   bool  
func (instance *DBConnection) GetBackgroundRefresh()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("DBConnection_GetBackgroundRefresh"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DBConnection) SetBackgroundRefresh(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("DBConnection_SetBackgroundRefresh"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets <see cref="ConnectionParameterCollection"/> for an ODBC or web query.
// Returns:
//   ConnectionParameterCollection  
func (instance *DBConnection) GetParameters()  (*ConnectionParameterCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("DBConnection_GetParameters"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ConnectionParameterCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteConnectionParameterCollection) 

	return result, nil 
}
// Gets the connection file.
// Returns:
//   string  
func (instance *DBConnection) GetConnectionFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("DBConnection_GetConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}


func (instance *DBConnection) ToExternalConnection() *ExternalConnection {
	parentClass := &ExternalConnection{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteDBConnection(dbconnection *DBConnection){
	runtime.SetFinalizer(dbconnection, nil)
	C.Delete_CObject(C.CString("Delete_DBConnection"),dbconnection.ptr)
	dbconnection.ptr = nil
}

// Class ExternalConnection 

// Specifies an external data connection
type ExternalConnection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ExternalConnection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the id of the connection.
// Returns:
//   int32  
func (instance *ExternalConnection) GetId()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ExternalConnection_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Returns:
//   int32  
func (instance *ExternalConnection) GetSourceType()  (ConnectionDataSourceType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ExternalConnection_GetSourceType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToConnectionDataSourceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ExternalConnection) SetSourceType(value ConnectionDataSourceType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ExternalConnection_SetSourceType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Returns:
//   string  
func (instance *ExternalConnection) GetSSOId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetSSOId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetSSOId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetSSOId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Returns:
//   bool  
func (instance *ExternalConnection) GetSavePassword()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_GetSavePassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetSavePassword(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetSavePassword"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Returns:
//   bool  
func (instance *ExternalConnection) GetSaveData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_GetSaveData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetSaveData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetSaveData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Returns:
//   bool  
func (instance *ExternalConnection) GetRefreshOnLoad()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_GetRefreshOnLoad"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetRefreshOnLoad(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetRefreshOnLoad"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Returns:
//   int32  
func (instance *ExternalConnection) GetReconnectionMethodType()  (ReConnectionMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ExternalConnection_GetReconnectionMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToReConnectionMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ExternalConnection) SetReconnectionMethodType(value ReConnectionMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ExternalConnection_SetReconnectionMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Returns:
//   bool  
func (instance *ExternalConnection) GetOnlyUseConnectionFile()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_GetOnlyUseConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetOnlyUseConnectionFile(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetOnlyUseConnectionFile"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Returns:
//   string  
func (instance *ExternalConnection) GetOdcFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetOdcFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetOdcFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetOdcFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Returns:
//   string  
func (instance *ExternalConnection) GetSourceFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetSourceFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetSourceFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetSourceFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Returns:
//   bool  
func (instance *ExternalConnection) IsNew()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_IsNew"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetIsNew(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetIsNew"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Returns:
//   string  
func (instance *ExternalConnection) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Returns:
//   bool  
func (instance *ExternalConnection) GetKeepAlive()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_GetKeepAlive"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetKeepAlive(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetKeepAlive"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Returns:
//   int32  
func (instance *ExternalConnection) GetRefreshInternal()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ExternalConnection_GetRefreshInternal"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ExternalConnection) SetRefreshInternal(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ExternalConnection_SetRefreshInternal"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the user description for this connection
// Returns:
//   string  
func (instance *ExternalConnection) GetConnectionDescription()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetConnectionDescription"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the user description for this connection
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetConnectionDescription(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetConnectionDescription"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Returns:
//   bool  
func (instance *ExternalConnection) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Returns:
//   int32  
func (instance *ExternalConnection) GetCredentialsMethodType()  (CredentialsMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ExternalConnection_GetCredentialsMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCredentialsMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ExternalConnection) SetCredentialsMethodType(value CredentialsMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ExternalConnection_SetCredentialsMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Returns:
//   bool  
func (instance *ExternalConnection) GetBackgroundRefresh()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnection_GetBackgroundRefresh"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ExternalConnection) SetBackgroundRefresh(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ExternalConnection_SetBackgroundRefresh"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets <see cref="ConnectionParameterCollection"/> for an ODBC or web query.
// Returns:
//   ConnectionParameterCollection  
func (instance *ExternalConnection) GetParameters()  (*ConnectionParameterCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ExternalConnection_GetParameters"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ConnectionParameterCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteConnectionParameterCollection) 

	return result, nil 
}
// Gets the type of this <see cref="ExternalConnection"/> object.
// Returns:
//   int32  
func (instance *ExternalConnection) GetClassType()  (ExternalConnectionClassType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ExternalConnection_GetClassType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToExternalConnectionClassType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the definition of power query formula.
// Returns:
//   PowerQueryFormula  
func (instance *ExternalConnection) GetPowerQueryFormula()  (*PowerQueryFormula,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ExternalConnection_GetPowerQueryFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormula{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormula) 

	return result, nil 
}
// Gets the connection file.
// Returns:
//   string  
func (instance *ExternalConnection) GetConnectionFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Returns:
//   string  
func (instance *ExternalConnection) GetCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the OLE DB command type.
// 1. Query specifies a cube name
// 2. Query specifies a SQL statement
// 3. Query specifies a table name
// 4. Query specifies that default information has been given, and it is up to the provider how to interpret.
// 5. Query is against a web based List Data Provider.
// Returns:
//   int32  
func (instance *ExternalConnection) GetCommandType()  (OLEDBCommandType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ExternalConnection_GetCommandType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOLEDBCommandType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the OLE DB command type.
// 1. Query specifies a cube name
// 2. Query specifies a SQL statement
// 3. Query specifies a table name
// 4. Query specifies that default information has been given, and it is up to the provider how to interpret.
// 5. Query is against a web based List Data Provider.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ExternalConnection) SetCommandType(value OLEDBCommandType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ExternalConnection_SetCommandType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Returns:
//   string  
func (instance *ExternalConnection) GetConnectionString()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetConnectionString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetConnectionString(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetConnectionString"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Returns:
//   string  
func (instance *ExternalConnection) GetSecondCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ExternalConnection_GetSecondCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ExternalConnection) SetSecondCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("ExternalConnection_SetSecondCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteExternalConnection(externalconnection *ExternalConnection){
	runtime.SetFinalizer(externalconnection, nil)
	C.Delete_CObject(C.CString("Delete_ExternalConnection"),externalconnection.ptr)
	externalconnection.ptr = nil
}

// Class ExternalConnectionCollection 

// Specifies the <see cref="ExternalConnection"/> collection
type ExternalConnectionCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ExternalConnectionCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ExternalConnectionCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="ExternalConnection"/> element at the specified index.
// Parameters:
//   index - int32 
// Returns:
//   ExternalConnection  
func (instance *ExternalConnectionCollection) Get_Int(index int32)  (*ExternalConnection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("ExternalConnectionCollection_Get_Integer"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ExternalConnection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteExternalConnection) 

	return result, nil 
}
// Gets the <see cref="ExternalConnection"/> element with the specified name.
// Parameters:
//   connectionName - string 
// Returns:
//   ExternalConnection  
func (instance *ExternalConnectionCollection) Get_String(connectionname string)  (*ExternalConnection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBU(C.CString("ExternalConnectionCollection_Get_String"), instance.ptr, C.CString(connectionname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ExternalConnection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteExternalConnection) 

	return result, nil 
}
// Gets the <see cref="ExternalConnection"/> element with the specified id.
// Parameters:
//   connId - int32 
// Returns:
//   ExternalConnection  
func (instance *ExternalConnectionCollection) GetExternalConnectionById(connid int32)  (*ExternalConnection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("ExternalConnectionCollection_GetExternalConnectionById"), instance.ptr, C.int(connid))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ExternalConnection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteExternalConnection) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *ExternalConnectionCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ExternalConnectionCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteExternalConnectionCollection(externalconnectioncollection *ExternalConnectionCollection){
	runtime.SetFinalizer(externalconnectioncollection, nil)
	C.Delete_CObject(C.CString("Delete_ExternalConnectionCollection"),externalconnectioncollection.ptr)
	externalconnectioncollection.ptr = nil
}

// Class WebQueryConnection 

// Specifies the properties for a web query source. A web query will retrieve data from HTML tables,
// and can also supply HTTP "Get" parameters to be processed by the web server in generating the HTML by
// including the parameters and parameter elements.
type WebQueryConnection struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ExternalConnection 
func NewWebQueryConnection(src *ExternalConnection) ( *WebQueryConnection, error) {
	webqueryconnection := &WebQueryConnection{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_WebQueryConnection"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		webqueryconnection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(webqueryconnection, DeleteWebQueryConnection)
		return webqueryconnection, nil
	} else {
		webqueryconnection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return webqueryconnection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of this <see cref="ExternalConnection"/> object.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetClassType()  (ExternalConnectionClassType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebQueryConnection_GetClassType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToExternalConnectionClassType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// true if the web query source is XML (versus HTML), otherwise false.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsXml()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsXml"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// true if the web query source is XML (versus HTML), otherwise false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsXml(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsXml"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This flag exists for backward compatibility with older existing spreadsheet files, and is set
// to true if this web query was created in Microsoft Excel 97.
// This is an optional attribute that can be ignored.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsXl97()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsXl97"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// This flag exists for backward compatibility with older existing spreadsheet files, and is set
// to true if this web query was created in Microsoft Excel 97.
// This is an optional attribute that can be ignored.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsXl97(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsXl97"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This flag exists for backward compatibility with older existing spreadsheet files, and is set
// to true if this web query was refreshed in a spreadsheet application newer than or equal
// to Microsoft Excel 2000.
// This is an optional attribute that can be ignored.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsXl2000()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsXl2000"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// This flag exists for backward compatibility with older existing spreadsheet files, and is set
// to true if this web query was refreshed in a spreadsheet application newer than or equal
// to Microsoft Excel 2000.
// This is an optional attribute that can be ignored.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsXl2000(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsXl2000"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// URL to use to refresh external data.
// Returns:
//   string  
func (instance *WebQueryConnection) GetUrl()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetUrl"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// URL to use to refresh external data.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetUrl(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetUrl"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the connection file.
// Returns:
//   string  
func (instance *WebQueryConnection) GetConnectionFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether dates should be imported into cells in the worksheet as text rather than dates.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsTextDates()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsTextDates"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether dates should be imported into cells in the worksheet as text rather than dates.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsTextDates(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsTextDates"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Flag indicating that XML source data should be imported instead of the HTML table itself.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsXmlSourceData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsXmlSourceData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating that XML source data should be imported instead of the HTML table itself.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsXmlSourceData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsXmlSourceData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns or sets the string used with the post method of inputting data into a web server
// to return data from a web query.
// Returns:
//   string  
func (instance *WebQueryConnection) GetPost()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetPost"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns or sets the string used with the post method of inputting data into a web server
// to return data from a web query.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetPost(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetPost"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Flag indicating whether data contained within HTML PRE tags in the web page is
// parsed into columns when you import the page into a query table.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsParsePre()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsParsePre"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether data contained within HTML PRE tags in the web page is
// parsed into columns when you import the page into a query table.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsParsePre(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsParsePre"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Flag indicating whether web queries should only work on HTML tables.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsHtmlTables()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsHtmlTables"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether web queries should only work on HTML tables.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsHtmlTables(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsHtmlTables"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// How to handle formatting from the HTML source when bringing web query data into the
// worksheet. Relevant when sourceData is True.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetHtmlFormat()  (HtmlFormatHandlingType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebQueryConnection_GetHtmlFormat"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToHtmlFormatHandlingType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// How to handle formatting from the HTML source when bringing web query data into the
// worksheet. Relevant when sourceData is True.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebQueryConnection) SetHtmlFormat(value HtmlFormatHandlingType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("WebQueryConnection_SetHtmlFormat"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Flag indicating whether to parse all tables inside a PRE block with the same width settings
// as the first row.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsSameSettings()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsSameSettings"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether to parse all tables inside a PRE block with the same width settings
// as the first row.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsSameSettings(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsSameSettings"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The URL of the user-facing web page showing the web query data. This URL is persisted
// in the case that sourceData="true" and url has been redirected to reference an XML file.
// Then the user-facing page can be shown in the UI, and the XML data can be retrieved
// behind the scenes.
// Returns:
//   string  
func (instance *WebQueryConnection) GetEditWebPage()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetEditWebPage"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The URL of the user-facing web page showing the web query data. This URL is persisted
// in the case that sourceData="true" and url has been redirected to reference an XML file.
// Then the user-facing page can be shown in the UI, and the XML data can be retrieved
// behind the scenes.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetEditWebPage(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetEditWebPage"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Flag indicating whether consecutive delimiters should be treated as just one delimiter.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsConsecutive()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsConsecutive"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Flag indicating whether consecutive delimiters should be treated as just one delimiter.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsConsecutive(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsConsecutive"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the id of the connection.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetId()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebQueryConnection_GetId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetSourceType()  (ConnectionDataSourceType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebQueryConnection_GetSourceType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToConnectionDataSourceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or Sets the external connection DataSource type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebQueryConnection) SetSourceType(value ConnectionDataSourceType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("WebQueryConnection_SetSourceType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Returns:
//   string  
func (instance *WebQueryConnection) GetSSOId()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetSSOId"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Identifier for Single Sign On (SSO) used for authentication between an intermediate
// spreadsheetML server and the external data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetSSOId(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetSSOId"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Returns:
//   bool  
func (instance *WebQueryConnection) GetSavePassword()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_GetSavePassword"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the password is to be saved as part of the connection string; otherwise, False.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetSavePassword(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetSavePassword"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Returns:
//   bool  
func (instance *WebQueryConnection) GetSaveData()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_GetSaveData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the external data fetched over the connection to populate a table is to be saved
// with the workbook; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetSaveData(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetSaveData"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Returns:
//   bool  
func (instance *WebQueryConnection) GetRefreshOnLoad()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_GetRefreshOnLoad"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if this connection should be refreshed when opening the file; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetRefreshOnLoad(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetRefreshOnLoad"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetReconnectionMethodType()  (ReConnectionMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebQueryConnection_GetReconnectionMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToReConnectionMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies what the spreadsheet application should do when a connection fails.
// The default value is ReConnectionMethodType.Required.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebQueryConnection) SetReconnectionMethodType(value ReConnectionMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("WebQueryConnection_SetReconnectionMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Returns:
//   bool  
func (instance *WebQueryConnection) GetOnlyUseConnectionFile()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_GetOnlyUseConnectionFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the spreadsheet application should always and only use the
// connection information in the external connection file indicated by the odcFile attribute
// when the connection is refreshed.  If false, then the spreadsheet application
// should follow the procedure indicated by the reconnectionMethod attribute
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetOnlyUseConnectionFile(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetOnlyUseConnectionFile"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Returns:
//   string  
func (instance *WebQueryConnection) GetOdcFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetOdcFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the full path to external connection file from which this connection was
// created. If a connection fails during an attempt to refresh data, and reconnectionMethod=1,
// then the spreadsheet application will try again using information from the external connection file
// instead of the connection object embedded within the workbook.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetOdcFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetOdcFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Returns:
//   string  
func (instance *WebQueryConnection) GetSourceFile()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetSourceFile"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Used when the external data source is file-based.
// When a connection to such a data source fails, the spreadsheet application attempts to connect directly to this file. May be
// expressed in URI or system-specific file path notation.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetSourceFile(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetSourceFile"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsNew()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsNew"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True if the connection has not been refreshed for the first time; otherwise, false.
// This state can happen when the user saves the file before a query has finished returning.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsNew(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsNew"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Returns:
//   string  
func (instance *WebQueryConnection) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the name of the connection. Each connection must have a unique name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Returns:
//   bool  
func (instance *WebQueryConnection) GetKeepAlive()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_GetKeepAlive"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// True when the spreadsheet application should make efforts to keep the connection
// open. When false, the application should close the connection after retrieving the
// information.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetKeepAlive(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetKeepAlive"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetRefreshInternal()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("WebQueryConnection_GetRefreshInternal"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the number of minutes between automatic refreshes of the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebQueryConnection) SetRefreshInternal(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("WebQueryConnection_SetRefreshInternal"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the user description for this connection
// Returns:
//   string  
func (instance *WebQueryConnection) GetConnectionDescription()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetConnectionDescription"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the user description for this connection
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetConnectionDescription(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetConnectionDescription"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Returns:
//   bool  
func (instance *WebQueryConnection) IsDeleted()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_IsDeleted"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the associated workbook connection has been deleted.  true if the
// connection has been deleted; otherwise, false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetIsDeleted(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetIsDeleted"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetCredentialsMethodType()  (CredentialsMethodType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebQueryConnection_GetCredentialsMethodType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCredentialsMethodType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the authentication method to be used when establishing (or re-establishing) the connection.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebQueryConnection) SetCredentialsMethodType(value CredentialsMethodType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("WebQueryConnection_SetCredentialsMethodType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Returns:
//   bool  
func (instance *WebQueryConnection) GetBackgroundRefresh()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("WebQueryConnection_GetBackgroundRefresh"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the connection can be refreshed in the background (asynchronously).
// true if preferred usage of the connection is to refresh asynchronously in the background;
// false if preferred usage of the connection is to refresh synchronously in the foreground.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *WebQueryConnection) SetBackgroundRefresh(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("WebQueryConnection_SetBackgroundRefresh"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets <see cref="ConnectionParameterCollection"/> for an ODBC or web query.
// Returns:
//   ConnectionParameterCollection  
func (instance *WebQueryConnection) GetParameters()  (*ConnectionParameterCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebQueryConnection_GetParameters"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ConnectionParameterCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteConnectionParameterCollection) 

	return result, nil 
}
// Gets the definition of power query formula.
// Returns:
//   PowerQueryFormula  
func (instance *WebQueryConnection) GetPowerQueryFormula()  (*PowerQueryFormula,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("WebQueryConnection_GetPowerQueryFormula"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PowerQueryFormula{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePowerQueryFormula) 

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Returns:
//   string  
func (instance *WebQueryConnection) GetCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The string containing the database command to pass to the data provider API that will
// interact with the external source in order to retrieve data
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the OLE DB command type.
// 1. Query specifies a cube name
// 2. Query specifies a SQL statement
// 3. Query specifies a table name
// 4. Query specifies that default information has been given, and it is up to the provider how to interpret.
// 5. Query is against a web based List Data Provider.
// Returns:
//   int32  
func (instance *WebQueryConnection) GetCommandType()  (OLEDBCommandType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("WebQueryConnection_GetCommandType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToOLEDBCommandType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the OLE DB command type.
// 1. Query specifies a cube name
// 2. Query specifies a SQL statement
// 3. Query specifies a table name
// 4. Query specifies that default information has been given, and it is up to the provider how to interpret.
// 5. Query is against a web based List Data Provider.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *WebQueryConnection) SetCommandType(value OLEDBCommandType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("WebQueryConnection_SetCommandType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Returns:
//   string  
func (instance *WebQueryConnection) GetConnectionString()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetConnectionString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// The connection information string is used to make contact with an OLE DB or ODBC data source.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetConnectionString(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetConnectionString"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Returns:
//   string  
func (instance *WebQueryConnection) GetSecondCommand()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("WebQueryConnection_GetSecondCommand"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a second command text string that is persisted when PivotTable server-based
// page fields are in use.
// For ODBC connections, serverCommand is usually a broader query than command (no
// WHERE clause is present in the former). Based on these 2 commands(Command and ServerCommand),
// parameter UI can be populated and parameterized queries can be constructed
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *WebQueryConnection) SetSecondCommand(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("WebQueryConnection_SetSecondCommand"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *WebQueryConnection) ToExternalConnection() *ExternalConnection {
	parentClass := &ExternalConnection{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteWebQueryConnection(webqueryconnection *WebQueryConnection){
	runtime.SetFinalizer(webqueryconnection, nil)
	C.Delete_CObject(C.CString("Delete_WebQueryConnection"),webqueryconnection.ptr)
	webqueryconnection.ptr = nil
}
