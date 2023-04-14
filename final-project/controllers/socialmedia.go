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

func CreateSocialmedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socialmedia := models.Socialmedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID

	err := db.Debug().Create(&Socialmedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Socialmedia)
}

func UpdateSocialmedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Socialmedia := models.Socialmedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID
	Socialmedia.ID = uint(socialmediaId)

	err := db.Model(&Socialmedia).Where("id = ?", socialmediaId).Updates(models.Socialmedia{Title: Socialmedia.Title, Description: Socialmedia.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Socialmedia)
}

func GetAllSocialmedia(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	socialmedias := []models.Socialmedia{}
	err := db.Find(&socialmedias).Error

	if err != nil {
		fmt.Println("Error getting user datas:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"socialmedia": socialmedias,
	})
}

func GetOneSocialmedia(ctx *gin.Context) {
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaId"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	socialmedia := models.Socialmedia{}
	err := db.First(&socialmedia, "id = ?", socialmediaID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"socialmedia": socialmedia,
	})
}
