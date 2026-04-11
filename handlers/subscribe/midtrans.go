package subscribe

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/srv-api/merchant/constant"
	res "github.com/srv-api/util/s/response"
)

type MidtransNotification struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	StatusCode        string `json:"status_code"`
	GrossAmount       string `json:"gross_amount"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

func (h *domainHandler) MidtransCallback(c echo.Context) error {
	var notificationPayload MidtransNotification

	if err := c.Bind(&notificationPayload); err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	internalStatus, err := mapMidtransStatusToInternal(notificationPayload.TransactionStatus)
	if err != nil {
		log.Printf("MidtransCallback: unknown transaction status received: %s", notificationPayload.TransactionStatus)
		return res.ErrorResponse(err).Send(c)
	}

	if err := h.serviceSubscribe.UpdateStatus(notificationPayload.OrderID, internalStatus); err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(map[string]string{
		"message": "Payment status updated successfully",
		"orderId": notificationPayload.OrderID,
		"status":  internalStatus,
	}).Send(c)
}

func mapMidtransStatusToInternal(midtransStatus string) (string, error) {
	switch midtransStatus {
	case "capture", "settlement":
		return constant.StatusPaid, nil
	case "deny", "cancel", "expire":
		return constant.StatusFailed, nil
	case "pending":
		return constant.StatusPending, nil
	default:
		return "", fmt.Errorf("unknown midtrans status: %s", midtransStatus)
	}
}
