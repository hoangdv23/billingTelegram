package databases

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var (
	dbBilling *gorm.DB
	db136     *gorm.DB
	dbLog     *gorm.DB
	once      sync.Once
)

// Hàm nạp cấu hình từ tệp .env
func LoadConfig() (DBConfig, DBConfig, DBConfig) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Không thể nạp tệp .env: %v", err)
	}

	// Cấu hình cho Billing DB
	billingConfig := DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}

	// Cấu hình cho DB136
	db136Config := DBConfig{
		Username: os.Getenv("DB_136_USERNAME"),
		Password: os.Getenv("DB_136_PASSWORD"),
		Host:     os.Getenv("DB_136_HOST"),
		Port:     os.Getenv("DB_136_PORT"),
		Database: os.Getenv("DB_136_NAME"),
	}

	// Cấu hình cho DB Log
	logConfig := DBConfig{
		Username: os.Getenv("DB_LOG_USERNAME"),
		Password: os.Getenv("DB_LOG_PASSWORD"),
		Host:     os.Getenv("DB_LOG_HOST"),
		Port:     os.Getenv("DB_LOG_PORT"),
		Database: os.Getenv("DB_LOG_NAME"),
	}

	return billingConfig, db136Config, logConfig
}

// Hàm kết nối đến cơ sở dữ liệu sử dụng GORM
func connectToDB(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		config.Username, config.Password, config.Host, config.Port, config.Database)

	// Kết nối đến cơ sở dữ liệu MySQL với GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("không thể kết nối tới DB (%s): %v", config.Database, err)
	}

	log.Printf("Kết nối thành công tới DB: %s", config.Database)
	return db, nil
}

// Hàm khởi tạo kết nối Billing DB
func InitBillingDB() {
	once.Do(func() {
		billingConfig, _, _ := LoadConfig()
		var err error
		dbBilling, err = connectToDB(billingConfig)
		if err != nil {
			log.Fatalf("Lỗi khi khởi tạo Billing DB: %v", err)
		}
		if dbBilling == nil {
			log.Fatal("Kết nối tới Billing DB bị nil!")
		}
	})
}

// Hàm khởi tạo kết nối DB136
func InitDB136() {
	once.Do(func() {
		_, db136Config, _ := LoadConfig()
		var err error
		db136, err = connectToDB(db136Config)
		if err != nil {
			log.Fatalf("Lỗi khi khởi tạo DB136: %v", err)
		}
		if db136 == nil {
			log.Fatal("Kết nối tới DB136 bị nil!")
		}
	})
}

// Hàm khởi tạo kết nối DB Log
func InitDBLog() {
	once.Do(func() {
		_, _, logConfig := LoadConfig()
		var err error
		dbLog, err = connectToDB(logConfig)
		if err != nil {
			log.Fatalf("Lỗi khi khởi tạo DB Log: %v", err)
		}
		if dbLog == nil {
			log.Fatal("Kết nối tới DB Log bị nil!")
		}
	})
}

// Lấy kết nối Billing DB
func GetBillingDB() *gorm.DB {
	if dbBilling == nil {
		InitBillingDB()
	}
	return dbBilling
}

// Lấy kết nối DB136
func GetDB136() *gorm.DB {
	if db136 == nil {
		InitDB136()
	}
	return db136
}

// Lấy kết nối DB Log
func GetDBLog() *gorm.DB {
	if dbLog == nil {
		InitDBLog()
	}
	return dbLog
}
