package main

import "flag"

func main()  {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

}

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter out any links w/ a different domain
   5. Find all pages (BFS)
   6. print out XML
*/
