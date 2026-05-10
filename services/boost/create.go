package boost

import (
	"time"

	dto "github.com/srv-api/detail/dto"
	util "github.com/srv-api/util/s"
)

func (s *boostService) Create(req dto.BoostRequest) (dto.BoostResponse, error) {
	now := time.Now()

	// otomatis 30 menit
	startTime := now
	endTime := now.Add(30 * time.Minute)

	create := dto.BoostRequest{
		ID:        util.GenerateRandomString(),
		UserID:    req.UserID,
		DetailID:  req.DetailID,
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: now,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.BoostResponse{}, err
	}

	return created, nil
}
