package product

import (
	"errors"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) BulkEdit(c echo.Context) error {
	var req dto.BulkEditRequest

	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	UpdatedBy, ok := c.Get("UpdatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, errors.New("updated_by not found")).Send(c)
	}
	req.UpdatedBy = UpdatedBy

	if len(req.Items) == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("no items to update")).Send(c)
	}

	data, err := b.serviceProduct.BulkEdit(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
