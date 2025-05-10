package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abisalde/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Printf("🚀 Starting server at Port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
