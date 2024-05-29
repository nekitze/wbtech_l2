package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func PrintCurrentTime(server string) error {
	time, err := ntp.Time(server)
	if err != nil {
		return err
	}

	fmt.Println(time)

	return nil
}

func main() {
	timeServer := "0.beevik-ntp.pool.ntp.org"

	err := PrintCurrentTime(timeServer)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
