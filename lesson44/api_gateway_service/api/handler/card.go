package handler

import (
	"net/http"

	"api_gateway_service/genproto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCard(ctx *gin.Context) {
	stn := &genproto.Void{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.Metro.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *Handler) GetByIdCard(ctx *gin.Context) {
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

func (h *Handler) UpdateCard(ctx *gin.Context) {
	card := &genproto.Void{}

	err := ctx.BindJSON(&card)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Metro.Update(&card, ctx.Param("id"))
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

	err := h.Metro.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")

}

func (h *Handler) GetAll(ctx *gin.Context) {
	req := &genproto.Void{}

	resp, err := h.Metro.GetAllCard(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
