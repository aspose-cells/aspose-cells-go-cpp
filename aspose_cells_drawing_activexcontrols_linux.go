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

/**************Enum ActiveXPersistenceType *****************/

// Represents the persistence method to persist an ActiveX control.
type ActiveXPersistenceType int32

const(
// The data is stored as xml data.
ActiveXPersistenceType_PropertyBag ActiveXPersistenceType = 0 

// The data is stored as a storage binary data.
ActiveXPersistenceType_Storage ActiveXPersistenceType = 1 

// The data is stored as a stream binary data.
ActiveXPersistenceType_Stream ActiveXPersistenceType = 2 

// The data is stored as a streaminit binary data.
ActiveXPersistenceType_StreamInit ActiveXPersistenceType = 3 
)

func Int32ToActiveXPersistenceType(value int32)(ActiveXPersistenceType ,error){
	switch value {
		case 0:  return ActiveXPersistenceType_PropertyBag, nil  
		case 1:  return ActiveXPersistenceType_Storage, nil  
		case 2:  return ActiveXPersistenceType_Stream, nil  
		case 3:  return ActiveXPersistenceType_StreamInit, nil  
		default:
			return 0 ,fmt.Errorf("invalid ActiveXPersistenceType value: %d", value)
	}
}

/**************Enum ControlBorderType *****************/

// Represents the border type of the ActiveX control.
type ControlBorderType int32

const(
// No border.
ControlBorderType_None ControlBorderType = 0 

// The single line.
ControlBorderType_Single ControlBorderType = 1 
)

func Int32ToControlBorderType(value int32)(ControlBorderType ,error){
	switch value {
		case 0:  return ControlBorderType_None, nil  
		case 1:  return ControlBorderType_Single, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlBorderType value: %d", value)
	}
}

/**************Enum ControlCaptionAlignmentType *****************/

// Represents the position of the Caption relative to the control.
type ControlCaptionAlignmentType int32

const(
// The left of the control.
ControlCaptionAlignmentType_Left ControlCaptionAlignmentType = 0 

// The right of the control.
ControlCaptionAlignmentType_Right ControlCaptionAlignmentType = 1 
)

func Int32ToControlCaptionAlignmentType(value int32)(ControlCaptionAlignmentType ,error){
	switch value {
		case 0:  return ControlCaptionAlignmentType_Left, nil  
		case 1:  return ControlCaptionAlignmentType_Right, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlCaptionAlignmentType value: %d", value)
	}
}

/**************Enum ControlListStyle *****************/

// Represents the visual appearance of the list in a ListBox or ComboBox.
type ControlListStyle int32

const(
// Displays a list in which the background of an item is highlighted when it is selected.
ControlListStyle_Plain ControlListStyle = 0 

// Displays a list in which an option button or a checkbox next to each entry displays the selection state of that item.
ControlListStyle_Option ControlListStyle = 1 
)

func Int32ToControlListStyle(value int32)(ControlListStyle ,error){
	switch value {
		case 0:  return ControlListStyle_Plain, nil  
		case 1:  return ControlListStyle_Option, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlListStyle value: %d", value)
	}
}

/**************Enum ControlMatchEntryType *****************/

// Represents how a ListBox or ComboBox searches its list as the user types.
type ControlMatchEntryType int32

const(
// The control searches for the next entry that starts with the character entered.
// Repeatedly typing the same letter cycles through all entries beginning with that letter.
ControlMatchEntryType_FirstLetter ControlMatchEntryType = 0 

// As each character is typed, the control searches for an entry matching all characters entered.
ControlMatchEntryType_Complete ControlMatchEntryType = 1 

// The list will not be searched when characters are typed.
ControlMatchEntryType_None ControlMatchEntryType = 2 
)

func Int32ToControlMatchEntryType(value int32)(ControlMatchEntryType ,error){
	switch value {
		case 0:  return ControlMatchEntryType_FirstLetter, nil  
		case 1:  return ControlMatchEntryType_Complete, nil  
		case 2:  return ControlMatchEntryType_None, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlMatchEntryType value: %d", value)
	}
}

/**************Enum ControlMousePointerType *****************/

// Represents the type of icon displayed as the mouse pointer for the control.
type ControlMousePointerType int32

const(
// Standard pointer.
ControlMousePointerType_Default ControlMousePointerType = 0 

// Arrow.
ControlMousePointerType_Arrow ControlMousePointerType = 1 

// Cross-hair pointer.
ControlMousePointerType_Cross ControlMousePointerType = 2 

// I-beam.
ControlMousePointerType_IBeam ControlMousePointerType = 3 

// Double arrow pointing northeast and southwest.
ControlMousePointerType_SizeNESW ControlMousePointerType = 6 

// Double arrow pointing north and south.
ControlMousePointerType_SizeNS ControlMousePointerType = 7 

// Double arrow pointing northwest and southeast.
ControlMousePointerType_SizeNWSE ControlMousePointerType = 8 

// Double arrow pointing west and east.
ControlMousePointerType_SizeWE ControlMousePointerType = 9 

// Up arrow.
ControlMousePointerType_UpArrow ControlMousePointerType = 10 

// Hourglass.
ControlMousePointerType_HourGlass ControlMousePointerType = 11 

// "Not” symbol (circle with a diagonal line) on top of the object being dragged.
ControlMousePointerType_NoDrop ControlMousePointerType = 12 

// Arrow with an hourglass.
ControlMousePointerType_AppStarting ControlMousePointerType = 13 

// Arrow with a question mark.
ControlMousePointerType_Help ControlMousePointerType = 14 

// "Size-all” cursor (arrows pointing north, south, east, and west).
ControlMousePointerType_SizeAll ControlMousePointerType = 15 

// Uses the icon specified by the MouseIcon property.
ControlMousePointerType_Custom ControlMousePointerType = 99 
)

func Int32ToControlMousePointerType(value int32)(ControlMousePointerType ,error){
	switch value {
		case 0:  return ControlMousePointerType_Default, nil  
		case 1:  return ControlMousePointerType_Arrow, nil  
		case 2:  return ControlMousePointerType_Cross, nil  
		case 3:  return ControlMousePointerType_IBeam, nil  
		case 6:  return ControlMousePointerType_SizeNESW, nil  
		case 7:  return ControlMousePointerType_SizeNS, nil  
		case 8:  return ControlMousePointerType_SizeNWSE, nil  
		case 9:  return ControlMousePointerType_SizeWE, nil  
		case 10:  return ControlMousePointerType_UpArrow, nil  
		case 11:  return ControlMousePointerType_HourGlass, nil  
		case 12:  return ControlMousePointerType_NoDrop, nil  
		case 13:  return ControlMousePointerType_AppStarting, nil  
		case 14:  return ControlMousePointerType_Help, nil  
		case 15:  return ControlMousePointerType_SizeAll, nil  
		case 99:  return ControlMousePointerType_Custom, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlMousePointerType value: %d", value)
	}
}

/**************Enum ControlPictureAlignmentType *****************/

// Represents the alignment of the picture inside the Form or Image.
type ControlPictureAlignmentType int32

const(
// The top left corner.
ControlPictureAlignmentType_TopLeft ControlPictureAlignmentType = 0 

// The top right corner.
ControlPictureAlignmentType_TopRight ControlPictureAlignmentType = 1 

// The center.
ControlPictureAlignmentType_Center ControlPictureAlignmentType = 2 

// The bottom left corner.
ControlPictureAlignmentType_BottomLeft ControlPictureAlignmentType = 3 

// The bottom right corner.
ControlPictureAlignmentType_BottomRight ControlPictureAlignmentType = 4 
)

func Int32ToControlPictureAlignmentType(value int32)(ControlPictureAlignmentType ,error){
	switch value {
		case 0:  return ControlPictureAlignmentType_TopLeft, nil  
		case 1:  return ControlPictureAlignmentType_TopRight, nil  
		case 2:  return ControlPictureAlignmentType_Center, nil  
		case 3:  return ControlPictureAlignmentType_BottomLeft, nil  
		case 4:  return ControlPictureAlignmentType_BottomRight, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlPictureAlignmentType value: %d", value)
	}
}

/**************Enum ControlPicturePositionType *****************/

// Represents the location of the control's picture relative to its caption.
type ControlPicturePositionType int32

const(
// The picture appears to the left of the caption.
// The caption is aligned with the top of the picture.
ControlPicturePositionType_LeftTop ControlPicturePositionType = 131072 

// The picture appears to the left of the caption.
// The caption is centered relative to the picture.
ControlPicturePositionType_LeftCenter ControlPicturePositionType = 327683 

// The picture appears to the left of the caption.
// The caption is aligned with the bottom of the picture.
ControlPicturePositionType_LeftBottom ControlPicturePositionType = 524294 

// The picture appears to the right of the caption.
// The caption is aligned with the top of the picture.
ControlPicturePositionType_RightTop ControlPicturePositionType = 2 

// The picture appears to the right of the caption.
// The caption is centered relative to the picture.
ControlPicturePositionType_RightCenter ControlPicturePositionType = 196613 

// The picture appears to the right of the caption.
// The caption is aligned with the bottom of the picture.
ControlPicturePositionType_RightBottom ControlPicturePositionType = 393224 

// The picture appears above the caption.
// The caption is aligned with the left edge of the picture.
ControlPicturePositionType_AboveLeft ControlPicturePositionType = 393216 

// The picture appears above the caption.
// The caption is centered below the picture.
ControlPicturePositionType_AboveCenter ControlPicturePositionType = 458753 

// The picture appears above the caption.
// The caption is aligned with the right edge of the picture.
ControlPicturePositionType_AboveRight ControlPicturePositionType = 524290 

// The picture appears below the caption.
// The caption is aligned with the left edge of the picture.
ControlPicturePositionType_BelowLeft ControlPicturePositionType = 6 

// The picture appears below the caption.
// The caption is centered above the picture.
ControlPicturePositionType_BelowCenter ControlPicturePositionType = 65543 

// The picture appears below the caption.
// The caption is aligned with the right edge of the picture.
ControlPicturePositionType_BelowRight ControlPicturePositionType = 131080 

// The picture appears in the center of the control.
// The caption is centered horizontally and vertically on top of the picture.
ControlPicturePositionType_Center ControlPicturePositionType = 262148 
)

func Int32ToControlPicturePositionType(value int32)(ControlPicturePositionType ,error){
	switch value {
		case 131072:  return ControlPicturePositionType_LeftTop, nil  
		case 327683:  return ControlPicturePositionType_LeftCenter, nil  
		case 524294:  return ControlPicturePositionType_LeftBottom, nil  
		case 2:  return ControlPicturePositionType_RightTop, nil  
		case 196613:  return ControlPicturePositionType_RightCenter, nil  
		case 393224:  return ControlPicturePositionType_RightBottom, nil  
		case 393216:  return ControlPicturePositionType_AboveLeft, nil  
		case 458753:  return ControlPicturePositionType_AboveCenter, nil  
		case 524290:  return ControlPicturePositionType_AboveRight, nil  
		case 6:  return ControlPicturePositionType_BelowLeft, nil  
		case 65543:  return ControlPicturePositionType_BelowCenter, nil  
		case 131080:  return ControlPicturePositionType_BelowRight, nil  
		case 262148:  return ControlPicturePositionType_Center, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlPicturePositionType value: %d", value)
	}
}

/**************Enum ControlPictureSizeMode *****************/

// Represents how to display the picture.
type ControlPictureSizeMode int32

const(
// Crops any part of the picture that is larger than the control's boundaries.
ControlPictureSizeMode_Clip ControlPictureSizeMode = 0 

// Stretches the picture to fill the control's area.
// This setting distorts the picture in either the horizontal or vertical direction.
ControlPictureSizeMode_Stretch ControlPictureSizeMode = 1 

// Enlarges the picture, but does not distort the picture in either the horizontal or vertical direction.
ControlPictureSizeMode_Zoom ControlPictureSizeMode = 3 
)

func Int32ToControlPictureSizeMode(value int32)(ControlPictureSizeMode ,error){
	switch value {
		case 0:  return ControlPictureSizeMode_Clip, nil  
		case 1:  return ControlPictureSizeMode_Stretch, nil  
		case 3:  return ControlPictureSizeMode_Zoom, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlPictureSizeMode value: %d", value)
	}
}

/**************Enum ControlScrollBarType *****************/

// Represents the type of scroll bar.
type ControlScrollBarType int32

const(
// Displays no scroll bars.
ControlScrollBarType_None ControlScrollBarType = 0 

// Displays a horizontal scroll bar.
ControlScrollBarType_Horizontal ControlScrollBarType = 1 

// Displays a vertical scroll bar.
ControlScrollBarType_BarsVertical ControlScrollBarType = 2 

// Displays both a horizontal and a vertical scroll bar.
ControlScrollBarType_BarsBoth ControlScrollBarType = 3 
)

func Int32ToControlScrollBarType(value int32)(ControlScrollBarType ,error){
	switch value {
		case 0:  return ControlScrollBarType_None, nil  
		case 1:  return ControlScrollBarType_Horizontal, nil  
		case 2:  return ControlScrollBarType_BarsVertical, nil  
		case 3:  return ControlScrollBarType_BarsBoth, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlScrollBarType value: %d", value)
	}
}

/**************Enum ControlScrollOrientation *****************/

// Represents type of scroll orientation
type ControlScrollOrientation int32

const(
// Control is rendered horizontally when the control's width is greater than its height.
// Control is rendered vertically otherwise.
ControlScrollOrientation_Auto ControlScrollOrientation = 3 

// Control is rendered vertically.
ControlScrollOrientation_Vertical ControlScrollOrientation = 0 

// Control is rendered horizontally.
ControlScrollOrientation_Horizontal ControlScrollOrientation = 1 
)

func Int32ToControlScrollOrientation(value int32)(ControlScrollOrientation ,error){
	switch value {
		case 3:  return ControlScrollOrientation_Auto, nil  
		case 0:  return ControlScrollOrientation_Vertical, nil  
		case 1:  return ControlScrollOrientation_Horizontal, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlScrollOrientation value: %d", value)
	}
}

/**************Enum ControlSpecialEffectType *****************/

// Represents the type of special effect.
type ControlSpecialEffectType int32

const(
// Flat
ControlSpecialEffectType_Flat ControlSpecialEffectType = 0 

// Raised
ControlSpecialEffectType_Raised ControlSpecialEffectType = 1 

// Sunken
ControlSpecialEffectType_Sunken ControlSpecialEffectType = 2 

// Etched
ControlSpecialEffectType_Etched ControlSpecialEffectType = 3 

// Bump
ControlSpecialEffectType_Bump ControlSpecialEffectType = 6 
)

func Int32ToControlSpecialEffectType(value int32)(ControlSpecialEffectType ,error){
	switch value {
		case 0:  return ControlSpecialEffectType_Flat, nil  
		case 1:  return ControlSpecialEffectType_Raised, nil  
		case 2:  return ControlSpecialEffectType_Sunken, nil  
		case 3:  return ControlSpecialEffectType_Etched, nil  
		case 6:  return ControlSpecialEffectType_Bump, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlSpecialEffectType value: %d", value)
	}
}

/**************Enum ControlType *****************/

// Represents all type of ActiveX control.
type ControlType int32

const(
// Button
ControlType_CommandButton ControlType = 0 

// ComboBox
ControlType_ComboBox ControlType = 1 

// CheckBox
ControlType_CheckBox ControlType = 2 

// ListBox
ControlType_ListBox ControlType = 3 

// TextBox
ControlType_TextBox ControlType = 4 

// Spinner
ControlType_SpinButton ControlType = 5 

// RadioButton
ControlType_RadioButton ControlType = 6 

// Label
ControlType_Label ControlType = 7 

// Image
ControlType_Image ControlType = 8 

// ToggleButton
ControlType_ToggleButton ControlType = 9 

// ScrollBar
ControlType_ScrollBar ControlType = 10 

// ScrollBar
ControlType_BarCode ControlType = 11 

// Unknown
ControlType_Unknown ControlType = 12 
)

func Int32ToControlType(value int32)(ControlType ,error){
	switch value {
		case 0:  return ControlType_CommandButton, nil  
		case 1:  return ControlType_ComboBox, nil  
		case 2:  return ControlType_CheckBox, nil  
		case 3:  return ControlType_ListBox, nil  
		case 4:  return ControlType_TextBox, nil  
		case 5:  return ControlType_SpinButton, nil  
		case 6:  return ControlType_RadioButton, nil  
		case 7:  return ControlType_Label, nil  
		case 8:  return ControlType_Image, nil  
		case 9:  return ControlType_ToggleButton, nil  
		case 10:  return ControlType_ScrollBar, nil  
		case 11:  return ControlType_BarCode, nil  
		case 12:  return ControlType_Unknown, nil  
		default:
			return 0 ,fmt.Errorf("invalid ControlType value: %d", value)
	}
}

/**************Enum DropButtonStyle *****************/

// Represents the symbol displayed on the drop button.
type DropButtonStyle int32

const(
// Displays a button with no symbol.
DropButtonStyle_Plain DropButtonStyle = 0 

// Displays a button with a down arrow.
DropButtonStyle_Arrow DropButtonStyle = 1 

// Displays a button with an ellipsis (...).
DropButtonStyle_Ellipsis DropButtonStyle = 2 

// Displays a button with a horizontal line like an underscore character.
DropButtonStyle_Reduce DropButtonStyle = 3 
)

func Int32ToDropButtonStyle(value int32)(DropButtonStyle ,error){
	switch value {
		case 0:  return DropButtonStyle_Plain, nil  
		case 1:  return DropButtonStyle_Arrow, nil  
		case 2:  return DropButtonStyle_Ellipsis, nil  
		case 3:  return DropButtonStyle_Reduce, nil  
		default:
			return 0 ,fmt.Errorf("invalid DropButtonStyle value: %d", value)
	}
}

/**************Enum InputMethodEditorMode *****************/

// Represents the default run-time mode of the Input Method Editor.
type InputMethodEditorMode int32

const(
// Does not control IME.
InputMethodEditorMode_NoControl InputMethodEditorMode = 0 

// IME on.
InputMethodEditorMode_On InputMethodEditorMode = 1 

// IME off. English mode.
InputMethodEditorMode_Off InputMethodEditorMode = 2 

// IME off.User can't turn on IME by keyboard.
InputMethodEditorMode_Disable InputMethodEditorMode = 3 

// IME on with Full-width hiragana mode.
InputMethodEditorMode_Hiragana InputMethodEditorMode = 4 

// IME on with Full-width katakana mode.
InputMethodEditorMode_Katakana InputMethodEditorMode = 5 

// IME on with Half-width katakana mode.
InputMethodEditorMode_KatakanaHalf InputMethodEditorMode = 6 

// IME on with Full-width Alphanumeric mode.
InputMethodEditorMode_AlphaFull InputMethodEditorMode = 7 

// IME on with Half-width Alphanumeric mode.
InputMethodEditorMode_Alpha InputMethodEditorMode = 8 

// IME on with Full-width hangul mode.
InputMethodEditorMode_HangulFull InputMethodEditorMode = 9 

// IME on with Half-width hangul mode.
InputMethodEditorMode_Hangul InputMethodEditorMode = 10 

// IME on with Full-width hanzi mode.
InputMethodEditorMode_HanziFull InputMethodEditorMode = 11 

// IME on with Half-width hanzi mode.
InputMethodEditorMode_Hanzi InputMethodEditorMode = 12 
)

func Int32ToInputMethodEditorMode(value int32)(InputMethodEditorMode ,error){
	switch value {
		case 0:  return InputMethodEditorMode_NoControl, nil  
		case 1:  return InputMethodEditorMode_On, nil  
		case 2:  return InputMethodEditorMode_Off, nil  
		case 3:  return InputMethodEditorMode_Disable, nil  
		case 4:  return InputMethodEditorMode_Hiragana, nil  
		case 5:  return InputMethodEditorMode_Katakana, nil  
		case 6:  return InputMethodEditorMode_KatakanaHalf, nil  
		case 7:  return InputMethodEditorMode_AlphaFull, nil  
		case 8:  return InputMethodEditorMode_Alpha, nil  
		case 9:  return InputMethodEditorMode_HangulFull, nil  
		case 10:  return InputMethodEditorMode_Hangul, nil  
		case 11:  return InputMethodEditorMode_HanziFull, nil  
		case 12:  return InputMethodEditorMode_Hanzi, nil  
		default:
			return 0 ,fmt.Errorf("invalid InputMethodEditorMode value: %d", value)
	}
}

/**************Enum ShowDropButtonType *****************/

// Specifies when to show the drop button
type ShowDropButtonType int32

const(
// Never show the drop button.
ShowDropButtonType_Never ShowDropButtonType = 0 

// Show the drop button when the control has the focus.
ShowDropButtonType_Focus ShowDropButtonType = 1 

// Always show the drop button.
ShowDropButtonType_Always ShowDropButtonType = 2 
)

func Int32ToShowDropButtonType(value int32)(ShowDropButtonType ,error){
	switch value {
		case 0:  return ShowDropButtonType_Never, nil  
		case 1:  return ShowDropButtonType_Focus, nil  
		case 2:  return ShowDropButtonType_Always, nil  
		default:
			return 0 ,fmt.Errorf("invalid ShowDropButtonType value: %d", value)
	}
}
// Class ActiveXControl 

// Represents the ActiveX control.
type ActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControlBase 
func NewActiveXControl(src *ActiveXControlBase) ( *ActiveXControl, error) {
	activexcontrol := &ActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_ActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		activexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(activexcontrol, DeleteActiveXControl)
		return activexcontrol, nil
	} else {
		activexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return activexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *ActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *ActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *ActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *ActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *ActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *ActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteActiveXControl(activexcontrol *ActiveXControl){
	runtime.SetFinalizer(activexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_ActiveXControl"),activexcontrol.ptr)
	activexcontrol.ptr = nil
}

// Class ActiveXControlBase 

// Represents the ActiveX control.
type ActiveXControlBase struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ActiveXControlBase) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControlBase_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ActiveXControlBase) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ActiveXControlBase_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ActiveXControlBase) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ActiveXControlBase_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ActiveXControlBase_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ActiveXControlBase) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ActiveXControlBase_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ActiveXControlBase_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ActiveXControlBase) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ActiveXControlBase_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ActiveXControlBase_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ActiveXControlBase) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ActiveXControlBase_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ActiveXControlBase_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ActiveXControlBase) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ActiveXControlBase_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ActiveXControlBase) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ActiveXControlBase_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ActiveXControlBase_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ActiveXControlBase) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ActiveXControlBase_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ActiveXControlBase_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ActiveXControlBase) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ActiveXControlBase_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ActiveXControlBase_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ActiveXControlBase) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ActiveXControlBase_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ActiveXControlBase_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ActiveXControlBase) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControlBase_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControlBase_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ActiveXControlBase) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ActiveXControlBase_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ActiveXControlBase_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ActiveXControlBase) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ActiveXControlBase_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}



func DeleteActiveXControlBase(activexcontrolbase *ActiveXControlBase){
	runtime.SetFinalizer(activexcontrolbase, nil)
	C.Delete_CObject(C.CString("Delete_ActiveXControlBase"),activexcontrolbase.ptr)
	activexcontrolbase.ptr = nil
}

// Class CheckBoxActiveXControl 

// Represents a CheckBox ActiveX control.
type CheckBoxActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewCheckBoxActiveXControl(src *ActiveXControl) ( *CheckBoxActiveXControl, error) {
	checkboxactivexcontrol := &CheckBoxActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_CheckBoxActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		checkboxactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(checkboxactivexcontrol, DeleteCheckBoxActiveXControl)
		return checkboxactivexcontrol, nil
	} else {
		checkboxactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return checkboxactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the group's name.
// Returns:
//   string  
func (instance *CheckBoxActiveXControl) GetGroupName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CheckBoxActiveXControl_GetGroupName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the group's name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetGroupName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CheckBoxActiveXControl_SetGroupName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the position of the Caption relative to the control.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetAlignment()  (ControlCaptionAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlCaptionAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the position of the Caption relative to the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetAlignment(value ControlCaptionAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsWordWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsWordWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsWordWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the descriptive text that appears on a control.
// Returns:
//   string  
func (instance *CheckBoxActiveXControl) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CheckBoxActiveXControl_GetCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetCaption(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CheckBoxActiveXControl_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetPicturePosition()  (ControlPicturePositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetPicturePosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPicturePositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetPicturePosition(value ControlPicturePositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetPicturePosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the data of the picture.
// Returns:
//   []byte  
func (instance *CheckBoxActiveXControl) GetPicture()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CheckBoxActiveXControl_GetPicture"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the data of the picture.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetPicture(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("CheckBoxActiveXControl_SetPicture"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the accelerator key for the control.
// Returns:
//   byte  
func (instance *CheckBoxActiveXControl) GetAccelerator()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("CheckBoxActiveXControl_GetAccelerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the accelerator key for the control.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetAccelerator(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("CheckBoxActiveXControl_SetAccelerator"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if the control is checked or not.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetValue()  (CheckValueType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCheckValueType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates if the control is checked or not.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetValue(value CheckValueType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetValue"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how the specified control will display Null values.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsTripleState()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsTripleState"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates how the specified control will display Null values.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsTripleState(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsTripleState"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *CheckBoxActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("CheckBoxActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *CheckBoxActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CheckBoxActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *CheckBoxActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("CheckBoxActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *CheckBoxActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CheckBoxActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("CheckBoxActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CheckBoxActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CheckBoxActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *CheckBoxActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CheckBoxActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CheckBoxActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *CheckBoxActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CheckBoxActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CheckBoxActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *CheckBoxActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("CheckBoxActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("CheckBoxActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *CheckBoxActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("CheckBoxActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("CheckBoxActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("CheckBoxActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("CheckBoxActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("CheckBoxActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("CheckBoxActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *CheckBoxActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CheckBoxActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CheckBoxActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *CheckBoxActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *CheckBoxActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteCheckBoxActiveXControl(checkboxactivexcontrol *CheckBoxActiveXControl){
	runtime.SetFinalizer(checkboxactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_CheckBoxActiveXControl"),checkboxactivexcontrol.ptr)
	checkboxactivexcontrol.ptr = nil
}

// Class ComboBoxActiveXControl 

// Represents a ComboBox ActiveX control.
type ComboBoxActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewComboBoxActiveXControl(src *ActiveXControl) ( *ComboBoxActiveXControl, error) {
	comboboxactivexcontrol := &ComboBoxActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_ComboBoxActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		comboboxactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(comboboxactivexcontrol, DeleteComboBoxActiveXControl)
		return comboboxactivexcontrol, nil
	} else {
		comboboxactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return comboboxactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the maximum number of characters
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetMaxLength()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetMaxLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the maximum number of characters
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetMaxLength(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetMaxLength"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the width in unit of points.
// Returns:
//   float64  
func (instance *ComboBoxActiveXControl) GetListWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ComboBoxActiveXControl_GetListWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the width in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetListWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ComboBoxActiveXControl_SetListWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents how the Value property is determined for a ComboBox or ListBox
// when the MultiSelect properties value (fmMultiSelectSingle).
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetBoundColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetBoundColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents how the Value property is determined for a ComboBox or ListBox
// when the MultiSelect properties value (fmMultiSelectSingle).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetBoundColumn(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetBoundColumn"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the column in a ComboBox or ListBox to display to the user.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetTextColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetTextColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the column in a ComboBox or ListBox to display to the user.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetTextColumn(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetTextColumn"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the number of columns to display in a ComboBox or ListBox.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetColumnCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetColumnCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the number of columns to display in a ComboBox or ListBox.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetColumnCount(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetColumnCount"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the maximum number of rows to display in the list.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetListRows()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetListRows"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the maximum number of rows to display in the list.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetListRows(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetListRows"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how a ListBox or ComboBox searches its list as the user types.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetMatchEntry()  (ControlMatchEntryType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetMatchEntry"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMatchEntryType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates how a ListBox or ComboBox searches its list as the user types.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetMatchEntry(value ControlMatchEntryType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetMatchEntry"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the symbol displayed on the drop button
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetDropButtonStyle()  (DropButtonStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetDropButtonStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDropButtonStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the symbol displayed on the drop button
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetDropButtonStyle(value DropButtonStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetDropButtonStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the symbol displayed on the drop button
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetShowDropButtonTypeWhen()  (ShowDropButtonType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetShowDropButtonTypeWhen"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToShowDropButtonType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the symbol displayed on the drop button
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetShowDropButtonTypeWhen(value ShowDropButtonType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetShowDropButtonTypeWhen"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the visual appearance.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetListStyle()  (ControlListStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetListStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlListStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the visual appearance.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetListStyle(value ControlListStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetListStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the type of border used by the control.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetBorderStyle()  (ControlBorderType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetBorderStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlBorderType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the type of border used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetBorderStyle(value ControlBorderType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetBorderStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetBorderOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetBorderOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetBorderOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetBorderOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the user can type into the control.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsEditable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsEditable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the user can type into the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsEditable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsEditable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether column headings are displayed.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) GetShowColumnHeads()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_GetShowColumnHeads"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether column headings are displayed.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetShowColumnHeads(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetShowColumnHeads"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether dragging and dropping is enabled for the control.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsDragBehaviorEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsDragBehaviorEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether dragging and dropping is enabled for the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsDragBehaviorEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsDragBehaviorEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies selection behavior when entering the control.
// True specifies that the selection remains unchanged from last time the control was active.
// False specifies that all the text in the control will be selected when entering the control.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) GetEnterFieldBehavior()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_GetEnterFieldBehavior"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies selection behavior when entering the control.
// True specifies that the selection remains unchanged from last time the control was active.
// False specifies that all the text in the control will be selected when entering the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetEnterFieldBehavior(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetEnterFieldBehavior"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the user can select a line of text by clicking in the region to the left of the text.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) GetSelectionMargin()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_GetSelectionMargin"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the user can select a line of text by clicking in the region to the left of the text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetSelectionMargin(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetSelectionMargin"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the value of the control.
// Returns:
//   string  
func (instance *ComboBoxActiveXControl) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ComboBoxActiveXControl_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the value of the control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ComboBoxActiveXControl_SetValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether selected text in the control appears highlighted when the control does not have focus.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) GetHideSelection()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_GetHideSelection"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether selected text in the control appears highlighted when the control does not have focus.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetHideSelection(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetHideSelection"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the column.
// Returns:
//   float64  
func (instance *ComboBoxActiveXControl) GetColumnWidths()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ComboBoxActiveXControl_GetColumnWidths"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the column.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetColumnWidths(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ComboBoxActiveXControl_SetColumnWidths"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the basic unit used to extend a selection.
// True specifies that the basic unit is a single character.
// false specifies that the basic unit is a whole word.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsAutoWordSelected()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsAutoWordSelected"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the basic unit used to extend a selection.
// True specifies that the basic unit is a single character.
// false specifies that the basic unit is a whole word.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsAutoWordSelected(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsAutoWordSelected"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *ComboBoxActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ComboBoxActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ComboBoxActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ComboBoxActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ComboBoxActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ComboBoxActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ComboBoxActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ComboBoxActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ComboBoxActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ComboBoxActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ComboBoxActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ComboBoxActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ComboBoxActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ComboBoxActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ComboBoxActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ComboBoxActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ComboBoxActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ComboBoxActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ComboBoxActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ComboBoxActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ComboBoxActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ComboBoxActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ComboBoxActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ComboBoxActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ComboBoxActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ComboBoxActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ComboBoxActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ComboBoxActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ComboBoxActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *ComboBoxActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteComboBoxActiveXControl(comboboxactivexcontrol *ComboBoxActiveXControl){
	runtime.SetFinalizer(comboboxactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_ComboBoxActiveXControl"),comboboxactivexcontrol.ptr)
	comboboxactivexcontrol.ptr = nil
}

// Class CommandButtonActiveXControl 

// Represents a command button.
type CommandButtonActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewCommandButtonActiveXControl(src *ActiveXControl) ( *CommandButtonActiveXControl, error) {
	commandbuttonactivexcontrol := &CommandButtonActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_CommandButtonActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		commandbuttonactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(commandbuttonactivexcontrol, DeleteCommandButtonActiveXControl)
		return commandbuttonactivexcontrol, nil
	} else {
		commandbuttonactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return commandbuttonactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CommandButtonActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Returns:
//   string  
func (instance *CommandButtonActiveXControl) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CommandButtonActiveXControl_GetCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetCaption(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CommandButtonActiveXControl_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetPicturePosition()  (ControlPicturePositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CommandButtonActiveXControl_GetPicturePosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPicturePositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetPicturePosition(value ControlPicturePositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CommandButtonActiveXControl_SetPicturePosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the data of the picture.
// Returns:
//   []byte  
func (instance *CommandButtonActiveXControl) GetPicture()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CommandButtonActiveXControl_GetPicture"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the data of the picture.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetPicture(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("CommandButtonActiveXControl_SetPicture"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the accelerator key for the control.
// Returns:
//   byte  
func (instance *CommandButtonActiveXControl) GetAccelerator()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("CommandButtonActiveXControl_GetAccelerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the accelerator key for the control.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetAccelerator(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("CommandButtonActiveXControl_SetAccelerator"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control takes the focus when clicked.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) GetTakeFocusOnClick()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_GetTakeFocusOnClick"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control takes the focus when clicked.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetTakeFocusOnClick(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetTakeFocusOnClick"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsWordWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsWordWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetIsWordWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CommandButtonActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CommandButtonActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *CommandButtonActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("CommandButtonActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CommandButtonActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CommandButtonActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *CommandButtonActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CommandButtonActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *CommandButtonActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("CommandButtonActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *CommandButtonActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("CommandButtonActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("CommandButtonActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("CommandButtonActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("CommandButtonActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *CommandButtonActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CommandButtonActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CommandButtonActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *CommandButtonActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("CommandButtonActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("CommandButtonActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *CommandButtonActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("CommandButtonActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("CommandButtonActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *CommandButtonActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("CommandButtonActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("CommandButtonActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("CommandButtonActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("CommandButtonActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("CommandButtonActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("CommandButtonActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *CommandButtonActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("CommandButtonActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("CommandButtonActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *CommandButtonActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *CommandButtonActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteCommandButtonActiveXControl(commandbuttonactivexcontrol *CommandButtonActiveXControl){
	runtime.SetFinalizer(commandbuttonactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_CommandButtonActiveXControl"),commandbuttonactivexcontrol.ptr)
	commandbuttonactivexcontrol.ptr = nil
}

// Class ImageActiveXControl 

// Represents the image control.
type ImageActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewImageActiveXControl(src *ActiveXControl) ( *ImageActiveXControl, error) {
	imageactivexcontrol := &ImageActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_ImageActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		imageactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(imageactivexcontrol, DeleteImageActiveXControl)
		return imageactivexcontrol, nil
	} else {
		imageactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return imageactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetBorderOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ImageActiveXControl_GetBorderOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetBorderOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ImageActiveXControl_SetBorderOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the type of border used by the control.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetBorderStyle()  (ControlBorderType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetBorderStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlBorderType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the type of border used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetBorderStyle(value ControlBorderType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetBorderStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets how to display the picture.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetPictureSizeMode()  (ControlPictureSizeMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetPictureSizeMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPictureSizeMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets how to display the picture.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetPictureSizeMode(value ControlPictureSizeMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetPictureSizeMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the data of the picture.
// Returns:
//   []byte  
func (instance *ImageActiveXControl) GetPicture()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ImageActiveXControl_GetPicture"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the data of the picture.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetPicture(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ImageActiveXControl_SetPicture"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the alignment of the picture inside the Form or Image.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetPictureAlignment()  (ControlPictureAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetPictureAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPictureAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the alignment of the picture inside the Form or Image.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetPictureAlignment(value ControlPictureAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetPictureAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the picture is tiled across the background.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsTiled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsTiled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the picture is tiled across the background.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsTiled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetIsTiled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *ImageActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ImageActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ImageActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ImageActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ImageActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ImageActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ImageActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ImageActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ImageActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ImageActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ImageActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ImageActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ImageActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ImageActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ImageActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ImageActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ImageActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ImageActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ImageActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ImageActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ImageActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ImageActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ImageActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ImageActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ImageActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ImageActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ImageActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ImageActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ImageActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ImageActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ImageActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ImageActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *ImageActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteImageActiveXControl(imageactivexcontrol *ImageActiveXControl){
	runtime.SetFinalizer(imageactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_ImageActiveXControl"),imageactivexcontrol.ptr)
	imageactivexcontrol.ptr = nil
}

// Class LabelActiveXControl 

// Represents the label ActiveX control.
type LabelActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewLabelActiveXControl(src *ActiveXControl) ( *LabelActiveXControl, error) {
	labelactivexcontrol := &LabelActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_LabelActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		labelactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(labelactivexcontrol, DeleteLabelActiveXControl)
		return labelactivexcontrol, nil
	} else {
		labelactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return labelactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Returns:
//   string  
func (instance *LabelActiveXControl) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LabelActiveXControl_GetCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetCaption(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LabelActiveXControl_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetPicturePosition()  (ControlPicturePositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetPicturePosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPicturePositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetPicturePosition(value ControlPicturePositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LabelActiveXControl_SetPicturePosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetBorderOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LabelActiveXControl_GetBorderOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetBorderOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LabelActiveXControl_SetBorderOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the type of border used by the control.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetBorderStyle()  (ControlBorderType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetBorderStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlBorderType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the type of border used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetBorderStyle(value ControlBorderType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LabelActiveXControl_SetBorderStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LabelActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the data of the picture.
// Returns:
//   []byte  
func (instance *LabelActiveXControl) GetPicture()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LabelActiveXControl_GetPicture"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the data of the picture.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetPicture(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LabelActiveXControl_SetPicture"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the accelerator key for the control.
// Returns:
//   byte  
func (instance *LabelActiveXControl) GetAccelerator()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("LabelActiveXControl_GetAccelerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the accelerator key for the control.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetAccelerator(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("LabelActiveXControl_SetAccelerator"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsWordWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsWordWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetIsWordWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LabelActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *LabelActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LabelActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LabelActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *LabelActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LabelActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *LabelActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("LabelActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *LabelActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("LabelActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("LabelActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("LabelActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("LabelActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *LabelActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LabelActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LabelActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *LabelActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("LabelActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("LabelActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *LabelActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("LabelActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("LabelActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *LabelActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("LabelActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("LabelActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LabelActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LabelActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("LabelActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("LabelActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *LabelActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *LabelActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("LabelActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("LabelActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *LabelActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *LabelActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteLabelActiveXControl(labelactivexcontrol *LabelActiveXControl){
	runtime.SetFinalizer(labelactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_LabelActiveXControl"),labelactivexcontrol.ptr)
	labelactivexcontrol.ptr = nil
}

// Class ListBoxActiveXControl 

// Represents a ListBox ActiveX control.
type ListBoxActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewListBoxActiveXControl(src *ActiveXControl) ( *ListBoxActiveXControl, error) {
	listboxactivexcontrol := &ListBoxActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_ListBoxActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		listboxactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(listboxactivexcontrol, DeleteListBoxActiveXControl)
		return listboxactivexcontrol, nil
	} else {
		listboxactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return listboxactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates specifies whether the control has vertical scroll bars, horizontal scroll bars, both, or neither.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetScrollBars()  (ControlScrollBarType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetScrollBars"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlScrollBarType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates specifies whether the control has vertical scroll bars, horizontal scroll bars, both, or neither.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetScrollBars(value ControlScrollBarType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetScrollBars"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the width in unit of points.
// Returns:
//   float64  
func (instance *ListBoxActiveXControl) GetListWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ListBoxActiveXControl_GetListWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the width in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetListWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ListBoxActiveXControl_SetListWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents how the Value property is determined for a ComboBox or ListBox
// when the MultiSelect properties value (fmMultiSelectSingle).
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetBoundColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ListBoxActiveXControl_GetBoundColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents how the Value property is determined for a ComboBox or ListBox
// when the MultiSelect properties value (fmMultiSelectSingle).
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetBoundColumn(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ListBoxActiveXControl_SetBoundColumn"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the column in a ComboBox or ListBox to display to the user.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetTextColumn()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ListBoxActiveXControl_GetTextColumn"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the column in a ComboBox or ListBox to display to the user.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetTextColumn(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ListBoxActiveXControl_SetTextColumn"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the number of columns to display in a ComboBox or ListBox.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetColumnCount()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ListBoxActiveXControl_GetColumnCount"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Represents the number of columns to display in a ComboBox or ListBox.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetColumnCount(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ListBoxActiveXControl_SetColumnCount"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how a ListBox or ComboBox searches its list as the user types.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetMatchEntry()  (ControlMatchEntryType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetMatchEntry"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMatchEntryType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates how a ListBox or ComboBox searches its list as the user types.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetMatchEntry(value ControlMatchEntryType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetMatchEntry"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the visual appearance.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetListStyle()  (ControlListStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetListStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlListStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the visual appearance.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetListStyle(value ControlListStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetListStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control permits multiple selections.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetSelectionType()  (SelectionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetSelectionType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToSelectionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether the control permits multiple selections.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetSelectionType(value SelectionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetSelectionType"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the value of the control.
// Returns:
//   string  
func (instance *ListBoxActiveXControl) GetValue()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ListBoxActiveXControl_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the value of the control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetValue(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ListBoxActiveXControl_SetValue"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the type of border used by the control.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetBorderStyle()  (ControlBorderType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetBorderStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlBorderType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the type of border used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetBorderStyle(value ControlBorderType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetBorderStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetBorderOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ListBoxActiveXControl_GetBorderOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetBorderOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ListBoxActiveXControl_SetBorderOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether column headings are displayed.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) GetShowColumnHeads()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_GetShowColumnHeads"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether column headings are displayed.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetShowColumnHeads(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetShowColumnHeads"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control will only show complete lines of text without showing any partial lines.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) GetIntegralHeight()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_GetIntegralHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will only show complete lines of text without showing any partial lines.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIntegralHeight(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetIntegralHeight"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the column.
// Returns:
//   float64  
func (instance *ListBoxActiveXControl) GetColumnWidths()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ListBoxActiveXControl_GetColumnWidths"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the column.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetColumnWidths(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ListBoxActiveXControl_SetColumnWidths"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *ListBoxActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ListBoxActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ListBoxActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ListBoxActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ListBoxActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ListBoxActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ListBoxActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ListBoxActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ListBoxActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ListBoxActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ListBoxActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ListBoxActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ListBoxActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ListBoxActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ListBoxActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ListBoxActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ListBoxActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ListBoxActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ListBoxActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ListBoxActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ListBoxActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ListBoxActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ListBoxActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ListBoxActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ListBoxActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ListBoxActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ListBoxActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ListBoxActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ListBoxActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ListBoxActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ListBoxActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *ListBoxActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteListBoxActiveXControl(listboxactivexcontrol *ListBoxActiveXControl){
	runtime.SetFinalizer(listboxactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_ListBoxActiveXControl"),listboxactivexcontrol.ptr)
	listboxactivexcontrol.ptr = nil
}

// Class RadioButtonActiveXControl 

// Represents a RadioButton ActiveX control.
type RadioButtonActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ToggleButtonActiveXControl 
func NewRadioButtonActiveXControl(src *ToggleButtonActiveXControl) ( *RadioButtonActiveXControl, error) {
	radiobuttonactivexcontrol := &RadioButtonActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_RadioButtonActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		radiobuttonactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(radiobuttonactivexcontrol, DeleteRadioButtonActiveXControl)
		return radiobuttonactivexcontrol, nil
	} else {
		radiobuttonactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return radiobuttonactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the group's name.
// Returns:
//   string  
func (instance *RadioButtonActiveXControl) GetGroupName()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("RadioButtonActiveXControl_GetGroupName"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the group's name.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetGroupName(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("RadioButtonActiveXControl_SetGroupName"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the position of the Caption relative to the control.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetAlignment()  (ControlCaptionAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetAlignment"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlCaptionAlignmentType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the position of the Caption relative to the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetAlignment(value ControlCaptionAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetAlignment"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsWordWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsWordWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsWordWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *RadioButtonActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("RadioButtonActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *RadioButtonActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("RadioButtonActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *RadioButtonActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("RadioButtonActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *RadioButtonActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("RadioButtonActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("RadioButtonActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *RadioButtonActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("RadioButtonActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("RadioButtonActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *RadioButtonActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("RadioButtonActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("RadioButtonActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *RadioButtonActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("RadioButtonActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("RadioButtonActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *RadioButtonActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("RadioButtonActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("RadioButtonActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("RadioButtonActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("RadioButtonActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("RadioButtonActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("RadioButtonActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the descriptive text that appears on a control.
// Returns:
//   string  
func (instance *RadioButtonActiveXControl) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("RadioButtonActiveXControl_GetCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetCaption(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("RadioButtonActiveXControl_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetPicturePosition()  (ControlPicturePositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetPicturePosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPicturePositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetPicturePosition(value ControlPicturePositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetPicturePosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the data of the picture.
// Returns:
//   []byte  
func (instance *RadioButtonActiveXControl) GetPicture()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("RadioButtonActiveXControl_GetPicture"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the data of the picture.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetPicture(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("RadioButtonActiveXControl_SetPicture"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the accelerator key for the control.
// Returns:
//   byte  
func (instance *RadioButtonActiveXControl) GetAccelerator()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("RadioButtonActiveXControl_GetAccelerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the accelerator key for the control.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetAccelerator(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("RadioButtonActiveXControl_SetAccelerator"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if the control is checked or not.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetValue()  (CheckValueType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("RadioButtonActiveXControl_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCheckValueType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates if the control is checked or not.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetValue(value CheckValueType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("RadioButtonActiveXControl_SetValue"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how the specified control will display Null values.
// Returns:
//   bool  
func (instance *RadioButtonActiveXControl) IsTripleState()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("RadioButtonActiveXControl_IsTripleState"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates how the specified control will display Null values.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsTripleState(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("RadioButtonActiveXControl_SetIsTripleState"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *RadioButtonActiveXControl) ToToggleButtonActiveXControl() *ToggleButtonActiveXControl {
	parentClass := &ToggleButtonActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *RadioButtonActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *RadioButtonActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteRadioButtonActiveXControl(radiobuttonactivexcontrol *RadioButtonActiveXControl){
	runtime.SetFinalizer(radiobuttonactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_RadioButtonActiveXControl"),radiobuttonactivexcontrol.ptr)
	radiobuttonactivexcontrol.ptr = nil
}

// Class ScrollBarActiveXControl 

// Represents the ScrollBar control.
type ScrollBarActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - SpinButtonActiveXControl 
func NewScrollBarActiveXControl(src *SpinButtonActiveXControl) ( *ScrollBarActiveXControl, error) {
	scrollbaractivexcontrol := &ScrollBarActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_ScrollBarActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		scrollbaractivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(scrollbaractivexcontrol, DeleteScrollBarActiveXControl)
		return scrollbaractivexcontrol, nil
	} else {
		scrollbaractivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return scrollbaractivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ScrollBarActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the amount by which the Position property changes
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetLargeChange()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetLargeChange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the amount by which the Position property changes
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetLargeChange(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetLargeChange"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ScrollBarActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ScrollBarActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ScrollBarActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ScrollBarActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ScrollBarActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *ScrollBarActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ScrollBarActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ScrollBarActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ScrollBarActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ScrollBarActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ScrollBarActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ScrollBarActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ScrollBarActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ScrollBarActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ScrollBarActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ScrollBarActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ScrollBarActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ScrollBarActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ScrollBarActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ScrollBarActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ScrollBarActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ScrollBarActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ScrollBarActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ScrollBarActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ScrollBarActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ScrollBarActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ScrollBarActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ScrollBarActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ScrollBarActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ScrollBarActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ScrollBarActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ScrollBarActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ScrollBarActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ScrollBarActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ScrollBarActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the minimum acceptable value.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetMin()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetMin"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the minimum acceptable value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetMin(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetMin"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the maximum acceptable value.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetMax()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetMax"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the maximum acceptable value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetMax(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetMax"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the value.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetPosition()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetPosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetPosition(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetPosition"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount by which the Position property changes
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetSmallChange()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ScrollBarActiveXControl_GetSmallChange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the amount by which the Position property changes
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetSmallChange(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ScrollBarActiveXControl_SetSmallChange"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets whether the SpinButton or ScrollBar is oriented vertically or horizontally.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetOrientation()  (ControlScrollOrientation,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ScrollBarActiveXControl_GetOrientation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlScrollOrientation(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets whether the SpinButton or ScrollBar is oriented vertically or horizontally.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ScrollBarActiveXControl) SetOrientation(value ControlScrollOrientation)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ScrollBarActiveXControl_SetOrientation"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ScrollBarActiveXControl) ToSpinButtonActiveXControl() *SpinButtonActiveXControl {
	parentClass := &SpinButtonActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *ScrollBarActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *ScrollBarActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteScrollBarActiveXControl(scrollbaractivexcontrol *ScrollBarActiveXControl){
	runtime.SetFinalizer(scrollbaractivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_ScrollBarActiveXControl"),scrollbaractivexcontrol.ptr)
	scrollbaractivexcontrol.ptr = nil
}

// Class SpinButtonActiveXControl 

// Represents the SpinButton control.
type SpinButtonActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewSpinButtonActiveXControl(src *ActiveXControl) ( *SpinButtonActiveXControl, error) {
	spinbuttonactivexcontrol := &SpinButtonActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_SpinButtonActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		spinbuttonactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(spinbuttonactivexcontrol, DeleteSpinButtonActiveXControl)
		return spinbuttonactivexcontrol, nil
	} else {
		spinbuttonactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return spinbuttonactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SpinButtonActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the minimum acceptable value.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetMin()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SpinButtonActiveXControl_GetMin"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the minimum acceptable value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetMin(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SpinButtonActiveXControl_SetMin"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the maximum acceptable value.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetMax()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SpinButtonActiveXControl_GetMax"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the maximum acceptable value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetMax(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SpinButtonActiveXControl_SetMax"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the value.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetPosition()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SpinButtonActiveXControl_GetPosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the value.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetPosition(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SpinButtonActiveXControl_SetPosition"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the amount by which the Position property changes
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetSmallChange()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SpinButtonActiveXControl_GetSmallChange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the amount by which the Position property changes
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetSmallChange(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SpinButtonActiveXControl_SetSmallChange"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets whether the SpinButton or ScrollBar is oriented vertically or horizontally.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetOrientation()  (ControlScrollOrientation,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SpinButtonActiveXControl_GetOrientation"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlScrollOrientation(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets whether the SpinButton or ScrollBar is oriented vertically or horizontally.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetOrientation(value ControlScrollOrientation)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SpinButtonActiveXControl_SetOrientation"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SpinButtonActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SpinButtonActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SpinButtonActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SpinButtonActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SpinButtonActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *SpinButtonActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SpinButtonActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SpinButtonActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SpinButtonActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *SpinButtonActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("SpinButtonActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SpinButtonActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *SpinButtonActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("SpinButtonActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *SpinButtonActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("SpinButtonActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("SpinButtonActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("SpinButtonActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("SpinButtonActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *SpinButtonActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SpinButtonActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SpinButtonActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *SpinButtonActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("SpinButtonActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("SpinButtonActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *SpinButtonActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("SpinButtonActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("SpinButtonActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *SpinButtonActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("SpinButtonActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("SpinButtonActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SpinButtonActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SpinButtonActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("SpinButtonActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("SpinButtonActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SpinButtonActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *SpinButtonActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("SpinButtonActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("SpinButtonActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *SpinButtonActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *SpinButtonActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteSpinButtonActiveXControl(spinbuttonactivexcontrol *SpinButtonActiveXControl){
	runtime.SetFinalizer(spinbuttonactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_SpinButtonActiveXControl"),spinbuttonactivexcontrol.ptr)
	spinbuttonactivexcontrol.ptr = nil
}

// Class TextBoxActiveXControl 

// Represents a text box ActiveX control.
type TextBoxActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewTextBoxActiveXControl(src *ActiveXControl) ( *TextBoxActiveXControl, error) {
	textboxactivexcontrol := &TextBoxActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_TextBoxActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		textboxactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(textboxactivexcontrol, DeleteTextBoxActiveXControl)
		return textboxactivexcontrol, nil
	} else {
		textboxactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return textboxactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the type of border used by the control.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetBorderStyle()  (ControlBorderType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetBorderStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlBorderType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the type of border used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetBorderStyle(value ControlBorderType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetBorderStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetBorderOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextBoxActiveXControl_GetBorderOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetBorderOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextBoxActiveXControl_SetBorderOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the maximum number of characters
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetMaxLength()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextBoxActiveXControl_GetMaxLength"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the maximum number of characters
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetMaxLength(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextBoxActiveXControl_SetMaxLength"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates specifies whether the control has vertical scroll bars, horizontal scroll bars, both, or neither.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetScrollBars()  (ControlScrollBarType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetScrollBars"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlScrollBarType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates specifies whether the control has vertical scroll bars, horizontal scroll bars, both, or neither.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetScrollBars(value ControlScrollBarType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetScrollBars"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets a character to be displayed in place of the characters entered.
// Returns:
//   byte  
func (instance *TextBoxActiveXControl) GetPasswordChar()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("TextBoxActiveXControl_GetPasswordChar"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets a character to be displayed in place of the characters entered.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetPasswordChar(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("TextBoxActiveXControl_SetPasswordChar"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the user can type into the control.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsEditable()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsEditable"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the user can type into the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsEditable(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsEditable"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control will only show complete lines of text without showing any partial lines.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) GetIntegralHeight()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_GetIntegralHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will only show complete lines of text without showing any partial lines.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIntegralHeight(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIntegralHeight"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether dragging and dropping is enabled for the control.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsDragBehaviorEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsDragBehaviorEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether dragging and dropping is enabled for the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsDragBehaviorEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsDragBehaviorEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the behavior of the ENTER key.
// True specifies that pressing ENTER will create a new line.
// False specifies that pressing ENTER will move the focus to the next object in the tab order.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) GetEnterKeyBehavior()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_GetEnterKeyBehavior"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the behavior of the ENTER key.
// True specifies that pressing ENTER will create a new line.
// False specifies that pressing ENTER will move the focus to the next object in the tab order.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetEnterKeyBehavior(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetEnterKeyBehavior"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies selection behavior when entering the control.
// True specifies that the selection remains unchanged from last time the control was active.
// False specifies that all the text in the control will be selected when entering the control.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) GetEnterFieldBehavior()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_GetEnterFieldBehavior"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies selection behavior when entering the control.
// True specifies that the selection remains unchanged from last time the control was active.
// False specifies that all the text in the control will be selected when entering the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetEnterFieldBehavior(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetEnterFieldBehavior"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether tab characters are allowed in the text of the control.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) GetTabKeyBehavior()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_GetTabKeyBehavior"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether tab characters are allowed in the text of the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetTabKeyBehavior(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetTabKeyBehavior"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether selected text in the control appears highlighted when the control does not have focus.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) GetHideSelection()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_GetHideSelection"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether selected text in the control appears highlighted when the control does not have focus.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetHideSelection(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetHideSelection"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the focus will automatically move to the next control when the user enters the maximum number of characters.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsAutoTab()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsAutoTab"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the focus will automatically move to the next control when the user enters the maximum number of characters.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsAutoTab(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsAutoTab"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can display more than one line of text.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsMultiLine()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsMultiLine"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can display more than one line of text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsMultiLine(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsMultiLine"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsWordWrapped()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsWordWrapped"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsWordWrapped"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set text of the control.
// Returns:
//   string  
func (instance *TextBoxActiveXControl) GetText()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextBoxActiveXControl_GetText"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set text of the control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetText(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextBoxActiveXControl_SetText"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the symbol displayed on the drop button
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetDropButtonStyle()  (DropButtonStyle,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetDropButtonStyle"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDropButtonStyle(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the symbol displayed on the drop button
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetDropButtonStyle(value DropButtonStyle)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetDropButtonStyle"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the symbol displayed on the drop button
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetShowDropButtonTypeWhen()  (ShowDropButtonType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetShowDropButtonTypeWhen"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToShowDropButtonType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Specifies the symbol displayed on the drop button
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetShowDropButtonTypeWhen(value ShowDropButtonType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetShowDropButtonTypeWhen"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Specifies the basic unit used to extend a selection.
// True specifies that the basic unit is a single character.
// false specifies that the basic unit is a whole word.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsAutoWordSelected()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsAutoWordSelected"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Specifies the basic unit used to extend a selection.
// True specifies that the basic unit is a single character.
// false specifies that the basic unit is a whole word.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsAutoWordSelected(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsAutoWordSelected"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *TextBoxActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("TextBoxActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *TextBoxActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("TextBoxActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *TextBoxActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("TextBoxActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *TextBoxActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("TextBoxActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("TextBoxActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("TextBoxActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("TextBoxActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *TextBoxActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextBoxActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextBoxActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *TextBoxActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("TextBoxActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("TextBoxActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *TextBoxActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextBoxActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextBoxActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *TextBoxActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("TextBoxActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("TextBoxActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextBoxActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextBoxActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("TextBoxActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("TextBoxActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *TextBoxActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("TextBoxActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("TextBoxActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *TextBoxActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *TextBoxActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteTextBoxActiveXControl(textboxactivexcontrol *TextBoxActiveXControl){
	runtime.SetFinalizer(textboxactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_TextBoxActiveXControl"),textboxactivexcontrol.ptr)
	textboxactivexcontrol.ptr = nil
}

// Class ToggleButtonActiveXControl 

// Represents a ToggleButton ActiveX control.
type ToggleButtonActiveXControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewToggleButtonActiveXControl(src *ActiveXControl) ( *ToggleButtonActiveXControl, error) {
	togglebuttonactivexcontrol := &ToggleButtonActiveXControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_ToggleButtonActiveXControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		togglebuttonactivexcontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(togglebuttonactivexcontrol, DeleteToggleButtonActiveXControl)
		return togglebuttonactivexcontrol, nil
	} else {
		togglebuttonactivexcontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return togglebuttonactivexcontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Returns:
//   string  
func (instance *ToggleButtonActiveXControl) GetCaption()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ToggleButtonActiveXControl_GetCaption"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and set the descriptive text that appears on a control.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetCaption(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ToggleButtonActiveXControl_SetCaption"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetPicturePosition()  (ControlPicturePositionType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetPicturePosition"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlPicturePositionType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and set the location of the control's picture relative to its caption.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetPicturePosition(value ControlPicturePositionType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ToggleButtonActiveXControl_SetPicturePosition"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the special effect of the control.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetSpecialEffect()  (ControlSpecialEffectType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetSpecialEffect"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlSpecialEffectType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the special effect of the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetSpecialEffect(value ControlSpecialEffectType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ToggleButtonActiveXControl_SetSpecialEffect"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the data of the picture.
// Returns:
//   []byte  
func (instance *ToggleButtonActiveXControl) GetPicture()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ToggleButtonActiveXControl_GetPicture"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the data of the picture.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetPicture(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ToggleButtonActiveXControl_SetPicture"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the accelerator key for the control.
// Returns:
//   byte  
func (instance *ToggleButtonActiveXControl) GetAccelerator()  (byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZGO(C.CString("ToggleButtonActiveXControl_GetAccelerator"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := byte(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the accelerator key for the control.
// Parameters:
//   value - byte 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetAccelerator(value byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZHB(C.CString("ToggleButtonActiveXControl_SetAccelerator"), instance.ptr, C.char(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if the control is checked or not.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetValue()  (CheckValueType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetValue"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToCheckValueType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates if the control is checked or not.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetValue(value CheckValueType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ToggleButtonActiveXControl_SetValue"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates how the specified control will display Null values.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsTripleState()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsTripleState"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates how the specified control will display Null values.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsTripleState(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetIsTripleState"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ToggleButtonActiveXControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *ToggleButtonActiveXControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ToggleButtonActiveXControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ToggleButtonActiveXControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *ToggleButtonActiveXControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ToggleButtonActiveXControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ToggleButtonActiveXControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("ToggleButtonActiveXControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *ToggleButtonActiveXControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("ToggleButtonActiveXControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("ToggleButtonActiveXControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("ToggleButtonActiveXControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("ToggleButtonActiveXControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *ToggleButtonActiveXControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ToggleButtonActiveXControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ToggleButtonActiveXControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *ToggleButtonActiveXControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("ToggleButtonActiveXControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("ToggleButtonActiveXControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *ToggleButtonActiveXControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ToggleButtonActiveXControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ToggleButtonActiveXControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *ToggleButtonActiveXControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("ToggleButtonActiveXControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("ToggleButtonActiveXControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ToggleButtonActiveXControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ToggleButtonActiveXControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("ToggleButtonActiveXControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("ToggleButtonActiveXControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *ToggleButtonActiveXControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("ToggleButtonActiveXControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("ToggleButtonActiveXControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *ToggleButtonActiveXControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *ToggleButtonActiveXControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteToggleButtonActiveXControl(togglebuttonactivexcontrol *ToggleButtonActiveXControl){
	runtime.SetFinalizer(togglebuttonactivexcontrol, nil)
	C.Delete_CObject(C.CString("Delete_ToggleButtonActiveXControl"),togglebuttonactivexcontrol.ptr)
	togglebuttonactivexcontrol.ptr = nil
}

// Class UnknownControl 

// Unknow control.
type UnknownControl struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - ActiveXControl 
func NewUnknownControl(src *ActiveXControl) ( *UnknownControl, error) {
	unknowncontrol := &UnknownControl{}
	var src_ptr unsafe.Pointer = nil
	if src != nil {
	  src_ptr =src.ptr
	}

	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("New_UnknownControl"),src_ptr)
	if CGoReturnPtr.error_no == 0 {
		unknowncontrol.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(unknowncontrol, DeleteUnknownControl)
		return unknowncontrol, nil
	} else {
		unknowncontrol.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return unknowncontrol, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *UnknownControl) IsNull()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_IsNull"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets the related data.
// Parameters:
//   relId - string 
// Returns:
//   []byte  
func (instance *UnknownControl) GetRelationshipData(relid string)  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZJR(C.CString("UnknownControl_GetRelationshipData"), instance.ptr, C.CString(relid))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets the binary data of the control.
// Returns:
//   []byte  
func (instance *UnknownControl) GetData()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("UnknownControl_GetData"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *UnknownControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("UnknownControl_GetType"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *UnknownControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_IsEnabled"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("UnknownControl_SetIsEnabled"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether data in the control is locked for editing.
// Returns:
//   bool  
func (instance *UnknownControl) IsLocked()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_IsLocked"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("UnknownControl_SetIsLocked"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control is transparent.
// Returns:
//   bool  
func (instance *UnknownControl) IsTransparent()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_IsTransparent"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("UnknownControl_SetIsTransparent"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Returns:
//   int32  
func (instance *UnknownControl) GetIMEMode()  (InputMethodEditorMode,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("UnknownControl_GetIMEMode"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToInputMethodEditorMode(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the default run-time mode of the Input Method Editor for the control as it receives focus.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *UnknownControl) SetIMEMode(value InputMethodEditorMode)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("UnknownControl_SetIMEMode"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Represents the font of the control.
// Returns:
//   Font  
func (instance *UnknownControl) GetFont()  (*Font,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("UnknownControl_GetFont"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Font{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteFont) 

	return result, nil 
}
// Represents how to align the text used by the control.
// Returns:
//   int32  
func (instance *UnknownControl) GetTextAlign()  (TextAlignmentType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("UnknownControl_GetTextAlign"), instance.ptr)
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
// Represents how to align the text used by the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *UnknownControl) SetTextAlign(value TextAlignmentType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("UnknownControl_SetTextAlign"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Returns:
//   bool  
func (instance *UnknownControl) IsAutoSize()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_IsAutoSize"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("UnknownControl_SetIsAutoSize"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *UnknownControl) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZI(C.CString("UnknownControl_GetWorkbook"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Workbook{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteWorkbook) 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Returns:
//   []byte  
func (instance *UnknownControl) GetMouseIcon()  ([]byte,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBK(C.CString("UnknownControl_GetMouseIcon"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}
// Gets and sets a custom icon to display as the mouse pointer for the control.
// Parameters:
//   value - []byte 
// Returns:
//   void  
func (instance *UnknownControl) SetMouseIcon(value []byte)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZBL(C.CString("UnknownControl_SetMouseIcon"), instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Returns:
//   int32  
func (instance *UnknownControl) GetMousePointer()  (ControlMousePointerType,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZJ(C.CString("UnknownControl_GetMousePointer"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToControlMousePointerType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets and sets the type of icon displayed as the mouse pointer for the control.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *UnknownControl) SetMousePointer(value ControlMousePointerType)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZK(C.CString("UnknownControl_SetMousePointer"), instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the linked cell.
// Returns:
//   string  
func (instance *UnknownControl) GetLinkedCell()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("UnknownControl_GetLinkedCell"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the linked cell.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *UnknownControl) SetLinkedCell(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("UnknownControl_SetLinkedCell"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the list fill range.
// Returns:
//   string  
func (instance *UnknownControl) GetListFillRange()  (string,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZL(C.CString("UnknownControl_GetListFillRange"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the list fill range.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *UnknownControl) SetListFillRange(value string)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZM(C.CString("UnknownControl_SetListFillRange"), instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the width of the control in unit of points.
// Returns:
//   float64  
func (instance *UnknownControl) GetWidth()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("UnknownControl_GetWidth"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the width of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *UnknownControl) SetWidth(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("UnknownControl_SetWidth"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the height of the control in unit of points.
// Returns:
//   float64  
func (instance *UnknownControl) GetHeight()  (float64,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAB(C.CString("UnknownControl_GetHeight"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the height of the control in unit of points.
// Parameters:
//   value - float64 
// Returns:
//   void  
func (instance *UnknownControl) SetHeight(value float64)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZAC(C.CString("UnknownControl_SetHeight"), instance.ptr, C.double(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the foreground.
// Returns:
//   int32  
func (instance *UnknownControl) GetForeOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("UnknownControl_GetForeOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the foreground.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *UnknownControl) SetForeOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("UnknownControl_SetForeOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets and sets the ole color of the background.
// Returns:
//   int32  
func (instance *UnknownControl) GetBackOleColor()  (int32,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZD(C.CString("UnknownControl_GetBackOleColor"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets and sets the ole color of the background.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *UnknownControl) SetBackOleColor(value int32)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZE(C.CString("UnknownControl_SetBackOleColor"), instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether this control is visible.
// Returns:
//   bool  
func (instance *UnknownControl) IsVisible()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_IsVisible"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("UnknownControl_SetIsVisible"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to show a shadow.
// Returns:
//   bool  
func (instance *UnknownControl) GetShadow()  (bool,  error)  {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZB(C.CString("UnknownControl_GetShadow"), instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := bool(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.CellsGoFunctoinZZZC(C.CString("UnknownControl_SetShadow"), instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func (instance *UnknownControl) ToActiveXControl() *ActiveXControl {
	parentClass := &ActiveXControl{}
	parentClass.ptr = instance.ptr
	return parentClass
}
func (instance *UnknownControl) ToActiveXControlBase() *ActiveXControlBase {
	parentClass := &ActiveXControlBase{}
	parentClass.ptr = instance.ptr
	return parentClass
}

func DeleteUnknownControl(unknowncontrol *UnknownControl){
	runtime.SetFinalizer(unknowncontrol, nil)
	C.Delete_CObject(C.CString("Delete_UnknownControl"),unknowncontrol.ptr)
	unknowncontrol.ptr = nil
}
