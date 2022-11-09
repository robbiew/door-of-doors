package main

import (
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/term"
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
	BL_Enabled string
	BL_Script  string
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
	ServerId   int
	ServerName string
	Desc       string
	Year       string
	DoorCode   string
}

var (
	menuKeys    []rune
	categories  []CategoryList
	doorsList   []DoorsList
	serversList []ServersList

	currCat     int
	currCatName string
	currCode    string
	currDoor    int
	currTitle   string
	currPage    int

	paginator bool

	lenList int

	shortTimer *time.Timer
	menuType   string
	idle       int

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
	currCat = 0
	currPage = 1

	// entry menu
	menuType = "category"

	//SQlite db
	initDb()

	// Get door32.sys, h, w as user object
	initDrop(dropPath)

	// Get ap config from config.ini
	initIni()
}

// Main input loop
func readWrapper(dataChan chan []byte, errorChan chan error) {
	shortTimer = newTimer(idle, func() {
		fmt.Println("\r\nYou've been idle for too long... exiting!")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	})

	// fd 0 is stdin - set to raw mode so return doesn't have to be pressed
	state, err := term.MakeRaw(0)
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	defer func() {
		if err := term.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	buf := make([]byte, 1024)
	reqLen, err := os.Stdin.Read(buf)
	if err != nil {
		errorChan <- err
		return
	}
	dataChan <- buf[:reqLen]
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

	db, _ := sql.Open("sqlite3", "./data.db") // Open the created SQLite File

	loop(db, dataChan, errorChan)
}
