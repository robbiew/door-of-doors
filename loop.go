package main

import (
	"database/sql"
	"os"

	"github.com/eiannone/keyboard"
)

func loop(db *sql.DB, dataChan chan []byte, errorChan chan error, f *os.File, logFile string) {

	clearScreen()
	header(U.W, currMenu)
	catMenu(db, "none")
	currY = 0
	shortTimer.Stop()
	// log.Println("time stopped...")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		// if currMenu == " launch" {
		// 	clearScreen()
		// 	header(U.W, currMenu)
		// 	s := doorByServer(db)
		// 	moveCursor(2, 24)
		// 	fmt.Printf("Launch: %v on %v", currTitle, s[currY].ServerName)

		// 	// writeLog(f, U.Alias, s.DoorTitle, s.ServerName)
		// 	// if s.ServerId == "1" {
		// 	// 	goldMine(U.Alias, C.GM_Tag, s.DoorCode, C.GM_Host, C.GM_Port, C.GM_script)
		// 	// }
		// 	// if s.ServerId == "2" {
		// 	// 	bbsLink(s.DoorCode, U.UserNum, C.BL_Script)
		// 	// }
		// 	// clearScreen()
		// 	// header(U.W, currMenu)
		// 	// catMenu(db, "none")
		// }

		if key == keyboard.KeyEsc || string(char) == "Q" || string(char) == "q" {
			if currMenu == "door" {
				currY = 0
				currMenu = "category"
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
				continue
			}
			if currMenu == "server" {
				currY = 0
				currMenu = "door"
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
				continue
			}
			if currMenu == "launch" {
				currY = 0
				currMenu = "server"
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
				continue
			} else {
				cursorShow()
				break
			}
		}

		if key == keyboard.KeyArrowUp {
			catMenu(db, "up")
			continue
		}
		if key == keyboard.KeyArrowDown {
			catMenu(db, "down")
			continue
		}

		if key == keyboard.KeyEnter {
			clearScreen()
			if currMenu == "door" {
				currMenu = "server"
				currY = 0
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
			}
			if currMenu == "server" {
				currMenu = "server"
				currY = 0
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
			}

			if currMenu == "category" {
				currMenu = "door"
				currY = 0
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
			}
			continue
		}
		continue

	}

}
