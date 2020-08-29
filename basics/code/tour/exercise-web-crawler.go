package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher) {
// 	// TODO: Fetch URLs in parallel.
// 	// TODO: Don't fetch the same URL twice.
// 	// This implementation doesn't do either:
// 	if depth <= 0 {
// 		return
// 	}
// 	body, urls, err := fetcher.Fetch(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("found: %s %q\n", url, body)
// 	for _, u := range urls {
// 		Crawl(u, depth-1, fetcher)
// 	}
// 	return
// }

//
// Concurrent crawler with shared state and mutex
//
type fetchState struct {
	mu      sync.Mutex      // mutex is in this struct only for convenience, it's not required
	fetched map[string]bool // Map of fetched urls
}

func MutexCrawl(url string, fetcher Fetcher, f *fetchState) {
	f.mu.Lock() // Lock needed so that 2 different threads can't fetch at the same time
	isFetched := f.fetched[url]
	f.fetched[ur] = true // change state after fetching current status as it will be processed if not fetched
	f.mu.Unlock()

	if isFetched {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go MutexCrawl(u, fetcher, f)
	}

}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
