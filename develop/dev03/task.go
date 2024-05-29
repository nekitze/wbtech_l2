package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Flags struct {
	Reversed     bool
	NumberSorted bool
	Deduplicated bool
	SortColumn   int
}

func sort(lines []string, flags Flags) []string {
	var sorted []string

	if flags.Deduplicated {
		seen := make(map[string]bool)
		for _, line := range lines {
			if _, ok := seen[line]; !ok {
				seen[line] = true
				sorted = append(sorted, line)
			}
		}
	} else {
		sorted = lines
	}

	slices.SortFunc(sorted, func(a, b string) int {
		return compareStrings(a, b, flags)
	})

	if flags.Reversed {
		slices.Reverse(sorted)
	}

	return sorted
}

func compareStrings(a, b string, flags Flags) int {
	s1 := a
	s2 := b

	if flags.SortColumn >= 0 {
		s1 = strings.Split(a, " ")[flags.SortColumn]
		s2 = strings.Split(b, " ")[flags.SortColumn]
	}

	if flags.NumberSorted {
		v1, _ := strconv.Atoi(s1)
		v2, _ := strconv.Atoi(s2)
		return cmp.Compare(v1, v2)
	}

	return cmp.Compare(s1, s2)
}

func readFile(file *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseFlags() Flags {
	reversed := flag.Bool("r", false, "reverse order")
	numberSorted := flag.Bool("n", false, "sort by number")
	deduplicated := flag.Bool("u", false, "remove duplicate lines")
	sortColumn := flag.Int("k", -1, "sort column")

	flag.Parse()

	return Flags{
		Reversed:     *reversed,
		NumberSorted: *numberSorted,
		Deduplicated: *deduplicated,
		SortColumn:   *sortColumn,
	}
}

// example : go run task.go -r task.go
func main() {
	filepath := os.Args[len(os.Args)-1]
	flags := parseFlags()

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := readFile(file)
	sorted := sort(lines, flags)
	for _, s := range sorted {
		fmt.Println(s)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
