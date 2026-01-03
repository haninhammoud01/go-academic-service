package entity

import "gorm.io/gorm"

// AutoMigrate runs all migrations
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Student{},
		&Lecturer{},
		&Course{},
		&Enrollment{},
	)
}
