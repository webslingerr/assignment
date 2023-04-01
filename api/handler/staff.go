package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Staff godoc
// @ID create_staff
// @Router /staff [POST]
// @Summary Create Staff
// @Description Create Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param Staff body models.CreateStaff true "CreateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStaff(c *gin.Context) {
	var createStaff models.CreateStaff

	err := c.ShouldBindJSON(&createStaff)
	if err != nil {
		h.handlerResponse(c, "create staff", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Staff().Create(context.Background(), &createStaff)
	if err != nil {
		h.handlerResponse(c, "storage create staff", http.StatusInternalServerError, err.Error())
		return
	}

	staff, err := h.storages.Staff().GetById(context.Background(), &models.StaffPrimaryKey{StaffId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id staff", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff", http.StatusCreated, staff)
} 

// Get By ID Staff godoc
// @ID get_by_id_staff
// @Router /staff/{id} [GET]
// @Summary Get By ID Staff
// @Description Get By ID Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStaff(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id staff", http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.storages.Staff().GetById(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id staff", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id staff", http.StatusOK, category)
}

// Get List Staff godoc
// @ID get_list_staff
// @Router /staff [GET]
// @Summary Get List Staff
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStaff(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list staff", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list staff", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Staff().GetList(context.Background(), &models.GetListStaffRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list staff", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list staff", http.StatusOK, resp)
}

// Get Update Staff godoc
// @ID update_staff
// @Router /staff/{id} [PUT]
// @Summary Update Staff
// @Description Update Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Staff body models.UpdateStaff true "UpdateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStaff(c *gin.Context) {
	var updateStaff models.UpdateStaff

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi update staff", http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&updateStaff)
	if err != nil {
		h.handlerResponse(c, "Update staff", http.StatusBadRequest, err.Error())
		return 
	}
	updateStaff.StaffId = idInt

	rowsAffected, err := h.storages.Staff().Update(context.Background(), &updateStaff)
	if err != nil {
		h.handlerResponse(c, "Storage update staff", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update staff", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Staff().GetById(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id staff", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update staff", http.StatusOK, resp)
}

// Delete Staff godoc
// @ID delete_staff
// @Router /staff/{id} [DELETE]
// @Summary Delete Staff
// @Description Delete Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Staff body models.StaffPrimaryKey true "DeleteStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete staff", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Staff().Delete(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete staff", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete staff", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete staff", http.StatusNoContent, "Deleted Successfully")
}