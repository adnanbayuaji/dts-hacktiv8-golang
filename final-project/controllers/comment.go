package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	appJSON = "application/json"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Title: Comment.Title, Description: Comment.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	comments := []models.Comment{}
	err := db.Find(&comments).Error

	if err != nil {
		fmt.Println("Error getting user datas:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"comment": comments,
	})
}

func GetOneComment(ctx *gin.Context) {
	commentID, _ := strconv.Atoi(ctx.Param("commentId"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	comment := models.Comment{}
	err := db.First(&comment, "id = ?", commentID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}
