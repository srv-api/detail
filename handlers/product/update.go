package product

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Update(c echo.Context) error {
	var req dto.ProductUpdateRequest
	var resp dto.ProductUpdateResponse

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	updatedBy, ok := c.Get("UpdatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ID = idUint
	req.UpdatedBy = updatedBy
	req.UserID = userid
	req.MerchantID = merchantId

	err = c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err = b.serviceProduct.Update(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}
