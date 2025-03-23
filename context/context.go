package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// * Spy for Store which we use in tests
type SpyStore struct {
	response string
}

// * We need a way to test we do not write any kind of response in the error case
type SpyResponseWriter struct {
	written bool
}

// * Methods that satisfy the ResponseWriter interface
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

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

// * Server code
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log error however you like
		}
		fmt.Fprint(w, data)
	}
}
