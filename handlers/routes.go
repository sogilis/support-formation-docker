package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Create_router(storagePath string) *mux.Router {
	// Routes order creation matter.
	r := mux.NewRouter()

	r.PathPrefix("/i/{slug}").Handler(GetImage{StoragePath: storagePath}).Methods("GET")
	r.PathPrefix("/i/upload").Handler(PostImage{StoragePath: storagePath}).Methods("POST")
	r.PathPrefix("/").Handler(GetIndex{BasePath: "/"}).Methods("GET")

	return r
}

func Launch(router *mux.Router, port uint) error {
	address := fmt.Sprintf("0.0.0.0:%v", port)
	log.Info("Launching HTTP server on address - ", address)

	srv := &http.Server{
		Handler: router,
		Addr:    address,
	}

	return srv.ListenAndServe()
}
