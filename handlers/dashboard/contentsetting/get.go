package contentsetting

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	resp, err := b.serviceContentSetting.Get(dto.ContentSettingRequest{})
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	return res.SuccessResponse(resp).Send(c)
}
