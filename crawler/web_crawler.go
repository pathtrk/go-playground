package main

import (
	"fmt"
	"sync"
	"time"
)

// Fetcher inplements Fetch method
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type webCrawler struct {
	urls    map[string]int
	results []string
	mux     sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (w *webCrawler) Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		w.results = append(w.results, fmt.Sprintln(err))
		return
	}
	w.results = append(w.results, fmt.Sprintf("found: %s %q\n", url, body))
	for _, u := range urls {
		w.mux.Lock()
		w.urls[u]++
		w.mux.Unlock()
		if w.urls[u] < 2 {
			w.Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func (w *webCrawler) Results() []string {
	w.mux.Lock()
	defer w.mux.Unlock()
	return w.results
}

func main() {
	w := webCrawler{urls: make(map[string]int)}
	go w.Crawl("http://golang.org/", 4, fetcher)
	time.Sleep(time.Second)
	for _, s := range w.Results() {
		fmt.Println(s)
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
