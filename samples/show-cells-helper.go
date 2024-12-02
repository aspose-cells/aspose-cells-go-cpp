package main

import (
	. "asposecells"
)

func ShowCellsHelper() {
	cell_name, _ := CellsHelper_CellIndexToName(3, 5)
	row, column, _ := CellsHelper_CellNameToIndex("C4")
	println(cell_name)
	println(row)
	println(column)
}
