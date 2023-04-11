package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Category godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body models.CreateCategory true "CreateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCategory(c *gin.Context) {
	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		h.handlerResponse(c, "create category", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Category().Create(context.Background(), &createCategory)
	if err != nil {
		h.handlerResponse(c, "storage create category", http.StatusInternalServerError, err.Error())
		return
	}

	category, err := h.storages.Category().GetById(context.Background(), &models.CategoryPrimaryKey{CategoryId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id category", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create category", http.StatusCreated, category)
} 

// Get By ID Category godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCategory(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id category", http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.storages.Category().GetById(context.Background(), &models.CategoryPrimaryKey{CategoryId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id category", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id category", http.StatusOK, category)
}

// Get List Category godoc
// @ID get_list_category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCategory(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list category", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list category", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Category().GetList(context.Background(), &models.GetListCategoryRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list category", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list category", http.StatusOK, resp)
}

// Get Update Category godoc
// @ID update_category
// @Router /category/{id} [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Category body models.UpdateCategory true "UpdateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCategory(c *gin.Context) {
	var updateCategory models.UpdateCategory

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi update category", http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "Update category", http.StatusBadRequest, err.Error())
		return 
	}
	updateCategory.CategoryId = idInt

	rowsAffected, err := h.storages.Category().Update(context.Background(), &updateCategory)
	if err != nil {
		h.handlerResponse(c, "Storage update category", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update category", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Category().GetById(context.Background(), &models.CategoryPrimaryKey{CategoryId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id category", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update category", http.StatusOK, resp)
}



// Delete Category godoc
// @ID delete_category
// @Router /category/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Category body models.CategoryPrimaryKey true "DeleteCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete category", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Category().Delete(context.Background(), &models.CategoryPrimaryKey{CategoryId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete category", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete category", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete category", http.StatusNoContent, "Deleted Successfully")
}