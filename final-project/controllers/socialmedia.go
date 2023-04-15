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

var (
	appJSON = "application/json"
)

// CreateSocialmedia godoc
// @Summary Post details for a given Id
// @Description Post details of Socialmedia corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param models.Socialmedia body models.Socialmedia true "create socialmedia"
// @Success 200 {object} models.Socialmedia
// @Router /socialmedias [post]
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

// UpdateSocialmedia godoc
// @Summary Update Socialmedia identified by the given Id
// @Description Update the Socialmedia corresponding to the input
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Id path int true "ID of the socialmedia to be updated"
// @Success 200 {object} models.Socialmedia
// @Router /socialmedias/{id} [patch]
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

	err := db.Model(&Socialmedia).Where("id = ?", socialmediaId).Updates(models.Socialmedia{Name: Socialmedia.Name, SocialMediaUrl: Socialmedia.SocialMediaUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Socialmedia)
}

// GetAllSocialmedia godoc
// @Summary Get details
// @Description Get details of all Socialmedia
// @Tags socialmedias
// @Accept json
// @Produce json
// @Success 200 {object} models.Socialmedia
// @Router /socialmedias [get]
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

// GetOneSocialmedia godoc
// @Summary Get details for a given Id
// @Description Get details of Socialmedia corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Id path int true "ID of the socialmedia"
// @Success 200 {object} models.Socialmedia
// @Router /socialmedias [get]
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

// DeleteSocialmedia godoc
// @Summary Delete socialmedia identified by the given Id
// @Description Delete the Socialmedia corresponding to the input
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Id path int true "ID of the socialmedia to be deleted"
// @Success 204 "No Content"
// @Router /socialmedias/{id} [delete]
func DeleteSocialmedia(ctx *gin.Context) {
	socialmediaId, _ := strconv.Atoi(ctx.Param("socialmediaId"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Socialmedia := models.Socialmedia{}
	err := db.Where("id = ?", socialmediaId).Delete(&Socialmedia).Error
	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Socialmedia with id %d has been successfully deleted", socialmediaId)

	ctx.JSON(http.StatusOK, "Deleted")
}
