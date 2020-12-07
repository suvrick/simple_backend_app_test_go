package controllers

/*
	CONTROLLER CRUD USER
*/

import (
	"kiss_web/database"
	"kiss_web/models"
	"kiss_web/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var repo repository.IUserRepository

// Init ...
func Init(router *gin.Engine) {

	_db := database.GetDB()

	if _db == nil {
		log.Panicln("Db pointer is nil!")
		return
	}

	repo = repository.NewUserRepository(_db)

	router.GET("user/all", getAllHandler)
	router.GET("user/get", getByIDHandler)
	router.GET("user/delete", deleteHandler)
	router.POST("user/update", updateHandler)
	router.POST("user/create", createHandler)
}

func getAllHandler(c *gin.Context) {
	users, err := repo.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
			"users":  users,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"users":  users,
	})
}

func updateHandler(c *gin.Context) {

	user := &models.User{}

	//Заполняем структуру из формы
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Invalid parameter",
			"user":   user,
		})
		return
	}

	//Вызов метода репозитория!!!
	err := repo.UpdateUser(user)

	//Отправляем ошибку
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
			"user":   user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func getByIDHandler(c *gin.Context) {

	strID := c.Query("id")

	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad parameter",
			"user":   models.User{},
		})
		return
	}

	user, err := repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": err.Error(),
			"user":   user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func createHandler(c *gin.Context) {

	login := c.PostForm("login")
	password := c.PostForm("password")

	if len(login) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Invalid parameter",
			"user":   models.User{},
		})
		return
	}

	user := &models.User{
		Login:    login,
		Password: password,
	}

	err := repo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
			"user":   user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"user":   user,
	})
}

func deleteHandler(c *gin.Context) {

	strID := c.Query("id")

	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad parameter",
		})
		return
	}

	err = repo.RemoveUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
