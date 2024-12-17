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

/**************Enum PivotAreaType *****************/
// Indicates the type of rule being used to describe an area or aspect of the PivotTable.
type PivotAreaType int32

const(
// No Pivot area.
PivotAreaType_None PivotAreaType = 0 

// Represents a header or item.
PivotAreaType_Normal PivotAreaType = 1 

// Represents something in the data area.
PivotAreaType_Data PivotAreaType = 2 

// Represents the whole PivotTable.
PivotAreaType_All PivotAreaType = 3 

// Represents the blank cells at the top-left of the PivotTable (top-right for RTL sheets).
PivotAreaType_Origin PivotAreaType = 4 

// Represents a field button.
PivotAreaType_Button PivotAreaType = 5 

// Represents the blank cells at the top-right of the PivotTable (top-left for RTL sheets).
PivotAreaType_TopRight PivotAreaType = 6 
)

func Int32ToPivotAreaType(value int32)(PivotAreaType ,error){
	switch value {
		case 0:  return PivotAreaType_None, nil  
		case 1:  return PivotAreaType_Normal, nil  
		case 2:  return PivotAreaType_Data, nil  
		case 3:  return PivotAreaType_All, nil  
		case 4:  return PivotAreaType_Origin, nil  
		case 5:  return PivotAreaType_Button, nil  
		case 6:  return PivotAreaType_TopRight, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotAreaType value: %d", value)
	}
}

/**************Enum PivotConditionFormatRuleType *****************/
// Represents PivotTable condition formatting rule type.
type PivotConditionFormatRuleType int32

const(
// Indicates that Top N conditional formatting is not evaluated
PivotConditionFormatRuleType_None PivotConditionFormatRuleType = 0 

// Indicates that Top N conditional formatting is
// evaluated across the entire scope range.
PivotConditionFormatRuleType_All PivotConditionFormatRuleType = 1 

// Indicates that Top N conditional formatting is evaluated for each row.
PivotConditionFormatRuleType_Row PivotConditionFormatRuleType = 2 

// Indicates that Top N conditional formatting is
// evaluated for each column.
PivotConditionFormatRuleType_Column PivotConditionFormatRuleType = 3 
)

func Int32ToPivotConditionFormatRuleType(value int32)(PivotConditionFormatRuleType ,error){
	switch value {
		case 0:  return PivotConditionFormatRuleType_None, nil  
		case 1:  return PivotConditionFormatRuleType_All, nil  
		case 2:  return PivotConditionFormatRuleType_Row, nil  
		case 3:  return PivotConditionFormatRuleType_Column, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotConditionFormatRuleType value: %d", value)
	}
}

/**************Enum PivotConditionFormatScopeType *****************/
// Represents PivotTable condition formatting scope type.
type PivotConditionFormatScopeType int32

const(
// Indicates that conditional formatting is applied to the selected data fields.
PivotConditionFormatScopeType_Data PivotConditionFormatScopeType = 0 

// Indicates that conditional formatting is applied to the selected PivotTable field intersections.
PivotConditionFormatScopeType_Field PivotConditionFormatScopeType = 1 

// Indicates that conditional formatting is applied to the selected cells.
PivotConditionFormatScopeType_Selection PivotConditionFormatScopeType = 2 
)

func Int32ToPivotConditionFormatScopeType(value int32)(PivotConditionFormatScopeType ,error){
	switch value {
		case 0:  return PivotConditionFormatScopeType_Data, nil  
		case 1:  return PivotConditionFormatScopeType_Field, nil  
		case 2:  return PivotConditionFormatScopeType_Selection, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotConditionFormatScopeType value: %d", value)
	}
}

/**************Enum PivotFieldDataDisplayFormat *****************/
// Represents data display format in the PivotTable data field.
type PivotFieldDataDisplayFormat int32

const(
// Represents normal display format.
PivotFieldDataDisplayFormat_Normal PivotFieldDataDisplayFormat = 0 

// Represents difference from display format.
PivotFieldDataDisplayFormat_DifferenceFrom PivotFieldDataDisplayFormat = 1 

// Represents percentage of display format.
PivotFieldDataDisplayFormat_PercentageOf PivotFieldDataDisplayFormat = 2 

// Represents percentage difference from  display format.
PivotFieldDataDisplayFormat_PercentageDifferenceFrom PivotFieldDataDisplayFormat = 3 

// Represents running total in display format.
PivotFieldDataDisplayFormat_RunningTotalIn PivotFieldDataDisplayFormat = 4 

// Represents percentage of row display format.
PivotFieldDataDisplayFormat_PercentageOfRow PivotFieldDataDisplayFormat = 5 

// Represents percentage of column display format.
PivotFieldDataDisplayFormat_PercentageOfColumn PivotFieldDataDisplayFormat = 6 

// Represents percentage of total display format.
PivotFieldDataDisplayFormat_PercentageOfTotal PivotFieldDataDisplayFormat = 7 

// Represents index display format.
PivotFieldDataDisplayFormat_Index PivotFieldDataDisplayFormat = 8 

// Represents percentage of parent row total display format.
PivotFieldDataDisplayFormat_PercentageOfParentRowTotal PivotFieldDataDisplayFormat = 9 

// Represents percentage of parent column total display format.
PivotFieldDataDisplayFormat_PercentageOfParentColumnTotal PivotFieldDataDisplayFormat = 10 

// Represents percentage of parent total display format.
PivotFieldDataDisplayFormat_PercentageOfParentTotal PivotFieldDataDisplayFormat = 11 

// Represents percentage of running total in display format.
PivotFieldDataDisplayFormat_PercentageOfRunningTotalIn PivotFieldDataDisplayFormat = 12 

// Represents smallest to largest display format.
PivotFieldDataDisplayFormat_RankSmallestToLargest PivotFieldDataDisplayFormat = 13 

// Represents largest to smallest display format.
PivotFieldDataDisplayFormat_RankLargestToSmallest PivotFieldDataDisplayFormat = 14 
)

func Int32ToPivotFieldDataDisplayFormat(value int32)(PivotFieldDataDisplayFormat ,error){
	switch value {
		case 0:  return PivotFieldDataDisplayFormat_Normal, nil  
		case 1:  return PivotFieldDataDisplayFormat_DifferenceFrom, nil  
		case 2:  return PivotFieldDataDisplayFormat_PercentageOf, nil  
		case 3:  return PivotFieldDataDisplayFormat_PercentageDifferenceFrom, nil  
		case 4:  return PivotFieldDataDisplayFormat_RunningTotalIn, nil  
		case 5:  return PivotFieldDataDisplayFormat_PercentageOfRow, nil  
		case 6:  return PivotFieldDataDisplayFormat_PercentageOfColumn, nil  
		case 7:  return PivotFieldDataDisplayFormat_PercentageOfTotal, nil  
		case 8:  return PivotFieldDataDisplayFormat_Index, nil  
		case 9:  return PivotFieldDataDisplayFormat_PercentageOfParentRowTotal, nil  
		case 10:  return PivotFieldDataDisplayFormat_PercentageOfParentColumnTotal, nil  
		case 11:  return PivotFieldDataDisplayFormat_PercentageOfParentTotal, nil  
		case 12:  return PivotFieldDataDisplayFormat_PercentageOfRunningTotalIn, nil  
		case 13:  return PivotFieldDataDisplayFormat_RankSmallestToLargest, nil  
		case 14:  return PivotFieldDataDisplayFormat_RankLargestToSmallest, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotFieldDataDisplayFormat value: %d", value)
	}
}

/**************Enum PivotFieldGroupType *****************/
// Represents the group type of pivot field.
type PivotFieldGroupType int32

const(
// No group
PivotFieldGroupType_None PivotFieldGroupType = 0 

// Grouped by DateTime range.
PivotFieldGroupType_DateTimeRange PivotFieldGroupType = 1 

// Grouped by numberic range.
PivotFieldGroupType_NumbericRange PivotFieldGroupType = 2 

// Grouped by discrete points.
PivotFieldGroupType_Discrete PivotFieldGroupType = 3 
)

func Int32ToPivotFieldGroupType(value int32)(PivotFieldGroupType ,error){
	switch value {
		case 0:  return PivotFieldGroupType_None, nil  
		case 1:  return PivotFieldGroupType_DateTimeRange, nil  
		case 2:  return PivotFieldGroupType_NumbericRange, nil  
		case 3:  return PivotFieldGroupType_Discrete, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotFieldGroupType value: %d", value)
	}
}

/**************Enum PivotFieldSubtotalType *****************/
// Summary description for PivotFieldSubtotalType.
type PivotFieldSubtotalType int32

const(
// Represents None subtotal type.
PivotFieldSubtotalType_None PivotFieldSubtotalType = 0 

// Represents Automatic subtotal type.
PivotFieldSubtotalType_Automatic PivotFieldSubtotalType = 1 

// Represents Sum subtotal type.
PivotFieldSubtotalType_Sum PivotFieldSubtotalType = 2 

// Represents Count subtotal type.
PivotFieldSubtotalType_Count PivotFieldSubtotalType = 4 

// Represents Average subtotal type.
PivotFieldSubtotalType_Average PivotFieldSubtotalType = 8 

// Represents Max subtotal type.
PivotFieldSubtotalType_Max PivotFieldSubtotalType = 16 

// Represents Min subtotal type.
PivotFieldSubtotalType_Min PivotFieldSubtotalType = 32 

// Represents Product subtotal type.
PivotFieldSubtotalType_Product PivotFieldSubtotalType = 64 

// Represents Count Nums subtotal type.
PivotFieldSubtotalType_CountNums PivotFieldSubtotalType = 128 

// Represents Stdev subtotal type.
PivotFieldSubtotalType_Stdev PivotFieldSubtotalType = 256 

// Represents Stdevp subtotal type.
PivotFieldSubtotalType_Stdevp PivotFieldSubtotalType = 512 

// Represents Var subtotal type.
PivotFieldSubtotalType_Var PivotFieldSubtotalType = 1024 

// Represents Varp subtotal type.
PivotFieldSubtotalType_Varp PivotFieldSubtotalType = 2048 
)

func Int32ToPivotFieldSubtotalType(value int32)(PivotFieldSubtotalType ,error){
	switch value {
		case 0:  return PivotFieldSubtotalType_None, nil  
		case 1:  return PivotFieldSubtotalType_Automatic, nil  
		case 2:  return PivotFieldSubtotalType_Sum, nil  
		case 4:  return PivotFieldSubtotalType_Count, nil  
		case 8:  return PivotFieldSubtotalType_Average, nil  
		case 16:  return PivotFieldSubtotalType_Max, nil  
		case 32:  return PivotFieldSubtotalType_Min, nil  
		case 64:  return PivotFieldSubtotalType_Product, nil  
		case 128:  return PivotFieldSubtotalType_CountNums, nil  
		case 256:  return PivotFieldSubtotalType_Stdev, nil  
		case 512:  return PivotFieldSubtotalType_Stdevp, nil  
		case 1024:  return PivotFieldSubtotalType_Var, nil  
		case 2048:  return PivotFieldSubtotalType_Varp, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotFieldSubtotalType value: %d", value)
	}
}

/**************Enum PivotFieldType *****************/
// Represents PivotTable field type.
type PivotFieldType int32

const(
// Presents base pivot field type.
PivotFieldType_Undefined PivotFieldType = 0 

// Presents row pivot field type.
PivotFieldType_Row PivotFieldType = 1 

// Presents column pivot field type.
PivotFieldType_Column PivotFieldType = 2 

// Presents page pivot field type.
PivotFieldType_Page PivotFieldType = 4 

// Presents data pivot field type.
PivotFieldType_Data PivotFieldType = 8 
)

func Int32ToPivotFieldType(value int32)(PivotFieldType ,error){
	switch value {
		case 0:  return PivotFieldType_Undefined, nil  
		case 1:  return PivotFieldType_Row, nil  
		case 2:  return PivotFieldType_Column, nil  
		case 4:  return PivotFieldType_Page, nil  
		case 8:  return PivotFieldType_Data, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotFieldType value: %d", value)
	}
}

/**************Enum PivotFilterType *****************/
// Represents PivotTable Filter type.
type PivotFilterType int32

const(
// Indicates the "begins with" filter for field captions.
PivotFilterType_CaptionBeginsWith PivotFilterType = 0 

// Indicates the "is between" filter for field captions.
PivotFilterType_CaptionBetween PivotFilterType = 1 

// Indicates the "contains" filter for field captions.
PivotFilterType_CaptionContains PivotFilterType = 2 

// Indicates the "ends with" filter for field captions.
PivotFilterType_CaptionEndsWith PivotFilterType = 3 

// Indicates the "equal" filter for field captions.
PivotFilterType_CaptionEqual PivotFilterType = 4 

// Indicates the "is greater than" filter for field captions.
PivotFilterType_CaptionGreaterThan PivotFilterType = 5 

// Indicates the "is greater than or equal to" filter for field captions.
PivotFilterType_CaptionGreaterThanOrEqual PivotFilterType = 6 

// Indicates the "is less than" filter for field captions.
PivotFilterType_CaptionLessThan PivotFilterType = 7 

// Indicates the "is less than or equal to" filter for field captions.
PivotFilterType_CaptionLessThanOrEqual PivotFilterType = 8 

// Indicates the "does not begin with" filter for field captions.
PivotFilterType_CaptionNotBeginsWith PivotFilterType = 9 

// Indicates the "is not between" filter for field captions.
PivotFilterType_CaptionNotBetween PivotFilterType = 10 

// Indicates the "does not contain" filter for field captions.
PivotFilterType_CaptionNotContains PivotFilterType = 11 

// Indicates the "does not end with" filter for field captions.
PivotFilterType_CaptionNotEndsWith PivotFilterType = 12 

// Indicates the "not equal" filter for field captions.
PivotFilterType_CaptionNotEqual PivotFilterType = 13 

// Indicates the "count" filter.
PivotFilterType_Count PivotFilterType = 14 

// Indicates the "between" filter for date values.
PivotFilterType_DateBetween PivotFilterType = 15 

// Indicates the "equals" filter for date values.
PivotFilterType_DateEqual PivotFilterType = 16 

// Indicates the "after" filter for date values.
PivotFilterType_DateAfter PivotFilterType = 17 

// Indicates the "after or equal to" filter for date values.
PivotFilterType_DateAfterOrEqual PivotFilterType = 18 

// Indicates the "not between" filter for date values.
PivotFilterType_DateNotBetween PivotFilterType = 19 

// Indicates the "does not equal" filter for date values.
PivotFilterType_DateNotEqual PivotFilterType = 20 

// Indicates the "before" filter for date values.
PivotFilterType_DateBefore PivotFilterType = 21 

// Indicates the "before or equal to" filter for date values.
PivotFilterType_DateBeforeOrEqual PivotFilterType = 22 

// Indicates the "last month" filter for date values.
PivotFilterType_LastMonth PivotFilterType = 23 

// Indicates the "last quarter" filter for date values.
PivotFilterType_LastQuarter PivotFilterType = 24 

// Indicates the "last week" filter for date values.
PivotFilterType_LastWeek PivotFilterType = 25 

// Indicates the "last year" filter for date values.
PivotFilterType_LastYear PivotFilterType = 26 

// Indicates the "January" filter for date values.
PivotFilterType_January PivotFilterType = 27 

// Indicates the "February" filter for date values.
PivotFilterType_February PivotFilterType = 28 

// Indicates the "March" filter for date values.
PivotFilterType_March PivotFilterType = 29 

// Indicates the "April" filter for date values.
PivotFilterType_April PivotFilterType = 30 

// Indicates the "May" filter for date values.
PivotFilterType_May PivotFilterType = 31 

// Indicates the "June" filter for date values.
PivotFilterType_June PivotFilterType = 32 

// Indicates the "July" filter for date values.
PivotFilterType_July PivotFilterType = 33 

// Indicates the "August" filter for date values.
PivotFilterType_August PivotFilterType = 34 

// Indicates the "September" filter for date values.
PivotFilterType_September PivotFilterType = 35 

// Indicates the "October" filter for date values.
PivotFilterType_October PivotFilterType = 36 

// Indicates the "November" filter for date values.
PivotFilterType_November PivotFilterType = 37 

// Indicates the "December" filter for date values.
PivotFilterType_December PivotFilterType = 38 

// Indicates the "next month" filter for date values.
PivotFilterType_NextMonth PivotFilterType = 39 

// Indicates the "next quarter" for date values.
PivotFilterType_NextQuarter PivotFilterType = 40 

// Indicates the "next week" for date values.
PivotFilterType_NextWeek PivotFilterType = 41 

// Indicates the "next year" filter for date values.
PivotFilterType_NextYear PivotFilterType = 42 

// Indicates the "percent" filter for numeric values.
PivotFilterType_Percent PivotFilterType = 43 

// Indicates the "first quarter" filter for date values.
PivotFilterType_Quarter1 PivotFilterType = 44 

// Indicates the "second quarter" filter for date values.
PivotFilterType_Quarter2 PivotFilterType = 45 

// Indicates the "third quarter" filter for date values.
PivotFilterType_Quarter3 PivotFilterType = 46 

// Indicates the "fourth quarter" filter for date values.
PivotFilterType_Quarter4 PivotFilterType = 47 

// Indicates the "sum" filter for numeric values.
PivotFilterType_Sum PivotFilterType = 48 

// Indicates the "this month" filter for date values.
PivotFilterType_ThisMonth PivotFilterType = 49 

// Indicates the "this quarter" filter for date values.
PivotFilterType_ThisQuarter PivotFilterType = 50 

// Indicates the "this week" filter for date values.
PivotFilterType_ThisWeek PivotFilterType = 51 

// Indicate the "this year" filter for date values.
PivotFilterType_ThisYear PivotFilterType = 52 

// Indicates the "today" filter for date values.
PivotFilterType_Today PivotFilterType = 53 

// Indicates the "tomorrow" filter for date values.
PivotFilterType_Tomorrow PivotFilterType = 54 

// Indicates the PivotTable filter is unknown to the application.
PivotFilterType_Unknown PivotFilterType = 55 

// Indicates the "Value between" filter for text and numeric values.
PivotFilterType_ValueBetween PivotFilterType = 56 

// Indicates the "value equal" filter for text and numeric values.
PivotFilterType_ValueEqual PivotFilterType = 57 

// Indicates the "value greater than" filter for text and numeric values.
PivotFilterType_ValueGreaterThan PivotFilterType = 58 

// Indicates the "value greater than or equal to" filter for text and numeric values.
PivotFilterType_ValueGreaterThanOrEqual PivotFilterType = 59 

// Indicates the "value less than" filter for text and numeric values.
PivotFilterType_ValueLessThan PivotFilterType = 60 

// Indicates the "value less than or equal to" filter for text and numeric values.
PivotFilterType_ValueLessThanOrEqual PivotFilterType = 61 

// Indicates the "value not between" filter for text and numeric values.
PivotFilterType_ValueNotBetween PivotFilterType = 62 

// Indicates the "value not equal" filter for text and numeric values.
PivotFilterType_ValueNotEqual PivotFilterType = 63 

// Indicates the "year-to-date" filter for date values.
PivotFilterType_YearToDate PivotFilterType = 64 

// Indicates the "yesterday" filter for date values.
PivotFilterType_Yesterday PivotFilterType = 65 

// No filter.
PivotFilterType_None PivotFilterType = 255 
)

func Int32ToPivotFilterType(value int32)(PivotFilterType ,error){
	switch value {
		case 0:  return PivotFilterType_CaptionBeginsWith, nil  
		case 1:  return PivotFilterType_CaptionBetween, nil  
		case 2:  return PivotFilterType_CaptionContains, nil  
		case 3:  return PivotFilterType_CaptionEndsWith, nil  
		case 4:  return PivotFilterType_CaptionEqual, nil  
		case 5:  return PivotFilterType_CaptionGreaterThan, nil  
		case 6:  return PivotFilterType_CaptionGreaterThanOrEqual, nil  
		case 7:  return PivotFilterType_CaptionLessThan, nil  
		case 8:  return PivotFilterType_CaptionLessThanOrEqual, nil  
		case 9:  return PivotFilterType_CaptionNotBeginsWith, nil  
		case 10:  return PivotFilterType_CaptionNotBetween, nil  
		case 11:  return PivotFilterType_CaptionNotContains, nil  
		case 12:  return PivotFilterType_CaptionNotEndsWith, nil  
		case 13:  return PivotFilterType_CaptionNotEqual, nil  
		case 14:  return PivotFilterType_Count, nil  
		case 15:  return PivotFilterType_DateBetween, nil  
		case 16:  return PivotFilterType_DateEqual, nil  
		case 17:  return PivotFilterType_DateAfter, nil  
		case 18:  return PivotFilterType_DateAfterOrEqual, nil  
		case 19:  return PivotFilterType_DateNotBetween, nil  
		case 20:  return PivotFilterType_DateNotEqual, nil  
		case 21:  return PivotFilterType_DateBefore, nil  
		case 22:  return PivotFilterType_DateBeforeOrEqual, nil  
		case 23:  return PivotFilterType_LastMonth, nil  
		case 24:  return PivotFilterType_LastQuarter, nil  
		case 25:  return PivotFilterType_LastWeek, nil  
		case 26:  return PivotFilterType_LastYear, nil  
		case 27:  return PivotFilterType_January, nil  
		case 28:  return PivotFilterType_February, nil  
		case 29:  return PivotFilterType_March, nil  
		case 30:  return PivotFilterType_April, nil  
		case 31:  return PivotFilterType_May, nil  
		case 32:  return PivotFilterType_June, nil  
		case 33:  return PivotFilterType_July, nil  
		case 34:  return PivotFilterType_August, nil  
		case 35:  return PivotFilterType_September, nil  
		case 36:  return PivotFilterType_October, nil  
		case 37:  return PivotFilterType_November, nil  
		case 38:  return PivotFilterType_December, nil  
		case 39:  return PivotFilterType_NextMonth, nil  
		case 40:  return PivotFilterType_NextQuarter, nil  
		case 41:  return PivotFilterType_NextWeek, nil  
		case 42:  return PivotFilterType_NextYear, nil  
		case 43:  return PivotFilterType_Percent, nil  
		case 44:  return PivotFilterType_Quarter1, nil  
		case 45:  return PivotFilterType_Quarter2, nil  
		case 46:  return PivotFilterType_Quarter3, nil  
		case 47:  return PivotFilterType_Quarter4, nil  
		case 48:  return PivotFilterType_Sum, nil  
		case 49:  return PivotFilterType_ThisMonth, nil  
		case 50:  return PivotFilterType_ThisQuarter, nil  
		case 51:  return PivotFilterType_ThisWeek, nil  
		case 52:  return PivotFilterType_ThisYear, nil  
		case 53:  return PivotFilterType_Today, nil  
		case 54:  return PivotFilterType_Tomorrow, nil  
		case 55:  return PivotFilterType_Unknown, nil  
		case 56:  return PivotFilterType_ValueBetween, nil  
		case 57:  return PivotFilterType_ValueEqual, nil  
		case 58:  return PivotFilterType_ValueGreaterThan, nil  
		case 59:  return PivotFilterType_ValueGreaterThanOrEqual, nil  
		case 60:  return PivotFilterType_ValueLessThan, nil  
		case 61:  return PivotFilterType_ValueLessThanOrEqual, nil  
		case 62:  return PivotFilterType_ValueNotBetween, nil  
		case 63:  return PivotFilterType_ValueNotEqual, nil  
		case 64:  return PivotFilterType_YearToDate, nil  
		case 65:  return PivotFilterType_Yesterday, nil  
		case 255:  return PivotFilterType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotFilterType value: %d", value)
	}
}

/**************Enum PivotGroupByType *****************/
// Represents group by type.
type PivotGroupByType int32

const(
// Group by numbers.
PivotGroupByType_Numbers PivotGroupByType = 0 

// Presents Seconds groupby type.
PivotGroupByType_Seconds PivotGroupByType = 1 

// Presents Minutes groupby type.
PivotGroupByType_Minutes PivotGroupByType = 2 

// Presents Hours groupby type.
PivotGroupByType_Hours PivotGroupByType = 3 

// Presents Days groupby type.
PivotGroupByType_Days PivotGroupByType = 4 

// Presents Months groupby type.
PivotGroupByType_Months PivotGroupByType = 5 

// Presents Quarters groupby type.
PivotGroupByType_Quarters PivotGroupByType = 6 

// Presents Years groupby type.
PivotGroupByType_Years PivotGroupByType = 7 
)

func Int32ToPivotGroupByType(value int32)(PivotGroupByType ,error){
	switch value {
		case 0:  return PivotGroupByType_Numbers, nil  
		case 1:  return PivotGroupByType_Seconds, nil  
		case 2:  return PivotGroupByType_Minutes, nil  
		case 3:  return PivotGroupByType_Hours, nil  
		case 4:  return PivotGroupByType_Days, nil  
		case 5:  return PivotGroupByType_Months, nil  
		case 6:  return PivotGroupByType_Quarters, nil  
		case 7:  return PivotGroupByType_Years, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotGroupByType value: %d", value)
	}
}

/**************Enum PivotItemPosition *****************/
// Represents base item Next/Previous/All position in the base field .
type PivotItemPosition int32

const(
// Represents the previous pivot item in the PivotField.
PivotItemPosition_Previous PivotItemPosition = 0 

// Represents the next pivot item in the PivotField.
PivotItemPosition_Next PivotItemPosition = 1 

// Shows values as the different format based the index of pivot item in the PivotField.
PivotItemPosition_Custom PivotItemPosition = 2 
)

func Int32ToPivotItemPosition(value int32)(PivotItemPosition ,error){
	switch value {
		case 0:  return PivotItemPosition_Previous, nil  
		case 1:  return PivotItemPosition_Next, nil  
		case 2:  return PivotItemPosition_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotItemPosition value: %d", value)
	}
}

/**************Enum PivotItemPositionType *****************/
// Represents the position type of the pivot base item in the base field when the ShowDataAs calculation is in use.
type PivotItemPositionType int32

const(
// Represents the previous pivot item in the PivotField.
PivotItemPositionType_Previous PivotItemPositionType = 0 

// Represents the next pivot item in the PivotField.
PivotItemPositionType_Next PivotItemPositionType = 1 

// Shows values as the different format based the index of pivot item in the PivotField.
PivotItemPositionType_Custom PivotItemPositionType = 2 
)

func Int32ToPivotItemPositionType(value int32)(PivotItemPositionType ,error){
	switch value {
		case 0:  return PivotItemPositionType_Previous, nil  
		case 1:  return PivotItemPositionType_Next, nil  
		case 2:  return PivotItemPositionType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotItemPositionType value: %d", value)
	}
}

/**************Enum PivotLineType *****************/
// Specifies the type of the PivotLine.
type PivotLineType int32

const(
// Regular PivotLine with pivot items.
PivotLineType_Regular PivotLineType = 0 

// Subtotal line.
PivotLineType_Subtotal PivotLineType = 1 

// Grand Total line.
PivotLineType_GrandTotal PivotLineType = 2 

// Blank line after each group.
PivotLineType_Blank PivotLineType = 3 
)

func Int32ToPivotLineType(value int32)(PivotLineType ,error){
	switch value {
		case 0:  return PivotLineType_Regular, nil  
		case 1:  return PivotLineType_Subtotal, nil  
		case 2:  return PivotLineType_GrandTotal, nil  
		case 3:  return PivotLineType_Blank, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotLineType value: %d", value)
	}
}

/**************Enum PivotMissingItemLimitType *****************/
// Represents number of items to retain per field.
type PivotMissingItemLimitType int32

const(
// The default number of unique items per PivotField allowed.
PivotMissingItemLimitType_Automatic PivotMissingItemLimitType = 0 

// The maximum number of unique items per PivotField allowed (>32,500).
PivotMissingItemLimitType_Max PivotMissingItemLimitType = 1 

// No unique items per PivotField allowed.
PivotMissingItemLimitType_None PivotMissingItemLimitType = 2 
)

func Int32ToPivotMissingItemLimitType(value int32)(PivotMissingItemLimitType ,error){
	switch value {
		case 0:  return PivotMissingItemLimitType_Automatic, nil  
		case 1:  return PivotMissingItemLimitType_Max, nil  
		case 2:  return PivotMissingItemLimitType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotMissingItemLimitType value: %d", value)
	}
}

/**************Enum PivotRefreshState *****************/
// The state for refreshing pivot tables.
type PivotRefreshState int32

const(
// Successfully refreshed
PivotRefreshState_Success PivotRefreshState = 0 

// Refresh failed because the data source is external.
PivotRefreshState_UnsupportedExternalDataSource PivotRefreshState = 1 
)

func Int32ToPivotRefreshState(value int32)(PivotRefreshState ,error){
	switch value {
		case 0:  return PivotRefreshState_Success, nil  
		case 1:  return PivotRefreshState_UnsupportedExternalDataSource, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotRefreshState value: %d", value)
	}
}

/**************Enum PivotTableAutoFormatType *****************/
// Represents PivotTable auto format type.
type PivotTableAutoFormatType int32

const(
// Represents None format type.
PivotTableAutoFormatType_None PivotTableAutoFormatType = 0 

// Represents Classic auto format type.
PivotTableAutoFormatType_Classic PivotTableAutoFormatType = 1 

// Represents Report1 format type.
PivotTableAutoFormatType_Report1 PivotTableAutoFormatType = 2 

// Represents Report2 format type.
PivotTableAutoFormatType_Report2 PivotTableAutoFormatType = 3 

// Represents Report3 format type.
PivotTableAutoFormatType_Report3 PivotTableAutoFormatType = 4 

// Represents Report4 format type.
PivotTableAutoFormatType_Report4 PivotTableAutoFormatType = 5 

// Represents Report5 format type.
PivotTableAutoFormatType_Report5 PivotTableAutoFormatType = 6 

// Represents Report6 format type.
PivotTableAutoFormatType_Report6 PivotTableAutoFormatType = 7 

// Represents Report7 format type.
PivotTableAutoFormatType_Report7 PivotTableAutoFormatType = 8 

// Represents Report8 format type.
PivotTableAutoFormatType_Report8 PivotTableAutoFormatType = 9 

// Represents Report9 format type.
PivotTableAutoFormatType_Report9 PivotTableAutoFormatType = 10 

// Represents Report10 format type.
PivotTableAutoFormatType_Report10 PivotTableAutoFormatType = 11 

// Represents Table1 format type.
PivotTableAutoFormatType_Table1 PivotTableAutoFormatType = 12 

// Represents Table2 format type.
PivotTableAutoFormatType_Table2 PivotTableAutoFormatType = 13 

// Represents Table3 format type.
PivotTableAutoFormatType_Table3 PivotTableAutoFormatType = 14 

// Represents Table4 format type.
PivotTableAutoFormatType_Table4 PivotTableAutoFormatType = 15 

// Represents Table5 format type.
PivotTableAutoFormatType_Table5 PivotTableAutoFormatType = 16 

// Represents Table6 format type.
PivotTableAutoFormatType_Table6 PivotTableAutoFormatType = 17 

// Represents Table7 format type.
PivotTableAutoFormatType_Table7 PivotTableAutoFormatType = 18 

// Represents Table8 format type.
PivotTableAutoFormatType_Table8 PivotTableAutoFormatType = 19 

// Represents Table9 format type.
PivotTableAutoFormatType_Table9 PivotTableAutoFormatType = 20 

// Represents Table10 format type.
PivotTableAutoFormatType_Table10 PivotTableAutoFormatType = 21 
)

func Int32ToPivotTableAutoFormatType(value int32)(PivotTableAutoFormatType ,error){
	switch value {
		case 0:  return PivotTableAutoFormatType_None, nil  
		case 1:  return PivotTableAutoFormatType_Classic, nil  
		case 2:  return PivotTableAutoFormatType_Report1, nil  
		case 3:  return PivotTableAutoFormatType_Report2, nil  
		case 4:  return PivotTableAutoFormatType_Report3, nil  
		case 5:  return PivotTableAutoFormatType_Report4, nil  
		case 6:  return PivotTableAutoFormatType_Report5, nil  
		case 7:  return PivotTableAutoFormatType_Report6, nil  
		case 8:  return PivotTableAutoFormatType_Report7, nil  
		case 9:  return PivotTableAutoFormatType_Report8, nil  
		case 10:  return PivotTableAutoFormatType_Report9, nil  
		case 11:  return PivotTableAutoFormatType_Report10, nil  
		case 12:  return PivotTableAutoFormatType_Table1, nil  
		case 13:  return PivotTableAutoFormatType_Table2, nil  
		case 14:  return PivotTableAutoFormatType_Table3, nil  
		case 15:  return PivotTableAutoFormatType_Table4, nil  
		case 16:  return PivotTableAutoFormatType_Table5, nil  
		case 17:  return PivotTableAutoFormatType_Table6, nil  
		case 18:  return PivotTableAutoFormatType_Table7, nil  
		case 19:  return PivotTableAutoFormatType_Table8, nil  
		case 20:  return PivotTableAutoFormatType_Table9, nil  
		case 21:  return PivotTableAutoFormatType_Table10, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotTableAutoFormatType value: %d", value)
	}
}

/**************Enum PivotTableSelectionType *****************/
// Specifies what can be selected in a PivotTable during a structured selection.
// These constants can be combined to select multiple types.
type PivotTableSelectionType int32

const(
// Data and labels
PivotTableSelectionType_DataAndLabel PivotTableSelectionType = 0 

// Only selects data
PivotTableSelectionType_DataOnly PivotTableSelectionType = 2 

// Only selects labels
PivotTableSelectionType_LabelOnly PivotTableSelectionType = 1 
)

func Int32ToPivotTableSelectionType(value int32)(PivotTableSelectionType ,error){
	switch value {
		case 0:  return PivotTableSelectionType_DataAndLabel, nil  
		case 2:  return PivotTableSelectionType_DataOnly, nil  
		case 1:  return PivotTableSelectionType_LabelOnly, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotTableSelectionType value: %d", value)
	}
}

/**************Enum PivotTableSourceType *****************/
// Represents data source type of the pivot table.
type PivotTableSourceType int32

const(
// Specifies that the source data is a range.
PivotTableSourceType_Sheet PivotTableSourceType = 1 

// Specifies that external source data is used.
PivotTableSourceType_External PivotTableSourceType = 2 

// Specifies that multiple consolidation ranges are used as the source data.
PivotTableSourceType_Consolidation PivotTableSourceType = 4 

// The source data is populated from a temporary internal structure.
PivotTableSourceType_Scenario PivotTableSourceType = 8 

// Unknown data source.
PivotTableSourceType_Unknown PivotTableSourceType = 9 
)

func Int32ToPivotTableSourceType(value int32)(PivotTableSourceType ,error){
	switch value {
		case 1:  return PivotTableSourceType_Sheet, nil  
		case 2:  return PivotTableSourceType_External, nil  
		case 4:  return PivotTableSourceType_Consolidation, nil  
		case 8:  return PivotTableSourceType_Scenario, nil  
		case 9:  return PivotTableSourceType_Unknown, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotTableSourceType value: %d", value)
	}
}

/**************Enum PivotTableStyleType *****************/
// Represents the pivot table style type.
type PivotTableStyleType int32

const(
PivotTableStyleType_None PivotTableStyleType = 0 


PivotTableStyleType_PivotTableStyleLight1 PivotTableStyleType = 1 


PivotTableStyleType_PivotTableStyleLight2 PivotTableStyleType = 2 


PivotTableStyleType_PivotTableStyleLight3 PivotTableStyleType = 3 


PivotTableStyleType_PivotTableStyleLight4 PivotTableStyleType = 4 


PivotTableStyleType_PivotTableStyleLight5 PivotTableStyleType = 5 


PivotTableStyleType_PivotTableStyleLight6 PivotTableStyleType = 6 


PivotTableStyleType_PivotTableStyleLight7 PivotTableStyleType = 7 


PivotTableStyleType_PivotTableStyleLight8 PivotTableStyleType = 8 


PivotTableStyleType_PivotTableStyleLight9 PivotTableStyleType = 9 


PivotTableStyleType_PivotTableStyleLight10 PivotTableStyleType = 10 


PivotTableStyleType_PivotTableStyleLight11 PivotTableStyleType = 11 


PivotTableStyleType_PivotTableStyleLight12 PivotTableStyleType = 12 


PivotTableStyleType_PivotTableStyleLight13 PivotTableStyleType = 13 


PivotTableStyleType_PivotTableStyleLight14 PivotTableStyleType = 14 


PivotTableStyleType_PivotTableStyleLight15 PivotTableStyleType = 15 


PivotTableStyleType_PivotTableStyleLight16 PivotTableStyleType = 16 


PivotTableStyleType_PivotTableStyleLight17 PivotTableStyleType = 17 


PivotTableStyleType_PivotTableStyleLight18 PivotTableStyleType = 18 


PivotTableStyleType_PivotTableStyleLight19 PivotTableStyleType = 19 


PivotTableStyleType_PivotTableStyleLight20 PivotTableStyleType = 20 


PivotTableStyleType_PivotTableStyleLight21 PivotTableStyleType = 21 


PivotTableStyleType_PivotTableStyleLight22 PivotTableStyleType = 22 


PivotTableStyleType_PivotTableStyleLight23 PivotTableStyleType = 23 


PivotTableStyleType_PivotTableStyleLight24 PivotTableStyleType = 24 


PivotTableStyleType_PivotTableStyleLight25 PivotTableStyleType = 25 


PivotTableStyleType_PivotTableStyleLight26 PivotTableStyleType = 26 


PivotTableStyleType_PivotTableStyleLight27 PivotTableStyleType = 27 


PivotTableStyleType_PivotTableStyleLight28 PivotTableStyleType = 28 


PivotTableStyleType_PivotTableStyleMedium1 PivotTableStyleType = 29 


PivotTableStyleType_PivotTableStyleMedium2 PivotTableStyleType = 30 


PivotTableStyleType_PivotTableStyleMedium3 PivotTableStyleType = 31 


PivotTableStyleType_PivotTableStyleMedium4 PivotTableStyleType = 32 


PivotTableStyleType_PivotTableStyleMedium5 PivotTableStyleType = 33 


PivotTableStyleType_PivotTableStyleMedium6 PivotTableStyleType = 34 


PivotTableStyleType_PivotTableStyleMedium7 PivotTableStyleType = 35 


PivotTableStyleType_PivotTableStyleMedium8 PivotTableStyleType = 36 


PivotTableStyleType_PivotTableStyleMedium9 PivotTableStyleType = 37 


PivotTableStyleType_PivotTableStyleMedium10 PivotTableStyleType = 38 


PivotTableStyleType_PivotTableStyleMedium11 PivotTableStyleType = 39 


PivotTableStyleType_PivotTableStyleMedium12 PivotTableStyleType = 40 


PivotTableStyleType_PivotTableStyleMedium13 PivotTableStyleType = 41 


PivotTableStyleType_PivotTableStyleMedium14 PivotTableStyleType = 42 


PivotTableStyleType_PivotTableStyleMedium15 PivotTableStyleType = 43 


PivotTableStyleType_PivotTableStyleMedium16 PivotTableStyleType = 44 


PivotTableStyleType_PivotTableStyleMedium17 PivotTableStyleType = 45 


PivotTableStyleType_PivotTableStyleMedium18 PivotTableStyleType = 46 


PivotTableStyleType_PivotTableStyleMedium19 PivotTableStyleType = 47 


PivotTableStyleType_PivotTableStyleMedium20 PivotTableStyleType = 48 


PivotTableStyleType_PivotTableStyleMedium21 PivotTableStyleType = 49 


PivotTableStyleType_PivotTableStyleMedium22 PivotTableStyleType = 50 


PivotTableStyleType_PivotTableStyleMedium23 PivotTableStyleType = 51 


PivotTableStyleType_PivotTableStyleMedium24 PivotTableStyleType = 52 


PivotTableStyleType_PivotTableStyleMedium25 PivotTableStyleType = 53 


PivotTableStyleType_PivotTableStyleMedium26 PivotTableStyleType = 54 


PivotTableStyleType_PivotTableStyleMedium27 PivotTableStyleType = 55 


PivotTableStyleType_PivotTableStyleMedium28 PivotTableStyleType = 56 


PivotTableStyleType_PivotTableStyleDark1 PivotTableStyleType = 57 


PivotTableStyleType_PivotTableStyleDark2 PivotTableStyleType = 58 


PivotTableStyleType_PivotTableStyleDark3 PivotTableStyleType = 59 


PivotTableStyleType_PivotTableStyleDark4 PivotTableStyleType = 60 


PivotTableStyleType_PivotTableStyleDark5 PivotTableStyleType = 61 


PivotTableStyleType_PivotTableStyleDark6 PivotTableStyleType = 62 


PivotTableStyleType_PivotTableStyleDark7 PivotTableStyleType = 63 


PivotTableStyleType_PivotTableStyleDark8 PivotTableStyleType = 64 


PivotTableStyleType_PivotTableStyleDark9 PivotTableStyleType = 65 


PivotTableStyleType_PivotTableStyleDark10 PivotTableStyleType = 66 


PivotTableStyleType_PivotTableStyleDark11 PivotTableStyleType = 67 


PivotTableStyleType_PivotTableStyleDark12 PivotTableStyleType = 68 


PivotTableStyleType_PivotTableStyleDark13 PivotTableStyleType = 69 


PivotTableStyleType_PivotTableStyleDark14 PivotTableStyleType = 70 


PivotTableStyleType_PivotTableStyleDark15 PivotTableStyleType = 71 


PivotTableStyleType_PivotTableStyleDark16 PivotTableStyleType = 72 


PivotTableStyleType_PivotTableStyleDark17 PivotTableStyleType = 73 


PivotTableStyleType_PivotTableStyleDark18 PivotTableStyleType = 74 


PivotTableStyleType_PivotTableStyleDark19 PivotTableStyleType = 75 


PivotTableStyleType_PivotTableStyleDark20 PivotTableStyleType = 76 


PivotTableStyleType_PivotTableStyleDark21 PivotTableStyleType = 77 


PivotTableStyleType_PivotTableStyleDark22 PivotTableStyleType = 78 


PivotTableStyleType_PivotTableStyleDark23 PivotTableStyleType = 79 


PivotTableStyleType_PivotTableStyleDark24 PivotTableStyleType = 80 


PivotTableStyleType_PivotTableStyleDark25 PivotTableStyleType = 81 


PivotTableStyleType_PivotTableStyleDark26 PivotTableStyleType = 82 


PivotTableStyleType_PivotTableStyleDark27 PivotTableStyleType = 83 


PivotTableStyleType_PivotTableStyleDark28 PivotTableStyleType = 84 


PivotTableStyleType_Custom PivotTableStyleType = 85 
)

func Int32ToPivotTableStyleType(value int32)(PivotTableStyleType ,error){
	switch value {
		case 0:  return PivotTableStyleType_None, nil  
		case 1:  return PivotTableStyleType_PivotTableStyleLight1, nil  
		case 2:  return PivotTableStyleType_PivotTableStyleLight2, nil  
		case 3:  return PivotTableStyleType_PivotTableStyleLight3, nil  
		case 4:  return PivotTableStyleType_PivotTableStyleLight4, nil  
		case 5:  return PivotTableStyleType_PivotTableStyleLight5, nil  
		case 6:  return PivotTableStyleType_PivotTableStyleLight6, nil  
		case 7:  return PivotTableStyleType_PivotTableStyleLight7, nil  
		case 8:  return PivotTableStyleType_PivotTableStyleLight8, nil  
		case 9:  return PivotTableStyleType_PivotTableStyleLight9, nil  
		case 10:  return PivotTableStyleType_PivotTableStyleLight10, nil  
		case 11:  return PivotTableStyleType_PivotTableStyleLight11, nil  
		case 12:  return PivotTableStyleType_PivotTableStyleLight12, nil  
		case 13:  return PivotTableStyleType_PivotTableStyleLight13, nil  
		case 14:  return PivotTableStyleType_PivotTableStyleLight14, nil  
		case 15:  return PivotTableStyleType_PivotTableStyleLight15, nil  
		case 16:  return PivotTableStyleType_PivotTableStyleLight16, nil  
		case 17:  return PivotTableStyleType_PivotTableStyleLight17, nil  
		case 18:  return PivotTableStyleType_PivotTableStyleLight18, nil  
		case 19:  return PivotTableStyleType_PivotTableStyleLight19, nil  
		case 20:  return PivotTableStyleType_PivotTableStyleLight20, nil  
		case 21:  return PivotTableStyleType_PivotTableStyleLight21, nil  
		case 22:  return PivotTableStyleType_PivotTableStyleLight22, nil  
		case 23:  return PivotTableStyleType_PivotTableStyleLight23, nil  
		case 24:  return PivotTableStyleType_PivotTableStyleLight24, nil  
		case 25:  return PivotTableStyleType_PivotTableStyleLight25, nil  
		case 26:  return PivotTableStyleType_PivotTableStyleLight26, nil  
		case 27:  return PivotTableStyleType_PivotTableStyleLight27, nil  
		case 28:  return PivotTableStyleType_PivotTableStyleLight28, nil  
		case 29:  return PivotTableStyleType_PivotTableStyleMedium1, nil  
		case 30:  return PivotTableStyleType_PivotTableStyleMedium2, nil  
		case 31:  return PivotTableStyleType_PivotTableStyleMedium3, nil  
		case 32:  return PivotTableStyleType_PivotTableStyleMedium4, nil  
		case 33:  return PivotTableStyleType_PivotTableStyleMedium5, nil  
		case 34:  return PivotTableStyleType_PivotTableStyleMedium6, nil  
		case 35:  return PivotTableStyleType_PivotTableStyleMedium7, nil  
		case 36:  return PivotTableStyleType_PivotTableStyleMedium8, nil  
		case 37:  return PivotTableStyleType_PivotTableStyleMedium9, nil  
		case 38:  return PivotTableStyleType_PivotTableStyleMedium10, nil  
		case 39:  return PivotTableStyleType_PivotTableStyleMedium11, nil  
		case 40:  return PivotTableStyleType_PivotTableStyleMedium12, nil  
		case 41:  return PivotTableStyleType_PivotTableStyleMedium13, nil  
		case 42:  return PivotTableStyleType_PivotTableStyleMedium14, nil  
		case 43:  return PivotTableStyleType_PivotTableStyleMedium15, nil  
		case 44:  return PivotTableStyleType_PivotTableStyleMedium16, nil  
		case 45:  return PivotTableStyleType_PivotTableStyleMedium17, nil  
		case 46:  return PivotTableStyleType_PivotTableStyleMedium18, nil  
		case 47:  return PivotTableStyleType_PivotTableStyleMedium19, nil  
		case 48:  return PivotTableStyleType_PivotTableStyleMedium20, nil  
		case 49:  return PivotTableStyleType_PivotTableStyleMedium21, nil  
		case 50:  return PivotTableStyleType_PivotTableStyleMedium22, nil  
		case 51:  return PivotTableStyleType_PivotTableStyleMedium23, nil  
		case 52:  return PivotTableStyleType_PivotTableStyleMedium24, nil  
		case 53:  return PivotTableStyleType_PivotTableStyleMedium25, nil  
		case 54:  return PivotTableStyleType_PivotTableStyleMedium26, nil  
		case 55:  return PivotTableStyleType_PivotTableStyleMedium27, nil  
		case 56:  return PivotTableStyleType_PivotTableStyleMedium28, nil  
		case 57:  return PivotTableStyleType_PivotTableStyleDark1, nil  
		case 58:  return PivotTableStyleType_PivotTableStyleDark2, nil  
		case 59:  return PivotTableStyleType_PivotTableStyleDark3, nil  
		case 60:  return PivotTableStyleType_PivotTableStyleDark4, nil  
		case 61:  return PivotTableStyleType_PivotTableStyleDark5, nil  
		case 62:  return PivotTableStyleType_PivotTableStyleDark6, nil  
		case 63:  return PivotTableStyleType_PivotTableStyleDark7, nil  
		case 64:  return PivotTableStyleType_PivotTableStyleDark8, nil  
		case 65:  return PivotTableStyleType_PivotTableStyleDark9, nil  
		case 66:  return PivotTableStyleType_PivotTableStyleDark10, nil  
		case 67:  return PivotTableStyleType_PivotTableStyleDark11, nil  
		case 68:  return PivotTableStyleType_PivotTableStyleDark12, nil  
		case 69:  return PivotTableStyleType_PivotTableStyleDark13, nil  
		case 70:  return PivotTableStyleType_PivotTableStyleDark14, nil  
		case 71:  return PivotTableStyleType_PivotTableStyleDark15, nil  
		case 72:  return PivotTableStyleType_PivotTableStyleDark16, nil  
		case 73:  return PivotTableStyleType_PivotTableStyleDark17, nil  
		case 74:  return PivotTableStyleType_PivotTableStyleDark18, nil  
		case 75:  return PivotTableStyleType_PivotTableStyleDark19, nil  
		case 76:  return PivotTableStyleType_PivotTableStyleDark20, nil  
		case 77:  return PivotTableStyleType_PivotTableStyleDark21, nil  
		case 78:  return PivotTableStyleType_PivotTableStyleDark22, nil  
		case 79:  return PivotTableStyleType_PivotTableStyleDark23, nil  
		case 80:  return PivotTableStyleType_PivotTableStyleDark24, nil  
		case 81:  return PivotTableStyleType_PivotTableStyleDark25, nil  
		case 82:  return PivotTableStyleType_PivotTableStyleDark26, nil  
		case 83:  return PivotTableStyleType_PivotTableStyleDark27, nil  
		case 84:  return PivotTableStyleType_PivotTableStyleDark28, nil  
		case 85:  return PivotTableStyleType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid PivotTableStyleType value: %d", value)
	}
}

/**************Enum ReserveMissingPivotItemType *****************/
// Represents how to keep the missing pivot items.
type ReserveMissingPivotItemType int32

const(
// Removes old missint pivot items and reserves visible items which the current data source does not contain as missing items.
ReserveMissingPivotItemType_Default ReserveMissingPivotItemType = 0 

// Reserves all missing items.
ReserveMissingPivotItemType_All ReserveMissingPivotItemType = 1 

// Removes all missing pivot items.
ReserveMissingPivotItemType_None ReserveMissingPivotItemType = 2 
)

func Int32ToReserveMissingPivotItemType(value int32)(ReserveMissingPivotItemType ,error){
	switch value {
		case 0:  return ReserveMissingPivotItemType_Default, nil  
		case 1:  return ReserveMissingPivotItemType_All, nil  
		case 2:  return ReserveMissingPivotItemType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid ReserveMissingPivotItemType value: %d", value)
	}
}
// Class CustomPiovtFieldGroupItem 

// Represents an item of custom grouped field.
type CustomPiovtFieldGroupItem struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CustomPiovtFieldGroupItem) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.CustomPiovtFieldGroupItem_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}


func DeleteCustomPiovtFieldGroupItem(custompiovtfieldgroupitem *CustomPiovtFieldGroupItem){
	runtime.SetFinalizer(custompiovtfieldgroupitem, nil)
	C.Delete_CustomPiovtFieldGroupItem(custompiovtfieldgroupitem.ptr)
	custompiovtfieldgroupitem.ptr = nil
}

// Class PivotArea 

// Presents the selected area of the PivotTable.
type PivotArea struct {
	ptr unsafe.Pointer
}

// Presents the selected area of the PivotTable.
// Parameters:
//   table - PivotTable 
func NewPivotArea(table *PivotTable) ( *PivotArea, error) {
	pivotarea := &PivotArea{}
	CGoReturnPtr := C.New_PivotArea(table.ptr)
	if CGoReturnPtr.error_no == 0 {
		pivotarea.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotarea, DeletePivotArea)
		return pivotarea, nil
	} else {
		pivotarea.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotarea, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotArea) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotArea_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets all filters for this PivotArea.
// Returns:
//   PivotAreaFilterCollection  
func (instance *PivotArea) GetFilters()  (*PivotAreaFilterCollection,  error)  {
	CGoReturnPtr := C.PivotArea_GetFilters( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotAreaFilterCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotAreaFilterCollection) 

	return result, nil 
}
// Select the area with filters.
// Parameters:
//   axisType - int32 
//   fieldPosition - int32 
//   selectionType - int32 
// Returns:
//   void  
func (instance *PivotArea) Select(axistype PivotFieldType, fieldposition int32, selectiontype PivotTableSelectionType)  error {
	CGoReturnPtr := C.PivotArea_Select( instance.ptr, C.int( int32(axistype)), C.int(fieldposition), C.int( int32(selectiontype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether only the data values (in the data area of the view) for an item
// selection are selected and does not include the item labels.
// Returns:
//   bool  
func (instance *PivotArea) GetOnlyData()  (bool,  error)  {
	CGoReturnPtr := C.PivotArea_GetOnlyData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether only the data values (in the data area of the view) for an item
// selection are selected and does not include the item labels.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotArea) SetOnlyData(value bool)  error {
	CGoReturnPtr := C.PivotArea_SetOnlyData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether only the data labels for an item selection are selected.
// Returns:
//   bool  
func (instance *PivotArea) GetOnlyLabel()  (bool,  error)  {
	CGoReturnPtr := C.PivotArea_GetOnlyLabel( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether only the data labels for an item selection are selected.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotArea) SetOnlyLabel(value bool)  error {
	CGoReturnPtr := C.PivotArea_SetOnlyLabel( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the row grand total is included.
// Returns:
//   bool  
func (instance *PivotArea) IsRowGrandIncluded()  (bool,  error)  {
	CGoReturnPtr := C.PivotArea_IsRowGrandIncluded( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the row grand total is included.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotArea) SetIsRowGrandIncluded(value bool)  error {
	CGoReturnPtr := C.PivotArea_SetIsRowGrandIncluded( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the column grand total is included.
// Returns:
//   bool  
func (instance *PivotArea) IsColumnGrandIncluded()  (bool,  error)  {
	CGoReturnPtr := C.PivotArea_IsColumnGrandIncluded( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the column grand total is included.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotArea) SetIsColumnGrandIncluded(value bool)  error {
	CGoReturnPtr := C.PivotArea_SetIsColumnGrandIncluded( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the region of the PivotTable to which this rule applies.
// Returns:
//   int32  
func (instance *PivotArea) GetAxisType()  (PivotFieldType,  error)  {
	CGoReturnPtr := C.PivotArea_GetAxisType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the region of the PivotTable to which this rule applies.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotArea) SetAxisType(value PivotFieldType)  error {
	CGoReturnPtr := C.PivotArea_SetAxisType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of selection rule.
// Returns:
//   int32  
func (instance *PivotArea) GetRuleType()  (PivotAreaType,  error)  {
	CGoReturnPtr := C.PivotArea_GetRuleType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotAreaType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of selection rule.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotArea) SetRuleType(value PivotAreaType)  error {
	CGoReturnPtr := C.PivotArea_SetRuleType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the rule refers to an area that is in outline mode.
// Returns:
//   bool  
func (instance *PivotArea) IsOutline()  (bool,  error)  {
	CGoReturnPtr := C.PivotArea_IsOutline( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the rule refers to an area that is in outline mode.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotArea) SetIsOutline(value bool)  error {
	CGoReturnPtr := C.PivotArea_SetIsOutline( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotArea(pivotarea *PivotArea){
	runtime.SetFinalizer(pivotarea, nil)
	C.Delete_PivotArea(pivotarea.ptr)
	pivotarea.ptr = nil
}

// Class PivotAreaFilter 

// Represents the filter of <see cref="PivotArea"/> for <see cref="PivotTable"/>.
type PivotAreaFilter struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPivotAreaFilter() ( *PivotAreaFilter, error) {
	pivotareafilter := &PivotAreaFilter{}
	CGoReturnPtr := C.New_PivotAreaFilter()
	if CGoReturnPtr.error_no == 0 {
		pivotareafilter.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotareafilter, DeletePivotAreaFilter)
		return pivotareafilter, nil
	} else {
		pivotareafilter.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotareafilter, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotAreaFilter) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotAreaFilter_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets which subtotal is set for this filter.
// Parameters:
//   subtotalType - int32 
// Returns:
//   bool  
func (instance *PivotAreaFilter) IsSubtotalSet(subtotaltype PivotFieldSubtotalType)  (bool,  error)  {
	CGoReturnPtr := C.PivotAreaFilter_IsSubtotalSet( instance.ptr, C.int( int32(subtotaltype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Subtotal for the filter.
// Parameters:
//   subtotalType - int32 
//   shown - bool 
// Returns:
//   void  
func (instance *PivotAreaFilter) SetSubtotals(subtotaltype PivotFieldSubtotalType, shown bool)  error {
	CGoReturnPtr := C.PivotAreaFilter_SetSubtotals( instance.ptr, C.int( int32(subtotaltype)), C.bool(shown))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this field has selection.
// Only works when the PivotTable is in Outline view.
// Returns:
//   bool  
func (instance *PivotAreaFilter) GetSelected()  (bool,  error)  {
	CGoReturnPtr := C.PivotAreaFilter_GetSelected( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether this field has selection.
// Only works when the PivotTable is in Outline view.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotAreaFilter) SetSelected(value bool)  error {
	CGoReturnPtr := C.PivotAreaFilter_SetSelected( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotAreaFilter(pivotareafilter *PivotAreaFilter){
	runtime.SetFinalizer(pivotareafilter, nil)
	C.Delete_PivotAreaFilter(pivotareafilter.ptr)
	pivotareafilter.ptr = nil
}

// Class PivotAreaFilterCollection 

// Represents the list of filters for <see cref="PivotArea"/>
type PivotAreaFilterCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPivotAreaFilterCollection() ( *PivotAreaFilterCollection, error) {
	pivotareafiltercollection := &PivotAreaFilterCollection{}
	CGoReturnPtr := C.New_PivotAreaFilterCollection()
	if CGoReturnPtr.error_no == 0 {
		pivotareafiltercollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotareafiltercollection, DeletePivotAreaFilterCollection)
		return pivotareafiltercollection, nil
	} else {
		pivotareafiltercollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotareafiltercollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotAreaFilterCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotAreaFilterCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets filter from the list by the index.
// Parameters:
//   index - int32 
// Returns:
//   PivotAreaFilter  
func (instance *PivotAreaFilterCollection) Get(index int32)  (*PivotAreaFilter,  error)  {
	CGoReturnPtr := C.PivotAreaFilterCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotAreaFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotAreaFilter) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *PivotAreaFilterCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotAreaFilterCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotAreaFilterCollection(pivotareafiltercollection *PivotAreaFilterCollection){
	runtime.SetFinalizer(pivotareafiltercollection, nil)
	C.Delete_PivotAreaFilterCollection(pivotareafiltercollection.ptr)
	pivotareafiltercollection.ptr = nil
}

// Class PivotDateTimeRangeGroupSettings 

// Represents the field grouped by date time range.
type PivotDateTimeRangeGroupSettings struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PivotFieldGroupSettings 
func NewPivotDateTimeRangeGroupSettings(src *PivotFieldGroupSettings) ( *PivotDateTimeRangeGroupSettings, error) {
	pivotdatetimerangegroupsettings := &PivotDateTimeRangeGroupSettings{}
	CGoReturnPtr := C.New_PivotDateTimeRangeGroupSettings(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		pivotdatetimerangegroupsettings.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotdatetimerangegroupsettings, DeletePivotDateTimeRangeGroupSettings)
		return pivotdatetimerangegroupsettings, nil
	} else {
		pivotdatetimerangegroupsettings.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotdatetimerangegroupsettings, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotDateTimeRangeGroupSettings) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotDateTimeRangeGroupSettings_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the data time group type.
// Returns:
//   int32  
func (instance *PivotDateTimeRangeGroupSettings) GetType()  (PivotFieldGroupType,  error)  {
	CGoReturnPtr := C.PivotDateTimeRangeGroupSettings_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldGroupType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the start date time of the group.
// Returns:
//   Date  
func (instance *PivotDateTimeRangeGroupSettings) GetStart()  (*Date,  error)  {
	CGoReturnPtr := C.PivotDateTimeRangeGroupSettings_GetStart( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets the end date time of the group.
// Returns:
//   Date  
func (instance *PivotDateTimeRangeGroupSettings) GetEnd()  (*Date,  error)  {
	CGoReturnPtr := C.PivotDateTimeRangeGroupSettings_GetEnd( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets the internal of the group.
// Returns:
//   float64  
func (instance *PivotDateTimeRangeGroupSettings) GetInterval()  (float64,  error)  {
	CGoReturnPtr := C.PivotDateTimeRangeGroupSettings_GetInterval( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Check whether the field is grouped by the type.
// Parameters:
//   type - int32 
// Returns:
//   bool  
func (instance *PivotDateTimeRangeGroupSettings) IsGroupedBy(type_ PivotGroupByType)  (bool,  error)  {
	CGoReturnPtr := C.PivotDateTimeRangeGroupSettings_IsGroupedBy( instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}


func DeletePivotDateTimeRangeGroupSettings(pivotdatetimerangegroupsettings *PivotDateTimeRangeGroupSettings){
	runtime.SetFinalizer(pivotdatetimerangegroupsettings, nil)
	C.Delete_PivotDateTimeRangeGroupSettings(pivotdatetimerangegroupsettings.ptr)
	pivotdatetimerangegroupsettings.ptr = nil
}

// Class PivotDiscreteGroupSettings 

// Rrepsents the discrete group of pivot field
type PivotDiscreteGroupSettings struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PivotFieldGroupSettings 
func NewPivotDiscreteGroupSettings(src *PivotFieldGroupSettings) ( *PivotDiscreteGroupSettings, error) {
	pivotdiscretegroupsettings := &PivotDiscreteGroupSettings{}
	CGoReturnPtr := C.New_PivotDiscreteGroupSettings(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		pivotdiscretegroupsettings.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotdiscretegroupsettings, DeletePivotDiscreteGroupSettings)
		return pivotdiscretegroupsettings, nil
	} else {
		pivotdiscretegroupsettings.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotdiscretegroupsettings, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotDiscreteGroupSettings) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotDiscreteGroupSettings_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the group type.
// Returns:
//   int32  
func (instance *PivotDiscreteGroupSettings) GetType()  (PivotFieldGroupType,  error)  {
	CGoReturnPtr := C.PivotDiscreteGroupSettings_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldGroupType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}


func DeletePivotDiscreteGroupSettings(pivotdiscretegroupsettings *PivotDiscreteGroupSettings){
	runtime.SetFinalizer(pivotdiscretegroupsettings, nil)
	C.Delete_PivotDiscreteGroupSettings(pivotdiscretegroupsettings.ptr)
	pivotdiscretegroupsettings.ptr = nil
}

// Class PivotField 

// Represents a field in a PivotTable report.
type PivotField struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotField) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the pivot items of the pivot field
// Returns:
//   PivotItemCollection  
func (instance *PivotField) GetPivotItems()  (*PivotItemCollection,  error)  {
	CGoReturnPtr := C.PivotField_GetPivotItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotItemCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotItemCollection) 

	return result, nil 
}
// Gets the group settings of the pivot field.
// Returns:
//   PivotFieldGroupSettings  
func (instance *PivotField) GetGroupSettings()  (*PivotFieldGroupSettings,  error)  {
	CGoReturnPtr := C.PivotField_GetGroupSettings( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldGroupSettings{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldGroupSettings) 

	return result, nil 
}
// Init the pivot items of the pivot field
// Returns:
//   void  
func (instance *PivotField) InitPivotItems()  error {
	CGoReturnPtr := C.PivotField_InitPivotItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Automatically group the field with internal
// Parameters:
//   interval - float64 
//   newField - bool 
// Returns:
//   void  
func (instance *PivotField) GroupBy_Double_Bool(interval float64, newfield bool)  error {
	CGoReturnPtr := C.PivotField_GroupBy_Double_Boolean( instance.ptr, C.double(interval), C.bool(newfield))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Group the file by number.
// Parameters:
//   start - float64 
//   end - float64 
//   interval - float64 
//   newField - bool 
// Returns:
//   bool  
func (instance *PivotField) GroupBy_Double_Double_Double_Bool(start float64, end float64, interval float64, newfield bool)  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GroupBy_Double_Double_Double_Boolean( instance.ptr, C.double(start), C.double(end), C.double(interval), C.bool(newfield))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Ungroup the pivot field.
// Returns:
//   void  
func (instance *PivotField) Ungroup()  error {
	CGoReturnPtr := C.PivotField_Ungroup( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the pivot filter of the pivot field by type
// Parameters:
//   type - int32 
// Returns:
//   PivotFilter  
func (instance *PivotField) GetPivotFilterByType(type_ PivotFilterType)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotField_GetPivotFilterByType( instance.ptr, C.int( int32(type_)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Clears filter setting on this pivot field.
// Returns:
//   void  
func (instance *PivotField) ClearFilter()  error {
	CGoReturnPtr := C.PivotField_ClearFilter( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Filters by values of data pivot field.
// Parameters:
//   valueFieldIndex - int32 
//   type - int32 
//   isTop - bool 
//   itemCount - int32 
// Returns:
//   PivotFilter  
func (instance *PivotField) FilterTop10(valuefieldindex int32, type_ PivotFilterType, istop bool, itemcount int32)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotField_FilterTop10( instance.ptr, C.int(valuefieldindex), C.int( int32(type_)), C.bool(istop), C.int(itemcount))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by values of data pivot field.
// Parameters:
//   valueFieldIndex - int32 
//   type - int32 
//   value1 - float64 
//   value2 - float64 
// Returns:
//   PivotFilter  
func (instance *PivotField) FilterByValue(valuefieldindex int32, type_ PivotFilterType, value1 float64, value2 float64)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotField_FilterByValue( instance.ptr, C.int(valuefieldindex), C.int( int32(type_)), C.double(value1), C.double(value2))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by captions of row or column pivot field.
// Parameters:
//   type - int32 
//   label1 - string 
//   label2 - string 
// Returns:
//   PivotFilter  
func (instance *PivotField) FilterByLabel(type_ PivotFilterType, label1 string, label2 string)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotField_FilterByLabel( instance.ptr, C.int( int32(type_)), C.CString(label1), C.CString(label2))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by date setting of row or column pivot field.
// Parameters:
//   type - int32 
//   dateTime1 - Date 
//   dateTime2 - Date 
// Returns:
//   PivotFilter  
func (instance *PivotField) FilterByDate(type_ PivotFilterType, datetime1 *Date, datetime2 *Date)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotField_FilterByDate( instance.ptr, C.int( int32(type_)), datetime1.ptr, datetime2.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Indicates whether the specified PivotTable field is calculated field.
// Returns:
//   bool  
func (instance *PivotField) IsCalculatedField()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsCalculatedField( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets formula of the calculated field .
// Returns:
//   string  
func (instance *PivotField) GetFormula()  (string,  error)  {
	CGoReturnPtr := C.PivotField_GetFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this field represents values fields.
// Returns:
//   bool  
func (instance *PivotField) IsValueFields()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsValueFields( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the PivotField index in the base PivotFields.
// Returns:
//   int32  
func (instance *PivotField) GetBaseIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetBaseIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the PivotField index in the base PivotFields.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetBaseIndex(value int32)  error {
	CGoReturnPtr := C.PivotField_SetBaseIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the index of <see cref="PivotField"/> in the region.
// Returns:
//   int32  
func (instance *PivotField) GetPosition()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetPosition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the region of the PivotTable that this field is displayed.
// Returns:
//   int32  
func (instance *PivotField) GetRegionType()  (PivotFieldType,  error)  {
	CGoReturnPtr := C.PivotField_GetRegionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the name of PivotField.
// Returns:
//   string  
func (instance *PivotField) GetName()  (string,  error)  {
	CGoReturnPtr := C.PivotField_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the name of PivotField.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotField) SetName(value string)  error {
	CGoReturnPtr := C.PivotField_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the PivotField display name.
// Returns:
//   string  
func (instance *PivotField) GetDisplayName()  (string,  error)  {
	CGoReturnPtr := C.PivotField_GetDisplayName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the PivotField display name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotField) SetDisplayName(value string)  error {
	CGoReturnPtr := C.PivotField_SetDisplayName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets whether the specified field shows that subtotals.
// Parameters:
//   subtotalType - int32 
//   shown - bool 
// Returns:
//   void  
func (instance *PivotField) SetSubtotals(subtotaltype PivotFieldSubtotalType, shown bool)  error {
	CGoReturnPtr := C.PivotField_SetSubtotals( instance.ptr, C.int( int32(subtotaltype)), C.bool(shown))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing specified subtotal.
// Parameters:
//   subtotalType - int32 
// Returns:
//   bool  
func (instance *PivotField) GetSubtotals(subtotaltype PivotFieldSubtotalType)  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetSubtotals( instance.ptr, C.int( int32(subtotaltype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field shows automatic subtotals. Default is true.
// Returns:
//   bool  
func (instance *PivotField) IsAutoSubtotals()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsAutoSubtotals( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field shows automatic subtotals. Default is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsAutoSubtotals(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsAutoSubtotals( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified field can be dragged to the column position.
// The default value is true.
// Returns:
//   bool  
func (instance *PivotField) GetDragToColumn()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetDragToColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field can be dragged to the column position.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetDragToColumn(value bool)  error {
	CGoReturnPtr := C.PivotField_SetDragToColumn( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified field can be dragged to the hide position.
// The default value is true.
// Returns:
//   bool  
func (instance *PivotField) GetDragToHide()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetDragToHide( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field can be dragged to the hide position.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetDragToHide(value bool)  error {
	CGoReturnPtr := C.PivotField_SetDragToHide( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified field can be dragged to the row position.
// The default value is true.
// Returns:
//   bool  
func (instance *PivotField) GetDragToRow()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetDragToRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field can be dragged to the row position.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetDragToRow(value bool)  error {
	CGoReturnPtr := C.PivotField_SetDragToRow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified field can be dragged to the page position.
// The default value is true.
// Returns:
//   bool  
func (instance *PivotField) GetDragToPage()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetDragToPage( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field can be dragged to the page position.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetDragToPage(value bool)  error {
	CGoReturnPtr := C.PivotField_SetDragToPage( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified field can be dragged to the data position.
// The default value is true.
// Returns:
//   bool  
func (instance *PivotField) GetDragToData()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetDragToData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified field can be dragged to the data position.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetDragToData(value bool)  error {
	CGoReturnPtr := C.PivotField_SetDragToData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// indicates whether the field can have multiple items
// selected in the page field
// The default value is false.
// Returns:
//   bool  
func (instance *PivotField) IsMultipleItemSelectionAllowed()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsMultipleItemSelectionAllowed( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// indicates whether the field can have multiple items
// selected in the page field
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsMultipleItemSelectionAllowed(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsMultipleItemSelectionAllowed( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether repeating labels of the field in the region.
// The default value is false.
// Returns:
//   bool  
func (instance *PivotField) IsRepeatItemLabels()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsRepeatItemLabels( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether repeating labels of the field in the region.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsRepeatItemLabels(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsRepeatItemLabels( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether including new items to the field in manual filter.
// The default value is false.
// Returns:
//   bool  
func (instance *PivotField) IsIncludeNewItemsInFilter()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsIncludeNewItemsInFilter( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether including new items to the field in manual filter.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsIncludeNewItemsInFilter(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsIncludeNewItemsInFilter( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether inserting page breaks after each item.
// The default value is false.
// Returns:
//   bool  
func (instance *PivotField) IsInsertPageBreaksBetweenItems()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsInsertPageBreaksBetweenItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether inserting page breaks after each item.
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsInsertPageBreaksBetweenItems(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsInsertPageBreaksBetweenItems( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether all items displays in the PivotTable report,
// even if they don't contain summary data.
// show items with no data
// The default value is false.
// Returns:
//   bool  
func (instance *PivotField) GetShowAllItems()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetShowAllItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether all items displays in the PivotTable report,
// even if they don't contain summary data.
// show items with no data
// The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetShowAllItems(value bool)  error {
	CGoReturnPtr := C.PivotField_SetShowAllItems( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sorts this pivot field.
// Parameters:
//   sortType - int32 
//   fieldSortedBy - int32 
// Returns:
//   void  
func (instance *PivotField) SortBy_SortOrder_Int(sorttype SortOrder, fieldsortedby int32)  error {
	CGoReturnPtr := C.PivotField_SortBy_SortOrder_Integer( instance.ptr, C.int( int32(sorttype)), C.int(fieldsortedby))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sorts this pivot field.
// Parameters:
//   sortType - int32 
//   fieldSortedBy - int32 
//   dataType - int32 
//   cellName - string 
// Returns:
//   void  
func (instance *PivotField) SortBy_SortOrder_Int_PivotLineType_String(sorttype SortOrder, fieldsortedby int32, datatype PivotLineType, cellname string)  error {
	CGoReturnPtr := C.PivotField_SortBy_SortOrder_Integer_PivotLineType_String( instance.ptr, C.int( int32(sorttype)), C.int(fieldsortedby), C.int( int32(datatype)), C.CString(cellname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether a sort operation that will be applied to this pivot field is an autosort operation or a simple data sort.
// Returns:
//   bool  
func (instance *PivotField) GetNonAutoSortDefault()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetNonAutoSortDefault( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether a sort operation that will be applied to this pivot field is an autosort operation or a simple data sort.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetNonAutoSortDefault(value bool)  error {
	CGoReturnPtr := C.PivotField_SetNonAutoSortDefault( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified PivotTable field is automatically sorted.
// Returns:
//   bool  
func (instance *PivotField) IsAutoSort()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsAutoSort( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified PivotTable field is automatically sorted.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsAutoSort(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsAutoSort( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified PivotTable field is autosorted ascending.
// Returns:
//   bool  
func (instance *PivotField) IsAscendSort()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsAscendSort( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified PivotTable field is autosorted ascending.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsAscendSort(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsAscendSort( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets all settings of auto sorting
// Returns:
//   PivotFieldSortSetting  
func (instance *PivotField) GetSortSetting()  (*PivotFieldSortSetting,  error)  {
	CGoReturnPtr := C.PivotField_GetSortSetting( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldSortSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldSortSetting) 

	return result, nil 
}
// Represents the index of field which is auto sorted.
// -1 means PivotField itself,others means the position of the data fields.
// Returns:
//   int32  
func (instance *PivotField) GetAutoSortField()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetAutoSortField( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the index of field which is auto sorted.
// -1 means PivotField itself,others means the position of the data fields.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetAutoSortField(value int32)  error {
	CGoReturnPtr := C.PivotField_SetAutoSortField( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified PivotTable field is automatically shown,only valid for excel 2003.
// Returns:
//   bool  
func (instance *PivotField) IsAutoShow()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsAutoShow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified PivotTable field is automatically shown,only valid for excel 2003.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsAutoShow(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsAutoShow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the specified PivotTable field is autoshown ascending.
// Returns:
//   bool  
func (instance *PivotField) IsAscendShow()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsAscendShow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the specified PivotTable field is autoshown ascending.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetIsAscendShow(value bool)  error {
	CGoReturnPtr := C.PivotField_SetIsAscendShow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represent the number of top or bottom items
// that are automatically shown in the specified PivotTable field.
// Returns:
//   int32  
func (instance *PivotField) GetAutoShowCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetAutoShowCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represent the number of top or bottom items
// that are automatically shown in the specified PivotTable field.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetAutoShowCount(value int32)  error {
	CGoReturnPtr := C.PivotField_SetAutoShowCount( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents auto show field index. -1 means PivotField itself.
// It should be the index of the data fields.
// Returns:
//   int32  
func (instance *PivotField) GetAutoShowField()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetAutoShowField( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents auto show field index. -1 means PivotField itself.
// It should be the index of the data fields.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetAutoShowField(value int32)  error {
	CGoReturnPtr := C.PivotField_SetAutoShowField( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the function used to summarize the PivotTable data field.
// Returns:
//   int32  
func (instance *PivotField) GetFunction()  (ConsolidationFunction,  error)  {
	CGoReturnPtr := C.PivotField_GetFunction( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToConsolidationFunction(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the function used to summarize the PivotTable data field.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetFunction(value ConsolidationFunction)  error {
	CGoReturnPtr := C.PivotField_SetFunction( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Shows values of data field as different display format when the ShowDataAs calculation is in use.
// Parameters:
//   displayFormat - int32 
//   baseField - int32 
//   baseItemPositionType - int32 
//   baseItem - int32 
// Returns:
//   void  
func (instance *PivotField) ShowValuesAs(displayformat PivotFieldDataDisplayFormat, basefield int32, baseitempositiontype PivotItemPositionType, baseitem int32)  error {
	CGoReturnPtr := C.PivotField_ShowValuesAs( instance.ptr, C.int( int32(displayformat)), C.int(basefield), C.int( int32(baseitempositiontype)), C.int(baseitem))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the settings of showing values as when the ShowDataAs calculation is in use.
// Returns:
//   PivotShowValuesSetting  
func (instance *PivotField) GetShowValuesSetting()  (*PivotShowValuesSetting,  error)  {
	CGoReturnPtr := C.PivotField_GetShowValuesSetting( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotShowValuesSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotShowValuesSetting) 

	return result, nil 
}
// Represents the current page item showing for the page field (valid only for page fields).
// Returns:
//   int32  
func (instance *PivotField) GetCurrentPageItem()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetCurrentPageItem( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the current page item showing for the page field (valid only for page fields).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetCurrentPageItem(value int32)  error {
	CGoReturnPtr := C.PivotField_SetCurrentPageItem( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the built-in display format of numbers and dates.
// Returns:
//   int32  
func (instance *PivotField) GetNumber()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetNumber( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the built-in display format of numbers and dates.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotField) SetNumber(value int32)  error {
	CGoReturnPtr := C.PivotField_SetNumber( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether inserting blank line after each item.
// Returns:
//   bool  
func (instance *PivotField) GetInsertBlankRow()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetInsertBlankRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether inserting blank line after each item.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetInsertBlankRow(value bool)  error {
	CGoReturnPtr := C.PivotField_SetInsertBlankRow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// when ShowInOutlineForm is true, then display subtotals at the top of the list of items instead of at the bottom
// Returns:
//   bool  
func (instance *PivotField) GetShowSubtotalAtTop()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetShowSubtotalAtTop( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// when ShowInOutlineForm is true, then display subtotals at the top of the list of items instead of at the bottom
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetShowSubtotalAtTop(value bool)  error {
	CGoReturnPtr := C.PivotField_SetShowSubtotalAtTop( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether layout this field in outline form on the Pivot Table view
// Returns:
//   bool  
func (instance *PivotField) GetShowInOutlineForm()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetShowInOutlineForm( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether layout this field in outline form on the Pivot Table view
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetShowInOutlineForm(value bool)  error {
	CGoReturnPtr := C.PivotField_SetShowInOutlineForm( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the custom display format of numbers and dates.
// Returns:
//   string  
func (instance *PivotField) Get_NumberFormat()  (string,  error)  {
	CGoReturnPtr := C.PivotField_Get_NumberFormat( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the custom display format of numbers and dates.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotField) SetNumberFormat(value string)  error {
	CGoReturnPtr := C.PivotField_SetNumberFormat( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets whether the specific PivotItem is hidden.
// Parameters:
//   index - int32 
// Returns:
//   bool  
func (instance *PivotField) IsHiddenItem(index int32)  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsHiddenItem( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Sets whether the specific PivotItem in a data field is hidden.
// Parameters:
//   index - int32 
//   isHidden - bool 
// Returns:
//   void  
func (instance *PivotField) HideItem_Int_Bool(index int32, ishidden bool)  error {
	CGoReturnPtr := C.PivotField_HideItem_Integer_Boolean( instance.ptr, C.int(index), C.bool(ishidden))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets whether hidding the detail of  the specific PivotItem..
// Parameters:
//   index - int32 
// Returns:
//   bool  
func (instance *PivotField) IsHiddenItemDetail(index int32)  (bool,  error)  {
	CGoReturnPtr := C.PivotField_IsHiddenItemDetail( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Sets whether the specific PivotItem in a pivot field is hidden detail.
// Parameters:
//   index - int32 
//   isHiddenDetail - bool 
// Returns:
//   void  
func (instance *PivotField) HideItemDetail(index int32, ishiddendetail bool)  error {
	CGoReturnPtr := C.PivotField_HideItemDetail( instance.ptr, C.int(index), C.bool(ishiddendetail))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets whether the PivotItems in a pivot field is hidden detail.That is collapse/expand this field.
// Parameters:
//   isHiddenDetail - bool 
// Returns:
//   void  
func (instance *PivotField) HideDetail(ishiddendetail bool)  error {
	CGoReturnPtr := C.PivotField_HideDetail( instance.ptr, C.bool(ishiddendetail))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets whether the specific PivotItem in a data field is hidden.
// Parameters:
//   itemValue - string 
//   isHidden - bool 
// Returns:
//   void  
func (instance *PivotField) HideItem_String_Bool(itemvalue string, ishidden bool)  error {
	CGoReturnPtr := C.PivotField_HideItem_String_Boolean( instance.ptr, C.CString(itemvalue), C.bool(ishidden))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the count of the base items in this pivot field.
// Returns:
//   int32  
func (instance *PivotField) GetItemCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotField_GetItemCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Add a calculated formula item to the pivot field.
// Parameters:
//   name - string 
//   formula - string 
// Returns:
//   void  
func (instance *PivotField) AddCalculatedItem(name string, formula string)  error {
	CGoReturnPtr := C.PivotField_AddCalculatedItem( instance.ptr, C.CString(name), C.CString(formula))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether display labels from the next field in the same column on the Pivot Table view
// Returns:
//   bool  
func (instance *PivotField) GetShowCompact()  (bool,  error)  {
	CGoReturnPtr := C.PivotField_GetShowCompact( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether display labels from the next field in the same column on the Pivot Table view
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotField) SetShowCompact(value bool)  error {
	CGoReturnPtr := C.PivotField_SetShowCompact( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotField(pivotfield *PivotField){
	runtime.SetFinalizer(pivotfield, nil)
	C.Delete_PivotField(pivotfield.ptr)
	pivotfield.ptr = nil
}

// Class PivotFieldCollection 

// Represents a collection of all the PivotField objects
// in the PivotTable's specific PivotFields type.
type PivotFieldCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFieldCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the PivotFields type.
// Returns:
//   int32  
func (instance *PivotFieldCollection) GetType()  (PivotFieldType,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the count of the pivotFields.
// Returns:
//   int32  
func (instance *PivotFieldCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the PivotField Object at the specific index.
// Parameters:
//   index - int32 
// Returns:
//   PivotField  
func (instance *PivotFieldCollection) Get_Int(index int32)  (*PivotField,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotField{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotField) 

	return result, nil 
}
// Gets the PivotField Object of the specific name.
// Parameters:
//   name - string 
// Returns:
//   PivotField  
func (instance *PivotFieldCollection) Get_String(name string)  (*PivotField,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotField{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotField) 

	return result, nil 
}
// Adds a PivotField Object to the specific type PivotFields.
// Parameters:
//   baseFieldIndex - int32 
// Returns:
//   int32  
func (instance *PivotFieldCollection) AddByBaseIndex(basefieldindex int32)  (int32,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_AddByBaseIndex( instance.ptr, C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a PivotField Object to the specific type PivotFields.
// Parameters:
//   pivotField - PivotField 
// Returns:
//   int32  
func (instance *PivotFieldCollection) Add(pivotfield *PivotField)  (int32,  error)  {
	CGoReturnPtr := C.PivotFieldCollection_Add( instance.ptr, pivotfield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// clear all fields of PivotFieldCollection
// Returns:
//   void  
func (instance *PivotFieldCollection) Clear()  error {
	CGoReturnPtr := C.PivotFieldCollection_Clear( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Moves the PivotField from current position to destination position
// Parameters:
//   currPos - int32 
//   destPos - int32 
// Returns:
//   void  
func (instance *PivotFieldCollection) Move(currpos int32, destpos int32)  error {
	CGoReturnPtr := C.PivotFieldCollection_Move( instance.ptr, C.int(currpos), C.int(destpos))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotFieldCollection(pivotfieldcollection *PivotFieldCollection){
	runtime.SetFinalizer(pivotfieldcollection, nil)
	C.Delete_PivotFieldCollection(pivotfieldcollection.ptr)
	pivotfieldcollection.ptr = nil
}

// Class PivotFieldGroupSettings 

// Represents the group setting of pivot field.
type PivotFieldGroupSettings struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPivotFieldGroupSettings() ( *PivotFieldGroupSettings, error) {
	pivotfieldgroupsettings := &PivotFieldGroupSettings{}
	CGoReturnPtr := C.New_PivotFieldGroupSettings()
	if CGoReturnPtr.error_no == 0 {
		pivotfieldgroupsettings.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotfieldgroupsettings, DeletePivotFieldGroupSettings)
		return pivotfieldgroupsettings, nil
	} else {
		pivotfieldgroupsettings.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotfieldgroupsettings, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFieldGroupSettings) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFieldGroupSettings_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the group type of pivot field.
// Returns:
//   int32  
func (instance *PivotFieldGroupSettings) GetType()  (PivotFieldGroupType,  error)  {
	CGoReturnPtr := C.PivotFieldGroupSettings_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldGroupType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}


func DeletePivotFieldGroupSettings(pivotfieldgroupsettings *PivotFieldGroupSettings){
	runtime.SetFinalizer(pivotfieldgroupsettings, nil)
	C.Delete_PivotFieldGroupSettings(pivotfieldgroupsettings.ptr)
	pivotfieldgroupsettings.ptr = nil
}

// Class PivotFieldSortSetting 

// Represents the setting of sorting pivot fields.
type PivotFieldSortSetting struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFieldSortSetting) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the <see cref="SortOrder"/>.
// Returns:
//   int32  
func (instance *PivotFieldSortSetting) GetSortType()  (SortOrder,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_GetSortType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSortOrder(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether sorting the field by itself or data field.
// Returns:
//   bool  
func (instance *PivotFieldSortSetting) IsSortByLabels()  (bool,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_IsSortByLabels( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the index of the field sorted by.
// -1 means sorting the PivotField by the labels,others means sorting by the data field.
// Returns:
//   int32  
func (instance *PivotFieldSortSetting) GetFieldIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_GetFieldIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// The pivot line type sorted by.
// Returns:
//   int32  
func (instance *PivotFieldSortSetting) GetLineTypeSortedBy()  (PivotLineType,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_GetLineTypeSortedBy( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotLineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether a simple data sort operation will be applied.
// Returns:
//   bool  
func (instance *PivotFieldSortSetting) IsSimpleSort()  (bool,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_IsSimpleSort( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Sorts by the values in which row or column.
// Returns:
//   string  
func (instance *PivotFieldSortSetting) GetCell()  (string,  error)  {
	CGoReturnPtr := C.PivotFieldSortSetting_GetCell( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotFieldSortSetting(pivotfieldsortsetting *PivotFieldSortSetting){
	runtime.SetFinalizer(pivotfieldsortsetting, nil)
	C.Delete_PivotFieldSortSetting(pivotfieldsortsetting.ptr)
	pivotfieldsortsetting.ptr = nil
}

// Class PivotFilter 

// Represents a PivotFilter in PivotFilter Collection.
type PivotFilter struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFilter) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFilter_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether uses whole days in its filtering criteria.
// Returns:
//   bool  
func (instance *PivotFilter) GetUseWholeDay()  (bool,  error)  {
	CGoReturnPtr := C.PivotFilter_GetUseWholeDay( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether uses whole days in its filtering criteria.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotFilter) SetUseWholeDay(value bool)  error {
	CGoReturnPtr := C.PivotFilter_SetUseWholeDay( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the autofilter type of the pivot filter.
// Returns:
//   int32  
func (instance *PivotFilter) GetFilterType()  (PivotFilterType,  error)  {
	CGoReturnPtr := C.PivotFilter_GetFilterType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFilterType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the index of source field which this pivot filter is applied to.
// Returns:
//   int32  
func (instance *PivotFilter) GetFieldIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotFilter_GetFieldIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets top 10 setting of the filter.
// Returns:
//   Top10Filter  
func (instance *PivotFilter) GetTop10Value()  (*Top10Filter,  error)  {
	CGoReturnPtr := C.PivotFilter_GetTop10Value( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Top10Filter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTop10Filter) 

	return result, nil 
}
// Gets the category of this filter.
// Returns:
//   int32  
func (instance *PivotFilter) GetFilterCategory()  (FilterCategory,  error)  {
	CGoReturnPtr := C.PivotFilter_GetFilterCategory( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFilterCategory(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the string value1 of the label pivot filter.
// Returns:
//   string  
func (instance *PivotFilter) GetValue1()  (string,  error)  {
	CGoReturnPtr := C.PivotFilter_GetValue1( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the string value1 of the label pivot filter.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotFilter) SetValue1(value string)  error {
	CGoReturnPtr := C.PivotFilter_SetValue1( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the string value2 of the label pivot filter.
// Returns:
//   string  
func (instance *PivotFilter) GetValue2()  (string,  error)  {
	CGoReturnPtr := C.PivotFilter_GetValue2( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the string value2 of the label pivot filter.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotFilter) SetValue2(value string)  error {
	CGoReturnPtr := C.PivotFilter_SetValue2( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the index of value field in the value region.
// Returns:
//   int32  
func (instance *PivotFilter) GetValueFieldIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotFilter_GetValueFieldIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the index of value field in the value region.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotFilter) SetValueFieldIndex(value int32)  error {
	CGoReturnPtr := C.PivotFilter_SetValueFieldIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the index of the measure cube field.
// this property is used only by filters in OLAP pivots and specifies on which measure a value filter should apply.
// Returns:
//   int32  
func (instance *PivotFilter) GetMeasureCubeFieldIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotFilter_GetMeasureCubeFieldIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the member property field index of the pivot filter.
// Returns:
//   int32  
func (instance *PivotFilter) GetMemberPropertyFieldIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotFilter_GetMemberPropertyFieldIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the member property field index of the pivot filter.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotFilter) SetMemberPropertyFieldIndex(value int32)  error {
	CGoReturnPtr := C.PivotFilter_SetMemberPropertyFieldIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the name of the pivot filter.
// Returns:
//   string  
func (instance *PivotFilter) GetName()  (string,  error)  {
	CGoReturnPtr := C.PivotFilter_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the pivot filter.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotFilter) SetName(value string)  error {
	CGoReturnPtr := C.PivotFilter_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the Evaluation Order of the pivot filter.
// Returns:
//   int32  
func (instance *PivotFilter) GetEvaluationOrder()  (int32,  error)  {
	CGoReturnPtr := C.PivotFilter_GetEvaluationOrder( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the Evaluation Order of the pivot filter.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotFilter) SetEvaluationOrder(value int32)  error {
	CGoReturnPtr := C.PivotFilter_SetEvaluationOrder( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotFilter(pivotfilter *PivotFilter){
	runtime.SetFinalizer(pivotfilter, nil)
	C.Delete_PivotFilter(pivotfilter.ptr)
	pivotfilter.ptr = nil
}

// Class PivotFilterCollection 

// Represents a collection of all the PivotFilter objects
type PivotFilterCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFilterCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the pivotfilter object at the specific index.
// Parameters:
//   index - int32 
// Returns:
//   PivotFilter  
func (instance *PivotFilterCollection) Get(index int32)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by values of data pivot field.
// Parameters:
//   baseFieldIndex - int32 
//   valueFieldIndex - int32 
//   type - int32 
//   isTop - bool 
//   itemCount - int32 
// Returns:
//   PivotFilter  
func (instance *PivotFilterCollection) AddTop10Filter(basefieldindex int32, valuefieldindex int32, type_ PivotFilterType, istop bool, itemcount int32)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_AddTop10Filter( instance.ptr, C.int(basefieldindex), C.int(valuefieldindex), C.int( int32(type_)), C.bool(istop), C.int(itemcount))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by values of data pivot field.
// Parameters:
//   baseFieldIndex - int32 
//   valueFieldIndex - int32 
//   type - int32 
//   value1 - float64 
//   value2 - float64 
// Returns:
//   PivotFilter  
func (instance *PivotFilterCollection) AddValueFilter(basefieldindex int32, valuefieldindex int32, type_ PivotFilterType, value1 float64, value2 float64)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_AddValueFilter( instance.ptr, C.int(basefieldindex), C.int(valuefieldindex), C.int( int32(type_)), C.double(value1), C.double(value2))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by captions of row or column pivot field.
// Parameters:
//   baseFieldIndex - int32 
//   type - int32 
//   label1 - string 
//   label2 - string 
// Returns:
//   PivotFilter  
func (instance *PivotFilterCollection) AddLabelFilter(basefieldindex int32, type_ PivotFilterType, label1 string, label2 string)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_AddLabelFilter( instance.ptr, C.int(basefieldindex), C.int( int32(type_)), C.CString(label1), C.CString(label2))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Filters by date setting of row or column pivot field.
// Parameters:
//   baseFieldIndex - int32 
//   type - int32 
//   dateTime1 - Date 
//   dateTime2 - Date 
// Returns:
//   PivotFilter  
func (instance *PivotFilterCollection) AddDateFilter(basefieldindex int32, type_ PivotFilterType, datetime1 *Date, datetime2 *Date)  (*PivotFilter,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_AddDateFilter( instance.ptr, C.int(basefieldindex), C.int( int32(type_)), datetime1.ptr, datetime2.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilter{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilter) 

	return result, nil 
}
// Clear PivotFilter from the specific PivotField
// Parameters:
//   fieldIndex - int32 
// Returns:
//   void  
func (instance *PivotFilterCollection) ClearFilter(fieldindex int32)  error {
	CGoReturnPtr := C.PivotFilterCollection_ClearFilter( instance.ptr, C.int(fieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *PivotFilterCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotFilterCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotFilterCollection(pivotfiltercollection *PivotFilterCollection){
	runtime.SetFinalizer(pivotfiltercollection, nil)
	C.Delete_PivotFilterCollection(pivotfiltercollection.ptr)
	pivotfiltercollection.ptr = nil
}

// Class PivotFormatCondition 

// Represents a PivotTable Format Condition in PivotFormatCondition Collection.
type PivotFormatCondition struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFormatCondition) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFormatCondition_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Adds PivotTable conditional format limit in the data fields.
// Parameters:
//   fieldName - string 
// Returns:
//   void  
func (instance *PivotFormatCondition) AddDataAreaCondition_String(fieldname string)  error {
	CGoReturnPtr := C.PivotFormatCondition_AddDataAreaCondition_String( instance.ptr, C.CString(fieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds PivotTable conditional format limit in the data fields.
// Parameters:
//   dataField - PivotField 
// Returns:
//   void  
func (instance *PivotFormatCondition) AddDataAreaCondition_PivotField(datafield *PivotField)  error {
	CGoReturnPtr := C.PivotFormatCondition_AddDataAreaCondition_PivotField( instance.ptr, datafield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds PivotTable conditional format limit in the row fields.
// Parameters:
//   fieldName - string 
// Returns:
//   void  
func (instance *PivotFormatCondition) AddRowAreaCondition_String(fieldname string)  error {
	CGoReturnPtr := C.PivotFormatCondition_AddRowAreaCondition_String( instance.ptr, C.CString(fieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds PivotTable conditional format limit in the row fields.
// Parameters:
//   rowField - PivotField 
// Returns:
//   void  
func (instance *PivotFormatCondition) AddRowAreaCondition_PivotField(rowfield *PivotField)  error {
	CGoReturnPtr := C.PivotFormatCondition_AddRowAreaCondition_PivotField( instance.ptr, rowfield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds PivotTable conditional format limit in the column fields.
// Parameters:
//   fieldName - string 
// Returns:
//   void  
func (instance *PivotFormatCondition) AddColumnAreaCondition_String(fieldname string)  error {
	CGoReturnPtr := C.PivotFormatCondition_AddColumnAreaCondition_String( instance.ptr, C.CString(fieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds PivotTable conditional format limit in the column fields.
// Parameters:
//   columnField - PivotField 
// Returns:
//   void  
func (instance *PivotFormatCondition) AddColumnAreaCondition_PivotField(columnfield *PivotField)  error {
	CGoReturnPtr := C.PivotFormatCondition_AddColumnAreaCondition_PivotField( instance.ptr, columnfield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Sets conditional areas of PivotFormatCondition object.
// Returns:
//   void  
func (instance *PivotFormatCondition) SetConditionalAreas()  error {
	CGoReturnPtr := C.PivotFormatCondition_SetConditionalAreas( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get and set scope type for the pivot table condition format .
// Returns:
//   int32  
func (instance *PivotFormatCondition) GetScopeType()  (PivotConditionFormatScopeType,  error)  {
	CGoReturnPtr := C.PivotFormatCondition_GetScopeType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotConditionFormatScopeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Get and set scope type for the pivot table condition format .
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotFormatCondition) SetScopeType(value PivotConditionFormatScopeType)  error {
	CGoReturnPtr := C.PivotFormatCondition_SetScopeType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get and set rule type for the pivot table condition format .
// Returns:
//   int32  
func (instance *PivotFormatCondition) GetRuleType()  (PivotConditionFormatRuleType,  error)  {
	CGoReturnPtr := C.PivotFormatCondition_GetRuleType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotConditionFormatRuleType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Get and set rule type for the pivot table condition format .
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotFormatCondition) SetRuleType(value PivotConditionFormatRuleType)  error {
	CGoReturnPtr := C.PivotFormatCondition_SetRuleType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get formatconditions for the pivot table condition format .
// Returns:
//   FormatConditionCollection  
func (instance *PivotFormatCondition) GetFormatConditions()  (*FormatConditionCollection,  error)  {
	CGoReturnPtr := C.PivotFormatCondition_GetFormatConditions( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FormatConditionCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFormatConditionCollection) 

	return result, nil 
}


func DeletePivotFormatCondition(pivotformatcondition *PivotFormatCondition){
	runtime.SetFinalizer(pivotformatcondition, nil)
	C.Delete_PivotFormatCondition(pivotformatcondition.ptr)
	pivotformatcondition.ptr = nil
}

// Class PivotFormatConditionCollection 

// Represents PivotTable Format Conditions.
type PivotFormatConditionCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotFormatConditionCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotFormatConditionCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Adds a pivot FormatCondition to the collection.
// Returns:
//   int32  
func (instance *PivotFormatConditionCollection) Add()  (int32,  error)  {
	CGoReturnPtr := C.PivotFormatConditionCollection_Add( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the pivot FormatCondition object at the specific index.
// Parameters:
//   index - int32 
// Returns:
//   PivotFormatCondition  
func (instance *PivotFormatConditionCollection) Get(index int32)  (*PivotFormatCondition,  error)  {
	CGoReturnPtr := C.PivotFormatConditionCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFormatCondition{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFormatCondition) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *PivotFormatConditionCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotFormatConditionCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotFormatConditionCollection(pivotformatconditioncollection *PivotFormatConditionCollection){
	runtime.SetFinalizer(pivotformatconditioncollection, nil)
	C.Delete_PivotFormatConditionCollection(pivotformatconditioncollection.ptr)
	pivotformatconditioncollection.ptr = nil
}

// Class PivotItem 

// Represents a item in a PivotField report.
type PivotItem struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotItem) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotItem_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and Sets whether the pivot item is hidden.
// Returns:
//   bool  
func (instance *PivotItem) IsHidden()  (bool,  error)  {
	CGoReturnPtr := C.PivotItem_IsHidden( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and Sets whether the pivot item is hidden.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotItem) SetIsHidden(value bool)  error {
	CGoReturnPtr := C.PivotItem_SetIsHidden( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifying the position index in all the PivotItems,not the PivotItems under the same parent node.
// Returns:
//   int32  
func (instance *PivotItem) GetPosition()  (int32,  error)  {
	CGoReturnPtr := C.PivotItem_GetPosition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifying the position index in all the PivotItems,not the PivotItems under the same parent node.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotItem) SetPosition(value int32)  error {
	CGoReturnPtr := C.PivotItem_SetPosition( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifying the position index in the PivotItems under the same parent node.
// Returns:
//   int32  
func (instance *PivotItem) GetPositionInSameParentNode()  (int32,  error)  {
	CGoReturnPtr := C.PivotItem_GetPositionInSameParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifying the position index in the PivotItems under the same parent node.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotItem) SetPositionInSameParentNode(value int32)  error {
	CGoReturnPtr := C.PivotItem_SetPositionInSameParentNode( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Moves the item up or down
// Parameters:
//   count - int32 
//   isSameParent - bool 
// Returns:
//   void  
func (instance *PivotItem) Move(count int32, issameparent bool)  error {
	CGoReturnPtr := C.PivotItem_Move( instance.ptr, C.int(count), C.bool(issameparent))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets whether the detail of this pivot item is hidden.
// Returns:
//   bool  
func (instance *PivotItem) IsDetailHidden()  (bool,  error)  {
	CGoReturnPtr := C.PivotItem_IsDetailHidden( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets whether the detail of this pivot item is hidden.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotItem) SetIsDetailHidden(value bool)  error {
	CGoReturnPtr := C.PivotItem_SetIsDetailHidden( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this pivot item is a calculated formula item.
// Returns:
//   bool  
func (instance *PivotItem) IsCalculatedItem()  (bool,  error)  {
	CGoReturnPtr := C.PivotItem_IsCalculatedItem( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the formula of this calculated item.
// Only works when this item is calculated item.
// Returns:
//   string  
func (instance *PivotItem) GetFormula()  (string,  error)  {
	CGoReturnPtr := C.PivotItem_GetFormula( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the item is removed from the data source.
// Returns:
//   bool  
func (instance *PivotItem) IsMissing()  (bool,  error)  {
	CGoReturnPtr := C.PivotItem_IsMissing( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the string value of the pivot item
// If the value is null, it will return ""
// Returns:
//   string  
func (instance *PivotItem) GetStringValue()  (string,  error)  {
	CGoReturnPtr := C.PivotItem_GetStringValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the double value of the pivot item
// If the value is null or not number ,it will return 0
// Returns:
//   float64  
func (instance *PivotItem) GetDoubleValue()  (float64,  error)  {
	CGoReturnPtr := C.PivotItem_GetDoubleValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the date time value of the pivot item
// If the value is null ,it will return DateTime.MinValue
// Returns:
//   Date  
func (instance *PivotItem) GetDateTimeValue()  (*Date,  error)  {
	CGoReturnPtr := C.PivotItem_GetDateTimeValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets the value of the pivot item
// Returns:
//   Object  
func (instance *PivotItem) GetValue()  (*Object,  error)  {
	CGoReturnPtr := C.PivotItem_GetValue( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Object{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteObject) 

	return result, nil 
}
// Gets the name of the pivot item.
// Returns:
//   string  
func (instance *PivotItem) GetName()  (string,  error)  {
	CGoReturnPtr := C.PivotItem_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the pivot item.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotItem) SetName(value string)  error {
	CGoReturnPtr := C.PivotItem_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the index of the pivot item in cache field.
// Returns:
//   int32  
func (instance *PivotItem) GetIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotItem_GetIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the index of the pivot item in cache field.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotItem) SetIndex(value int32)  error {
	CGoReturnPtr := C.PivotItem_SetIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotItem(pivotitem *PivotItem){
	runtime.SetFinalizer(pivotitem, nil)
	C.Delete_PivotItem(pivotitem.ptr)
	pivotitem.ptr = nil
}

// Class PivotItemCollection 

// Represents all the <see cref="PivotItem"/> objects in the PivotField.
type PivotItemCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotItemCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotItemCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the PivotItem Object at the specific index.
// Parameters:
//   index - int32 
// Returns:
//   PivotItem  
func (instance *PivotItemCollection) Get_Int(index int32)  (*PivotItem,  error)  {
	CGoReturnPtr := C.PivotItemCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotItem{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotItem) 

	return result, nil 
}
// Gets the <see cref="PivotItem"/> by the specific name.
// Parameters:
//   itemValue - string 
// Returns:
//   PivotItem  
func (instance *PivotItemCollection) Get_String(itemvalue string)  (*PivotItem,  error)  {
	CGoReturnPtr := C.PivotItemCollection_Get_String( instance.ptr, C.CString(itemvalue))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotItem{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotItem) 

	return result, nil 
}
// Gets the count of the pivot items.
// Returns:
//   int32  
func (instance *PivotItemCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotItemCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Directly swap two items.
// Parameters:
//   index1 - int32 
//   index2 - int32 
// Returns:
//   void  
func (instance *PivotItemCollection) SwapItem(index1 int32, index2 int32)  error {
	CGoReturnPtr := C.PivotItemCollection_SwapItem( instance.ptr, C.int(index1), C.int(index2))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotItemCollection(pivotitemcollection *PivotItemCollection){
	runtime.SetFinalizer(pivotitemcollection, nil)
	C.Delete_PivotItemCollection(pivotitemcollection.ptr)
	pivotitemcollection.ptr = nil
}

// Class PivotNumbericRangeGroupSettings 

// Represents the numberic range group of the pivot field.
type PivotNumbericRangeGroupSettings struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PivotFieldGroupSettings 
func NewPivotNumbericRangeGroupSettings(src *PivotFieldGroupSettings) ( *PivotNumbericRangeGroupSettings, error) {
	pivotnumbericrangegroupsettings := &PivotNumbericRangeGroupSettings{}
	CGoReturnPtr := C.New_PivotNumbericRangeGroupSettings(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		pivotnumbericrangegroupsettings.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotnumbericrangegroupsettings, DeletePivotNumbericRangeGroupSettings)
		return pivotnumbericrangegroupsettings, nil
	} else {
		pivotnumbericrangegroupsettings.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotnumbericrangegroupsettings, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotNumbericRangeGroupSettings) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotNumbericRangeGroupSettings_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the group type.
// Returns:
//   int32  
func (instance *PivotNumbericRangeGroupSettings) GetType()  (PivotFieldGroupType,  error)  {
	CGoReturnPtr := C.PivotNumbericRangeGroupSettings_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldGroupType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the start number of the group.
// Returns:
//   float64  
func (instance *PivotNumbericRangeGroupSettings) GetStart()  (float64,  error)  {
	CGoReturnPtr := C.PivotNumbericRangeGroupSettings_GetStart( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the end number of the group.
// Returns:
//   float64  
func (instance *PivotNumbericRangeGroupSettings) GetEnd()  (float64,  error)  {
	CGoReturnPtr := C.PivotNumbericRangeGroupSettings_GetEnd( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the interval of the group.
// Returns:
//   float64  
func (instance *PivotNumbericRangeGroupSettings) GetInterval()  (float64,  error)  {
	CGoReturnPtr := C.PivotNumbericRangeGroupSettings_GetInterval( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotNumbericRangeGroupSettings(pivotnumbericrangegroupsettings *PivotNumbericRangeGroupSettings){
	runtime.SetFinalizer(pivotnumbericrangegroupsettings, nil)
	C.Delete_PivotNumbericRangeGroupSettings(pivotnumbericrangegroupsettings.ptr)
	pivotnumbericrangegroupsettings.ptr = nil
}

// Class PivotPageFields 

// Represents the pivot page field items
// if the pivot table data source is consolidation ranges.
// It only can contain up to 4 fields.
type PivotPageFields struct {
	ptr unsafe.Pointer
}

// Represents the pivot page field items.
func NewPivotPageFields() ( *PivotPageFields, error) {
	pivotpagefields := &PivotPageFields{}
	CGoReturnPtr := C.New_PivotPageFields()
	if CGoReturnPtr.error_no == 0 {
		pivotpagefields.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivotpagefields, DeletePivotPageFields)
		return pivotpagefields, nil
	} else {
		pivotpagefields.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivotpagefields, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotPageFields) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotPageFields_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the number of page fields.
// Returns:
//   int32  
func (instance *PivotPageFields) GetPageFieldCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotPageFields_GetPageFieldCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotPageFields(pivotpagefields *PivotPageFields){
	runtime.SetFinalizer(pivotpagefields, nil)
	C.Delete_PivotPageFields(pivotpagefields.ptr)
	pivotpagefields.ptr = nil
}

// Class PivotShowValuesSetting 

// Represents the settings about showing values as when the ShowDataAs calculation is in use.
type PivotShowValuesSetting struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotShowValuesSetting) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotShowValuesSetting_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents how to show values of a data field in the pivot report.
// Returns:
//   int32  
func (instance *PivotShowValuesSetting) GetCalculationType()  (PivotFieldDataDisplayFormat,  error)  {
	CGoReturnPtr := C.PivotShowValuesSetting_GetCalculationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotFieldDataDisplayFormat(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents how to show values of a data field in the pivot report.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotShowValuesSetting) SetCalculationType(value PivotFieldDataDisplayFormat)  error {
	CGoReturnPtr := C.PivotShowValuesSetting_SetCalculationType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the base field for a ShowDataAs calculation when the ShowDataAs calculation is in use.
// Returns:
//   int32  
func (instance *PivotShowValuesSetting) GetBaseFieldIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotShowValuesSetting_GetBaseFieldIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the base field for a ShowDataAs calculation when the ShowDataAs calculation is in use.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotShowValuesSetting) SetBaseFieldIndex(value int32)  error {
	CGoReturnPtr := C.PivotShowValuesSetting_SetBaseFieldIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents type of the base pivot item in the base field when the ShowDataAs calculation is in use.
// Valid only for data fields.
// Because PivotItemPosition.Custom is only for read,if you need to set PivotItemPosition.Custom,
// please set PivotField.BaseItemIndex attribute.
// Returns:
//   int32  
func (instance *PivotShowValuesSetting) GetBaseItemPositionType()  (PivotItemPositionType,  error)  {
	CGoReturnPtr := C.PivotShowValuesSetting_GetBaseItemPositionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotItemPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents type of the base pivot item in the base field when the ShowDataAs calculation is in use.
// Valid only for data fields.
// Because PivotItemPosition.Custom is only for read,if you need to set PivotItemPosition.Custom,
// please set PivotField.BaseItemIndex attribute.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotShowValuesSetting) SetBaseItemPositionType(value PivotItemPositionType)  error {
	CGoReturnPtr := C.PivotShowValuesSetting_SetBaseItemPositionType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the custom index of the pivot item in the base field when the ShowDataAs calculation is in use.
// Valid only for data fields.
// Returns:
//   int32  
func (instance *PivotShowValuesSetting) GetBaseItemIndex()  (int32,  error)  {
	CGoReturnPtr := C.PivotShowValuesSetting_GetBaseItemIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the custom index of the pivot item in the base field when the ShowDataAs calculation is in use.
// Valid only for data fields.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotShowValuesSetting) SetBaseItemIndex(value int32)  error {
	CGoReturnPtr := C.PivotShowValuesSetting_SetBaseItemIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotShowValuesSetting(pivotshowvaluessetting *PivotShowValuesSetting){
	runtime.SetFinalizer(pivotshowvaluessetting, nil)
	C.Delete_PivotShowValuesSetting(pivotshowvaluessetting.ptr)
	pivotshowvaluessetting.ptr = nil
}

// Class PivotTable 

// Summary description for PivotTable.
type PivotTable struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotTable) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Performs application-defined tasks associated with freeing, releasing, or
// resetting unmanaged resources.
// Returns:
//   void  
func (instance *PivotTable) Dispose()  error {
	CGoReturnPtr := C.PivotTable_Dispose( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether the PivotTable is compatible for Excel2003 when refreshing PivotTable,
// if true, a string must be less than or equal to 255 characters, so if the string is greater than 255 characters,
// it will be truncated. if false, a string will not have the aforementioned restriction.
// The default value is true.
// Returns:
//   bool  
func (instance *PivotTable) IsExcel2003Compatible()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_IsExcel2003Compatible( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies whether the PivotTable is compatible for Excel2003 when refreshing PivotTable,
// if true, a string must be less than or equal to 255 characters, so if the string is greater than 255 characters,
// it will be truncated. if false, a string will not have the aforementioned restriction.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetIsExcel2003Compatible(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetIsExcel2003Compatible( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the name of the last user who refreshed this PivotTable
// Returns:
//   string  
func (instance *PivotTable) GetRefreshedByWho()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetRefreshedByWho( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the last date time when the PivotTable was refreshed.
// Returns:
//   Date  
func (instance *PivotTable) GetRefreshDate()  (*Date,  error)  {
	CGoReturnPtr := C.PivotTable_GetRefreshDate( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Date{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets and sets the pivottable style name.
// Returns:
//   string  
func (instance *PivotTable) GetPivotTableStyleName()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetPivotTableStyleName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the pivottable style name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetPivotTableStyleName(value string)  error {
	CGoReturnPtr := C.PivotTable_SetPivotTableStyleName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the built-in pivot table style.
// Returns:
//   int32  
func (instance *PivotTable) GetPivotTableStyleType()  (PivotTableStyleType,  error)  {
	CGoReturnPtr := C.PivotTable_GetPivotTableStyleType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotTableStyleType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the built-in pivot table style.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTable) SetPivotTableStyleType(value PivotTableStyleType)  error {
	CGoReturnPtr := C.PivotTable_SetPivotTableStyleType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Copies named style from another pivot table.
// Parameters:
//   pivotTable - PivotTable 
// Returns:
//   void  
func (instance *PivotTable) CopyStyle(pivottable *PivotTable)  error {
	CGoReturnPtr := C.PivotTable_CopyStyle( instance.ptr, pivottable.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Show all the report filter pages according to PivotField, the PivotField must be located in the PageFields.
// Parameters:
//   pageField - PivotField 
// Returns:
//   void  
func (instance *PivotTable) ShowReportFilterPage(pagefield *PivotField)  error {
	CGoReturnPtr := C.PivotTable_ShowReportFilterPage( instance.ptr, pagefield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Show all the report filter pages according to PivotField's name, the PivotField must be located in the PageFields.
// Parameters:
//   fieldName - string 
// Returns:
//   void  
func (instance *PivotTable) ShowReportFilterPageByName(fieldname string)  error {
	CGoReturnPtr := C.PivotTable_ShowReportFilterPageByName( instance.ptr, C.CString(fieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Show all the report filter pages according to the position index in the PageFields
// Parameters:
//   posIndex - int32 
// Returns:
//   void  
func (instance *PivotTable) ShowReportFilterPageByIndex(posindex int32)  error {
	CGoReturnPtr := C.PivotTable_ShowReportFilterPageByIndex( instance.ptr, C.int(posindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes a field from specific field area
// Parameters:
//   fieldType - int32 
//   fieldName - string 
// Returns:
//   void  
func (instance *PivotTable) RemoveField_PivotFieldType_String(fieldtype PivotFieldType, fieldname string)  error {
	CGoReturnPtr := C.PivotTable_RemoveField_PivotFieldType_String( instance.ptr, C.int( int32(fieldtype)), C.CString(fieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes a field from specific field area
// Parameters:
//   fieldType - int32 
//   baseFieldIndex - int32 
// Returns:
//   void  
func (instance *PivotTable) RemoveField_PivotFieldType_Int(fieldtype PivotFieldType, basefieldindex int32)  error {
	CGoReturnPtr := C.PivotTable_RemoveField_PivotFieldType_Integer( instance.ptr, C.int( int32(fieldtype)), C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Remove field from specific field area
// Parameters:
//   fieldType - int32 
//   pivotField - PivotField 
// Returns:
//   void  
func (instance *PivotTable) RemoveField_PivotFieldType_PivotField(fieldtype PivotFieldType, pivotfield *PivotField)  error {
	CGoReturnPtr := C.PivotTable_RemoveField_PivotFieldType_PivotField( instance.ptr, C.int( int32(fieldtype)), pivotfield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds the field to the specific area.
// Parameters:
//   fieldType - int32 
//   fieldName - string 
// Returns:
//   int32  
func (instance *PivotTable) AddFieldToArea_PivotFieldType_String(fieldtype PivotFieldType, fieldname string)  (int32,  error)  {
	CGoReturnPtr := C.PivotTable_AddFieldToArea_PivotFieldType_String( instance.ptr, C.int( int32(fieldtype)), C.CString(fieldname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds the field to the specific area.
// Parameters:
//   fieldType - int32 
//   baseFieldIndex - int32 
// Returns:
//   int32  
func (instance *PivotTable) AddFieldToArea_PivotFieldType_Int(fieldtype PivotFieldType, basefieldindex int32)  (int32,  error)  {
	CGoReturnPtr := C.PivotTable_AddFieldToArea_PivotFieldType_Integer( instance.ptr, C.int( int32(fieldtype)), C.int(basefieldindex))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds the field to the specific area.
// Parameters:
//   fieldType - int32 
//   pivotField - PivotField 
// Returns:
//   int32  
func (instance *PivotTable) AddFieldToArea_PivotFieldType_PivotField(fieldtype PivotFieldType, pivotfield *PivotField)  (int32,  error)  {
	CGoReturnPtr := C.PivotTable_AddFieldToArea_PivotFieldType_PivotField( instance.ptr, C.int( int32(fieldtype)), pivotfield.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a calculated field to pivot field.
// Parameters:
//   name - string 
//   formula - string 
//   dragToDataArea - bool 
// Returns:
//   void  
func (instance *PivotTable) AddCalculatedField_String_String_Bool(name string, formula string, dragtodataarea bool)  error {
	CGoReturnPtr := C.PivotTable_AddCalculatedField_String_String_Boolean( instance.ptr, C.CString(name), C.CString(formula), C.bool(dragtodataarea))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds a calculated field to pivot field and drag it to data area.
// Parameters:
//   name - string 
//   formula - string 
// Returns:
//   void  
func (instance *PivotTable) AddCalculatedField_String_String(name string, formula string)  error {
	CGoReturnPtr := C.PivotTable_AddCalculatedField_String_String( instance.ptr, C.CString(name), C.CString(formula))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the specific pivot field list by the region.
// Parameters:
//   fieldType - int32 
// Returns:
//   PivotFieldCollection  
func (instance *PivotTable) GetFields(fieldtype PivotFieldType)  (*PivotFieldCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetFields( instance.ptr, C.int( int32(fieldtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldCollection) 

	return result, nil 
}
// Returns a PivotFields object that are currently shown as column fields.
// Returns:
//   PivotFieldCollection  
func (instance *PivotTable) GetColumnFields()  (*PivotFieldCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetColumnFields( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldCollection) 

	return result, nil 
}
// Returns a PivotFields object that are currently shown as row fields.
// Returns:
//   PivotFieldCollection  
func (instance *PivotTable) GetRowFields()  (*PivotFieldCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetRowFields( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldCollection) 

	return result, nil 
}
// Returns a PivotFields object that are currently shown as page fields.
// Returns:
//   PivotFieldCollection  
func (instance *PivotTable) GetPageFields()  (*PivotFieldCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetPageFields( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldCollection) 

	return result, nil 
}
// Gets a PivotField object that represents all the data fields in a PivotTable.
// Read-only.It would be init only when there are two or more data fields in the DataPiovtFiels.
// It only use to add DataPivotField to the PivotTable row/column area . Default is in row area.
// Returns:
//   PivotFieldCollection  
func (instance *PivotTable) GetDataFields()  (*PivotFieldCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetDataFields( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldCollection) 

	return result, nil 
}
// Gets a <see cref="PivotField"/> object that represents all the data fields in a PivotTable.
// Read-only.
// It would only be created when there are two or more data fields in the Data region.
// Defaultly it is in row region. You can drag it to the row/column region with PivotTable.AddFieldToArea() method .
// Returns:
//   PivotField  
func (instance *PivotTable) GetDataField()  (*PivotField,  error)  {
	CGoReturnPtr := C.PivotTable_GetDataField( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotField{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotField) 

	return result, nil 
}
// Returns all base pivot fields in the PivotTable.
// Returns:
//   PivotFieldCollection  
func (instance *PivotTable) GetBaseFields()  (*PivotFieldCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetBaseFields( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFieldCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFieldCollection) 

	return result, nil 
}
// Returns a list of pivot filters.
// Returns:
//   PivotFilterCollection  
func (instance *PivotTable) GetPivotFilters()  (*PivotFilterCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetPivotFilters( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFilterCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFilterCollection) 

	return result, nil 
}
// Returns a CellArea object that represents the range
// that contains the column area in the PivotTable report. Read-only.
// Returns:
//   CellArea  
func (instance *PivotTable) GetColumnRange()  (*CellArea,  error)  {
	CGoReturnPtr := C.PivotTable_GetColumnRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Returns a CellArea object that represents the range
// that contains the row area in the PivotTable report. Read-only.
// Returns:
//   CellArea  
func (instance *PivotTable) GetRowRange()  (*CellArea,  error)  {
	CGoReturnPtr := C.PivotTable_GetRowRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Returns a <see cref="CellArea"/> object that represents the range that contains the data area
// in the list between the header row and the insert row. Read-only.
// Returns:
//   CellArea  
func (instance *PivotTable) GetDataBodyRange()  (*CellArea,  error)  {
	CGoReturnPtr := C.PivotTable_GetDataBodyRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Returns a CellArea object that represents the range containing the entire PivotTable report,
// but doesn't include page fields. Read-only.
// Returns:
//   CellArea  
func (instance *PivotTable) GetTableRange1()  (*CellArea,  error)  {
	CGoReturnPtr := C.PivotTable_GetTableRange1( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Returns a CellArea object that represents the range containing the entire PivotTable report,
// includes page fields. Read-only.
// Returns:
//   CellArea  
func (instance *PivotTable) GetTableRange2()  (*CellArea,  error)  {
	CGoReturnPtr := C.PivotTable_GetTableRange2( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellArea) 

	return result, nil 
}
// Moves the PivotTable to a different location in the worksheet.
// Parameters:
//   row - int32 
//   column - int32 
// Returns:
//   void  
func (instance *PivotTable) Move_Int_Int(row int32, column int32)  error {
	CGoReturnPtr := C.PivotTable_Move_Integer_Integer( instance.ptr, C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Moves the PivotTable to a different location in the worksheet.
// Parameters:
//   destCellName - string 
// Returns:
//   void  
func (instance *PivotTable) Move_String(destcellname string)  error {
	CGoReturnPtr := C.PivotTable_Move_String( instance.ptr, C.CString(destcellname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report shows grand totals for columns.
// Returns:
//   bool  
func (instance *PivotTable) GetColumnGrand()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetColumnGrand( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report shows grand totals for columns.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetColumnGrand(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetColumnGrand( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report displays classic pivottable layout.
// (enables dragging fields in the grid)
// Returns:
//   bool  
func (instance *PivotTable) IsGridDropZones()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_IsGridDropZones( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report displays classic pivottable layout.
// (enables dragging fields in the grid)
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetIsGridDropZones(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetIsGridDropZones( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report shows grand totals for rows.
// Returns:
//   bool  
func (instance *PivotTable) GetRowGrand()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetRowGrand( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report shows grand totals for rows.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetRowGrand(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetRowGrand( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report displays a custom string if the value is null.
// Returns:
//   bool  
func (instance *PivotTable) GetDisplayNullString()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetDisplayNullString( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report displays a custom string if the value is null.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetDisplayNullString(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetDisplayNullString( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the string displayed in cells that contain null values
// when the DisplayNullString property is true.The default value is an empty string.
// Returns:
//   string  
func (instance *PivotTable) GetNullString()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetNullString( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the string displayed in cells that contain null values
// when the DisplayNullString property is true.The default value is an empty string.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetNullString(value string)  error {
	CGoReturnPtr := C.PivotTable_SetNullString( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report displays a custom string in cells that contain errors.
// Returns:
//   bool  
func (instance *PivotTable) GetDisplayErrorString()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetDisplayErrorString( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report displays a custom string in cells that contain errors.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetDisplayErrorString(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetDisplayErrorString( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the name of the value area field header in the PivotTable.
// Returns:
//   string  
func (instance *PivotTable) GetDataFieldHeaderName()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetDataFieldHeaderName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the value area field header in the PivotTable.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetDataFieldHeaderName(value string)  error {
	CGoReturnPtr := C.PivotTable_SetDataFieldHeaderName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the string displayed in cells that contain errors
// when the DisplayErrorString property is true.The default value is an empty string.
// Returns:
//   string  
func (instance *PivotTable) GetErrorString()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetErrorString( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the string displayed in cells that contain errors
// when the DisplayErrorString property is true.The default value is an empty string.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetErrorString(value string)  error {
	CGoReturnPtr := C.PivotTable_SetErrorString( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report is automatically formatted.
// Checkbox "autoformat table " which is in pivottable option for Excel 2003
// Returns:
//   bool  
func (instance *PivotTable) IsAutoFormat()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_IsAutoFormat( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report is automatically formatted.
// Checkbox "autoformat table " which is in pivottable option for Excel 2003
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetIsAutoFormat(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetIsAutoFormat( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether autofitting column width on update
// Returns:
//   bool  
func (instance *PivotTable) GetAutofitColumnWidthOnUpdate()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetAutofitColumnWidthOnUpdate( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether autofitting column width on update
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetAutofitColumnWidthOnUpdate(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetAutofitColumnWidthOnUpdate( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the auto format type of PivotTable.
// Returns:
//   int32  
func (instance *PivotTable) GetAutoFormatType()  (PivotTableAutoFormatType,  error)  {
	CGoReturnPtr := C.PivotTable_GetAutoFormatType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotTableAutoFormatType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the auto format type of PivotTable.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTable) SetAutoFormatType(value PivotTableAutoFormatType)  error {
	CGoReturnPtr := C.PivotTable_SetAutoFormatType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to add blank rows.
// This property only applies for the PivotTable auto format types which needs to add blank rows.
// Returns:
//   bool  
func (instance *PivotTable) GetHasBlankRows()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetHasBlankRows( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether to add blank rows.
// This property only applies for the PivotTable auto format types which needs to add blank rows.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetHasBlankRows(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetHasBlankRows( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// True if the specified PivotTable report's outer-row item, column item, subtotal, and grand total labels use merged cells.
// Returns:
//   bool  
func (instance *PivotTable) GetMergeLabels()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetMergeLabels( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// True if the specified PivotTable report's outer-row item, column item, subtotal, and grand total labels use merged cells.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetMergeLabels(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetMergeLabels( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether formatting is preserved when the PivotTable is refreshed or recalculated.
// Returns:
//   bool  
func (instance *PivotTable) GetPreserveFormatting()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetPreserveFormatting( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether formatting is preserved when the PivotTable is refreshed or recalculated.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetPreserveFormatting(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetPreserveFormatting( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets whether showing expand/collapse buttons.
// Returns:
//   bool  
func (instance *PivotTable) GetShowDrill()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowDrill( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets and sets whether showing expand/collapse buttons.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowDrill(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowDrill( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets whether drilldown is enabled.
// Returns:
//   bool  
func (instance *PivotTable) GetEnableDrilldown()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetEnableDrilldown( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets whether drilldown is enabled.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetEnableDrilldown(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetEnableDrilldown( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable Field dialog box is available
// when the user double-clicks the PivotTable field.
// Returns:
//   bool  
func (instance *PivotTable) GetEnableFieldDialog()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetEnableFieldDialog( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable Field dialog box is available
// when the user double-clicks the PivotTable field.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetEnableFieldDialog(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetEnableFieldDialog( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets whether enable the field list for the PivotTable.
// Returns:
//   bool  
func (instance *PivotTable) GetEnableFieldList()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetEnableFieldList( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets whether enable the field list for the PivotTable.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetEnableFieldList(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetEnableFieldList( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable Wizard is available.
// Returns:
//   bool  
func (instance *PivotTable) GetEnableWizard()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetEnableWizard( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable Wizard is available.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetEnableWizard(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetEnableWizard( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether hidden page field items in the PivotTable report
// are included in row and column subtotals, block totals, and grand totals.
// The default value is False.
// Returns:
//   bool  
func (instance *PivotTable) GetSubtotalHiddenPageItems()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetSubtotalHiddenPageItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether hidden page field items in the PivotTable report
// are included in row and column subtotals, block totals, and grand totals.
// The default value is False.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetSubtotalHiddenPageItems(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetSubtotalHiddenPageItems( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the text string label that is displayed in the grand total column or row heading.
// The default value is the string "Grand Total".
// Returns:
//   string  
func (instance *PivotTable) GetGrandTotalName()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetGrandTotalName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the text string label that is displayed in the grand total column or row heading.
// The default value is the string "Grand Total".
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetGrandTotalName(value string)  error {
	CGoReturnPtr := C.PivotTable_SetGrandTotalName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the PivotTable report is recalculated only at the user's request.
// Returns:
//   bool  
func (instance *PivotTable) GetManualUpdate()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetManualUpdate( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the PivotTable report is recalculated only at the user's request.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetManualUpdate(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetManualUpdate( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether the fields of a PivotTable can have multiple filters set on them.
// Returns:
//   bool  
func (instance *PivotTable) IsMultipleFieldFilters()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_IsMultipleFieldFilters( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether the fields of a PivotTable can have multiple filters set on them.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetIsMultipleFieldFilters(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetIsMultipleFieldFilters( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether the fields of a PivotTable can have multiple filters set on them.
// Returns:
//   int32  
func (instance *PivotTable) GetMissingItemsLimit()  (PivotMissingItemLimitType,  error)  {
	CGoReturnPtr := C.PivotTable_GetMissingItemsLimit( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotMissingItemLimitType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies a boolean value that indicates whether the fields of a PivotTable can have multiple filters set on them.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTable) SetMissingItemsLimit(value PivotMissingItemLimitType)  error {
	CGoReturnPtr := C.PivotTable_SetMissingItemsLimit( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether the user is allowed to edit the cells in the data area of the pivottable.
// Enable cell editing in the values area
// Returns:
//   bool  
func (instance *PivotTable) GetEnableDataValueEditing()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetEnableDataValueEditing( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether the user is allowed to edit the cells in the data area of the pivottable.
// Enable cell editing in the values area
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetEnableDataValueEditing(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetEnableDataValueEditing( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether tooltips should be displayed for PivotTable data cells.
// Returns:
//   bool  
func (instance *PivotTable) GetShowDataTips()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowDataTips( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether tooltips should be displayed for PivotTable data cells.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowDataTips(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowDataTips( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether member property information should be omitted from PivotTable tooltips.
// Returns:
//   bool  
func (instance *PivotTable) GetShowMemberPropertyTips()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowMemberPropertyTips( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether member property information should be omitted from PivotTable tooltips.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowMemberPropertyTips(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowMemberPropertyTips( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether showing values row.
// Returns:
//   bool  
func (instance *PivotTable) GetShowValuesRow()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowValuesRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether showing values row.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowValuesRow(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowValuesRow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether to include empty columns in the table
// Returns:
//   bool  
func (instance *PivotTable) GetShowEmptyCol()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowEmptyCol( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether to include empty columns in the table
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowEmptyCol(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowEmptyCol( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether to include empty rows in the table.
// Returns:
//   bool  
func (instance *PivotTable) GetShowEmptyRow()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowEmptyRow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether to include empty rows in the table.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowEmptyRow(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowEmptyRow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether fields in the PivotTable are sorted in non-default order in the field list.
// Returns:
//   bool  
func (instance *PivotTable) GetFieldListSortAscending()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetFieldListSortAscending( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether fields in the PivotTable are sorted in non-default order in the field list.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetFieldListSortAscending(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetFieldListSortAscending( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies a boolean value that indicates whether drill indicators should be printed.
// print expand/collapse buttons when displayed on pivottable.
// Returns:
//   bool  
func (instance *PivotTable) GetPrintDrill()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetPrintDrill( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a boolean value that indicates whether drill indicators should be printed.
// print expand/collapse buttons when displayed on pivottable.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetPrintDrill(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetPrintDrill( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the title of the alter text.
// Returns:
//   string  
func (instance *PivotTable) GetAltTextTitle()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetAltTextTitle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the title of the alter text.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetAltTextTitle(value string)  error {
	CGoReturnPtr := C.PivotTable_SetAltTextTitle( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the description of the alt text.
// Returns:
//   string  
func (instance *PivotTable) GetAltTextDescription()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetAltTextDescription( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the description of the alt text.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetAltTextDescription(value string)  error {
	CGoReturnPtr := C.PivotTable_SetAltTextDescription( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the name of the PivotTable
// Returns:
//   string  
func (instance *PivotTable) GetName()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the name of the PivotTable
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetName(value string)  error {
	CGoReturnPtr := C.PivotTable_SetName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the Column Header Caption of the PivotTable.
// Returns:
//   string  
func (instance *PivotTable) GetColumnHeaderCaption()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetColumnHeaderCaption( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the Column Header Caption of the PivotTable.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetColumnHeaderCaption(value string)  error {
	CGoReturnPtr := C.PivotTable_SetColumnHeaderCaption( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the indentation increment for compact axis and can be used to set the Report Layout to Compact Form.
// Returns:
//   int32  
func (instance *PivotTable) GetIndent()  (int32,  error)  {
	CGoReturnPtr := C.PivotTable_GetIndent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the indentation increment for compact axis and can be used to set the Report Layout to Compact Form.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTable) SetIndent(value int32)  error {
	CGoReturnPtr := C.PivotTable_SetIndent( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the Row Header Caption of the PivotTable.
// Returns:
//   string  
func (instance *PivotTable) GetRowHeaderCaption()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetRowHeaderCaption( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the Row Header Caption of the PivotTable.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetRowHeaderCaption(value string)  error {
	CGoReturnPtr := C.PivotTable_SetRowHeaderCaption( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether row header caption is shown in the PivotTable report
// Indicates whether Display field captions and filter drop downs
// Returns:
//   bool  
func (instance *PivotTable) GetShowRowHeaderCaption()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowRowHeaderCaption( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether row header caption is shown in the PivotTable report
// Indicates whether Display field captions and filter drop downs
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowRowHeaderCaption(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowRowHeaderCaption( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether consider built-in custom list when sort data
// Returns:
//   bool  
func (instance *PivotTable) GetCustomListSort()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetCustomListSort( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether consider built-in custom list when sort data
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetCustomListSort(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetCustomListSort( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the Format Conditions of the pivot table.
// Returns:
//   PivotFormatConditionCollection  
func (instance *PivotTable) GetPivotFormatConditions()  (*PivotFormatConditionCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetPivotFormatConditions( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotFormatConditionCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotFormatConditionCollection) 

	return result, nil 
}
// Gets and sets the order in which page fields are added to the PivotTable report's layout.
// Returns:
//   int32  
func (instance *PivotTable) GetPageFieldOrder()  (PrintOrderType,  error)  {
	CGoReturnPtr := C.PivotTable_GetPageFieldOrder( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPrintOrderType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the order in which page fields are added to the PivotTable report's layout.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTable) SetPageFieldOrder(value PrintOrderType)  error {
	CGoReturnPtr := C.PivotTable_SetPageFieldOrder( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the number of page fields in each column or row in the PivotTable report.
// Returns:
//   int32  
func (instance *PivotTable) GetPageFieldWrapCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotTable_GetPageFieldWrapCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the number of page fields in each column or row in the PivotTable report.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTable) SetPageFieldWrapCount(value int32)  error {
	CGoReturnPtr := C.PivotTable_SetPageFieldWrapCount( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets a string saved with the PivotTable report.
// Returns:
//   string  
func (instance *PivotTable) GetTag()  (string,  error)  {
	CGoReturnPtr := C.PivotTable_GetTag( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets a string saved with the PivotTable report.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PivotTable) SetTag(value string)  error {
	CGoReturnPtr := C.PivotTable_SetTag( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data for the PivotTable report is saved with the workbook.
// Returns:
//   bool  
func (instance *PivotTable) GetSaveData()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetSaveData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data for the PivotTable report is saved with the workbook.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetSaveData(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetSaveData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether Refresh Data when Opening File.
// Returns:
//   bool  
func (instance *PivotTable) GetRefreshDataOnOpeningFile()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetRefreshDataOnOpeningFile( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether Refresh Data when Opening File.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetRefreshDataOnOpeningFile(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetRefreshDataOnOpeningFile( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether Refreshing Data or not.
// Returns:
//   bool  
func (instance *PivotTable) GetRefreshDataFlag()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetRefreshDataFlag( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether Refreshing Data or not.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetRefreshDataFlag(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetRefreshDataFlag( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the data source type of the pivot table.
// Returns:
//   int32  
func (instance *PivotTable) GetSourceType()  (PivotTableSourceType,  error)  {
	CGoReturnPtr := C.PivotTable_GetSourceType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotTableSourceType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Refreshes pivottable's data and setting from it's data source.
// Returns:
//   int32  
func (instance *PivotTable) RefreshData()  (PivotRefreshState,  error)  {
	CGoReturnPtr := C.PivotTable_RefreshData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotRefreshState(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Refreshes pivottable's data and setting from it's data source with options.
// Parameters:
//   option - PivotTableRefreshOption 
// Returns:
//   int32  
func (instance *PivotTable) RefreshData_PivotTableRefreshOption(option *PivotTableRefreshOption)  (PivotRefreshState,  error)  {
	CGoReturnPtr := C.PivotTable_RefreshData_PivotTableRefreshOption( instance.ptr, option.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPivotRefreshState(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Calculates pivottable's data to cells.
// Returns:
//   void  
func (instance *PivotTable) CalculateData()  error {
	CGoReturnPtr := C.PivotTable_CalculateData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Calculating pivot tables with options
// Parameters:
//   option - PivotTableCalculateOption 
// Returns:
//   void  
func (instance *PivotTable) CalculateData_PivotTableCalculateOption(option *PivotTableCalculateOption)  error {
	CGoReturnPtr := C.PivotTable_CalculateData_PivotTableCalculateOption( instance.ptr, option.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Clear PivotTable's data and formatting
// Returns:
//   void  
func (instance *PivotTable) ClearData()  error {
	CGoReturnPtr := C.PivotTable_ClearData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Calculates pivottable's range.
// Returns:
//   void  
func (instance *PivotTable) CalculateRange()  error {
	CGoReturnPtr := C.PivotTable_CalculateRange( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Format all the cell in the pivottable area
// Parameters:
//   style - Style 
// Returns:
//   void  
func (instance *PivotTable) FormatAll(style *Style)  error {
	CGoReturnPtr := C.PivotTable_FormatAll( instance.ptr, style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Format the row data in the pivottable area
// Parameters:
//   row - int32 
//   style - Style 
// Returns:
//   void  
func (instance *PivotTable) FormatRow(row int32, style *Style)  error {
	CGoReturnPtr := C.PivotTable_FormatRow( instance.ptr, C.int(row), style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Formats selected area of the PivotTable.
// Parameters:
//   pivotArea - PivotArea 
//   style - Style 
// Returns:
//   void  
func (instance *PivotTable) Format_PivotArea_Style(pivotarea *PivotArea, style *Style)  error {
	CGoReturnPtr := C.PivotTable_Format_PivotArea_Style( instance.ptr, pivotarea.ptr, style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Show the detail of one item in the data region to a new Table.
// Parameters:
//   rowOffset - int32 
//   columnOffset - int32 
//   newSheet - bool 
//   destRow - int32 
//   destColumn - int32 
// Returns:
//   void  
func (instance *PivotTable) ShowDetail(rowoffset int32, columnoffset int32, newsheet bool, destrow int32, destcolumn int32)  error {
	CGoReturnPtr := C.PivotTable_ShowDetail( instance.ptr, C.int(rowoffset), C.int(columnoffset), C.bool(newsheet), C.int(destrow), C.int(destcolumn))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the collection of formats applied to PivotTable.
// Returns:
//   PivotTableFormatCollection  
func (instance *PivotTable) GetPivotFormats()  (*PivotTableFormatCollection,  error)  {
	CGoReturnPtr := C.PivotTable_GetPivotFormats( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotTableFormatCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotTableFormatCollection) 

	return result, nil 
}
// Format the cell in the pivottable area
// Parameters:
//   row - int32 
//   column - int32 
//   style - Style 
// Returns:
//   void  
func (instance *PivotTable) Format_Int_Int_Style(row int32, column int32, style *Style)  error {
	CGoReturnPtr := C.PivotTable_Format_Integer_Integer_Style( instance.ptr, C.int(row), C.int(column), style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether pivot item captions on the row area are repeated on each printed page for pivot fields in tabular form.
// Returns:
//   bool  
func (instance *PivotTable) GetRepeatItemsOnEachPrintedPage()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetRepeatItemsOnEachPrintedPage( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether pivot item captions on the row area are repeated on each printed page for pivot fields in tabular form.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetRepeatItemsOnEachPrintedPage(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetRepeatItemsOnEachPrintedPage( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the print titles for the worksheet are set based
// on the PivotTable report. The default value is false.
// Returns:
//   bool  
func (instance *PivotTable) GetPrintTitles()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetPrintTitles( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the print titles for the worksheet are set based
// on the PivotTable report. The default value is false.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetPrintTitles(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetPrintTitles( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether items in the row and column areas are visible
// when the data area of the PivotTable is empty. The default value is true.
// Returns:
//   bool  
func (instance *PivotTable) GetDisplayImmediateItems()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetDisplayImmediateItems( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether items in the row and column areas are visible
// when the data area of the PivotTable is empty. The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetDisplayImmediateItems(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetDisplayImmediateItems( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this PivotTable is selected.
// Returns:
//   bool  
func (instance *PivotTable) IsSelected()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_IsSelected( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether this PivotTable is selected.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetIsSelected(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetIsSelected( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the row header in the pivot table should have the style applied.
// Returns:
//   bool  
func (instance *PivotTable) GetShowPivotStyleRowHeader()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowPivotStyleRowHeader( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the row header in the pivot table should have the style applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowPivotStyleRowHeader(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowPivotStyleRowHeader( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the column header in the pivot table should have the style applied.
// Returns:
//   bool  
func (instance *PivotTable) GetShowPivotStyleColumnHeader()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowPivotStyleColumnHeader( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the column header in the pivot table should have the style applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowPivotStyleColumnHeader(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowPivotStyleColumnHeader( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether row stripe formatting is applied.
// Returns:
//   bool  
func (instance *PivotTable) GetShowPivotStyleRowStripes()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowPivotStyleRowStripes( instance.ptr)
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
func (instance *PivotTable) SetShowPivotStyleRowStripes(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowPivotStyleRowStripes( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether stripe formatting is applied for column.
// Returns:
//   bool  
func (instance *PivotTable) GetShowPivotStyleColumnStripes()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowPivotStyleColumnStripes( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether stripe formatting is applied for column.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowPivotStyleColumnStripes(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowPivotStyleColumnStripes( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the column formatting is applied.
// Returns:
//   bool  
func (instance *PivotTable) GetShowPivotStyleLastColumn()  (bool,  error)  {
	CGoReturnPtr := C.PivotTable_GetShowPivotStyleLastColumn( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the column formatting is applied.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTable) SetShowPivotStyleLastColumn(value bool)  error {
	CGoReturnPtr := C.PivotTable_SetShowPivotStyleLastColumn( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Layouts the PivotTable in compact form.
// Returns:
//   void  
func (instance *PivotTable) ShowInCompactForm()  error {
	CGoReturnPtr := C.PivotTable_ShowInCompactForm( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Layouts the PivotTable in outline form.
// Returns:
//   void  
func (instance *PivotTable) ShowInOutlineForm()  error {
	CGoReturnPtr := C.PivotTable_ShowInOutlineForm( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Layouts the PivotTable in tabular form.
// Returns:
//   void  
func (instance *PivotTable) ShowInTabularForm()  error {
	CGoReturnPtr := C.PivotTable_ShowInTabularForm( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Cell"/> object by the display name of PivotField.
// Parameters:
//   displayName - string 
// Returns:
//   Cell  
func (instance *PivotTable) GetCellByDisplayName(displayname string)  (*Cell,  error)  {
	CGoReturnPtr := C.PivotTable_GetCellByDisplayName( instance.ptr, C.CString(displayname))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Cell{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCell) 

	return result, nil 
}


func DeletePivotTable(pivottable *PivotTable){
	runtime.SetFinalizer(pivottable, nil)
	C.Delete_PivotTable(pivottable.ptr)
	pivottable.ptr = nil
}

// Class PivotTableCalculateOption 

// Rerepsents the options of calcuating the pivot table.
type PivotTableCalculateOption struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPivotTableCalculateOption() ( *PivotTableCalculateOption, error) {
	pivottablecalculateoption := &PivotTableCalculateOption{}
	CGoReturnPtr := C.New_PivotTableCalculateOption()
	if CGoReturnPtr.error_no == 0 {
		pivottablecalculateoption.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivottablecalculateoption, DeletePivotTableCalculateOption)
		return pivottablecalculateoption, nil
	} else {
		pivottablecalculateoption.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivottablecalculateoption, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotTableCalculateOption) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableCalculateOption_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether refreshing data source of the pivottable.
// Returns:
//   bool  
func (instance *PivotTableCalculateOption) GetRefreshData()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableCalculateOption_GetRefreshData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether refreshing data source of the pivottable.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTableCalculateOption) SetRefreshData(value bool)  error {
	CGoReturnPtr := C.PivotTableCalculateOption_SetRefreshData( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether refreshing charts are based on this pivot table.
// Returns:
//   bool  
func (instance *PivotTableCalculateOption) GetRefreshCharts()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableCalculateOption_GetRefreshCharts( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether refreshing charts are based on this pivot table.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PivotTableCalculateOption) SetRefreshCharts(value bool)  error {
	CGoReturnPtr := C.PivotTableCalculateOption_SetRefreshCharts( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents how to reserve missing pivot items.
// Returns:
//   int32  
func (instance *PivotTableCalculateOption) GetReserveMissingPivotItemType()  (ReserveMissingPivotItemType,  error)  {
	CGoReturnPtr := C.PivotTableCalculateOption_GetReserveMissingPivotItemType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToReserveMissingPivotItemType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents how to reserve missing pivot items.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTableCalculateOption) SetReserveMissingPivotItemType(value ReserveMissingPivotItemType)  error {
	CGoReturnPtr := C.PivotTableCalculateOption_SetReserveMissingPivotItemType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotTableCalculateOption(pivottablecalculateoption *PivotTableCalculateOption){
	runtime.SetFinalizer(pivottablecalculateoption, nil)
	C.Delete_PivotTableCalculateOption(pivottablecalculateoption.ptr)
	pivottablecalculateoption.ptr = nil
}

// Class PivotTableCollection 

// Represents the collection of all the PivotTable objects on the specified worksheet.
type PivotTableCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotTableCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Performs application-defined tasks associated with freeing, releasing, or
// resetting unmanaged resources.
// Returns:
//   void  
func (instance *PivotTableCollection) Dispose()  error {
	CGoReturnPtr := C.PivotTableCollection_Dispose( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Adds a new PivotTable.
// Parameters:
//   sourceData - string 
//   destCellName - string 
//   tableName - string 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_String_String_String(sourcedata string, destcellname string, tablename string)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_String_String_String( instance.ptr, C.CString(sourcedata), C.CString(destcellname), C.CString(tablename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable.
// Parameters:
//   sourceData - string 
//   destCellName - string 
//   tableName - string 
//   useSameSource - bool 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_String_String_String_Bool(sourcedata string, destcellname string, tablename string, usesamesource bool)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_String_String_String_Boolean( instance.ptr, C.CString(sourcedata), C.CString(destcellname), C.CString(tablename), C.bool(usesamesource))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable.
// Parameters:
//   sourceData - string 
//   row - int32 
//   column - int32 
//   tableName - string 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_String_Int_Int_String(sourcedata string, row int32, column int32, tablename string)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_String_Integer_Integer_String( instance.ptr, C.CString(sourcedata), C.int(row), C.int(column), C.CString(tablename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable.
// Parameters:
//   sourceData - string 
//   row - int32 
//   column - int32 
//   tableName - string 
//   useSameSource - bool 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_String_Int_Int_String_Bool(sourcedata string, row int32, column int32, tablename string, usesamesource bool)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_String_Integer_Integer_String_Boolean( instance.ptr, C.CString(sourcedata), C.int(row), C.int(column), C.CString(tablename), C.bool(usesamesource))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable.
// Parameters:
//   sourceData - string 
//   row - int32 
//   column - int32 
//   tableName - string 
//   useSameSource - bool 
//   isXlsClassic - bool 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_String_Int_Int_String_Bool_Bool(sourcedata string, row int32, column int32, tablename string, usesamesource bool, isxlsclassic bool)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_String_Integer_Integer_String_Boolean_Boolean( instance.ptr, C.CString(sourcedata), C.int(row), C.int(column), C.CString(tablename), C.bool(usesamesource), C.bool(isxlsclassic))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable.
// Parameters:
//   sourceData - string 
//   cell - string 
//   tableName - string 
//   useSameSource - bool 
//   isXlsClassic - bool 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_String_String_String_Bool_Bool(sourcedata string, cell string, tablename string, usesamesource bool, isxlsclassic bool)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_String_String_String_Boolean_Boolean( instance.ptr, C.CString(sourcedata), C.CString(cell), C.CString(tablename), C.bool(usesamesource), C.bool(isxlsclassic))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable based on another PivotTable.
// Parameters:
//   pivotTable - PivotTable 
//   destCellName - string 
//   tableName - string 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_PivotTable_String_String(pivottable *PivotTable, destcellname string, tablename string)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_PivotTable_String_String( instance.ptr, pivottable.ptr, C.CString(destcellname), C.CString(tablename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a new PivotTable based on another PivotTable.
// Parameters:
//   pivotTable - PivotTable 
//   row - int32 
//   column - int32 
//   tableName - string 
// Returns:
//   int32  
func (instance *PivotTableCollection) Add_PivotTable_Int_Int_String(pivottable *PivotTable, row int32, column int32, tablename string)  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Add_PivotTable_Integer_Integer_String( instance.ptr, pivottable.ptr, C.int(row), C.int(column), C.CString(tablename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the PivotTable report by index.
// Parameters:
//   index - int32 
// Returns:
//   PivotTable  
func (instance *PivotTableCollection) Get_Int(index int32)  (*PivotTable,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Get_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotTable) 

	return result, nil 
}
// Gets the PivotTable report by pivottable's name.
// Parameters:
//   name - string 
// Returns:
//   PivotTable  
func (instance *PivotTableCollection) Get_String(name string)  (*PivotTable,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Get_String( instance.ptr, C.CString(name))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotTable) 

	return result, nil 
}
// Gets the PivotTable report by pivottable's position.
// Parameters:
//   row - int32 
//   column - int32 
// Returns:
//   PivotTable  
func (instance *PivotTableCollection) Get_Int_Int(row int32, column int32)  (*PivotTable,  error)  {
	CGoReturnPtr := C.PivotTableCollection_Get_Integer_Integer( instance.ptr, C.int(row), C.int(column))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotTable{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotTable) 

	return result, nil 
}
// Clear all pivot tables.
// Returns:
//   void  
func (instance *PivotTableCollection) Clear()  error {
	CGoReturnPtr := C.PivotTableCollection_Clear( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Deletes the specified PivotTable and delete the PivotTable data
// Parameters:
//   pivotTable - PivotTable 
// Returns:
//   void  
func (instance *PivotTableCollection) Remove_PivotTable(pivottable *PivotTable)  error {
	CGoReturnPtr := C.PivotTableCollection_Remove_PivotTable( instance.ptr, pivottable.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Deletes the specified PivotTable
// Parameters:
//   pivotTable - PivotTable 
//   keepData - bool 
// Returns:
//   void  
func (instance *PivotTableCollection) Remove_PivotTable_Bool(pivottable *PivotTable, keepdata bool)  error {
	CGoReturnPtr := C.PivotTableCollection_Remove_PivotTable_Boolean( instance.ptr, pivottable.ptr, C.bool(keepdata))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Deletes the PivotTable at the specified index and delete the PivotTable data
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *PivotTableCollection) RemoveAt_Int(index int32)  error {
	CGoReturnPtr := C.PivotTableCollection_RemoveAt_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Deletes the PivotTable at the specified index
// Parameters:
//   index - int32 
//   keepData - bool 
// Returns:
//   void  
func (instance *PivotTableCollection) RemoveAt_Int_Bool(index int32, keepdata bool)  error {
	CGoReturnPtr := C.PivotTableCollection_RemoveAt_Integer_Boolean( instance.ptr, C.int(index), C.bool(keepdata))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   int32  
func (instance *PivotTableCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotTableCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotTableCollection(pivottablecollection *PivotTableCollection){
	runtime.SetFinalizer(pivottablecollection, nil)
	C.Delete_PivotTableCollection(pivottablecollection.ptr)
	pivottablecollection.ptr = nil
}

// Class PivotTableFormat 

// Represents the format defined in the PivotTable.
type PivotTableFormat struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotTableFormat) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableFormat_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the pivot area.
// Returns:
//   PivotArea  
func (instance *PivotTableFormat) GetPivotArea()  (*PivotArea,  error)  {
	CGoReturnPtr := C.PivotTableFormat_GetPivotArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotArea{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotArea) 

	return result, nil 
}
// Gets the formatted style.
// Returns:
//   Style  
func (instance *PivotTableFormat) GetStyle()  (*Style,  error)  {
	CGoReturnPtr := C.PivotTableFormat_GetStyle( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Style{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteStyle) 

	return result, nil 
}
// Sets the style of the pivot area.
// Parameters:
//   style - Style 
// Returns:
//   void  
func (instance *PivotTableFormat) SetStyle(style *Style)  error {
	CGoReturnPtr := C.PivotTableFormat_SetStyle( instance.ptr, style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotTableFormat(pivottableformat *PivotTableFormat){
	runtime.SetFinalizer(pivottableformat, nil)
	C.Delete_PivotTableFormat(pivottableformat.ptr)
	pivottableformat.ptr = nil
}

// Class PivotTableFormatCollection 

// Represents the collection of formats applied to PivotTable.
type PivotTableFormatCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotTableFormatCollection) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableFormatCollection_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the format by the index.
// Parameters:
//   index - int32 
// Returns:
//   PivotTableFormat  
func (instance *PivotTableFormatCollection) Get(index int32)  (*PivotTableFormat,  error)  {
	CGoReturnPtr := C.PivotTableFormatCollection_Get( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotTableFormat{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotTableFormat) 

	return result, nil 
}
// Add a <see cref="PivotTableFormat"/>.
// Returns:
//   int32  
func (instance *PivotTableFormatCollection) Add()  (int32,  error)  {
	CGoReturnPtr := C.PivotTableFormatCollection_Add( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Formats selected area.
// Parameters:
//   axisType - int32 
//   fieldPosition - int32 
//   subtotalType - int32 
//   selectionType - int32 
//   isGrandRow - bool 
//   isGrandColumn - bool 
//   style - Style 
// Returns:
//   PivotTableFormat  
func (instance *PivotTableFormatCollection) FormatArea(axistype PivotFieldType, fieldposition int32, subtotaltype PivotFieldSubtotalType, selectiontype PivotTableSelectionType, isgrandrow bool, isgrandcolumn bool, style *Style)  (*PivotTableFormat,  error)  {
	CGoReturnPtr := C.PivotTableFormatCollection_FormatArea( instance.ptr, C.int( int32(axistype)), C.int(fieldposition), C.int( int32(subtotaltype)), C.int( int32(selectiontype)), C.bool(isgrandrow), C.bool(isgrandcolumn), style.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &PivotTableFormat{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeletePivotTableFormat) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *PivotTableFormatCollection) GetCount()  (int32,  error)  {
	CGoReturnPtr := C.PivotTableFormatCollection_GetCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePivotTableFormatCollection(pivottableformatcollection *PivotTableFormatCollection){
	runtime.SetFinalizer(pivottableformatcollection, nil)
	C.Delete_PivotTableFormatCollection(pivottableformatcollection.ptr)
	pivottableformatcollection.ptr = nil
}

// Class PivotTableRefreshOption 

// Represents the options of refreshing data source of the pivot table.
type PivotTableRefreshOption struct {
	ptr unsafe.Pointer
}

// Represents the options of refreshing data source of the pivot table.
func NewPivotTableRefreshOption() ( *PivotTableRefreshOption, error) {
	pivottablerefreshoption := &PivotTableRefreshOption{}
	CGoReturnPtr := C.New_PivotTableRefreshOption()
	if CGoReturnPtr.error_no == 0 {
		pivottablerefreshoption.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pivottablerefreshoption, DeletePivotTableRefreshOption)
		return pivottablerefreshoption, nil
	} else {
		pivottablerefreshoption.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pivottablerefreshoption, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PivotTableRefreshOption) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PivotTableRefreshOption_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents how to reserve missing pivot items.
// Returns:
//   int32  
func (instance *PivotTableRefreshOption) GetReserveMissingPivotItemType()  (ReserveMissingPivotItemType,  error)  {
	CGoReturnPtr := C.PivotTableRefreshOption_GetReserveMissingPivotItemType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToReserveMissingPivotItemType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents how to reserve missing pivot items.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *PivotTableRefreshOption) SetReserveMissingPivotItemType(value ReserveMissingPivotItemType)  error {
	CGoReturnPtr := C.PivotTableRefreshOption_SetReserveMissingPivotItemType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePivotTableRefreshOption(pivottablerefreshoption *PivotTableRefreshOption){
	runtime.SetFinalizer(pivottablerefreshoption, nil)
	C.Delete_PivotTableRefreshOption(pivottablerefreshoption.ptr)
	pivottablerefreshoption.ptr = nil
}
