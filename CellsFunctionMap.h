
// Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
// Powered by Aspose.Cells.

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#ifdef _WIN32
#include <guiddef.h>
#else
#include <string.h>
#endif

#define ASPOSE_CELLS_LIBRARY

#ifdef _WIN32
	#ifdef ASPOSE_CELLS_LIBRARY
		#define ASPOSE_CELLS_API __declspec(dllexport)
		#define ASPOSE_CELLS_API2
	#else
		#define ASPOSE_CELLS_API __declspec(dllimport)
		#define ASPOSE_CELLS_API2
	#endif
	#define ASPOSE_CELLS_API3
#else
	#ifdef ASPOSE_CELLS_LIBRARY
		#define ASPOSE_CELLS_API
		#define ASPOSE_CELLS_API2 __attribute__ ((visibility("default")))
	#else
		#define ASPOSE_CELLS_API
		#define ASPOSE_CELLS_API2
	#endif
	#define ASPOSE_CELLS_API3 __attribute__ ((visibility("default")))
#endif

#ifndef ASPOSE_CELLS_CGO
#define ASPOSE_CELLS_CGO

#ifdef __cplusplus
extern "C" {
#endif

#ifdef _WIN32
    ASPOSE_CELLS_API typedef struct {
        int8_t return_value;
        int error_no;
        char* error_message;
    } c_return_int8_value;
    ASPOSE_CELLS_API typedef struct {
        int16_t return_value;
        int error_no;
        char* error_message;
    } c_return_int16_value;
    ASPOSE_CELLS_API typedef struct {
        int return_value;
        int error_no;
        char* error_message;
    } c_return_int_value;
    ASPOSE_CELLS_API typedef struct {
        uint16_t return_value;
        int error_no;
        char* error_message;
    } c_return_uint16_value;
    ASPOSE_CELLS_API typedef struct {
        uint32_t return_value;
        int error_no;
        char* error_message;
    } c_return_uint_value;
    ASPOSE_CELLS_API typedef struct {
        uint64_t return_value;
        int error_no;
        char* error_message;
    } c_return_ulong_value;
    ASPOSE_CELLS_API typedef struct {
        bool return_value;
        int error_no;
        char* error_message;
    } c_return_bool_value;
    ASPOSE_CELLS_API typedef struct {
        char return_value;
        int error_no;
        char* error_message;
    } c_return_char_value;
    ASPOSE_CELLS_API typedef struct {
        float return_value;
        int error_no;
        char* error_message;
    } c_return_float_value;    
    ASPOSE_CELLS_API typedef struct {
        double return_value;
        int error_no;
        char* error_message;
    } c_return_double_value;
    ASPOSE_CELLS_API typedef struct {
        long return_value;
        int error_no;
        char* error_message;
    } c_return_long_value;

    ASPOSE_CELLS_API typedef struct {
        int error_no;
        char* error_message;
    } c_return_void_value;

    ASPOSE_CELLS_API typedef struct {
        char* return_value;
        int data_length;
        int error_no;
        char* error_message;
    } c_return_string_value;

    ASPOSE_CELLS_API typedef struct {
        void* return_value;
        int column_length;
        int row_length;
        size_t size;
        int error_no;
        char* error_message;
    } c_return_ptr_value;
#else
    typedef struct {
        int8_t return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_int8_value;
    typedef struct {
        int16_t return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_int16_value;
    typedef struct {
        int return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_int_value;
    typedef struct {
        uint16_t return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_uint16_value;
    typedef struct {
        uint32_t return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_uint_value;
    typedef struct {
        uint64_t return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_ulong_value;
    typedef struct {
        bool return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_bool_value;
    typedef struct {
        char return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_char_value;
    typedef struct {
        float return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_float_value;    
    typedef struct {
        double return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_double_value;
    typedef struct {
        long return_value;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_long_value;

    typedef struct {
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_void_value;

    typedef struct {
        char* return_value;
        int data_length;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_string_value;

    typedef struct {
        void* return_value;
        int column_length;
        int row_length;
        size_t size;
        int error_no;
        char* error_message;
    } ASPOSE_CELLS_API c_return_ptr_value;

#endif

ASPOSE_CELLS_API void Startup();
/*****************Class Stream ****************/
ASPOSE_CELLS_API void* Get_Date(int year, int month, int day, int hour, int minute, int second);
ASPOSE_CELLS_API void Delete_GetDate(void* date);
ASPOSE_CELLS_API c_return_ptr_value* New_Object_Null();

/**************Enumerator ReferredAreaEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* ReferredAreaEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* ReferredAreaEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* ReferredAreaEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_ReferredAreaEnumerator(void* instance_ptr);

/**************Enumerator CellEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* CellEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* CellEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* CellEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_CellEnumerator(void* instance_ptr);

/**************Enumerator ExternalLinkEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* ExternalLinkEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* ExternalLinkEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* ExternalLinkEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_ExternalLinkEnumerator(void* instance_ptr);

/**************Enumerator RowEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* RowEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* RowEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* RowEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_RowEnumerator(void* instance_ptr);

/**************Enumerator TextParagraphEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* TextParagraphEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* TextParagraphEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* TextParagraphEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_TextParagraphEnumerator(void* instance_ptr);

/**************Enumerator ChartPointEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* ChartPointEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* ChartPointEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* ChartPointEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_ChartPointEnumerator(void* instance_ptr);

/**************Enumerator PivotFieldEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* PivotFieldEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* PivotFieldEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* PivotFieldEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_PivotFieldEnumerator(void* instance_ptr);

/**************Enumerator PivotItemEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* PivotItemEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* PivotItemEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* PivotItemEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_PivotItemEnumerator(void* instance_ptr);

/**************Enumerator DigitalSignatureEnumerator *****************/
ASPOSE_CELLS_API c_return_ptr_value* DigitalSignatureEnumerator_GetCurrent(void* instance_ptr);
ASPOSE_CELLS_API c_return_bool_value* DigitalSignatureEnumerator_MoveNext(void* instance_ptr);
ASPOSE_CELLS_API c_return_void_value* DigitalSignatureEnumerator_Reset(void* instance_ptr);
ASPOSE_CELLS_API void Delete_DigitalSignatureEnumerator(void* instance_ptr);

/***************************************************************/
ASPOSE_CELLS_API c_return_void_value* Cell_PutValue_Null(void* instance_ptr);


ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZZA(const char *name  );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZZB(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZC(const char *name ,void* param1,bool param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZZD(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZE(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZF(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZZG(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZH(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZI(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZZJ(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZZK(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZL(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZZM(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZN(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZZO(const char *name ,void* param1,bool param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZP(const char *name ,void* param1,int32_t param2,char* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZQ(const char *name ,void* param1,int32_t param2,int param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,int32_t param8,int32_t param9 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZR(const char *name ,void* param1,int32_t param2,bool param3,bool param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZS(const char *name ,void* param1,int32_t param2,int param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZT(const char *name ,void* param1,int32_t param2,void* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZU(const char *name ,void* param1,int32_t param2,int param3,void* param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZV(const char *name ,void* param1,int32_t param2,int param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZW(const char *name ,void* param1,int32_t param2,int param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZZX(const char *name ,void* param1,int32_t param2,int param3,void* param4,bool param5,int param6,void* param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZZY(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAA(const char *name ,void* param1,bool param2 );
ASPOSE_CELLS_API c_return_double_value* CellsGoFunctoinZZAB(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAC(const char *name ,void* param1,double param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAD(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAE(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAF(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAG(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAH(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZAI(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAJ(const char *name ,void* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAK(const char *name ,void* param1,char* param2,bool param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAL(const char *name ,void* param1,char* param2,bool param3 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZAM(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_float_value* CellsGoFunctoinZZAN(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAO(const char *name ,void* param1,void* param2,bool param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAP(const char *name ,void* param1,void* param2,void* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAQ(const char *name ,void* param1,char* param2,void* param3 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZAR(const char *name ,void* param1,bool param2,bool param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAS(const char *name ,void* param1,char* param2,void* param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAT(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAU(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZAV(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,void* param5,void* param6,int32_t param7,int32_t param8 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAW(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAX(const char *name ,void* param1,bool param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZAY(const char *name ,void* param1,char* param2,void* param3,bool param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBA(const char *name ,void* param1,char* param2,void* param3,void* param4,int32_t param5,int32_t param6,bool param7,bool param8 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBB(const char *name ,void* param1,char* param2,void* param3,void* param4,int32_t param5,int32_t param6,bool param7,bool param8,void* param9 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBC(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4,char* param5,void* param6,int32_t param7,int32_t param8 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBD(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4,bool param5,void* param6,int32_t param7,int32_t param8 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBE(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,void* param8,int32_t param9,int32_t param10 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBF(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,bool param6,void* param7,int32_t param8,int32_t param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBG(const char *name ,void* param1,int32_t param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBH(const char *name ,void* param1,char* param2,char* param3,void* param4 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZBI(const char *name ,void* param1,bool param2 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZBJ(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBK(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBL(const char *name ,void* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBM(const char *name ,int32_t param1,int32_t param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBN(const char *name ,char* param1,char* param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZBO(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBP(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_long_value* CellsGoFunctoinZZBQ(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZBR(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBS(const char *name ,void* param1,char* param2,char* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBT(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBU(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZBV(const char *name ,void* param1,int32_t param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBW(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5,bool param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZBX(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZBY(const char *name ,void* param1,char* param2,char* param3,bool param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCA(const char *name ,void* param1,void* param2,int32_t param3,char* param4,bool param5,int32_t param6,int32_t param7 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCB(const char *name ,void* param1,char* param2,void* param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCC(const char *name ,void* param1,void* param2,int32_t param3,void* param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCD(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCE(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,bool param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCF(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,bool param6,bool param7 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCG(const char *name ,void* param1,int32_t param2,double param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCH(const char *name ,void* param1,int32_t param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCI(const char *name ,void* param1,int32_t param2,int32_t param3,double param4 );
ASPOSE_CELLS_API c_return_double_value* CellsGoFunctoinZZCJ(const char *name ,void* param1,int32_t param2,bool param3,int param4 );
ASPOSE_CELLS_API c_return_double_value* CellsGoFunctoinZZCK(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZCL(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCM(const char *name ,void* param1,int32_t param2,void* param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCN(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5,void* param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCO(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCP(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCQ(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCR(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5,void* param6,void* param7 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCS(const char *name ,void* param1,bool param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCT(const char *name ,void* param1,int32_t param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCU(const char *name ,void* param1,int32_t param2,bool param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZCV(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZCW(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZCX(const char *name ,void* param1,int32_t param2,int32_t param3 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZCY(const char *name ,void* param1,int32_t param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZDA(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDB(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,char* param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDC(const char *name ,void* param1,void* param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDD(const char *name ,void* param1,void* param2,void* param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDE(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDF(const char *name ,void* param1,void* param2,int32_t param3,int param4,bool param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDG(const char *name ,void* param1,void* param2,int param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDH(const char *name ,void* param1,void* param2,int32_t param3,int param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDI(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDJ(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDK(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDL(const char *name ,void* param1,void* param2,int32_t param3,int param4,void* param5,int32_t param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDM(const char *name ,void* param1,void* param2,int32_t param3,int param4,void* param5,int32_t param6,bool param7,bool param8,bool param9 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDN(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,bool param6,void* param7,int32_t param8 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDO(const char *name ,void* param1,bool param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDP(const char *name ,void* param1,int32_t param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZDQ(const char *name ,void* param1,int32_t param2,int32_t param3,int param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZDR(const char *name  );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDS(const char *name ,int param1 );
ASPOSE_CELLS_API c_return_double_value* CellsGoFunctoinZZDT(const char *name  );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDU(const char *name ,double param1 );
ASPOSE_CELLS_API c_return_double_value* CellsGoFunctoinZZDV(const char *name ,char* param1,void* param2,double param3 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZDW(const char *name  );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZDX(const char *name ,char* param1,int32_t* param2,int32_t* param3 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZDY(const char *name ,int32_t param1,int32_t param2 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZEA(const char *name ,int32_t param1 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEB(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZEC(const char *name ,double param1,bool param2 );
ASPOSE_CELLS_API c_return_double_value* CellsGoFunctoinZZED(const char *name ,void* param1,bool param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZEE(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZEF(const char *name ,void* param1,int32_t param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZEG(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZEH(const char *name ,char* param1,char param2 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZEI(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZEJ(const char *name  );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZEK(const char *name ,bool param1 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEL(const char *name ,void* param1,int32_t param2,int32_t param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEM(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZEN(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEO(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZEP(const char *name ,void* param1,uint8_t param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZEQ(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4,void* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZER(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4,void* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZES(const char *name ,void* param1,char* param2,char* param3,void* param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZET(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZEU(const char *name ,int param1,int32_t param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEV(const char *name ,void* param1,int param2,int32_t param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEW(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZEX(const char *name ,void* param1,int param2,char* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZEY(const char *name ,void* param1,int param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFA(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFB(const char *name ,void* param1,void* param2,void* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZFC(const char *name ,void* param1,int32_t param2,int param3,char* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZFD(const char *name ,void* param1,int32_t param2,int param3,int param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZFE(const char *name ,void* param1,int32_t param2,int param3,int param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZFF(const char *name ,void* param1,int32_t param2,int param3,void* param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFG(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFH(const char *name ,void* param1,void* param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFI(const char *name ,int param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFJ(const char *name ,bool param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFK(const char *name ,int param1 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZFL(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZFM(const char *name ,void* param1,int param2,bool param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZFN(const char *name ,void* param1,char* param2,char* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZFO(const char *name ,void* param1,char* param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZFP(const char *name ,void* param1,int param2,char* param3,void* param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFQ(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFR(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFS(const char *name ,void* param1,int32_t param2,char* param3 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZFT(const char *name ,void* param1,int32_t param2,char* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZFU(const char *name ,int param1 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZFV(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZFW(const char *name ,int param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFX(const char *name ,char* param1,bool param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZFY(const char *name ,char* param1,bool param2,bool param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGA(const char *name ,char* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZGB(const char *name ,char* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGC(const char *name ,char* param1,bool param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGD(const char *name ,void* param1,int32_t param2,bool param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGE(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZGF(const char *name  );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZGG(const char *name ,int32_t param1,int32_t param2,void* param3 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZGH(const char *name ,void* param1,bool param2,bool param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGI(const char *name ,void* param1,char* param2,char* param3,bool param4,bool param5 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZGJ(const char *name ,void* param1,int32_t param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZGK(const char *name ,void* param1,void* param2,int param3,int param4,char* param5,char* param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGL(const char *name ,void* param1,int param2,int param3,char* param4,char* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGM(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZGN(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_char_value* CellsGoFunctoinZZGO(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGP(const char *name ,void* param1,char* param2,char* param3,bool param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGQ(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGR(const char *name ,void* param1,char* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGS(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGT(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,char* param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGU(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,char* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGV(const char *name ,void* param1,char* param2,char* param3,char* param4,char* param5,char* param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGW(const char *name ,void* param1,char* param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZGX(const char *name ,void* param1,void* param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZGY(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZHA(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZHB(const char *name ,void* param1,char param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZHC(const char *name ,void* param1,int param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHD(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHE(const char *name ,void* param1,int param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHF(const char *name ,int8_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHG(const char *name ,uint8_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHH(const char *name ,int16_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHI(const char *name ,uint16_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHJ(const char *name ,int32_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHK(const char *name ,uint32_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHL(const char *name ,int64_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHM(const char *name ,uint64_t param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHN(const char *name ,float param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHO(const char *name ,double param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHP(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHQ(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHR(const char *name ,void* param1,int32_t param2,int32_t param3 );
ASPOSE_CELLS_API c_return_int8_value* CellsGoFunctoinZZHS(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_int16_value* CellsGoFunctoinZZHT(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_uint16_value* CellsGoFunctoinZZHU(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_uint_value* CellsGoFunctoinZZHV(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_ulong_value* CellsGoFunctoinZZHW(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZHX(const char *name ,void* param1,double param2,double param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZHY(const char *name ,void* param1,int32_t param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIA(const char *name ,void* param1,bool param2,bool param3,bool param4,int32_t param5,void* param6,int32_t param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIB(const char *name ,void* param1,bool param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIC(const char *name ,void* param1,bool param2,bool param3,bool param4,int32_t param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZID(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIE(const char *name ,void* param1,char* param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIF(const char *name ,void* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIG(const char *name ,void* param1,int param2,void* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIH(const char *name ,void* param1,void* param2,int32_t param3,void* param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZII(const char *name ,void* param1,int param2,int param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIJ(const char *name ,void* param1,int param2,int param3,void* param4 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZIK(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIL(const char *name ,void* param1,bool param2,bool param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZIM(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIN(const char *name ,void* param1,int param2,char* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIO(const char *name ,void* param1,bool param2,char* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIP(const char *name ,void* param1,char* param2,char* param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIQ(const char *name ,void* param1,int param2,void* param3,void* param4 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZIR(const char *name ,void* param1,int param2,int param3,void* param4 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZIS(const char *name ,void* param1,int param2,int param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIT(const char *name ,void* param1,void* param2,void* param3,int param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIU(const char *name ,int param1,double param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZIV(const char *name ,void* param1,char* param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZIW(const char *name ,void* param1,char* param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZIX(const char *name ,void* param1,void* param2,int param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZIY(const char *name ,void* param1,void* param2,bool param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJA(const char *name ,void* param1,void* param2,int32_t param3,bool param4,bool param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZJB(const char *name ,char* param1,void* param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZJC(const char *name ,void* param1,int32_t param2,void* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJD(const char *name ,void* param1,char* param2,int param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZJE(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZJF(const char *name ,void* param1,char* param2,char* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZJG(const char *name ,void* param1,char* param2,double param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZJH(const char *name ,void* param1,char* param2,void* param3,int32_t param4,bool param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZJI(const char *name ,void* param1,bool param2,void* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZJJ(const char *name ,void* param1,int32_t param2,void* param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJK(const char *name ,void* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZJL(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJM(const char *name ,void* param1,bool param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZJN(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJO(const char *name ,void* param1,int param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZJP(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJQ(const char *name ,void* param1,char* param2,char* param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJR(const char *name ,void* param1,void* param2,int32_t param3,char* param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZJS(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZJT(const char *name ,void* param1,int32_t* param2,int32_t* param3,int32_t* param4,int32_t* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJU(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,void* param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJV(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJW(const char *name ,void* param1,bool param2,char* param3,char* param4,char* param5,bool param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJX(const char *name ,void* param1,bool param2,bool param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZJY(const char *name ,void* param1,int param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKA(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZKB(const char *name ,void* param1,char* param2,bool param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKC(const char *name ,void* param1,char* param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKD(const char *name ,void* param1,char* param2,void* param3,void* param4,int32_t param5,int32_t param6,void* param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKE(const char *name ,void* param1,char* param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKF(const char *name ,void* param1,char* param2,void* param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKG(const char *name ,void* param1,char* param2,void* param3,void* param4,int32_t param5,int32_t param6,int32_t param7,int32_t param8,void* param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKH(const char *name ,void* param1,char* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKI(const char *name ,void* param1,int32_t param2,int param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKJ(const char *name ,void* param1,int32_t param2,int param3,char* param4 );
ASPOSE_CELLS_API c_return_string_value* CellsGoFunctoinZZKK(const char *name ,void* param1,int32_t param2,char* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKL(const char *name ,void* param1,char* param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZKM(const char *name ,void* param1,char* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKN(const char *name ,char* param1,char* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKO(const char *name ,char* param1,char* param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKP(const char *name ,void* param1,void* param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKQ(const char *name ,void* param1,void* param2,char* param3,char* param4,char* param5,int param6 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKR(const char *name ,void* param1,int32_t param2,char* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKS(const char *name ,int param1,void* param2,void* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZKT(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZKU(const char *name ,int32_t param1 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZKV(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKW(const char *name ,void* param1,void* param2,double param3,int param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKX(const char *name ,void* param1,void* param2,double param3,void* param4,double param5,int param6,int32_t param7 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZKY(const char *name ,void* param1,int param2,int param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLA(const char *name ,void* param1,int param2,double param3,int param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLB(const char *name ,void* param1,float param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLC(const char *name ,void* param1,int param2,int param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLD(const char *name ,void* param1,double param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLE(const char *name ,void* param1,double param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLF(const char *name ,void* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLG(const char *name ,void* param1,bool param2,void* param3,int32_t param4,char* param5,bool param6,char* param7 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZLH(const char *name ,void* param1,bool param2,void* param3,int32_t param4,char* param5,bool param6,char* param7,bool param8 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZLI(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,void* param6,int32_t param7 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZLJ(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,void* param6,int32_t param7,char* param8 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZLK(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4,int32_t param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZLL(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4,int32_t param5,int32_t param6,int32_t param7 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZLM(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLN(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLO(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLP(const char *name ,void* param1,int param2,char* param3,char* param4,int32_t param5,bool param6,bool param7,int32_t param8,int32_t param9,int32_t param10,int32_t param11 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLQ(const char *name ,void* param1,int param2,char* param3,char* param4,int32_t param5,bool param6,bool param7,int32_t param8,int32_t param9,int32_t param10,int32_t param11,int32_t param12,int32_t param13 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLR(const char *name ,void* param1,int param2,char* param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,int32_t param8,int32_t param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLS(const char *name ,void* param1,int param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,void* param7,int32_t param8,bool param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLT(const char *name ,void* param1,int param2,int param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,void* param8,int32_t param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLU(const char *name ,void* param1,int param2,int param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLV(const char *name ,void* param1,int param2,int param3,double param4,double param5,double param6,double param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLW(const char *name ,void* param1,int param2,int param3,double param4,double param5,double param6,double param7,void* param8,int32_t param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLX(const char *name ,void* param1,int param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,int32_t param8 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZLY(const char *name ,void* param1,int param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMA(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,void* param6,int32_t param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMB(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4,int32_t param5,int32_t param6,int32_t param7 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMC(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,void* param8,int32_t param9,void* param10,int32_t param11 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMD(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,char* param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZME(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,void* param8,int32_t param9 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMF(const char *name ,void* param1,void* param2,void* param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMG(const char *name ,void* param1,void* param2,void* param3,int32_t param4,int32_t param5,bool param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMH(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6,int32_t param7,void* param8,int32_t param9 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMI(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMJ(const char *name ,void* param1,float param2,float param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMK(const char *name ,void* param1,float param2,float param3,float param4,float param5,float param6,float param7 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZML(const char *name ,void* param1,float param2,float param3,float param4,float param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMM(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZMN(const char *name ,void* param1,int param2,double param3 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZMO(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMP(const char *name ,void* param1,int16_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMQ(const char *name ,void* param1,int64_t param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMR(const char *name ,void* param1,char* param2,float param3,float param4,int param5,int param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZMS(const char *name ,void* param1,float param2,float param3,int param4,int param5 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZMT(const char *name ,void* param1,int param2,bool param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZMU(const char *name ,void* param1,int param2,int32_t param3,int32_t param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZMV(const char *name ,void* param1,void* param2,int32_t param3,char* param4,bool param5,int32_t param6,int32_t param7,int32_t param8,int32_t param9 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZMW(const char *name ,void* param1,int param2,char* param3,bool param4,int32_t param5,int32_t param6,int32_t param7,int32_t param8 );
ASPOSE_CELLS_API c_return_float_value* CellsGoFunctoinZZMX(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZMY(const char *name ,void* param1,int32_t param2,char* param3,bool param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZNA(const char *name ,void* param1,char* param2,bool param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZNB(const char *name ,void* param1,char* param2,bool param3,bool param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNC(const char *name ,void* param1,char* param2,bool param3,void* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZND(const char *name ,void* param1,int param2,char* param3,bool param4,void* param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNE(const char *name ,void* param1,char* param2,void* param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNF(const char *name ,void* param1,char* param2,bool param3 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNG(const char *name ,void* param1,char* param2,double param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNH(const char *name ,char* param1,void* param2,char* param3,void* param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNI(const char *name ,char* param1,void* param2,int32_t param3,int32_t param4,void* param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNJ(const char *name ,char* param1,void* param2,int32_t param3 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNK(const char *name ,void* param1,int param2,int32_t param3,int param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNL(const char *name ,void* param1,int32_t param2,int32_t param3,int param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNM(const char *name ,void* param1,double param2,bool param3 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZNN(const char *name ,void* param1,void* param2,void* param3,void* param4,int32_t param5,double param6,bool param7 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZNO(const char *name ,void* param1,double param2,double param3,double param4,bool param5 );
ASPOSE_CELLS_API c_return_bool_value* CellsGoFunctoinZZNP(const char *name ,void* param1,void* param2,int32_t param3,bool param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNQ(const char *name ,void* param1,int32_t param2,int param3,bool param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNR(const char *name ,void* param1,int32_t param2,int param3,double param4,double param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNS(const char *name ,void* param1,int param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNT(const char *name ,void* param1,int param2,void* param3,void* param4 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNU(const char *name ,void* param1,int param2,int32_t param3,int param4,char* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZNV(const char *name ,void* param1,int param2,int32_t param3,int param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNW(const char *name ,void* param1,int32_t param2,int32_t param3,int param4,bool param5,int32_t param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNX(const char *name ,void* param1,int32_t param2,int32_t param3,int param4,double param5,double param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZNY(const char *name ,void* param1,int32_t param2,int param3,char* param4,char* param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOA(const char *name ,void* param1,int32_t param2,int param3,void* param4,void* param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZOB(const char *name ,void* param1,int32_t param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOC(const char *name ,void* param1,int param2,void* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOD(const char *name ,void* param1,void* param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZOE(const char *name ,void* param1,int32_t param2,int32_t param3,bool param4,int32_t param5,int32_t param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOF(const char *name ,void* param1,char* param2,char* param3,char* param4,bool param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOG(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,char* param5,bool param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOH(const char *name ,void* param1,char* param2,int32_t param3,int32_t param4,char* param5,bool param6,bool param7 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOI(const char *name ,void* param1,char* param2,char* param3,char* param4,bool param5,bool param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOJ(const char *name ,void* param1,void* param2,char* param3,char* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOK(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,char* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOL(const char *name ,void* param1,void* param2,int32_t param3,bool param4,void* param5,char* param6,char* param7 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOM(const char *name ,void* param1,void* param2,int32_t param3,bool param4,void* param5,int32_t param6,int32_t param7,char* param8 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZON(const char *name ,void* param1,int param2,int32_t param3,int param4,int param5,bool param6,bool param7,void* param8 );
ASPOSE_CELLS_API c_return_float_value* CellsGoFunctoinZZOO(const char *name ,void* param1,int param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZOP(const char *name ,void* param1,void* param2,float param3,float param4,float param5,float param6 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOQ(const char *name ,char* param1,float param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOR(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOS(const char *name ,void* param1,int32_t param2 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZOT(const char *name ,void* param1,void* param2,int32_t param3,void* param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOU(const char *name ,void* param1,int32_t param2,char* param3,char* param4,void* param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOV(const char *name ,char* param1,char* param2,char* param3,void* param4 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZOW(const char *name ,bool param1,bool param2 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZOX(const char *name ,void* param1,int32_t param2,int32_t param3,void* param4,bool param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZOY(const char *name ,void* param1,int32_t param2,int32_t param3,char* param4,bool param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPA(const char *name ,void* param1,int32_t param2,int32_t param3,int32_t param4,int32_t param5,bool param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPB(const char *name ,void* param1,int32_t param2,int32_t param3,int param4,char* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPC(const char *name ,void* param1,char* param2,bool param3,int param4,char* param5,int param6 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPD(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPE(const char *name ,void* param1,void* param2,char* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPF(const char *name ,void* param1,void* param2,int32_t param3,int32_t param4,void* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPG(const char *name ,void* param1,void* param2,char* param3,void* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPH(const char *name ,void* param1,void* param2,int32_t param3,char* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPI(const char *name ,void* param1,void* param2,void* param3,char* param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPJ(const char *name ,void* param1,void* param2,void* param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZPK(const char *name ,void* param1,char* param2,void* param3,int32_t param4 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPL(const char *name ,void* param1,void* param2,char* param3 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPM(const char *name ,void* param1,char* param2,char* param3,void* param4,int32_t param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPN(const char *name ,void* param1,char* param2,char* param3,char* param4,char* param5 );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPO(const char *name ,void* param1,char* param2,bool param3,int32_t param4,int32_t param5 );
ASPOSE_CELLS_API c_return_ptr_value* CellsGoFunctoinZZPP(const char *name  );
ASPOSE_CELLS_API c_return_int_value* CellsGoFunctoinZZPQ(const char *name ,void* param1 );
ASPOSE_CELLS_API c_return_void_value* CellsGoFunctoinZZPR(const char *name ,void* param1,uint8_t param2 );
ASPOSE_CELLS_API void Delete_CObject(const char* name , void* instance_ptr);
#ifdef __cplusplus
}
#endif
#endif
