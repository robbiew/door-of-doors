package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
	gd "github.com/robbiew/door-of-doors/common"
	"github.com/robbiew/door-of-doors/timeout"
)

var (
	dropPath string
)

// launches external program (goldmine-connect)
func goGoldmine(un string, host string, port string, tag string) {
	prg := "./goldmine-connect"
	arg1 := host
	arg2 := port
	arg3 := un
	arg4 := tag
	arg5 := "lord"

	gd.ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2+" "+arg3+" "+arg4+" "+arg5)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking

}

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

}

func main() {

	// Get door32.sys, h, w as user object
	u := gd.Initialize(dropPath)
	c := gd.ConfGoldMine("./config.ini")

	gd.ClearScreen()
	gd.MoveCursor(0, 0)

	// Exit if no ANSI capabilities (sorry!)
	if u.Emulation != 1 {
		fmt.Println("Sorry, ANSI is required to use this...")
		time.Sleep(time.Duration(2) * time.Second)
		os.Exit(0)
	}

	timeout.StartTimer(u.TimeLeft)

	// A reliable keyboard library to detect key presses
	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		timeout.StartTimer(timeout.UpdatedTime)

		gd.MoveCursor(0, 0)
		gd.HeaderBar(u.W)

		gd.MoveCursor(0, 0)
		if timeout.Warning {

			fmt.Printf(gd.BgCyan+gd.RedHi+"Almost out of time! %v mins"+gd.Reset, u.TimeLeft)
		} else {

			fmt.Printf(gd.BgCyan+gd.CyanHi+"Time Left: %v mins"+gd.Reset, u.TimeLeft)
		}

		// A Test Menu
		gd.MoveCursor(0, 3)

		fmt.Fprintf(os.Stdout, gd.Cyan+"["+gd.YellowHi+"Q"+gd.Cyan+"] "+gd.Reset+gd.Magenta+"Quit\r\n")

		fmt.Fprintf(os.Stdout, gd.Reset+"\r\nCommand? ")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if string(char) == "t" || string(char) == "T" {
			goGoldmine(u.Alias, c.Host, c.Port, c.Tag)
		}
		if string(char) == "q" || string(char) == "Q" || key == keyboard.KeyEsc {
			break
		}
		gd.ClearScreen()
		continue
	}
}
