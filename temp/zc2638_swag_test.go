package temp

import (
	"github.com/gin-gonic/gin"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/swagger"
	"io"
	"net/http"
	"testing"
)

func handle(c *gin.Context) {
	io.WriteString(c.Writer, "Insert your code here")
}

// Category example from the swagger pet store
type Category struct {
	ID   int64  `json:"category"`
	Name string `json:"name"`
}

// Pet example from the swagger pet store
type Pet struct {
	ID        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []string `json:"tags"`
}

func TestZc2638SwagGin(t *testing.T) {
	post := endpoint.New("post", "/pet", endpoint.Summary("Add a new pet to the store"),
		endpoint.Handler(handle),
		endpoint.Description("Additional information on adding a pet to the store"),
		endpoint.Body(Pet{}, "Pet object that needs to be added to the store", true),
		endpoint.Response(http.StatusOK, "Successfully added pet", endpoint.Schema(Pet{})),
	)
	get := endpoint.New("get", "/pet/{petId}", endpoint.Summary("Find pet by ID"),
		endpoint.Handler(handle),
		endpoint.Path("petId", "integer", "ID of pet to return", true),
		endpoint.Response(http.StatusOK, "successful operation", endpoint.Schema(Pet{})),
	)

	api := swag.New(
		swag.Endpoints(post, get),
	)

	router := gin.New()
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = swag.ColonPath(path)

		router.Handle(endpoint.Method, path, h)
	})

	enableCors := true
	router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))

	http.ListenAndServe(":8080", router)
}
