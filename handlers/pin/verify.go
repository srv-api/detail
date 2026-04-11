package pin

import (
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) VerifyPIN(c echo.Context) error {
	var req dto.VerifyPinRequest

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userid
	req.MerchantID = merchantId
	req.CreatedBy = createdBy

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err := h.servicePin.VerifyPIN(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
