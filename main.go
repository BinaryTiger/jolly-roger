/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"binarytiger/jolly_roger/cmd"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func main() {
	saveRawHook()
	cmd.Execute()
}

func serve() {
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
			if vendor := chi.URLParam(r, "vendor"); vendor != "" {
				w.Write([]byte(fmt.Sprintf("received for %s", vendor)))
			} else {
				w.Write([]byte("received at top level"))
			}
		})

	})
	http.ListenAndServe(":3000", r)
}

func saveRawHook() {
	var version string
	db, _ := sql.Open("sqlite3", "file:local.db") // #TODO load as config
	db.QueryRow(`SELECT sqlite_version()`).Scan(&version)
	fmt.Println(version)
}
