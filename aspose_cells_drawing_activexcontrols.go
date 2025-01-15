// +build windows

// Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
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
	CGoReturnPtr := C.New_ActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Returns:
//   bool  
func (instance *ActiveXControl) IsEnabled()  (bool,  error)  {
	
	CGoReturnPtr := C.ActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetWorkbook( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_GetMouseIcon( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetMouseIcon( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.ActiveXControl_GetMousePointer( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetMousePointer( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ActiveXControl_GetLinkedCell( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetLinkedCell( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetListFillRange( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetListFillRange( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_GetWidth( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetWidth( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetHeight( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetHeight( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetForeOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetForeOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetBackOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControl_SetBackOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ActiveXControl_IsVisible( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControl_SetIsVisible( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControl_GetShadow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControl) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControl_SetShadow( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteActiveXControl(activexcontrol *ActiveXControl){
	runtime.SetFinalizer(activexcontrol, nil)
	C.Delete_ActiveXControl(activexcontrol.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the <see cref="Workbook"/> object.
// Returns:
//   Workbook  
func (instance *ActiveXControlBase) GetWorkbook()  (*Workbook,  error)  {
	
	CGoReturnPtr := C.ActiveXControlBase_GetWorkbook( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetMouseIcon( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetMouseIcon( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetMousePointer( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetMousePointer( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetLinkedCell( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetLinkedCell( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetListFillRange( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetListFillRange( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetWidth( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetWidth( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetHeight( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetHeight( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetForeOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetForeOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetBackOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ActiveXControlBase_SetBackOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_IsVisible( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether this control is visible.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetIsVisible(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControlBase_SetIsVisible( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetShadow( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether to show a shadow.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ActiveXControlBase) SetShadow(value bool)  error {
	
	CGoReturnPtr := C.ActiveXControlBase_SetShadow( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ActiveXControlBase_GetData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}


func DeleteActiveXControlBase(activexcontrolbase *ActiveXControlBase){
	runtime.SetFinalizer(activexcontrolbase, nil)
	C.Delete_ActiveXControlBase(activexcontrolbase.ptr)
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
	CGoReturnPtr := C.New_CheckBoxActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *CheckBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetGroupName( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetGroupName( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetAlignment( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetAlignment( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsWordWrapped( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIsWordWrapped( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetCaption( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetCaption( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetPicturePosition( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetPicturePosition( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetPicture( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetPicture( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetAccelerator( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetAccelerator( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetValue( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetValue( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsTripleState( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates how the specified control will display Null values.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsTripleState(value bool)  error {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIsTripleState( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.CheckBoxActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CheckBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CheckBoxActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteCheckBoxActiveXControl(checkboxactivexcontrol *CheckBoxActiveXControl){
	runtime.SetFinalizer(checkboxactivexcontrol, nil)
	C.Delete_CheckBoxActiveXControl(checkboxactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_ComboBoxActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ComboBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetMaxLength( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetMaxLength( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetListWidth( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetListWidth( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetBoundColumn( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetBoundColumn( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetTextColumn( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetTextColumn( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetColumnCount( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetColumnCount( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetListRows( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetListRows( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetMatchEntry( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetMatchEntry( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetDropButtonStyle( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetDropButtonStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetShowDropButtonTypeWhen( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetShowDropButtonTypeWhen( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetListStyle( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetListStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetBorderStyle( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetBorderStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetBorderOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetBorderOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsEditable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the user can type into the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsEditable(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsEditable( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetShowColumnHeads( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether column headings are displayed.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetShowColumnHeads(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetShowColumnHeads( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsDragBehaviorEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether dragging and dropping is enabled for the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsDragBehaviorEnabled(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsDragBehaviorEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetEnterFieldBehavior( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetEnterFieldBehavior( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetSelectionMargin( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the user can select a line of text by clicking in the region to the left of the text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetSelectionMargin(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetSelectionMargin( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetValue( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetValue( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetHideSelection( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether selected text in the control appears highlighted when the control does not have focus.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetHideSelection(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetHideSelection( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetColumnWidths( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetColumnWidths( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsAutoWordSelected( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsAutoWordSelected( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.ComboBoxActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ComboBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.ComboBoxActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteComboBoxActiveXControl(comboboxactivexcontrol *ComboBoxActiveXControl){
	runtime.SetFinalizer(comboboxactivexcontrol, nil)
	C.Delete_ComboBoxActiveXControl(comboboxactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_CommandButtonActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *CommandButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetCaption( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetCaption( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetPicturePosition( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetPicturePosition( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetPicture( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetPicture( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetAccelerator( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetAccelerator( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetTakeFocusOnClick( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control takes the focus when clicked.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetTakeFocusOnClick(value bool)  error {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetTakeFocusOnClick( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_IsWordWrapped( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetIsWordWrapped( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.CommandButtonActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *CommandButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.CommandButtonActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteCommandButtonActiveXControl(commandbuttonactivexcontrol *CommandButtonActiveXControl){
	runtime.SetFinalizer(commandbuttonactivexcontrol, nil)
	C.Delete_CommandButtonActiveXControl(commandbuttonactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_ImageActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ImageActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.ImageActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.ImageActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetBorderOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetBorderOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetBorderStyle( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetBorderStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetPictureSizeMode( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetPictureSizeMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetPicture( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetPicture( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetPictureAlignment( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetPictureAlignment( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_IsTiled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the picture is tiled across the background.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsTiled(value bool)  error {
	
	CGoReturnPtr := C.ImageActiveXControl_SetIsTiled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ImageActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.ImageActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ImageActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.ImageActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ImageActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.ImageActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.ImageActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ImageActiveXControl_GetData( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := C.GoBytes(unsafe.Pointer(CGoReturnPtr.return_value), C.int(CGoReturnPtr.column_length))
	 

	return result, nil 
}


func DeleteImageActiveXControl(imageactivexcontrol *ImageActiveXControl){
	runtime.SetFinalizer(imageactivexcontrol, nil)
	C.Delete_ImageActiveXControl(imageactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_LabelActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *LabelActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.LabelActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetCaption( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetCaption( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetPicturePosition( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetPicturePosition( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetBorderOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetBorderOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetBorderStyle( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetBorderStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetPicture( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetPicture( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetAccelerator( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetAccelerator( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_IsWordWrapped( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.LabelActiveXControl_SetIsWordWrapped( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.LabelActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.LabelActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.LabelActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.LabelActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.LabelActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *LabelActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.LabelActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteLabelActiveXControl(labelactivexcontrol *LabelActiveXControl){
	runtime.SetFinalizer(labelactivexcontrol, nil)
	C.Delete_LabelActiveXControl(labelactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_ListBoxActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ListBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetScrollBars( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetScrollBars( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetListWidth( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetListWidth( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetBoundColumn( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetBoundColumn( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetTextColumn( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetTextColumn( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetColumnCount( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetColumnCount( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetMatchEntry( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetMatchEntry( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetListStyle( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetListStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetSelectionType( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetSelectionType( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetValue( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetValue( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetBorderStyle( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetBorderStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetBorderOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetBorderOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetShowColumnHeads( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether column headings are displayed.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetShowColumnHeads(value bool)  error {
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetShowColumnHeads( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetIntegralHeight( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will only show complete lines of text without showing any partial lines.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIntegralHeight(value bool)  error {
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetIntegralHeight( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetColumnWidths( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetColumnWidths( instance.ptr, C.double(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.ListBoxActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ListBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.ListBoxActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteListBoxActiveXControl(listboxactivexcontrol *ListBoxActiveXControl){
	runtime.SetFinalizer(listboxactivexcontrol, nil)
	C.Delete_ListBoxActiveXControl(listboxactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_RadioButtonActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *RadioButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetGroupName( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetGroupName( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetAlignment( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetAlignment( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_IsWordWrapped( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetIsWordWrapped( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetCaption( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetCaption( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetPicturePosition( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetPicturePosition( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetPicture( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetPicture( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetAccelerator( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetAccelerator( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_GetValue( instance.ptr)
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetValue( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.RadioButtonActiveXControl_IsTripleState( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates how the specified control will display Null values.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RadioButtonActiveXControl) SetIsTripleState(value bool)  error {
	
	CGoReturnPtr := C.RadioButtonActiveXControl_SetIsTripleState( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteRadioButtonActiveXControl(radiobuttonactivexcontrol *RadioButtonActiveXControl){
	runtime.SetFinalizer(radiobuttonactivexcontrol, nil)
	C.Delete_RadioButtonActiveXControl(radiobuttonactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_ScrollBarActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ScrollBarActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetLargeChange( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_SetLargeChange( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetMin( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_SetMin( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetMax( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_SetMax( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetPosition( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_SetPosition( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetSmallChange( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_SetSmallChange( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_GetOrientation( instance.ptr)
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
	
	CGoReturnPtr := C.ScrollBarActiveXControl_SetOrientation( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteScrollBarActiveXControl(scrollbaractivexcontrol *ScrollBarActiveXControl){
	runtime.SetFinalizer(scrollbaractivexcontrol, nil)
	C.Delete_ScrollBarActiveXControl(scrollbaractivexcontrol.ptr)
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
	CGoReturnPtr := C.New_SpinButtonActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *SpinButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetMin( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetMin( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetMax( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetMax( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetPosition( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetPosition( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetSmallChange( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetSmallChange( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetOrientation( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetOrientation( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.SpinButtonActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *SpinButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.SpinButtonActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteSpinButtonActiveXControl(spinbuttonactivexcontrol *SpinButtonActiveXControl){
	runtime.SetFinalizer(spinbuttonactivexcontrol, nil)
	C.Delete_SpinButtonActiveXControl(spinbuttonactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_TextBoxActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *TextBoxActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetBorderStyle( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetBorderStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetBorderOleColor( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetBorderOleColor( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetMaxLength( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetMaxLength( instance.ptr, C.int(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetScrollBars( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetScrollBars( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetPasswordChar( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetPasswordChar( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsEditable( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the user can type into the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsEditable(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsEditable( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetIntegralHeight( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will only show complete lines of text without showing any partial lines.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIntegralHeight(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIntegralHeight( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsDragBehaviorEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether dragging and dropping is enabled for the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsDragBehaviorEnabled(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsDragBehaviorEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetEnterKeyBehavior( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetEnterKeyBehavior( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetEnterFieldBehavior( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetEnterFieldBehavior( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetTabKeyBehavior( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether tab characters are allowed in the text of the control.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetTabKeyBehavior(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetTabKeyBehavior( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetHideSelection( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether selected text in the control appears highlighted when the control does not have focus.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetHideSelection(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetHideSelection( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsAutoTab( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the focus will automatically move to the next control when the user enters the maximum number of characters.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsAutoTab(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsAutoTab( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsMultiLine( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can display more than one line of text.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsMultiLine(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsMultiLine( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsWordWrapped( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the contents of the control automatically wrap at the end of a line.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsWordWrapped(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsWordWrapped( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetText( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetText( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetDropButtonStyle( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetDropButtonStyle( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetShowDropButtonTypeWhen( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetShowDropButtonTypeWhen( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsAutoWordSelected( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsAutoWordSelected( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.TextBoxActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *TextBoxActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.TextBoxActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteTextBoxActiveXControl(textboxactivexcontrol *TextBoxActiveXControl){
	runtime.SetFinalizer(textboxactivexcontrol, nil)
	C.Delete_TextBoxActiveXControl(textboxactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_ToggleButtonActiveXControl(src.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the type of the ActiveX control.
// Returns:
//   int32  
func (instance *ToggleButtonActiveXControl) GetType()  (ControlType,  error)  {
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetCaption( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetCaption( instance.ptr, C.CString(value))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetPicturePosition( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetPicturePosition( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetSpecialEffect( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetSpecialEffect( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetPicture( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetPicture( instance.ptr, unsafe.Pointer(&value[0]), C.int( len(value)))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetAccelerator( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetAccelerator( instance.ptr, C.char(value))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetValue( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetValue( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_IsTripleState( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates how the specified control will display Null values.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsTripleState(value bool)  error {
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetIsTripleState( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ToggleButtonActiveXControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.ToggleButtonActiveXControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteToggleButtonActiveXControl(togglebuttonactivexcontrol *ToggleButtonActiveXControl){
	runtime.SetFinalizer(togglebuttonactivexcontrol, nil)
	C.Delete_ToggleButtonActiveXControl(togglebuttonactivexcontrol.ptr)
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
	CGoReturnPtr := C.New_UnknownControl(src.ptr)
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
	
	CGoReturnPtr := C.UnknownControl_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the related data.
// Parameters:
//   relId - string 
// Returns:
//   []byte  
func (instance *UnknownControl) GetRelationshipData(relid string)  ([]byte,  error)  {
	
	CGoReturnPtr := C.UnknownControl_GetRelationshipData( instance.ptr, C.CString(relid))
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
	
	CGoReturnPtr := C.UnknownControl_GetData( instance.ptr)
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
	
	CGoReturnPtr := C.UnknownControl_GetType( instance.ptr)
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
	
	CGoReturnPtr := C.UnknownControl_IsEnabled( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control can receive the focus and respond to user-generated events.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsEnabled(value bool)  error {
	
	CGoReturnPtr := C.UnknownControl_SetIsEnabled( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.UnknownControl_IsLocked( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether data in the control is locked for editing.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsLocked(value bool)  error {
	
	CGoReturnPtr := C.UnknownControl_SetIsLocked( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.UnknownControl_IsTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control is transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsTransparent(value bool)  error {
	
	CGoReturnPtr := C.UnknownControl_SetIsTransparent( instance.ptr, C.bool(value))
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
	
	CGoReturnPtr := C.UnknownControl_GetIMEMode( instance.ptr)
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
	
	CGoReturnPtr := C.UnknownControl_SetIMEMode( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.UnknownControl_GetFont( instance.ptr)
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
	
	CGoReturnPtr := C.UnknownControl_GetTextAlign( instance.ptr)
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
	
	CGoReturnPtr := C.UnknownControl_SetTextAlign( instance.ptr, C.int( int32(value)))
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
	
	CGoReturnPtr := C.UnknownControl_IsAutoSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the control will automatically resize to display its entire contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *UnknownControl) SetIsAutoSize(value bool)  error {
	
	CGoReturnPtr := C.UnknownControl_SetIsAutoSize( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteUnknownControl(unknowncontrol *UnknownControl){
	runtime.SetFinalizer(unknowncontrol, nil)
	C.Delete_UnknownControl(unknowncontrol.ptr)
	unknowncontrol.ptr = nil
}
