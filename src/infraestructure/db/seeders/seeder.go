package seeders

import (
	"log"
	"strings"
	"sync"

	"gorm.io/gorm"
)

// SeedData inserts data into the database. It takes a slice of data of type T, a string of a unique column
// and a function that returns the value of the unique column for each record. If the record already exists, it will not be inserted.
// The function will print the status of the insertion for each record.
//
// This function is useful when inserting a small amount of data, as it is easier to read and debug.
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

// SeedDataConcurrent inserts data concurrently using goroutines. It takes a slice of data of type T, a string of unique columns
// and a function that returns the values of the unique columns for each record. If the record already exists, it will not be inserted.
// The function will print the status of the insertion for each record.
//
// This function is useful when inserting a large amount of data, as it can speed up the process.
//
// The number of concurrent goroutines is set to 10, but can be changed by modifying the size of the semaphore channel.
func SeedDataConcurrent[T any](db *gorm.DB, data []T, uniqueColumns string, getValues func(T) interface{}) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // controla cuántas goroutines corren en paralelo (10 aquí)

	for _, record := range data {
		wg.Add(1)
		semaphore <- struct{}{} // ocupa un lugar

		go func(record T) {
			defer wg.Done()
			defer func() { <-semaphore }() // libera el lugar

			var existing T
			query := db

			values := getValues(record)

			// Si se pasan múltiples columnas
			switch v := values.(type) {
			case []interface{}:
				cols := splitColumns(uniqueColumns)
				for i, col := range cols {
					query = query.Where(col+" = ?", v[i])
				}
			default:
				query = query.Where(uniqueColumns+" = ?", values)
			}

			result := query.First(&existing)

			if result.Error != nil {
				if err := db.Create(&record).Error; err != nil {
					log.Printf("❌ Error inserting: %v\n", err)
				} else {
					log.Printf("✅ Inserted: %+v\n", record)
				}
			} else {
				log.Printf("⚠️ Exists: %+v\n", record)
			}
		}(record)
	}

	wg.Wait()
	log.Println("✅ Concurrent seeding complete")
}

func splitColumns(cols string) []string {
	var result []string
	for _, c := range strings.Split(cols, ",") {
		result = append(result, strings.TrimSpace(c))
	}
	return result
}
