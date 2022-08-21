package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	var urls []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		u := scanner.Text()

		if u != "" {
			urls = append(urls, u)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for _, url := range urls {
		host := host(url)

		if host == "" {
			fmt.Fprintln(os.Stderr, "Cannot parse as URL: "+url)
			continue
		}

		fmt.Println(host)
	}
}

func host(u string) string {
	parsed, err := url.Parse(u)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	return parsed.Host
}
