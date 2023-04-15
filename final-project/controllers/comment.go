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

// CreateComment godoc
// @Summary Post details for a given Id
// @Description Post details of Comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Router /comments/{photoId} [post]
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
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment.PhotoID = uint(photoId)

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

// UpdateComment godoc
// @Summary Update Comment identified by the given Id
// @Description Update the Comment corresponding to the input
// @Tags comments
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment to be updated"
// @Success 200 {object} models.Comment
// @Router /comments/{id}/{photoId} [patch]
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
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment.PhotoID = uint(photoId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

// GetAllComment godoc
// @Summary Get details
// @Description Get details of all Comment
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments [get]
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

// GetOneComment godoc
// @Summary Get details for a given Id
// @Description Get details of Comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /comments [get]
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

// DeleteComment godoc
// @Summary Delete comment identified by the given Id
// @Description Delete the Comment corresponding to the input
// @Tags comments
// @Accept json
// @Produce json
// @Param Id path int true "ID of the comment to be deleted"
// @Success 204 "No Content"
// @Router /comments/{id} [delete]
func DeleteComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Comment := models.Comment{}
	err := db.Where("id = ?", commentId).Delete(&Comment).Error
	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Comment with id %d has been successfully deleted", commentId)

	ctx.JSON(http.StatusOK, "Deleted")
}
