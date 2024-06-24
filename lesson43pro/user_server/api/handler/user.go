package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"user_pro/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	stn := models.User{}

	err := c.ShouldBindJSON(&stn)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.User.Create(&stn)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	c.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	stat, err := h.User.GetById(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, stat)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var card models.User

	err := c.BindJSON(&card)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.User.Update(&card, c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(200, "SUCCESS")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}

	err := h.User.Delete(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(200, "SUCCESS")

}

func (h *Handler) GetAllUser(c *gin.Context) {
	allUser := models.User{}
	allUser.Name = c.Query("name")
	allUser.Phone = c.Query("phone")
	ages := c.Query("age")
	if len(ages) == 0 {
		ages = "0"
	}
	age, err := strconv.Atoi(ages)
	allUser.Age = age
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.User.GetAll(allUser)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
