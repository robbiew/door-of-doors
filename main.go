package main

import (
	"bufio"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"

	gd "github.com/robbiew/door-of-doors/common"
	"golang.org/x/term"
)

var (
	dropPath string
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
	// c := GetConfig()

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

	// Categories menu

	db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
	// defer db.Close()

	gd.MoveCursor(2, 6)
	fmt.Println("Select a Category:")

	count := 0
	yLoc1 := 8
	yLoc2 := 8

	gd.MoveCursor(2, 8)
	categories := categoryList(db)
	for i := 1; i < len(categories); i++ {

		if count < 10 {
			gd.MoveCursor(2, yLoc1)
			fmt.Printf("[%d] %s\n", i, categories[i].CategoryName)
			yLoc1++
		}
		if count > 10 {
			gd.MoveCursor(40, yLoc2)
			fmt.Printf("[%d] %s\n", i, categories[i].CategoryName)
			yLoc2++
		}
		count++
	}

	// keys := doorsByCategory(db, 1)
	// fmt.Println(keys)

	for {

		gd.MoveCursor(0, 0)
		gd.HeaderBar(u.W, u.Alias, u.TimeLeft)
		fmt.Println(gd.Reset)

		gd.MoveCursor(2, 23)

		fmt.Print("-> ")

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
			gd.MoveCursor(0, 0)
			gd.HeaderBar(u.W, u.Alias, u.TimeLeft)
			fmt.Println(gd.Reset)

			gd.MoveCursor(2, 23)

			fmt.Print("-> ")
			r, _, err := in.ReadRune()
			if err != nil {
				log.Println("stdin:", err)
				break
			}
			// fmt.Printf("read rune %q\r\n", r)
			if r == 'q' || r == 'Q' {
				term.Restore(0, state)
				os.Exit(0)
				continue
			}
			if unicode.IsDigit(r) {
				fmt.Printf("number: %q       ", r)
				continue

			}
			continue
		}

		// inputs = append(inputs, string(xxx)
		// if len(inputs) > 2 {
		// 	break
		// }

		// fmt.Println("Your number is:", i)

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
