package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"net/http"
	"testing"
)

//	func TestRun(t *testing.T) {
//		ctx, cancel := context.WithCancel(context.Background())
//		eg, ctx := errgroup.WithContext(ctx)
//		eg.Go(func() error {
//			return run(ctx)
//		})
//		in := "message"
//		rsp, err := http.Get("http://localhost:18080/" + in)
//		if err != nil {
//			t.Errorf("failed to get: %+v", err)
//		}
//		defer rsp.Body.Close()
//		got, err := io.ReadAll(rsp.Body)
//		if err != nil {
//			t.Fatalf("failed to read body: %v", err)
//		}
//		want := fmt.Sprintf("Hello, %s!", in)
//		if string(got) != want {
//			t.Errorf("want %q, but got %q", want, got)
//		}
//		cancel()
//		if err := eg.Wait(); err != nil {
//			t.Fatal(err)
//		}
//	}
func TestRun(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, listener)
	})
	in := "message"
	url := fmt.Sprintf("http://%s/%s", listener.Addr().String(), in)
	t.Logf("try request to %q", url)
	response, err := http.Get(url)

	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer response.Body.Close()
	got, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
