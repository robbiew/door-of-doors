package main

import (
	"bufio"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"

	gd "github.com/robbiew/door-of-doors/common"
	"golang.org/x/term"
)

var (
	dropPath   string
	menuKeys   []rune
	categories []CategoryList
	currTab    int
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
	currTab = 0
	initDb()

}

func header(w int, tab int) {
	if w == 80 {
		if tab == 0 {
			gd.PrintAnsiLoc("art/tab1.ans", 0, 1)
			fmt.Printf(gd.Reset)
			gd.MoveCursor(70, 2)
			fmt.Printf(gd.BgBlue + gd.BlueHi + " [Q] Quit " + gd.Reset)
		}
	}
	if w > 80 {
		fmt.Fprintf(os.Stdout, " ")
	}
}

func catHeader(w int, cat int) {

	if w == 80 {
		fn := fmt.Sprint(cat)
		gd.PrintAnsiLoc("art/"+fn+".ans", 0, 1)
		fmt.Printf(gd.Reset)
		gd.MoveCursor(70, 2)
		fmt.Printf(gd.BgRed + gd.RedHi + " [Q] Quit " + gd.Reset)
	}
	if w > 80 {
		fmt.Fprintf(os.Stdout, " ")
	}

}

func prompt(w int, alias string, timeLeft int, color string) {
	if w == 80 {
		gd.PrintStringLoc(alias+" - "+fmt.Sprint(timeLeft)+" mins left"+gd.Reset, 3, 23, gd.BlackHi, gd.BgBlack)
		gd.MoveCursor(3, 24)
		if color == "blue" {
			gd.PrintAnsi("art/prompt-blue.ans", 0, 1)
			gd.MoveCursor(6, 24)
			fmt.Printf(gd.BgBlue + "  " + gd.Reset)
			gd.MoveCursor(6, 24)
		}
		if color == "red" {
			gd.PrintAnsi("art/prompt-red.ans", 0, 1)
			gd.MoveCursor(6, 24)
			fmt.Printf(gd.BgRed + "  " + gd.Reset)
			gd.MoveCursor(6, 24)
		}
	}
	if w > 80 {
		fmt.Fprintf(os.Stdout, " ")
	}
}

func catMenu(db *sql.DB, w int, alias string, tleft int) {
	count := 0
	yLoc1 := 8
	yLoc2 := 8
	xLoc1 := 2
	xLoc2 := 34

	categories = categoryList(db)
	for i := 0; i < len(categories); i++ {

		if count < 14 {
			gd.MoveCursor(xLoc1, yLoc1)
			fmt.Printf(gd.BlackHi+"["+gd.White+"%d"+gd.BlackHi+"]"+gd.Reset+gd.BlueHi+" %s\n"+gd.Reset, i+1, categories[i].CategoryName)
			yLoc1++
		}
		if count >= 14 {
			gd.MoveCursor(xLoc2, yLoc2)
			fmt.Printf(gd.BlackHi+"["+gd.White+"%d"+gd.BlackHi+"]"+gd.Reset+gd.BlueHi+" %s\n"+gd.Reset, i+1, categories[i].CategoryName)
			yLoc2++
		}
		count++
	}
	gd.MoveCursor(3, 24)
	prompt(w, alias, tleft, "blue")

}

func main() {
	// Get door32.sys, h, w as user object
	u := gd.Initialize(dropPath)
	c := GetConfig()

	gd.ClearScreen()

	// Exit if no ANSI capabilities (sorry!)
	if u.Emulation != 1 {
		fmt.Println("Sorry, ANSI is required to use this...")
		time.Sleep(time.Duration(2) * time.Second)
		os.Exit(0)
	}

	// Categories menu
	db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File

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
		header(u.W, currTab)
		gd.PrintStringLoc(" "+c.Menu_Title+" "+c.Version+" "+gd.Reset, 2, 2, gd.BlueHi, gd.BgBlue)
		catMenu(db, u.W, u.Alias, u.TimeLeft)

		// show anything typed in prompt so far
		s := string(menuKeys)
		gd.MoveCursor(6, 24)
		fmt.Printf(gd.BgBlue+gd.BgBlueHi+"%v"+gd.Reset, s)

		r, _, err := in.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			break
		}
		if r == 'q' || r == 'Q' {
			term.Restore(0, state)
			os.Exit(0)
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
					doorMenu(db, i, u.W, u.Alias, u.TimeLeft)
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
					fmt.Printf(gd.BgBlue+gd.BgBlueHi+"%v"+gd.Reset, s)
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
					gd.MoveCursor(6, 24)
					s := string(menuKeys)
					fmt.Printf(gd.BgBlue+gd.BgBlueHi+"%v"+gd.Reset, s)
					gd.MoveCursor(6, 24)
					fmt.Printf("     ")
					gd.MoveCursor(6, 24)
					fmt.Printf(gd.Red+"Select from 1 to %v"+gd.Reset, len(categories)-1)
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
					fmt.Printf(gd.BgBlue+gd.BgBlueHi+"%v"+gd.Reset, s)
					time.Sleep(100 * time.Millisecond)
					menuKeys = nil
					gd.MoveCursor(6, 24)
					fmt.Printf("                   ")
					gd.MoveCursor(6, 24)
					// show list
					gd.ClearScreen()
					doorMenu(db, i, u.W, u.Alias, u.TimeLeft)
					continue
				}
			}
			continue
		}
		continue
	}

	// for i := 0; i < len(categories); i++ {
	// 	if string(char) == categories[i].CategoryId {

	// 		fmt.Println("pressed")
	// 		gd.Pause()
	// 	}

	// }
	// if string(char) == "b" || string(char) == "B" {
	// 	BbsLink("lord", u.UserNum, c.BL_Script)
	// }
	// if string(char) == "g" || string(char) == "G" {
	// 	GoldMine(u.Alias, c.GM_Tag, "lord", c.GM_Host, c.GM_Port, c.GM_script)
	// }
	// if string(char) == "d" || string(char) == "D" {
	// 	DoorParty("lord", u.Alias, c.DP_Script)
	// }

	// if string(char) == "q" || string(char) == "Q" || key == keyboard.KeyEsc {
	// 	break
	// }
	// gd.ClearScreen()
	// continue

}
