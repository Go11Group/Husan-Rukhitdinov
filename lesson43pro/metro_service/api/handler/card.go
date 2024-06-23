package handler

import (
	"fmt"
	"net/http"
	"pro/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCard(ctx *gin.Context) {
	stn := models.Card{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Card.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdCard(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	stat, err := h.Card.GetById(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, stat)
}

func (h *Handler) UpdateCard(ctx *gin.Context) {
	var card models.Card

	err := ctx.BindJSON(&card)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Card.Update(&card, ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
}

func (h *Handler) DeleteCard(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}

	err := h.Card.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")

}

func (h *Handler) GetAllCard(ctx *gin.Context) {
	allCard := models.Card{}
	allCard.Number = ctx.Param("number")
	allCard.UserId = ctx.Param("user_id")
	
	res, err := h.Card.GetAll(allCard)
	if err != nil{
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}
