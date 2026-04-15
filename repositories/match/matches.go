package repository

import (
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

type MatchRepository interface {
	CreateMatch(user1, user2 string) error
	GetMatches(userID string) ([]map[string]interface{}, error)
}

type matchRepository struct {
	DB *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepository{DB: db}
}

// insert match (hindari duplicate)
func (r *matchRepository) CreateMatch(user1, user2 string) error {
	match := entity.Match{
		User1ID: user1,
		User2ID: user2,
	}

	return r.DB.
		Where("user1_id = ? AND user2_id = ? OR user1_id = ? AND user2_id = ?", user1, user2, user2, user1).
		FirstOrCreate(&match).Error
}

// ambil semua match user
func (r *matchRepository) GetMatches(userID string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	query := `
	SELECT 
		CASE 
			WHEN m.user1_id = ? THEN m.user2_id
			ELSE m.user1_id
		END as user_id,
		a.full_name,
		COALESCE(uf.file_path, '') as profile_picture
	FROM matches m
	JOIN access_doors a 
		ON a.id = CASE 
			WHEN m.user1_id = ? THEN m.user2_id
			ELSE m.user1_id
		END
	LEFT JOIN uploaded_files uf 
		ON uf.user_id = a.id AND uf.deleted_at IS NULL
	WHERE m.user1_id = ? OR m.user2_id = ?
	ORDER BY m.created_at DESC
	`

	err := r.DB.Raw(query, userID, userID, userID, userID).
		Scan(&results).Error

	return results, err
}
