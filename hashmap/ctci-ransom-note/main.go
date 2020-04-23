package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) string {
	magazineKeyword := map[string]int{}
	noteKeyword := map[string]int{}
	for _, kw := range magazine {
		if _, existed := magazineKeyword[kw]; !existed {
			magazineKeyword[kw] = 1
		} else {
			magazineKeyword[kw] += 1
		}
	}
	for _, text := range note {
		magCount, existed := magazineKeyword[text]
		if !existed {
			fmt.Println("not existed")
			return "No"
		}
		noteKeyword[text] += 1
		if noteKeyword[text] > magCount {
			fmt.Println("keyword > magCount", text, noteKeyword[text], magCount)
			return "No"
		}
	}
	log.Println("noteMap", noteKeyword)
	return "Yes"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
