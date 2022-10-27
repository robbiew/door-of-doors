package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	gd "github.com/robbiew/door-of-doors/common"
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
	c := GetConfig()

	gd.ClearScreen()
	gd.MoveCursor(0, 0)

	// Exit if no ANSI capabilities (sorry!)
	if u.Emulation != 1 {
		fmt.Println("Sorry, ANSI is required to use this...")
		time.Sleep(time.Duration(2) * time.Second)
		os.Exit(0)
	}

	// A reliable keyboard library to detect key presses
	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {

		gd.MoveCursor(0, 0)
		gd.HeaderBar(u.W)

		gd.MoveCursor(0, 0)

		// A Test Menu
		gd.MoveCursor(0, 3)

		// fmt.Fprintf(os.Stdout, gd.Cyan+"["+gd.YellowHi+"Q"+gd.Cyan+"] "+gd.Reset+gd.Magenta+"Quit\r\n")
		// fmt.Fprintf(os.Stdout, gd.Cyan+"["+gd.YellowHi+"G"+gd.Cyan+"] "+gd.Reset+gd.Magenta+"Gold Mine LORD test\r\n")
		// fmt.Fprintf(os.Stdout, gd.Cyan+"["+gd.YellowHi+"B"+gd.Cyan+"] "+gd.Reset+gd.Magenta+"BBSLink LORD test\r\n")
		// fmt.Fprintf(os.Stdout, gd.Cyan+"["+gd.YellowHi+"D"+gd.Cyan+"] "+gd.Reset+gd.Magenta+"Door Party LORD test\r\n")
		sqliteDatabase, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
		defer sqliteDatabase.Close()                          // Defer Closing the database
		displayDoors(sqliteDatabase)
		fmt.Fprintf(os.Stdout, gd.Reset+"\r\nCommand? ")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if string(char) == "b" || string(char) == "B" {
			BbsLink("lord", u.UserNum, c.BL_Script)
		}
		if string(char) == "g" || string(char) == "G" {
			GoldMine(u.Alias, c.GM_Tag, "lord", c.GM_Host, c.GM_Port, c.GM_script)
		}
		if string(char) == "d" || string(char) == "D" {
			DoorParty("lord", u.Alias, c.DP_Script)
		}

		if string(char) == "q" || string(char) == "Q" || key == keyboard.KeyEsc {
			break
		}
		gd.ClearScreen()
		continue
	}
}
