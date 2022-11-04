package main

import (
	"database/sql"
	"fmt"
)

func catMenu(db *sql.DB) {
	currCode = "MAIN1"
	ClearScreen()
	MoveCursor(0, 0)
	header(U.W)

	categories = categoryList(db)

	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 3
	xLoc2 := 36

	PrintStringLoc(" "+C.Menu_Title+" "+C.Version+" "+Reset, 2, 2, RedHi, BgBlack)

	for i := 0; i < len(categories); i++ {
		if count < 14 {
			MoveCursor(xLoc1, yLoc1)
			if count < 9 {
				fmt.Printf(White+" %d"+BlackHi+"..."+Reset+RedHi+"%s"+Reset, i+1, categories[i].CategoryName)
			} else {
				fmt.Printf(White+"%d"+BlackHi+"..."+Reset+RedHi+"%s"+Reset, i+1, categories[i].CategoryName)
			}
			yLoc1++
		}
		if count >= 14 {
			MoveCursor(xLoc2, yLoc2)
			fmt.Printf(White+"%d"+BlackHi+"..."+Reset+RedHi+"%s"+Reset, i+1, categories[i].CategoryName)
			yLoc2++
		}
		count++
	}
	MoveCursor(3, 24)
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
