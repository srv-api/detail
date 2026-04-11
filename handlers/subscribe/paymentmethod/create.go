package paymentmethod

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) Create(c echo.Context) error {
	userID, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	merchantID, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	paymentMethod := c.FormValue("payment_method")
	statusStr := c.FormValue("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	// file
	file, err := c.FormFile("image")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	src, err := file.Open()
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	defer src.Close()

	filename := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), file.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
	}

	req := dto.PaymentMethodRequest{
		UserID:        userID,
		MerchantID:    merchantID,
		CreatedBy:     createdBy,
		PaymentMethod: paymentMethod,
		Status:        status,
		FileName:      file.Filename,
		FilePath:      filename,
	}

	resp, err := h.servicePayment.Create(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
