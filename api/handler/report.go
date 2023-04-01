package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Send Product godoc
// @ID send_product
// @Router /report/send_product [PUT]
// @Summary Send Product
// @Description Send Product To Another Store
// @Tags Report
// @Accept json
// @Produce json
// @Param report body models.SendProduct true "SendProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) SendProductToStore(c *gin.Context) {
	var sendProduct models.SendProduct

	err := c.ShouldBindJSON(&sendProduct)
	if err != nil {
		h.handlerResponse(c, "Bind Json error send product to store", http.StatusBadRequest, err.Error())
		return
	}

	err = h.storages.Report().SendProduct(context.Background(), &sendProduct)
	if err != nil {
		h.handlerResponse(c, "Storage report  send product", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get store by id", http.StatusOK, "Success")
}

// Get List Staff Report godoc
// @ID get_list_staff_report
// @Router /report/staff_report [GET]
// @Summary Get List Staff Report
// @Description Get List Staff Report
// @Tags Report
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStaffReport(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Staff report", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Staff report", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Report().StaffReport(context.Background(), &models.StaffListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})

	if err != nil {
		h.handlerResponse(c, "Storage staff report", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Staff report", http.StatusOK, resp)
}

// Total Sum Order godoc
// @ID total_sum_order
// @Router /report/total_sum [GET]
// @Summary Total Sum Order
// @Description Total Sum Order
// @Tags Report
// @Accept json
// @Produce json
// @Param order_id query string true "order_id"
// @Param promocode_name query string false "promocode_name"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) OrderTotalSum(c *gin.Context) {
	var orderSum models.OrderTotalSum

	orderId, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		h.handlerResponse(c, "Atoi error order total sum", http.StatusBadRequest, err.Error())
		return
	}

	orderSum.OrderId = orderId
	orderSum.PromocodeName = c.Query("promocode_name")

	totalSum, err := h.storages.Report().OrderTotalSum(context.Background(), &orderSum)
	if err != nil {
		h.handlerResponse(c, "Storage order total sum", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Hello World")

	h.handlerResponse(c, "Order total sum", http.StatusOK, totalSum)
}
