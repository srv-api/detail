package userdetail

import (
	"errors"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Explore(c echo.Context) error {
	var req dto.UserDetailRequest
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	// Validasi
	if req.Latitude == 0 || req.Longitude == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("latitude and longitude required")).Send(c)
	}
	if req.Radius <= 0 {
		req.Radius = 50 // default radius
	}

	// Ambil user_id dari JWT
	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userId

	// Panggil service dengan request yang sudah diisi
	resp, err := b.serviceUserDetail.Explore(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
