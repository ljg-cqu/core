// https://github.com/swaggest/openapi-go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swaggest/openapi-go/openapi3"
	"log"
	"net/http"
	"time"
)

func main() {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("Things API").
		WithVersion("1.2.3").
		WithDescription("Put something here")

	type req struct {
		ID     string `path:"id" example:"XXX-XXXXX"`
		Locale string `query:"locale" pattern:"^[a-z]{2}-[A-Z]{2}$"`
		Title  string `json:"string"`
		Amount uint   `json:"amount"`
		Items  []struct {
			Count uint   `json:"count"`
			Name  string `json:"name"`
		} `json:"items"`
	}

	type resp struct {
		ID     string `json:"id" example:"XXX-XXXXX"`
		Amount uint   `json:"amount"`
		Items  []struct {
			Count uint   `json:"count"`
			Name  string `json:"name"`
		} `json:"items"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	putOp := openapi3.Operation{}

	cobra.CheckErr(reflector.SetRequest(&putOp, new(req), http.MethodPut))
	cobra.CheckErr(reflector.SetJSONResponse(&putOp, new(resp), http.StatusOK))
	cobra.CheckErr(reflector.SetJSONResponse(&putOp, new([]resp), http.StatusConflict))
	cobra.CheckErr(reflector.Spec.AddOperation(http.MethodPut, "/things/{id}", putOp))

	getOp := openapi3.Operation{}

	cobra.CheckErr(reflector.SetRequest(&getOp, new(req), http.MethodGet))
	cobra.CheckErr(reflector.SetJSONResponse(&getOp, new(resp), http.StatusOK))
	cobra.CheckErr(reflector.Spec.AddOperation(http.MethodGet, "/things/{id}", getOp))

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(schema))
}
