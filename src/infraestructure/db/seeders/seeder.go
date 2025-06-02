package seeders

import (
	"log"

	"gorm.io/gorm"
)

func SeedData[T any](db *gorm.DB, data []T, uniqueColumn string, getValue func(T) interface{}) {
	for _, record := range data {
		var existingRecord T
		result := db.Where(uniqueColumn+" = ?", getValue(record)).First(&existingRecord)

		if result.Error != nil {
			if err := db.Create(&record).Error; err != nil {
				log.Printf("❌ Error inserting record: %v\n", err)
			} else {
				log.Printf("✅ Inserted record: %+v\n", record)
			}
		} else {
			log.Printf("⚠️ Already exists: %+v\n", record)
		}
	}

	log.Println("Seed data completed")
}
