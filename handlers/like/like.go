package handler

import (
	"net/http"

	"github.com/srv-api/detail/dto"
	service "github.com/srv-api/detail/services/like"

	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	Service service.LikeService
}

func NewLikeHandler(s service.LikeService) *LikeHandler {
	return &LikeHandler{Service: s}
}

func (h *LikeHandler) LikeUser(c echo.Context) error {
	var req dto.LikeRequest

	// bind request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}

	// ambil user dari JWT middleware
	userID := c.Get("UserId").(string)

	res, err := h.Service.LikeUser(userID, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": res.Message,
		"data":    res,
	})
}
