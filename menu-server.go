package main

import (
	"database/sql"
	"fmt"
)

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
