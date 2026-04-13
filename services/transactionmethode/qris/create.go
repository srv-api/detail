package qris

import (
	dto "github.com/srv-api/detail/dto"
)

func (s *qrisService) Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error) {

	create := dto.CoQrisRequest{
		QrisName:  req.QrisName,
		Link:      req.Link,
		File:      req.File,
		Status:    req.Status,
		UserID:    req.UserID,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.CoQrisResponse{}, err
	}

	response := dto.CoQrisResponse{
		QrisName:  created.QrisName,
		Link:      created.Link,
		FilePath:  created.FilePath,
		Status:    created.Status,
		UserID:    created.UserID,
		DetailID:  created.DetailID,
		CreatedBy: created.CreatedBy,
	}

	return response, nil
}
