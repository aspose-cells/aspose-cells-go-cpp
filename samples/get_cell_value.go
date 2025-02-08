package main

func GetCellValue() {

	_, _, cells := create_table_data()

	cell, err5 := cells.Get_String("A1")
	if err5 != nil {
		println(err5)
	}
	print("Cell A1 value:")
	println(cell.GetStringValue())

	cell, err5 = cells.Get_String("B3")
	if err5 != nil {
		println(err5)
	}
	print("Cell B3 value:")
	println(cell.GetStringValue())

}
