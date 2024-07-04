package main

import (
	"fmt"
	"log"
	"net/http"
	"token/models"
	jt "token/token"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/register", Register)
	router.GET("/user", GetUser)

	router.Run(":5050")
}

func Register(c *gin.Context) {

	req := models.RegisterReq{}

	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, gin.H{
			"error": "invalid data inserted",
		})
		return
	}

	token := jt.GenToken(&models.TokenReq{
		UserId:   "1",
		UserName: req.UserName,
		Email:    req.Email,
	})

	c.IndentedJSON(http.StatusCreated, models.TokenRes{
		Token: *token,
	})

}

func GetUser(c *gin.Context) {
	token := c.GetHeader("authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "not authorized",
		})

		return
	}

	if err := jt.ExtaracClaims(token); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "invalid token",
		})

		return
	}

	fmt.Println("token : ", token)

	c.IndentedJSON(200, gin.H{
		"message": "success !!!",
	})

}
