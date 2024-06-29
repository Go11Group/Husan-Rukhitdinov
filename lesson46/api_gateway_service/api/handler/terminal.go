package handler

import (
	"net/http"

	"api_gateway_service/genproto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTerminal(ctx *gin.Context) {
	stn := &genproto.Terminal{}

	resp, err := h.Metro.Create(ctx, stn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetByIdTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.String(http.StatusBadRequest, "ID xato kiritildi!")
	}
	ter, err := h.Metro.GetById(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, ter)
}

func (h *Handler) UpdateTerminal(ctx *gin.Context) {
	terminal := &genproto.Terminal

	err := ctx.BindJSON(&terminal)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	err = h.Metro.Update(&terminal, ctx.Param("id"))
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

	err := h.Metro.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(200, "SUCCESS")

}

func (h *Handler) GetAllTerminal(ctx *gin.Context) {
	req := &genproto.Terminal{}

	res, err := h.Metro.GetAll(ctx, req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}
