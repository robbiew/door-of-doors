package main

import (
	"database/sql"
	"fmt"

	gd "github.com/robbiew/door-of-doors/common"
)

func catMenu(db *sql.DB, w int, alias string, tleft int) {
	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 2
	xLoc2 := 34

	categories = categoryList(db)
	for i := 0; i < len(categories); i++ {
		if count < 14 {
			gd.MoveCursor(xLoc1, yLoc1)
			fmt.Printf(gd.BlackHi+"["+gd.White+"%d"+gd.BlackHi+"]"+gd.Reset+gd.BlueHi+" %s\n"+gd.Reset, i+1, categories[i].CategoryName)
			yLoc1++
		}
		if count >= 14 {
			gd.MoveCursor(xLoc2, yLoc2)
			fmt.Printf(gd.BlackHi+"["+gd.White+"%d"+gd.BlackHi+"]"+gd.Reset+gd.BlueHi+" %s\n"+gd.Reset, i+1, categories[i].CategoryName)
			yLoc2++
		}
		count++
	}
	gd.MoveCursor(3, 24)
	prompt(w, alias, tleft, "blue")

}
