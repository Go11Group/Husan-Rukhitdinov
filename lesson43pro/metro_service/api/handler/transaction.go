package handler

import (
	"fmt"
	"pro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	stn := models.Transaction{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Transaction.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdTransaction(ctx *gin.Context){
	id := ctx.Param("id")
	if id == ""{
		ctx.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	tran,err := h.Transaction.GetById(id)
	if err  != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, tran)
}

func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	var transac models.Transaction

	err := ctx.BindJSON(&transac)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Transaction.Update(&transac, ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
}

func (h *Handler) DeleteTransction(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}
	
	err := h.Transaction.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
	
}

func (h *Handler) GetAllTransaction(ctx *gin.Context) {
	allTransaction := models.Transaction{}
	allTransaction.CardID = ctx.Param("card_id")
	allTransaction.Amount = ctx.Param("amount")
	allTransaction.TerminalId = ctx.Param("terminal_id")
	
	res, err := h.Transaction.GetAll(allTransaction)
	if err != nil{
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}