package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Flags struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	Numeration bool
	Regexp     string
}

func printData(data []string, flags Flags) {
	var count int
	for _, line := range data {
		if flags.Count {
			count++
			continue
		}

		fmt.Println(line)
	}

	if flags.Count {
		fmt.Println(count)
	}
}

func filterLines(data []string, flags Flags) ([]string, error) {
	var result []string

	regx, err := regexp.Compile(flags.Regexp)
	if err != nil {
		return nil, err
	}

	if flags.Numeration {
		for i, line := range data {
			data[i] = strconv.Itoa(i) + "\t" + line
		}
	}

	for i, line := range data {
		if (regx.MatchString(line) && !flags.Invert && !flags.Fixed) ||
			(!regx.MatchString(line) && flags.Invert && !flags.Fixed) ||
			(flags.Fixed && line == flags.Regexp && !flags.Invert) ||
			(flags.Fixed && line != flags.Regexp && flags.Invert) {
			result = append(result, line)

			if flags.Before > 0 || flags.Context > 0 {
				for j := i - 1; j >= i-flags.Before && j >= 0; j-- {
					result = slices.Insert(result, len(result)-1, data[j])
				}
			}

			if flags.After > 0 || flags.Context > 0 {
				for j := i + 1; j <= i+flags.After && j < len(data); j++ {
					result = append(result, data[j])
				}
			}
		}
	}

	return result, nil
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
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой")
	numeration := flag.Bool("n", false, "напечатать номер строки")
	regx := flag.String("e", "", "регулярное выражение")

	flag.Parse()
	return Flags{
		After:      *after,
		Before:     *before,
		Context:    *context,
		Count:      *count,
		IgnoreCase: *ignoreCase,
		Invert:     *invert,
		Fixed:      *fixed,
		Numeration: *numeration,
		Regexp:     *regx,
	}
}

func main() {
	filepath := os.Args[len(os.Args)-1]
	flags := parseFlags()

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	data := readFile(file)
	_ = file.Close()

	result, err := filterLines(data, flags)
	if err != nil {
		panic(err)
	}

	printData(result, flags)
}
