package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB 			*gorm.DB
	AppConfig 	*Config
)

type Config struct {
	AppPort 			string
	DBHost				string
	DBPort				string
	DBUser				string
	DBPassword			string
	DBName				string
	JWTSecret			string
	JWTExpireMinutes	string
	JWTRefreshToken		string
	JWTExpire			string
}

// function load
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	AppConfig = &Config{
		AppPort: getEnv("PORT","3030"),
		DBHost: getEnv("DB_HOST","localhost"),
		DBPort: getEnv("DB_PORT","5432"),
		DBUser: getEnv("DB_USER","postgres"),
		DBPassword: getEnv("DB_PASSWORD","ROOT"),
		DBName: getEnv("DB_NAME","project_management"),
		JWTSecret: getEnv("JWT_SECRET","supersecret"),
		JWTExpireMinutes: getEnv("JWT_EXPIRY_MINUTES","6000"),
		JWTRefreshToken: getEnv("REFRESH_TOKEN_EXPIRED","24h"),
		JWTExpire: getEnv("JWT_EXPIRED","2h"),
	}
}

// ambil value yg ada di dotenv
func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return fallback
	}
}

// connect to db
func ConnectDB() {
	cfg := AppConfig

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",cfg.DBHost,cfg.DBPort,cfg.DBUser,cfg.DBPassword,cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

}
