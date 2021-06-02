package main

import (
	"golang.org/x/sync/errgroup"
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"time"
)

type hands struct {
}

func (hands) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(bytes.NewBufferString("hello").Bytes())
	w.Write(bytes.NewBufferString(r.RemoteAddr).Bytes())

	time.Sleep(time.Second * 2)
}

func main() {
	var eg errgroup.Group
	ports := []string{":9800", ":9801", ":9802"}
	handl := hands{}
	http.Handle("/", handl)
	for _, port := range ports {
		port := port
		go func(port string) {
			http.ListenAndServe(port, nil)
		}(port)
	}
	urls := []string{
		"http://localhost:9800",
		"http://localhost:9801",
		"http://localhost:9802",
	}
	for _, url := range urls {
		eg.Go(func() error {
			url := url
			client := http.Client{}
			res, err := client.Get(url)
			if err != nil {
				return err
			}
			resp, _ := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", resp)
			return nil
		})
	}
	if err := eg.Wait(); err == nil {
		fmt.Printf("success")
	}


}
