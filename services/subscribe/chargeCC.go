package subscribe

import (
	"errors"

	dto "github.com/srv-api/merchant/dto"
)

func (s *subscribeService) CardPayment(req dto.CreditCardChargeRequest) (*dto.TokenizeResponse, error) {
	resp, err := s.Repo.CardPayment(req)
	if err != nil {
		return nil, err
	}

	if resp.Status != "success" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}
