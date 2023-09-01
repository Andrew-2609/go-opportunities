package config

import (
	"os"

	"github.com/Andrew-2609/go-opportunities/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found. Creating it...")

		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("DB file directory creation error: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("DB file creation error: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})

	if err != nil {
		logger.Errorf("SQLite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
