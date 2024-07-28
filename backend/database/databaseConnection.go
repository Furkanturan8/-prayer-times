package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"namaz-vakitleri/models"
	"namaz-vakitleri/pkg/config"
)

func DBInstance(cfg *config.Config) (*gorm.DB, error) {
	// MySQL DSN oluştur
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQLUsername, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDBName)

	// MySQL veritabanına bağlantı aç
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Automigrate models
	err = db.AutoMigrate(
		&models.City{},
		&models.Timings{},
		&models.GregorianDate{},
		&models.HijriDate{},
		&models.Phrases{},
		&models.Contact{},
	)
	if err != nil {
		return nil, fmt.Errorf("AutoMigrate failed: %v", err)
	}

	fmt.Printf("Connected to MySQL database: %s\n", cfg.MySQLDBName)

	return db, nil
}
