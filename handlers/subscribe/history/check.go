package history

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) CheckExpire(c echo.Context) error {
	orderID := c.Param("order_id")
	if orderID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "order_id is required"})
	}

	if err := h.serviceHistory.ExpireTransaction(orderID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction expired"})

}
