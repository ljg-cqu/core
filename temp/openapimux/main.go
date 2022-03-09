// https://github.com/mnf-group/openapimux
package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	mux "github.com/mnf-group/openapimux"
	"log"
	"net/http"
)

type fooHandler struct{}

func (f fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	r, err := mux.NewRouter("./openapi.yaml")
	if err != nil {
		panic(err)
	}

	r.UseHandlers(map[string]http.Handler{
		"getFoo": fooHandler{},
	})

	r.UseMiddleware(
		middleware.Recoverer,
		middleware.RequestID,
		//middleware.DefaultCompress,
	)

	r.ErrorHandler = func(w http.ResponseWriter, r *http.Request, data string, code int) {
		w.WriteHeader(code)
		if code == http.StatusInternalServerError {
			fmt.Println("Fatal:", data)
			w.Write([]byte("Oops"))
		} else {
			w.Write([]byte(data))
		}
	}

	log.Fatal(http.ListenAndServe(":8080", r))
}
