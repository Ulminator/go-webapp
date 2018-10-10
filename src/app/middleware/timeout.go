package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	r.WithContext(ctx)        // modified request to have new timeout context
	ch := make(chan struct{}) // channel receives a signal if request finishes
	go func() {               // call the next handler using a Go routine
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{} // if successful, a signal is sent to ch channel
	}()
	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
