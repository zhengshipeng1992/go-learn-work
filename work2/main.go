package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func HttpServer(srv *http.Server) error {
	http.HandleFunc("/", handleServer)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

func handleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world \n")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, gCtx := errgroup.WithContext(ctx)

	srv := &http.Server{Addr: ":8080"}

	group.Go(func() error {
		return HttpServer(srv)
	})

	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(gCtx)
	})

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			case <-chanel:
				cancel()
			}
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group err:", err)
	}
	fmt.Println("all group done!")
}
