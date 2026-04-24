package main

import (
	"book_inventory/app"
	"book_inventory/auth"
	"book_inventory/db"
	"book_inventory/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	// home
	router.GET("/", auth.HomeHandler)

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	// add book
	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	// update book
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBook)

	// delete book
	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
