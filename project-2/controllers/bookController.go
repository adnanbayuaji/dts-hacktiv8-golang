package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"project-2/database"
	"project-2/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func CreateBook(ctx *gin.Context) {
	var newBook Book
	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	Book := models.Book{
		Title:       newBook.Title,
		Author:      newBook.Author,
		Description: newBook.Desc,
	}

	err = db.Create(&Book).Error

	if err != nil {
		fmt.Println("Error creating book data:", err)
		return
	}

	fmt.Println("New Book Data:", Book)
	ctx.JSON(http.StatusCreated, "Created")
}

func UpdateBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var updatedBook Book

	err := ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	book := models.Book{}
	err = db.Model(&book).Where("id = ?", bookID).Updates(models.Book{Title: updatedBook.Title, Author: updatedBook.Author, Description: updatedBook.Desc}).Error

	if err != nil {
		fmt.Println("Error updating book data:", err)
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	db := database.GetDB()

	book := models.Book{}

	err := db.First(&book, "id = ?", bookID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Book data not found")
			return
		}
		print("Error finding book:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetBooks(ctx *gin.Context) {
	db := database.GetDB()

	books := []models.Book{}

	err := db.Find(&books).Error

	if err != nil {
		fmt.Println("Error getting book datas with products:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": books,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	db := database.GetDB()
	book := models.Book{}
	err := db.Where("id = ?", bookID).Delete(&book).Error
	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Book with id %d has been successfully deleted", bookID)

	ctx.JSON(http.StatusOK, "Deleted")
}
