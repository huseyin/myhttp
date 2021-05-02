package main

import (
	"context"
	"io"
	"reflect"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_result_output(t *testing.T) {
	type fields struct {
		url  string
		err  error
		body []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				url:  tt.fields.url,
				err:  tt.fields.err,
				body: tt.fields.body,
			}
			if got := r.output(); got != tt.want {
				t.Errorf("result.output() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_work_init(t *testing.T) {
	type fields struct {
		urls     []string
		c        int
		initOnce sync.Once
		results  chan *result
		w        io.Writer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &work{
				urls:     tt.fields.urls,
				c:        tt.fields.c,
				initOnce: tt.fields.initOnce,
				results:  tt.fields.results,
				w:        tt.fields.w,
			}
			b.init()
		})
	}
}

func Test_work_run(t *testing.T) {
	type fields struct {
		urls     []string
		c        int
		initOnce sync.Once
		results  chan *result
		w        io.Writer
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &work{
				urls:     tt.fields.urls,
				c:        tt.fields.c,
				initOnce: tt.fields.initOnce,
				results:  tt.fields.results,
				w:        tt.fields.w,
			}
			b.run(tt.args.ctx)
		})
	}
}

func Test_work_makeRequest(t *testing.T) {
	type fields struct {
		urls     []string
		c        int
		initOnce sync.Once
		results  chan *result
		w        io.Writer
	}
	type args struct {
		ctx    context.Context
		urlStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &work{
				urls:     tt.fields.urls,
				c:        tt.fields.c,
				initOnce: tt.fields.initOnce,
				results:  tt.fields.results,
				w:        tt.fields.w,
			}
			if gotRes := b.makeRequest(tt.args.ctx, tt.args.urlStr); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("work.makeRequest() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_work_writer(t *testing.T) {
	type fields struct {
		urls     []string
		c        int
		initOnce sync.Once
		results  chan *result
		w        io.Writer
	}
	tests := []struct {
		name   string
		fields fields
		want   io.Writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &work{
				urls:     tt.fields.urls,
				c:        tt.fields.c,
				initOnce: tt.fields.initOnce,
				results:  tt.fields.results,
				w:        tt.fields.w,
			}
			if got := in0.String(); got != tt.want {
				t.Errorf("work.writer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_work_print(t *testing.T) {
	type fields struct {
		urls     []string
		c        int
		initOnce sync.Once
		results  chan *result
		w        io.Writer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &work{
				urls:     tt.fields.urls,
				c:        tt.fields.c,
				initOnce: tt.fields.initOnce,
				results:  tt.fields.results,
				w:        tt.fields.w,
			}
			b.print()
		})
	}
}

func TestMD5(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.b); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errAndExit(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errAndExit(tt.args.msg)
		})
	}
}
