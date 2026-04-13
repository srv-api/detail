package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	Get(req dto.PaymentMethodRequest) ([]dto.PaymentMethodResponse, error)
	GetById(req dto.GetByIdPaymentRequest) (*dto.PaymentMethodResponse, error)
	Delete(req dto.DeletePaymentRequest) (dto.DeletePaymentResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.PaymentMethodUpdateRequest) (dto.PaymentMethodUpdateResponse, error)
	SaveImage(img entity.UploadedPayment) error
	GetPicture(req dto.GetPaymentploadRequest) (*dto.GetPaymentUploadResponse, error)
}

type paymentmethodRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) DomainRepository {
	return &paymentmethodRepository{
		DB: DB,
	}
}
