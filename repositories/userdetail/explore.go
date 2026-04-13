package userdetail

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *userdetailRepository) Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error) {
	var results []dto.ExploreUserResponse

	// 1. Ambil preferensi user yang login
	var currentPref entity.UserDetail
	if err := r.DB.Where("user_id = ?", req.UserID).First(&currentPref).Error; err != nil {
		return nil, err
	}

	// 2. Tentukan radius (dari request, preferensi, atau default 50 km)
	radius := req.Radius
	if radius <= 0 {
		radius = currentPref.Radius
		if radius <= 0 {
			radius = 50
		}
	}

	// 3. Tentukan rentang usia
	minAge := currentPref.MinAge
	maxAge := currentPref.MaxAge
	if minAge == 0 {
		minAge = 18
	}
	if maxAge == 0 {
		maxAge = 99
	}

	// 4. Query dengan subquery untuk menghitung distance
	//    Gunakan LEAST/GREATEST untuk menghindari domain error pada acos
	query := `
		SELECT * FROM (
			SELECT 
				ud.id,
				ud.user_id,
				u.full_name,
				u.gender,
				ud.latitude,
				ud.longitude,
				ud.bio,
				ud.radius,
				ud.min_age,
				ud.max_age,
				ud.gender_target,
				u.age,
				(6371 * acos(
					LEAST(1, GREATEST(-1,
						cos(radians(?)) * cos(radians(ud.latitude)) * cos(radians(ud.longitude) - radians(?)) +
						sin(radians(?)) * sin(radians(ud.latitude))
					))
				)) AS distance
			FROM user_details ud
			JOIN access_doors u ON u.id = ud.user_id
			WHERE ud.user_id != ?
				AND ud.latitude IS NOT NULL AND ud.longitude IS NOT NULL
				AND u.age BETWEEN ? AND ?
		) AS sub
		WHERE distance <= ?
		ORDER BY distance
	`

	err := r.DB.Raw(query,
		req.Latitude,   // untuk cos(radians(?)) pertama
		req.Longitude,  // untuk cos(radians(ud.longitude) - radians(?))
		req.Latitude,   // untuk sin(radians(?))
		req.UserID,     // exclude diri sendiri
		minAge, maxAge, // filter umur
		radius, // filter jarak
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// 5. Filter gender target jika tidak 'all'
	if currentPref.GenderTarget != "both" {
		filtered := make([]dto.ExploreUserResponse, 0, len(results))
		for _, u := range results {
			if u.Gender == currentPref.GenderTarget {
				filtered = append(filtered, u)
			}
		}
		results = filtered
	}

	return results, nil
}
