package main

import (
	"database/sql"
	"fmt"
)

func catMenu(db *sql.DB) {
	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 2
	xLoc2 := 34

	ClearScreen()

	MoveCursor(0, 0)
	mainHeader(U.W, currTab)
	PrintStringLoc(" "+C.Menu_Title+" "+C.Version+" "+Reset, 2, 2, BlueHi, BgBlue)

	categories = categoryList(db)
	for i := 0; i < len(categories); i++ {
		if count < 14 {
			MoveCursor(xLoc1, yLoc1)
			fmt.Printf(BlackHi+"["+White+"%d"+BlackHi+"]"+Reset+BlueHi+" %s\n"+Reset, i+1, categories[i].CategoryName)
			yLoc1++
		}
		if count >= 14 {
			MoveCursor(xLoc2, yLoc2)
			fmt.Printf(BlackHi+"["+White+"%d"+BlackHi+"]"+Reset+BlueHi+" %s\n"+Reset, i+1, categories[i].CategoryName)
			yLoc2++
		}
		count++
	}
	MoveCursor(3, 24)
	prompt("blue")

}
