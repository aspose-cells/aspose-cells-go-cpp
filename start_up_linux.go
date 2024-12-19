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
	"unsafe"
)
var GlobalVar bool = false
func init() {
	if GlobalVar == false {
		GlobalVar = true
		C.Startup()
	}
}

func Version() {
	println("---Version: v24.12.2---")
} 
type UUID struct {
	ptr unsafe.Pointer
}

type Stream struct {
	ptr unsafe.Pointer
}

type Vector struct {
	ptr unsafe.Pointer
	data_type string
}

type Enumerator struct {
	ptr unsafe.Pointer
	data_type string
}
