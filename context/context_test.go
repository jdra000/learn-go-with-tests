package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		var (
			// ctx is the Context for this handler
			ctx    context.Context
			// Calling cancel closes the ctx.Done channel
			cancel context.CancelFunc
		)

		// WithCancel returns a copy of parent whose Done channel is closed as soon as parent.Done is closed or cancel is called
		ctx, cancel = context.WithCancel(request.Context()) 
		// call Cancel after the function elapses
		time.AfterFunc(5*time.Millisecond, cancel)
		// WithContext returns a new group and an associated Context derived from its parameter
		request = request.WithContext(ctx) // New context in request
		
		// In order to check if the reponse is written or not, we created the SpyResponse with our customized ResponseWriter interface
		// It should not be written since context should be Done 
		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})

}
