package role

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/helpers"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.RoleUserRequest
	paginationDTO := helpers.GeneratePaginationRequest(c)

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	DetailID, ok := c.Get("DetailID").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	paginationDTO.DetailID = DetailID
	paginationDTO.UserID = userid

	if err := c.Bind(&paginationDTO); err != nil {
		return c.JSON(400, "Invalid request")
	}

	products, err := b.serviceRole.Get(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return c.JSON(http.StatusOK, products)
}
