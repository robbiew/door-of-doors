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

	if currMenu == "door" {
		doors = doorsByCategory(db, currCat)
		lenList = len(doors)
		lightbar = "art/lightbar-long.ans"
		blank = "                                          "

	}
	if currMenu == "server" {

		// currTitle = doors[currY].DoorTitle

		serversList = doorByServer(db)
		lenList = len(serversList)
		lightbar = "art/lightbar.ans"
		blank = "                               "

	}
	if currMenu == "category" {
		categories = categoryList(db)
		lenList = len(categories)
		lightbar = "art/lightbar.ans"
		blank = "                               "
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
				if currMenu == "category" {
					currCat = categories[i].CategoryId
					currCatName = categories[i].CategoryName
					fmt.Print(bgCyan + cyanHi + categories[i].CategoryName + reset)
				}
				if currMenu == "door" {
					currTitle = doors[i].DoorTitle
					fmt.Print(bgCyan + cyanHi + doors[i].DoorTitle + reset)
				}
				if currMenu == "server" {
					fmt.Print(bgCyan + cyanHi + serversList[i].ServerName + reset)
				}
				yLoc++

			} else {
				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				printAnsiLoc("art/seperator.ans", xLoc, yLoc)
				moveCursor(xLoc+2, yLoc)
				if currMenu == "category" {
					fmt.Printf(cyanHi + categories[i].CategoryName + reset)
				}
				if currMenu == "door" {
					fmt.Printf(cyanHi + doors[i].DoorTitle + reset)
				}
				if currMenu == "server" {
					fmt.Print(cyanHi + serversList[i].ServerName + reset)
				}
				yLoc++

			}
		}
		i++
		continue
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

	if currMenu == "door" {
		moveCursor(doorDescX, doorDescY)
		fmt.Print("                                ")
		moveCursor(doorDescX, doorDescY)
		fmt.Print(doors[currY].DoorTitle)
		moveCursor(catDescX, catDescY)
		fmt.Print(currCatName)
	}

	if currMenu == "category" {
		currcat := doorsByCategory(db, currCat)
		doorCount := len(currcat)
		moveCursor(catDescX, catDescY)
		fmt.Print("                                               ")
		moveCursor(catDescX, catDescY)
		fmt.Printf("%v: %v games", categories[currY].CategoryName, doorCount)
	}
	if currMenu == "server" {
		moveCursor(doorDescX, doorDescY)
		fmt.Print("                                ")
		moveCursor(doorDescX, doorDescY)
		fmt.Print(currTitle)
		moveCursor(doorDescX, doorDescY)
		fmt.Print("                                               ")
		moveCursor(doorDescX, doorDescY)
		fmt.Printf("%v description", serversList[currY].ServerName)

	}

	bottomY := listHeight + scrollY
	printAnsiLoc("art/arrow-up.ans", xLoc, scrollY)
	printAnsiLoc("art/arrow-down.ans", xLoc, bottomY)

}
