package pin

import (
	"errors"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/detail/dto"
	res "github.com/srv-api/util/s/response"
	"gorm.io/gorm"
)

func (h *domainHandler) GetPinStatus(c echo.Context) error {
	var req dto.PinRequest

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	DetailID, ok := c.Get("DetailId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userid
	req.DetailID = DetailID
	req.CreatedBy = createdBy

	status, err := h.servicePin.GetPinStatus(req)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}

	return res.SuccessResponse(map[string]bool{
		"is_pin_enabled": status.IsPinEnabled,
	}).Send(c)
}
