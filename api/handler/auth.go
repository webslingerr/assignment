package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_User
// @Router /register [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param User body models.CreateUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) RegisterUser(c *gin.Context) {
	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "create User", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.User().Create(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage create User", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User", http.StatusCreated, nil)
}

// Create User godoc
// @ID create_User
// @Router /login [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param User body models.LoginUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) LoginUser(c *gin.Context) {
	var loginUser models.LoginUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		h.handlerResponse(c, "login User", http.StatusBadRequest, err.Error())
		return
	}

	exists, err := h.storages.User().GetById(context.Background(), &loginUser)
	if err != nil {
		h.handlerResponse(c, "storage login User", http.StatusBadRequest, err.Error())
		return
	}

	if !exists {
		h.handlerResponse(c, "username or password invalid", http.StatusBadRequest, nil)
		return
	}

	if exists {
		tokenString, err := helper.MakeJWT(&loginUser)
		if err != nil {
			h.handlerResponse(c, "generate token", http.StatusBadRequest, err.Error())
			return
		}
		h.handlerResponse(c, "generated token", http.StatusOK, tokenString)
	}
}
