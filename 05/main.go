package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	firstLine := readFirstLine(file)
	seedsJoined := strings.Split(firstLine, ": ")[1]
	seeds := strings.Split(seedsJoined, " ")
	
	fmt.Print(strings.Join(seeds, ","))
	
}

func readFirstLine(file *os.File) string {
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	} else if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ""
}
