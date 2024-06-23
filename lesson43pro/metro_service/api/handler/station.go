package handler

import (
	"fmt"
	"pro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateStation(ctx *gin.Context) {
	stn := models.CreateStation{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdStation(ctx *gin.Context){
	id := ctx.Param("id")
	if id == ""{
		ctx.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	stat,err := h.Station.GetById(id)
	if err  != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, stat)
}

func (h *Handler) UpdateStation(ctx *gin.Context) {
	station :=  models.CreateStation{}

	err := ctx.BindJSON(&station)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Station.Update(&station, ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
}

func (h *Handler) DeleteStation(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}
	
	err := h.Station.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")
	
}

func (h *Handler) GetAllStation(ctx *gin.Context) {
	allStation := models.CreateStation{}
	allStation.Name = ctx.Param("name")
	
	res, err := h.Station.GetAll(allStation)
	if err != nil{
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}