package userdetail

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *userdetailRepository) Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error) {
	var results []dto.ExploreUserResponse

	// Formula Haversine dalam kilometer (earth radius = 6371 km)
	// distance = 6371 * acos( cos(radians(lat_user)) * cos(radians(lat_target)) * cos(radians(lon_target) - radians(lon_user)) + sin(radians(lat_user)) * sin(radians(lat_target)) )
	// Kita gunakan raw SQL karena lebih efisien

	query := `
        SELECT 
            ud.user_id,
            u.full_name,
            TIMESTAMPDIFF(YEAR, u.birthdate, CURDATE()) as age,
            u.gender,
            ud.latitude,
            ud.longitude,
            (6371 * acos(
                cos(radians(?)) * cos(radians(ud.latitude)) * cos(radians(ud.longitude) - radians(?)) +
                sin(radians(?)) * sin(radians(ud.latitude))
            )) AS distance
        FROM user_details ud
        JOIN users u ON u.id = ud.user_id
        WHERE ud.user_id != ?
            AND ud.latitude IS NOT NULL AND ud.longitude IS NOT NULL
            AND TIMESTAMPDIFF(YEAR, u.birthdate, CURDATE()) BETWEEN ? AND ?   -- min_age, max_age dari preferensi user? Bisa diambil dari user_detail current user
        HAVING distance <= ?
        ORDER BY distance
    `

	// Kita perlu preferensi current user: radius, min_age, max_age, gender_target
	// Ambil dulu preferensi user yang sedang login
	var currentPref entity.UserDetail
	if err := r.DB.Where("user_id = ?", req.UserID).First(&currentPref).Error; err != nil {
		return nil, err
	}

	// Gunakan radius dari request (bisa override dengan preferensi user? Terserah)
	radius := req.Radius
	if radius <= 0 {
		radius = currentPref.Radius
	}

	err := r.DB.Raw(query,
		req.Latitude, req.Longitude, req.Latitude,
		req.UserID,
		currentPref.MinAge, currentPref.MaxAge,
		radius,
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Filter gender target (jika tidak 'all')
	if currentPref.GenderTarget != "all" {
		filtered := make([]dto.ExploreUserResponse, 0)
		for _, u := range results {
			if u.GenderTarget == currentPref.GenderTarget {
				filtered = append(filtered, u)
			}
		}
		results = filtered
	}

	return results, nil
}
