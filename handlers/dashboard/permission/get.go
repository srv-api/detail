package permission

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.RoleUserRequest

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userid

	products, err := b.servicePermission.Get(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return c.JSON(http.StatusOK, products)
}
