package main

import (
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"context"
	"net"
	"io"
	"time"
	"bytes"
)

func init() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, []os.Signal{syscall.SIGINT, syscall.SIGQUIT}...)
	go func() {
		for sig := range sigchan {
			fmt.Printf("catch signal %s\n", sig)
		}
	}()
}
func doConn(ctx context.Context, con *net.Conn) error {
	defer (*con).Close()
	fmt.Printf("deal remote:%s\n", (*con).RemoteAddr())
	var str bytes.Buffer
	for {
		b1 := make([]byte, 1)
		_, err := (*con).Read(b1)
		if err != nil {
			if err == io.EOF {

			} else {
				return err
			}
		}
		if b1[0] == '\n' {
			fmt.Printf("%s\n", str.String())
			str = bytes.Buffer{}
		}
		str.Write(b1)
	}
	<-ctx.Done()
	return nil
}

func doPing(con *net.Conn) error {
	tick := time.NewTicker(time.Second * 5)
	for {
		select {
		case _ = <-tick.C:
			_, err := (*con).Write([]byte("ping"))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	// 建立server监听端口
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("listen failed err:%s\n", err)
		return
	}
	//fmt.Printf("prepared at :9999\n")
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept err:%s\n", err)
			break
		}
		root := context.Background()
		eg2, ctx := errgroup.WithContext(root)
		// doConn模拟 处理http请求
		eg2.Go(func() error {
			return doConn(ctx, &con)
		})
		// doPing模拟 处理http请求
		//eg2.Go(func() error {
		//	return doPing(&con)
		//})
		// 处理请求结果
		if err = eg2.Wait(); err != nil {
			fmt.Printf("recv err:%s\n", err)
			return
		}
		con.Close()
	}
}
