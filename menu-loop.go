package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

func loop(db *sql.DB, dataChan chan []byte, errorChan chan error) {
	if menuType == "category" {
		mainHeader(U.W, currTab)
		catMenu(db)
	}
	if menuType == "door" {
		doorMenu(db, dataChan, errorChan)
	}

	for {
		shortTimer.Stop()
		log.Println("time stopped...")

		go readWrapper(dataChan, errorChan)

		// show anything typed in prompt so far
		s := string(menuKeys)
		MoveCursor(6, 24)
		fmt.Printf(BgBlue+BgBlueHi+"%v"+Reset, s)

		r, _ := utf8.DecodeRune(<-dataChan)

		if r == 'q' || r == 'Q' {
			if menuType == "category" {
				log.Println("category exit")
				os.Exit(0)
			}
			if menuType == "door" {
				menuType = "category"
				catMenu(db)
			}
		}
		if r == '\b' {
			if len(menuKeys) > 0 {
				menuKeys = menuKeys[:len(menuKeys)-1]
			}
			MoveCursor(6, 24)
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
					MoveCursor(5, 24)
					fmt.Printf("                   ")
					MoveCursor(5, 24)
					// show list
					ClearScreen()
					shortTimer.Stop()
					log.Println("time stopped...")
					menuType = "door"
					currCat = i
					doorMenu(db, dataChan, errorChan)
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
					MoveCursor(6, 24)
					fmt.Printf(BgBlue+BgBlueHi+"%v"+Reset, s)
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
				if i > len(categories)-1 {
					menuKeys = append(menuKeys, r)
					MoveCursor(6, 24)
					s := string(menuKeys)
					fmt.Printf(BgBlue+BgBlueHi+"%v"+Reset, s)
					MoveCursor(6, 24)
					fmt.Printf("     ")
					MoveCursor(6, 24)
					fmt.Printf(Red+" Select from 1 to %v"+Reset, len(categories))
					time.Sleep(1 * time.Second)
					MoveCursor(6, 24)
					fmt.Printf("                               ")
					MoveCursor(6, 24)
					// wipe the slice so it starts over
					menuKeys = nil
					continue

				} else {
					// second key, it's valid, so load the category list!
					MoveCursor(6, 24)
					fmt.Printf("     ")
					MoveCursor(6, 24)
					fmt.Printf(BgBlue+BgBlueHi+"%v"+Reset, s)
					time.Sleep(100 * time.Millisecond)
					menuKeys = nil
					MoveCursor(6, 24)
					fmt.Printf("                   ")
					MoveCursor(6, 24)
					// show list
					ClearScreen()
					shortTimer.Stop()
					log.Println("time stopped...")
					menuType = "door"
					currCat = i
					doorMenu(db, dataChan, errorChan)
					continue
				}
			}
			continue
		}
		continue
	}
}
