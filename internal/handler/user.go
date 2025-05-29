package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"responsible_employee/internal/model"
	"responsible_employee/internal/utils"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.SignUp(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.NewErrorResponse(c, http.StatusConflict, "Пользователь с таким никнеймом уже существует!")
			return
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Пользователь успешно зарегистрирован!"})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input model.SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.SignIn(input)
	if err != nil {
		if err.Error() == "Пользователя с таким никнеймом не существует!" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		} else if err.Error() == "Неверный пароль!" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsersSortedByPoints(c *gin.Context) {
	users, err := h.services.GetUsersSortedByPoints()

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, users)
}
