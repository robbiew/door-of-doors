package main

import (
	"database/sql"
	"fmt"
)

func doorMenu(db *sql.DB) {
	clearScreen()
	header(U.W)

	doorsList = doorsByCategory(db, currCat)

	categories = categoryList(db)
	currCatName = categories[currCat-1].CategoryName

	moveCursor(3, 6)
	fmt.Print(whiteHi + currCatName + ":" + reset)

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 3
	xLoc2 := 32

	var SelectionData string

	for i := 0; i < len(doorsList); i++ {

		SelectionData = doorsList[i].DoorTitle

		if count < 14 {
			moveCursor(xLoc1, yLoc1)
			if count < 9 {
				moveCursor(xLoc1, yLoc1)
				fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
			} else {
				fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
			}
			yLoc1++
		}
		if count >= 14 {
			moveCursor(xLoc2, yLoc2)
			fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
			yLoc2++
		}
		count++
	}

	moveCursor(3, 24)
	prompt("red")

}
