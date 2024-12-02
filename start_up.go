// +build windows

package asposecells

// #cgo CXXFLAGS: -std=c++11
// #cgo CFLAGS: -I.
// #cgo windows LDFLAGS: -L./lib/win_x86_64 -lAspose.Cells.CWrapper
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