package qris

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *qrisService) Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error) {

	create := dto.CoQrisRequest{
		QrisName:   req.QrisName,
		Link:       req.Link,
		File:       req.File,
		Status:     req.Status,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.CoQrisResponse{}, err
	}

	response := dto.CoQrisResponse{
		QrisName:   created.QrisName,
		Link:       created.Link,
		FilePath:   created.FilePath,
		Status:     created.Status,
		UserID:     created.UserID,
		MerchantID: created.MerchantID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}
