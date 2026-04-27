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
	req.UserID = userID

	res, err := h.Service.LikeUser(req)
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
func (h *LikeHandler) Me(c echo.Context) error {
	var req dto.LikeRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}
	userID := c.Get("UserId").(string)
	req.UserID = userID

	medsos, err := h.Service.Me(req)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to get medsos")
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   medsos,
	})
}
