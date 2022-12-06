package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func loop(db *sql.DB, dataChan chan []byte, errorChan chan error, f *os.File, logFile string) {

	clearScreen()
	catMenu(db, "none")
	currY = 0
	shortTimer.Stop()

	// log.Println("time stopped...")

	for {

		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc || string(char) == "Q" || string(char) == "q" {
			if currMenu == "door" {
				currY = 0
				currStart = 0
				currMenu = "category"
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
				continue
			}

			if currMenu == "server" {
				currY = 0
				currStart = 0
				currMenu = "door"
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
			if currMenu == "category" {
				categories = nil
				currY = 0
				currStart = 0
				currMenu = "door"
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
				continue
			}
			if currMenu == "door" {
				doors = nil
				currY = 0
				currStart = 0
				currMenu = "server"
				clearScreen()
				header(U.W, currMenu)
				catMenu(db, "none")
				continue
			}
			if currMenu == "server" {
				currMenu = "server"

				writeLog(f, U.Alias, servers[currY].DoorTitle, servers[currY].ServerName)

				moveCursor(2, 24)
				fmt.Printf(yellow+"Launching %v on %v"+reset, servers[currY].DoorTitle, servers[currY].ServerName)
				moveCursor(0, 0)
				time.Sleep(time.Second * 1)

				cursorShow()

				keyboard.Close()

				if servers[currY].ServerId == "1" {
					goldMine(U.Alias, C.GM_Tag, servers[currY].DoorCode, C.GM_Host, C.GM_Port, C.GM_script)
					clearScreen()
					cursorHide()
					header(U.W, currMenu)
					catMenu(db, "none")
					continue
				}
				if servers[currY].ServerId == "2" {
					bbsLink(servers[currY].DoorCode, U.UserNum, C.BL_Script)
					clearScreen()
					cursorHide()
					header(U.W, currMenu)
					catMenu(db, "none")
					continue
				}
				if servers[currY].ServerId == "3" {
					doorParty(servers[currY].DoorCode, U.Alias, C.DP_Script)
					clearScreen()
					cursorHide()
					header(U.W, currMenu)
					catMenu(db, "none")
					continue
				}

			}

			continue
		}
	}
}
