package subscribe

import (
	"context"
	"fmt"
	"strconv"
	"time"

	paypal "github.com/plutov/paypal/v4"
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
	util "github.com/srv-api/util/s"
)

func (r *subscribeRepository) CreatePaypalOrder(req dto.PaypalCreateRequest) (*paypal.Order, error) {
	orderIntent := paypal.OrderIntentCapture

	purchaseUnit := paypal.PurchaseUnitRequest{
		Amount: &paypal.PurchaseUnitAmount{
			Currency: req.Currency,
			Value:    req.Amount,
		},
	}

	appContext := &paypal.ApplicationContext{
		ReturnURL: "https://cashpay.my.id/paypal-success",
		CancelURL: "https://cashpay.my.id/cancel",
	}

	// 1. Buat order PayPal
	order, err := r.paypalClient.CreateOrder(
		context.Background(),
		orderIntent,
		[]paypal.PurchaseUnitRequest{purchaseUnit},
		nil,
		appContext,
	)
	if err != nil {
		return nil, err
	}

	// 2. Ambil approval URL untuk disimpan ke DB
	var approvalUrl string
	for _, link := range order.Links {
		if link.Rel == "approve" {
			approvalUrl = link.Href
			break
		}
	}
	floatAmount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %v", err)
	}
	grossAmount := int64(floatAmount * 100)

	// 3. Simpan order ke DB
	tx := entity.Subscribe{
		ID:              util.GenerateRandomString(),
		UserID:          req.UserID,
		MerchantID:      req.MerchantID,
		CreatedBy:       req.CreatedBy,
		OrderID:         order.ID,
		GrossAmount:     grossAmount,
		PaymentType:     "paypal",
		Status:          "PENDING",
		TransactionTime: time.Now(),
		Url:             approvalUrl,
	}

	if err := r.DB.Create(&tx).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *subscribeRepository) CapturePaypalOrder(orderID string) (*paypal.CaptureOrderResponse, error) {
	return r.paypalClient.CaptureOrder(context.Background(), orderID, paypal.CaptureOrderRequest{})
}

func (r *subscribeRepository) UpdateSubscribeStatus(orderID string, status string) error {
	return r.DB.Model(&entity.Subscribe{}).
		Where("order_id = ?", orderID).
		Update("status", status).Error
}
