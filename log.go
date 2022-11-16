package main

import (
	"bufio"
	"fmt"
	"io"
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

func ReadExactLine(fileName string, lineNumber int) string {

	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occurred! ", err)
	}

	br := bufio.NewReader(inputFile)
	for i := 1; i < lineNumber; i++ {
		_, _ = br.ReadString('\n')
	}
	str, _ := br.ReadString('\n')
	fmt.Println("Line is ", str)
	return str
}

func lastUsers() {

	file := "activity.log"
	// check if there are any log entries, if not -- skip
	fi, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
	}
	if fi.Size() == 0 {

	} else {
		file := "activity.log"

		b, err := os.ReadFile(file)
		if err != nil {
			fmt.Print(err)
		}

		moveCursor(2, 21)
		PrintAnsi("art/bullet.ans", 0, 1)

		ct := lineCounter(string(b), '\n')

		y := 21
		i := 0
		for i < 10 {
			fmt.Print(ct)
			moveCursor(4, y)
			x := ReadExactLine(file, ct)
			fmt.Print(x)
			ct--
			i++
			y++
		}

		// date, error := time.Parse("2006/01/02 15:04:05 ", last[0])
		// if error != nil {
		// 	fmt.Println(error)
		// 	return
		// }

		// s := timeago.English.Format(date)
		// moveCursor(4, 21)
		// fmt.Printf(cyanHi+"%v played %v on %v"+blackHi+" %v", last[1], last[2], last[3], s)
		// moveCursor(4, 22)

	}

	// res1 := strings.SplitN(string(last), "|", 4)

}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.

// func Readln(r *bufio.Reader) ([]string, error) {

// 	var (
// 		isPrefix bool  = true
// 		err      error = nil
// 		line     []byte
// 		last     []string
// 	)
// 	for isPrefix && err == nil {
// 		line, isPrefix, err = r.ReadLine()
// 		last = append(last, string(line))
// 	}
// 	return last, err
// }

func ReadLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

func lineCounter(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
