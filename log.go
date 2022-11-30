package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func writeLog(f *os.File, user string, door string, server string) {
	log.SetOutput(f)
	log.Println("|" + user + "|" + door + "|" + server)
}

func dup_count(list []string) map[string]int {
	duplicate_frequency := make(map[string]int)
	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func getTopDoors() {
	filename := "activity.log"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	reader := bufio.NewReader(file)
	ct := lineCounter(string(b), '\n')
	i := 0

	var list []string
	for i < ct {
		line, _ := reader.ReadString('\n')
		res1 := strings.SplitN(string(line), "|", 4)
		logItem := res1[2] + " on " + res1[3]
		list = append(list, logItem)
		i++
	}

	dup_map := dup_count(list)

	keys := make([]string, 0, len(dup_map))
	for k := range dup_map {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return dup_map[keys[i]] > dup_map[keys[j]]
	})

	i = 0
	y := 8
	for _, k := range keys {
		// sort.Strings(keys)
		if i < 5 {
			splitr := strings.SplitN(string(k), " on ", 2)
			moveCursor(3, y)
			fmt.Printf(reset+cyan+"%v."+reset, i+1)
			moveCursor(6, y)
			fmt.Printf(cyanHi+"%v"+reset, splitr[0])
			moveCursor(41, y)
			fmt.Printf(greenHi+"%v"+reset, splitr[1])
			moveCursor(55, y)

			if len(fmt.Sprint(dup_map[k])) == 1 {
				fmt.Printf(yellowHi+"  %d"+reset+blackHi+" plays"+reset, dup_map[k])
			}
			if len(fmt.Sprint(dup_map[k])) > 1 && len(fmt.Sprint(dup_map[k])) <= 2 {
				fmt.Printf(yellowHi+" %d"+reset+blackHi+" plays"+reset, dup_map[k])
			}
			if len(fmt.Sprint(dup_map[k])) > 2 {
				fmt.Printf(yellowHi+"%d"+reset+blackHi+" plays"+reset, dup_map[k])
			}
			i++
			y++
		}
	}
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
	// fmt.Println("Line is ", str)
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
		ct := lineCounter(string(b), '\n')

		y := 16
		i := 0
		for i < 6 && i <= ct {
			moveCursor(3, y)
			PrintAnsi("art/bullet.ans", 0, 1)
			x := ReadExactLine(file, ct)

			res1 := strings.SplitN(string(x), "|", 4)

			moveCursor(5, y)
			fmt.Printf(reset+yellowHi+"%v "+reset+cyan+"played "+whiteHi+"%v "+reset+cyan+"on "+cyanHi+"%v", res1[1], res1[2], res1[3])
			ct--
			i++
			y++
		}

	}

}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.

func Readln(r *bufio.Reader) ([]string, error) {

	var (
		isPrefix bool  = true
		err      error = nil
		line     []byte
		last     []string
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		last = append(last, string(line))
	}
	return last, err
}

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

func isLogEmpty() bool {
	file := "activity.log"
	var empty bool
	// does file exist?
	if fsize, err := os.Stat(file); err == nil {
		// is it empty?
		if fsize.Size() == 0 {
			empty = true
		} else {
			empty = false
		}

	} else {
		empty = false
	}
	return empty

}

func showStats() {
	// skip if there's no records
	if !isLogEmpty() {
		errorChan := make(chan error)
		dataChan := make(chan []byte)

		PrintAnsi("art/stats.ans", 0, 24)
		getTopDoors()
		lastUsers()

		moveCursor(2, 23)
		fmt.Println(reset + green + "[" + greenHi + "Hit a Key" + reset + green + "]" + reset)
		moveCursor(2, 23)

		for {

			go readWrapper(dataChan, errorChan)

			r, _ := utf8.DecodeRune(<-dataChan)

			if r != '~' {
				break
			}

		}
	}

}
