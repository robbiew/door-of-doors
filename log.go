package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func writeLog(f *os.File, user string, door string, server string) {
	log.SetOutput(f)
	log.Println("|" + user + "|" + door + "|" + server)
}

func getLogCount(file string, title string, server string) int {

	var err error

	fh, err := os.Open(file)
	if err != nil {
		log.Fatalf("Could not open file '%v': %v", file, err)
		os.Exit(1)
	}
	reader := bufio.NewReader(fh)
	count := 0
	for {
		line, _ := reader.ReadString('\n')
		// fmt.Print(line)
		if strings.Count(line, title) == 1 && strings.Count(line, server) == 1 {
			count++
		}
		if line == "" {
			break
		}

	}
	return count
}

func lastUser() []string {
	file := "activity.log"

	var err error

	fh, err := os.Open(file)
	if err != nil {
		log.Fatalf("Could not open file '%v': %v", file, err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fh)
	var last []byte
	for scanner.Scan() {
		last = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res1 := strings.SplitN(string(last), "|", 4)

	return res1

}
