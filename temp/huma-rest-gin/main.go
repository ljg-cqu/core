// https://github.com/danielgtaylor/huma/
package main

import (
	"context"
	"fmt"
	"github.com/danielgtaylor/huma"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/rest/nethttp"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var version = "v1.0.0"

func main() {
	// router basing on huma
	humaR := huma.New("Contract review API\nThis is just a demonstration of contract review API.", version)
	humaR.DocsHandler(huma.SwaggerUIHandler(huma.New("Test API", version)))

	fileUploadResource := humaR.Resource("/file_upload")
	RunFileUploader(fileUploadResource)

	// router basing on gin
	ginR := gin.Default()
	ginR.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	ginR.GET("/ping2", gin.WrapF(Ping2))
	ginR.POST("/jsonBody", gin.WrapH(nethttp.NewHandler(jsonBody())))

	// starting servers

	ctx, shutdown := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)

	_ = gctx

	g.Go(func() error {
		log.Println("http://localhost:8888/docs")
		return http.ListenAndServe(":8888", humaR)
	})

	g.Go(func() error {
		log.Println("http://localhost:8082/ping")
		return ginR.Run(":8082")
		//return http.ListenAndServe(":8080", ginR)
	})
	//
	//g.Go(func() error {
	//	log.Println("http://localhost:8011/docs")
	//	return http.ListenAndServe(":8011", NewRouter())
	//})

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

func Ping2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong2"))
}
