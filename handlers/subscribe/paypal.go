package subscribe

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/detail/dto"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) PayPal(c echo.Context) error {
	var req dto.PaypalCreateRequest
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
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	res, err := h.serviceSubscribe.CreatePaypalOrder(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *domainHandler) CapturePaypalOrder(c echo.Context) error {
	orderID := c.Param("order_id")
	if orderID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Order ID is required")
	}

	res, err := h.serviceSubscribe.CapturePaypalOrder(orderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
