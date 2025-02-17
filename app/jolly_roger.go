package jolly_roger

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var persistence StorageEngine

func Serve() {
	persistence, err := NewFromViperSettings()

	if err != nil {
		// TODO meaninful error handling
		fmt.Printf("could not load persistence config: %s", err)
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alive"))
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/{vendor}", func(r chi.Router) {
		r.Post("/receive", func(w http.ResponseWriter, r *http.Request) {
			vendor := chi.URLParam(r, "vendor")
			if vendor == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("vendor parameter required"))
				return
			}

			//io.Copy(os.Stdout, r.Body)
			// TODO: Read request body
			payload, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to read request body: %v", err)))
				return
			}

			err = persistence.Store(vendor, payload)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to save webhook: %v", err)))
				return
			}

			w.Write([]byte(fmt.Sprintf("received and saved webhook for %s", vendor)))
		})

	})
	http.ListenAndServe(":3000", r)
}
