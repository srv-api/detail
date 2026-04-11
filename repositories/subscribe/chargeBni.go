package subscribe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
	util "github.com/srv-api/util/s"
	"gorm.io/gorm"
)

func (r *subscribeRepository) ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error) {
	var existingTx entity.Subscribe
	err := r.DB.
		Where("user_id = ? AND status = ?", req.UserID, "pending").
		Order("created_at DESC").
		First(&existingTx).Error

	if err == nil {
		return nil, errors.New("There is still an active transaction. Please check your Payment history.")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Gagal mengecek transaksi sebelumnya")
	}

	payload := map[string]interface{}{
		"payment_type": "bank_transfer",
		"transaction_details": map[string]interface{}{
			"order_id":     time.Now().Format("20060102150405"),
			"gross_amount": req.GrossAmount,
		},
		"bank_transfer": map[string]interface{}{
			"bank": "bni",
		},
		"custom_expiry": map[string]interface{}{
			"order_time":      time.Now().Format("2006-01-02 15:04:05 -0700"),
			"expiry_duration": 4,
			"unit":            "minutes",
		},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("failed to encode payload")
	}

	httpReq, err := http.NewRequest("POST", dto.GetMidtransEndpoint(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, errors.New("failed to create HTTP request")
	}

	auth := base64.StdEncoding.EncodeToString([]byte(dto.GetMidtransServerKey() + ":"))
	httpReq.Header.Set("Authorization", "Basic "+auth)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.New("failed to contact Midtrans")
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed to read Midtrans response")
	}

	var parsed dto.VAResponse
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return nil, errors.New("invalid response from Midtrans")
	}
	// Utility untuk gabung VA dan bank
	var vaNumber, bank string
	if len(parsed.VANumbers) > 0 {
		vaNumber = parsed.VANumbers[0].VANumber
		bank = parsed.VANumbers[0].Bank
	}
	tx := entity.Subscribe{
		ID:              util.GenerateRandomString(),
		MerchantID:      req.MerchantID,
		UserID:          req.UserID,
		CreatedBy:       req.CreatedBy,
		OrderID:         parsed.OrderID,
		TransactionID:   parsed.TransactionID,
		GrossAmount:     req.GrossAmount,
		PaymentType:     parsed.PaymentType,
		Status:          parsed.TransactionStatus,
		VA:              vaNumber,
		Bank:            bank,
		TransactionTime: parseTime(parsed.TransactionTime),
		ExpiryTime:      parseTime(parsed.ExpiryTime),
	}

	if err := r.DB.Create(&tx).Error; err != nil {
		return nil, err
	}

	return &parsed, nil
}
