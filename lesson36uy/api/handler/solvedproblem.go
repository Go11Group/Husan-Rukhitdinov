package handler

import (
	"my_pro/model"

	"github.com/gin-gonic/gin"
)

// SOLVED_PROBLEM
func (h *Handler) CreateSProblem(c *gin.Context) {

	solveProblem := model.SolvedProblem{}

	err := h.solvedproblem.CreateSolvedProblems(solveProblem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, solveProblem)
}

func (h *Handler) UpdateSProblem(c *gin.Context) {

	id := c.Param("id")
	solProblem := model.SolvedProblem{}
	if err := c.ShouldBindJSON(&solProblem); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.solvedproblem.UpdateSolvedProblems(solProblem, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "SolvedProblem update successfully!"})
}

func (h *Handler) DeleteSProblem(c *gin.Context) {

	id := c.Param("id")
	err := h.solvedproblem.DeleteSolvedProblems(id)

	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, gin.H{"message": "SolvedProblem delete successfully!"})
}

func (h *Handler) ReadAllSProblem(c *gin.Context) {
	err, res := h.solvedproblem.ReadSolvedProblems()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)

}
