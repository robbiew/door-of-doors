package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

type User struct {
	Alias     string
	TimeLeft  int
	Emulation int
	NodeNum   int
	UserNum   int
	H         int
	W         int
}

type DoorConfig struct {
	Menu_Title string
	Version    string
	GM_Host    string
	GM_Port    string
	GM_Tag     string
	GM_Enabled string
	GM_script  string
	DP_Enabled string
	DP_Script  string
	BL_Script  string
	BL_Enabled string
}

type CategoryList struct {
	CategoryId   int
	CategoryName string
	CategoryCode string
}

type DoorsList struct {
	DoorTitle string
}

type ServersList struct {
	DoorTitle  string
	ServerId   string
	ServerName string
	Desc       string
	Year       string
	DoorCode   string
}

var (
	categories  []CategoryList
	doorsList   []DoorsList
	serversList []ServersList

	currCat     int
	currCatName string
	currDoor    int
	currTitle   string
	currY       int
	saveY       int
	currStart   int

	lenList    int
	listHeight int
	menuType   string

	shortTimer *time.Timer

	idle int

	U *User
	C *DoorConfig
)

func init() {

	// if user doesn't press a key in X seconds
	idle = 240

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

	dropPath := *pathPtr

	//  initial state
	currCat = 0
	currY = 0
	saveY = 0
	currStart = 0
	listHeight = 10

	// entry menu
	menuType = "category"
	cursorHide()

	//SQlite db
	initDb()

	// Get door32.sys, h, w as user object
	initDrop(dropPath)

	// Get ap config from config.ini
	initIni()
}

func main() {
	errorChan := make(chan error)
	dataChan := make(chan []byte)

	clearScreen()

	// Exit if no ANSI capabilities (sorry!)
	if U.Emulation != 1 {
		fmt.Println("Sorry, ANSI is required to use this...")
		time.Sleep(time.Duration(2) * time.Second)
		os.Exit(0)
	}

	// Start the idle timer
	shortTimer = newTimer(idle, func() {
		fmt.Println("\r\nYou've been idle for too long... exiting!")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	})

	// Logging
	logFille := "activity.log"
	f, err := os.OpenFile(logFille, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// Database
	db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File
	categories = categoryList(db)
	lenList = len(categories)

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	showStats()
	clearScreen()
	loop(db, dataChan, errorChan, f, logFille)
}
