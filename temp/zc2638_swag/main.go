package main

import (
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"net/http"
)

func handlePet(w http.ResponseWriter, _ *http.Request) {
	// your code here
}

type Pet struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []string `json:"tags"`
}

func main() {
	// define our endpoints
	//
	post := endpoint.New("post", "/pet", "Add a new pet to the store",
		endpoint.Handler(handlePet),
		endpoint.Description("Additional information on adding a pet to the store"),
		endpoint.Body(Pet{}, "Pet object that needs to be added to the store", true),
		endpoint.Response(http.StatusOK, Pet{}, "Successfully added pet"),
	)
	get := endpoint.New("get", "/pet/{petId}", "Find pet by ID",
		endpoint.Path(handlePet, "integer", "ID of pet to return", true),
		endpoint.Response(http.StatusOK, Pet{}, "successful operation"),
	)

	// define the swagger api that will contain our endpoints
	//
	api := swag.New(
		swag.Title("Swagger Petstore"),
		swag.Endpoints(post, get),
	)

	// iterate over each endpoint and add them to the default server mux
	//
	for path, endpoints := range api.Paths {
		http.Handle(path, endpoints)
	}

	// use the api to server the swagger.json file
	//
	enableCors := true
	http.Handle("/swagger", api.Handler(enableCors))
	http.ListenAndServe(":8086", nil)
}
