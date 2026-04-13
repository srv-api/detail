package roleuser

import (
	"errors"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) BulkDelete(c echo.Context) error {
	var req dto.BulkDeleteRequest

	// Bind data dari body JSON
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	deletedBy, ok := c.Get("DeletedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.DeletedBy = deletedBy

	// Validasi ID, pastikan tidak kosong
	if len(req.ID) == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("invalid id")).Send(c)
	}

	// Panggil service untuk menghapus produk
	data, err := b.serviceRoleUser.BulkDelete(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
