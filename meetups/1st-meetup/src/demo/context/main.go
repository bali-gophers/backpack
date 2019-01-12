package main

import (
	"context"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type ContextKey string

const (
	ReqIDKey ContextKey = "request_id"
)

func main() {
	fmt.Println("Starting server ...")

	middleware := logging()

	http.Handle("/hello", middleware(http.HandlerFunc(hello)))
	http.ListenAndServe(":8080", nil)
}

func logging() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			uid := uuid.NewV4()
			ctx := context.WithValue(context.Background(), ReqIDKey, uid.String())
			reqWithCtx := req.WithContext(ctx)
			logContext(ctx, "[%s] %s", req.Method, req.URL.Path)
			h.ServeHTTP(w, reqWithCtx)
		})
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if err := processSatu(ctx); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err := processDua(ctx); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err := processTiga(ctx); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func processSatu(ctx context.Context) error {
	logContext(ctx, "invoked processSatu()")
	return nil
}

func processDua(ctx context.Context) error {
	logContext(ctx, "invoked processDua()")
	return nil
}

func processTiga(ctx context.Context) error {
	logContext(ctx, "invoked processTiga()")
	return nil
}

func logContext(ctx context.Context, format string, args ...interface{}) {
	val := ctx.Value(ReqIDKey).(string)
	fmt.Printf("ReqID: %s - %s\n", val, fmt.Sprintf(format, args...))
}
