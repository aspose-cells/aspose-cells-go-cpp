// +build linux

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

/**************Enum ColorDepth *****************/
// Enumerates Bit Depth Type for tiff image.
type ColorDepth int32

const(
// Default value, not set value.
ColorDepth_Default ColorDepth = 0 

// 1 bit per pixel
ColorDepth_Format1bpp ColorDepth = 1 

// 4 bits per pixel
ColorDepth_Format4bpp ColorDepth = 4 

// 8 bits per pixel
ColorDepth_Format8bpp ColorDepth = 8 

// 24 bits per pixel
ColorDepth_Format24bpp ColorDepth = 24 

// 32 bits per pixel
ColorDepth_Format32bpp ColorDepth = 32 
)

func Int32ToColorDepth(value int32)(ColorDepth ,error){
	switch value {
		case 0:  return ColorDepth_Default, nil  
		case 1:  return ColorDepth_Format1bpp, nil  
		case 4:  return ColorDepth_Format4bpp, nil  
		case 8:  return ColorDepth_Format8bpp, nil  
		case 24:  return ColorDepth_Format24bpp, nil  
		case 32:  return ColorDepth_Format32bpp, nil  
		default:
			return 0 ,fmt.Errorf("invalid ColorDepth value: %d", value)
	}
}

/**************Enum CommentTitleType *****************/
// Represents comment title type while rendering when comment is set to display at end of sheet.
type CommentTitleType int32

const(
// Represents comment title cell.
CommentTitleType_Cell CommentTitleType = 0 

// Represents comment title comment.
CommentTitleType_Comment CommentTitleType = 1 

// Represents comment title note.
CommentTitleType_Note CommentTitleType = 2 

// Represents comment title reply.
CommentTitleType_Reply CommentTitleType = 3 
)

func Int32ToCommentTitleType(value int32)(CommentTitleType ,error){
	switch value {
		case 0:  return CommentTitleType_Cell, nil  
		case 1:  return CommentTitleType_Comment, nil  
		case 2:  return CommentTitleType_Note, nil  
		case 3:  return CommentTitleType_Reply, nil  
		default:
			return 0 ,fmt.Errorf("invalid CommentTitleType value: %d", value)
	}
}

/**************Enum DrawObjectEnum *****************/
// Indicate Cell or Image of DrawObject.
type DrawObjectEnum int32

const(
// Indicate DrawObject is an Image
DrawObjectEnum_Image DrawObjectEnum = 0 

// indicate DrawObject is an Cell
DrawObjectEnum_Cell DrawObjectEnum = 1 
)

func Int32ToDrawObjectEnum(value int32)(DrawObjectEnum ,error){
	switch value {
		case 0:  return DrawObjectEnum_Image, nil  
		case 1:  return DrawObjectEnum_Cell, nil  
		default:
			return 0 ,fmt.Errorf("invalid DrawObjectEnum value: %d", value)
	}
}

/**************Enum ImageBinarizationMethod *****************/
// Specifies the method used to binarize image.
type ImageBinarizationMethod int32

const(
// Specifies threshold method.
ImageBinarizationMethod_Threshold ImageBinarizationMethod = 0 

// Specifies dithering using Floyd-Steinberg error diffusion method.
ImageBinarizationMethod_FloydSteinbergDithering ImageBinarizationMethod = 1 
)

func Int32ToImageBinarizationMethod(value int32)(ImageBinarizationMethod ,error){
	switch value {
		case 0:  return ImageBinarizationMethod_Threshold, nil  
		case 1:  return ImageBinarizationMethod_FloydSteinbergDithering, nil  
		default:
			return 0 ,fmt.Errorf("invalid ImageBinarizationMethod value: %d", value)
	}
}

/**************Enum PdfCompliance *****************/
// Allowing user to set PDF conversion's Compatibility
type PdfCompliance int32

const(
// Pdf format compatible with PDF 1.4
PdfCompliance_Pdf14 PdfCompliance = 0 

// Pdf format compatible with PDF 1.5
PdfCompliance_Pdf15 PdfCompliance = 3 

// Pdf format compatible with PDF 1.6
PdfCompliance_Pdf16 PdfCompliance = 4 

// Pdf format compatible with PDF 1.7
PdfCompliance_Pdf17 PdfCompliance = 5 

// Pdf format compatible with PDF/A-1b(ISO 19005-1)
PdfCompliance_PdfA1b PdfCompliance = 1 

// Pdf format compatible with PDF/A-1a(ISO 19005-1)
PdfCompliance_PdfA1a PdfCompliance = 2 

// Pdf format compatible with PDF/A-2b(ISO 19005-2)
PdfCompliance_PdfA2b PdfCompliance = 6 

// Pdf format compatible with PDF/A-2u(ISO 19005-2)
PdfCompliance_PdfA2u PdfCompliance = 7 

// Pdf format compatible with PDF/A-2a(ISO 19005-2)
PdfCompliance_PdfA2a PdfCompliance = 8 

// Pdf format compatible with PDF/A-3b(ISO 19005-3)
PdfCompliance_PdfA3b PdfCompliance = 9 

// Pdf format compatible with PDF/A-3u(ISO 19005-3)
PdfCompliance_PdfA3u PdfCompliance = 10 

// Pdf format compatible with PDF/A-3a(ISO 19005-3)
PdfCompliance_PdfA3a PdfCompliance = 11 
)

func Int32ToPdfCompliance(value int32)(PdfCompliance ,error){
	switch value {
		case 0:  return PdfCompliance_Pdf14, nil  
		case 3:  return PdfCompliance_Pdf15, nil  
		case 4:  return PdfCompliance_Pdf16, nil  
		case 5:  return PdfCompliance_Pdf17, nil  
		case 1:  return PdfCompliance_PdfA1b, nil  
		case 2:  return PdfCompliance_PdfA1a, nil  
		case 6:  return PdfCompliance_PdfA2b, nil  
		case 7:  return PdfCompliance_PdfA2u, nil  
		case 8:  return PdfCompliance_PdfA2a, nil  
		case 9:  return PdfCompliance_PdfA3b, nil  
		case 10:  return PdfCompliance_PdfA3u, nil  
		case 11:  return PdfCompliance_PdfA3a, nil  
		default:
			return 0 ,fmt.Errorf("invalid PdfCompliance value: %d", value)
	}
}

/**************Enum PdfCompressionCore *****************/
// Specifies a type of compression applied to all content in the PDF file except images.
type PdfCompressionCore int32

const(
// None
PdfCompressionCore_None PdfCompressionCore = 0 

// Rle
PdfCompressionCore_Rle PdfCompressionCore = 1 

// Lzw
PdfCompressionCore_Lzw PdfCompressionCore = 2 

// Flate
PdfCompressionCore_Flate PdfCompressionCore = 3 
)

func Int32ToPdfCompressionCore(value int32)(PdfCompressionCore ,error){
	switch value {
		case 0:  return PdfCompressionCore_None, nil  
		case 1:  return PdfCompressionCore_Rle, nil  
		case 2:  return PdfCompressionCore_Lzw, nil  
		case 3:  return PdfCompressionCore_Flate, nil  
		default:
			return 0 ,fmt.Errorf("invalid PdfCompressionCore value: %d", value)
	}
}

/**************Enum PdfCustomPropertiesExport *****************/
// Specifies the way <see cref="CustomDocumentPropertyCollection"/> are exported to PDF file.
type PdfCustomPropertiesExport int32

const(
// No custom properties are exported.
PdfCustomPropertiesExport_None PdfCustomPropertiesExport = 0 

// Custom properties are exported as entries in Info dictionary.
PdfCustomPropertiesExport_Standard PdfCustomPropertiesExport = 1 
)

func Int32ToPdfCustomPropertiesExport(value int32)(PdfCustomPropertiesExport ,error){
	switch value {
		case 0:  return PdfCustomPropertiesExport_None, nil  
		case 1:  return PdfCustomPropertiesExport_Standard, nil  
		default:
			return 0 ,fmt.Errorf("invalid PdfCustomPropertiesExport value: %d", value)
	}
}

/**************Enum PdfFontEncoding *****************/
// Represents pdf embedded font encoding.
type PdfFontEncoding int32

const(
// Represents use Identity-H encoding for all embedded fonts in pdf.
PdfFontEncoding_Identity PdfFontEncoding = 0 

// Represents prefer to use WinAnsiEncoding for TrueType fonts with characters 32-127,
// otherwise, Identity-H encoding will be used for embedded fonts in pdf.
PdfFontEncoding_AnsiPrefer PdfFontEncoding = 1 
)

func Int32ToPdfFontEncoding(value int32)(PdfFontEncoding ,error){
	switch value {
		case 0:  return PdfFontEncoding_Identity, nil  
		case 1:  return PdfFontEncoding_AnsiPrefer, nil  
		default:
			return 0 ,fmt.Errorf("invalid PdfFontEncoding value: %d", value)
	}
}

/**************Enum PdfOptimizationType *****************/
// Specifies a type of optimization.
type PdfOptimizationType int32

const(
// High print quality
PdfOptimizationType_Standard PdfOptimizationType = 0 

// File size is more important than print quality
PdfOptimizationType_MinimumSize PdfOptimizationType = 1 
)

func Int32ToPdfOptimizationType(value int32)(PdfOptimizationType ,error){
	switch value {
		case 0:  return PdfOptimizationType_Standard, nil  
		case 1:  return PdfOptimizationType_MinimumSize, nil  
		default:
			return 0 ,fmt.Errorf("invalid PdfOptimizationType value: %d", value)
	}
}

/**************Enum TiffCompression *****************/
// Specifies what type of compression to apply when saving images into TIFF format file.
type TiffCompression int32

const(
// Specifies no compression.
TiffCompression_CompressionNone TiffCompression = 0 

// Specifies the RLE compression scheme.
TiffCompression_CompressionRle TiffCompression = 1 

// Specifies the LZW compression scheme.
TiffCompression_CompressionLZW TiffCompression = 2 

// Specifies the CCITT3 compression scheme.
TiffCompression_CompressionCCITT3 TiffCompression = 3 

// Specifies the CCITT4 compression scheme.
TiffCompression_CompressionCCITT4 TiffCompression = 4 
)

func Int32ToTiffCompression(value int32)(TiffCompression ,error){
	switch value {
		case 0:  return TiffCompression_CompressionNone, nil  
		case 1:  return TiffCompression_CompressionRle, nil  
		case 2:  return TiffCompression_CompressionLZW, nil  
		case 3:  return TiffCompression_CompressionCCITT3, nil  
		case 4:  return TiffCompression_CompressionCCITT4, nil  
		default:
			return 0 ,fmt.Errorf("invalid TiffCompression value: %d", value)
	}
}
// Class DrawObject 

// DrawObject will be initialized and returned when rendering.
type DrawObject struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *DrawObject) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.DrawObject_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates the Cell object when rendering.
// All properties of cell can be accessed.
// Returns:
//   Cell  
func (instance *DrawObject) GetCell()  (*Cell,  error)  {
	CGoReturnPtr := C.DrawObject_GetCell( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Cell{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCell) 

	return result, nil 
}
// Indicates the Shape object when rendering.
// All properties of shape can be accessed.
// Returns:
//   Shape  
func (instance *DrawObject) GetShape()  (*Shape,  error)  {
	CGoReturnPtr := C.DrawObject_GetShape( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Shape{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteShape) 

	return result, nil 
}
// Indicates the type of DrawObject.
// Returns:
//   int32  
func (instance *DrawObject) GetType()  (DrawObjectEnum,  error)  {
	CGoReturnPtr := C.DrawObject_GetType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDrawObjectEnum(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates the page index of DrawObject.
// Page index is based on zero.
// One Sheet contains several pages when rendering.
// Returns:
//   int32  
func (instance *DrawObject) GetCurrentPage()  (int32,  error)  {
	CGoReturnPtr := C.DrawObject_GetCurrentPage( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates total pages in current rendering.
// Returns:
//   int32  
func (instance *DrawObject) GetTotalPages()  (int32,  error)  {
	CGoReturnPtr := C.DrawObject_GetTotalPages( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicates current sheet index of DrawObject.
// Returns:
//   int32  
func (instance *DrawObject) GetSheetIndex()  (int32,  error)  {
	CGoReturnPtr := C.DrawObject_GetSheetIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteDrawObject(drawobject *DrawObject){
	runtime.SetFinalizer(drawobject, nil)
	C.Delete_DrawObject(drawobject.ptr)
	drawobject.ptr = nil
}

// Class DrawObjectEventHandler 

// Interface to get DrawObject and Bound when rendering.
type DrawObjectEventHandler struct {
	ptr unsafe.Pointer
}


// Implements this interface to get DrawObject and Bound when rendering.
// Parameters:
//   drawObject - DrawObject 
//   x - float32 
//   y - float32 
//   width - float32 
//   height - float32 
// Returns:
//   void  
func (instance *DrawObjectEventHandler) Draw(drawobject *DrawObject, x float32, y float32, width float32, height float32)  error {
	CGoReturnPtr := C.DrawObjectEventHandler_Draw( instance.ptr, drawobject.ptr, C.float(x), C.float(y), C.float(width), C.float(height))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteDrawObjectEventHandler(drawobjecteventhandler *DrawObjectEventHandler){
	runtime.SetFinalizer(drawobjecteventhandler, nil)
	C.Delete_DrawObjectEventHandler(drawobjecteventhandler.ptr)
	drawobjecteventhandler.ptr = nil
}

// Class ImageOrPrintOptions 

// Allows to specify options when rendering worksheet to images, printing worksheet or rendering chart to image.
type ImageOrPrintOptions struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewImageOrPrintOptions() ( *ImageOrPrintOptions, error) {
	imageorprintoptions := &ImageOrPrintOptions{}
	CGoReturnPtr := C.New_ImageOrPrintOptions()
	if CGoReturnPtr.error_no == 0 {
		imageorprintoptions.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(imageorprintoptions, DeleteImageOrPrintOptions)
		return imageorprintoptions, nil
	} else {
		imageorprintoptions.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return imageorprintoptions, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// If PrintWithStatusDialog = true , there will be a dialog that shows current print status.
// else no such dialog will show.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetPrintWithStatusDialog(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetPrintWithStatusDialog( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If PrintWithStatusDialog = true , there will be a dialog that shows current print status.
// else no such dialog will show.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetPrintWithStatusDialog()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetPrintWithStatusDialog( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets the horizontal resolution for generated images, in dots per inch.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetHorizontalResolution()  (int32,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetHorizontalResolution( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the horizontal resolution for generated images, in dots per inch.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetHorizontalResolution(value int32)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetHorizontalResolution( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the vertical resolution for generated images, in dots per inch.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetVerticalResolution()  (int32,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetVerticalResolution( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the vertical resolution for generated images, in dots per inch.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetVerticalResolution(value int32)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetVerticalResolution( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the type of compression to apply only when saving pages to the <c>Tiff</c> format.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetTiffCompression()  (TiffCompression,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetTiffCompression( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTiffCompression(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the type of compression to apply only when saving pages to the <c>Tiff</c> format.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetTiffCompression(value TiffCompression)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetTiffCompression( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets bit depth to apply only when saving pages to the <c>Tiff</c> format.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetTiffColorDepth()  (ColorDepth,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetTiffColorDepth( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToColorDepth(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets bit depth to apply only when saving pages to the <c>Tiff</c> format.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetTiffColorDepth(value ColorDepth)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetTiffColorDepth( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets method used while converting images to 1 bpp format
// when <see cref="ImageType"/> is Tiff and <see cref="TiffCompression"/> is equal to Ccitt3 or Ccitt4.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetTiffBinarizationMethod()  (ImageBinarizationMethod,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetTiffBinarizationMethod( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToImageBinarizationMethod(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets method used while converting images to 1 bpp format
// when <see cref="ImageType"/> is Tiff and <see cref="TiffCompression"/> is equal to Ccitt3 or Ccitt4.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetTiffBinarizationMethod(value ImageBinarizationMethod)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetTiffBinarizationMethod( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates which pages will not be printed.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetPrintingPage()  (PrintingPageType,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetPrintingPage( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToPrintingPageType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Indicates which pages will not be printed.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetPrintingPage(value PrintingPageType)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetPrintingPage( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets a value determining the quality of the generated  images
// to apply only when saving pages to the <c>Jpeg</c> format. The default value is 100
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetQuality()  (int32,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetQuality( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets a value determining the quality of the generated  images
// to apply only when saving pages to the <c>Jpeg</c> format. The default value is 100
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetQuality(value int32)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetQuality( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the format of the generated images.
// default value: PNG.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetImageType()  (ImageType,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetImageType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToImageType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets the format of the generated images.
// default value: PNG.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetImageType(value ImageType)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetImageType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If OnePagePerSheet is true , all content of one sheet will output to only one page in result.
// The paper size of pagesetup will be invalid, and the other settings of pagesetup
// will still take effect.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetOnePagePerSheet()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetOnePagePerSheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// If OnePagePerSheet is true , all content of one sheet will output to only one page in result.
// The paper size of pagesetup will be invalid, and the other settings of pagesetup
// will still take effect.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetOnePagePerSheet(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetOnePagePerSheet( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If AllColumnsInOnePagePerSheet is true , all column content of one sheet will output to only one page in result.
// The width of paper size of pagesetup will be invalid, and the other settings of pagesetup
// will still take effect.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetAllColumnsInOnePagePerSheet()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetAllColumnsInOnePagePerSheet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// If AllColumnsInOnePagePerSheet is true , all column content of one sheet will output to only one page in result.
// The width of paper size of pagesetup will be invalid, and the other settings of pagesetup
// will still take effect.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetAllColumnsInOnePagePerSheet(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetAllColumnsInOnePagePerSheet( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Implements this interface to get DrawObject and Bound when rendering.
// Returns:
//   DrawObjectEventHandler  
func (instance *ImageOrPrintOptions) GetDrawObjectEventHandler()  (*DrawObjectEventHandler,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetDrawObjectEventHandler( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &DrawObjectEventHandler{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteDrawObjectEventHandler) 

	return result, nil 
}
// Implements this interface to get DrawObject and Bound when rendering.
// Parameters:
//   value - DrawObjectEventHandler 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetDrawObjectEventHandler(value *DrawObjectEventHandler)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetDrawObjectEventHandler( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicate the filename of embedded image in svg.
// This should be full path with directory like "c:\\xpsEmbedded"
// Returns:
//   string  
func (instance *ImageOrPrintOptions) GetEmbededImageNameInSvg()  (string,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetEmbededImageNameInSvg( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Indicate the filename of embedded image in svg.
// This should be full path with directory like "c:\\xpsEmbedded"
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetEmbededImageNameInSvg(value string)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetEmbededImageNameInSvg( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// if this property is true, the generated svg will fit to view port.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetSVGFitToViewPort()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetSVGFitToViewPort( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// if this property is true, the generated svg will fit to view port.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetSVGFitToViewPort(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetSVGFitToViewPort( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// If this property is true , one Area will be output, and no scale will take effect.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetOnlyArea()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetOnlyArea( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// If this property is true , one Area will be output, and no scale will take effect.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetOnlyArea(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetOnlyArea( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates if the background of generated image should be transparent.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetTransparent()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetTransparent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates if the background of generated image should be transparent.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetTransparent(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetTransparent( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to only substitute the font of character when the cell font is not compatibility for it.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetIsFontSubstitutionCharGranularity(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetIsFontSubstitutionCharGranularity( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to only substitute the font of character when the cell font is not compatibility for it.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) IsFontSubstitutionCharGranularity()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_IsFontSubstitutionCharGranularity( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets the 0-based index of the first page to save.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetPageIndex(value int32)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetPageIndex( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the 0-based index of the first page to save.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetPageIndex()  (int32,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetPageIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets the number of pages to save.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetPageCount(value int32)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetPageCount( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the number of pages to save.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetPageCount()  (int32,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Sets desired width and height of image.
// Parameters:
//   desiredWidth - int32 
//   desiredHeight - int32 
//   keepAspectRatio - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetDesiredSize(desiredwidth int32, desiredheight int32, keepaspectratio bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetDesiredSize( instance.ptr, C.int(desiredwidth), C.int(desiredheight), C.bool(keepaspectratio))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to optimize the output elements.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) IsOptimized()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_IsOptimized( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether to optimize the output elements.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetIsOptimized(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetIsOptimized( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// When characters in the Excel are Unicode and not be set with correct font in cell style,
// They may appear as block in pdf,image.
// Set the DefaultFont such as MingLiu or MS Gothic to show these characters.
// If this property is not set, Aspose.Cells will use system default font to show these unicode characters.
// Returns:
//   string  
func (instance *ImageOrPrintOptions) GetDefaultFont()  (string,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetDefaultFont( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// When characters in the Excel are Unicode and not be set with correct font in cell style,
// They may appear as block in pdf,image.
// Set the DefaultFont such as MingLiu or MS Gothic to show these characters.
// If this property is not set, Aspose.Cells will use system default font to show these unicode characters.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetDefaultFont(value string)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetDefaultFont( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// When characters in the Excel are Unicode and not be set with correct font in cell style,
// They may appear as block in pdf,image.
// Set this to true to try to use workbook's default font to show these characters first.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetCheckWorkbookDefaultFont()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetCheckWorkbookDefaultFont( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// When characters in the Excel are Unicode and not be set with correct font in cell style,
// They may appear as block in pdf,image.
// Set this to true to try to use workbook's default font to show these characters first.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetCheckWorkbookDefaultFont(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetCheckWorkbookDefaultFont( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether to output a blank page when there is nothing to print.
// Returns:
//   bool  
func (instance *ImageOrPrintOptions) GetOutputBlankPageWhenNothingToPrint()  (bool,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetOutputBlankPageWhenNothingToPrint( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether to output a blank page when there is nothing to print.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetOutputBlankPageWhenNothingToPrint(value bool)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetOutputBlankPageWhenNothingToPrint( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets gridline type.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetGridlineType()  (GridlineType,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetGridlineType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToGridlineType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets gridline type.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetGridlineType(value GridlineType)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetGridlineType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets displaying text type when the text width is larger than cell width.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetTextCrossType()  (TextCrossType,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetTextCrossType( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToTextCrossType(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets displaying text type when the text width is larger than cell width.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetTextCrossType(value TextCrossType)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetTextCrossType( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets default edit language.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetDefaultEditLanguage()  (DefaultEditLanguage,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetDefaultEditLanguage( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToDefaultEditLanguage(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Gets or sets default edit language.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetDefaultEditLanguage(value DefaultEditLanguage)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetDefaultEditLanguage( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets the sheets to render. Default is all visible sheets in the workbook: <see cref="Aspose.Cells.Rendering.SheetSet.Visible"/>.
// Returns:
//   SheetSet  
func (instance *ImageOrPrintOptions) GetSheetSet()  (*SheetSet,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetSheetSet( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SheetSet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSheetSet) 

	return result, nil 
}
// Gets or sets the sheets to render. Default is all visible sheets in the workbook: <see cref="Aspose.Cells.Rendering.SheetSet.Visible"/>.
// Parameters:
//   value - SheetSet 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetSheetSet(value *SheetSet)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetSheetSet( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Setting for rendering Emf metafile.
// Returns:
//   int32  
func (instance *ImageOrPrintOptions) GetEmfRenderSetting()  (EmfRenderSetting,  error)  {
	CGoReturnPtr := C.ImageOrPrintOptions_GetEmfRenderSetting( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result , err := Int32ToEmfRenderSetting(int32(CGoReturnPtr.return_value)) 
	if err != nil {
		return 0, err
	}

	return result, nil 
}
// Setting for rendering Emf metafile.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *ImageOrPrintOptions) SetEmfRenderSetting(value EmfRenderSetting)  error {
	CGoReturnPtr := C.ImageOrPrintOptions_SetEmfRenderSetting( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteImageOrPrintOptions(imageorprintoptions *ImageOrPrintOptions){
	runtime.SetFinalizer(imageorprintoptions, nil)
	C.Delete_ImageOrPrintOptions(imageorprintoptions.ptr)
	imageorprintoptions.ptr = nil
}

// Class PageEndSavingArgs 

// Info for a page ends saving process.
type PageEndSavingArgs struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PageSavingArgs 
func NewPageEndSavingArgs(src *PageSavingArgs) ( *PageEndSavingArgs, error) {
	pageendsavingargs := &PageEndSavingArgs{}
	CGoReturnPtr := C.New_PageEndSavingArgs(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		pageendsavingargs.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pageendsavingargs, DeletePageEndSavingArgs)
		return pageendsavingargs, nil
	} else {
		pageendsavingargs.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pageendsavingargs, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PageEndSavingArgs) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PageEndSavingArgs_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets a value indicating whether having more pages to be output.
// The default value is true.
// Returns:
//   bool  
func (instance *PageEndSavingArgs) GetHasMorePages()  (bool,  error)  {
	CGoReturnPtr := C.PageEndSavingArgs_GetHasMorePages( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets a value indicating whether having more pages to be output.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PageEndSavingArgs) SetHasMorePages(value bool)  error {
	CGoReturnPtr := C.PageEndSavingArgs_SetHasMorePages( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Current page index, zero based.
// Returns:
//   int32  
func (instance *PageEndSavingArgs) GetPageIndex()  (int32,  error)  {
	CGoReturnPtr := C.PageEndSavingArgs_GetPageIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Total page count.
// Returns:
//   int32  
func (instance *PageEndSavingArgs) GetPageCount()  (int32,  error)  {
	CGoReturnPtr := C.PageEndSavingArgs_GetPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePageEndSavingArgs(pageendsavingargs *PageEndSavingArgs){
	runtime.SetFinalizer(pageendsavingargs, nil)
	C.Delete_PageEndSavingArgs(pageendsavingargs.ptr)
	pageendsavingargs.ptr = nil
}

// Class PageSavingArgs 

// Info for a page saving process.
type PageSavingArgs struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PageSavingArgs) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PageSavingArgs_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Current page index, zero based.
// Returns:
//   int32  
func (instance *PageSavingArgs) GetPageIndex()  (int32,  error)  {
	CGoReturnPtr := C.PageSavingArgs_GetPageIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Total page count.
// Returns:
//   int32  
func (instance *PageSavingArgs) GetPageCount()  (int32,  error)  {
	CGoReturnPtr := C.PageSavingArgs_GetPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePageSavingArgs(pagesavingargs *PageSavingArgs){
	runtime.SetFinalizer(pagesavingargs, nil)
	C.Delete_PageSavingArgs(pagesavingargs.ptr)
	pagesavingargs.ptr = nil
}

// Class PageStartSavingArgs 

// Info for a page starts saving process.
type PageStartSavingArgs struct {
	ptr unsafe.Pointer
}

// Constructs from a parent object.
// Parameters:
//   src - PageSavingArgs 
func NewPageStartSavingArgs(src *PageSavingArgs) ( *PageStartSavingArgs, error) {
	pagestartsavingargs := &PageStartSavingArgs{}
	CGoReturnPtr := C.New_PageStartSavingArgs(src.ptr)
	if CGoReturnPtr.error_no == 0 {
		pagestartsavingargs.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pagestartsavingargs, DeletePageStartSavingArgs)
		return pagestartsavingargs, nil
	} else {
		pagestartsavingargs.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pagestartsavingargs, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PageStartSavingArgs) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PageStartSavingArgs_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets a value indicating whether the page should be output.
// The default value is true.
// Returns:
//   bool  
func (instance *PageStartSavingArgs) IsToOutput()  (bool,  error)  {
	CGoReturnPtr := C.PageStartSavingArgs_IsToOutput( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets a value indicating whether the page should be output.
// The default value is true.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PageStartSavingArgs) SetIsToOutput(value bool)  error {
	CGoReturnPtr := C.PageStartSavingArgs_SetIsToOutput( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Current page index, zero based.
// Returns:
//   int32  
func (instance *PageStartSavingArgs) GetPageIndex()  (int32,  error)  {
	CGoReturnPtr := C.PageStartSavingArgs_GetPageIndex( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Total page count.
// Returns:
//   int32  
func (instance *PageStartSavingArgs) GetPageCount()  (int32,  error)  {
	CGoReturnPtr := C.PageStartSavingArgs_GetPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeletePageStartSavingArgs(pagestartsavingargs *PageStartSavingArgs){
	runtime.SetFinalizer(pagestartsavingargs, nil)
	C.Delete_PageStartSavingArgs(pagestartsavingargs.ptr)
	pagestartsavingargs.ptr = nil
}

// Class PdfBookmarkEntry 

// PdfBookmarkEntry is an entry in pdf bookmark.
// if Text property of current instance is null or "",
// current instance will be hidden and children will be inserted on current level.
type PdfBookmarkEntry struct {
	ptr unsafe.Pointer
}

// Default constructor.
func NewPdfBookmarkEntry() ( *PdfBookmarkEntry, error) {
	pdfbookmarkentry := &PdfBookmarkEntry{}
	CGoReturnPtr := C.New_PdfBookmarkEntry()
	if CGoReturnPtr.error_no == 0 {
		pdfbookmarkentry.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(pdfbookmarkentry, DeletePdfBookmarkEntry)
		return pdfbookmarkentry, nil
	} else {
		pdfbookmarkentry.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return pdfbookmarkentry, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *PdfBookmarkEntry) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.PdfBookmarkEntry_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Title of a bookmark.
// Returns:
//   string  
func (instance *PdfBookmarkEntry) GetText()  (string,  error)  {
	CGoReturnPtr := C.PdfBookmarkEntry_GetText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Title of a bookmark.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PdfBookmarkEntry) SetText(value string)  error {
	CGoReturnPtr := C.PdfBookmarkEntry_SetText( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// The cell to which the bookmark link.
// Returns:
//   Cell  
func (instance *PdfBookmarkEntry) GetDestination()  (*Cell,  error)  {
	CGoReturnPtr := C.PdfBookmarkEntry_GetDestination( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Cell{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteCell) 

	return result, nil 
}
// The cell to which the bookmark link.
// Parameters:
//   value - Cell 
// Returns:
//   void  
func (instance *PdfBookmarkEntry) SetDestination(value *Cell)  error {
	CGoReturnPtr := C.PdfBookmarkEntry_SetDestination( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets name of destination.
// Returns:
//   string  
func (instance *PdfBookmarkEntry) GetDestinationName()  (string,  error)  {
	CGoReturnPtr := C.PdfBookmarkEntry_GetDestinationName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets name of destination.
// Parameters:
//   value - string 
// Returns:
//   void  
func (instance *PdfBookmarkEntry) SetDestinationName(value string)  error {
	CGoReturnPtr := C.PdfBookmarkEntry_SetDestinationName( instance.ptr, C.CString(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// When this property is true, the bookmarkentry will expand, otherwise it will collapse.
// Returns:
//   bool  
func (instance *PdfBookmarkEntry) IsOpen()  (bool,  error)  {
	CGoReturnPtr := C.PdfBookmarkEntry_IsOpen( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// When this property is true, the bookmarkentry will expand, otherwise it will collapse.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfBookmarkEntry) SetIsOpen(value bool)  error {
	CGoReturnPtr := C.PdfBookmarkEntry_SetIsOpen( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// When this property is true, the bookmarkentry will collapse, otherwise it will expand.
// Returns:
//   bool  
func (instance *PdfBookmarkEntry) IsCollapse()  (bool,  error)  {
	CGoReturnPtr := C.PdfBookmarkEntry_IsCollapse( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// When this property is true, the bookmarkentry will collapse, otherwise it will expand.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *PdfBookmarkEntry) SetIsCollapse(value bool)  error {
	CGoReturnPtr := C.PdfBookmarkEntry_SetIsCollapse( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeletePdfBookmarkEntry(pdfbookmarkentry *PdfBookmarkEntry){
	runtime.SetFinalizer(pdfbookmarkentry, nil)
	C.Delete_PdfBookmarkEntry(pdfbookmarkentry.ptr)
	pdfbookmarkentry.ptr = nil
}

// Class RenderingFont 

// Font for rendering.
type RenderingFont struct {
	ptr unsafe.Pointer
}

// Initializes a new instance of the <see cref="RenderingFont"/>
// Parameters:
//   fontName - string 
//   fontSize - float32 
func NewRenderingFont(fontname string, fontsize float32) ( *RenderingFont, error) {
	renderingfont := &RenderingFont{}
	CGoReturnPtr := C.New_RenderingFont(C.CString(fontname), C.float(fontsize))
	if CGoReturnPtr.error_no == 0 {
		renderingfont.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(renderingfont, DeleteRenderingFont)
		return renderingfont, nil
	} else {
		renderingfont.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return renderingfont, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RenderingFont) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RenderingFont_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets name of the font.
// Returns:
//   string  
func (instance *RenderingFont) GetName()  (string,  error)  {
	CGoReturnPtr := C.RenderingFont_GetName( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets size of the font in points.
// Returns:
//   float32  
func (instance *RenderingFont) GetSize()  (float32,  error)  {
	CGoReturnPtr := C.RenderingFont_GetSize( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets bold for the font.
// Returns:
//   bool  
func (instance *RenderingFont) GetBold()  (bool,  error)  {
	CGoReturnPtr := C.RenderingFont_GetBold( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets bold for the font.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RenderingFont) SetBold(value bool)  error {
	CGoReturnPtr := C.RenderingFont_SetBold( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets italic for the font.
// Returns:
//   bool  
func (instance *RenderingFont) GetItalic()  (bool,  error)  {
	CGoReturnPtr := C.RenderingFont_GetItalic( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets italic for the font.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RenderingFont) SetItalic(value bool)  error {
	CGoReturnPtr := C.RenderingFont_SetItalic( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets color for the font.
// Returns:
//   Color  
func (instance *RenderingFont) GetColor()  (*Color,  error)  {
	CGoReturnPtr := C.RenderingFont_GetColor( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &Color{}
	result.ptr = CGoReturnPtr.return_value 

	return result, nil 
}
// Gets or sets color for the font.
// Parameters:
//   value - Color 
// Returns:
//   void  
func (instance *RenderingFont) SetColor(value *Color)  error {
	CGoReturnPtr := C.RenderingFont_SetColor( instance.ptr, value.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteRenderingFont(renderingfont *RenderingFont){
	runtime.SetFinalizer(renderingfont, nil)
	C.Delete_RenderingFont(renderingfont.ptr)
	renderingfont.ptr = nil
}

// Class RenderingWatermark 

// Watermark for rendering.
type RenderingWatermark struct {
	ptr unsafe.Pointer
}

// Creates instance of text watermark.
// Parameters:
//   text - string 
//   renderingFont - RenderingFont 
func NewRenderingWatermark(text string, renderingfont *RenderingFont) ( *RenderingWatermark, error) {
	renderingwatermark := &RenderingWatermark{}
	CGoReturnPtr := C.New_RenderingWatermark(C.CString(text), renderingfont.ptr)
	if CGoReturnPtr.error_no == 0 {
		renderingwatermark.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(renderingwatermark, DeleteRenderingWatermark)
		return renderingwatermark, nil
	} else {
		renderingwatermark.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return renderingwatermark, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *RenderingWatermark) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.RenderingWatermark_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets or sets roation of the watermark in degrees.
// Returns:
//   float32  
func (instance *RenderingWatermark) GetRotation()  (float32,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetRotation( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets roation of the watermark in degrees.
// Parameters:
//   value - float32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetRotation(value float32)  error {
	CGoReturnPtr := C.RenderingWatermark_SetRotation( instance.ptr, C.float(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets scale relative to target page in percent.
// Returns:
//   int32  
func (instance *RenderingWatermark) GetScaleToPagePercent()  (int32,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetScaleToPagePercent( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets scale relative to target page in percent.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetScaleToPagePercent(value int32)  error {
	CGoReturnPtr := C.RenderingWatermark_SetScaleToPagePercent( instance.ptr, C.int(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets opacity of the watermark in range [0, 1].
// Returns:
//   float32  
func (instance *RenderingWatermark) GetOpacity()  (float32,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetOpacity( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets opacity of the watermark in range [0, 1].
// Parameters:
//   value - float32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetOpacity(value float32)  error {
	CGoReturnPtr := C.RenderingWatermark_SetOpacity( instance.ptr, C.float(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Indicates whether the watermark is placed behind page contents.
// Returns:
//   bool  
func (instance *RenderingWatermark) IsBackground()  (bool,  error)  {
	CGoReturnPtr := C.RenderingWatermark_IsBackground( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Indicates whether the watermark is placed behind page contents.
// Parameters:
//   value - bool 
// Returns:
//   void  
func (instance *RenderingWatermark) SetIsBackground(value bool)  error {
	CGoReturnPtr := C.RenderingWatermark_SetIsBackground( instance.ptr, C.bool(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets text of the watermark.
// Returns:
//   string  
func (instance *RenderingWatermark) GetText()  (string,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetText( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  "", err
	}
	result := C.GoString(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets font of the watermark.
// Returns:
//   RenderingFont  
func (instance *RenderingWatermark) GetFont()  (*RenderingFont,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetFont( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &RenderingFont{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteRenderingFont) 

	return result, nil 
}
// Gets or sets horizontal alignment of the watermark to the page.
// Returns:
//   int32  
func (instance *RenderingWatermark) GetHAlignment()  (TextAlignmentType,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetHAlignment( instance.ptr)
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
// Gets or sets horizontal alignment of the watermark to the page.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetHAlignment(value TextAlignmentType)  error {
	CGoReturnPtr := C.RenderingWatermark_SetHAlignment( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets vertical alignment of the watermark to the page.
// Returns:
//   int32  
func (instance *RenderingWatermark) GetVAlignment()  (TextAlignmentType,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetVAlignment( instance.ptr)
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
// Gets or sets vertical alignment of the watermark to the page.
// Parameters:
//   value - int32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetVAlignment(value TextAlignmentType)  error {
	CGoReturnPtr := C.RenderingWatermark_SetVAlignment( instance.ptr, C.int( int32(value)))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets offset value to <see cref="HAlignment"/>
// Returns:
//   float32  
func (instance *RenderingWatermark) GetOffsetX()  (float32,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetOffsetX( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets offset value to <see cref="HAlignment"/>
// Parameters:
//   value - float32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetOffsetX(value float32)  error {
	CGoReturnPtr := C.RenderingWatermark_SetOffsetX( instance.ptr, C.float(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Gets or sets offset value to <see cref="VAlignment"/>
// Returns:
//   float32  
func (instance *RenderingWatermark) GetOffsetY()  (float32,  error)  {
	CGoReturnPtr := C.RenderingWatermark_GetOffsetY( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets or sets offset value to <see cref="VAlignment"/>
// Parameters:
//   value - float32 
// Returns:
//   void  
func (instance *RenderingWatermark) SetOffsetY(value float32)  error {
	CGoReturnPtr := C.RenderingWatermark_SetOffsetY( instance.ptr, C.float(value))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteRenderingWatermark(renderingwatermark *RenderingWatermark){
	runtime.SetFinalizer(renderingwatermark, nil)
	C.Delete_RenderingWatermark(renderingwatermark.ptr)
	renderingwatermark.ptr = nil
}

// Class SheetPrintingPreview 

// Worksheet printing preview.
type SheetPrintingPreview struct {
	ptr unsafe.Pointer
}

// The construct of SheetPrintingPreview
// Parameters:
//   sheet - Worksheet 
//   options - ImageOrPrintOptions 
func NewSheetPrintingPreview(sheet *Worksheet, options *ImageOrPrintOptions) ( *SheetPrintingPreview, error) {
	sheetprintingpreview := &SheetPrintingPreview{}
	CGoReturnPtr := C.New_SheetPrintingPreview(sheet.ptr, options.ptr)
	if CGoReturnPtr.error_no == 0 {
		sheetprintingpreview.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(sheetprintingpreview, DeleteSheetPrintingPreview)
		return sheetprintingpreview, nil
	} else {
		sheetprintingpreview.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return sheetprintingpreview, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SheetPrintingPreview) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.SheetPrintingPreview_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Evaluate the total page count of this worksheet
// Returns:
//   int32  
func (instance *SheetPrintingPreview) GetEvaluatedPageCount()  (int32,  error)  {
	CGoReturnPtr := C.SheetPrintingPreview_GetEvaluatedPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteSheetPrintingPreview(sheetprintingpreview *SheetPrintingPreview){
	runtime.SetFinalizer(sheetprintingpreview, nil)
	C.Delete_SheetPrintingPreview(sheetprintingpreview.ptr)
	sheetprintingpreview.ptr = nil
}

// Class SheetRender 

// Represents a worksheet render which can render worksheet to various images such as (BMP, PNG, JPEG, TIFF..)
// The constructor of this class , must be used after modification of pagesetup, cell style.
type SheetRender struct {
	ptr unsafe.Pointer
}

// the construct of SheetRender, need worksheet and ImageOrPrintOptions as params
// Parameters:
//   worksheet - Worksheet 
//   options - ImageOrPrintOptions 
func NewSheetRender(worksheet *Worksheet, options *ImageOrPrintOptions) ( *SheetRender, error) {
	sheetrender := &SheetRender{}
	CGoReturnPtr := C.New_SheetRender(worksheet.ptr, options.ptr)
	if CGoReturnPtr.error_no == 0 {
		sheetrender.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(sheetrender, DeleteSheetRender)
		return sheetrender, nil
	} else {
		sheetrender.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return sheetrender, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SheetRender) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.SheetRender_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the total page count of current worksheet.
// Returns:
//   int32  
func (instance *SheetRender) GetPageCount()  (int32,  error)  {
	CGoReturnPtr := C.SheetRender_GetPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Gets calculated page scale of the sheet.
// Returns the set scale if <see cref="PageSetup.Zoom"/> is set. Otherwise, returns the calculated scale according to <see cref="PageSetup.FitToPagesWide"/> and <see cref="PageSetup.FitToPagesTall"/>.
// Returns:
//   float64  
func (instance *SheetRender) GetPageScale()  (float64,  error)  {
	CGoReturnPtr := C.SheetRender_GetPageScale( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := float64(CGoReturnPtr.return_value) 

	return result, nil 
}
// Render certain page to a file.
// Parameters:
//   pageIndex - int32 
//   fileName - string 
// Returns:
//   void  
func (instance *SheetRender) ToImage_Int_String(pageindex int32, filename string)  error {
	CGoReturnPtr := C.SheetRender_ToImage_Integer_String( instance.ptr, C.int(pageindex), C.CString(filename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Render whole worksheet as Tiff Image to a file.
// Parameters:
//   filename - string 
// Returns:
//   void  
func (instance *SheetRender) ToTiff_String(filename string)  error {
	CGoReturnPtr := C.SheetRender_ToTiff_String( instance.ptr, C.CString(filename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Releases resources created and used for rendering.
// Returns:
//   void  
func (instance *SheetRender) Dispose()  error {
	CGoReturnPtr := C.SheetRender_Dispose( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteSheetRender(sheetrender *SheetRender){
	runtime.SetFinalizer(sheetrender, nil)
	C.Delete_SheetRender(sheetrender.ptr)
	sheetrender.ptr = nil
}

// Class SheetSet 

// Describes a set of sheets.
type SheetSet struct {
	ptr unsafe.Pointer
}


// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *SheetSet) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.SheetSet_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets a set with active sheet of the workbook.
// Returns:
//   SheetSet  
func SheetSet_GetActive()  (*SheetSet,  error)  {
	CGoReturnPtr := C.SheetSet_GetActive()
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SheetSet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSheetSet) 

	return result, nil 
}
// Gets a set with visible sheets of the workbook in their original order.
// Returns:
//   SheetSet  
func SheetSet_GetVisible()  (*SheetSet,  error)  {
	CGoReturnPtr := C.SheetSet_GetVisible()
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SheetSet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSheetSet) 

	return result, nil 
}
// Gets a set with all sheets of the workbook in their original order.
// Returns:
//   SheetSet  
func SheetSet_GetAll()  (*SheetSet,  error)  {
	CGoReturnPtr := C.SheetSet_GetAll()
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  nil, err
	}
	result := &SheetSet{}
	result.ptr = CGoReturnPtr.return_value 
	runtime.SetFinalizer(result, DeleteSheetSet) 

	return result, nil 
}


func DeleteSheetSet(sheetset *SheetSet){
	runtime.SetFinalizer(sheetset, nil)
	C.Delete_SheetSet(sheetset.ptr)
	sheetset.ptr = nil
}

// Class WorkbookPrintingPreview 

// Workbook printing preview.
type WorkbookPrintingPreview struct {
	ptr unsafe.Pointer
}

// The construct of WorkbookPrintingPreview
// Parameters:
//   workbook - Workbook 
//   options - ImageOrPrintOptions 
func NewWorkbookPrintingPreview(workbook *Workbook, options *ImageOrPrintOptions) ( *WorkbookPrintingPreview, error) {
	workbookprintingpreview := &WorkbookPrintingPreview{}
	CGoReturnPtr := C.New_WorkbookPrintingPreview(workbook.ptr, options.ptr)
	if CGoReturnPtr.error_no == 0 {
		workbookprintingpreview.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(workbookprintingpreview, DeleteWorkbookPrintingPreview)
		return workbookprintingpreview, nil
	} else {
		workbookprintingpreview.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return workbookprintingpreview, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WorkbookPrintingPreview) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.WorkbookPrintingPreview_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Evaluate the total page count of this workbook
// Returns:
//   int32  
func (instance *WorkbookPrintingPreview) GetEvaluatedPageCount()  (int32,  error)  {
	CGoReturnPtr := C.WorkbookPrintingPreview_GetEvaluatedPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}


func DeleteWorkbookPrintingPreview(workbookprintingpreview *WorkbookPrintingPreview){
	runtime.SetFinalizer(workbookprintingpreview, nil)
	C.Delete_WorkbookPrintingPreview(workbookprintingpreview.ptr)
	workbookprintingpreview.ptr = nil
}

// Class WorkbookRender 

// Represents a Workbook render.
// The constructor of this class , must be used after modification of pagesetup, cell style.
type WorkbookRender struct {
	ptr unsafe.Pointer
}

// The construct of WorkbookRender
// Parameters:
//   workbook - Workbook 
//   options - ImageOrPrintOptions 
func NewWorkbookRender(workbook *Workbook, options *ImageOrPrintOptions) ( *WorkbookRender, error) {
	workbookrender := &WorkbookRender{}
	CGoReturnPtr := C.New_WorkbookRender(workbook.ptr, options.ptr)
	if CGoReturnPtr.error_no == 0 {
		workbookrender.ptr = CGoReturnPtr.return_value
		runtime.SetFinalizer(workbookrender, DeleteWorkbookRender)
		return workbookrender, nil
	} else {
		workbookrender.ptr = nil
		err := errors.New(C.GoString(CGoReturnPtr.error_message))
		return workbookrender, err
	}	
}

// Checks whether the implementation object is nullptr.
// Returns:
//   bool  
func (instance *WorkbookRender) IsNull()  (bool,  error)  {
	CGoReturnPtr := C.WorkbookRender_IsNull( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  true, err
	}
	result := CGoReturnPtr.return_value != C.bool(true) 

	return result, nil 
}
// Gets the total page count of workbook.
// Returns:
//   int32  
func (instance *WorkbookRender) GetPageCount()  (int32,  error)  {
	CGoReturnPtr := C.WorkbookRender_GetPageCount( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  0, err
	}
	result := int32(CGoReturnPtr.return_value) 

	return result, nil 
}
// Render whole workbook as Tiff Image to a file.
// Parameters:
//   filename - string 
// Returns:
//   void  
func (instance *WorkbookRender) ToImage_String(filename string)  error {
	CGoReturnPtr := C.WorkbookRender_ToImage_String( instance.ptr, C.CString(filename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Render certain page to a file.
// Parameters:
//   pageIndex - int32 
//   fileName - string 
// Returns:
//   void  
func (instance *WorkbookRender) ToImage_Int_String(pageindex int32, filename string)  error {
	CGoReturnPtr := C.WorkbookRender_ToImage_Integer_String( instance.ptr, C.int(pageindex), C.CString(filename))
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}
// Releases resources created and used for rendering.
// Returns:
//   void  
func (instance *WorkbookRender) Dispose()  error {
	CGoReturnPtr := C.WorkbookRender_Dispose( instance.ptr)
	if CGoReturnPtr.error_no != 0 {
		err := errors.New(C.GoString(CGoReturnPtr.error_message))	
		return  err
	}

	return nil 
}


func DeleteWorkbookRender(workbookrender *WorkbookRender){
	runtime.SetFinalizer(workbookrender, nil)
	C.Delete_WorkbookRender(workbookrender.ptr)
	workbookrender.ptr = nil
}
