package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{
	{ID: 1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go"},
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.ID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func UpdateBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var updatedBook Book

	err := ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i, book := range BookDatas {
		if bookID == book.ID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].ID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func GetBooks(ctx *gin.Context) {
	var bookData []Book

	for _, book := range BookDatas {
		detailBook := Book{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Desc:   book.Desc,
		}
		bookData = append(bookData, detailBook)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, "Deleted")
}
