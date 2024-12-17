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

/**************Enum EquationCharacterPositionType *****************/

// Specifies the position of a particular subobject within its parent
type EquationCharacterPositionType int32

const(
// At the top of the parent object
EquationCharacterPositionType_Top EquationCharacterPositionType = 0 

// At the bottom of the parent object
EquationCharacterPositionType_Bottom EquationCharacterPositionType = 1 
)

func Int32ToEquationCharacterPositionType(value int32)(EquationCharacterPositionType ,error){
	switch value {
		case 0:  return EquationCharacterPositionType_Top, nil  
		case 1:  return EquationCharacterPositionType_Bottom, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationCharacterPositionType value: %d", value)
	}
}

/**************Enum EquationCombiningCharacterType *****************/

// Type of combining characters.
type EquationCombiningCharacterType int32

const(
// Use unknown type when not found in existing type.
EquationCombiningCharacterType_Unknown EquationCombiningCharacterType = -1 

// "̇" Unicode: u0307
// Combining Dot Above
EquationCombiningCharacterType_DotAbove EquationCombiningCharacterType = 0 

// "̈" Unicode: u0308
// Combining Diaeresis
EquationCombiningCharacterType_Diaeresis EquationCombiningCharacterType = 1 

// "⃛" Unicode: u20db
// Combining Three Dots Above
EquationCombiningCharacterType_ThreeDotsAbove EquationCombiningCharacterType = 2 

// "̂" Unicode: u0302
// Combining Circumflex Accent
EquationCombiningCharacterType_CircumflexAccent EquationCombiningCharacterType = 3 

// "̌" Unicode: u030c
// Combining Caron
EquationCombiningCharacterType_Caron EquationCombiningCharacterType = 4 

// "́" Unicode: u0301
// Combining Acute Accent
EquationCombiningCharacterType_AcuteAccent EquationCombiningCharacterType = 5 

// "̀" Unicode: u0300
// Combining Grave Accent
EquationCombiningCharacterType_GraveAccent EquationCombiningCharacterType = 6 

// "̆" Unicode: u0306
// Combining Combining Breve
EquationCombiningCharacterType_Breve EquationCombiningCharacterType = 7 

// "̃" Unicode: u0303
// Combining Tilde
EquationCombiningCharacterType_Tilde EquationCombiningCharacterType = 8 

// "̅" Unicode: u0305
// Combining Overline
EquationCombiningCharacterType_Overline EquationCombiningCharacterType = 9 

// "̿" Unicode: u033f
// Combining Double Overline
EquationCombiningCharacterType_DoubleOverline EquationCombiningCharacterType = 10 

// "⏞" Unicode: u23de
// Combining Top Curly Bracket
EquationCombiningCharacterType_TopCurlyBracket EquationCombiningCharacterType = 11 

// "⏟" Unicode: u23df
// Combining Bottom Curly Bracket
EquationCombiningCharacterType_BottomCurlyBracket EquationCombiningCharacterType = 12 

// "⃖" Unicode: u20d6
// Combining Left Arrow Above
EquationCombiningCharacterType_LeftArrowAbove EquationCombiningCharacterType = 13 

// "⃗" Unicode: u20d7
// Combining Right Arrow Above
EquationCombiningCharacterType_RightArrowAbove EquationCombiningCharacterType = 14 

// "⃡" Unicode: u20e1
// Combining Left Right Arrow Above
EquationCombiningCharacterType_LeftRightArrowAbove EquationCombiningCharacterType = 15 

// "⃐" Unicode: u20d0
// Combining Left Harpoon Above
EquationCombiningCharacterType_LeftHarpoonAbove EquationCombiningCharacterType = 16 

// "⃑" Unicode: u20d1
// Combining Right Harpoon Above
EquationCombiningCharacterType_RightHarpoonAbove EquationCombiningCharacterType = 17 

// "←" Unicode: u2190
// Leftwards Arrow
EquationCombiningCharacterType_LeftwardsArrow EquationCombiningCharacterType = 18 

// "→" Unicode: u2192
// Rightwards Arrow
EquationCombiningCharacterType_RightwardsArrow EquationCombiningCharacterType = 19 

// "↔" Unicode: u2194
// Left Right Arrow
EquationCombiningCharacterType_LeftRightArrow EquationCombiningCharacterType = 20 

// "⇐" Unicode: u21d0
// Leftwards Double Arrow
EquationCombiningCharacterType_LeftwardsDoubleArrow EquationCombiningCharacterType = 21 

// "⇒" Unicode: u21d2
// Rightwards Double Arrow
EquationCombiningCharacterType_RightwardsDoubleArrow EquationCombiningCharacterType = 22 

// "⇔" Unicode: u21d4
// Left Right Double Arrow
EquationCombiningCharacterType_LeftRightDoubleArrow EquationCombiningCharacterType = 23 
)

func Int32ToEquationCombiningCharacterType(value int32)(EquationCombiningCharacterType ,error){
	switch value {
		case -1:  return EquationCombiningCharacterType_Unknown, nil  
		case 0:  return EquationCombiningCharacterType_DotAbove, nil  
		case 1:  return EquationCombiningCharacterType_Diaeresis, nil  
		case 2:  return EquationCombiningCharacterType_ThreeDotsAbove, nil  
		case 3:  return EquationCombiningCharacterType_CircumflexAccent, nil  
		case 4:  return EquationCombiningCharacterType_Caron, nil  
		case 5:  return EquationCombiningCharacterType_AcuteAccent, nil  
		case 6:  return EquationCombiningCharacterType_GraveAccent, nil  
		case 7:  return EquationCombiningCharacterType_Breve, nil  
		case 8:  return EquationCombiningCharacterType_Tilde, nil  
		case 9:  return EquationCombiningCharacterType_Overline, nil  
		case 10:  return EquationCombiningCharacterType_DoubleOverline, nil  
		case 11:  return EquationCombiningCharacterType_TopCurlyBracket, nil  
		case 12:  return EquationCombiningCharacterType_BottomCurlyBracket, nil  
		case 13:  return EquationCombiningCharacterType_LeftArrowAbove, nil  
		case 14:  return EquationCombiningCharacterType_RightArrowAbove, nil  
		case 15:  return EquationCombiningCharacterType_LeftRightArrowAbove, nil  
		case 16:  return EquationCombiningCharacterType_LeftHarpoonAbove, nil  
		case 17:  return EquationCombiningCharacterType_RightHarpoonAbove, nil  
		case 18:  return EquationCombiningCharacterType_LeftwardsArrow, nil  
		case 19:  return EquationCombiningCharacterType_RightwardsArrow, nil  
		case 20:  return EquationCombiningCharacterType_LeftRightArrow, nil  
		case 21:  return EquationCombiningCharacterType_LeftwardsDoubleArrow, nil  
		case 22:  return EquationCombiningCharacterType_RightwardsDoubleArrow, nil  
		case 23:  return EquationCombiningCharacterType_LeftRightDoubleArrow, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationCombiningCharacterType value: %d", value)
	}
}

/**************Enum EquationDelimiterShapeType *****************/

// This specifies the shape of delimiters in the delimiter object.
type EquationDelimiterShapeType int32

const(
// The divider is centered around the entire height of its content.
EquationDelimiterShapeType_Centered EquationDelimiterShapeType = 0 

// The divider is altered to exactly match their contents' height.
EquationDelimiterShapeType_Match EquationDelimiterShapeType = 1 
)

func Int32ToEquationDelimiterShapeType(value int32)(EquationDelimiterShapeType ,error){
	switch value {
		case 0:  return EquationDelimiterShapeType_Centered, nil  
		case 1:  return EquationDelimiterShapeType_Match, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationDelimiterShapeType value: %d", value)
	}
}

/**************Enum EquationFractionType *****************/

// This specifies the display style of the fraction bar.
type EquationFractionType int32

const(
// This specifies that the numerator is above and the denominator below is separated by a bar in the middle.
EquationFractionType_Bar EquationFractionType = 0 

// This specifies that the numerator is above and the denominator below is not separated by a bar in the middle.
EquationFractionType_NoBar EquationFractionType = 1 

// This specifies that the numerator is on the left and the denominator is on the right, separated by a '/' in between.
EquationFractionType_Linear EquationFractionType = 2 

// This specifies that the numerator is on the upper left and the denominator is on the lower right, separated by a "/".
EquationFractionType_Skewed EquationFractionType = 3 
)

func Int32ToEquationFractionType(value int32)(EquationFractionType ,error){
	switch value {
		case 0:  return EquationFractionType_Bar, nil  
		case 1:  return EquationFractionType_NoBar, nil  
		case 2:  return EquationFractionType_Linear, nil  
		case 3:  return EquationFractionType_Skewed, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationFractionType value: %d", value)
	}
}

/**************Enum EquationHorizontalJustificationType *****************/

// This specifies the default horizontal justification of equations in the document.
type EquationHorizontalJustificationType int32

const(
// Centered
EquationHorizontalJustificationType_Center EquationHorizontalJustificationType = 0 

// Centered as Group
EquationHorizontalJustificationType_CenterGroup EquationHorizontalJustificationType = 1 

// Left Justified
EquationHorizontalJustificationType_Left EquationHorizontalJustificationType = 2 

// Right Justified
EquationHorizontalJustificationType_Right EquationHorizontalJustificationType = 3 
)

func Int32ToEquationHorizontalJustificationType(value int32)(EquationHorizontalJustificationType ,error){
	switch value {
		case 0:  return EquationHorizontalJustificationType_Center, nil  
		case 1:  return EquationHorizontalJustificationType_CenterGroup, nil  
		case 2:  return EquationHorizontalJustificationType_Left, nil  
		case 3:  return EquationHorizontalJustificationType_Right, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationHorizontalJustificationType value: %d", value)
	}
}

/**************Enum EquationLimitLocationType *****************/

// Specifies the limit location on an operator.
type EquationLimitLocationType int32

const(
// Specifies that the limit is centered above or below the operator.
EquationLimitLocationType_UndOvr EquationLimitLocationType = 0 

// Specifies that the limit is on the right side of the operator.
EquationLimitLocationType_SubSup EquationLimitLocationType = 1 
)

func Int32ToEquationLimitLocationType(value int32)(EquationLimitLocationType ,error){
	switch value {
		case 0:  return EquationLimitLocationType_UndOvr, nil  
		case 1:  return EquationLimitLocationType_SubSup, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationLimitLocationType value: %d", value)
	}
}

/**************Enum EquationMathematicalOperatorType *****************/

// Mathematical Operators Type
type EquationMathematicalOperatorType int32

const(
// Use unknown type when not found in existing type.
EquationMathematicalOperatorType_Unknown EquationMathematicalOperatorType = -1 

// "∀" Unicode:\u2200
EquationMathematicalOperatorType_ForAll EquationMathematicalOperatorType = 0 

// "∁" Unicode:\u2201
EquationMathematicalOperatorType_Complement EquationMathematicalOperatorType = 1 

// "∂" Unicode:\u2202
EquationMathematicalOperatorType_PartialDifferential EquationMathematicalOperatorType = 2 

// "∃" Unicode:\u2203
EquationMathematicalOperatorType_Exists EquationMathematicalOperatorType = 3 

// "∄" Unicode:\u2204
EquationMathematicalOperatorType_NotExists EquationMathematicalOperatorType = 4 

// "∅" Unicode:\u2205
EquationMathematicalOperatorType_EmptySet EquationMathematicalOperatorType = 5 

// "∆" Unicode:\u2206
EquationMathematicalOperatorType_Increment EquationMathematicalOperatorType = 6 

// "∇" Unicode:\u2207
EquationMathematicalOperatorType_Nabla EquationMathematicalOperatorType = 7 

// "∈" Unicode:\u2208
EquationMathematicalOperatorType_ElementOf EquationMathematicalOperatorType = 8 

// "∉" Unicode:\u2209
EquationMathematicalOperatorType_NotAnElementOf EquationMathematicalOperatorType = 9 

// "∊" Unicode:\u220a
EquationMathematicalOperatorType_SmallElementOf EquationMathematicalOperatorType = 10 

// "∋" Unicode:\u220b
EquationMathematicalOperatorType_Contain EquationMathematicalOperatorType = 11 

// "∌" Unicode:\u220c
EquationMathematicalOperatorType_NotContain EquationMathematicalOperatorType = 12 

// "∍" Unicode:\u220d
EquationMathematicalOperatorType_SmallContain EquationMathematicalOperatorType = 13 

// "∎" Unicode:\u220e
EquationMathematicalOperatorType_EndOfProof EquationMathematicalOperatorType = 14 

// "∏" Unicode:\u220f
EquationMathematicalOperatorType_NaryProduct EquationMathematicalOperatorType = 15 

// "∐" Unicode:\u2210
EquationMathematicalOperatorType_NaryCoproduct EquationMathematicalOperatorType = 16 

// "∑" Unicode:\u2211
EquationMathematicalOperatorType_NarySummation EquationMathematicalOperatorType = 17 

// "∧" Unicode:\u2227
EquationMathematicalOperatorType_LogicalAnd EquationMathematicalOperatorType = 18 

// "∨" Unicode:\u2228
EquationMathematicalOperatorType_LogicalOr EquationMathematicalOperatorType = 19 

// "∩" Unicode:\u2229
EquationMathematicalOperatorType_Intersection EquationMathematicalOperatorType = 20 

// "∪" Unicode:\u222a
EquationMathematicalOperatorType_Union EquationMathematicalOperatorType = 21 

// "∫" Unicode:\u222b
EquationMathematicalOperatorType_Integral EquationMathematicalOperatorType = 22 

// "∬" Unicode:\u222c
EquationMathematicalOperatorType_DoubleIntegral EquationMathematicalOperatorType = 23 

// "∭" Unicode:\u222d
EquationMathematicalOperatorType_TripleIntegral EquationMathematicalOperatorType = 24 

// "∮" Unicode:\u222e
EquationMathematicalOperatorType_ContourIntegral EquationMathematicalOperatorType = 25 

// "∯" Unicode:\u222f
EquationMathematicalOperatorType_SurfaceIntegral EquationMathematicalOperatorType = 26 

// "∰" Unicode:\u2230
EquationMathematicalOperatorType_VolumeIntegral EquationMathematicalOperatorType = 27 

// "∱" Unicode:\u2231
EquationMathematicalOperatorType_Clockwise EquationMathematicalOperatorType = 28 

// "∲" Unicode:\u2232
EquationMathematicalOperatorType_ClockwiseContourIntegral EquationMathematicalOperatorType = 29 

// "∳" Unicode:\u2233
EquationMathematicalOperatorType_AnticlockwiseContourIntegral EquationMathematicalOperatorType = 30 

// "⋀" Unicode:\u22c0
EquationMathematicalOperatorType_NaryLogicalAnd EquationMathematicalOperatorType = 31 

// "⋁" Unicode:\u22c1
EquationMathematicalOperatorType_NaryLogicalOr EquationMathematicalOperatorType = 32 

// "⋂" Unicode:\u22c2
EquationMathematicalOperatorType_NaryIntersection EquationMathematicalOperatorType = 33 

// "⋃" Unicode:\u22c3
EquationMathematicalOperatorType_NaryUnion EquationMathematicalOperatorType = 34 
)

func Int32ToEquationMathematicalOperatorType(value int32)(EquationMathematicalOperatorType ,error){
	switch value {
		case -1:  return EquationMathematicalOperatorType_Unknown, nil  
		case 0:  return EquationMathematicalOperatorType_ForAll, nil  
		case 1:  return EquationMathematicalOperatorType_Complement, nil  
		case 2:  return EquationMathematicalOperatorType_PartialDifferential, nil  
		case 3:  return EquationMathematicalOperatorType_Exists, nil  
		case 4:  return EquationMathematicalOperatorType_NotExists, nil  
		case 5:  return EquationMathematicalOperatorType_EmptySet, nil  
		case 6:  return EquationMathematicalOperatorType_Increment, nil  
		case 7:  return EquationMathematicalOperatorType_Nabla, nil  
		case 8:  return EquationMathematicalOperatorType_ElementOf, nil  
		case 9:  return EquationMathematicalOperatorType_NotAnElementOf, nil  
		case 10:  return EquationMathematicalOperatorType_SmallElementOf, nil  
		case 11:  return EquationMathematicalOperatorType_Contain, nil  
		case 12:  return EquationMathematicalOperatorType_NotContain, nil  
		case 13:  return EquationMathematicalOperatorType_SmallContain, nil  
		case 14:  return EquationMathematicalOperatorType_EndOfProof, nil  
		case 15:  return EquationMathematicalOperatorType_NaryProduct, nil  
		case 16:  return EquationMathematicalOperatorType_NaryCoproduct, nil  
		case 17:  return EquationMathematicalOperatorType_NarySummation, nil  
		case 18:  return EquationMathematicalOperatorType_LogicalAnd, nil  
		case 19:  return EquationMathematicalOperatorType_LogicalOr, nil  
		case 20:  return EquationMathematicalOperatorType_Intersection, nil  
		case 21:  return EquationMathematicalOperatorType_Union, nil  
		case 22:  return EquationMathematicalOperatorType_Integral, nil  
		case 23:  return EquationMathematicalOperatorType_DoubleIntegral, nil  
		case 24:  return EquationMathematicalOperatorType_TripleIntegral, nil  
		case 25:  return EquationMathematicalOperatorType_ContourIntegral, nil  
		case 26:  return EquationMathematicalOperatorType_SurfaceIntegral, nil  
		case 27:  return EquationMathematicalOperatorType_VolumeIntegral, nil  
		case 28:  return EquationMathematicalOperatorType_Clockwise, nil  
		case 29:  return EquationMathematicalOperatorType_ClockwiseContourIntegral, nil  
		case 30:  return EquationMathematicalOperatorType_AnticlockwiseContourIntegral, nil  
		case 31:  return EquationMathematicalOperatorType_NaryLogicalAnd, nil  
		case 32:  return EquationMathematicalOperatorType_NaryLogicalOr, nil  
		case 33:  return EquationMathematicalOperatorType_NaryIntersection, nil  
		case 34:  return EquationMathematicalOperatorType_NaryUnion, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationMathematicalOperatorType value: %d", value)
	}
}

/**************Enum EquationNodeType *****************/

// Equation node type.
// Notice:
// (1)[1-99] Currently there is only one node in the scope, and its enumeration value is 1. The node it specifies is used to store mathematical text.
// (2)[100-199] Indicates that the node is a component of some special function nodes.
// (3)[200-] Indicates that the node has some special functions.
type EquationNodeType int32

const(
// UnKnow
EquationNodeType_UnKnow EquationNodeType = 0 

// specifies a node that stores math text
EquationNodeType_Text EquationNodeType = 1 

// Specifies a Base component
EquationNodeType_Base EquationNodeType = 100 

// Specifies a Denominator component
EquationNodeType_Denominator EquationNodeType = 101 

// Specifies a Numerator component
EquationNodeType_Numerator EquationNodeType = 102 

// Specifies a FunctionName component
EquationNodeType_FunctionName EquationNodeType = 103 

// Specifies a Subscript component
EquationNodeType_Subscript EquationNodeType = 104 

// Specifies a Superscript component
EquationNodeType_Superscript EquationNodeType = 105 

// Specifies a Degree component
EquationNodeType_Degree EquationNodeType = 106 

// Specifies a MatrixRow component.A single row of the matrix
EquationNodeType_MatrixRow EquationNodeType = 107 

// Represents a sub-object of Lower-Limit function or Upper-Limit function
EquationNodeType_Limit EquationNodeType = 108 

// Specifies a mathematical paragraph(oMathPara).
EquationNodeType_EquationParagraph EquationNodeType = 200 

// Specifies the Lower-Limit function
EquationNodeType_LowerLimit EquationNodeType = 217 

// Specifies the Upper-Limit function
EquationNodeType_UpperLimit EquationNodeType = 218 

// Specifies an equation or mathematical expression(OMath).
EquationNodeType_Mathematical EquationNodeType = 201 

// Specifies fractional equation
EquationNodeType_Fraction EquationNodeType = 202 

// Specifies function equation
EquationNodeType_Function EquationNodeType = 203 

// Specifies delimiter equation
EquationNodeType_Delimiter EquationNodeType = 204 

// Specifies n-ary operator equation
EquationNodeType_Nary EquationNodeType = 205 

// Specifies the radical equation
EquationNodeType_Radical EquationNodeType = 206 

// Specifies superscript equation
EquationNodeType_Sup EquationNodeType = 207 

// Specifies subscript equation
EquationNodeType_Sub EquationNodeType = 208 

// Specifies an equation with superscripts and subscripts to the right of the operands.
EquationNodeType_SubSup EquationNodeType = 209 

// Specifies an equation with superscripts and subscripts to the left of the operands.
EquationNodeType_PreSubSup EquationNodeType = 210 

// Specifies accent equation
EquationNodeType_Accent EquationNodeType = 211 

// Specifies bar equation
EquationNodeType_Bar EquationNodeType = 212 

// Specifies border box equation
EquationNodeType_BorderBox EquationNodeType = 213 

// Specifies box equation
EquationNodeType_Box EquationNodeType = 214 

// Specifies Group-Character equation
EquationNodeType_GroupChr EquationNodeType = 215 

// Specifies the Matrix equation,
EquationNodeType_Matrix EquationNodeType = 216 

// Specifies the Equation-Array function. The function consists of one or more equations.
EquationNodeType_ArrayEquation EquationNodeType = 317 
)

func Int32ToEquationNodeType(value int32)(EquationNodeType ,error){
	switch value {
		case 0:  return EquationNodeType_UnKnow, nil  
		case 1:  return EquationNodeType_Text, nil  
		case 100:  return EquationNodeType_Base, nil  
		case 101:  return EquationNodeType_Denominator, nil  
		case 102:  return EquationNodeType_Numerator, nil  
		case 103:  return EquationNodeType_FunctionName, nil  
		case 104:  return EquationNodeType_Subscript, nil  
		case 105:  return EquationNodeType_Superscript, nil  
		case 106:  return EquationNodeType_Degree, nil  
		case 107:  return EquationNodeType_MatrixRow, nil  
		case 108:  return EquationNodeType_Limit, nil  
		case 200:  return EquationNodeType_EquationParagraph, nil  
		case 217:  return EquationNodeType_LowerLimit, nil  
		case 218:  return EquationNodeType_UpperLimit, nil  
		case 201:  return EquationNodeType_Mathematical, nil  
		case 202:  return EquationNodeType_Fraction, nil  
		case 203:  return EquationNodeType_Function, nil  
		case 204:  return EquationNodeType_Delimiter, nil  
		case 205:  return EquationNodeType_Nary, nil  
		case 206:  return EquationNodeType_Radical, nil  
		case 207:  return EquationNodeType_Sup, nil  
		case 208:  return EquationNodeType_Sub, nil  
		case 209:  return EquationNodeType_SubSup, nil  
		case 210:  return EquationNodeType_PreSubSup, nil  
		case 211:  return EquationNodeType_Accent, nil  
		case 212:  return EquationNodeType_Bar, nil  
		case 213:  return EquationNodeType_BorderBox, nil  
		case 214:  return EquationNodeType_Box, nil  
		case 215:  return EquationNodeType_GroupChr, nil  
		case 216:  return EquationNodeType_Matrix, nil  
		case 317:  return EquationNodeType_ArrayEquation, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationNodeType value: %d", value)
	}
}

/**************Enum EquationVerticalJustificationType *****************/

// This specifies the default vertical justification of equations in the document.
type EquationVerticalJustificationType int32

const(
// top
EquationVerticalJustificationType_Top EquationVerticalJustificationType = 0 

// center
EquationVerticalJustificationType_Center EquationVerticalJustificationType = 1 

// bottom
EquationVerticalJustificationType_Bottom EquationVerticalJustificationType = 2 
)

func Int32ToEquationVerticalJustificationType(value int32)(EquationVerticalJustificationType ,error){
	switch value {
		case 0:  return EquationVerticalJustificationType_Top, nil  
		case 1:  return EquationVerticalJustificationType_Center, nil  
		case 2:  return EquationVerticalJustificationType_Bottom, nil  
		default:
			return 0 ,fmt.Errorf("invalid EquationVerticalJustificationType value: %d", value)
	}
}
// Class AccentEquationNode 

// This class specifies an accent equation, consisting of a base component and a combining diacritic.
type AccentEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewAccentEquationNode(src *EquationNode) ( *AccentEquationNode, error) {
	accentequationnode := &AccentEquationNode{}
	CGoReturnPtr := C.New_AccentEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		accentequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(accentequationnode, DeleteAccentEquationNode)
		return accentequationnode, nil
	} else {
		accentequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return accentequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *AccentEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.AccentEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This attribute specifies the type of combining diacritical mark attached to the base of the accent function. The default accent character is U+0302.
// It is strongly recommended to use attribute AccentType to set accent character.
// Use this property setting if you cannot find the character you need in a known type.
// Returns:
//   string  
func (instance *AccentEquationNode) GetAccentCharacter()  (string,  error)  {
	CGoReturnPtr := C.AccentEquationNode_GetAccentCharacter( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// This attribute specifies the type of combining diacritical mark attached to the base of the accent function. The default accent character is U+0302.
// It is strongly recommended to use attribute AccentType to set accent character.
// Use this property setting if you cannot find the character you need in a known type.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *AccentEquationNode) SetAccentCharacter(value string)  error {
	CGoReturnPtr := C.AccentEquationNode_SetAccentCharacter( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specify combining characters by type value.
// Returns:
//   int32  
func (instance *AccentEquationNode) GetAccentCharacterType()  (EquationCombiningCharacterType,  error)  {
	CGoReturnPtr := C.AccentEquationNode_GetAccentCharacterType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationCombiningCharacterType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specify combining characters by type value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *AccentEquationNode) SetAccentCharacterType(value EquationCombiningCharacterType)  error {
	CGoReturnPtr := C.AccentEquationNode_SetAccentCharacterType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *AccentEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.AccentEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *AccentEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *AccentEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.AccentEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *AccentEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.AccentEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *AccentEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.AccentEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *AccentEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *AccentEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.AccentEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *AccentEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *AccentEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *AccentEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *AccentEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *AccentEquationNode) Remove()  error {
	CGoReturnPtr := C.AccentEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *AccentEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.AccentEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *AccentEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.AccentEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *AccentEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.AccentEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *AccentEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.AccentEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *AccentEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.AccentEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func AccentEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.AccentEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteAccentEquationNode(accentequationnode *AccentEquationNode){
	runtime.SetFinalizer(accentequationnode, nil)
	C.Delete_AccentEquationNode(accentequationnode.ptr)
	accentequationnode.ptr = nil
}

// Class ArrayEquationNode 

// Specifies the Equation-Array function, an object consisting of one or more equations.
type ArrayEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewArrayEquationNode(src *EquationNode) ( *ArrayEquationNode, error) {
	arrayequationnode := &ArrayEquationNode{}
	CGoReturnPtr := C.New_ArrayEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		arrayequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(arrayequationnode, DeleteArrayEquationNode)
		return arrayequationnode, nil
	} else {
		arrayequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return arrayequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ArrayEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *ArrayEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *ArrayEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *ArrayEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.ArrayEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *ArrayEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *ArrayEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *ArrayEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *ArrayEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.ArrayEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *ArrayEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *ArrayEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *ArrayEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *ArrayEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *ArrayEquationNode) Remove()  error {
	CGoReturnPtr := C.ArrayEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *ArrayEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.ArrayEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *ArrayEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.ArrayEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *ArrayEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.ArrayEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *ArrayEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *ArrayEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func ArrayEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.ArrayEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteArrayEquationNode(arrayequationnode *ArrayEquationNode){
	runtime.SetFinalizer(arrayequationnode, nil)
	C.Delete_ArrayEquationNode(arrayequationnode.ptr)
	arrayequationnode.ptr = nil
}

// Class BarEquationNode 

// This class specifies the bar equation, consisting of a base argument and an overbar or underbar.
type BarEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewBarEquationNode(src *EquationNode) ( *BarEquationNode, error) {
	barequationnode := &BarEquationNode{}
	CGoReturnPtr := C.New_BarEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		barequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(barequationnode, DeleteBarEquationNode)
		return barequationnode, nil
	} else {
		barequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return barequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *BarEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.BarEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This attribute specifies the position of the bar in the bar object
// Returns:
//   int32  
func (instance *BarEquationNode) GetBarPosition()  (EquationCharacterPositionType,  error)  {
	CGoReturnPtr := C.BarEquationNode_GetBarPosition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationCharacterPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This attribute specifies the position of the bar in the bar object
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *BarEquationNode) SetBarPosition(value EquationCharacterPositionType)  error {
	CGoReturnPtr := C.BarEquationNode_SetBarPosition( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *BarEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.BarEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *BarEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *BarEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.BarEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *BarEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.BarEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *BarEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.BarEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BarEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *BarEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.BarEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BarEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BarEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BarEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *BarEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *BarEquationNode) Remove()  error {
	CGoReturnPtr := C.BarEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *BarEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.BarEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *BarEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.BarEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *BarEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.BarEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *BarEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.BarEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *BarEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.BarEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func BarEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BarEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteBarEquationNode(barequationnode *BarEquationNode){
	runtime.SetFinalizer(barequationnode, nil)
	C.Delete_BarEquationNode(barequationnode.ptr)
	barequationnode.ptr = nil
}

// Class BorderBoxEquationNode 

// This class specifies the Border Box function, consisting of a border drawn around an equation.
type BorderBoxEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewBorderBoxEquationNode(src *EquationNode) ( *BorderBoxEquationNode, error) {
	borderboxequationnode := &BorderBoxEquationNode{}
	CGoReturnPtr := C.New_BorderBoxEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		borderboxequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(borderboxequationnode, DeleteBorderBoxEquationNode)
		return borderboxequationnode, nil
	} else {
		borderboxequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return borderboxequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *BorderBoxEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *BorderBoxEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *BorderBoxEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *BorderBoxEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.BorderBoxEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *BorderBoxEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *BorderBoxEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BorderBoxEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *BorderBoxEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.BorderBoxEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BorderBoxEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BorderBoxEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BorderBoxEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *BorderBoxEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *BorderBoxEquationNode) Remove()  error {
	CGoReturnPtr := C.BorderBoxEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *BorderBoxEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.BorderBoxEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *BorderBoxEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.BorderBoxEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *BorderBoxEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.BorderBoxEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *BorderBoxEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *BorderBoxEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func BorderBoxEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BorderBoxEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteBorderBoxEquationNode(borderboxequationnode *BorderBoxEquationNode){
	runtime.SetFinalizer(borderboxequationnode, nil)
	C.Delete_BorderBoxEquationNode(borderboxequationnode.ptr)
	borderboxequationnode.ptr = nil
}

// Class BoxEquationNode 

// This class specifies the box function, which is used to group components of an equation.
type BoxEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewBoxEquationNode(src *EquationNode) ( *BoxEquationNode, error) {
	boxequationnode := &BoxEquationNode{}
	CGoReturnPtr := C.New_BoxEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		boxequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(boxequationnode, DeleteBoxEquationNode)
		return boxequationnode, nil
	} else {
		boxequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return boxequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *BoxEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.BoxEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *BoxEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.BoxEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *BoxEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *BoxEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.BoxEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *BoxEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.BoxEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *BoxEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.BoxEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BoxEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *BoxEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.BoxEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BoxEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BoxEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *BoxEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *BoxEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *BoxEquationNode) Remove()  error {
	CGoReturnPtr := C.BoxEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *BoxEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.BoxEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *BoxEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.BoxEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *BoxEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.BoxEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *BoxEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.BoxEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *BoxEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.BoxEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func BoxEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.BoxEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteBoxEquationNode(boxequationnode *BoxEquationNode){
	runtime.SetFinalizer(boxequationnode, nil)
	C.Delete_BoxEquationNode(boxequationnode.ptr)
	boxequationnode.ptr = nil
}

// Class DelimiterEquationNode 

// This class specifies the delimiter equation, consisting of opening and closing delimiters (such as parentheses, braces, brackets, and vertical bars), and a component contained inside.
// The delimiter may have more than one component, with a designated separator character between each component.
type DelimiterEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewDelimiterEquationNode(src *EquationNode) ( *DelimiterEquationNode, error) {
	delimiterequationnode := &DelimiterEquationNode{}
	CGoReturnPtr := C.New_DelimiterEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		delimiterequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(delimiterequationnode, DeleteDelimiterEquationNode)
		return delimiterequationnode, nil
	} else {
		delimiterequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return delimiterequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DelimiterEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Delimiter beginning character.
// Returns:
//   string  
func (instance *DelimiterEquationNode) GetBeginChar()  (string,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetBeginChar( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Delimiter beginning character.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DelimiterEquationNode) SetBeginChar(value string)  error {
	CGoReturnPtr := C.DelimiterEquationNode_SetBeginChar( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Delimiter ending character.
// Returns:
//   string  
func (instance *DelimiterEquationNode) GetEndChar()  (string,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetEndChar( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Delimiter ending character.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DelimiterEquationNode) SetEndChar(value string)  error {
	CGoReturnPtr := C.DelimiterEquationNode_SetEndChar( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns:
//   bool  
func (instance *DelimiterEquationNode) GetNaryGrow()  (bool,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetNaryGrow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *DelimiterEquationNode) SetNaryGrow(value bool)  error {
	CGoReturnPtr := C.DelimiterEquationNode_SetNaryGrow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Delimiter separator character.
// Returns:
//   string  
func (instance *DelimiterEquationNode) GetSeparatorChar()  (string,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetSeparatorChar( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Delimiter separator character.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *DelimiterEquationNode) SetSeparatorChar(value string)  error {
	CGoReturnPtr := C.DelimiterEquationNode_SetSeparatorChar( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the shape of delimiters in the delimiter object.
// Returns:
//   int32  
func (instance *DelimiterEquationNode) GetDelimiterShape()  (EquationDelimiterShapeType,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetDelimiterShape( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationDelimiterShapeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the shape of delimiters in the delimiter object.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *DelimiterEquationNode) SetDelimiterShape(value EquationDelimiterShapeType)  error {
	CGoReturnPtr := C.DelimiterEquationNode_SetDelimiterShape( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *DelimiterEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *DelimiterEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *DelimiterEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.DelimiterEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *DelimiterEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *DelimiterEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *DelimiterEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *DelimiterEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.DelimiterEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *DelimiterEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *DelimiterEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *DelimiterEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *DelimiterEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *DelimiterEquationNode) Remove()  error {
	CGoReturnPtr := C.DelimiterEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *DelimiterEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.DelimiterEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *DelimiterEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.DelimiterEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *DelimiterEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.DelimiterEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *DelimiterEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *DelimiterEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func DelimiterEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.DelimiterEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteDelimiterEquationNode(delimiterequationnode *DelimiterEquationNode){
	runtime.SetFinalizer(delimiterequationnode, nil)
	C.Delete_DelimiterEquationNode(delimiterequationnode.ptr)
	delimiterequationnode.ptr = nil
}

// Class EquationComponentNode 

// This class specifies the components of an equation or mathematical expression.
// Different types of components combined into different equations.
// For example, a fraction consists of two parts, a numerator component and a denominator component.
// For more component types, please refer to 'EquationNodeType'.
type EquationComponentNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewEquationComponentNode(src *EquationNode) ( *EquationComponentNode, error) {
	equationcomponentnode := &EquationComponentNode{}
	CGoReturnPtr := C.New_EquationComponentNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		equationcomponentnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(equationcomponentnode, DeleteEquationComponentNode)
		return equationcomponentnode, nil
	} else {
		equationcomponentnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return equationcomponentnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *EquationComponentNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.EquationComponentNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *EquationComponentNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.EquationComponentNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *EquationComponentNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *EquationComponentNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.EquationComponentNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *EquationComponentNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.EquationComponentNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *EquationComponentNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.EquationComponentNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationComponentNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *EquationComponentNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.EquationComponentNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationComponentNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationComponentNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationComponentNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *EquationComponentNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *EquationComponentNode) Remove()  error {
	CGoReturnPtr := C.EquationComponentNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *EquationComponentNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.EquationComponentNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *EquationComponentNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.EquationComponentNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *EquationComponentNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.EquationComponentNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *EquationComponentNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.EquationComponentNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *EquationComponentNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.EquationComponentNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func EquationComponentNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationComponentNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteEquationComponentNode(equationcomponentnode *EquationComponentNode){
	runtime.SetFinalizer(equationcomponentnode, nil)
	C.Delete_EquationComponentNode(equationcomponentnode.ptr)
	equationcomponentnode.ptr = nil
}

// Class EquationNode 

// Abstract class for deriving other equation nodes.
type EquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - FontSetting 
func NewEquationNode(src *FontSetting) ( *EquationNode, error) {
	equationnode := &EquationNode{}
	CGoReturnPtr := C.New_EquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		equationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(equationnode, DeleteEquationNode)
		return equationnode, nil
	} else {
		equationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return equationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *EquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.EquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *EquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *EquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.EquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *EquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.EquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *EquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.EquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *EquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.EquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *EquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *EquationNode) Remove()  error {
	CGoReturnPtr := C.EquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *EquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.EquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *EquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.EquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *EquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.EquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *EquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.EquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *EquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.EquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *EquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.EquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func EquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Gets the start index of the characters.
// Returns:
//   int32  
func (instance *EquationNode) GetStartIndex()  (int32,  error)  {
	CGoReturnPtr := C.EquationNode_GetStartIndex( instance.ptr)
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
func (instance *EquationNode) GetLength()  (int32,  error)  {
	CGoReturnPtr := C.EquationNode_GetLength( instance.ptr)
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
func (instance *EquationNode) GetFont()  (*Font,  error)  {
	CGoReturnPtr := C.EquationNode_GetFont( instance.ptr)
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
func (instance *EquationNode) SetWordArtStyle(style PresetWordArtStyle)  error {
	CGoReturnPtr := C.EquationNode_SetWordArtStyle( instance.ptr, C.int( int32(style)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Returns the text options.
// Returns:
//   TextOptions  
func (instance *EquationNode) GetTextOptions()  (*TextOptions,  error)  {
	CGoReturnPtr := C.EquationNode_GetTextOptions( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &TextOptions{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteTextOptions) 

	return result, nil 
}


func DeleteEquationNode(equationnode *EquationNode){
	runtime.SetFinalizer(equationnode, nil)
	C.Delete_EquationNode(equationnode.ptr)
	equationnode.ptr = nil
}

// Class EquationNodeParagraph 

// This class specifies a mathematical paragraph containing one or more MathEquationNode(OMath) elements.
type EquationNodeParagraph struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewEquationNodeParagraph(src *EquationNode) ( *EquationNodeParagraph, error) {
	equationnodeparagraph := &EquationNodeParagraph{}
	CGoReturnPtr := C.New_EquationNodeParagraph(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		equationnodeparagraph.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(equationnodeparagraph, DeleteEquationNodeParagraph)
		return equationnodeparagraph, nil
	} else {
		equationnodeparagraph.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return equationnodeparagraph, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *EquationNodeParagraph) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This specifies justification of the math paragraph (a series of adjacent equations within the same paragraph). A math paragraph can be Left Justified, Right Justified, Centered, or Centered as Group. By default, the math paragraph is Centered as Group. This means that the equations can be aligned with respect to each other, but the entire group of equations is centered as a whole.
// Returns:
//   int32  
func (instance *EquationNodeParagraph) GetJustification()  (EquationHorizontalJustificationType,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_GetJustification( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationHorizontalJustificationType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This specifies justification of the math paragraph (a series of adjacent equations within the same paragraph). A math paragraph can be Left Justified, Right Justified, Centered, or Centered as Group. By default, the math paragraph is Centered as Group. This means that the equations can be aligned with respect to each other, but the entire group of equations is centered as a whole.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *EquationNodeParagraph) SetJustification(value EquationHorizontalJustificationType)  error {
	CGoReturnPtr := C.EquationNodeParagraph_SetJustification( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *EquationNodeParagraph) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *EquationNodeParagraph) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *EquationNodeParagraph) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.EquationNodeParagraph_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *EquationNodeParagraph) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *EquationNodeParagraph) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNodeParagraph) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *EquationNodeParagraph) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.EquationNodeParagraph_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNodeParagraph) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNodeParagraph) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *EquationNodeParagraph) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *EquationNodeParagraph) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *EquationNodeParagraph) Remove()  error {
	CGoReturnPtr := C.EquationNodeParagraph_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *EquationNodeParagraph) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.EquationNodeParagraph_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *EquationNodeParagraph) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.EquationNodeParagraph_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *EquationNodeParagraph) RemoveAllChildren()  error {
	CGoReturnPtr := C.EquationNodeParagraph_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *EquationNodeParagraph) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *EquationNodeParagraph) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func EquationNodeParagraph_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.EquationNodeParagraph_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteEquationNodeParagraph(equationnodeparagraph *EquationNodeParagraph){
	runtime.SetFinalizer(equationnodeparagraph, nil)
	C.Delete_EquationNodeParagraph(equationnodeparagraph.ptr)
	equationnodeparagraph.ptr = nil
}

// Class FractionEquationNode 

// This class  specifies the fraction equation, consisting of a numerator and denominator separated by a fraction bar. The fraction bar can be horizontal or diagonal, depending on the fraction properties. The fraction equation is also used to represent the stack function, which places one element above another, with no fraction bar.
type FractionEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewFractionEquationNode(src *EquationNode) ( *FractionEquationNode, error) {
	fractionequationnode := &FractionEquationNode{}
	CGoReturnPtr := C.New_FractionEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		fractionequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(fractionequationnode, DeleteFractionEquationNode)
		return fractionequationnode, nil
	} else {
		fractionequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return fractionequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *FractionEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.FractionEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This specifies the type of fraction ; the default is 'Bar'.
// Returns:
//   int32  
func (instance *FractionEquationNode) GetFractionType()  (EquationFractionType,  error)  {
	CGoReturnPtr := C.FractionEquationNode_GetFractionType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationFractionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This specifies the type of fraction ; the default is 'Bar'.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *FractionEquationNode) SetFractionType(value EquationFractionType)  error {
	CGoReturnPtr := C.FractionEquationNode_SetFractionType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *FractionEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.FractionEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *FractionEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *FractionEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.FractionEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *FractionEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.FractionEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *FractionEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.FractionEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FractionEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *FractionEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.FractionEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FractionEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FractionEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FractionEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *FractionEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *FractionEquationNode) Remove()  error {
	CGoReturnPtr := C.FractionEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *FractionEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.FractionEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *FractionEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.FractionEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *FractionEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.FractionEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *FractionEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.FractionEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *FractionEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.FractionEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func FractionEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FractionEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteFractionEquationNode(fractionequationnode *FractionEquationNode){
	runtime.SetFinalizer(fractionequationnode, nil)
	C.Delete_FractionEquationNode(fractionequationnode.ptr)
	fractionequationnode.ptr = nil
}

// Class FunctionEquationNode 

// This class specifies the Function-Apply equation, which consists of a function name and an argument acted upon.
// The types of the name and argument components are 'EquationNodeType.FunctionName' and 'EquationNodeType.Base' respectively.
type FunctionEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewFunctionEquationNode(src *EquationNode) ( *FunctionEquationNode, error) {
	functionequationnode := &FunctionEquationNode{}
	CGoReturnPtr := C.New_FunctionEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		functionequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(functionequationnode, DeleteFunctionEquationNode)
		return functionequationnode, nil
	} else {
		functionequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return functionequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *FunctionEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *FunctionEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *FunctionEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *FunctionEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.FunctionEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *FunctionEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *FunctionEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FunctionEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *FunctionEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.FunctionEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FunctionEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FunctionEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *FunctionEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *FunctionEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *FunctionEquationNode) Remove()  error {
	CGoReturnPtr := C.FunctionEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *FunctionEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.FunctionEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *FunctionEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.FunctionEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *FunctionEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.FunctionEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *FunctionEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *FunctionEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func FunctionEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.FunctionEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteFunctionEquationNode(functionequationnode *FunctionEquationNode){
	runtime.SetFinalizer(functionequationnode, nil)
	C.Delete_FunctionEquationNode(functionequationnode.ptr)
	functionequationnode.ptr = nil
}

// Class GroupCharacterEquationNode 

// This class specifies the Group-Character function, consisting of a character drawn above or below text, often with the purpose of visually grouping items.
type GroupCharacterEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewGroupCharacterEquationNode(src *EquationNode) ( *GroupCharacterEquationNode, error) {
	groupcharacterequationnode := &GroupCharacterEquationNode{}
	CGoReturnPtr := C.New_GroupCharacterEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		groupcharacterequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(groupcharacterequationnode, DeleteGroupCharacterEquationNode)
		return groupcharacterequationnode, nil
	} else {
		groupcharacterequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return groupcharacterequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *GroupCharacterEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies a symbol(default U+23DF).
// It is strongly recommended to use attribute ChrType to set accent character.
// Use this property setting if you cannot find the character you need in a known type.
// Returns:
//   string  
func (instance *GroupCharacterEquationNode) GetGroupChr()  (string,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetGroupChr( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies a symbol(default U+23DF).
// It is strongly recommended to use attribute ChrType to set accent character.
// Use this property setting if you cannot find the character you need in a known type.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) SetGroupChr(value string)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_SetGroupChr( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specify combining characters by type value.
// Returns:
//   int32  
func (instance *GroupCharacterEquationNode) GetChrType()  (EquationCombiningCharacterType,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetChrType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationCombiningCharacterType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specify combining characters by type value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) SetChrType(value EquationCombiningCharacterType)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_SetChrType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This attribute specifies the position of the character in the object
// Returns:
//   int32  
func (instance *GroupCharacterEquationNode) GetPosition()  (EquationCharacterPositionType,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetPosition( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationCharacterPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This attribute specifies the position of the character in the object
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) SetPosition(value EquationCharacterPositionType)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_SetPosition( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This attribute, combined with pos of groupChrPr, specifies the vertical layout of the groupChr object. Where pos specifies the position of the grouping character, vertJc specifies the alignment of the object with respect to the baseline.
// Returns:
//   int32  
func (instance *GroupCharacterEquationNode) GetVertJc()  (EquationCharacterPositionType,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetVertJc( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationCharacterPositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This attribute, combined with pos of groupChrPr, specifies the vertical layout of the groupChr object. Where pos specifies the position of the grouping character, vertJc specifies the alignment of the object with respect to the baseline.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) SetVertJc(value EquationCharacterPositionType)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_SetVertJc( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *GroupCharacterEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *GroupCharacterEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *GroupCharacterEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *GroupCharacterEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *GroupCharacterEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *GroupCharacterEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *GroupCharacterEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *GroupCharacterEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *GroupCharacterEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) Remove()  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *GroupCharacterEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.GroupCharacterEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *GroupCharacterEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *GroupCharacterEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func GroupCharacterEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.GroupCharacterEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteGroupCharacterEquationNode(groupcharacterequationnode *GroupCharacterEquationNode){
	runtime.SetFinalizer(groupcharacterequationnode, nil)
	C.Delete_GroupCharacterEquationNode(groupcharacterequationnode.ptr)
	groupcharacterequationnode.ptr = nil
}

// Class LimLowUppEquationNode 

// This class specifies the limit function.
type LimLowUppEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewLimLowUppEquationNode(src *EquationNode) ( *LimLowUppEquationNode, error) {
	limlowuppequationnode := &LimLowUppEquationNode{}
	CGoReturnPtr := C.New_LimLowUppEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		limlowuppequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(limlowuppequationnode, DeleteLimLowUppEquationNode)
		return limlowuppequationnode, nil
	} else {
		limlowuppequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return limlowuppequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LimLowUppEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *LimLowUppEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *LimLowUppEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *LimLowUppEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.LimLowUppEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *LimLowUppEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *LimLowUppEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *LimLowUppEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *LimLowUppEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.LimLowUppEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *LimLowUppEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *LimLowUppEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *LimLowUppEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *LimLowUppEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *LimLowUppEquationNode) Remove()  error {
	CGoReturnPtr := C.LimLowUppEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *LimLowUppEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.LimLowUppEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *LimLowUppEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.LimLowUppEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *LimLowUppEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.LimLowUppEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *LimLowUppEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *LimLowUppEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func LimLowUppEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.LimLowUppEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteLimLowUppEquationNode(limlowuppequationnode *LimLowUppEquationNode){
	runtime.SetFinalizer(limlowuppequationnode, nil)
	C.Delete_LimLowUppEquationNode(limlowuppequationnode.ptr)
	limlowuppequationnode.ptr = nil
}

// Class MathematicalEquationNode 

// This class specifies an equation or mathematical expression. All mathematical text of equations or mathematical expressions are contained by this class.
type MathematicalEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewMathematicalEquationNode(src *EquationNode) ( *MathematicalEquationNode, error) {
	mathematicalequationnode := &MathematicalEquationNode{}
	CGoReturnPtr := C.New_MathematicalEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		mathematicalequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(mathematicalequationnode, DeleteMathematicalEquationNode)
		return mathematicalequationnode, nil
	} else {
		mathematicalequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return mathematicalequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *MathematicalEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *MathematicalEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *MathematicalEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *MathematicalEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.MathematicalEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *MathematicalEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *MathematicalEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MathematicalEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *MathematicalEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.MathematicalEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MathematicalEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MathematicalEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MathematicalEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *MathematicalEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *MathematicalEquationNode) Remove()  error {
	CGoReturnPtr := C.MathematicalEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *MathematicalEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.MathematicalEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *MathematicalEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.MathematicalEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *MathematicalEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.MathematicalEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *MathematicalEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *MathematicalEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func MathematicalEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MathematicalEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteMathematicalEquationNode(mathematicalequationnode *MathematicalEquationNode){
	runtime.SetFinalizer(mathematicalequationnode, nil)
	C.Delete_MathematicalEquationNode(mathematicalequationnode.ptr)
	mathematicalequationnode.ptr = nil
}

// Class MatrixEquationNode 

// This class specifies the Matrix equation, consisting of one or more elements laid out in one or more rows and one or more columns.
type MatrixEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewMatrixEquationNode(src *EquationNode) ( *MatrixEquationNode, error) {
	matrixequationnode := &MatrixEquationNode{}
	CGoReturnPtr := C.New_MatrixEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		matrixequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(matrixequationnode, DeleteMatrixEquationNode)
		return matrixequationnode, nil
	} else {
		matrixequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return matrixequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *MatrixEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This attribute specifies the justification of the matrix. Text outside of the matrix can be aligned with the bottom, top, or center of a matrix function. Default, the matrix assumes center justification.
// Returns:
//   int32  
func (instance *MatrixEquationNode) GetBaseJc()  (EquationVerticalJustificationType,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_GetBaseJc( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationVerticalJustificationType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This attribute specifies the justification of the matrix. Text outside of the matrix can be aligned with the bottom, top, or center of a matrix function. Default, the matrix assumes center justification.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *MatrixEquationNode) SetBaseJc(value EquationVerticalJustificationType)  error {
	CGoReturnPtr := C.MatrixEquationNode_SetBaseJc( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This attribute specifies the Hide Placeholders property on a matrix. When this property is on, placeholders do not appear in the matrix.Default, placeholders do appear such that the locations where text can be inserted are made visible.
// Returns:
//   bool  
func (instance *MatrixEquationNode) IsHidePlaceholder()  (bool,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_IsHidePlaceholder( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This attribute specifies the Hide Placeholders property on a matrix. When this property is on, placeholders do not appear in the matrix.Default, placeholders do appear such that the locations where text can be inserted are made visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *MatrixEquationNode) SetIsHidePlaceholder(value bool)  error {
	CGoReturnPtr := C.MatrixEquationNode_SetIsHidePlaceholder( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *MatrixEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *MatrixEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *MatrixEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.MatrixEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *MatrixEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *MatrixEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MatrixEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *MatrixEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.MatrixEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MatrixEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MatrixEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *MatrixEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *MatrixEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *MatrixEquationNode) Remove()  error {
	CGoReturnPtr := C.MatrixEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *MatrixEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.MatrixEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *MatrixEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.MatrixEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *MatrixEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.MatrixEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *MatrixEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *MatrixEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func MatrixEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.MatrixEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteMatrixEquationNode(matrixequationnode *MatrixEquationNode){
	runtime.SetFinalizer(matrixequationnode, nil)
	C.Delete_MatrixEquationNode(matrixequationnode.ptr)
	matrixequationnode.ptr = nil
}

// Class NaryEquationNode 

// This class specifies an n-ary operator equation consisting of an n-ary operator, a base (or operand), and optional upper and lower bounds.
type NaryEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewNaryEquationNode(src *EquationNode) ( *NaryEquationNode, error) {
	naryequationnode := &NaryEquationNode{}
	CGoReturnPtr := C.New_NaryEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		naryequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(naryequationnode, DeleteNaryEquationNode)
		return naryequationnode, nil
	} else {
		naryequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return naryequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *NaryEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.NaryEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Whether to display the lower bound
// Returns:
//   bool  
func (instance *NaryEquationNode) IsHideSubscript()  (bool,  error)  {
	CGoReturnPtr := C.NaryEquationNode_IsHideSubscript( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Whether to display the lower bound
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NaryEquationNode) SetIsHideSubscript(value bool)  error {
	CGoReturnPtr := C.NaryEquationNode_SetIsHideSubscript( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Whether to display the upper bound
// Returns:
//   bool  
func (instance *NaryEquationNode) IsHideSuperscript()  (bool,  error)  {
	CGoReturnPtr := C.NaryEquationNode_IsHideSuperscript( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Whether to display the upper bound
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NaryEquationNode) SetIsHideSuperscript(value bool)  error {
	CGoReturnPtr := C.NaryEquationNode_SetIsHideSuperscript( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This attribute specifies the location of limits in n-ary operators. Limits can be either centered above and below the n-ary operator, or positioned just to the right of the operator.
// Returns:
//   int32  
func (instance *NaryEquationNode) GetLimitLocation()  (EquationLimitLocationType,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetLimitLocation( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationLimitLocationType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// This attribute specifies the location of limits in n-ary operators. Limits can be either centered above and below the n-ary operator, or positioned just to the right of the operator.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *NaryEquationNode) SetLimitLocation(value EquationLimitLocationType)  error {
	CGoReturnPtr := C.NaryEquationNode_SetLimitLocation( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// an n-ary operator.e.g "∑".
// It is strongly recommended to use attribute NaryOperatorType to set n-ary operator.
// Use this property setting if you cannot find the character you need in a known type.
// Returns:
//   string  
func (instance *NaryEquationNode) GetNaryOperator()  (string,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetNaryOperator( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// an n-ary operator.e.g "∑".
// It is strongly recommended to use attribute NaryOperatorType to set n-ary operator.
// Use this property setting if you cannot find the character you need in a known type.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *NaryEquationNode) SetNaryOperator(value string)  error {
	CGoReturnPtr := C.NaryEquationNode_SetNaryOperator( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// an n-ary operator.e.g "∑"
// Returns:
//   int32  
func (instance *NaryEquationNode) GetNaryOperatorType()  (EquationMathematicalOperatorType,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetNaryOperatorType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationMathematicalOperatorType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// an n-ary operator.e.g "∑"
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *NaryEquationNode) SetNaryOperatorType(value EquationMathematicalOperatorType)  error {
	CGoReturnPtr := C.NaryEquationNode_SetNaryOperatorType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// This attribute specifies the growth property of n-ary operators at the document level. When off, n-ary operators such as integrals and summations do not grow to match the size of their operand height. When on, the n-ary operator grows vertically to match its operand height.
// Returns:
//   bool  
func (instance *NaryEquationNode) GetNaryGrow()  (bool,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetNaryGrow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// This attribute specifies the growth property of n-ary operators at the document level. When off, n-ary operators such as integrals and summations do not grow to match the size of their operand height. When on, the n-ary operator grows vertically to match its operand height.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *NaryEquationNode) SetNaryGrow(value bool)  error {
	CGoReturnPtr := C.NaryEquationNode_SetNaryGrow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *NaryEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.NaryEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *NaryEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *NaryEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.NaryEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *NaryEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.NaryEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *NaryEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.NaryEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *NaryEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *NaryEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.NaryEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *NaryEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *NaryEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *NaryEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *NaryEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *NaryEquationNode) Remove()  error {
	CGoReturnPtr := C.NaryEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *NaryEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.NaryEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *NaryEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.NaryEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *NaryEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.NaryEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *NaryEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *NaryEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.NaryEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func NaryEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.NaryEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteNaryEquationNode(naryequationnode *NaryEquationNode){
	runtime.SetFinalizer(naryequationnode, nil)
	C.Delete_NaryEquationNode(naryequationnode.ptr)
	naryequationnode.ptr = nil
}

// Class RadicalEquationNode 

// This class specifies the radical equation, consisting of an optional degree deg(EquationNodeType.Degree) and a base.
type RadicalEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewRadicalEquationNode(src *EquationNode) ( *RadicalEquationNode, error) {
	radicalequationnode := &RadicalEquationNode{}
	CGoReturnPtr := C.New_RadicalEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		radicalequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(radicalequationnode, DeleteRadicalEquationNode)
		return radicalequationnode, nil
	} else {
		radicalequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return radicalequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RadicalEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Whether to hide the degree of radicals.
// Returns:
//   bool  
func (instance *RadicalEquationNode) IsDegHide()  (bool,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_IsDegHide( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Whether to hide the degree of radicals.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadicalEquationNode) SetIsDegHide(value bool)  error {
	CGoReturnPtr := C.RadicalEquationNode_SetIsDegHide( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *RadicalEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *RadicalEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *RadicalEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.RadicalEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *RadicalEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *RadicalEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *RadicalEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *RadicalEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.RadicalEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *RadicalEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *RadicalEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *RadicalEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *RadicalEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *RadicalEquationNode) Remove()  error {
	CGoReturnPtr := C.RadicalEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *RadicalEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.RadicalEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *RadicalEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.RadicalEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *RadicalEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.RadicalEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *RadicalEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *RadicalEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func RadicalEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.RadicalEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteRadicalEquationNode(radicalequationnode *RadicalEquationNode){
	runtime.SetFinalizer(radicalequationnode, nil)
	C.Delete_RadicalEquationNode(radicalequationnode.ptr)
	radicalequationnode.ptr = nil
}

// Class SubSupEquationNode 

// This class specifies an equation that can optionally be superscript or subscript.
// There are four main forms of this equation, superscript，subscript，superscript and subscript placed to the left of the base, superscript and subscript placed to the right of the base.
type SubSupEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewSubSupEquationNode(src *EquationNode) ( *SubSupEquationNode, error) {
	subsupequationnode := &SubSupEquationNode{}
	CGoReturnPtr := C.New_SubSupEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		subsupequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(subsupequationnode, DeleteSubSupEquationNode)
		return subsupequationnode, nil
	} else {
		subsupequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return subsupequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SubSupEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *SubSupEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *SubSupEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *SubSupEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.SubSupEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *SubSupEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *SubSupEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *SubSupEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *SubSupEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.SubSupEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *SubSupEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *SubSupEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *SubSupEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *SubSupEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *SubSupEquationNode) Remove()  error {
	CGoReturnPtr := C.SubSupEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *SubSupEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.SubSupEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *SubSupEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.SubSupEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *SubSupEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.SubSupEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *SubSupEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *SubSupEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func SubSupEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.SubSupEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteSubSupEquationNode(subsupequationnode *SubSupEquationNode){
	runtime.SetFinalizer(subsupequationnode, nil)
	C.Delete_SubSupEquationNode(subsupequationnode.ptr)
	subsupequationnode.ptr = nil
}

// Class TextRunEquationNode 

// This class in the equation node is used to store the actual content(a sequence of mathematical text) of the equation.
// Usually a node object per character.
type TextRunEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewTextRunEquationNode(src *EquationNode) ( *TextRunEquationNode, error) {
	textrunequationnode := &TextRunEquationNode{}
	CGoReturnPtr := C.New_TextRunEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		textrunequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(textrunequationnode, DeleteTextRunEquationNode)
		return textrunequationnode, nil
	} else {
		textrunequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return textrunequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextRunEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Set the content of the text node(Usually a node object per character).
// Returns:
//   string  
func (instance *TextRunEquationNode) GetText()  (string,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_GetText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Set the content of the text node(Usually a node object per character).
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextRunEquationNode) SetText(value string)  error {
	CGoReturnPtr := C.TextRunEquationNode_SetText( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *TextRunEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *TextRunEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *TextRunEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.TextRunEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *TextRunEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *TextRunEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *TextRunEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *TextRunEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.TextRunEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *TextRunEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *TextRunEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *TextRunEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *TextRunEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *TextRunEquationNode) Remove()  error {
	CGoReturnPtr := C.TextRunEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *TextRunEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.TextRunEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *TextRunEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.TextRunEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *TextRunEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.TextRunEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *TextRunEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *TextRunEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func TextRunEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.TextRunEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteTextRunEquationNode(textrunequationnode *TextRunEquationNode){
	runtime.SetFinalizer(textrunequationnode, nil)
	C.Delete_TextRunEquationNode(textrunequationnode.ptr)
	textrunequationnode.ptr = nil
}

// Class UnknowEquationNode 

// Equation node class of unknown type
type UnknowEquationNode struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - EquationNode 
func NewUnknowEquationNode(src *EquationNode) ( *UnknowEquationNode, error) {
	unknowequationnode := &UnknowEquationNode{}
	CGoReturnPtr := C.New_UnknowEquationNode(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		unknowequationnode.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(unknowequationnode, DeleteUnknowEquationNode)
		return unknowequationnode, nil
	} else {
		unknowequationnode.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return unknowequationnode, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *UnknowEquationNode) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Determine whether the current equation node is equal to the specified node
// Parameters:
//   obj - Object 
// Returns:
//   bool  
func (instance *UnknowEquationNode) Equals(obj *Object)  (bool,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_Equals( instance.ptr, obj.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Specifies the parent node of the current node
// Returns:
//   EquationNode  
func (instance *UnknowEquationNode) GetParentNode()  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_GetParentNode( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Specifies the parent node of the current node
// Parameters:
//   value - EquationNode 
// Returns:
//   void  
func (instance *UnknowEquationNode) SetParentNode(value *EquationNode)  error {
	CGoReturnPtr := C.UnknowEquationNode_SetParentNode( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Convert this equtation to LaTeX expression.
// Returns:
//   string  
func (instance *UnknowEquationNode) ToLaTeX()  (string,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_ToLaTeX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Convert this equtation to MathML expression.
// Returns:
//   string  
func (instance *UnknowEquationNode) ToMathML()  (string,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_ToMathML( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Insert a node of the specified type at the end of the child node list of the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *UnknowEquationNode) AddChild_EquationNodeType(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_AddChild_EquationNodeType( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node at the end of the current node's list of child nodes.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *UnknowEquationNode) AddChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.UnknowEquationNode_AddChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Inserts a node of the specified type at the specified index position in the current node's child node list.
// Parameters:
//   index - int32 
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *UnknowEquationNode) InsertChild(index int32, equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_InsertChild( instance.ptr, C.int(index), C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node after the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *UnknowEquationNode) InsertAfter(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_InsertAfter( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Inserts the specified node before the current node.
// Parameters:
//   equationType - int32 
// Returns:
//   EquationNode  
func (instance *UnknowEquationNode) InsertBefore(equationtype EquationNodeType)  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_InsertBefore( instance.ptr, C.int( int32(equationtype)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Returns the node at the specified index among the children of the current node.
// Parameters:
//   index - int32 
// Returns:
//   EquationNode  
func (instance *UnknowEquationNode) GetChild(index int32)  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_GetChild( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}
// Removes itself from the parent.
// Returns:
//   void  
func (instance *UnknowEquationNode) Remove()  error {
	CGoReturnPtr := C.UnknowEquationNode_Remove( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the specified node from the current node's children.
// Parameters:
//   node - EquationNode 
// Returns:
//   void  
func (instance *UnknowEquationNode) RemoveChild_EquationNode(node *EquationNode)  error {
	CGoReturnPtr := C.UnknowEquationNode_RemoveChild_EquationNode( instance.ptr, node.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes the node at the specified index from the current node's children.
// Parameters:
//   index - int32 
// Returns:
//   void  
func (instance *UnknowEquationNode) RemoveChild_Int(index int32)  error {
	CGoReturnPtr := C.UnknowEquationNode_RemoveChild_Integer( instance.ptr, C.int(index))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Removes all the child nodes of the current node.
// Returns:
//   void  
func (instance *UnknowEquationNode) RemoveAllChildren()  error {
	CGoReturnPtr := C.UnknowEquationNode_RemoveAllChildren( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the type of the node.
// Returns:
//   int32  
func (instance *UnknowEquationNode) GetType()  (TextNodeType,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_GetType( instance.ptr)
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
// Get the equation type of the current node
// Returns:
//   int32  
func (instance *UnknowEquationNode) GetEquationType()  (EquationNodeType,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_GetEquationType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEquationNodeType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Create a node of the specified type.
// Parameters:
//   equationType - int32 
//   workbook - Workbook 
//   parent - EquationNode 
// Returns:
//   EquationNode  
func UnknowEquationNode_CreateNode(equationtype EquationNodeType, workbook *Workbook, parent *EquationNode)  (*EquationNode,  error)  {
	CGoReturnPtr := C.UnknowEquationNode_CreateNode(C.int( int32(equationtype)), workbook.ptr, parent.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &EquationNode{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteEquationNode) 

	return result, nil 
}


func DeleteUnknowEquationNode(unknowequationnode *UnknowEquationNode){
	runtime.SetFinalizer(unknowequationnode, nil)
	C.Delete_UnknowEquationNode(unknowequationnode.ptr)
	unknowequationnode.ptr = nil
}
