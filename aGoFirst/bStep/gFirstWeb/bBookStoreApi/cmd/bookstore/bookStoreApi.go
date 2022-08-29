package main

import (
	"context"
	_ "github.com/dog/bBookStoreApi/internal/store"
	"github.com/dog/bBookStoreApi/server"
	"github.com/dog/bBookStoreApi/store/factory"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8089", s)

	errChain, err := src.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
	}
	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChain:
		log.Println("web server run failed :", err)
		return
	case <-c:
		log.Println("bookstore program is exiting")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("bookstore program exit error : ", err)
		return
	}

	log.Println("bookstore program exit ok")

}
