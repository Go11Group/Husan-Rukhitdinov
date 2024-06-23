package handler

import (
	"fmt"
	"pro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateTerminal(ctx *gin.Context) {
	stn := models.Terminal{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Terminal.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdTerminal(ctx *gin.Context){
	id := ctx.Param("id")
	if id == ""{
		ctx.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	ter,err := h.Terminal.GetById(id)
	if err  != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, ter)
}

func (h *Handler) UpdateTerminal(ctx *gin.Context) {
	var terminal models.Terminal

	err := ctx.BindJSON(&terminal)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Terminal.Update(&terminal, ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
}

func (h *Handler) DeleteTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}
	
	err := h.Terminal.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
	
}

func (h *Handler) GetAllTerminal(ctx *gin.Context) {
	allTerminal := models.Terminal{}
	allTerminal.StationId = ctx.Param("station_id")
	
	res, err := h.Terminal.GetAll(allTerminal)
	if err != nil{
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}