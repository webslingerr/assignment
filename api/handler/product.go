package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Product godoc
// @ID create_product
// @Router /product [POST]
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body models.CreateProduct true "CreateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {
	var createProduct models.CreateProduct

	err := c.ShouldBindJSON(&createProduct)
	if err != nil {
		h.handlerResponse(c, "create product", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Product().Create(context.Background(), &createProduct)
	if err != nil {
		h.handlerResponse(c, "storage create product", http.StatusInternalServerError, err.Error())
		return
	}

	product, err := h.storages.Product().GetById(context.Background(), &models.ProductPrimaryKey{ProductId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id product", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create product", http.StatusCreated, product)
}

// Get By ID Product godoc
// @ID get_by_id_product
// @Router /product/{id} [GET]
// @Summary Get By ID Product
// @Description Get By ID Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id product", http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.storages.Product().GetById(context.Background(), &models.ProductPrimaryKey{ProductId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id product", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id product", http.StatusOK, product)
}

// Get List Product godoc
// @ID get_list_product
// @Router /product [GET]
// @Summary Get List Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListProduct(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list product", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list product", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Product().GetList(context.Background(), &models.GetListProductRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list product", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list product", http.StatusOK, resp)
}

// Get Update Product godoc
// @ID update_product
// @Router /product/{id} [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Product body models.UpdateProduct true "UpdateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {
	var updateProduct models.UpdateProduct

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi update product", http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&updateProduct)
	if err != nil {
		h.handlerResponse(c, "Update product", http.StatusBadRequest, err.Error())
		return
	}
	updateProduct.ProductId = idInt

	rowsAffected, err := h.storages.Product().Update(context.Background(), &updateProduct)
	if err != nil {
		h.handlerResponse(c, "Storage update product", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update product", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Product().GetById(context.Background(), &models.ProductPrimaryKey{ProductId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id product", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update product", http.StatusOK, resp)
}

// Delete Product godoc
// @ID delete_product
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Product body models.ProductPrimaryKey true "DeleteProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete product", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Product().Delete(context.Background(), &models.ProductPrimaryKey{ProductId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete product", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete product", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete customer", http.StatusNoContent, "Deleted Successfully")
}
