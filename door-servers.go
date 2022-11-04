package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Launches a bash script that uses rlogin to connect to Door Party local server
func doorParty(doorCode string, un string, script string) {
	prg := script
	arg1 := fmt.Sprint(un)
	arg2 := doorCode

	ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}

// Launches a bash script that uses rlogin to connect to Gold Mine remote server
func goldMine(userName string, tag string, doorCode string, host string, port string, script string) {
	prg := script
	arg1 := userName
	arg2 := tag
	arg3 := doorCode
	arg4 := host
	arg5 := port

	ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2+" "+arg3+" "+arg4+" "+arg5)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}

// Launches a bash script that uses rlogin to connect to BBS Link server
func bbsLink(doorCode string, un int, script string) {
	prg := script
	arg1 := doorCode
	arg2 := fmt.Sprint(un)

	ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}
