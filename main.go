package main

import (
	"books-app/controllers"
	"books-app/driver"
	"books-app/models"
	"books-app/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/books", utils.TokenVerifyMiddleWare(controller.GetBooks(db))).Methods("GET")
	router.HandleFunc("/books/{id}", utils.TokenVerifyMiddleWare(controller.GetBook(db))).Methods("GET")
	router.HandleFunc("/books", utils.TokenVerifyMiddleWare(controller.AddBook(db))).Methods("POST")
	router.HandleFunc("/books", utils.TokenVerifyMiddleWare(controller.UpdateBook(db))).Methods("PUT")
	router.HandleFunc("/books/{id}", utils.TokenVerifyMiddleWare(controller.RemoveBook(db))).Methods("DELETE")

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	router.HandleFunc("/users", utils.TokenVerifyMiddleWare(controller.GetUsers(db))).Methods("GET")
	router.HandleFunc("/users/{id}", utils.TokenVerifyMiddleWare(controller.GetUser(db))).Methods("GET")
	router.HandleFunc("/users/{id}", utils.TokenVerifyMiddleWare(controller.RemoveUser(db))).Methods("DELETE")

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
