package handler

import (
	"fmt"

	"my_pro/storage/postgres"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	user          *postgres.CrudUsersRepo
	problem       *postgres.CrudProblemsRepo
	solvedproblem *postgres.CrudSolvedProblemRepo
}

func NewHandler(user *postgres.CrudUsersRepo, problem *postgres.CrudProblemsRepo, solved *postgres.CrudSolvedProblemRepo) *Handler {
	return &Handler{user: user, problem: problem, solvedproblem: solved}
}

func (h *Handler) HandlerAPI() {

	router := gin.Default()

	userGroup := router.Group("/api/user")

	userGroup.POST("/createuser", h.CreateUser)
	userGroup.GET("/readusers", h.ReadAllUser)
	userGroup.PUT("updateuser", h.UpdateUser)
	userGroup.DELETE("deleteuser", h.DeleteUser)

	problemGroup := router.Group("/api/problem")

	problemGroup.POST("/createproblem",h.CreateProblem)
	problemGroup.GET("/readproblem",h.ReadAllProblem)
	problemGroup.PUT("/updateproblem",h.UpdateProblem)
	problemGroup.DELETE("/deleteproblem",h.DeleteProblem)

	solProblem := router.Group("api/solvedProblem")

	solProblem.POST("/createsolproblem",h.CreateSProblem)
	solProblem.GET("/readsolproblem",h.ReadAllSProblem)
	solProblem.PUT("/updatesolproblem",h.UpdateSProblem)
	solProblem.DELETE("/deletesolproblem",h.DeleteSProblem)
	
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
