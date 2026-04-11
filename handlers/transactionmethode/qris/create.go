package qris

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) Create(c echo.Context) error {
	var resp dto.CoQrisResponse

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

	// Parse file from request
	file, err := c.FormFile("image") // Ensure "image" matches the form-data field name
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	// Validate file size (2 MB limit)
	const maxFileSize = 2 * 1024 * 1024 // 2 MB
	if file.Size > maxFileSize {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "File size exceeds the 2MB limit"})
	}

	req := dto.CoQrisRequest{
		UserID:     userid,
		CreatedBy:  createdBy,
		File:       file,
		MerchantID: merchantId,
	}

	resp, err = h.serviceQris.Create(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}
