package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Customer godoc
// @ID create_customer
// @Router /customer [POST]
// @Summary Create Customer
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param Customer body models.CreateCustomer true "CreateCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCustomer(c *gin.Context) {
	var createCustomer models.CreateCustomer

	err := c.ShouldBindJSON(&createCustomer)
	if err != nil {
		h.handlerResponse(c, "create customer", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Customer().Create(context.Background(), &createCustomer)
	if err != nil {
		h.handlerResponse(c, "storage create customer", http.StatusInternalServerError, err.Error())
		return
	}

	customer, err := h.storages.Customer().GetById(context.Background(), &models.CustomerPrimaryKey{CustomerId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id customer", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create customer", http.StatusCreated, customer)
} 

// Get By ID Customer godoc
// @ID get_by_id_customer
// @Router /customer/{id} [GET]
// @Summary Get By ID Customer
// @Description Get By ID Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCustomer(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id customer", http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.storages.Customer().GetById(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id customer", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id customer", http.StatusOK, customer)
}

// Get List Customer godoc
// @ID get_list_customer
// @Router /customer [GET]
// @Summary Get List Customer
// @Description Get List Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCustomer(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list customer", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list customer", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Customer().GetList(context.Background(), &models.GetListCustomerRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list customer", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list customer", http.StatusOK, resp)
}

// Get Update Customer godoc
// @ID update_customer
// @Router /customer/{id} [PUT]
// @Summary Update Customer
// @Description Update Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Customer body models.UpdateCustomer true "UpdateCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCustomer(c *gin.Context) {
	var updateCustomer models.UpdateCustomer

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi update customer", http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&updateCustomer)
	if err != nil {
		h.handlerResponse(c, "Update custoemr", http.StatusBadRequest, err.Error())
		return 
	}
	updateCustomer.CustomerId = idInt

	rowsAffected, err := h.storages.Customer().Update(context.Background(), &updateCustomer)
	if err != nil {
		h.handlerResponse(c, "Storage update customer", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update customer", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Customer().GetById(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id customer", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update customer", http.StatusOK, resp)
}

// Delete Customer godoc
// @ID delete_customer
// @Router /customer/{id} [DELETE]
// @Summary Delete Customer
// @Description Delete Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Customer body models.CustomerPrimaryKey true "DeleteCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete customer", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Customer().Delete(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete customer", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete customer", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete customer", http.StatusNoContent, "Deleted Successfully")
}