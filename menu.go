package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

func catMenu(db *sql.DB, arrow string) {
	listHeight = 8
	yLoc := U.H - listHeight

	scrollY := yLoc - 1
	xLoc := 2

	lightbarLen := 48
	doorDescLen := uint(34)

	doorDescX := 42
	doorDescY := yLoc

	doScroll := false
	var listMax int

	// blank out description areas before writing

	blankDesc := strings.Repeat(" ", int(doorDescLen))
	blankLightbar := strings.Repeat(" ", lightbarLen)
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

	// iterate through the  list
	i := 0
	for i <= listMax+currStart {
		if i >= 0 && i < lenList && i >= currStart {
			if i == currY {
				moveCursor(xLoc, yLoc)
				fmt.Print(blankLightbar)
				printAnsiLoc("art/seperator-on.ans", xLoc, yLoc)
				moveCursor(xLoc+1, yLoc)
				printAnsiLoc(lightbar, xLoc+1, yLoc)
				moveCursor(xLoc+2, yLoc)
				if currMenu == "category" {
					currCat = categories[i].CategoryId
					currCatName = categories[i].CategoryName
					currDoors := doorsByCategory(db, currCat)
					doorCount := len(currDoors)

					fmt.Printf(reset + bgCyan + cyanHi + currCatName + reset)
					moveCursor(doorDescX, yLoc)
					fmt.Printf(reset+greenHi+"%v games ", doorCount)

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
				fmt.Print(blankLightbar)
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
	printAnsiLoc("art/arrow-down.ans", xLoc, yLoc)

	// print some things after the scroll area has been printed
	if currMenu == "door" {
		// Game Desc
		i := 0
		y := doorDescY
		for i < listHeight-2 {
			printMultiStringAt(blankDesc, doorDescX, y+1)
			i++
			y++
		}

		printMultiStringAt(blankDesc, doorDescX, doorDescY)
		moveCursor(doorDescX, doorDescY)
		fmt.Printf(greenHi+"%v "+reset, doors[currY].DoorTitle)

		if len(doors[currY].DoorYear) > 0 {
			fmt.Printf(greenHi+"("+reset+green+"%v"+greenHi+")"+reset, doors[currY].DoorYear)
		}

		wrapped := wordwrap.WrapString(doors[currY].DoorDesc, doorDescLen)
		printMultiStringAt(green+wrapped+reset, doorDescX, doorDescY+1)

	}

	if currMenu == "server" {

		// Server Description
		i := 0
		y := doorDescY
		for i < listHeight {
			printMultiStringAt(blankDesc, doorDescX, y)
			i++
			y++
		}
		moveCursor(doorDescX, doorDescY)
		doors = doorsByCategory(db, currCat)
		fmt.Printf(greenHi+"%v "+reset, doors[currY].DoorTitle)
		if len(doors[currY].DoorYear) > 0 {
			fmt.Printf(greenHi+"("+reset+green+"%v"+greenHi+") "+reset, doors[currY].DoorYear)
		}

		wrapped := wordwrap.WrapString(servers[currY].ServerDesc, doorDescLen)
		printMultiStringAt(blackHi+wrapped+reset, doorDescX, doorDescY+1)

	}

}
