package cron

import (
	"log"
	"time"

	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

func StartDailyReset(db *gorm.DB) {
	// Jalankan setiap jam 00:00
	ticker := time.NewTicker(24 * time.Hour)

	go func() {
		for range ticker.C {
			// Tunggu sampai jam 00:00
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			time.Sleep(time.Until(next))

			// Reset swipe semua user
			result := db.Model(&entity.UserLimit{}).
				Where("updated_at < ?", time.Now().Add(-24*time.Hour)).
				Updates(map[string]interface{}{
					"remaining_swipe":      50,
					"remaining_super_like": 1,
					"updated_at":           time.Now(),
				})

			if result.Error != nil {
				log.Printf("Failed to reset daily swipe: %v", result.Error)
			} else {
				log.Printf("Reset daily swipe for %d users", result.RowsAffected)
			}
		}
	}()
}
