package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func loadDataFromUrl(client *http.Client, url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	u := flag.String("u", "", "target url")
	o := flag.String("o", "output.html", "output filename")
	flag.Parse()

	if len(*u) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide target url! Example: `-u google.com`")
		return
	}

	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	data, err := loadDataFromUrl(client, *u)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = os.WriteFile(*o, data, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
