package subscribe

import (
	"net/http"

	dto "github.com/srv-api/detail/dto"
	res "github.com/srv-api/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) ChargePermata(c echo.Context) error {
	var req dto.ChargeRequest
	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	DetailID, ok := c.Get("DetailID").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.DetailID = DetailID
	req.UserID = userid
	req.CreatedBy = createdBy

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	// Call the service to charge the BCA account
	response, err := h.serviceSubscribe.ChargePermata(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Return the response to the client
	return c.JSON(http.StatusOK, response)
}
