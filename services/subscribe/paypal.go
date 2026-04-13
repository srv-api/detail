package subscribe

import (
	"errors"
	"strings"

	"github.com/srv-api/merchant/dto"
)

// func (s *subscribeService) CreatePaypalOrder(req dto.PaypalCreateRequest) (*dto.PaypalOrderResponse, error) {
// 	order, err := s.Repo.CreatePaypalOrder(req.Amount, req.Currency)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ambil approval URL dari links
// 	var approvalLink string
// 	for _, link := range order.Links {
// 		if link.Rel == "approve" {
// 			approvalLink = link.Href
// 			break
// 		}
// 	}

// 	return &dto.PaypalOrderResponse{
// 		ID:     order.ID,
// 		Status: order.Status,
// 		Link:   approvalLink,
// 	}, nil
// }

func (s *subscribeService) CreatePaypalOrder(req dto.PaypalCreateRequest) (*dto.PaypalOrderResponse, error) {
	if req.Amount == "" {
		return nil, errors.New("missing required fields: amount")
	}

	result, err := s.Repo.CreatePaypalOrder(req) // *paypal.Order
	if err != nil {
		return nil, err
	}

	var approvalLink string
	for _, link := range result.Links {
		if link.Rel == "approve" {
			approvalLink = link.Href
			break
		}
	}

	return &dto.PaypalOrderResponse{
		ID:     result.ID,
		Status: result.Status,
		Link:   approvalLink,
	}, nil
}

func (s *subscribeService) CapturePaypalOrder(orderID string) (*dto.PaypalCaptureResponse, error) {
	capture, err := s.Repo.CapturePaypalOrder(orderID)
	if err != nil {
		return nil, err
	}

	// Jika berhasil dan statusnya COMPLETED, update status di DB
	if strings.ToUpper(capture.Status) == "COMPLETED" {
		if err := s.Repo.UpdateSubscribeStatus(orderID, "settlement"); err != nil {
			return nil, err
		}
	}

	// 2. Update akun user (perpanjang 30 hari)
	if err := s.Repo.UpdateUserVerified(orderID); err != nil {
		return nil, err
	}

	return &dto.PaypalCaptureResponse{
		Status: capture.Status,
	}, nil
}
