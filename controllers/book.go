package controllers

import (
	"books-app/models"
	"books-app/repository/book"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var books []models.Book

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		c := user.(*jwt.Token).Claims.(jwt.MapClaims)
		var userId int = int(c["userId"].(float64))

		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		book.UserId = userId

		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		c := user.(*jwt.Token).Claims.(jwt.MapClaims)
		
		var userId int = int(c["userId"].(float64))
		var book models.Book

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		book.UserId = userId

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		c := user.(*jwt.Token).Claims.(jwt.MapClaims)
		
		var userId int = int(c["userId"].(float64))
		var book models.Book
		var bookID int

		book.UserId = userId

		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		bookID = bookRepo.AddBook(db, book)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		c := user.(*jwt.Token).Claims.(jwt.MapClaims)
		
		var userId int = int(c["userId"].(float64))
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)

		book.UserId = userId
		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		bookRepo := bookRepository.BookRepository{}
		
		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		rowsDeleted := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
