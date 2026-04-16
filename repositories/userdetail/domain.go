package userdetail

import (
	dto "github.com/srv-api/detail/dto"
	limit "github.com/srv-api/detail/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error)
	LongLat(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
	GetById(req dto.GetUserDetailByIdRequest) (*dto.UserDetailRequest, error)
	Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
	GetUserLimit(userID string) (*limit.UserLimit, error)
}

type userdetailRepository struct {
	DB *gorm.DB
}

func NewUserDetailRepository(DB *gorm.DB) DomainRepository {
	return &userdetailRepository{
		DB: DB,
	}
}
