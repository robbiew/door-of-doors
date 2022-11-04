package main

import (
	"database/sql"
	"fmt"
)

func doorMenu(db *sql.DB) {
	currCode = getCatCode(currCat, categories)
	ClearScreen()
	MoveCursor(0, 0)
	header(U.W)

	doorsList := doorsByCategory(db)

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 3
	xLoc2 := 36

	var data string

	for i := 0; i < len(doorsList); i++ {

		data = doorsList[i].DoorTitle

		if count < 14 {
			MoveCursor(xLoc1, yLoc1)
			if count < 9 {
				MoveCursor(xLoc1, yLoc1)
				fmt.Printf(White+" %d"+BlackHi+"..."+Reset+RedHi+"%s\n"+Reset, i+1, data)

			} else {
				fmt.Printf(White+"%d"+BlackHi+"..."+Reset+RedHi+"%s\n"+Reset, i+1, data)
			}
			yLoc1++
		}
		if count >= 14 {
			MoveCursor(xLoc2, yLoc2)
			fmt.Printf(White+"%d"+BlackHi+"..."+Reset+RedHi+"%s\n"+Reset, i+1, data)
			yLoc2++
		}
		count++
	}

	MoveCursor(3, 24)
	prompt("red")

}
