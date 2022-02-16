package db

import (
  "main/models"
  
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func GetDbConnection() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
  if err != nil {
    panic("Error during DB connection")
  }
  return db
}

func MigrateSchemas() {
  db := GetDbConnection();
  db.AutoMigrate(&models.Anime{}, &models.Episode{})
}