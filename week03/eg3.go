package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func fn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/quit" {
		// special path to exit
		os.Exit(1)
	}
	fmt.Printf("%s\n", r.URL.Path)
	w.Write([]byte("OK"))
}

func httpstart() error {
	serv := http.NewServeMux()
	serv.HandleFunc("/", fn)

	return http.ListenAndServe(":4000", serv)
}

func reg() error {
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGCHLD, syscall.SIGABRT, syscall.SIGHUP}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, sigs...)
	for sig := range sigchan {
		fmt.Printf("recv signal: %s\n", sig)
	}
	return nil
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return httpstart()
	})
	eg.Go(func() error {
		<-ctx.Done()
		fmt.Printf("http server will shutdown\n")
		return errors.New("shutdown")
	})
	eg.Go(reg)
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
