package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/simple_CRUD_API_With_Golang/services"
	"log"
	"net/http"
)

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", services.GetMoives).Methods("GET")
	r.HandleFunc("/movies", services.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", services.GetMoiveById).Methods("GET")
	r.HandleFunc("/movies/{id}", services.UpdataMoiveByID).Methods("PUT")
	r.HandleFunc("/movies/{id}", services.DeleteMovieByID).Methods("DELETE")

	fmt.Println("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
