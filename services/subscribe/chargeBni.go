package subscribe

import (
	"errors"

	dto "github.com/srv-api/merchant/dto"
)

func (s *subscribeService) ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error) {
	if req.GrossAmount <= 0 {
		return nil, errors.New("missing required fields: amount")
	}

	resp, err := s.Repo.ChargeBni(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}
