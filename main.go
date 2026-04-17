package main

import (
	"book_inventory/app"
	"book_inventory/db"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	// home
	router.GET("/books", handler.GetBooks)

	router.Run()
}
