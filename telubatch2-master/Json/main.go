package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Customer struct {
	Address struct {
		City   string `json:"city"`
		State  string `json:"state"`
		Street string `json:"street"`
		Zip    string `json:"zip"`
	} `json:"address"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Customer

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	// fmt.Fprintln(w, "First Name : "+request.FirstName)
	// fmt.Fprintln(w, "City Name : "+request.Address.City)
	if r.Method == "POST" {

		//Tugas insert kan ke table Employeees
		stmt, err := db.Prepare("INSERT INTO employees (LastName,FirstName,Address, City) VALUES (?,?,?,?)")

		_, err = stmt.Exec(request.LastName, request.FirstName, request.Address, request.Address.City)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Berhasil Ditambahkan")
		}
	}

}

func getToCustomers(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Customer

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	// fmt.Fprintln(w, "First Name : "+request.FirstName)
	// fmt.Fprintln(w, "City Name : "+request.Address.City)
	if r.Method == "POST" {

		//Tugas insert kan ke table Employeees
		stmt, err := db.Prepare("INSERT INTO customers (LastName,FirstName,Address,City) VALUES (?,?,?,?)")

		_, err = stmt.Exec(request.LastName, request.FirstName)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Berhasil Ditambahkan")
		}
	}

}

func getToOrders(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Customer

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	// fmt.Fprintln(w, "First Name : "+request.FirstName)
	// fmt.Fprintln(w, "City Name : "+request.Address.City)
	if r.Method == "POST" {

		//Tugas insert kan ke table Employeees
		stmt, err := db.Prepare("INSERT INTO orders (LastName,FirstName) VALUES (?,?)")

		_, err = stmt.Exec(request.LastName, request.FirstName)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Berhasil Ditambahkan")
		}
	}

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/customers", getCustomers).Methods("POST")
	r.HandleFunc("/customers", getToCustomers).Methods("POST")
	r.HandleFunc("/customers", getToOrders).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
