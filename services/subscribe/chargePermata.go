package subscribe

import (
	"errors"

	dto "github.com/srv-api/merchant/dto"
)

func (s *subscribeService) ChargePermata(req dto.ChargeRequest) (*dto.VAPermataResponse, error) {
	if req.GrossAmount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	resp, err := s.Repo.ChargePermata(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}
