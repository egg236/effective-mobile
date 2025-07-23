package api

import (
	"bytes"
	"effective-mobile/app"
	"effective-mobile/config"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"time"
)

type Server struct {
	app app.App
	http.Server
}

func NewServer(a app.App) *Server {
	s := &Server{app: a}
	s.Addr = fmt.Sprintf(":%s", config.Cfg.Port())

	mux := mux2.NewRouter()
	mux.Use(Middleware)
	mux.HandleFunc("/records", s.handleReadRecords).Methods(http.MethodGet)
	mux.HandleFunc("/records/sum", s.handleReadRecordsSum).Methods(http.MethodGet)

	mux.HandleFunc("/record/{recordID:[0-9]+}", s.handleReadRecordByID).Methods(http.MethodGet)

	mux.HandleFunc("/record", s.handleCreateRecord).Methods(http.MethodPost)

	mux.HandleFunc("/record", s.handleUpdateRecord).Methods(http.MethodPut)

	mux.HandleFunc("/record/{recordID:[0-9]+}", s.handleDeleteRecord).Methods(http.MethodDelete)

	s.Handler = mux

	slog.Info("server initialized", "address", s.Addr)

	return s
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("fatal error while processing request", "error", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		id := rand.Int() % 1000
		timeStart := time.Now()

		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			slog.Error("failed to read request body", "error", err.Error())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		rw := &ResponseWrapper{core: w}
		slog.Info("Request received", "id", id, "method", r.Method, "path", r.URL.RequestURI(), "body", string(body))
		next.ServeHTTP(rw, r)
		slog.Info("Request processed", "id", id, "method", r.Method, "path", r.URL.RequestURI(), "status", rw.Status, "time", time.Since(timeStart), "body", string(rw.Body))
	})
}

type ResponseWrapper struct {
	core   http.ResponseWriter
	Status int
	Body   []byte
}

func (rw *ResponseWrapper) Write(b []byte) (int, error) {
	if rw.Body == nil {
		rw.Body = make([]byte, 0)
		rw.Body = append(rw.Body, b...)
	} else {
		rw.Body = append(rw.Body, b...)
	}

	return rw.core.Write(b)
}

func (rw *ResponseWrapper) WriteHeader(statusCode int) {
	rw.Status = statusCode
	rw.core.WriteHeader(statusCode)
}

func (rw *ResponseWrapper) Header() http.Header {
	return rw.core.Header()
}
