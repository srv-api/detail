package purchase

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/detail/dto"
)

func (h *purchaseHandler) VerifyPurchase(c echo.Context) error {
	var req dto.VerifyPurchaseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.PurchaseResponse{
			Success: false,
			Message: "Invalid request format",
		})
	}

	// Validasi request
	if req.UserID == "" || req.PurchaseToken == "" || req.ProductID == "" {
		return c.JSON(http.StatusBadRequest, dto.PurchaseResponse{
			Success: false,
			Message: "Missing required fields",
		})
	}

	// Proses purchase
	response, err := h.servicePurchase.ProcessPurchase(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.PurchaseResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

// Endpoint tambahan untuk cek status premium
func (h *purchaseHandler) GetUserPremiumStatus(c echo.Context) error {
	userID := c.Param("userId")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User ID required",
		})
	}

	// Call service method to get user status
	// status, err := h.purchaseService.GetUserPremiumStatus(userID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"userId": userID,
		// "isPremium": status,
	})
}
