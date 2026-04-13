package history

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *historyRepository) GetById(req dto.GetHistory) (*dto.VAResponse, error) {
	tr := entity.Subscribe{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.VAResponse{
		OrderID:           tr.OrderID,
		TransactionID:     tr.TransactionID,
		TransactionStatus: tr.Status,
		PaymentType:       tr.PaymentType,
		TransactionTime:   tr.TransactionTime.Format("2006-01-02 15:04:05"),
		VANumbers: []struct {
			Bank     string `json:"bank"`
			VANumber string `json:"va_number"`
		}{
			{
				Bank:     tr.Bank,
				VANumber: tr.VA,
			},
		},
		ExpiryTime: tr.ExpiryTime.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
