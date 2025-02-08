package main

func GetStyle() {
	_, _, cells := create_table_data()

	cell, err5 := cells.Get_String("A1")
	if err5 != nil {
		println(err5)
	}
	style, _ := cell.GetStyle()
	isnull, _ := style.IsNull()
	if isnull {
		print("Get Style address: ")
		print(style)
	}
}
