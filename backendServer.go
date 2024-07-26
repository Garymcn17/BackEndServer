package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID	string `json:"id"`
	Username 	string `json:"username"`
	Email string `json:"email"`
}

var users []User



func main() {
	users = append(users, User{ID: "1", Username: "Gary Tester", Email: "TesterGary@gmail.com"})
	users = append(users, User{ID: "2", Username: "John Smalls", Email: "Jsmalls@gmail.com"})


	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/createUser", createUsers).Methods("POST")
	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(users)
}

func createUsers(w http.ResponseWriter, r *http.Request){
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

// func handler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Hello, world!")
// }