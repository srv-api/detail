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

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
		})
	}

	userID := c.Get("UserId").(string)
	req.UserID = userID

	// 🔥 LOG
	println("=== HANDLER LIKE ===")
	println("UserID from JWT:", userID)
	println("Req UserID:", req.UserID)
	println("Req TargetUserID:", req.TargetUserID)
	println("Req IsSuperLike:", req.IsSuperLike)

	res, err := h.Service.LikeUser(req)
	// ... sisanya sama
}
