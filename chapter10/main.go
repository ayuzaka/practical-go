package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/time/rate"
)

type Comment struct {
	Message  string `validate:"required,min=1,max=140"`
	UserName string `validate:"required,min=1,max=15"`
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode

		log.Printf("%d %s", statusCode, http.StatusText(statusCode))
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

var limiter = rate.NewLimiter(rate.Every(time.Second/1), 10)

func LimitHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start %s\n", r.URL)
		next.ServeHTTP(w, r)
		log.Printf("finish %s\n", r.URL)
	})
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func pani(w http.ResponseWriter, r *http.Request) {
	panic("panic occur")
}

func main() {
	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			mutex.RLock()
			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()

		case http.MethodPost:
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			validate := validator.New()
			if err := validate.Struct(c); err != nil {
				var out []string
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					for _, fe := range ve {
						switch fe.Field() {
						case "Message":
							out = append(out, fmt.Sprintf("Message は1~140文字です"))
						case "UserName":
							out = append(out, fmt.Sprintf("UserName は1~15文字です"))
						}
					}
				}

				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, strings.Join(out, ",")), http.StatusBadRequest)
				return
			}

			mutex.Lock()
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))

		default:
			http.Error(w, `{"status":"permits only GET or POST"}`, http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		word := r.FormValue("searchword")
		log.Printf("searchword = %s\n", word)

		words, ok := r.Form["searchword"]
		log.Printf("search words = %v has values %v\n", words, ok)

		log.Print("all queries")
		for key, values := range r.Form {
			log.Printf(" %s: %v\n", key, values)
		}

	})

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 * 1024 * 1024)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		f, h, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(h.Filename)

		o, err := os.Create(h.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer o.Close()

		_, err = io.Copy(o, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		value := r.PostFormValue("data")
		log.Printf(" value = %s\n", value)
	})

	h := MiddlewareLogging(http.HandlerFunc(Health))
	http.Handle("/health", http.TimeoutHandler(h, 5, "request timeout"))

	http.Handle("/logger", wrapHandlerWithLogging(http.HandlerFunc(Health)))
	http.Handle("/panic", Recovery(http.HandlerFunc(pani)))

	http.Handle("/limit", LimitHandler(http.HandlerFunc(Health)))

	http.ListenAndServe(":8888", nil)
}
