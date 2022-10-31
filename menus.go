package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"

	gd "github.com/robbiew/door-of-doors/common"
	"golang.org/x/term"
)

func menuOptions(db *sql.DB, menuType string, selection int) int {

	var catList []CategoryList
	var doorList []DoorsList
	var max int

	if menuType == "category" {
		catList = categoryList(db)
		max = len(catList)
	}
	if menuType == "door" {
		doorList = doorsByCategory(db, selection)
	}
	if menuType == "server" {
	}

	count := 0
	yLoc1 := 8
	yLoc2 := 8

	gd.MoveCursor(2, 8)

	for i := 0; i < max; i++ {
		if count < 10 {
			gd.MoveCursor(2, yLoc1)
		}
		if count >= 10 {
			gd.MoveCursor(40, yLoc2)
		}
		if menuType == "category" {
			fmt.Printf("[%d] %s\n", i+1, catList[i].CategoryName)
		}
		if menuType == "door" {
			fmt.Printf("[%d] %s\n", i+1, doorList[i].DoorTitle)
		}
		if count < 10 {
			yLoc1++
		}
		if count >= 10 {
			yLoc2++
		}

		count++
	}
	return max
}

func menuMain(db *sql.DB, menuType string, selection int) {

	gd.MoveCursor(2, 6)
	fmt.Println("Select:")

	maxLen := menuOptions(db, menuType, 0)

	gd.MoveCursor(0, 0)
	fmt.Println(gd.Reset)
	// fd 0 is stdin
	state, err := term.MakeRaw(0)
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	defer func() {
		if err := term.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	gd.MoveCursor(2, 23)
	fmt.Print("-> ")
	in := bufio.NewReader(os.Stdin)

	for {
		gd.MoveCursor(2, 6)
		fmt.Println("Select:")
		if menuType == "category" {
			menuOptions(db, "category", 0)
		}
		if menuType == "door" {
			menuOptions(db, "door", selection)
		}

		gd.MoveCursor(0, 0)
		fmt.Println(gd.Reset)

		r, _, err := in.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			break
		}
		if r == 'q' || r == 'Q' {
			term.Restore(0, state)
			os.Exit(0)
			continue
		}
		// User hit return on a single digit number in the list, let's load a category
		if r == '\n' || r == '\r' {

			s := string(menuKeys)
			// i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			gd.MoveCursor(5, 23)
			fmt.Printf("View Category %v...", s)
			menuKeys = nil
			time.Sleep(1 * time.Second)
			gd.MoveCursor(5, 23)
			fmt.Printf("                   ")
			gd.MoveCursor(5, 23)
			// show list
			menuOptions(db, "door", selection)
			continue
		}
		// Make sure it's a number greater than 0, otherwise don't respond
		if unicode.IsDigit(r) {
			if int(r-'0') != 0 {
				if len(menuKeys) <= 0 {
					menuKeys = append(menuKeys, r)
					s := string(menuKeys)
					gd.MoveCursor(5, 23)
					fmt.Printf("%v", s)
					continue
				}
			}
			// we collect a key press in raw mode, save it to a slice, then print the slice
			if len(menuKeys) == 1 {
				menuKeys = append(menuKeys, r)
				s := string(menuKeys)
				i, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				// User entered a number greater than what's in the list
				if i > unicode.MaxASCII-1 {
					menuKeys = append(menuKeys, r)
					gd.MoveCursor(5, 23)
					s := string(menuKeys)
					fmt.Printf("%v", s)
					gd.MoveCursor(5, 23)
					fmt.Printf("     ")
					gd.MoveCursor(5, 23)
					fmt.Printf(gd.Red+"Select from 1 to %v"+gd.Reset, maxLen-1)
					time.Sleep(1 * time.Second)
					gd.MoveCursor(5, 23)
					fmt.Printf("                               ")
					gd.MoveCursor(5, 23)
					// wipe the slice so it starts over
					menuKeys = nil
					continue
					// second key, it's valid, so load the category list!
				} else {
					gd.MoveCursor(5, 23)
					fmt.Printf("     ")
					gd.MoveCursor(5, 23)
					fmt.Printf("%v", s)
					gd.MoveCursor(5, 23)
					time.Sleep(100 * time.Millisecond)
					fmt.Printf("View Category %v...", s)
					menuKeys = nil
					time.Sleep(1 * time.Second)
					gd.MoveCursor(5, 23)
					fmt.Printf("                   ")
					gd.MoveCursor(5, 23)
					// show list
					menuOptions(db, "door", i)
					continue
				}
			}
			continue
		}
		continue
	}

}
