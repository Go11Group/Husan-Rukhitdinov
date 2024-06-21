package handler

import (
	"fmt"
	"my_pro/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// USER
func (h *Handler) CreateUser(c *gin.Context) {

	user := model.Users{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("dfjdfm00000-----",err)	
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.user.CreateUsers(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	user := model.Users{}
	user.UserID = id

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.user.UpdateUsers(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "User updated successfully"})
}
func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	err := h.user.DeleteUsers(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "User Deleted successfully"})

}

func (h *Handler) ReadAllUser(c *gin.Context) {
	a, err := h.user.ReadUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, a)

}
