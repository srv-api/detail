package subscribe

import (
	"encoding/json"

	"github.com/srv-api/merchant/dto"
)

func (s *subscribeService) CancelPay(req dto.GetorderID) (map[string]interface{}, int, error) {
	body, status, err := s.Repo.CancelPay(req)
	if err != nil {
		return nil, status, err
	}

	var midtransResp dto.MidtransCancelResponse
	if err := json.Unmarshal(body, &midtransResp); err != nil {
		return nil, status, err
	}

	// update DB
	if err := s.Repo.UpdateSubscribeByOrderID(midtransResp); err != nil {
		return nil, status, err
	}

	// return JSON as native map
	var result map[string]interface{}
	_ = json.Unmarshal(body, &result)
	return result, status, nil
}
