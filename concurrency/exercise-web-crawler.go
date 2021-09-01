package main

import (
    "fmt"
    "sync"
)

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, c chan bool) {
    if depth <= 0 {
        close(c)
        return
    }
    body, urls, err := fetcher.Fetch(url)
    visitedUrls.mu.Lock()
    visitedUrls.v[url] = true
    if err != nil {
        fmt.Println(err)
        visitedUrls.mu.Unlock()
        close(c)
        return
    }
    visitedUrls.mu.Unlock()
    fmt.Printf("found: %s %q\n", url, body)

    // Prepare an array of channels for each thread
    var channels []chan bool
    for _, u := range urls {
        // If an URL key doesn't exist or it hasn't been visited yet, create a channel and 
        // send off the coroutine
        visitedUrls.mu.Lock()
        if visitedMapRes, visitedMapOk := visitedUrls.v[u]; visitedMapOk {
            visitedUrls.mu.Unlock()
            if !visitedMapRes {
              c := make(chan bool)
              channels = append(channels, c)
              go Crawl(u, depth-1, fetcher, c)
            }
        } else {
            //fmt.Println(u, "does not exist")
            visitedUrls.mu.Unlock()
              c := make(chan bool)
              channels = append(channels, c)
            go Crawl(u, depth-1, fetcher, c)
            //fmt.Println("go Crawl finished")
        }
    }

    // Wait for all sent coroutines
    for _, craw_chan := range channels {
        if _, open := <-craw_chan; !open {
            continue
        }
    }
    close(c)
    return
}

func main() {
    c := make(chan bool)
    Crawl("https://golang.org/", 4, fetcher, c)
    if _, open := <-c; !open {
        return
    }
}

type SafeVisitorUrlCounter struct {
    mu sync.Mutex
    v map[string]bool
}

var visitedUrls = SafeVisitorUrlCounter{v: make(map[string]bool)}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    visited bool
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
                false,
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
                false,
                "Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
                false,
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
                false,
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
