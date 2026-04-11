package product

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/product/entity"
)

func (b *productRepository) GetPicture(req dto.GetProductUploadRequest) (*dto.GetProductUploadResponse, error) {
	tr := entity.UploadedFile{
		FileName: req.FileName,
	}

	if err := b.DB.Where("file_name = ?", tr.FileName).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetProductUploadResponse{
		FilePath: tr.FilePath,
	}

	return response, nil
}
