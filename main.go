package main

import (
	"net/http"
	"gqljokes/handlers"
)

func main() {
	http.Handle("/graphql", http.HandlerFunc(handlers.GraphQLHandler))
	http.ListenAndServe(":8080", nil)
}
