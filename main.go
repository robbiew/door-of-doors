package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	gd "github.com/robbiew/door-of-doors/common"
	"github.com/robbiew/door-of-doors/go-input"
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
	gd.MoveCursor(0, 0)

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
	gd.MoveCursor(0, 6)
	db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
	// defer db.Close()                          // Defer Closing the database

	categories := doorCategories(db)
	for i := 0; i < len(categories); i++ {
		fmt.Printf("[%d] %s\n", i, categories[i].CategoryName)
	}

	// keys := doorsByCategory(db, 1)
	// fmt.Println(keys)

	fmt.Fprintf(os.Stdout, gd.Reset+"\r\nCommand? ")

	for {

		ui := &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		}

		query := "Which language do you prefer to use?"
		lang, err := ui.Select(query, []string{"go", "Go", "golang"}, &input.Options{

			Loop: true,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Answer is %s\n", lang)
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
