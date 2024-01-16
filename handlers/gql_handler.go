package handlers

import (
	"net/http"
	"github.com/graphql-go/handler"
	"gqljokes/schemas"
)

// NewGraphQLHandler creates a new GraphQL handler with the given schema.
func NewGraphQLHandler() http.Handler {
	h := handler.New(&handler.Config{
		Schema: &schemas.Schema,
		Pretty: true,
	})
	return h
}

// GraphQLHandler is an HTTP handler for serving GraphQL requests.
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests for GraphQL queries
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h := NewGraphQLHandler()
	h.ServeHTTP(w, r)
}
