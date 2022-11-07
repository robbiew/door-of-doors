package main

import (
	"database/sql"
	"fmt"
)

func catMenu(db *sql.DB) {
	clearScreen()
	currCode = "MAIN"
	header(U.W)

	moveCursor(3, 6)
	fmt.Print(whiteHi + "Select a category:" + reset)
	categories = categoryList(db)

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 3
	xLoc2 := 36

	// printStringLoc(" "+C.Menu_Title+" "+C.Version+" "+reset, 2, 2, redHi, bgBlack)

	for i := 0; i < len(categories); i++ {
		if count < 14 {
			moveCursor(xLoc1, yLoc1)
			if count < 9 {
				fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
			} else {
				fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
			}
			yLoc1++
		}
		if count >= 14 {
			moveCursor(xLoc2, yLoc2)
			fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
			yLoc2++
		}
		count++
	}
	moveCursor(3, 24)
	prompt("red")

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
