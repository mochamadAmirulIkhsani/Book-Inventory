package auth

import (
	"book_inventory/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context) {
	var credential models.Login
	err := c.Bind(&credential)

	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Username/password is invalid request",
		})
	}

	if credential.Username != models.USER || credential.Password != models.PASSWORD {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "Username/password is invalid",
		})
	} else {
		// token

		claim := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			Issuer:    "book inventoory",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(models.SECRET))

		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "Username/password is invalid",
			})
			c.Abort()
		}

		q := url.Values{}
		q.Set("auth", token)
		location := url.URL{Path: "/books", RawFragment: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	}
}
