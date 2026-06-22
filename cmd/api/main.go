package main

import (
	"fmt"
	"net/http"

	"github.com/RAHULBONEY/goApi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal("Error: ", err)
	}

}
