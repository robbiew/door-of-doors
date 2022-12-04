package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func loop(db *sql.DB, dataChan chan []byte, errorChan chan error, f *os.File, logFile string) {

	clearScreen()
	header(U.W)
	catMenu(db, "none", "category")

	shortTimer.Stop()
	// log.Println("time stopped...")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc || string(char) == "Q" || string(char) == "q" {
			if menuType == "door" {
				menuType = "category"
				currY = saveY
				clearScreen()
				header(U.W)
				catMenu(db, "none")
				continue
			}
			if menuType == "server" {
				menuType = "door"
				currY = saveY
				clearScreen()
				header(U.W)
				catMenu(db, "none")
				continue
			}
			if menuType == "category" {
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
			saveY = currY
			currY = 0

			if menuType == "category" {
				menuType = "door"

			}
			if menuType == "door" {
				menuType = "server"

			}
			if menuType == "server" {
				fmt.Print("launching door...")

			}
			// clearScreen()
			header(U.W)
			catMenu(db, "none")
			return
		}

	}

}

// 		shortTimer.Stop()
// 		// log.Println("time stopped...")
// 		if menuType == "category" {
// 			menuType = "door"
// 			currCat = i
// 			doorMenu(db)
// 			continue
// 		}
// 		if menuType == "door" {
// 			currDoor = i - 1
// 			menuType = "server"
// 			serverMenu(db)
// 			continue
// 		}

// 		if menuType == "server" {
// 			menuType = "door"
// 			currCat = i
// 			s := serversList[i-1]
// 			writeLog(f, U.Alias, s.DoorTitle, s.ServerName)
// 			clearScreen()
// 			if s.ServerId == "1" {
// 				goldMine(U.Alias, C.GM_Tag, s.DoorCode, C.GM_Host, C.GM_Port, C.GM_script)
// 			}
// 			if s.ServerId == "2" {
// 				bbsLink(s.DoorCode, U.UserNum, C.BL_Script)
// 			}
// 			clearScreen()

// 			doorMenu(db)
// 			continue
// 		}
// 		continue
// 	}

// }
// continue
// }
// 		continue
// 	}

// }
