package contentsetting

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/detail/dto"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) Update(c echo.Context) error {
	var req dto.UpdateContentSettingRequest

	// ✅ Ambil ID dari query param
	id := c.QueryParam("id")
	if id == "" {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, errors.New("missing id")).Send(c)
	}

	// ✅ Ambil UserID, UpdatedBy, DetailIDdari context (biasanya dari JWT middleware)
	userID, ok := c.Get("UserId").(string)
	if !ok || userID == "" {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, errors.New("invalid user")).Send(c)
	}

	updatedBy, ok := c.Get("UpdatedBy").(string)
	if !ok || updatedBy == "" {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, errors.New("invalid updater")).Send(c)
	}

	DetailID, ok := c.Get("DetailId").(string)
	if !ok || DetailID == "" {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, errors.New("invalid merchant")).Send(c)
	}

	// ✅ Bind body JSON ke struct
	if err := c.Bind(&req); err != nil {
		return res.Response(c, http.StatusBadRequest, res.ResponseModel{
			Status:  false,
			Message: "invalid request body: " + err.Error(),
			Data:    nil,
		})
	}

	// ✅ Set field tambahan dari context
	req.ID = id
	req.UserID = userID
	req.DetailID = DetailID
	req.UpdatedBy = updatedBy

	// ✅ Panggil service
	result, err := h.serviceContentSetting.Update(req)
	if err != nil {
		return res.Response(c, http.StatusInternalServerError, res.ResponseModel{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// ✅ Return hasil sukses
	return res.SuccessResponse(result).Send(c)
}
