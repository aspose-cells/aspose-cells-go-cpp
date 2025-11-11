// +build windows

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/win_x86_64" -L"${SRCDIR}/lib/win_x86_64" -lAspose.Cells.CWrapper
// #include <CellsFunctionMap.h>
import "C"
import (
	"fmt"  
 	
	"errors"	
	"runtime"
	"unsafe" 
)

/**************Enum BulletType *****************/

// Represents the type of the bullet.
type BulletType int32

const(
// No bullet.
BulletType_None BulletType = 0 

// Character bullet.
BulletType_Character BulletType = 1 

// Image bullet.
BulletType_Picture BulletType = 2 

// Automatic numbered bullet.
BulletType_AutoNumbered BulletType = 3 
)

func Int32ToBulletType(value int32)(BulletType ,error){
	switch value {
		case 0:  return BulletType_None, nil  
		case 1:  return BulletType_Character, nil  
		case 2:  return BulletType_Picture, nil  
		case 3:  return BulletType_AutoNumbered, nil  
		default:
			return 0 ,fmt.Errorf("invalid BulletType value: %d", value)
	}
}

/**************Enum LineSpaceSizeType *****************/

// Represents the unit type of line space size.
type LineSpaceSizeType int32

const(
// Represents in unit of a percentage of the text size.
LineSpaceSizeType_Percentage LineSpaceSizeType = 0 

// Represents in unit of points.
LineSpaceSizeType_Points LineSpaceSizeType = 1 
)

func Int32ToLineSpaceSizeType(value int32)(LineSpaceSizeType ,error){
	switch value {
		case 0:  return LineSpaceSizeType_Percentage, nil  
		case 1:  return LineSpaceSizeType_Points, nil  
		default:
			return 0 ,fmt.Errorf("invalid LineSpaceSizeType value: %d", value)
	}
}

/**************Enum ShapeTextVerticalAlignmentType *****************/

// It corresponds to "Format Shape - Text Options - Text Box - Vertical Alignment" in Excel.
type ShapeTextVerticalAlignmentType int32

const(
ShapeTextVerticalAlignmentType_Top ShapeTextVerticalAlignmentType = 0 


ShapeTextVerticalAlignmentType_Middle ShapeTextVerticalAlignmentType = 1 


ShapeTextVerticalAlignmentType_Bottom ShapeTextVerticalAlignmentType = 2 


ShapeTextVerticalAlignmentType_TopCentered ShapeTextVerticalAlignmentType = 3 


ShapeTextVerticalAlignmentType_MiddleCentered ShapeTextVerticalAlignmentType = 4 


ShapeTextVerticalAlignmentType_BottomCentered ShapeTextVerticalAlignmentType = 5 


ShapeTextVerticalAlignmentType_Left ShapeTextVerticalAlignmentType = 6 


ShapeTextVerticalAlignmentType_Center ShapeTextVerticalAlignmentType = 7 


ShapeTextVerticalAlignmentType_Right ShapeTextVerticalAlignmentType = 8 


ShapeTextVerticalAlignmentType_LeftMiddle ShapeTextVerticalAlignmentType = 9 


ShapeTextVerticalAlignmentType_CenterMiddle ShapeTextVerticalAlignmentType = 10 


ShapeTextVerticalAlignmentType_RightMiddle ShapeTextVerticalAlignmentType = 11 
)

func Int32ToShapeTextVerticalAlignmentType(value int32)(ShapeTextVerticalAlignmentType ,error){
	switch value {
		case 0:  return ShapeTextVerticalAlignmentType_Top, nil  
		case 1:  return ShapeTextVerticalAlignmentType_Middle, nil  
		case 2:  return ShapeTextVerticalAlignmentType_Bottom, nil  
		case 3:  return ShapeTextVerticalAlignmentType_TopCentered, nil  
		case 4:  return ShapeTextVerticalAlignmentType_MiddleCentered, nil  
		case 5:  return ShapeTextVerticalAlignmentType_BottomCentered, nil  
		case 6:  return ShapeTextVerticalAlignmentType_Left, nil  
		case 7:  return ShapeTextVerticalAlignmentType_Center, nil  
		case 8:  return ShapeTextVerticalAlignmentType_Right, nil  
		case 9:  return ShapeTextVerticalAlignmentType_LeftMiddle, nil  
		case 10:  return ShapeTextVerticalAlignmentType_CenterMiddle, nil  
		case 11:  return ShapeTextVerticalAlignmentType_RightMiddle, nil  
		default:
			return 0 ,fmt.Errorf("invalid ShapeTextVerticalAlignmentType value: %d", value)
	}
}

/**************Enum TextAutonumberScheme *****************/

// Represents all automatic number scheme.
type TextAutonumberScheme int32

const(
TextAutonumberScheme_None TextAutonumberScheme = 0 

// (a), (b), (c), …
TextAutonumberScheme_AlphaLcParenBoth TextAutonumberScheme = 1 

// a), b), c), …
TextAutonumberScheme_AlphaLcParenR TextAutonumberScheme = 2 

// a., b., c., …
TextAutonumberScheme_AlphaLcPeriod TextAutonumberScheme = 3 

// (A), (B), (C), …
TextAutonumberScheme_AlphaUcParenBoth TextAutonumberScheme = 4 

// A), B), C), …
TextAutonumberScheme_AlphaUcParenR TextAutonumberScheme = 5 

// A., B., C., …
TextAutonumberScheme_AlphaUcPeriod TextAutonumberScheme = 6 

// Bidi Arabic 1 (AraAlpha) with ANSI minus symbol
TextAutonumberScheme_Arabic1Minus TextAutonumberScheme = 7 

// Bidi Arabic 2 (AraAbjad) with ANSI minus symbol
TextAutonumberScheme_Arabic2Minus TextAutonumberScheme = 8 

// Dbl-byte Arabic numbers w/ double-byte period
TextAutonumberScheme_ArabicDbPeriod TextAutonumberScheme = 9 

// Dbl-byte Arabic numbers
TextAutonumberScheme_ArabicDbPlain TextAutonumberScheme = 10 

// (1), (2), (3), …
TextAutonumberScheme_ArabicParenBoth TextAutonumberScheme = 11 

// 1), 2), 3), …
TextAutonumberScheme_ArabicParenR TextAutonumberScheme = 12 

// 1., 2., 3., …
TextAutonumberScheme_ArabicPeriod TextAutonumberScheme = 13 

// 1, 2, 3, …
TextAutonumberScheme_ArabicPlain TextAutonumberScheme = 14 

// Dbl-byte circle numbers (1-10 circle[0x2460-], 11-arabic numbers)
TextAutonumberScheme_CircleNumDbPlain TextAutonumberScheme = 15 

// Wingdings black circle numbers
TextAutonumberScheme_CircleNumWdBlackPlain TextAutonumberScheme = 16 

// Wingdings white circle numbers (0-10 circle[0x0080-],11- arabic numbers)
TextAutonumberScheme_CircleNumWdWhitePlain TextAutonumberScheme = 17 

// EA: Simplified Chinese w/ single-byte period
TextAutonumberScheme_Ea1ChsPeriod TextAutonumberScheme = 18 

// EA: Simplified Chinese (TypeA 1-99, TypeC 100-)
TextAutonumberScheme_Ea1ChsPlain TextAutonumberScheme = 19 

// EA: Traditional Chinese w/ single-byte period
TextAutonumberScheme_Ea1ChtPeriod TextAutonumberScheme = 20 

// EA: Traditional Chinese (TypeA 1-19, TypeC 20-)
TextAutonumberScheme_Ea1ChtPlain TextAutonumberScheme = 21 

// EA: Japanese w/ double-byte period
TextAutonumberScheme_Ea1JpnChsDbPeriod TextAutonumberScheme = 22 

// EA: Japanese/Korean w/ single-byte period
TextAutonumberScheme_Ea1JpnKorPeriod TextAutonumberScheme = 23 

// EA: Japanese/Korean (TypeC 1-)
TextAutonumberScheme_Ea1JpnKorPlain TextAutonumberScheme = 24 

// Bidi Hebrew 2 with ANSI minus symbol
TextAutonumberScheme_Hebrew2Minus TextAutonumberScheme = 25 

// Hindi alphabet period - consonants
TextAutonumberScheme_HindiAlpha1Period TextAutonumberScheme = 26 

// Hindi alphabet period - vowels
TextAutonumberScheme_HindiAlphaPeriod TextAutonumberScheme = 27 

// Hindi numerical parentheses - right
TextAutonumberScheme_HindiNumParenR TextAutonumberScheme = 28 

// Hindi numerical period
TextAutonumberScheme_HindiNumPeriod TextAutonumberScheme = 29 

// (i), (ii), (iii), …
TextAutonumberScheme_RomanLcParenBoth TextAutonumberScheme = 30 

// i), ii), iii), …
TextAutonumberScheme_RomanLcParenR TextAutonumberScheme = 31 

// i., ii., iii., …
TextAutonumberScheme_RomanLcPeriod TextAutonumberScheme = 32 

// (I), (II), (III), …
TextAutonumberScheme_RomanUcParenBoth TextAutonumberScheme = 33 

// I), II), III), …
TextAutonumberScheme_RomanUcParenR TextAutonumberScheme = 34 

// I., II., III., …
TextAutonumberScheme_RomanUcPeriod TextAutonumberScheme = 35 

// Thai alphabet parentheses - both
TextAutonumberScheme_ThaiAlphaParenBoth TextAutonumberScheme = 36 

// Thai alphabet parentheses - right
TextAutonumberScheme_ThaiAlphaParenR TextAutonumberScheme = 37 

// Thai alphabet period
TextAutonumberScheme_ThaiAlphaPeriod TextAutonumberScheme = 38 

// Thai numerical parentheses - both
TextAutonumberScheme_ThaiNumParenBoth TextAutonumberScheme = 39 

// Thai numerical parentheses - right
TextAutonumberScheme_ThaiNumParenR TextAutonumberScheme = 40 

// Thai numerical period
TextAutonumberScheme_ThaiNumPeriod TextAutonumberScheme = 41 
)

func Int32ToTextAutonumberScheme(value int32)(TextAutonumberScheme ,error){
	switch value {
		case 0:  return TextAutonumberScheme_None, nil  
		case 1:  return TextAutonumberScheme_AlphaLcParenBoth, nil  
		case 2:  return TextAutonumberScheme_AlphaLcParenR, nil  
		case 3:  return TextAutonumberScheme_AlphaLcPeriod, nil  
		case 4:  return TextAutonumberScheme_AlphaUcParenBoth, nil  
		case 5:  return TextAutonumberScheme_AlphaUcParenR, nil  
		case 6:  return TextAutonumberScheme_AlphaUcPeriod, nil  
		case 7:  return TextAutonumberScheme_Arabic1Minus, nil  
		case 8:  return TextAutonumberScheme_Arabic2Minus, nil  
		case 9:  return TextAutonumberScheme_ArabicDbPeriod, nil  
		case 10:  return TextAutonumberScheme_ArabicDbPlain, nil  
		case 11:  return TextAutonumberScheme_ArabicParenBoth, nil  
		case 12:  return TextAutonumberScheme_ArabicParenR, nil  
		case 13:  return TextAutonumberScheme_ArabicPeriod, nil  
		case 14:  return TextAutonumberScheme_ArabicPlain, nil  
		case 15:  return TextAutonumberScheme_CircleNumDbPlain, nil  
		case 16:  return TextAutonumberScheme_CircleNumWdBlackPlain, nil  
		case 17:  return TextAutonumberScheme_CircleNumWdWhitePlain, nil  
		case 18:  return TextAutonumberScheme_Ea1ChsPeriod, nil  
		case 19:  return TextAutonumberScheme_Ea1ChsPlain, nil  
		case 20:  return TextAutonumberScheme_Ea1ChtPeriod, nil  
		case 21:  return TextAutonumberScheme_Ea1ChtPlain, nil  
		case 22:  return TextAutonumberScheme_Ea1JpnChsDbPeriod, nil  
		case 23:  return TextAutonumberScheme_Ea1JpnKorPeriod, nil  
		case 24:  return TextAutonumberScheme_Ea1JpnKorPlain, nil  
		case 25:  return TextAutonumberScheme_Hebrew2Minus, nil  
		case 26:  return TextAutonumberScheme_HindiAlpha1Period, nil  
		case 27:  return TextAutonumberScheme_HindiAlphaPeriod, nil  
		case 28:  return TextAutonumberScheme_HindiNumParenR, nil  
		case 29:  return TextAutonumberScheme_HindiNumPeriod, nil  
		case 30:  return TextAutonumberScheme_RomanLcParenBoth, nil  
		case 31:  return TextAutonumberScheme_RomanLcParenR, nil  
		case 32:  return TextAutonumberScheme_RomanLcPeriod, nil  
		case 33:  return TextAutonumberScheme_RomanUcParenBoth, nil  
		case 34:  return TextAutonumberScheme_RomanUcParenR, nil  
		case 35:  return TextAutonumberScheme_RomanUcPeriod, nil  
		case 36:  return TextAutonumberScheme_ThaiAlphaParenBoth, nil  
		case 37:  return TextAutonumberScheme_ThaiAlphaParenR, nil  
		case 38:  return TextAutonumberScheme_ThaiAlphaPeriod, nil  
		case 39:  return TextAutonumberScheme_ThaiNumParenBoth, nil  
		case 40:  return TextAutonumberScheme_ThaiNumParenR, nil  
		case 41:  return TextAutonumberScheme_ThaiNumPeriod, nil  
		default:
			return 0 ,fmt.Errorf("invalid TextAutonumberScheme value: %d", value)
	}
}

/**************Enum TextFontAlignType *****************/

// Represents the different types of font alignment.
type TextFontAlignType int32

const(
// When the text flow is horizontal or simple vertical same as fontBaseline
// but for other vertical modes same as fontCenter.
TextFontAlignType_Automatic TextFontAlignType = 0 

// The letters are anchored to the very bottom of a single line.
TextFontAlignType_Bottom TextFontAlignType = 1 

// The letters are anchored to the bottom baseline of a single line.
TextFontAlignType_Baseline TextFontAlignType = 2 

// The letters are anchored between the two baselines of a single line.
TextFontAlignType_Center TextFontAlignType = 3 

// The letters are anchored to the top baseline of a single line.
TextFontAlignType_Top TextFontAlignType = 4 
)

func Int32ToTextFontAlignType(value int32)(TextFontAlignType ,error){
	switch value {
		case 0:  return TextFontAlignType_Automatic, nil  
		case 1:  return TextFontAlignType_Bottom, nil  
		case 2:  return TextFontAlignType_Baseline, nil  
		case 3:  return TextFontAlignType_Center, nil  
		case 4:  return TextFontAlignType_Top, nil  
		default:
			return 0 ,fmt.Errorf("invalid TextFontAlignType value: %d", value)
	}
}

/**************Enum TextNodeType *****************/

// Represents the node type.
type TextNodeType int32

const(
// Represents the text node.
TextNodeType_TextRun TextNodeType = 0 

// Represents the text paragraph.
TextNodeType_TextParagraph TextNodeType = 1 

// Represents the equation text.
TextNodeType_Equation TextNodeType = 2 
)

func Int32ToTextNodeType(value int32)(TextNodeType ,error){
	switch value {
		case 0:  return TextNodeType_TextRun, nil  
		case 1:  return TextNodeType_TextParagraph, nil  
		case 2:  return TextNodeType_Equation, nil  
		default:
			return 0 ,fmt.Errorf("invalid TextNodeType value: %d", value)
	}
}

/**************Enum TextTabAlignmentType *****************/

// Represents the text tab alignment types.
type TextTabAlignmentType int32

const(
// The text at this tab stop is center aligned.
TextTabAlignmentType_Center TextTabAlignmentType = 0 

// At this tab stop, the decimals are lined up.
TextTabAlignmentType_Decimal TextTabAlignmentType = 1 

// The text at this tab stop is left aligned.
TextTabAlignmentType_Left TextTabAlignmentType = 2 

// The text at this tab stop is right aligned.
TextTabAlignmentType_Right TextTabAlignmentType = 3 
)

func Int32ToTextTabAlignmentType(value int32)(TextTabAlignmentType ,error){
	switch value {
		case 0:  return TextTabAlignmentType_Center, nil  
		case 1:  return TextTabAlignmentType_Decimal, nil  
		case 2:  return TextTabAlignmentType_Left, nil  
		case 3:  return TextTabAlignmentType_Right, nil  
		default:
			return 0 ,fmt.Errorf("invalid TextTabAlignmentType value: %d", value)
	}
}

/**************Enum TextVerticalType *****************/

// Represents the text direct type.
type TextVerticalType int32

const(
// East Asian Vertical display.
TextVerticalType_Vertical TextVerticalType = 0 

// Horizontal text.
TextVerticalType_Horizontal TextVerticalType = 1 

// Displayed vertical and the text flows top down then LEFT to RIGHT
TextVerticalType_VerticalLeftToRight TextVerticalType = 2 

// Each line is 90 degrees rotated clockwise
TextVerticalType_Vertical90 TextVerticalType = 3 

// Each line is 270 degrees rotated clockwise
TextVerticalType_Vertical270 TextVerticalType = 4 

// Determines if all of the text is vertical
TextVerticalType_Stacked TextVerticalType = 5 

// Specifies that vertical WordArt should be shown from right to left rather than left to right.
TextVerticalType_StackedRightToLeft TextVerticalType = 6 
)

func Int32ToTextVerticalType(value int32)(TextVerticalType ,error){
	switch value {
		case 0:  return TextVerticalType_Vertical, nil  
		case 1:  return TextVerticalType_Horizontal, nil  
		case 2:  return TextVerticalType_VerticalLeftToRight, nil  
		case 3:  return TextVerticalType_Vertical90, nil  
		case 4:  return TextVerticalType_Vertical270, nil  
		case 5:  return TextVerticalType_Stacked, nil  
		case 6:  return TextVerticalType_StackedRightToLeft, nil  
		default:
			return 0 ,fmt.Errorf("invalid TextVerticalType value: %d", value)
	}
}
// Class AutoNumberedBulletValue 

// Represents automatic numbered bullet.
type AutoNumberedBulletValue struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewAutoNumberedBulletValue() ( *AutoNumberedBulletValue, error) {
	autonumberedbulletvalue := &AutoNumberedBulletValue{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_AutoNumberedBulletValue"),)
	if CGoReturnPtr.error_no == 0 {
		autonumberedbulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(autonumberedbulletvalue, DeleteAutoNumberedBulletValue)
		return autonumberedbulletvalue, nil
	} else {
		autonumberedbulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return autonumberedbulletvalue, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - BulletValue 
func NewAutoNumberedBulletValue_BulletValue(src *BulletValue) ( *AutoNumberedBulletValue, error) {
	autonumberedbulletvalue := &AutoNumberedBulletValue{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_AutoNumberedBulletValue_BulletValue"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		autonumberedbulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(autonumberedbulletvalue, DeleteAutoNumberedBulletValue)
		return autonumberedbulletvalue, nil
	} else {
		autonumberedbulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return autonumberedbulletvalue, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *AutoNumberedBulletValue) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("AutoNumberedBulletValue_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the bullet.
// Returns:
//   int32  
func (instance *AutoNumberedBulletValue) GetType()  (BulletType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("AutoNumberedBulletValue_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBulletType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the starting number of the bullet.
// Returns:
//   int32  
func (instance *AutoNumberedBulletValue) GetStartAt()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("AutoNumberedBulletValue_GetStartAt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the starting number of the bullet.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *AutoNumberedBulletValue) SetStartAt(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("AutoNumberedBulletValue_SetStartAt"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the scheme of automatic number.
// Returns:
//   int32  
func (instance *AutoNumberedBulletValue) GetAutonumberScheme()  (TextAutonumberScheme,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("AutoNumberedBulletValue_GetAutonumberScheme"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAutonumberScheme(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Represents the scheme of automatic number.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *AutoNumberedBulletValue) SetAutonumberScheme(value TextAutonumberScheme)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("AutoNumberedBulletValue_SetAutonumberScheme"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *AutoNumberedBulletValue) ToBulletValue() *BulletValue {
	parentClass := &BulletValue{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteAutoNumberedBulletValue(autonumberedbulletvalue *AutoNumberedBulletValue){
	runtime.SetFinalizer(autonumberedbulletvalue, nil)
	C.Delete_CObject(C.CString("Delete_AutoNumberedBulletValue"),autonumberedbulletvalue.ptr)
	autonumberedbulletvalue.ptr = nil
}

// Class Bullet 

// Represents the bullet points should be applied to a paragraph.
type Bullet struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *Bullet) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("Bullet_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the value of bullet.
// Returns:
//   BulletValue  
func (instance *Bullet) GetBulletValue()  (*BulletValue,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("Bullet_GetBulletValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &BulletValue{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteBulletValue) 

	return result, nil 
}
// Gets and sets the type of bullet.
// Returns:
//   int32  
func (instance *Bullet) GetType()  (BulletType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("Bullet_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBulletType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of bullet.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *Bullet) SetType(value BulletType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("Bullet_SetType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Get and sets the name of the font.
// Returns:
//   string  
func (instance *Bullet) GetFontName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("Bullet_GetFontName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Get and sets the name of the font.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *Bullet) SetFontName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("Bullet_SetFontName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteBullet(bullet *Bullet){
	runtime.SetFinalizer(bullet, nil)
	C.Delete_CObject(C.CString("Delete_Bullet"),bullet.ptr)
	bullet.ptr = nil
}

// Class BulletValue 

// Represents the value of the bullet.
type BulletValue struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *BulletValue) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("BulletValue_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the bullet's value.
// Returns:
//   int32  
func (instance *BulletValue) GetType()  (BulletType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("BulletValue_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBulletType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}



func DeleteBulletValue(bulletvalue *BulletValue){
	runtime.SetFinalizer(bulletvalue, nil)
	C.Delete_CObject(C.CString("Delete_BulletValue"),bulletvalue.ptr)
	bulletvalue.ptr = nil
}

// Class CharacterBulletValue 

// Represents the character bullet.
type CharacterBulletValue struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewCharacterBulletValue() ( *CharacterBulletValue, error) {
	characterbulletvalue := &CharacterBulletValue{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_CharacterBulletValue"),)
	if CGoReturnPtr.error_no == 0 {
		characterbulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(characterbulletvalue, DeleteCharacterBulletValue)
		return characterbulletvalue, nil
	} else {
		characterbulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return characterbulletvalue, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - BulletValue 
func NewCharacterBulletValue_BulletValue(src *BulletValue) ( *CharacterBulletValue, error) {
	characterbulletvalue := &CharacterBulletValue{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_CharacterBulletValue_BulletValue"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		characterbulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(characterbulletvalue, DeleteCharacterBulletValue)
		return characterbulletvalue, nil
	} else {
		characterbulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return characterbulletvalue, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CharacterBulletValue) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CharacterBulletValue_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the bullet.
// Returns:
//   int32  
func (instance *CharacterBulletValue) GetType()  (BulletType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CharacterBulletValue_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBulletType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets character of the bullet.
// Returns:
//   byte  
func (instance *CharacterBulletValue) GetCharacter()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("CharacterBulletValue_GetCharacter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets character of the bullet.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *CharacterBulletValue) SetCharacter(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("CharacterBulletValue_SetCharacter"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *CharacterBulletValue) ToBulletValue() *BulletValue {
	parentClass := &BulletValue{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteCharacterBulletValue(characterbulletvalue *CharacterBulletValue){
	runtime.SetFinalizer(characterbulletvalue, nil)
	C.Delete_CObject(C.CString("Delete_CharacterBulletValue"),characterbulletvalue.ptr)
	characterbulletvalue.ptr = nil
}

// Class FontSettingCollection 

// Represents the list of <see cref="FontSetting"/>.
type FontSettingCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *FontSettingCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("FontSettingCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Sets the preset WordArt style.
// Parameters:
//   style - int32 
// Returns:
//   void  
func (instance *FontSettingCollection) SetWordArtStyle(style PresetWordArtStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("FontSettingCollection_SetWordArtStyle"), instance.ptr, C.int( int32(style)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the alignment setting of the text body.
// Returns:
//   ShapeTextAlignment  
func (instance *FontSettingCollection) GetTextAlignment()  (*ShapeTextAlignment,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("FontSettingCollection_GetTextAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShapeTextAlignment{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShapeTextAlignment) 

	return result, nil 
}
// Gets all paragraphs.
// Returns:
//   TextParagraphCollection  
func (instance *FontSettingCollection) GetTextParagraphs()  (*TextParagraphCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("FontSettingCollection_GetTextParagraphs"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextParagraphCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextParagraphCollection) 

	return result, nil 
}
// Gets the enumerator of the paragraphs.
// Returns:
//   unsafe.Pointer  
func (instance *FontSettingCollection) GetParagraphEnumerator()  (*TextParagraphEnumerator,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAW(C.CString("FontSettingCollection_GetParagraphEnumerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextParagraphEnumerator{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteTextParagraphEnumerator)
	 

	return result, nil 
}
// Gets and sets the text of the shape.
// Returns:
//   string  
func (instance *FontSettingCollection) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("FontSettingCollection_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the text of the shape.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *FontSettingCollection) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("FontSettingCollection_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Appends the text.
// Parameters:
//   text - string 
// Returns:
//   void  
func (instance *FontSettingCollection) AppendText(text string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("FontSettingCollection_AppendText"), instance.ptr, C.CString(text))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Insert index at the position.
// Parameters:
//   index - int32 
//   text - string 
// Returns:
//   void  
func (instance *FontSettingCollection) InsertText(index int32, text string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZP(C.CString("FontSettingCollection_InsertText"), instance.ptr, C.int(index), C.CString(text))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Replace the text.
// Parameters:
//   index - int32 
//   count - int32 
//   text - string 
// Returns:
//   void  
func (instance *FontSettingCollection) Replace_Int_Int_String(index int32, count int32, text string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMM(C.CString("FontSettingCollection_Replace_Integer_Integer_String"), instance.ptr, C.int(index), C.int(count), C.CString(text))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Replace the text.
// Parameters:
//   oldValue - string 
//   newValue - string 
// Returns:
//   void  
func (instance *FontSettingCollection) Replace_String_String(oldvalue string, newvalue string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZFN(C.CString("FontSettingCollection_Replace_String_String"), instance.ptr, C.CString(oldvalue), C.CString(newvalue))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Delete some characters.
// Parameters:
//   index - int32 
//   count - int32 
// Returns:
//   void  
func (instance *FontSettingCollection) DeleteText(index int32, count int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZCH(C.CString("FontSettingCollection_DeleteText"), instance.ptr, C.int(index), C.int(count))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the html string which contains data and some formats in this shape.
// Returns:
//   string  
func (instance *FontSettingCollection) GetHtmlString()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("FontSettingCollection_GetHtmlString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the html string which contains data and some formats in this shape.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *FontSettingCollection) SetHtmlString(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("FontSettingCollection_SetHtmlString"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Format the text with font setting.
// Parameters:
//   startIndex - int32 
//   length - int32 
//   font - Font 
//   flag - StyleFlag 
// Returns:
//   void  
func (instance *FontSettingCollection) Format(startindex int32, length int32, font *Font, flag *StyleFlag)  error {
	
	var font_ptr unsafe.Pointer = nil
	if font != nil {
	  font_ptr =font.ptr
	}
	var flag_ptr unsafe.Pointer = nil
	if flag != nil {
	  flag_ptr =flag.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZEQ(C.CString("FontSettingCollection_Format"), instance.ptr, C.int(startindex), C.int(length), font_ptr, flag_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="FontSetting"/> by the index.
// Parameters:
//   index - int32 
// Returns:
//   FontSetting  
func (instance *FontSettingCollection) Get(index int32)  (*FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("FontSettingCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FontSetting{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFontSetting) 

	return result, nil 
}
// Clear all setting.
// Returns:
//   void  
func (instance *FontSettingCollection) Clear()  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZF(C.CString("FontSettingCollection_Clear"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *FontSettingCollection) Equals(obj *Object)  (bool,  error)  {
	
	var obj_ptr unsafe.Pointer = nil
	if obj != nil {
	  obj_ptr =obj.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZBJ(C.CString("FontSettingCollection_Equals"), instance.ptr, obj_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *FontSettingCollection) GetHashCode()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("FontSettingCollection_GetHashCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *FontSettingCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("FontSettingCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteFontSettingCollection(fontsettingcollection *FontSettingCollection){
	runtime.SetFinalizer(fontsettingcollection, nil)
	C.Delete_CObject(C.CString("Delete_FontSettingCollection"),fontsettingcollection.ptr)
	fontsettingcollection.ptr = nil
}

// Class NoneBulletValue 

// Represents no bullet.
type NoneBulletValue struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewNoneBulletValue() ( *NoneBulletValue, error) {
	nonebulletvalue := &NoneBulletValue{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_NoneBulletValue"),)
	if CGoReturnPtr.error_no == 0 {
		nonebulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(nonebulletvalue, DeleteNoneBulletValue)
		return nonebulletvalue, nil
	} else {
		nonebulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nonebulletvalue, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - BulletValue 
func NewNoneBulletValue_BulletValue(src *BulletValue) ( *NoneBulletValue, error) {
	nonebulletvalue := &NoneBulletValue{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_NoneBulletValue_BulletValue"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		nonebulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(nonebulletvalue, DeleteNoneBulletValue)
		return nonebulletvalue, nil
	} else {
		nonebulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return nonebulletvalue, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *NoneBulletValue) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("NoneBulletValue_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the bullet's value.
// Returns:
//   int32  
func (instance *NoneBulletValue) GetType()  (BulletType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("NoneBulletValue_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBulletType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}


func (instance *NoneBulletValue) ToBulletValue() *BulletValue {
	parentClass := &BulletValue{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteNoneBulletValue(nonebulletvalue *NoneBulletValue){
	runtime.SetFinalizer(nonebulletvalue, nil)
	C.Delete_CObject(C.CString("Delete_NoneBulletValue"),nonebulletvalue.ptr)
	nonebulletvalue.ptr = nil
}

// Class PictureBulletValue 

// Represents the value of the image bullet.
type PictureBulletValue struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPictureBulletValue() ( *PictureBulletValue, error) {
	picturebulletvalue := &PictureBulletValue{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_PictureBulletValue"),)
	if CGoReturnPtr.error_no == 0 {
		picturebulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(picturebulletvalue, DeletePictureBulletValue)
		return picturebulletvalue, nil
	} else {
		picturebulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return picturebulletvalue, err
	}	
}
// Constructs from a parent object.
// Parameters:
//   src - BulletValue 
func NewPictureBulletValue_BulletValue(src *BulletValue) ( *PictureBulletValue, error) {
	picturebulletvalue := &PictureBulletValue{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_PictureBulletValue_BulletValue"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		picturebulletvalue.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(picturebulletvalue, DeletePictureBulletValue)
		return picturebulletvalue, nil
	} else {
		picturebulletvalue.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return picturebulletvalue, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PictureBulletValue) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("PictureBulletValue_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the bullet's value.
// Returns:
//   int32  
func (instance *PictureBulletValue) GetType()  (BulletType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("PictureBulletValue_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToBulletType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets image data of the bullet.
// Returns:
//   []byte  
func (instance *PictureBulletValue) GetImageData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("PictureBulletValue_GetImageData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets image data of the bullet.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *PictureBulletValue) SetImageData(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("PictureBulletValue_SetImageData"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *PictureBulletValue) ToBulletValue() *BulletValue {
	parentClass := &BulletValue{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeletePictureBulletValue(picturebulletvalue *PictureBulletValue){
	runtime.SetFinalizer(picturebulletvalue, nil)
	C.Delete_CObject(C.CString("Delete_PictureBulletValue"),picturebulletvalue.ptr)
	picturebulletvalue.ptr = nil
}

// Class ShapeTextAlignment 

// Represents the setting of shape's text alignment;
type ShapeTextAlignment struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ShapeTextAlignment) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ShapeTextAlignment_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text wrapped type of the shape which contains text.
// Returns:
//   bool  
func (instance *ShapeTextAlignment) IsTextWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ShapeTextAlignment_IsTextWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the text wrapped type of the shape which contains text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetIsTextWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ShapeTextAlignment_SetIsTextWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether rotating text with shape.
// Returns:
//   bool  
func (instance *ShapeTextAlignment) GetRotateTextWithShape()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ShapeTextAlignment_GetRotateTextWithShape"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether rotating text with shape.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetRotateTextWithShape(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ShapeTextAlignment_SetRotateTextWithShape"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text vertical overflow type of the text box.
// Returns:
//   int32  
func (instance *ShapeTextAlignment) GetTextVerticalOverflow()  (TextOverflowType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ShapeTextAlignment_GetTextVerticalOverflow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextOverflowType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text vertical overflow type of the text box.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetTextVerticalOverflow(value TextOverflowType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ShapeTextAlignment_SetTextVerticalOverflow"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal overflow type of the text box.
// Returns:
//   int32  
func (instance *ShapeTextAlignment) GetTextHorizontalOverflow()  (TextOverflowType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ShapeTextAlignment_GetTextHorizontalOverflow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextOverflowType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal overflow type of the text box.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetTextHorizontalOverflow(value TextOverflowType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ShapeTextAlignment_SetTextHorizontalOverflow"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the rotation of the shape.
// Returns:
//   float64  
func (instance *ShapeTextAlignment) GetRotationAngle()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ShapeTextAlignment_GetRotationAngle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the rotation of the shape.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetRotationAngle(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ShapeTextAlignment_SetRotationAngle"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text direction.
// Returns:
//   int32  
func (instance *ShapeTextAlignment) GetTextVerticalType()  (TextVerticalType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ShapeTextAlignment_GetTextVerticalType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextVerticalType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text direction.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetTextVerticalType(value TextVerticalType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ShapeTextAlignment_SetTextVerticalType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the shape is locked when worksheet is protected.
// Returns:
//   bool  
func (instance *ShapeTextAlignment) IsLockedText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ShapeTextAlignment_IsLockedText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the shape is locked when worksheet is protected.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetIsLockedText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ShapeTextAlignment_SetIsLockedText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if size of shape is adjusted automatically according to its content.
// Returns:
//   bool  
func (instance *ShapeTextAlignment) GetAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ShapeTextAlignment_GetAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates if size of shape is adjusted automatically according to its content.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ShapeTextAlignment_SetAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the transform type of text.
// Returns:
//   int32  
func (instance *ShapeTextAlignment) GetTextShapeType()  (AutoShapeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ShapeTextAlignment_GetTextShapeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToAutoShapeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the transform type of text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetTextShapeType(value AutoShapeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ShapeTextAlignment_SetTextShapeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the top margin in unit of Points
// Returns:
//   float64  
func (instance *ShapeTextAlignment) GetTopMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ShapeTextAlignment_GetTopMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the top margin in unit of Points
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetTopMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ShapeTextAlignment_SetTopMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the bottom margin in unit of Points
// Returns:
//   float64  
func (instance *ShapeTextAlignment) GetBottomMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ShapeTextAlignment_GetBottomMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the bottom margin in unit of Points
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetBottomMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ShapeTextAlignment_SetBottomMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the left margin in unit of Points
// Returns:
//   float64  
func (instance *ShapeTextAlignment) GetLeftMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ShapeTextAlignment_GetLeftMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the left margin in unit of Points
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetLeftMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ShapeTextAlignment_SetLeftMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the right margin in unit of Points
// Returns:
//   float64  
func (instance *ShapeTextAlignment) GetRightMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ShapeTextAlignment_GetRightMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the right margin in unit of Points
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetRightMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ShapeTextAlignment_SetRightMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the margin of the text frame is automatic.
// Returns:
//   bool  
func (instance *ShapeTextAlignment) IsAutoMargin()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ShapeTextAlignment_IsAutoMargin"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the margin of the text frame is automatic.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetIsAutoMargin(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ShapeTextAlignment_SetIsAutoMargin"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the number of columns of text in the bounding rectangle.
// Returns:
//   int32  
func (instance *ShapeTextAlignment) GetNumberOfColumns()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ShapeTextAlignment_GetNumberOfColumns"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the number of columns of text in the bounding rectangle.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ShapeTextAlignment) SetNumberOfColumns(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ShapeTextAlignment_SetNumberOfColumns"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determines whether this instance has the same value as another specified <see cref="ShapeTextAlignment"/> object.
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *ShapeTextAlignment) Equals(obj *Object)  (bool,  error)  {
	
	var obj_ptr unsafe.Pointer = nil
	if obj != nil {
	  obj_ptr =obj.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZBJ(C.CString("ShapeTextAlignment_Equals"), instance.ptr, obj_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *ShapeTextAlignment) GetHashCode()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ShapeTextAlignment_GetHashCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteShapeTextAlignment(shapetextalignment *ShapeTextAlignment){
	runtime.SetFinalizer(shapetextalignment, nil)
	C.Delete_CObject(C.CString("Delete_ShapeTextAlignment"),shapetextalignment.ptr)
	shapetextalignment.ptr = nil
}

// Class TextBoxOptions 

// Represents the text options of the shape
type TextBoxOptions struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextBoxOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// It corresponds to "Format Shape - Text Options - Text Box - Vertical Alignment" in Excel.
// Returns:
//   int32  
func (instance *TextBoxOptions) GetShapeTextVerticalAlignment()  (ShapeTextVerticalAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxOptions_GetShapeTextVerticalAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToShapeTextVerticalAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// It corresponds to "Format Shape - Text Options - Text Box - Vertical Alignment" in Excel.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxOptions) SetShapeTextVerticalAlignment(value ShapeTextVerticalAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextBoxOptions_SetShapeTextVerticalAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to resize the shape to fit the text
// Returns:
//   bool  
func (instance *TextBoxOptions) GetResizeToFitText()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxOptions_GetResizeToFitText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to resize the shape to fit the text
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxOptions) SetResizeToFitText(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxOptions_SetResizeToFitText"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the text display direction within a given text body.
// It corresponds to "Format Shape - Text Options - Text Box - Text direction" in Excel
// Returns:
//   int32  
func (instance *TextBoxOptions) GetShapeTextDirection()  (TextVerticalType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxOptions_GetShapeTextDirection"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextVerticalType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the text display direction within a given text body.
// It corresponds to "Format Shape - Text Options - Text Box - Text direction" in Excel
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxOptions) SetShapeTextDirection(value TextVerticalType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextBoxOptions_SetShapeTextDirection"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the left margin in unit of Points.
// Returns:
//   float64  
func (instance *TextBoxOptions) GetLeftMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextBoxOptions_GetLeftMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the left margin in unit of Points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextBoxOptions) SetLeftMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextBoxOptions_SetLeftMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the right margin in unit of Points.
// Returns:
//   float64  
func (instance *TextBoxOptions) GetRightMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextBoxOptions_GetRightMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the right margin in unit of Points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextBoxOptions) SetRightMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextBoxOptions_SetRightMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the top margin in unit of Points.
// Returns:
//   float64  
func (instance *TextBoxOptions) GetTopMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextBoxOptions_GetTopMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the top margin in unit of Points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextBoxOptions) SetTopMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextBoxOptions_SetTopMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the bottom margin in unit of Points
// Returns:
//   float64  
func (instance *TextBoxOptions) GetBottomMarginPt()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextBoxOptions_GetBottomMarginPt"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the bottom margin in unit of Points
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextBoxOptions) SetBottomMarginPt(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextBoxOptions_SetBottomMarginPt"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether allow text to overflow shape.
// Returns:
//   bool  
func (instance *TextBoxOptions) GetAllowTextToOverflow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxOptions_GetAllowTextToOverflow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Whether allow text to overflow shape.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxOptions) SetAllowTextToOverflow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxOptions_SetAllowTextToOverflow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies text wrapping within a shape.
// False - No wrapping will occur on text body.
// True - Wrapping will occur on text body.
// Returns:
//   bool  
func (instance *TextBoxOptions) GetWrapTextInShape()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxOptions_GetWrapTextInShape"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies text wrapping within a shape.
// False - No wrapping will occur on text body.
// True - Wrapping will occur on text body.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxOptions) SetWrapTextInShape(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxOptions_SetWrapTextInShape"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteTextBoxOptions(textboxoptions *TextBoxOptions){
	runtime.SetFinalizer(textboxoptions, nil)
	C.Delete_CObject(C.CString("Delete_TextBoxOptions"),textboxoptions.ptr)
	textboxoptions.ptr = nil
}

// Class TextOptions 

// Represents the text options.
type TextOptions struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - Font 
func NewTextOptions(src *Font) ( *TextOptions, error) {
	textoptions := &TextOptions{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_TextOptions"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		textoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(textoptions, DeleteTextOptions)
		return textoptions, nil
	} else {
		textoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return textoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextOptions) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the shape.
// Returns:
//   string  
func (instance *TextOptions) GetName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextOptions_GetName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the name of the shape.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextOptions) SetName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("TextOptions_SetName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the user interface language.
// Returns:
//   int32  
func (instance *TextOptions) GetLanguageCode()  (CountryCode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextOptions_GetLanguageCode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCountryCode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the user interface language.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetLanguageCode(value CountryCode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextOptions_SetLanguageCode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the latin name.
// Returns:
//   string  
func (instance *TextOptions) GetLatinName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextOptions_GetLatinName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the latin name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextOptions) SetLatinName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("TextOptions_SetLatinName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the FarEast name.
// Returns:
//   string  
func (instance *TextOptions) GetFarEastName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextOptions_GetFarEastName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the FarEast name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextOptions) SetFarEastName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZN(C.CString("TextOptions_SetFarEastName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the fill format of the text.
// Returns:
//   FillFormat  
func (instance *TextOptions) GetFill()  (*FillFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextOptions_GetFill"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &FillFormat{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFillFormat) 

	return result, nil 
}
// Represents the outline format of the text.
// Returns:
//   LineFormat  
func (instance *TextOptions) GetOutline()  (*LineFormat,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextOptions_GetOutline"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &LineFormat{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteLineFormat) 

	return result, nil 
}
// Represents a <see cref="ShadowEffect"/> object that specifies shadow effect for the chart element or shape.
// Returns:
//   ShadowEffect  
func (instance *TextOptions) GetShadow()  (*ShadowEffect,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextOptions_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ShadowEffect{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShadowEffect) 

	return result, nil 
}
// Gets or sets the color of underline.
// Returns:
//   CellsColor  
func (instance *TextOptions) GetUnderlineColor()  (*CellsColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextOptions_GetUnderlineColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &CellsColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCellsColor) 

	return result, nil 
}
// Gets or sets the color of underline.
// Parameters:
//   value - CellsColor 
// Returns:
//   void  
func (instance *TextOptions) SetUnderlineColor(value *CellsColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("TextOptions_SetUnderlineColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the minimum font size at which character kerning will occur for this text run.
// Returns:
//   float64  
func (instance *TextOptions) GetKerning()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextOptions_GetKerning"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the minimum font size at which character kerning will occur for this text run.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextOptions) SetKerning(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextOptions_SetKerning"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the spacing between characters within a text run.
// Returns:
//   float64  
func (instance *TextOptions) GetSpacing()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextOptions_GetSpacing"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the spacing between characters within a text run.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextOptions) SetSpacing(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextOptions_SetSpacing"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represent the character set.
// Returns:
//   int32  
func (instance *TextOptions) GetCharset()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextOptions_GetCharset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represent the character set.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetCharset(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextOptions_SetCharset"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the font is italic.
// Returns:
//   bool  
func (instance *TextOptions) IsItalic()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsItalic"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the font is italic.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextOptions) SetIsItalic(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextOptions_SetIsItalic"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the font is bold.
// Returns:
//   bool  
func (instance *TextOptions) IsBold()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsBold"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the font is bold.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextOptions) SetIsBold(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextOptions_SetIsBold"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text caps type.
// Returns:
//   int32  
func (instance *TextOptions) GetCapsType()  (TextCapsType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextOptions_GetCapsType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextCapsType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text caps type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetCapsType(value TextCapsType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextOptions_SetCapsType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the strike type of the text.
// Returns:
//   int32  
func (instance *TextOptions) GetStrikeType()  (TextStrikeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextOptions_GetStrikeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextStrikeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets the strike type of the text.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetStrikeType(value TextStrikeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextOptions_SetStrikeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the font is single strikeout.
// Returns:
//   bool  
func (instance *TextOptions) IsStrikeout()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsStrikeout"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the font is single strikeout.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextOptions) SetIsStrikeout(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextOptions_SetIsStrikeout"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the script offset,in unit of percentage
// Returns:
//   float64  
func (instance *TextOptions) GetScriptOffset()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextOptions_GetScriptOffset"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the script offset,in unit of percentage
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextOptions) SetScriptOffset(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextOptions_SetScriptOffset"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the font is super script.
// Returns:
//   bool  
func (instance *TextOptions) IsSuperscript()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsSuperscript"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the font is super script.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextOptions) SetIsSuperscript(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextOptions_SetIsSuperscript"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value indicating whether the font is subscript.
// Returns:
//   bool  
func (instance *TextOptions) IsSubscript()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsSubscript"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value indicating whether the font is subscript.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextOptions) SetIsSubscript(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextOptions_SetIsSubscript"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the font underline type.
// Returns:
//   int32  
func (instance *TextOptions) GetUnderline()  (FontUnderlineType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextOptions_GetUnderline"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFontUnderlineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the font underline type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetUnderline(value FontUnderlineType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextOptions_SetUnderline"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the double size of the font.
// Returns:
//   float64  
func (instance *TextOptions) GetDoubleSize()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextOptions_GetDoubleSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the double size of the font.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextOptions) SetDoubleSize(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextOptions_SetDoubleSize"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the size of the font.
// Returns:
//   int32  
func (instance *TextOptions) GetSize()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextOptions_GetSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the size of the font.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetSize(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextOptions_SetSize"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the theme color.
// Returns:
//   ThemeColor  
func (instance *TextOptions) GetThemeColor()  (*ThemeColor,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextOptions_GetThemeColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &ThemeColor{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteThemeColor) 

	return result, nil 
}
// Gets and sets the theme color.
// Parameters:
//   value - ThemeColor 
// Returns:
//   void  
func (instance *TextOptions) SetThemeColor(value *ThemeColor)  error {
	
	var value_ptr unsafe.Pointer = nil
	if value != nil {
	  value_ptr =value.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("TextOptions_SetThemeColor"), instance.ptr, value_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the <see cref="Color"/> of the font.
// Returns:
//   Color  
func (instance *TextOptions) GetColor()  (*Color,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAD(C.CString("TextOptions_GetColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets the <see cref="Color"/> of the font.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *TextOptions) SetColor(value *Color)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAE(C.CString("TextOptions_SetColor"), instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the color with a 32-bit ARGB value.
// Returns:
//   int32  
func (instance *TextOptions) GetArgbColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextOptions_GetArgbColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the color with a 32-bit ARGB value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetArgbColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextOptions_SetArgbColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Checks if two fonts are equals.
// Parameters:
//   font - Font 
// Returns:
//   bool  
func (instance *TextOptions) Equals(font *Font)  (bool,  error)  {
	
	var font_ptr unsafe.Pointer = nil
	if font != nil {
	  font_ptr =font.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZBJ(C.CString("TextOptions_Equals"), instance.ptr, font_ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the normalization of height that is to be applied to the text run.
// Returns:
//   bool  
func (instance *TextOptions) IsNormalizeHeights()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextOptions_IsNormalizeHeights"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the normalization of height that is to be applied to the text run.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextOptions) SetIsNormalizeHeights(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextOptions_SetIsNormalizeHeights"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the scheme type of the font.
// Returns:
//   int32  
func (instance *TextOptions) GetSchemeType()  (FontSchemeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextOptions_GetSchemeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToFontSchemeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the scheme type of the font.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextOptions) SetSchemeType(value FontSchemeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextOptions_SetSchemeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns a string represents the current Cell object.
// Returns:
//   string  
func (instance *TextOptions) ToString()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextOptions_ToString"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}


func (instance *TextOptions) ToFont() *Font {
	parentClass := &Font{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteTextOptions(textoptions *TextOptions){
	runtime.SetFinalizer(textoptions, nil)
	C.Delete_CObject(C.CString("Delete_TextOptions"),textoptions.ptr)
	textoptions.ptr = nil
}

// Class TextParagraph 

// Represents the text paragraph setting.
type TextParagraph struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - FontSetting 
func NewTextParagraph(src *FontSetting) ( *TextParagraph, error) {
	textparagraph := &TextParagraph{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("New_TextParagraph"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		textparagraph.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(textparagraph, DeleteTextParagraph)
		return textparagraph, nil
	} else {
		textparagraph.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return textparagraph, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextParagraph) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextParagraph_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the bullet.
// Returns:
//   Bullet  
func (instance *TextParagraph) GetBullet()  (*Bullet,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextParagraph_GetBullet"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Bullet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteBullet) 

	return result, nil 
}
// Gets the type of text node.
// Returns:
//   int32  
func (instance *TextParagraph) GetType()  (TextNodeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextParagraph_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be used within a paragraph.
// Returns:
//   int32  
func (instance *TextParagraph) GetLineSpaceSizeType()  (LineSpaceSizeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextParagraph_GetLineSpaceSizeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineSpaceSizeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be used within a paragraph.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextParagraph) SetLineSpaceSizeType(value LineSpaceSizeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextParagraph_SetLineSpaceSizeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount of vertical white space that will be used within a paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetLineSpace()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetLineSpace"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be used within a paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetLineSpace(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetLineSpace"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount of vertical white space that will be present after a paragraph.
// Returns:
//   int32  
func (instance *TextParagraph) GetSpaceAfterSizeType()  (LineSpaceSizeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextParagraph_GetSpaceAfterSizeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineSpaceSizeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be present after a paragraph.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextParagraph) SetSpaceAfterSizeType(value LineSpaceSizeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextParagraph_SetSpaceAfterSizeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount of vertical white space that will be present after a paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetSpaceAfter()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetSpaceAfter"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be present after a paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetSpaceAfter(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetSpaceAfter"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount of vertical white space that will be present before a paragraph.
// Returns:
//   int32  
func (instance *TextParagraph) GetSpaceBeforeSizeType()  (LineSpaceSizeType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextParagraph_GetSpaceBeforeSizeType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToLineSpaceSizeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be present before a paragraph.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextParagraph) SetSpaceBeforeSizeType(value LineSpaceSizeType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextParagraph_SetSpaceBeforeSizeType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount of vertical white space that will be present before a paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetSpaceBefore()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetSpaceBefore"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the amount of vertical white space that will be present before a paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetSpaceBefore(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetSpaceBefore"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets tab stop list.
// Returns:
//   TextTabStopCollection  
func (instance *TextParagraph) GetStops()  (*TextTabStopCollection,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextParagraph_GetStops"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextTabStopCollection{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextTabStopCollection) 

	return result, nil 
}
// Specifies whether a Latin word can be broken in half and wrapped onto the next line without a hyphen being added.
// Returns:
//   bool  
func (instance *TextParagraph) IsLatinLineBreak()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextParagraph_IsLatinLineBreak"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether a Latin word can be broken in half and wrapped onto the next line without a hyphen being added.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextParagraph) SetIsLatinLineBreak(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextParagraph_SetIsLatinLineBreak"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether an East Asian word can be broken in half and wrapped onto the next line without a hyphen being added.
// Returns:
//   bool  
func (instance *TextParagraph) IsEastAsianLineBreak()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextParagraph_IsEastAsianLineBreak"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether an East Asian word can be broken in half and wrapped onto the next line without a hyphen being added.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextParagraph) SetIsEastAsianLineBreak(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextParagraph_SetIsEastAsianLineBreak"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies whether punctuation is to be forcefully laid out on a line of text or put on a different line of text.
// Returns:
//   bool  
func (instance *TextParagraph) IsHangingPunctuation()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextParagraph_IsHangingPunctuation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies whether punctuation is to be forcefully laid out on a line of text or put on a different line of text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextParagraph) SetIsHangingPunctuation(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextParagraph_SetIsHangingPunctuation"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the right margin of the paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetRightMargin()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetRightMargin"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the right margin of the paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetRightMargin(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetRightMargin"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the left margin of the paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetLeftMargin()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetLeftMargin"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the left margin of the paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetLeftMargin(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetLeftMargin"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the indent size that will be applied to the first line of text in the paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetFirstLineIndent()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetFirstLineIndent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the indent size that will be applied to the first line of text in the paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetFirstLineIndent(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetFirstLineIndent"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determines where vertically on a line of text the actual words are positioned. This deals
// with vertical placement of the characters with respect to the baselines.
// Returns:
//   int32  
func (instance *TextParagraph) GetFontAlignType()  (TextFontAlignType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextParagraph_GetFontAlignType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextFontAlignType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Determines where vertically on a line of text the actual words are positioned. This deals
// with vertical placement of the characters with respect to the baselines.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextParagraph) SetFontAlignType(value TextFontAlignType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextParagraph_SetFontAlignType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the text horizontal alignment type of the paragraph.
// Returns:
//   int32  
func (instance *TextParagraph) GetAlignmentType()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextParagraph_GetAlignmentType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the text horizontal alignment type of the paragraph.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextParagraph) SetAlignmentType(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextParagraph_SetAlignmentType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default size for a tab character within this paragraph.
// Returns:
//   float64  
func (instance *TextParagraph) GetDefaultTabSize()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextParagraph_GetDefaultTabSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the default size for a tab character within this paragraph.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextParagraph) SetDefaultTabSize(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextParagraph_SetDefaultTabSize"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets all text runs in this paragraph.
// If this paragraph is empty, return paragraph itself.
// Returns:
//   []FontSetting  
func (instance *TextParagraph) GetChildren()  ([]FontSetting,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZY(C.CString("TextParagraph_GetChildren"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result:= make([]FontSetting, CGoReturnPtr.column_length)
	for i := 0; i < int(CGoReturnPtr.column_length); i++ {
	   offset := uintptr(C.size_t(i)) * uintptr(CGoReturnPtr.size)
	   goObject := &FontSetting{}
	   goObject.ptr =unsafe.Pointer(uintptr( unsafe.Pointer(CGoReturnPtr.return_value)) + offset)
	   result[i] = *goObject
	}
	 

	return result, nil 
}
// Gets the start index of the characters.
// Returns:
//   int32  
func (instance *TextParagraph) GetStartIndex()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextParagraph_GetStartIndex"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the length of the characters.
// Returns:
//   int32  
func (instance *TextParagraph) GetLength()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextParagraph_GetLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Returns the font of this object.
// Returns:
//   Font  
func (instance *TextParagraph) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextParagraph_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Sets the preset WordArt style.
// Parameters:
//   style - int32 
// Returns:
//   void  
func (instance *TextParagraph) SetWordArtStyle(style PresetWordArtStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextParagraph_SetWordArtStyle"), instance.ptr, C.int( int32(style)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the text options.
// Returns:
//   TextOptions  
func (instance *TextParagraph) GetTextOptions()  (*TextOptions,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextParagraph_GetTextOptions"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}


func (instance *TextParagraph) ToFontSetting() *FontSetting {
	parentClass := &FontSetting{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteTextParagraph(textparagraph *TextParagraph){
	runtime.SetFinalizer(textparagraph, nil)
	C.Delete_CObject(C.CString("Delete_TextParagraph"),textparagraph.ptr)
	textparagraph.ptr = nil
}

// Class TextParagraphCollection 

// Represents all text paragraph.
type TextParagraphCollection struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextParagraphCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextParagraphCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the count of text paragraphs.
// Returns:
//   int32  
func (instance *TextParagraphCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextParagraphCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="TextParagraph"/> object at specific index.
// Parameters:
//   index - int32 
// Returns:
//   TextParagraph  
func (instance *TextParagraphCollection) Get(index int32)  (*TextParagraph,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("TextParagraphCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextParagraph{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextParagraph) 

	return result, nil 
}
// Gets the enumerator of the paragraphs.
// Returns:
//   unsafe.Pointer  
func (instance *TextParagraphCollection) GetEnumerator()  (*TextParagraphEnumerator,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAW(C.CString("TextParagraphCollection_GetEnumerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextParagraphEnumerator{}
	result.ptr = CGoReturnPtr.return_value
	runtime.SetFinalizer(result, DeleteTextParagraphEnumerator)
	 

	return result, nil 
}



func DeleteTextParagraphCollection(textparagraphcollection *TextParagraphCollection){
	runtime.SetFinalizer(textparagraphcollection, nil)
	C.Delete_CObject(C.CString("Delete_TextParagraphCollection"),textparagraphcollection.ptr)
	textparagraphcollection.ptr = nil
}

// Class TextTabStop 

// Represents tab stop.
type TextTabStop struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextTabStop) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextTabStop_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the alignment that is to be applied to text using this tab stop.
// Returns:
//   int32  
func (instance *TextTabStop) GetTabAlignment()  (TextTabAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextTabStop_GetTabAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextTabAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the alignment that is to be applied to text using this tab stop.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextTabStop) SetTabAlignment(value TextTabAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextTabStop_SetTabAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the position of the tab stop relative to the left margin.
// Returns:
//   float64  
func (instance *TextTabStop) GetTabPosition()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextTabStop_GetTabPosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the position of the tab stop relative to the left margin.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextTabStop) SetTabPosition(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextTabStop_SetTabPosition"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}



func DeleteTextTabStop(texttabstop *TextTabStop){
	runtime.SetFinalizer(texttabstop, nil)
	C.Delete_CObject(C.CString("Delete_TextTabStop"),texttabstop.ptr)
	texttabstop.ptr = nil
}

// Class TextTabStopCollection 

// Represents the list of all tab stops.
type TextTabStopCollection struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewTextTabStopCollection() ( *TextTabStopCollection, error) {
	texttabstopcollection := &TextTabStopCollection{}
	CGoReturnPtr := C.CellsGoFunctoinZZZA(C.CString("New_TextTabStopCollection"),)
	if CGoReturnPtr.error_no == 0 {
		texttabstopcollection.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(texttabstopcollection, DeleteTextTabStopCollection)
		return texttabstopcollection, nil
	} else {
		texttabstopcollection.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return texttabstopcollection, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextTabStopCollection) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextTabStopCollection_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Adds a tab stop.
// Parameters:
//   tabAlignment - int32 
//   tabPosition - float64 
// Returns:
//   int32  
func (instance *TextTabStopCollection) Add(tabalignment TextTabAlignmentType, tabposition float64)  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZMN(C.CString("TextTabStopCollection_Add"), instance.ptr, C.int( int32(tabalignment)), C.double(tabposition))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets <see cref="TextTabStop"/> by the index.
// Parameters:
//   index - int32 
// Returns:
//   TextTabStop  
func (instance *TextTabStopCollection) Get(index int32)  (*TextTabStop,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAG(C.CString("TextTabStopCollection_Get"), instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextTabStop{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextTabStop) 

	return result, nil 
}
// Returns:
//   int32  
func (instance *TextTabStopCollection) GetCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextTabStopCollection_GetCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}



func DeleteTextTabStopCollection(texttabstopcollection *TextTabStopCollection){
	runtime.SetFinalizer(texttabstopcollection, nil)
	C.Delete_CObject(C.CString("Delete_TextTabStopCollection"),texttabstopcollection.ptr)
	texttabstopcollection.ptr = nil
}
