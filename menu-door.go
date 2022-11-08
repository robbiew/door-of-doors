package main

import (
	"database/sql"
	"fmt"
)

func doorMenu(db *sql.DB) {
	clearScreen()
	currCode = "DOOR"
	header(U.W)
	paginator = false

	doorsList = doorsByCategory(db, currCat)
	lenList = len(doorsList)
	categories = categoryList(db)
	currCatName = categories[currCat-1].CategoryName

	moveCursor(3, 6)
	fmt.Print(whiteHi + currCatName + ":" + reset)

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 3
	xLoc2 := 43

	var SelectionData string

	for i := 0; i < len(doorsList); i++ {

		SelectionData = doorsList[i].DoorTitle

		if count < 12 {
			moveCursor(xLoc1, yLoc1)
			if count < 9 {
				moveCursor(xLoc1, yLoc1)
				fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
			} else {
				fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
			}
			yLoc1++
		}
		if count >= 12 && count < 24 {
			moveCursor(xLoc2, yLoc2)
			fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
			yLoc2++
		}
		if count > 23 {
			paginator = true
		}
		count++

	}

	if lenList > 23 {

		moveCursor(26, 21)
		if currPage == 1 {
			fmt.Print("[,]...Prev/Next " + bgRed + whiteHi + " 1 " + reset + whiteHi + " 2 " + reset)
		}
		if currPage == 2 {
			fmt.Print("[,]...Prev/Next " + whiteHi + " 1 " + reset + bgRed + whiteHi + " 2 " + reset)
		}
	}

	moveCursor(3, 24)
	prompt("red")

}
