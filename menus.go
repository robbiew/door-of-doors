package main

import (
	"database/sql"
	"fmt"
)

func catMenu(db *sql.DB, arrow string) {
	yLoc := 13
	scrollY := yLoc
	xLoc := 2

	doorDescX := 48
	doorDescY := yLoc

	catDescX := 2
	catDescY := 2

	doScroll := false
	var listMax int
	var lightbar string
	var blank string

	if menuType == "category" {
		categories = categoryList(db)
		lenList = len(categories)
		lightbar = "art/lightbar.ans"
		blank = "                               "
	}

	// var realCat int
	if menuType == "door" {
		doorsList = doorsByCategory(db, currCat)
		lenList = len(doorsList)
		lightbar = "art/lightbar-long.ans"
		blank = "                                          "
	}

	// if there are more items in list than display height
	if lenList > listHeight {
		doScroll = true
	}

	if doScroll {
		listMax = listHeight
	} else {
		listMax = lenList - 1
	}

	// Arrow keys move the light bars
	if arrow == "up" {
		if currY > 0 {
			currY--
		}
		if currStart <= lenList && currStart > 0 {
			currStart--
		}
	}
	if arrow == "down" {
		if currY < listMax || currY >= listMax && currStart+listMax < lenList-1 {
			currY++
		}
		if currY > listMax && currStart+listMax < lenList-1 {
			currStart++
		}
	}

	// iterate through the  list
	i := 0
	for i <= listMax+currStart {
		if i >= 0 && i < lenList && i >= currStart {
			if i == currY {
				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				printAnsiLoc("art/seperator.ans", xLoc, yLoc)
				moveCursor(xLoc+1, yLoc)
				printAnsiLoc(lightbar, xLoc+1, yLoc)
				moveCursor(xLoc+2, yLoc)
				if menuType == "category" {
					currCat = categories[i].CategoryId
					fmt.Print(bgCyan + cyanHi + categories[i].CategoryName + reset)
				}
				if menuType == "door" {
					fmt.Print(bgCyan + cyanHi + doorsList[i].DoorTitle + reset)
				}
				yLoc++
			} else {

				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				printAnsiLoc("art/seperator.ans", xLoc, yLoc)
				moveCursor(xLoc+2, yLoc)
				if menuType == "category" {
					fmt.Printf(cyanHi + categories[i].CategoryName + reset)
				}
				if menuType == "door" {
					fmt.Printf(cyanHi + doorsList[i].DoorTitle + reset)
				}
				yLoc++
			}
		}
		i++
	}
	if lenList < listHeight {
		x := 0
		for x < (listHeight - lenList) {
			moveCursor(xLoc+1, yLoc)
			printAnsiLoc("art/seperator.ans", xLoc, yLoc)
			fmt.Print("\r\n")
			x++
			yLoc++
		}
	}

	if menuType == "door" {
		moveCursor(doorDescX, doorDescY)
		fmt.Print("                                ")
		moveCursor(doorDescX, doorDescY)
		fmt.Print(doorsList[currY].DoorTitle)
		moveCursor(catDescX, catDescY)
		fmt.Print(currCatName)
	}
	if menuType == "category" {
		currcat := doorsByCategory(db, currCat)
		doorCount := len(currcat)
		moveCursor(catDescX, catDescY)
		fmt.Print("                                               ")
		moveCursor(catDescX, catDescY)
		fmt.Printf("%v: %v games", categories[currY].CategoryName, doorCount)
		currCatName = categories[currY].CategoryName
	}

	bottomY := listHeight + scrollY
	printAnsiLoc("art/arrow-up.ans", xLoc, scrollY)
	printAnsiLoc("art/arrow-down.ans", xLoc, bottomY)

}

// func serverMenu(db *sql.DB) {
// 	clearScreen()
// 	header(U.W)

// 	// get title
// 	currTitle = doorsList[currDoor].DoorTitle
// 	serversList = doorByServer(db)
// 	lenList = len(serversList)

// 	moveCursor(3, 6)
// 	fmt.Print(whiteHi + "Play " + yellowHi + currTitle + whiteHi + " on:" + reset)

// 	moveCursor(39, 8)
// 	fmt.Print(currTitle + " (" + serversList[0].Year + ")")
// 	moveCursor(39, 10)
// 	d := (justifyText(serversList[0].Desc, 38))
// 	printMultiStringAt(d, 39, 10)

// 	count := 0
// 	xLoc1 := 3
// 	yLoc1 := 8

// 	var data string

// 	for i := 0; i < len(serversList); i++ {
// 		data = serversList[i].ServerName
// 		moveCursor(xLoc1, yLoc1)
// 		fmt.Printf(yellowHi+"%d"+cyanHi+"..."+reset+yellowHi+"%s\n"+reset, i+1, data)
// 		yLoc1++
// 		count++
// 	}

// 	moveCursor(3, 24)
// 	prompt("red")

// }
