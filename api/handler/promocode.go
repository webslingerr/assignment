package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Promocode godoc
// @ID create_promocode
// @Router /promocode [POST]
// @Summary Create Promocode
// @Description Create Promocode
// @Tags Promocode
// @Accept json
// @Produce json
// @Param Promocode body models.CreatePromocode true "CreatePromocodeRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreatePromocode(c *gin.Context) {
	var createPromocode models.CreatePromocode

	err := c.ShouldBindJSON(&createPromocode)
	if err != nil {
		h.handlerResponse(c, "create promocode", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Promocode().Create(context.Background(), &createPromocode)
	if err != nil {
		h.handlerResponse(c, "storage create promocode", http.StatusInternalServerError, err.Error())
		return
	}

	promocode, err := h.storages.Promocode().GetById(context.Background(), &models.PromocodePrimaryKey{PromocodeId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id promocode", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create promocode", http.StatusCreated, promocode)
} 

// Get By ID Promocode godoc
// @ID get_by_id_promocode
// @Router /promocode/{id} [GET]
// @Summary Get By ID Promocode
// @Description Get By ID Promocode
// @Tags Promocode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdPromocode(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id promocode", http.StatusBadRequest, err.Error())
		return
	}

	promocode, err := h.storages.Promocode().GetById(context.Background(), &models.PromocodePrimaryKey{PromocodeId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id promocode", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id promocode", http.StatusOK, promocode)
}

// Get List Promocode godoc
// @ID get_list_promocode
// @Router /promocode [GET]
// @Summary Get List Promocode
// @Description Get List Promocode
// @Tags Promocode
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListPromocode(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list promocode", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list promocode", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Promocode().GetList(context.Background(), &models.GetListPromocodeRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list promocode", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list promocode", http.StatusOK, resp)
}

// Delete Promocode godoc
// @ID delete_promocode
// @Router /promocode/{id} [DELETE]
// @Summary Delete Promocode
// @Description Delete Promocode
// @Tags Promocode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Promocode body models.PromocodePrimaryKey true "DeletePromocodeRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeletePromocode(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete promocode", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Promocode().Delete(context.Background(), &models.PromocodePrimaryKey{PromocodeId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete promocode", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete promocode", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete promocode", http.StatusNoContent, "Deleted Successfully")
}