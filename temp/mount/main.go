package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"net/http"
	"time"
)

func restService() *web.Service {
	s := web.DefaultService()

	// Init API documentation schema.
	s.OpenAPI.Info.Title = "Basic Example"
	s.OpenAPI.Info.WithDescription("This app showcases a trivial REST API.")
	s.OpenAPI.Info.Version = "v1.2.3"

	// Setup middlewares.
	s.Use(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	// Declare input port type.
	type helloInput struct {
		Locale string `query:"locale" default:"en-US" pattern:"^[a-z]{2}-[A-Z]{2}$" enum:"ru-RU,en-US"`
		Name   string `path:"name" minLength:"3"` // Field tags define parameter location and JSON schema constraints.
	}

	// Declare output port type.
	type helloOutput struct {
		Now     time.Time `header:"X-Now" json:"-"`
		Message string    `json:"message"`
	}

	messages := map[string]string{
		"en-US": "Hello, %s!",
		"ru-RU": "Привет, %s!",
	}

	// Create use case interactor with references to input/output types and interaction function.
	u := usecase.NewIOI(new(helloInput), new(helloOutput), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*helloInput)
			out = output.(*helloOutput)
		)

		msg, available := messages[in.Locale]
		if !available {
			return status.Wrap(errors.New("unknown locale"), status.InvalidArgument)
		}

		out.Message = fmt.Sprintf(msg, in.Name)
		out.Now = time.Now()

		return nil
	})

	// Describe use case interactor.
	u.SetTitle("Greeter")
	u.SetDescription("Greeter greets you.")

	u.SetExpectedErrors(status.InvalidArgument)

	// Add use case handler to router.
	s.Get("/rest/hello/{name}", u)

	// Swagger UI endpoint at /docs.
	s.Docs("/rest/docs", swgui.New)
	return s
}

func main() {
	admin := http.NewServeMux()

	// Prefix all paths with the mount point. A ServeMux matches
	// the full path, even when invoked from another ServeMux.
	mountPoint := "/admin"
	admin.HandleFunc(mountPoint+"/", root)
	admin.HandleFunc(mountPoint+"/foo", foo)

	// Add a trailing "/" to the mount point to indicate a subtree match.
	http.Handle(mountPoint+"/", admin)
	http.HandleFunc("/ping", pong)

	// mount gin
	ginR := gin.Default()
	ginR.GET("/gin/", func(ctx *gin.Context) {
		ctx.JSON(200, "gin root")
	})
	ginR.GET("/gin/echo", func(ctx *gin.Context) {
		ctx.JSON(200, "gin echo")
	})
	http.Handle("/gin/", ginR.Handler())

	// mount swagest/rest service
	http.Handle("/rest/", restService())

	http.ListenAndServe(":4567", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Admin: ROOT")
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Admin: FOO")
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!")
}
