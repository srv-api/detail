package subscribe

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	dto "github.com/srv-api/merchant/dto"
)

func (s *subscribeService) ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error) {
	if req.GrossAmount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	resp, err := s.Repo.ChargeQris(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}

func (s *subscribeService) CheckTransactionStatus(request dto.CreateTransactionRequest) (dto.TransactionStatusResponse, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return dto.TransactionStatusResponse{}, fmt.Errorf("midtrans server key is not set")
	}

	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", request.OrderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.TransactionStatusResponse{}, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.TransactionStatusResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.TransactionStatusResponse{}, err
	}

	var result dto.TransactionStatusResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return dto.TransactionStatusResponse{}, err
	}

	// Update DB status
	if err := s.Repo.UpdateStatus(request.OrderID, result.TransactionStatus); err != nil {
		return dto.TransactionStatusResponse{}, err
	}

	if result.TransactionStatus == "settlement" {
		if err := s.Repo.UpdateUserVerified(request.OrderID); err != nil {
			return dto.TransactionStatusResponse{}, err
		}
	}

	return result, nil
}
