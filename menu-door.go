package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"

	gd "github.com/robbiew/door-of-doors/common"
)

func doorMenu(db *sql.DB, cat int, w int, alias string, tleft int, dataChan chan []byte, errorChan chan error) {
	count := 0
	yLoc1 := 7
	yLoc2 := 7
	xLoc1 := 2
	xLoc2 := 29

	gd.ClearScreen()

	gd.MoveCursor(0, 0)
	catHeader(w, cat)

	doorsList := doorsByCategory(db, cat)
	for i := 0; i < len(doorsList); i++ {

		if count < 15 {
			gd.MoveCursor(xLoc1, yLoc1)
			fmt.Printf(gd.BlackHi+"["+gd.White+"%d"+gd.BlackHi+"]"+gd.Reset+gd.RedHi+" %s\n"+gd.Reset, i+1, doorsList[i].DoorTitle)
			yLoc1++
		}
		if count >= 15 {
			gd.MoveCursor(xLoc2, yLoc2)
			fmt.Printf(gd.BlackHi+"["+gd.White+"%d"+gd.BlackHi+"]"+gd.Reset+gd.RedHi+" %s\n"+gd.Reset, i+1, doorsList[i].DoorTitle)
			yLoc2++
		}
		count++
	}

	gd.MoveCursor(3, 24)
	prompt(w, alias, tleft, "red")

	for {
		shortTimer.Stop()
		log.Println("time stopped...")

		go readWrapper(dataChan, errorChan)

		// show anything typed in prompt so far
		s := string(menuKeys)
		gd.MoveCursor(6, 24)
		fmt.Printf(gd.BgRed+gd.BgRedHi+"%v"+gd.Reset, s)

		r, _ := utf8.DecodeRune(<-dataChan)

		if r == 'q' || r == 'Q' {
			gd.ClearScreen()
			break
		}

		if r == '\b' {
			if len(menuKeys) > 0 {
				menuKeys = menuKeys[:len(menuKeys)-1]
			}
			gd.MoveCursor(6, 24)

		}

		// User hit return on a single digit number in the list, let's load a category
		if len(menuKeys) != 0 && r == '\n' || r == '\r' {
			s := string(menuKeys)
			if len(menuKeys) > 0 {
				i, err := strconv.Atoi(s)
				if err != nil {
					fmt.Println(err)
				}
				menuKeys = nil
				if i != 0 {
					gd.MoveCursor(5, 24)
					fmt.Printf("                   ")
					gd.MoveCursor(5, 24)
					// show list
					gd.ClearScreen()
					shortTimer.Stop()
					log.Println("time stopped...")
					// doorMenu(db, i, u.W, u.Alias, u.TimeLeft, dataChan, errorChan)
				}
			}
			continue
		}

		// Make sure it's a number greater than 0, otherwise don't respond
		if unicode.IsDigit(r) {
			if int(r-'0') != 0 {
				if len(menuKeys) <= 0 {
					menuKeys = append(menuKeys, r)
					s := string(menuKeys)
					gd.MoveCursor(6, 24)
					fmt.Printf(gd.BgRed+gd.BgRedHi+"%v"+gd.Reset, s)
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
				if i > len(doorsList)-1 {
					menuKeys = append(menuKeys, r)
					gd.MoveCursor(6, 24)
					s := string(menuKeys)
					fmt.Printf("%v", s)
					gd.MoveCursor(6, 24)
					fmt.Printf("     ")
					gd.MoveCursor(6, 24)
					fmt.Printf(gd.RedHi+" Select from 1 to %v"+gd.Reset, len(doorsList))
					time.Sleep(1 * time.Second)
					gd.MoveCursor(6, 24)
					fmt.Printf("                               ")
					gd.MoveCursor(6, 24)
					// wipe the slice so it starts over
					menuKeys = nil
					continue

				} else {
					// second key, it's valid, so load the category list!
					gd.MoveCursor(6, 24)
					fmt.Printf("     ")
					gd.MoveCursor(6, 24)
					fmt.Printf(gd.BgRed+gd.BgRedHi+"%v"+gd.Reset, s)
					time.Sleep(100 * time.Millisecond)
					menuKeys = nil
					gd.MoveCursor(6, 24)
					fmt.Printf("                   ")
					gd.MoveCursor(6, 24)
					// show list
					gd.ClearScreen()
					shortTimer.Stop()
					log.Println("time stopped...")
					// doorMenu(db, i, u.W, u.Alias, u.TimeLeft, dataChan, errorChan)
					continue
				}
			}
			continue
		}
		continue
	}

}
