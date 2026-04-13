package subscribe

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) CheckTransactionStatus(c echo.Context) error {
	// Getting order_id from URL params
	orderID := c.Param("order_id")

	// Get UserID and CreatedBy from context (assumes these values are set during authentication)
	userID, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	// Creating the DTO for the request
	req := dto.CreateTransactionRequest{
		UserID:    userID,
		CreatedBy: createdBy,
		OrderID:   orderID,
	}

	// Call the service method to check transaction status
	result, err := h.serviceSubscribe.CheckTransactionStatus(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	// Return successful response with transaction status
	return res.SuccessResponse(result).Send(c)
}
