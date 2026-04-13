package paymentmethod

import (
	"fmt"
	"os"

	dto "github.com/srv-api/merchant/dto"
)

func (b *paymentmethodService) GetPicture(req dto.GetPaymentploadRequest) (*dto.GetPaymentUploadResponse, error) {
	// Ambil data dari repository
	transaction, err := b.Repo.GetPicture(req)
	if err != nil {
		return nil, err
	}

	// Pastikan path file benar
	filePath := "./" + transaction.FilePath // Tambahkan prefix untuk path lokal
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found")
	}

	transaction.FilePath = filePath
	return transaction, nil
}
