package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var library []Book

	router := gin.Default()

	// ... (kode lainnya tetap sama)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"library": library})
	})

	router.POST("/add", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBind(&book); err != nil {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
			return
		}
		library = append(library, book)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func createBook() (Book, error) {
	var title, author, isbn string

	fmt.Print("Enter Book Title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter Author: ")
	fmt.Scanln(&author)

	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&isbn)

	err := validateISBN(isbn)
	if err != nil {
		return Book{}, err
	}

	return Book{Title: title, Author: author, ISBN: isbn}, nil
}
func validateISBN(isbn string) error {
	if len(isbn) != 13 {
		return errors.New("ISBN must be 13 characters long")
	}
	return nil
}
