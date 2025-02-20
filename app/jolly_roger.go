package jolly_roger

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"jolly_roger/app/storage"
	"jolly_roger/app/stream"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/spf13/viper"
)

var persistence storage.StorageEngine
var streamOut stream.StreamEngine

func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Read config from app root directory
		viper.AddConfigPath(".")
		viper.SetConfigType("toml")
		viper.SetConfigName("jolly_roger.toml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("no config found")
	}
}

func Serve() {
	persistence, err := storage.NewFromViperSettings()

	if err != nil {
		// TODO meaninful error handling
		fmt.Printf("could not load persistence config: %s", err)
		return
	}

	streamOut, err := stream.NewFromViperSettings()

	if err != nil {
		// TODO meaninful error handling
		fmt.Printf("could not load streaming config: %s", err)
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

	r.Route("/{vendor}", func(r chi.Router) {
		r.Post("/receive", func(w http.ResponseWriter, r *http.Request) {
			vendor := chi.URLParam(r, "vendor")
			if vendor == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("vendor parameter required"))
				return
			}

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

			err = streamOut.Pub(vendor, payload)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("failed to stream webhook: %v", err)))
				return
			}

			w.Write([]byte(fmt.Sprintf("received and saved webhook for %s", vendor)))
		})

	})
	http.ListenAndServe(":3000", r)
}
