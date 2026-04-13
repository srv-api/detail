package permission

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) GetById(c echo.Context) error {
	var req dto.GetPermissionByIdRequest

	// idStr := c.Param("id")

	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, fmt.Errorf("invalid ID format: %v", err)).Send(c)
	}
	req.ID = uint(idUint)

	transaction, err := b.servicePermission.GetById(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return res.SuccessResponse(transaction).Send(c)

}
