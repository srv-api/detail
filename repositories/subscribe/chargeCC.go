package subscribe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
	util "github.com/srv-api/util/s"
)

func (r *subscribeRepository) CardPayment(req dto.CreditCardChargeRequest) (*dto.TokenizeResponse, error) {
	// Siapkan data request
	requestBody := dto.CreditCardChargeRequest{
		PaymentType: "credit_card",
		TransactionDetails: dto.TransactionDetails{
			OrderID:     time.Now().Format("20060102150405"),
			GrossAmount: req.TransactionDetails.GrossAmount,
		},
		CreditCard: dto.CreditCardData{
			TokenID:        req.CreditCard.TokenID,
			Authentication: true,
		},
	}

	// Encode ke JSON
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Buat request HTTP
	httpReq, err := http.NewRequest("POST", dto.GetMidtransEndpoint(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set header yang wajib
	auth := base64.StdEncoding.EncodeToString([]byte(dto.GetMidtransServerKey() + ":"))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+auth)

	// Kirim request
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.New("failed to contact Midtrans")
	}
	defer resp.Body.Close()

	// === DEBUG: Baca raw response body ===
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read Midtrans response")
	}
	// Print raw response ke console/log supaya bisa dicek
	fmt.Println("Midtrans raw response:", string(respBody))

	// Reset resp.Body supaya bisa decode ulang JSON dari data yang sudah dibaca
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	// Parse JSON response ke struct
	var response dto.TokenizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Jika Midtrans kembalikan status gagal, kembalikan error
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, errors.New("Midtrans error: " + response.StatusMessage)
	}

	// Simpan transaksi ke DB
	tx := entity.CreditCard{
		ID:            util.GenerateRandomString(),
		UserID:        req.UserID,
		CreatedBy:     req.CreatedBy,
		GrossAmount:   req.TransactionDetails.GrossAmount,
		TokenID:       response.TokenID,
		TransactionID: response.TransactionID,
		Status:        response.Status,
	}

	if err := r.DB.Create(&tx).Error; err != nil {
		return nil, err
	}

	return &response, nil
}
