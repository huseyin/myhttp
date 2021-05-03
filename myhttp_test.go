package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	var url string

	body := "<html><body>Hello, world!</body></html>"

	handler := func(w http.ResponseWriter, r *http.Request) {
		url = "http://" + r.Host

		fmt.Fprintf(w, body)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	w := &work{
		urls: []string{server.URL},
		c:    1,
		w:    io.Discard,
	}

	w.init()
	w.doPool()
	go w.finish()

	for res := range w.results {
		if url != res.url {
			t.Errorf("%v received, %v expected.", url, res.url)
		}

		if res.err != nil {
			t.Errorf("%v received, nil expected.", res.err)
		}

		if body != string(res.body) {
			t.Errorf("%v received, %v expected.", body, string(res.body))
		}
	}
}

func TestOutput(t *testing.T) {
	tests := []struct {
		name string
		src  *result
		want string
	}{
		{
			name: "test case without err",
			src: &result{
				url:  "https://example.com",
				err:  nil,
				body: []byte("test body"),
			},
			want: "https://example.com bbf9afe7431caf5f89a608bc31e8d822",
		},
		{
			name: "test case with err",
			src: &result{
				url:  "https://example.com",
				err:  errors.New("test error"),
				body: []byte("test body"),
			},
			want: "https://example.com test error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.src.output(); got != tt.want {
				t.Errorf("%v received, %v expected.", got, tt.want)
			}
		})
	}
}

func TestMD5(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "test case",
			data: []byte("test string"),
			want: "6f8db599de986fab7a21625b7916589c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.data); got != tt.want {
				t.Errorf("%v received, %v expected.", got, tt.want)
			}
		})
	}
}
