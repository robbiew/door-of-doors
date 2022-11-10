package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"gopkg.in/ini.v1"
)

// CREDIT TO https://github.com/k0kubun/go-ansi for some of these sequences.

const (
	Esc = "\u001B["
	Osc = "\u001B]"
	Bel = "\u0007"
)

// Common fonts, supported by SyncTerm
const (
	Mosoul          = Esc + "0;38 D"
	Potnoodle       = Esc + "0;37 D"
	Microknight     = Esc + "0;41 D"
	Microknightplus = Esc + "0;39 D"
	Topaz           = Esc + "0;42 D"
	Topazplus       = Esc + "0;40 D"
	Ibm             = Esc + "0;0 D"
	Ibmthin         = Esc + "0;26 D"
)

// Symbols
var (
	Heart        = string([]rune{'\u0003'})
	ArrowUpDown  = string([]rune{'\u0017'})
	ArrowUp      = string([]rune{'\u0018'})
	ArrowDown    = string([]rune{'\u0019'})
	ArrowDownFat = string([]rune{'\u001F'})
	ArrowRight   = string([]rune{'\u0010'})
	ArrowLeft    = string([]rune{'\u0011'})
	Block        = string([]rune{'\u0219'})

	modalH int // in case height is odd
	modalW int // in case width is odd
)

// Common ANSI escapes sequences. These should be used when the desired action
// is only needed once; otherwise, use the functions (e.g. moving a cursor
// several lines/columns). See: https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
const (
	cursorBackward = Esc + "D"
	cursorPrevLine = Esc + "F"
	cursorLeft     = Esc + "G"
	cursorTop      = Esc + "d"
	cursorTopLeft  = Esc + "H"

	cursorBlinkEnable  = Esc + "?12h"
	cursorBlinkDisable = Esc + "?12I"

	scrollUp   = Esc + "S"
	scrollDown = Esc + "T"

	textInsertChar = Esc + "@"
	textDeleteChar = Esc + "P"
	textEraseChar  = Esc + "X"
	textInsertLine = Esc + "L"
	textDeleteLine = Esc + "M"

	eraseRight  = Esc + "K"
	eraseLeft   = Esc + "1K"
	eraseLine   = Esc + "2K"
	eraseDown   = Esc + "J"
	eraseUp     = Esc + "1J"
	eraseScreen = Esc + "2J"

	black     = Esc + "30m"
	red       = Esc + "31m"
	green     = Esc + "32m"
	yellow    = Esc + "33m"
	blue      = Esc + "34m"
	magenta   = Esc + "35m"
	cyan      = Esc + "36m"
	white     = Esc + "37m"
	blackHi   = Esc + "30;1m"
	redHi     = Esc + "31;1m"
	greenHi   = Esc + "32;1m"
	yellowHi  = Esc + "33;1m"
	blueHi    = Esc + "34;1m"
	magentaHi = Esc + "35;1m"
	cyanHi    = Esc + "36;1m"
	whiteHi   = Esc + "37;1m"

	bgBlack     = Esc + "40m"
	bgRed       = Esc + "41m"
	bgGreen     = Esc + "42m"
	bgYellow    = Esc + "43m"
	bgBlue      = Esc + "44m"
	bgMagenta   = Esc + "45m"
	bgCyan      = Esc + "46m"
	bgWhite     = Esc + "47m"
	bgBlackHi   = Esc + "40;1m"
	bgredHi     = Esc + "41;1m"
	bgGreenHi   = Esc + "42;1m"
	bgYellowHi  = Esc + "43;1m"
	bgblueHi    = Esc + "44;1m"
	bgMagentaHi = Esc + "45;1m"
	bgCyanHi    = Esc + "46;1m"
	bgWhiteHi   = Esc + "47;1m"

	reset = Esc + "0m"
)

func getInfo(game string) string {

	inf, err := ini.Load("info.ini")
	if err != nil {
		fmt.Printf("Fail to read CONFIG file: %v", err)
		os.Exit(1)
	}

	desc := inf.Section(game).Key("desc").String()
	return desc

}

// Get info from the Drop File, h, w
func initDrop(path string) {
	alias, timeLeft, emulation, nodeNum, userNum := dropFileData(path)
	h, w := GetTermSize()
	u := User{alias, timeLeft, emulation, nodeNum, userNum, h, w}
	U = &u

}

func initIni() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read CONFIG file: %v", err)
		os.Exit(1)
	}

	menu_title := cfg.Section("general").Key("title").String()
	version := cfg.Section("general").Key("version").String()
	gm_host := cfg.Section("goldmine").Key("host").String()
	gm_port := cfg.Section("goldmine").Key("port").String()
	gm_tag := cfg.Section("goldmine").Key("tag").String()
	gm_enabled := cfg.Section("goldmine").Key("enabled").String()
	gm_script := cfg.Section("goldmine").Key("script").String()
	dp_script := cfg.Section("doorparty").Key("script").String()
	dp_enabled := cfg.Section("doorparty").Key("enabled").String()
	bl_script := cfg.Section("bbslink").Key("script").String()
	bl_enabled := cfg.Section("bbslink").Key("enabled").String()

	c := DoorConfig{menu_title, version, gm_host, gm_port, gm_tag, gm_enabled, gm_script, dp_script, dp_enabled, bl_script, bl_enabled}
	C = &c

}

func header(w int) {
	if w == 80 {
		printAnsiLoc("art/"+currCode+".ans", 0, 1)
		moveCursor(70, 2)
	}
	if w > 80 {
		fmt.Fprintf(os.Stdout, " ")
	}
}

func prompt(color string) {
	if U.W == 80 {
		printStringLoc(U.Alias+" - "+fmt.Sprint(U.TimeLeft)+" mins left"+reset, 3, 23, blackHi, bgBlack)
		moveCursor(3, 24)
		if color == "blue" {
			PrintAnsi("art/prompt-blue.ans", 0, 1)
			moveCursor(6, 24)
			fmt.Printf(bgBlue + "  " + reset)
			moveCursor(6, 24)
		}
		if color == "red" {
			PrintAnsi("art/prompt-red.ans", 0, 1)
			moveCursor(6, 24)
			fmt.Printf(bgRed + "  " + reset)
			moveCursor(6, 24)
		}
	}
	if U.W > 80 {
		fmt.Fprintf(os.Stdout, " ")
	}
}

// newTimer boots a user after being idle too long
func newTimer(seconds int, action func()) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func() {
		<-timer.C
		action()

	}()
	// log.Println("time started...")
	return timer
}

func truncateText(s string, max int) string {
	if len(s) > max {
		r := 0
		for i := range s {
			r++
			if r > max-3 {
				return s[:i] + "..."
			}
		}
	}
	return s
}

// Move cursor to X, Y location
func moveCursor(x int, y int) {
	fmt.Printf(Esc+"%d;%df", y, x)
}

func eraseToLeft() {
	fmt.Println(eraseLeft)
}

// Erase the screen
func clearScreen() {
	fmt.Println(eraseScreen)
	moveCursor(0, 0)
}

// Move the cursor n cells to up.
func cursorUp(n int) {
	fmt.Printf(Esc+"%dA", n)
}

// Move the cursor n cells to down.
func cursorDown(n int) {
	fmt.Printf(Esc+"%dB", n)
}

// Move the cursor n cells to right.
func cursorForward(n int) {
	fmt.Printf(Esc+"%dC", n)
}

// Move the cursor n cells to left.
func cursorBack(n int) {
	fmt.Printf(Esc+"%dD", n)
}

// Move cursor to beginning of the line n lines down.
func cursorNextLine(n int) {
	fmt.Printf(Esc+"%dE", n)
}

// Move cursor to beginning of the line n lines up.
func cursorPreviousLine(n int) {
	fmt.Printf(Esc+"%dF", n)
}

// Move cursor horizontally to x.
func cursorHorizontalAbsolute(x int) {
	fmt.Printf(Esc+"%dG", x)
}

// Show the cursor.
func cursorShow() {
	fmt.Print(Esc + "?25h")
}

// Hide the cursor.
func cursorHide() {
	fmt.Print(Esc + "?25l")
}

// Save the screen.
func saveScreen() {
	fmt.Print(Esc + "?47h")
}

// Restore the saved screen.
func restoreScreen() {
	fmt.Print(Esc + "?47l")
}

func dropFileData(path string) (string, int, int, int, int) {
	// path needs to include trailing slash!
	var dropAlias string
	var dropTimeLeft string
	var dropEmulation string
	var nodeNum string
	var userNum string

	file, err := os.Open(strings.ToLower(path + "door32.sys"))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	count := 0
	for _, line := range text {
		if count == 4 {
			userNum = line
		}
		if count == 6 {
			dropAlias = line
		}
		if count == 8 {
			dropTimeLeft = line
		}
		if count == 9 {
			dropEmulation = line
		}
		if count == 10 {
			nodeNum = line
		}
		if count == 11 {
			break
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	timeInt, err := strconv.Atoi(dropTimeLeft) // return as int
	if err != nil {
		log.Fatal(err)
	}

	emuInt, err := strconv.Atoi(dropEmulation) // return as int
	if err != nil {
		log.Fatal(err)
	}
	nodeInt, err := strconv.Atoi(nodeNum) // return as int
	if err != nil {
		log.Fatal(err)
	}

	userNumInt, err := strconv.Atoi(userNum) // return as int
	if err != nil {
		log.Fatal(err)
	}

	return dropAlias, timeInt, emuInt, nodeInt, userNumInt
}

/*
Get the terminal size
- Send a cursor position that we know is way too large
- Terminal sends back the largest row + col size
- Read in the result
*/
func GetTermSize() (int, int) {
	// Set the terminal to raw mode so we aren't waiting for CLRF rom user (to be undone with `-raw`)
	rawMode := exec.Command("/bin/stty", "raw")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()

	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "\033[999;999f") // larger than any known term size
	fmt.Fprintf(os.Stdout, "\033[6n")       // ansi escape code for reporting cursor location
	text, _ := reader.ReadString('R')

	// Set the terminal back from raw mode to 'cooked'
	rawModeOff := exec.Command("/bin/stty", "-raw")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()

	// check for the desired output
	if strings.Contains(string(text), ";") {
		re := regexp.MustCompile(`\d+;\d+`)
		line := re.FindString(string(text))

		s := strings.Split(line, ";")
		sh, sw := s[0], s[1]

		ih, err := strconv.Atoi(sh)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		iw, err := strconv.Atoi(sw)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		h := ih
		w := iw

		clearScreen()

		return h, w

	} else {
		// couldn't detect, so let's just set 80 x 25 to be safe
		h := 80
		w := 25

		return h, w
	}
}

func PrintAnsi(artfile string, delay int, height int) {
	b, err := os.ReadFile(artfile) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	noSauce := trimStringFromSauce(string(b)) // strip off the SAUCE metadata
	s := bufio.NewScanner(strings.NewReader(string(noSauce)))

	i := 1

	for s.Scan() {
		fmt.Print(s.Text())
		time.Sleep(time.Duration(delay) * time.Millisecond)
		if i < height {
			fmt.Fprintf(os.Stdout, "\r\n")
		} else {
			moveCursor(0, 0)
			break
		}
		i++
	}
}

func trimStringFromSauce(s string) string {
	if idx := strings.Index(s, "COMNT"); idx != -1 {
		string := s
		delimiter := "COMNT"
		leftOfDelimiter := strings.Split(string, delimiter)[0]
		trim := trimLastChar(leftOfDelimiter)
		return trim
	}
	if idx := strings.Index(s, "SAUCE00"); idx != -1 {
		string := s
		delimiter := "SAUCE00"
		leftOfDelimiter := strings.Split(string, delimiter)[0]
		trim := trimLastChar(leftOfDelimiter)
		return trim
	}
	return s
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func printAnsiLoc(artfile string, x int, y int) {
	b, err := os.ReadFile(artfile) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	noSauce := trimStringFromSauce(string(b)) // strip off the SAUCE metadata
	s := bufio.NewScanner(strings.NewReader(string(noSauce)))
	yLoc := y

	for s.Scan() {
		fmt.Printf(Esc + strconv.Itoa(yLoc) + ";" + strconv.Itoa(x) + "f" + s.Text())
		yLoc++
	}
}

// Print text at an X, Y location
func printStringLoc(text string, x int, y int, fg string, bg string) {
	fmt.Fprintf(os.Stdout, reset+Esc+strconv.Itoa(y)+";"+strconv.Itoa(x)+"f"+bg+fg+text+reset)

}

// Horizontally center some text.
func centerText(s string, w int) {
	fmt.Fprintf(os.Stdout, (fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))))
}

func absCenterArt(artfile string, l int) {
	artY := (modalH / 2) - 2
	artLen := l / 2
	artX := (modalW - modalW/2) - artLen

	noSauce := trimStringFromSauce(artfile) // strip off the SAUCE metadata
	s := bufio.NewScanner(strings.NewReader(string(noSauce)))

	for s.Scan() {
		fmt.Fprintf(os.Stdout, Esc+strconv.Itoa(artY)+";"+strconv.Itoa(artX)+"f")
		fmt.Println(s.Text())
		artY++
	}
}

func justifyText(text string, maxWidth int) string {
	// Split the text into words
	words := strings.Split(text, " ")
	// Keep track of current line and previous lines
	var lines, current string
	// Iterate through each word
	for _, word := range words {
		// If the word fits, add to current
		if len(current)+len(word) <= maxWidth {
			// All words need a space after, expect for the last word
			// we will trim the space for the last word
			current += word + " "
			continue
		}
		// Otherwise, word did not fit so lets finish this line
		// Trim extra space from the last word
		current = current[:len(current)-1]
		// Distribute any extra spaces and add current line to previous lines
		lines += distributeSpaces(current, maxWidth-len(current))
		// The next iteration current line will be this word that did not fit
		current = word + " "
	}
	// return the previous lines and any current line we were working on
	return lines + current
}

// distributeSpaces will spread extra amount of whitespaces into the gaps
// between the words in s
func distributeSpaces(s string, extra int) string {
	// Nothing to do, return early
	if extra < 1 {
		return s + "\n"
	}
	// Split into tokens as it's easier to work with
	words := strings.Split(s, " ")
	// Edge case where only 1 word is on a line
	// we can just put all extra space at the end then return
	if len(words) == 1 {
		for extra > 0 {
			words[0] += " "
			extra--
		}
		return words[0] + "\n"
	}
	// calculate how many spaces between each word and any extra that dont
	// divide evenly into the gaps
	gaps := len(words) - 1
	spaces := extra + gaps
	spacePerGap, extraSpace := spaces/gaps, spaces%gaps
	// Distribute the spaces that dont fit evenly
	for i := 0; extraSpace > 0; i++ {
		words[i] += " "
		extraSpace--
	}
	// Make a tmp spacer for the whitespace that does divide evenly
	tmp := ""
	for spacePerGap > 0 {
		tmp += " "
		spacePerGap--
	}
	// Join the words with the tmp spacer
	return strings.Join(words, tmp) + "\n"
}

func printMultiStringAt(text string, x int, y int) {

	xLoc := x
	yLoc := y
	for _, line := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		moveCursor(xLoc, yLoc)
		fmt.Println(line)
		yLoc++
	}

}