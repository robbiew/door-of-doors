package main

import (
	"fmt"
	"os"
	"os/exec"

	gd "github.com/robbiew/door-of-doors/common"
	"gopkg.in/ini.v1"
)

type DoorConfigs struct {
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

func GetConfig() DoorConfigs {

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

	c := DoorConfigs{

		Menu_Title: menu_title,
		Version:    version,
		GM_Host:    gm_host,
		GM_Port:    gm_port,
		GM_Tag:     gm_tag,
		GM_Enabled: gm_enabled,
		GM_script:  gm_script,

		DP_Script:  dp_script,
		DP_Enabled: dp_enabled,

		BL_Script:  bl_script,
		BL_Enabled: bl_enabled,
	}

	return c
}

// Launches a bash script that uses rlogin to connect to Door Party local server
func DoorParty(door string, un string, script string) {
	prg := script
	arg1 := fmt.Sprint(un)
	arg2 := door

	gd.ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}

// Launches a bash script that uses rlogin to connect to Gold Mine remote server
func GoldMine(userName string, tag string, doorCode string, host string, port string, script string) {
	prg := script
	arg1 := userName
	arg2 := tag
	arg3 := doorCode
	arg4 := host
	arg5 := port

	gd.ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2+" "+arg3+" "+arg4+" "+arg5)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}

// Launches a bash script that uses rlogin to connect to BBS Link server
func BbsLink(door string, un int, script string) {
	prg := script
	arg1 := door
	arg2 := fmt.Sprint(un)

	gd.ClearScreen()

	cmd := exec.Command("bash", "-c", prg+" "+arg1+" "+arg2)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking
}
