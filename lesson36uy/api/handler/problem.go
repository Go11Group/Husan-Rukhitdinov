package handler

import (
	"my_pro/model"

	"github.com/gin-gonic/gin"
)

// PROBLEM
func (h *Handler) CreateProblem(c *gin.Context) {

	var problem model.Problem
	err := h.problem.CreateProblems(problem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, problem)

}

func (h *Handler) UpdateProblem(c *gin.Context) {

	id := c.Param("id")
	problem := model.Problem{}
	problem.ProblemID = id

	if err := c.ShouldBindJSON(&problem); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.problem.UpdateProblems(problem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "Problem update successfully", "problem": problem})
}

func (h *Handler) DeleteProblem(c *gin.Context) {

	id := c.Param("id")
	err := h.problem.DeleteProblems(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "Problem Deleted successfully"})

}

func (h *Handler) ReadAllProblem(c *gin.Context) {
	a, err := h.problem.ReadProblems()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, a)

}
