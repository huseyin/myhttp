package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
)

const defaultTimeout = 10 * time.Second

func main() {
	flag.Parse()
	parallel := flag.Int("parallel", 10, "The number of parallel requests.")

	urls := flag.Args()
	if len(urls) == 0 {
		errAndExit("You must specify at least one URL.")
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Listen to the signals to shutdown gracefully.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()

	w := &work{c: *parallel, urls: urls}
	w.run(ctx)
}

type result struct {
	url  string
	err  error
	body []byte
}

// output returns a string that is formatted with URL and MD5-encoded
// body. If there is an error, it will return the error message
// instead of body.
func (r *result) output() string {
	if r.err != nil {
		return fmt.Sprintf("%s %v", r.url, r.err)
	}
	return fmt.Sprintf("%s %s", r.url, MD5(r.body))
}

type work struct {
	// urls is a collection of the URLs that will be used.
	urls []string

	// c is the concurrency level
	c int

	initOnce sync.Once
	results  chan *result

	// w is where results will be written
	w io.Writer
}

// init initializes the internal data structures.
func (b *work) init() {
	b.initOnce.Do(func() {
		b.results = make(chan *result)
		b.w = b.writer()
	})
}

// run makes all the requests, prints the result. It blocks until
// all work is done.
func (b *work) run(ctx context.Context) {
	b.init()

	var wg sync.WaitGroup
	wg.Add(len(b.urls))

	urls := make(chan string, len(b.urls))
	for _, url := range b.urls {
		urls <- url
	}
	close(urls)

	for i := 0; i < b.c; i++ {
		go func() {
			for url := range urls {
				b.results <- b.makeRequest(ctx, url)
				wg.Done()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(b.results)
	}()

	b.print()
}

// makeRequest makes a GET request with the provided context, returns a result
// which is populated with the HTTP response.
func (b *work) makeRequest(ctx context.Context, urlStr string) (res *result) {
	res = &result{url: urlStr}

	u, err := url.Parse(urlStr)
	if err != nil {
		res.err = err
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		res.err = err
		return
	}

	client := &http.Client{Timeout: defaultTimeout}
	resp, err := client.Do(req)
	if err != nil {
		res.err = err
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.err = err
		return
	}
	res.body = body

	return
}

// writer returns the io.Writer which is used to write result.
// If nil, it defaults to stdout.
func (b *work) writer() io.Writer {
	if b.w == nil {
		return os.Stdout
	}
	return b.w
}

// print prints the results to where output points.
func (b *work) print() {
	for res := range b.results {
		fmt.Fprintln(b.writer(), res.output())
	}
}

// MD5 encodes a given byte array into a MD5 hash,
// returns it as hex string.
func MD5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func errAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
