package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Brand godoc
// @ID create_brand
// @Router /brand [POST]
// @Summary Create Brand
// @Description Create Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param Brand body models.CreateBrand true "CreateBrandRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateBrand(c *gin.Context) {
	var createBrand models.CreateBrand

	err := c.ShouldBindJSON(&createBrand)
	if err != nil {
		h.handlerResponse(c, "create brand", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Brand().Create(context.Background(), &createBrand)
	if err != nil {
		h.handlerResponse(c, "storage create brand", http.StatusInternalServerError, err.Error())
		return
	}

	brand, err := h.storages.Brand().GetById(context.Background(), &models.BrandPrimaryKey{BrandId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id brand", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create brand", http.StatusCreated, brand)
}

// Get By ID Brand godoc
// @ID get_by_id_brand
// @Router /brand/{id} [GET]
// @Summary Get By ID Brand
// @Description Get By ID Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"Order
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdBrand(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id brand", http.StatusBadRequest, err.Error())
		return
	}

	brand, err := h.storages.Brand().GetById(context.Background(), &models.BrandPrimaryKey{BrandId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id brand", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id brand", http.StatusOK, brand)
}

// Get List Brand godoc
// @ID get_list_brand
// @Router /brand [GET]
// @Summary Get List Brand
// @Description Get List Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListBrand(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list brand", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list brand", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Brand().GetList(context.Background(), &models.GetListBrandRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list brand", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list brand", http.StatusOK, resp)
}

// Get Update Brand godoc
// @ID update_brand
// @Router /brand/{id} [PUT]
// @Summary Update Brand
// @Description Update Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Brand body models.UpdateBrand true "UpdateBrandRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateBrand(c *gin.Context) {
	var updateBrand models.UpdateBrand

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi update brand", http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&updateBrand)
	if err != nil {
		h.handlerResponse(c, "Update brand", http.StatusBadRequest, err.Error())
		return
	}
	updateBrand.BrandId = idInt

	rowsAffected, err := h.storages.Brand().Update(context.Background(), &updateBrand)
	if err != nil {
		h.handlerResponse(c, "Storage update brand", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update brand", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Brand().GetById(context.Background(), &models.BrandPrimaryKey{BrandId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id brand", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update brand", http.StatusOK, resp)
}

// Delete Brand godoc
// @ID delete_brand
// @Router /brand/{id} [DELETE]
// @Summary Delete Brand
// @Description Delete Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Brand body models.BrandPrimaryKey true "DeleteBrandRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete brand", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Brand().Delete(context.Background(), &models.BrandPrimaryKey{BrandId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete brand", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete brand", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete brand", http.StatusNoContent, "Deleted Successfully")
}
