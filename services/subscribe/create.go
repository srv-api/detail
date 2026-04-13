package subscribe

import (
	"fmt"
	"os"

	"github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/snap"
	"github.com/srv-api/merchant/constant"

	dto "github.com/srv-api/merchant/dto"
	util "github.com/srv-api/util/s"
)

func (s *subscribeService) Create(req dto.SubscribeRequest) (dto.SubscribeResponse, error) {
	orderID := util.GenerateRandomString()

	create := dto.SubscribeRequest{
		ID:          util.GenerateRandomString(),
		UserID:      req.UserID,
		CreatedBy:   req.CreatedBy,
		OrderID:     orderID,
		GrossAmount: req.GrossAmount,
		Status:      constant.StatusInitiated,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.SubscribeResponse{}, err
	}

	// Init Snap client
	midtransClient := snap.Client{}
	midtransClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// Mapping dari user input ke SnapPaymentType
	paymentType := mapPaymentType(req.PaymentType)

	// Build Snap Request
	transactionReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  create.OrderID,
			GrossAmt: int64(create.GrossAmount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: create.CreatedBy,
		},
		EnabledPayments: []snap.SnapPaymentType{paymentType}, // This is the correct type

	}

	// Create Snap Transaction

	snapTokenResp, err := midtransClient.CreateTransaction(transactionReq)
	if snapTokenResp == nil || snapTokenResp.Token == "" {
		_ = s.Repo.UpdateStatus(orderID, constant.StatusFailed)
		return dto.SubscribeResponse{}, nil
	}

	_ = s.Repo.UpdateStatus(orderID, constant.StatusPending)

	response := dto.SubscribeResponse{
		ID:          created.ID,
		OrderID:     orderID,
		GrossAmount: create.GrossAmount,
		Token:       snapTokenResp.Token,
		UserID:      created.UserID,
		CreatedBy:   created.CreatedBy,
		Status:      constant.StatusPending,
		RedirectURL: snapTokenResp.RedirectURL,
	}

	return response, nil
}

func (s *subscribeService) UpdateStatus(orderID string, transactionStatus string) error {
	var status string

	switch transactionStatus {
	case "settlement", "capture":
		status = "paid"
	case "pending":
		status = "pending"
	case "deny", "cancel", "expire":
		status = "failed"
	case "refund":
		status = "refunded"
	case "partial_refund":
		status = "partial_refunded"
	default:
		return fmt.Errorf("unknown transaction status: %s", transactionStatus)
	}

	if err := s.Repo.UpdateStatus(orderID, status); err != nil {
		return err
	}

	if transactionStatus == "settlement" {
		if err := s.Repo.UpdateUserVerified(orderID); err != nil {
			return err
		}
	}

	return nil
}

func mapPaymentType(userInput string) snap.SnapPaymentType {
	switch userInput {

	case "gopay":
		return snap.PaymentTypeGopay
	case "bca_va":
		return snap.PaymentTypeBCAVA
	case "bni_va":
		return snap.PaymentTypeBNIVA
	case "permata_va":
		return snap.PaymentTypePermataVA
	case "bank_transfer":
		return snap.PaymentTypeBankTransfer
	case "alfamart":
		return snap.PaymentTypeAlfamart
	case "indomaret":
		return snap.PaymentTypeIndomaret
	default:
		return snap.PaymentTypeGopay // default QRIS
	}
}
