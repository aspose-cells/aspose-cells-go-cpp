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
 	
)

/**************Enum MarkdownTableHeaderType *****************/

// Represents the header type of the table in the markdown file.
type MarkdownTableHeaderType int32

const(
// First row as header of the table.
MarkdownTableHeaderType_FirstRow MarkdownTableHeaderType = 0 

// Column name (such as A,B,C...) as header of the table.
MarkdownTableHeaderType_ColumnHeader MarkdownTableHeaderType = 1 

// An empty header row.
MarkdownTableHeaderType_Empty MarkdownTableHeaderType = 2 
)

func Int32ToMarkdownTableHeaderType(value int32)(MarkdownTableHeaderType ,error){
	switch value {
		case 0:  return MarkdownTableHeaderType_FirstRow, nil  
		case 1:  return MarkdownTableHeaderType_ColumnHeader, nil  
		case 2:  return MarkdownTableHeaderType_Empty, nil  
		default:
			return 0 ,fmt.Errorf("invalid MarkdownTableHeaderType value: %d", value)
	}
}
