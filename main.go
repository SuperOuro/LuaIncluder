package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var IncludeRegex *regexp.Regexp = regexp.MustCompile(`--include\s["|'](?P<Path>.*)["|']`)
var PathIndex int = IncludeRegex.SubexpIndex("Path")

func timer() func() {
	start := time.Now();
	return func() {
		fmt.Printf("took %v\n", time.Since(start));
	}
}

func FindIncludeAndReplace(Lines *[]string, Line string, i int, ch chan int) {
	matches := IncludeRegex.FindStringSubmatch(Line)

	text, err := os.ReadFile(matches[PathIndex])
	if err != nil {
		log.Fatal(err)
	}

	ch <- 1
	(*Lines)[i] = string(text)
}

func StartInclude(MainFile []byte) string {
	Lines, GoroutinesNumber := strings.Split(string(MainFile), "\n"), 0
	ch := make(chan int)

	for i, Line := range Lines {
		if IncludeRegex.MatchString(Line) {
			go FindIncludeAndReplace(&Lines, Line, i, ch)
			GoroutinesNumber++
		}
	}

	for i := 1; i <= GoroutinesNumber; i++ {
		<-ch
	}

	return strings.Join(Lines, "\n");
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected main file path, got nil.")
	}

	defer timer()();
	MainFilePath := os.Args[1]

	MainFile, err := os.ReadFile(MainFilePath)
	if err != nil {
		log.Fatal(err)
	}

	output := StartInclude(MainFile);
	for IncludeRegex.MatchString(output) {
		output = StartInclude([]byte(output));
	}

	err = os.WriteFile("output.lua", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
