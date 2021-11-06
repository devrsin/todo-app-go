package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sinyakovskiy/todo-app"
	"net/http"
)

//register
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

//auth
func (h *Handler) signIn(c *gin.Context) {

}
