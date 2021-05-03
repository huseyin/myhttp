package main

import (
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

const (
	defaultTimeout = 10 * time.Second
)

func main() {
	var parallel = flag.Int("parallel", 10, "The number of parallel requests.")

	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		errAndExit("You must specify at least one URL.")
	}

	w := &work{c: *parallel, urls: urls}
	w.run()

	// Listen to the signals to shutdown gracefully.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		w.stop()
	}()
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
	urls []string

	c int // concurrency level

	pool     *sync.WaitGroup
	initOnce sync.Once
	stopCh   chan struct{}
	results  chan *result

	w io.Writer
}

// init initializes the internal data structures.
func (b *work) init() {
	b.initOnce.Do(func() {
		b.results = make(chan *result)
		b.stopCh = make(chan struct{}, b.c)
		b.w = b.writer()
		b.pool = &sync.WaitGroup{}
		b.pool.Add(len(b.urls))
	})
}

// run makes all the requests, prints the result. It blocks
// until all work is done.
func (b *work) run() {
	b.init()
	b.doPool()
	go func() {
		b.finish()
	}()
	b.print()
}

// stop gracefully shuts down the workers when a stop
// signal received.
func (b *work) stop() {
	for i := 0; i < b.c; i++ {
		b.stopCh <- struct{}{}
	}
}

// finish closes the chan of results and waits for the workers.
func (b *work) finish() {
	b.pool.Wait()
	close(b.results)
}

func (b *work) doPool() {
	urls := make(chan string, len(b.urls))
	for _, url := range b.urls {
		urls <- url
	}
	close(urls)

	for i := 0; i < b.c; i++ {
		go func() {
			for url := range urls {
				select {
				case <-b.stopCh:
					return
				default:
					b.results <- b.do(url)
					b.pool.Done()
				}
			}
		}()
	}
}

func (b *work) do(urlStr string) (res *result) {
	res = &result{url: urlStr}

	u, err := url.Parse(urlStr)
	if err != nil {
		res.err = err
		return
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
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
