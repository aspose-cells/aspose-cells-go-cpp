// +build windows

/* ----------------------------------------------------------------
 * Copyright (c) 2001-2025 Aspose Pty Ltd. All Rights Reserved.
 * Powered by Aspose.Cells.
 * ---------------------------------------------------------------*/


package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}/lib/win_x86_64" -L"${SRCDIR}/lib/win_x86_64" -lAspose.Cells.CWrapper
// #include <AsposeCellsCWrapper.h>
import "C"
import (
	"unsafe"
)
var GlobalVar bool = false
func init() {
	if GlobalVar == false {
		GlobalVar = true
		C.Startup()
	}
}

func Version() string {
	println("---Version: v25.4.0---")
	return "v25.4.0"
} 
type UUID struct {
	ptr unsafe.Pointer
}

