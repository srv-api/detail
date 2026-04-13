package userdetail

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *userdetailRepository) Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error) {
	var results []dto.ExploreUserResponse

	// Ambil preferensi user yang sedang login (radius, min_age, max_age, gender_target)
	var currentPref entity.UserDetail
	if err := r.DB.Where("user_id = ?", req.UserID).First(&currentPref).Error; err != nil {
		return nil, err
	}

	// Gunakan radius dari request, fallback ke preferensi user, lalu default 50 km
	radius := req.Radius
	if radius <= 0 {
		radius = currentPref.Radius
		if radius <= 0 {
			radius = 50
		}
	}

	// Gunakan min_age dan max_age dari preferensi user (bisa juga dari request jika ada)
	minAge := currentPref.MinAge
	maxAge := currentPref.MaxAge
	if minAge == 0 {
		minAge = 18
	}
	if maxAge == 0 {
		maxAge = 99
	}

	// Query SQL dengan Haversine formula
	query := `
        SELECT 
            ud.user_id,
            u.full_name,
            u.gender,
            ud.latitude,
            ud.longitude,
            ud.bio,
            ud.radius as user_radius,
            ud.min_age,
            ud.max_age,
            ud.gender_target,
            TIMESTAMPDIFF(YEAR, u.birthdate, CURDATE()) as age,
            (6371 * acos(
                cos(radians(?)) * cos(radians(ud.latitude)) * cos(radians(ud.longitude) - radians(?)) +
                sin(radians(?)) * sin(radians(ud.latitude))
            )) AS distance
        FROM user_details ud
        JOIN users u ON u.id = ud.user_id
        WHERE ud.user_id != ?
            AND ud.latitude IS NOT NULL AND ud.longitude IS NOT NULL
            AND TIMESTAMPDIFF(YEAR, u.birthdate, CURDATE()) BETWEEN ? AND ?
        HAVING distance <= ?
        ORDER BY distance
    `

	err := r.DB.Raw(query,
		req.Latitude, req.Longitude, req.Latitude,
		req.UserID,
		minAge, maxAge,
		radius,
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Filter berdasarkan gender target (jika tidak 'all')
	if currentPref.GenderTarget != "all" {
		filtered := make([]dto.ExploreUserResponse, 0)
		for _, u := range results {
			// Bandingkan gender user dengan gender target yang diinginkan
			if u.Gender == currentPref.GenderTarget {
				filtered = append(filtered, u)
			}
		}
		results = filtered
	}

	return results, nil
}
