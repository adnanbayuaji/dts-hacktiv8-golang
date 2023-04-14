package controllers

import (
	"errors"
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"password": User.Password,
		"age":      User.Age,
	})

}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func CreateUser(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"password": User.Password,
		"age":      User.Age,
	})
}

func UpdateUser(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	updatedUser := models.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&updatedUser)
	} else {
		ctx.ShouldBind(&updatedUser)
	}

	user := models.User{}

	err := db.Model(&user).Where("id = ?", userID).Updates(models.User{Username: updatedUser.Username, Email: updatedUser.Email, Password: updatedUser.Password, Age: updatedUser.Age}).Error

	if err != nil {
		fmt.Println("Error updating User data:", err)
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func GetUser(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	user := models.User{}
	err := db.First(&user, "id = ?", userID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUsers(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	users := []models.User{}
	err := db.Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}

func DeleteUser(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))

	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	user := models.User{}
	err := db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("User with id %d has been successfully deleted", userID)

	ctx.JSON(http.StatusOK, "Deleted")
}
