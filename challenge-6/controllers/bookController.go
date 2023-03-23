package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "adm"
	dbname   = "db-go-sql"
)

func ConnectDatabase() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func CreateBook(ctx *gin.Context) {
	db = ConnectDatabase()
	defer db.Close()

	var newBook Book
	var book = Book{}

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	INSERT INTO books (title, author, description)	
	VALUES ($1, $2, $3)
	RETURNING *
	`

	err = db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Desc).Scan(&book.ID, &book.Title, &book.Author, &book.Desc)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Book Data : %+v\n", book)
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

	db = ConnectDatabase()
	defer db.Close()

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, bookID, updatedBook.Title, updatedBook.Author, updatedBook.Desc)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount:", count)

	ctx.JSON(http.StatusOK, "Updated")
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	var results = []Book{}

	db = ConnectDatabase()
	defer db.Close()

	sqlStatement := `
	SELECT *
	FROM books
	WHERE id = $1;
	`

	rows, err := db.Query(sqlStatement, bookID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": results,
	})
}

func GetBooks(ctx *gin.Context) {
	var results = []Book{}

	db = ConnectDatabase()
	defer db.Close()

	sqlStatement := "SELECT * from books"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": results,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	db = ConnectDatabase()
	defer db.Close()

	sqlStatement := `
	DELETE from books
	WHERE id = $1
	`

	res, err := db.Exec(sqlStatement, bookID)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)

	ctx.JSON(http.StatusOK, "Deleted")
}
