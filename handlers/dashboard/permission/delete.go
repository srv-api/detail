package permission

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Delete(c echo.Context) error {
	var req dto.DeletePermissionRequest
	deletedBy, ok := c.Get("DeletedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.DeletedBy = deletedBy

	// Ambil ID dari parameter dan konversi menjadi uint
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, fmt.Errorf("invalid ID format: %v", err)).Send(c)
	}

	req.ID = uint(idUint)

	data, err := b.servicePermission.Delete(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
