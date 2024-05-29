package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func cutLine(line, delimiter string, fields int, separated bool) string {
	columns := strings.Split(line, delimiter)

	if separated && !strings.Contains(line, delimiter) {
		return ""
	} else if len(columns) < fields {
		return line
	} else {
		return columns[fields-1]
	}
}

func main() {
	fields := flag.Int("f", 0, "выбрать колонки")
	delimiter := flag.String("d", " ", "указать разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	if *fields <= 0 {
		panic("-f should be > 0")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = cutLine(line, *delimiter, *fields, *separated)
		fmt.Println(line)
	}
}
