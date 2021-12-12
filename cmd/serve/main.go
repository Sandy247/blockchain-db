package main

import (
	"fmt"
	"net/http"

	"github.com/Sandy247/blockchain-db/internal/controllers"
)

func main() {
	c := controllers.NewController()
	mux := http.NewServeMux()
	mux.Handle("/", controllers.RootHandler(c.RequestsHandler))

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "localhost", "5000"),
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
