package userdetail

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error)
	LongLat(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
	GetById(req dto.GetUserDetailByIdRequest) (*dto.UserDetailRequest, error)
	Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
}

type userdetailRepository struct {
	DB *gorm.DB
}

func NewUserDetailRepository(DB *gorm.DB) DomainRepository {
	return &userdetailRepository{
		DB: DB,
	}
}
