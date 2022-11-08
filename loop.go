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
		catMenu(db)

	}
	if menuType == "door" {
		doorMenu(db)

	}
	if menuType == "server" {
		serverMenu(db)

	}

	for {
		shortTimer.Stop()
		log.Println("time stopped...")

		go readWrapper(dataChan, errorChan)

		// show anything typed in prompt so far
		s := string(menuKeys)
		moveCursor(6, 24)
		fmt.Printf(bgRed+redHi+"%v"+reset, s)

		r, _ := utf8.DecodeRune(<-dataChan)

		if r == 'q' || r == 'Q' {
			if menuType == "category" {
				log.Println("category exit")
				os.Exit(0)
			}
			if menuType == "door" {
				menuType = "category"
				catMenu(db)
				continue
			}
			if menuType == "server" {
				menuType = "door"
				doorMenu(db)
				continue
			}
		}
		if r == '\b' {
			if len(menuKeys) > 0 {
				menuKeys = menuKeys[:len(menuKeys)-1]
			}
			moveCursor(6, 24)
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
					moveCursor(5, 24)
					fmt.Printf("                   ")
					moveCursor(5, 24)
					// show list
					clearScreen()
					shortTimer.Stop()
					log.Println("time stopped...")
					if menuType == "category" {
						menuType = "door"
						currCat = i
						doorMenu(db)
						continue
					}
					if menuType == "door" {
						currDoor = i - 1
						menuType = "server"
						serverMenu(db)
						continue
					}

					if menuType == "server" {
						menuType = "door"
						currCat = i
						doorMenu(db)
						continue
					}

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
					moveCursor(6, 24)
					fmt.Printf(bgRed+redHi+"%v"+reset, s)
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

				if i > lenList {
					menuKeys = append(menuKeys, r)
					moveCursor(6, 24)
					s := string(menuKeys)
					fmt.Printf(bgRed+redHi+"%v"+reset, s)
					moveCursor(6, 24)
					fmt.Printf("     ")
					moveCursor(6, 24)
					fmt.Printf(red+" Select from 1 to %v"+reset, lenList)
					time.Sleep(1 * time.Second)
					moveCursor(6, 24)
					fmt.Printf("                               ")
					moveCursor(6, 24)
					fmt.Printf(bgRed + "  " + reset)

					// wipe the slice so it starts over
					menuKeys = nil
					continue

				} else {
					// second key, it's valid, so load the category list!
					moveCursor(6, 24)
					fmt.Printf("     ")
					moveCursor(6, 24)
					fmt.Printf(bgRed+redHi+"%v"+reset, s)
					time.Sleep(100 * time.Millisecond)
					menuKeys = nil
					moveCursor(6, 24)
					fmt.Printf("                   ")
					moveCursor(6, 24)
					clearScreen()

					shortTimer.Stop()
					log.Println("time stopped...")
					if menuType == "category" {
						menuType = "door"
						currCat = i
						doorMenu(db)
						continue
					}
					if menuType == "door" {
						currDoor = i - 1
						menuType = "server"
						serverMenu(db)
						continue
					}

					if menuType == "server" {
						menuType = "door"
						currCat = i
						doorMenu(db)
						continue
					}
					continue
				}
			}
			continue
		}
		continue
	}
}
