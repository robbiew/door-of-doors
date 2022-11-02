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

func doorMenu(db *sql.DB, cat int, w int, alias string, tleft int) {
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

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			break
		}
		if r == 'q' || r == 'Q' {
			gd.ClearScreen()
			break
		}

		if r == '\b' {
			if len(menuKeys) > 0 {
				menuKeys = menuKeys[:len(menuKeys)-1]
			}
			log.Println("backscpace")
			gd.MoveCursor(5, 23)
			time.Sleep(time.Second * 1)
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
					gd.MoveCursor(5, 23)
					fmt.Printf("                   ")
					gd.MoveCursor(5, 23)
					// show list
					gd.ClearScreen()
					// doorMenu(db, i, w, alias, tleft)
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
				if i > len(doorsList)-1 {
					menuKeys = append(menuKeys, r)
					gd.MoveCursor(5, 23)
					s := string(menuKeys)
					fmt.Printf("%v", s)
					gd.MoveCursor(5, 23)
					fmt.Printf("     ")
					gd.MoveCursor(5, 23)
					fmt.Printf(gd.Red+"Select from 1 to %v"+gd.Reset, len(doorsList)-1)
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
					// doorMenu(db, i)
					continue
				}
			}
			continue
		}
		continue
	}

}
