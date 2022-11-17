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
	lenList = len(categories)

	recordsPerCol := 14

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 2
	xLoc2 := 34

	for i := 0; i < len(categories); i++ {
		if count < recordsPerCol {
			moveCursor(xLoc1, yLoc1)
			if count < 9 {
				fmt.Printf(white+" %d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
			} else {
				fmt.Printf(white+"%d"+blackHi+"..."+reset+redHi+"%s"+reset, i+1, categories[i].CategoryName)
			}
			yLoc1++
		}
		if count >= recordsPerCol {
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
