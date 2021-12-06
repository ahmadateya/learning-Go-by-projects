package main

import (
	"flag"
	"fmt"
	"sitemap-builder/link"
	"net/http"
)

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter out any links w/ a different domain
   5. Find all pages (BFS)
   6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	/*
		io.Copy.copy data from reader to a writer
		os.Stdout is the writer that writes to the terminal
	*/
	links, _ := link.Parse(resp.Body)
	fmt.Println(links)
}
