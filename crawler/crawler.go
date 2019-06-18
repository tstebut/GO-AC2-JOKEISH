// file: crawl.go
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jackdanger/collectlinks"
)

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}
	retrieve(args[0])
}

func retrieve(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}
}
