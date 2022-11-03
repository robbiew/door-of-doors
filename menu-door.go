package main

import (
	"database/sql"
	"fmt"
)

func doorMenu(db *sql.DB, dataChan chan []byte, errorChan chan error) {
	count := 0
	yLoc1 := 7
	yLoc2 := 7
	xLoc1 := 2
	xLoc2 := 29

	ClearScreen()

	MoveCursor(0, 0)
	catHeader(U.W)

	doorsList := doorsByCategory(db)
	for i := 0; i < len(doorsList); i++ {

		if count < 15 {
			MoveCursor(xLoc1, yLoc1)
			fmt.Printf(BlackHi+"["+White+"%d"+BlackHi+"]"+Reset+RedHi+" %s\n"+Reset, i+1, doorsList[i].DoorTitle)
			yLoc1++
		}
		if count >= 15 {
			MoveCursor(xLoc2, yLoc2)
			fmt.Printf(BlackHi+"["+White+"%d"+BlackHi+"]"+Reset+RedHi+" %s\n"+Reset, i+1, doorsList[i].DoorTitle)
			yLoc2++
		}
		count++
	}

	MoveCursor(3, 24)
	prompt("red")

}
