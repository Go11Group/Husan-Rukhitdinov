package handler

import (
	"fmt"
	"net/http"

	"api_gateway_service/genproto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStation(ctx *gin.Context) {
	stn := &genproto.Void{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Metro.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdStation(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	stat, err := h.Metro.GetById(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, stat)
}

func (h *Handler) UpdateStation(ctx *gin.Context) {
	station := &genproto.Void{}

	err := ctx.BindJSON(&station)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Metro.Update(&station, ctx.Param("id"))
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

	err := h.Metro.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")

}

func (h *Handler) GetAllStation(ctx *gin.Context) {
	req := &genproto.Void{}

	res, err := h.Metro.GetAll(ctx, req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}
