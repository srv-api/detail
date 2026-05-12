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

			// Reset swipe untuk user NON-premium saja menggunakan JOIN
			result := db.Table("user_limits").
				Joins("JOIN user_details ON user_limits.user_id = user_details.user_id").
				Where("user_details.is_premium = ?", false).
				Where("user_limits.updated_at < ?", time.Now().Add(-24*time.Hour)).
				Updates(map[string]interface{}{
					"user_limits.remaining_swipe":      50,
					"user_limits.remaining_super_like": 1,
					"user_limits.updated_at":           time.Now(),
				})

			if result.Error != nil {
				log.Printf("Failed to reset daily swipe: %v", result.Error)
			} else {
				log.Printf("Reset daily swipe for %d non-premium users", result.RowsAffected)
			}
		}
	}()
}

func StartTenMinutesJob(db *gorm.DB) {
	// Hitung 10 menit dari sekarang
	delay := 10 * time.Minute

	log.Printf("Cronjob akan berjalan dalam 10 menit pada: %v", time.Now().Add(delay))

	time.AfterFunc(delay, func() {
		// Eksekusi task setelah 10 menit
		log.Println("Menjalankan cronjob setelah 10 menit...")

		// Contoh task: reset atau update data
		result := db.Model(&entity.UserLimit{}).
			Where("updated_at < ?", time.Now().Add(-24*time.Hour)).
			Updates(map[string]interface{}{
				"remaining_swipe":      50,
				"remaining_super_like": 1,
				"updated_at":           time.Now(),
			})

		if result.Error != nil {
			log.Printf("Error: %v", result.Error)
		} else {
			log.Printf("Berhasil menjalankan task, affected rows: %d", result.RowsAffected)
		}
	})
}
