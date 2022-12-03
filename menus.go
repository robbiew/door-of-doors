package main

import (
	"database/sql"
	"fmt"
)

func catMenu(db *sql.DB, arrow string) {
	categories = categoryList(db)
	lenList = len(categories)
	yLoc := 8
	xLoc := 2
	const blank = "                                "

	if arrow == "up" {
		if currY-1 > 0 {
			currY--
		}
		if currStart <= lenList && currStart > 0 {
			currStart--
		}
	}
	if arrow == "down" {
		if currY < listHeight || currY >= listHeight && currStart+listHeight < lenList-1 {
			currY++
		}
		if currY > listHeight && currStart+listHeight < lenList-1 {
			currStart++
		}
	}
	// moveCursor(0, 0)
	// fmt.Printf("\r\ncurrY: %v    \r\ncurrStart: %v   ", currY, currStart)
	// moveCursor(0, 0)

	i := 0

	for i <= listHeight+currStart {
		if i >= 0 && i < lenList && i > currStart {
			if i == currY {
				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				moveCursor(xLoc, yLoc)
				printAnsiLoc("art/bullet-on.ans", xLoc, yLoc)
				moveCursor(xLoc+1, yLoc)
				fmt.Print(bgCyan + cyanHi + " " + categories[i].CategoryName + " " + fmt.Sprint(currY) + " " + fmt.Sprint(i) + " " + reset)
				yLoc++
			} else {
				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				moveCursor(xLoc, yLoc)
				fmt.Printf("  " + categories[i].CategoryName + " " + fmt.Sprint(currY) + " " + fmt.Sprint(i) + " ")
				yLoc++
			}

		}

		i++
	}

	// for i := 0; i < len(categories); i++ {
	// 	if count < listHeight {
	// 		moveCursor(xLoc1, yLoc1)
	// 		if count < 9 {
	// 			fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
	// 		} else {
	// 			fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
	// 		}
	// 		yLoc1++
	// 	}
	// 	if count >= listHeight {
	// 		moveCursor(xLoc2, yLoc2)
	// 		fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
	// 		yLoc2++
	// 	}
	// 	count++
	// }
	// moveCursor(3, 24)
	// prompt("red")

}

func getCatCode(id int, CategoryList []CategoryList) (result string) {
	for _, categoryList := range CategoryList {
		if categoryList.CategoryId == id {
			result = categoryList.CategoryCode

			break
		}
	}
	return result
}

func doorMenu(db *sql.DB) {
	clearScreen()
	currCode = "DOOR"
	header(U.W)

	realCat := categories[currCat-1].CategoryId
	categories = categoryList(db)

	doorsList = doorsByCategory(db, realCat)
	lenList = len(doorsList)

	currCatName = categories[currCat-1].CategoryName

	moveCursor(3, 6)
	fmt.Print(whiteHi + currCatName + ":" + reset)

	recordsPerCol := 12

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 2
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
		if lenList > (recordsPerCol * 2) {
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

func serverMenu(db *sql.DB) {
	clearScreen()
	currCode = "SERVER"
	header(U.W)

	// get title
	currTitle = doorsList[currDoor].DoorTitle
	serversList = doorByServer(db)
	lenList = len(serversList)

	moveCursor(3, 6)
	fmt.Print(whiteHi + "Play " + yellowHi + currTitle + whiteHi + " on:" + reset)

	moveCursor(39, 8)
	fmt.Print(currTitle + " (" + serversList[0].Year + ")")
	moveCursor(39, 10)
	d := (justifyText(serversList[0].Desc, 38))
	printMultiStringAt(d, 39, 10)

	count := 0
	xLoc1 := 3
	yLoc1 := 8

	var data string

	for i := 0; i < len(serversList); i++ {
		data = serversList[i].ServerName
		moveCursor(xLoc1, yLoc1)
		fmt.Printf(yellowHi+"%d"+cyanHi+"..."+reset+yellowHi+"%s\n"+reset, i+1, data)
		yLoc1++
		count++
	}

	moveCursor(3, 24)
	prompt("red")

}
