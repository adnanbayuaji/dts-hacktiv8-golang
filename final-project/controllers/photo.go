package controllers

import (
	"errors"
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePhoto godoc
// @Summary Post details for a given Id
// @Description Post details of Photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "create photo"
// @Success 200 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Update Photo identified by the given Id
// @Description Update the Photo corresponding to the input
// @Tags photos
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo to be updated"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [patch]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// GetAllPhoto godoc
// @Summary Get details
// @Description Get details of all Photo
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetAllPhoto(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	photos := []models.Photo{}
	err := db.Find(&photos).Error

	if err != nil {
		fmt.Println("Error getting user datas:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"photo": photos,
	})
}

// GetOnePhoto godoc
// @Summary Get details for a given Id
// @Description Get details of Photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetOnePhoto(ctx *gin.Context) {
	photoID, _ := strconv.Atoi(ctx.Param("photoId"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	photo := models.Photo{}
	err := db.First(&photo, "id = ?", photoID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"photo": photo,
	})
}

// DeletePhoto godoc
// @Summary Delete photo identified by the given Id
// @Description Delete the Photo corresponding to the input
// @Tags photos
// @Accept json
// @Produce json
// @Param Id path int true "ID of the photo to be deleted"
// @Success 204 "No Content"
// @Router /photos/{id} [delete]
func DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Photo := models.Photo{}
	err := db.Where("id = ?", photoId).Delete(&Photo).Error
	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Photo with id %d has been successfully deleted", photoId)

	ctx.JSON(http.StatusOK, "Deleted")
}
