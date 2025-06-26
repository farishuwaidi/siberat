package handler

import (
	"Siberat/dto"
	"Siberat/entity"
	"Siberat/helper"
	"Siberat/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetAllUser(c *gin.Context){
	users, err := h.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			StatusCode:  http.StatusInternalServerError,
			Message: "failed to retrieve users",
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "List of users",
		Data: users,
	}))
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusBadRequest,
			Message: "invalid user id",
		}))
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusNotFound,
			Message: "User not found",
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "User retrieved successfully",
		Data: user,
	}))
}

func (h *UserHandler) UpdatedUser(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusBadRequest,
			Message: "invalid user id",
		}))
		return
	}
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusBadRequest,
			Message:  err.Error(),
		}))
		return
	}

	if err := h.service.UpdateUser(id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusInternalServerError,
			Message: "failed to update user",
		}))	
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "User Update successfuly",
	}))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusBadRequest,
			Message: "Invalid User ID",
		}))
		return
	}

	//ambil user login dari contex (set oleh middleware)
	loggedInUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusUnauthorized,
			Message: "Unauthorized",
		}))
		return
	}
	user := loggedInUser.(entity.User)

	//tidak dapat menghapus diri sendiri
	if user.ID == id {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusBadRequest,
			Message: "You can't delete your self",
		}))
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			StatusCode: http.StatusNotFound,
			Message: "User not found or already deleted",
		}))
		return
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "User Deleted Successfuly",
	}))
}