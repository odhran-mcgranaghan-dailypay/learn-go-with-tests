package context

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
)

type Store interface {
    Fetch(ctx context.Context) (string, error)
}

type SpyStore struct {
    response string
}

// Microextension - context values
// some request ID key, define a type for it,
// create a default const value of 0
type key int
const requestIDKey key = 0

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
    data := make(chan string, 1)

    go func() {
        var result string
        for _, c := range s.response {
            select {
            case <-ctx.Done():
                log.Println("spy store got cancelled")
                return
            default:
                time.Sleep(10 * time.Millisecond)
                result += string(c)
            }
        }
        data <- result
    }()

    select {
    case <-ctx.Done():
        return "", ctx.Err()
    case res := <-data:
        return res, nil
    }
}

func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		// Microextension
		// create a new context with a value of request ID key
		// withValue function returns a new context with the value

        // storing short lived data is deemed okay
		ctx := context.WithValue(r.Context(), requestIDKey, "12345")

        data, err := store.Fetch(ctx)
        if err != nil {
            return
        }
        requestID := ctx.Value(requestIDKey).(string)
        fmt.Fprintf(w, "Request ID: %s, Data: %s", requestID, data)    
	}
}

func main() {
    store := &SpyStore{response: "hello, world"}
    svr := Server(store)
    http.ListenAndServe(":8080", svr)
}
