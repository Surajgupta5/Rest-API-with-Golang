package router

import (
	"restapiwithgo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newfile", middleware.Createfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/delete", middleware.Deletefile).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/rename", middleware.Renamefile).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/get", middleware.Get).Methods("GET", "OPTIONS")

	return router
}

