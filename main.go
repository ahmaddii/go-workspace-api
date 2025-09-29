package main

// Today we created an Endpoint Rest API with get post put or delete methods ok

// in which we are users endpoint data in json format on the browser which is our client

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct { // for now creating dummy user as we dont have real DB

	Id   int    `json:id`
	Name string `json:name`
}

// Fake DB Function which has slice of data stored

func getUsersFromDb() []User {

	return []User{

		{Id: 1, Name: "Ahmad"},
		{Id: 2, Name: "Ali"},
		{Id: 3, Name: "Fahad"},
		{Id: 4, Name: "Abuzer"},
	}

}

func getUserData(w http.ResponseWriter, r *http.Request) {

	users := getUsersFromDb()

	usersJson, err := json.Marshal(users)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Header().Set("contenttype", "application/json")

	w.Write(usersJson) // actually send data to the client or on browser

}

func main() {

	r := mux.NewRouter() // create new router

	// register route

	r.HandleFunc("/users", getUserData).Methods(http.MethodGet)

	//r.HandleFunc("/users/{id}",getUserId).Methods(http.MethodGet) for getting one object data from an users end point wtih Id =2

	//r.HandleFunc("/users", createUser).Methods(http.MethodPost) if user want to save the new user and after that send confirmation

	//r.HandleFunc("/user/{id}",updateUser).Methods(http.MethodPut) if you want to update specific person data from name ahmd to adeel for id 1

	//r.HandleFunc("user/{id}",deleteUser).Methods(http.MethodDelete) api deletes that specifc id user and return the succes msg

	// users is the end point where you get or fetch the data from get user data func

	// path users and handler where to send getuser data and method get fetch data of users

	http.ListenAndServe(":8000", r)



}
