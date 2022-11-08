package main

import (
	"database/sql"
	"fmt"
)

func doorMenu(db *sql.DB) {
	clearScreen()
	currCode = "DOOR"
	header(U.W)

	doorsList = doorsByCategory(db, currCat)
	lenList = len(doorsList)
	categories = categoryList(db)
	currCatName = categories[currCat-1].CategoryName

	moveCursor(3, 6)
	fmt.Print(whiteHi + currCatName + ":" + reset)

	recordsPerCol := 12

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 3
	xLoc2 := 43

	var SelectionData string

	if currPage == 1 {
		for i := 0; i < len(doorsList); i++ {
			SelectionData = doorsList[i].DoorTitle
			if count < recordsPerCol {
				moveCursor(xLoc1, yLoc1)
				if count < 9 {
					moveCursor(xLoc1, yLoc1)
					fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
				} else {
					fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
				}
				yLoc1++
			}
			if count >= recordsPerCol && count <= (recordsPerCol*2)-1 {
				moveCursor(xLoc2, yLoc2)
				fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
				yLoc2++
			}
			count++
		}
		if lenList > (recordsPerCol*2)-1 {
			paginator = true
			moveCursor(3, 6)
			fmt.Print(whiteHi + currCatName + " (1/2):" + reset)
			moveCursor(26, 21)
			fmt.Print(white + "[" + blackHi + "," + reset + white + "]" + cyanHi + "..." + redHi + "Prev/Next " + bgRed + whiteHi + " 1 " + reset + whiteHi + " 2 " + reset)
		}
	}

	if currPage == 2 {

		clearScreen()
		header(U.W)
		moveCursor(3, 6)
		fmt.Print(whiteHi + currCatName + " (2/2):" + reset)

		count := 24

		for i := count; i < len(doorsList); i++ {
			SelectionData = doorsList[i].DoorTitle
			if count < recordsPerCol*3 {
				moveCursor(xLoc1, yLoc1)
				fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
				yLoc1++
			}
			if count >= (recordsPerCol*3)-1 {
				moveCursor(xLoc2, yLoc2)
				fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s\n"+reset, i+1, SelectionData)
				yLoc2++
			}
			count++
		}
		moveCursor(26, 21)
		fmt.Print(white + "[" + blackHi + "," + reset + white + "]" + cyanHi + "..." + redHi + "Prev/Next " + bgBlack + whiteHi + " 1 " + reset + bgRed + whiteHi + " 2 " + reset)
	}

	moveCursor(3, 24)
	prompt("red")

}
