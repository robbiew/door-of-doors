package main

import (
	"database/sql"
	"fmt"

	"github.com/mitchellh/go-wordwrap"
)

func catMenu(db *sql.DB, arrow string) {
	yLoc := 14
	scrollY := yLoc - 1
	xLoc := 2

	listHeight = 9

	doorDescX := 46
	doorDescY := yLoc

	catDescX := 2
	catDescY := 2

	doScroll := false
	var listMax int

	// blank out description areas before writing
	blankDesc :=
		"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n" +
			"                               \r\n"

	blank := "                                          "
	lightbar := "art/lightbar-long.ans"

	if currMenu == "category" {
		categories = categoryList(db)
		lenList = len(categories)
	}
	if currMenu == "door" {
		doors = doorsByCategory(db, currCat)
		lenList = len(doors)
	}
	if currMenu == "server" {
		servers = doorByServer(db)
		lenList = len(servers)
	}

	// if there are more items in list than display height
	if lenList > listHeight {
		doScroll = true
	}

	if doScroll {
		listMax = listHeight - 1
	} else {
		listMax = lenList
	}

	// Arrow keys move the light bars
	if arrow == "up" {
		if currY > 0 {
			currY--
		}
		if currStart <= lenList && currStart > 0 && currY+currStart <= listMax-1 || currStart > currY {
			currStart--
		}

	}
	if arrow == "down" {
		if currY < lenList-1 || currY >= listMax && currStart+listMax < lenList-1 {
			currY++
		}
		if currY > listMax && currStart+listMax < lenList-1 {
			currStart++
		}

	}

	// debug positions
	// moveCursor(2, 25)
	// fmt.Print("                                                                          ")
	// moveCursor(2, 25)
	// fmt.Printf("DEBUG: listMax: %v, currStart: %v, lenList: %v, currY: %v", listMax, currStart, lenList, currY)

	// scroll indicators
	printAnsiLoc("art/arrow-up.ans", xLoc, scrollY)
	printAnsiLoc("art/arrow-down.ans", xLoc, scrollY+listHeight+1)

	// iterate through the  list
	i := 0
	for i <= listMax+currStart {
		if i >= 0 && i < lenList && i >= currStart {
			if i == currY {
				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				printAnsiLoc("art/seperator-on.ans", xLoc, yLoc)

				moveCursor(xLoc+1, yLoc)
				printAnsiLoc(lightbar, xLoc+1, yLoc)
				moveCursor(xLoc+2, yLoc)
				if currMenu == "category" {
					currCat = categories[i].CategoryId
					currCatName = categories[i].CategoryName
					fmt.Printf(reset + bgCyan + cyanHi + currCatName + reset)

				}
				if currMenu == "door" {
					currTitle = doors[i].DoorTitle
					fmt.Printf(reset + bgCyan + cyanHi + currTitle + reset)

				}
				if currMenu == "server" {
					serverTitle = servers[i].ServerName
					fmt.Printf(reset + bgCyan + cyanHi + serverTitle + reset)

				}
				yLoc++

			} else {
				moveCursor(xLoc, yLoc)
				fmt.Print(blank)
				printAnsiLoc("art/seperator.ans", xLoc, yLoc)
				moveCursor(xLoc+2, yLoc)
				if currMenu == "server" {
					serverTitle = servers[i].ServerName
					fmt.Printf(reset + cyanHi + serverTitle + reset)
				}
				if currMenu == "category" {
					fmt.Printf(reset + cyanHi + categories[i].CategoryName + reset)
				}
				if currMenu == "door" {
					fmt.Printf(reset + cyanHi + doors[i].DoorTitle + reset)
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
			if x > 0 && x < (listHeight-lenList) {
				fmt.Print("\r\n")
			}
			x++
			yLoc++
		}
	}

	// print some things after the scroll area has been printed
	if currMenu == "door" {
		currcat := doorsByCategory(db, currCat)
		doorCount := len(currcat)
		moveCursor(catDescX, catDescY)
		fmt.Printf(reset+"%v: %v games", currCatName, doorCount)
		if len(doors[currY].DoorDesc) > 1 {
			printMultiStringAt(blankDesc, doorDescX, doorDescY)
			// Game Title
			moveCursor(doorDescX, doorDescY)
			fmt.Printf(white+"%v (%v)"+reset, doors[currY].DoorTitle, doors[currY].DoorYear)

			// Game Description
			wrapped := wordwrap.WrapString(doors[currY].DoorDesc, 30)
			printMultiStringAt(wrapped, doorDescX, doorDescY+1)
		} else {
			printMultiStringAt(blankDesc, doorDescX, doorDescY)

		}
	}

	if currMenu == "category" {
		currcat := doorsByCategory(db, currCat)
		doorCount := len(currcat)
		moveCursor(catDescX, catDescY)
		fmt.Print("                                               ")
		moveCursor(catDescX, catDescY)
		fmt.Printf(reset+"%v: %v games", currCatName, doorCount)
	}

	if currMenu == "server" {
		// Game Title
		printMultiStringAt(blankDesc, doorDescX, doorDescY)
		moveCursor(catDescX, catDescY)
		fmt.Print(reset + currTitle)

		// Server Description
		wrapped := wordwrap.WrapString(servers[currY].ServerDesc, 30)
		printMultiStringAt(wrapped, doorDescX, doorDescY)

	}

}
