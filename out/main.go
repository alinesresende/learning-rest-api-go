package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/users", getUsers).Methods("GET")

	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", muxRouter))

}

type User struct {
	Id   int
	Name string
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id:   1,
			Name: "Aline",
		}})
}
