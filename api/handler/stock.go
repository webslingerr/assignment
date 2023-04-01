package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Stock godoc
// @ID create_stock
// @Router /stock [POST]
// @Summary Create Stock
// @Description Create Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param stock body models.CreateStock true "CreateStockRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStock(c *gin.Context) {
	var createStock models.CreateStock

	err := c.ShouldBindJSON(&createStock) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "Create product", http.StatusBadRequest, err.Error())
		return
	}

	storeId, err := h.storages.Stock().Create(context.Background(), &createStock)
	if err != nil {
		h.handlerResponse(c, "Storage stock create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Stock().GetById(context.Background(), &models.StockPrimaryKey{StoreId: storeId})
	if err != nil {
		h.handlerResponse(c, "Storage stock get by id", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Create stock", http.StatusCreated, resp)
}

// Get By ID Stock godoc
// @ID get_by_id_stock
// @Router /stock/{id} [GET]
// @Summary Get By ID Stock
// @Description Get By ID Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStock(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id stock", http.StatusBadRequest, err.Error())
		return
	}

	stock, err := h.storages.Stock().GetById(context.Background(), &models.StockPrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id stock", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id stock", http.StatusOK, stock)
}

// Get List Stock godoc
// @ID get_list_stock
// @Router /stock [GET]
// @Summary Get List Stock
// @Description Get List Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStock(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list stock", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list stock", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Stock().GetList(context.Background(), &models.GetListStockRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list stock", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list stock", http.StatusOK, resp)
}

// Get Update Stock godoc
// @ID update_stock
// @Router /stock [PUT]
// @Summary Update Stock
// @Description Update Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param storeId query string true "storeId"
// @Param productId query string true "productId"
// @Param Stock body models.UpdateStock true "UpdateStockRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStock(c *gin.Context) {
	var updateStock models.UpdateStock

	storeId, err := strconv.Atoi(c.Query("storeId"))
	if err != nil {
		h.handlerResponse(c, "Atoi update stock store id", http.StatusBadRequest, "invalid offset")
		return
	}

	productId, err := strconv.Atoi(c.Query("productId"))
	if err != nil {
		h.handlerResponse(c, "Atoi update stock product id", http.StatusBadRequest, "invalid limit")
		return
	}

	err = c.ShouldBindJSON(&updateStock)
	if err != nil {
		h.handlerResponse(c, "Update staff", http.StatusBadRequest, err.Error())
		return
	}
	updateStock.StoreId = storeId
	updateStock.ProductId = productId

	rowsAffected, err := h.storages.Stock().Update(context.Background(), &updateStock)
	if err != nil {
		h.handlerResponse(c, "Storage update stock", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update stock", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Update stock", http.StatusOK, "Updated Successfully")
}

// DELETE Stock godoc
// @ID delete_stock
// @Router /stock/{id} [DELETE]
// @Summary Delete Stock
// @Description Delete Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param stock body models.StockPrimaryKey true "DeleteStockRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStock(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Storage stock delete", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Store().Delete(context.Background(), &models.StorePrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage store delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage store delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "Delete store", http.StatusNoContent, nil)
}