package main

import (
	//"assignments/databases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"ID"`
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Email string `json:"Email"`
	DOB   string `json:"DOB"`
}

var Users []User

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var user User
	json.Unmarshal(reqBody, &user)
	// update our global Articles array to include
	// our new Article
	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	i, err := strconv.Atoi(key)
	if err != nil {
		log.Panic(err)
	}

	for _, User := range Users {
		if User.ID == i {
			json.NewEncoder(w).Encode(User)
		}
	}
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(Users)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/user/{id}", returnSingleUser)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", createNewUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	//databases.CreateUsersTable()
	fmt.Println("Rest API v2.0 - Mux Routers")
	Users = []User{
		User{ID: 1, Fname: "John", Lname: "Ray", Email: "john@gmail.com", DOB: "19-04-1998"},
		User{ID: 2, Fname: "Will", Lname: "Smith", Email: "will@gmail.com", DOB: "12-03-1990"},
	}

	handleRequests()
}
