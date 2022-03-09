package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// router basing on chi
	chiR := chi.NewRouter()
	chiR.Use(middleware.Logger)
	chiR.Get("/chi_welcome", Welcom)

	// router basing on gin
	ginR := gin.Default()
	ginR.GET("/gin_welcome", gin.WrapF(Welcom))

	// starting servers

	ctx, shutdown := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)

	_ = gctx

	g.Go(func() error {
		log.Println("http://localhost:3000/chi_welcome")
		return http.ListenAndServe(":3000", chiR)
	})

	g.Go(func() error {
		log.Println("http://localhost:8081/gin_welcome")
		return ginR.Run(":8081")
		//return http.ListenAndServe(":8080", ginR)
	})

	// Handle graceful shutdown.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Gracefully shutting down the server...")

	shutdown()

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func Welcom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
