package main

import (
	//"assignments/databases"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"ID"`
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Email string `json:"Email"`
	DOB   string `json:"DOB"`
}

var Users []User

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	json.NewEncoder(w).Encode(Users)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", returnAllUsers)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	//databases.CreateUsersTable()

	Users = []User{
		User{ID: 1, Fname: "John", Lname: "Ray", Email: "john@gmail.com", DOB: "19-04-1998"},
		User{ID: 2, Fname: "Will", Lname: "Smith", Email: "will@gmail.com", DOB: "12-03-1990"},
	}

	handleRequests()
}
