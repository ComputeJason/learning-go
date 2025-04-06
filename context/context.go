package context

import (
	"context"
	"fmt"
	"net/http"
)

// explore how context helps us manage long running processes

// At Google, we require that Go programmers pass a Context parameter
// as the first argument to every function on the call path between
// incoming and outgoing requests. This allows Go code developed by many
// different teams to interoperate well.

// although having this dependency does make our function signature messy
// and should be changed at a language level, but not for now.

// its argued that you shouldn't pass values around in context because passed
// around in an untyped map and is not statically checked. However also argued you
// can you it to pass values orthogonal (unrelated/unused) to the request like
// traceID, can could potentially be useful, but if included in the function parameters
// would be pretty messy & unused too.

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}