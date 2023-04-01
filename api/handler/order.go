package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param Order body models.CreateOrder true "CreateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {
	var createOrder models.CreateOrder

	err := c.ShouldBindJSON(&createOrder)
	if err != nil {
		h.handlerResponse(c, "create order", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Order().Create(context.Background(), &createOrder)
	if err != nil {
		h.handlerResponse(c, "storage create order", http.StatusInternalServerError, err.Error())
		return
	}

	staff, err := h.storages.Order().GetById(context.Background(), &models.OrderPrimaryKey{OrderId: id})
	if err != nil {
		h.handlerResponse(c, "storage get by id order", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create order", http.StatusCreated, staff)
} 

// Get By ID Order godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err get by id staff", http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.storages.Order().GetById(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id order", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get by id order", http.StatusOK, category)
}

// Get List Order godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListOrder(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Get list order", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Get list order", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Order().GetList(context.Background(), &models.GetListOrderRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "Storage get list order", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get list order", http.StatusOK, resp)
}

// Get Update Order godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Order body models.UpdateOrder true "UpdateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateOrder(c *gin.Context) {
	var updateOrder models.UpdateOrder

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi update order", http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBindJSON(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "Update order", http.StatusBadRequest, err.Error())
		return 
	}
	updateOrder.OrderId = idInt

	rowsAffected, err := h.storages.Order().Update(context.Background(), &updateOrder)
	if err != nil {
		h.handlerResponse(c, "Storage update order", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage update order", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Order().GetById(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage get by id order", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update order", http.StatusOK, resp)
}

// Delete Order godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Order body models.OrderPrimaryKey true "DeleteOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi delete order", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storages.Order().Delete(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "Storage delete order", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage delete order", http.StatusBadRequest, "no rows affected")
		return
	}

	h.handlerResponse(c, "Delete order", http.StatusNoContent, "Deleted Successfully")
}

// Create Order Item godoc
// @ID create_order_item
// @Router /order_item [POST]
// @Summary Create Order Item
// @Description Create Order Item
// @Tags Order
// @Accept json
// @Produce json
// @Param order_item body models.CreateOrderItem true "CreateOrderItemRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrderItem(c *gin.Context) {

	var createOrderItem models.CreateOrderItem

	err := c.ShouldBindJSON(&createOrderItem) 
	if err != nil {
		h.handlerResponse(c, "Create order item", http.StatusBadRequest, err.Error())
		return
	}

	err = h.storages.Order().CheckStock(context.Background(), &createOrderItem)
	if err != nil {
		h.handlerResponse(c, "Check stock", http.StatusBadRequest, err.Error())
		return
	}

	err = h.storages.Order().AddOrderItem(context.Background(), &createOrderItem)
	if err != nil {
		h.handlerResponse(c, "Storage order create item", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Create order item", http.StatusCreated, "Added Successfully")
}

// DELETE Order Item godoc
// @ID delete_order_item
// @Router /order_item/{id} [DELETE]
// @Summary Delete Order Item
// @Description Delete Order Item
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param item_id query string true "item_id"
// @Param orderItem body models.OrderItemPrimaryKey true "DeleteOrderItemRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrderItem(c *gin.Context) {

	id := c.Param("id")
	itemId := c.Query("item_id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "Atoi err delete order item", http.StatusBadRequest, "id incorrect")
		return
	}

	idItemInt, err := strconv.Atoi(itemId)
	if err != nil {
		h.handlerResponse(c, "Atoi err delete order item", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Order().RemoveOrderItem(context.Background(), &models.OrderItemPrimaryKey{OrderId: idInt, ItemId: idItemInt})
	if err != nil {
		h.handlerResponse(c, "Storage order delete", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "Storage order item delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "Delete order", http.StatusNoContent, "Deleted succesfully")
}