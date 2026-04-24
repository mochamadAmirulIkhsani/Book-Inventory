package app

import (
	"book_inventory/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) GetBooks(c *gin.Context) {
	var books []models.Books
	auth := c.Query("auth")

	h.DB.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": books,
		"auth":    auth,
	})
}

func (h *Handler) GetBookById(c *gin.Context) {
	bookId := c.Param("id")

	var books models.Books

	if h.DB.Find(&books, "id=?", bookId).Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "book.html", gin.H{
		"title":   books.Title,
		"payload": books,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) AddBook(c *gin.Context) {
	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title": "add book",
		"auth":  c.Query("auth"),
	})
}

func (h *Handler) PostBook(c *gin.Context) {
	var books models.Books

	c.Bind(&books)
	h.DB.Create(&books)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")

	if h.DB.Find(&books, "id=?", bookId).Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "not found",
			"auth":  c.Query("auth"),
		})
	}

	c.HTML(http.StatusOK, "formBook.html", gin.H{
		"title":   "add book",
		"payload": books,
		"auth":    c.Query("auth"),
	})
}

func (h *Handler) PutBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")

	if h.DB.Find(&books, "id=?", bookId).Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "not found",
			"auth":  c.Query("auth"),
		})
	}

	var reqBook = books
	c.Bind(&reqBook)

	h.DB.Model(&books).Where("id=?", bookId).Updates(reqBook)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/book/%s?auth=%s", bookId, c.PostForm("auth")))

}

func (h *Handler) DeleteBook(c *gin.Context) {
	var books models.Books

	bookId := c.Param("id")
	h.DB.Delete(&books, "id=?", bookId)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/books?auth=%s", c.PostForm("auth")))
}
