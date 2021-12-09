package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sitemap-builder/link"
	"strings"
)

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter out any links w/ a different domain
   5. Find all pages (BFS)
   6. print out XML
*/
/*
	Learnings:
		io.Copy.copy data from reader to a writer
		os.Stdout is the writer that writes to the terminal
*/
func main() {
	urlFlag := flag.String("url", "https://go.dev", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func bfs(urlString string, maxDepth int) []string {
	// a map of pages we've already seen
	// it's a map because we don't want to add duplicates, and we want it's fast lookup
	// it's better to use empty struct (struct{}) than map[string]bool if you don't care about the value
	// the empty struct don't use memory

	seen := make(map[string]struct{})
	// a queue is the queue of the pages/links we are visiting now
	// nextQueue is the queue of the pages/links we are visiting next
	// after finishing links in the queue, saving the links to the nextQueue
	// we add the nextQueue to the queue and repeat the process
	var queue map[string]struct{}
	nextQueue := map[string]struct{}{
		urlString: struct{}{},
	}
	for i := 0; i < maxDepth; i++ {
		// we add the nextQueue to the queue
		// we make the nextQueue empty
		queue, nextQueue = nextQueue, make(map[string]struct{})
		for currentPageUrl, _ := range queue {
			// currentPageUrl is the current level url
			if _, ok := seen[currentPageUrl]; !ok {
				// struct{}{} consists of 2 parts
				// first part struct{} which is the type itself
				// and the second part is {} which means we are instantiating an empty struct
				seen[currentPageUrl] = struct{}{}
			}
			// l is the links found in current level page,
			// so, we can add them to the nextQueue
			for _, l := range getUrls(urlString) {
				nextQueue[l] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for l := range seen {
		ret = append(ret, l)
	}
	return ret
}
func getUrls(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host: reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	// name it links instead of hrefs because it's same as the name of the function
	var hSlice []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hSlice = append(hSlice, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hSlice = append(hSlice, l.Href)
		}
	}
	return hSlice
}

func filter(links []string, keepFn func(string) bool) []string {
	var fSlice []string
	for _, l := range links {
		if keepFn(l) {
			fSlice = append(fSlice, l)
		}
	}
	return fSlice
}

func withPrefix(prefix string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, prefix)
	}
}