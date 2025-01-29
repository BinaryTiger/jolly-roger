package jolly_roger

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Serve() {
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

			// TODO: Read request body
			var payload []byte
			_, err := r.Body.Read(payload)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("failed to read request body"))
				return
			}

			// TODO: Insert webhook into database
			// Example SQL: INSERT INTO webhooks (vendor, payload, received_at) VALUES ($1, $2, NOW())
			/*
			if _, err := db.Exec(
				"INSERT INTO webhooks (vendor, payload, received_at) VALUES ($1, $2, NOW())",
				vendor,
				payload,
			); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("failed to save webhook"))
				return
			}
			*/

			w.Write([]byte(fmt.Sprintf("received and saved webhook for %s", vendor)))
		})

	})
	http.ListenAndServe(":3000", r)
}
