package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"time"

	gd "github.com/robbiew/door-of-doors/common"
)

var (
	dropPath string
	menuKeys []rune
)

func init() {

	// Use FLAG to get command line paramenters
	pathPtr := flag.String("path", "", "path to door32.sys file")
	required := []string{"path"}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, "missing path to door32.sys directory: -%s \n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
	dropPath = *pathPtr

	initDb()

}

func main() {
	// Get door32.sys, h, w as user object
	u := gd.Initialize(dropPath)
	db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File

	gd.ClearScreen()

	// Exit if no ANSI capabilities (sorry!)
	if u.Emulation != 1 {
		fmt.Println("Sorry, ANSI is required to use this...")
		time.Sleep(time.Duration(2) * time.Second)
		os.Exit(0)
	}

	gd.MoveCursor(0, 0)
	gd.HeaderBar(u.W, u.Alias, u.TimeLeft)
	fmt.Println(gd.Reset)

	menuMain(db, "category", 0)
}

// func categoryDoorList(db *sql.DB, cat int) {
// 	gd.ClearScreen()

// 	gd.MoveCursor(2, 6)
// 	fmt.Println("Select a Game:")

// 	count := 0
// 	yLoc1 := 8
// 	yLoc2 := 8

// 	gd.MoveCursor(2, 8)
// 	doorsList := doorsByCategory(db, cat)

// 	// fmt.Println(doorsList)
// 	for i := 0; i < len(doorsList); i++ {

// 		if count < 10 {
// 			gd.MoveCursor(2, yLoc1)
// 			fmt.Printf("[%d] %s\n", i+1, doorsList[i].DoorTitle)
// 			yLoc1++
// 		}
// 		if count >= 10 {
// 			gd.MoveCursor(40, yLoc2)
// 			fmt.Printf("[%d] %s\n", i+1, doorsList[i].DoorTitle)
// 			yLoc2++
// 		}
// 		count++
// 	}
// 	// fd 0 is stdin
// 	state, err := term.MakeRaw(0)
// 	if err != nil {
// 		log.Fatalln("setting stdin to raw:", err)
// 	}
// 	defer func() {
// 		if err := term.Restore(0, state); err != nil {
// 			log.Println("warning, failed to restore terminal:", err)
// 		}
// 	}()

// 	gd.MoveCursor(2, 23)
// 	fmt.Print("-> ")
// 	in := bufio.NewReader(os.Stdin)
// 	for {
// 		r, _, err := in.ReadRune()
// 		if err != nil {
// 			log.Println("stdin:", err)
// 			break
// 		}
// 		if r == 'q' || r == 'Q' {
// 			break
// 		}
// 		// User hit return on a single digit number in the list, let's load a category
// 		if r == '\n' || r == '\r' {

// 			s := string(menuKeys)
// 			i, err := strconv.Atoi(s)
// 			if err != nil {
// 				panic(err)
// 			}
// 			gd.MoveCursor(5, 23)
// 			fmt.Printf("View Category %v...", s)
// 			menuKeys = nil
// 			time.Sleep(1 * time.Second)
// 			gd.MoveCursor(5, 23)
// 			fmt.Printf("                   ")
// 			gd.MoveCursor(5, 23)
// 			// show list
// 			categoryDoorList(db, i)
// 			continue
// 		}
// 		// Make sure it's a number greater than 0, otherwise don't respond
// 		if unicode.IsDigit(r) {
// 			if int(r-'0') != 0 {
// 				if len(menuKeys) <= 0 {
// 					menuKeys = append(menuKeys, r)
// 					s := string(menuKeys)
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf("%v", s)
// 					continue
// 				}
// 			}
// 			// we collect a key press in raw mode, save it to a slice, then print the slice
// 			if len(menuKeys) == 1 {
// 				menuKeys = append(menuKeys, r)
// 				s := string(menuKeys)
// 				i, err := strconv.Atoi(s)
// 				if err != nil {
// 					panic(err)
// 				}
// 				// User entered a number greater than what's in the list
// 				if i > len(doorsList)-1 {
// 					menuKeys = append(menuKeys, r)
// 					gd.MoveCursor(5, 23)
// 					s := string(menuKeys)
// 					fmt.Printf("%v", s)
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf("     ")
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf(gd.Red+"Select from 1 to %v"+gd.Reset, len(doorsList)-1)
// 					time.Sleep(1 * time.Second)
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf("                               ")
// 					gd.MoveCursor(5, 23)
// 					// wipe the slice so it starts over
// 					menuKeys = nil
// 					continue
// 					// second key, it's valid, so load the category list!
// 				} else {
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf("     ")
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf("%v", s)
// 					gd.MoveCursor(5, 23)
// 					time.Sleep(100 * time.Millisecond)
// 					fmt.Printf("View Category %v...", s)
// 					menuKeys = nil
// 					time.Sleep(1 * time.Second)
// 					gd.MoveCursor(5, 23)
// 					fmt.Printf("                   ")
// 					gd.MoveCursor(5, 23)
// 					// show list
// 					categoryDoorList(db, i)
// 					continue
// 				}
// 			}
// 			continue
// 		}
// 		continue
// 	}

// }
