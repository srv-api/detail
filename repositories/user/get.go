package user

import (
	"fmt"
	"math"
	"strings"

	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/helpers"
	util "github.com/srv-api/util/s"
)

func (r *userRepository) Get(req *dto.Pagination) (dto.UserPaginationResponse, int) {
	var users []entity.AccessDoor

	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0

	// Ubah offset agar sesuai dengan page yang dimulai dari 1
	offset := (req.Page - 1) * req.Limit

	// Ambil data sesuai limit, offset, dan urutan
	find := r.DB.Preload("Verified").Preload("Merchant").Preload("ProfilePicture").Limit(req.Limit).Offset(offset).Order(req.Sort)

	// Generate where query untuk search
	if req.Searchs != nil {
		for _, value := range req.Searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				find = find.Where(fmt.Sprintf("%s = ?", column), query)
			case "contains":
				find = find.Where(fmt.Sprintf("%s LIKE ?", column), "%"+query+"%")
			case "in":
				find = find.Where(fmt.Sprintf("%s IN (?)", column), strings.Split(query, ","))
			}
		}
	}

	find = find.Find(&users)

	// Periksa jika ada error saat pengambilan data
	if errFind := find.Error; errFind != nil {
		return dto.UserPaginationResponse{}, totalPages
	}

	req.Rows = users

	// Hitung total data
	if errCount := r.DB.Model(&entity.AccessDoor{}).Count(&totalRows).Error; errCount != nil {
		return dto.UserPaginationResponse{}, totalPages
	}

	for i := range users {
		users[i].FullName = helpers.TruncateString(users[i].FullName, 47)
	}

	req.TotalRows = int(totalRows)

	// Hitung total halaman berdasarkan limit
	totalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages
	// Hitung `fromRow` dan `toRow` untuk page saat ini
	if req.Page == 1 {
		// Untuk halaman pertama
		fromRow = 1
		toRow = req.Limit
	} else {
		if req.Page <= totalPages {
			fromRow = (req.Page-1)*req.Limit + 1
			toRow = req.Page * req.Limit
		}
	}

	// Pastikan `toRow` tidak melebihi `totalRows`
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	// Set hasil akhir
	req.FromRow = fromRow
	req.ToRow = toRow
	var userResponses []dto.UserResponse
	for _, u := range users {

		decryptedWa, err := util.Decrypt(u.Whatsapp)
		if err != nil {
			return dto.UserPaginationResponse{}, totalPages

		}

		decryptedEmail, err := util.Decrypt(u.Email)
		if err != nil {
			return dto.UserPaginationResponse{}, totalPages

		}
		userResp := dto.UserResponse{
			ID:       u.ID,
			FullName: u.FullName,
			Whatsapp: decryptedWa,
			Email:    decryptedEmail,
			Verified: dto.UserVerified{
				ID:             u.Verified.ID,
				UserID:         u.Verified.UserID,
				Token:          u.Verified.Token,
				Verified:       u.Verified.Verified,
				StatusAccount:  u.Verified.StatusAccount,
				AccountExpired: u.Verified.AccountExpired,
				Otp:            u.Verified.Otp,
				ExpiredAt:      u.Verified.ExpiredAt,
			},
			UserDetail: dto.UserDetailResponse{
				ID:             u.Merchant.ID,
				UserID:         u.Merchant.UserID,
				Latitude:       u.Merchant.Latitude,
				Longitude:      u.Merchant.Longitude,
				Radius:         u.Merchant.Radius,
				MinAge:         u.Merchant.MinAge,
				MaxAge:         u.Merchant.MaxAge,
				GenderTarget:   u.Merchant.GenderTarget,
				ProfilePicture: "http://103.150.227.223:2356" + u.ProfilePicture.FileName,
				UpdatedAt:      u.Merchant.UpdatedAt,
			},
		}
		userResponses = append(userResponses, userResp)
	}
	response := dto.UserPaginationResponse{
		Limit:        req.Limit,
		Page:         req.Page,
		Sort:         req.Sort,
		TotalRows:    req.TotalRows,
		TotalPages:   req.TotalPages,
		FirstPage:    req.FirstPage,
		PreviousPage: req.PreviousPage,
		NextPage:     req.NextPage,
		LastPage:     req.LastPage,
		FromRow:      req.FromRow,
		ToRow:        req.ToRow,
		Data:         userResponses,
		Searchs:      req.Searchs,
	}

	return response, totalPages
}
