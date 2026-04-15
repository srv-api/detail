package handler

import (
	"net/http"

	service "github.com/srv-api/detail/services/match"

	"github.com/labstack/echo/v4"
)

type MatchHandler struct {
	Service service.MatchService
}

func NewMatchHandler(s service.MatchService) *MatchHandler {
	return &MatchHandler{Service: s}
}

func (h *MatchHandler) GetMatches(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	if userIDInterface == nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	userID := userIDInterface.(string)

	results, err := h.Service.GetMatches(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get matches",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   results,
	})
}
