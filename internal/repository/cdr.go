package repository

import (
	"billing/internal/entity"
	"fmt"
	"time"
	"sync"
	"gorm.io/gorm"
	"log"
)

type CDRRepository interface {
	GetCdr1900OUT(vas string, year int, month int) ([]entity.Cdr, error)
	GetCdrVasOUT( year int, month int) ([]entity.Cdr, error)
	Getcdr1900IN(vas string, year int, month int) ([]entity.Cdr, error)
	GetCdrVasIN( year int, month int) ([]entity.Cdr, error)
}

type cdrRepository struct {
	db *gorm.DB
}

func NewCDRRepository(db *gorm.DB) CDRRepository {
	return &cdrRepository{db: db}
}

// Lấy cdr DIGITEL gọi 1900 nhà mạng
func (r *cdrRepository) GetCdr1900OUT(vas string, year int, month int) ([]entity.Cdr, error) {
	var allRecords []entity.Cdr
	lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := lastDay.Day()

	var wg sync.WaitGroup
	mu := sync.Mutex{} 

	for day := 1; day <= daysInMonth; day++ {
		wg.Add(1)
		go func(day int) {
			defer wg.Done()
			tableName := fmt.Sprintf("cdr%04d%02d%02d", year, month, day)
			var cdrRecords []entity.Cdr

			// Thực hiện truy vấn
			err := r.db.Table(tableName).
					Select("caller", "callee", "time", "duration","minute","cost", "callee_gw").
					Where("Callee LIKE ?", vas+"%").
					Or("Callee LIKE ?", "concat(84,"+vas+")%").
					Where("call_type LIKE ?", "OUT_VAS").
					Find(&cdrRecords).Error

			if err != nil {
				// Log lỗi nếu có nhưng không làm gián đoạn toàn bộ quá trình
				log.Printf("Error querying table %s: %v", tableName, err)
				return
			}
			// Cập nhật kết quả vào allRecords một cách thread-safe
			mu.Lock()
			allRecords = append(allRecords, cdrRecords...)
			mu.Unlock()
		}(day)
	}
	wg.Wait()
	return allRecords, nil
}



func (r *cdrRepository) GetCdrVasOUT( year int, month int) ([]entity.Cdr, error) {
	var allRecords []entity.Cdr
	lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	daysInMonth := lastDay.Day()

	var wg sync.WaitGroup
	mu := sync.Mutex{} 

	for day := 1; day <= daysInMonth; day++ {
		wg.Add(1)
		go func(day int) {
			defer wg.Done()
			tableName := fmt.Sprintf("cdr%04d%02d%02d", year, month, day)
			var cdrRecords []entity.Cdr

			// Thực hiện truy vấn
			err := r.db.Table(tableName).
					Select("caller", "callee", "time", "duration","minute","cost", "callee_gw").
					Where("Callee LIKE ?", "1900%").
					Or("Callee LIKE ?", "841900%").
					Or("Callee LIKE ?", "1800%").
					Or("Callee LIKE ?", "841800%").
					Where("call_type LIKE ?", "OUT_VAS").
					Find(&cdrRecords).Error

			if err != nil {
				// Log lỗi nếu có nhưng không làm gián đoạn toàn bộ quá trình
				log.Printf("Error querying table %s: %v", tableName, err)
				return
			}
			// Cập nhật kết quả vào allRecords một cách thread-safe
			mu.Lock()
			allRecords = append(allRecords, cdrRecords...)
			mu.Unlock()
		}(day)
	}
	wg.Wait()
	return allRecords, nil
}

func (r *cdrRepository) Getcdr1900IN(vas string, year int, month int) ([]entity.Cdr, error) {
    var allRecords []entity.Cdr
    tableName := fmt.Sprintf("cdrdvgtgt%04d%02d", year, month)

    // Thực hiện truy vấn
    err := r.db.Table(tableName).
        Select("caller", "callee", "time", "duration","minute","cost", "caller_gw").
        Where("Callee LIKE ?", vas+"%").
        Or("Callee LIKE ?", "concat(84,"+vas+")%").
        Where("call_type LIKE ?", "OUT_VAS").
        Find(&allRecords).Error // Lưu kết quả vào allRecords
    if err != nil {
        // Log lỗi nếu có nhưng không làm gián đoạn toàn bộ quá trình
        log.Printf("Error querying table %s: %v", tableName, err)
        return nil, err // Trả về lỗi nếu có
    }

    // Trả về kết quả nếu không có lỗi
    return allRecords, nil
}


func (r *cdrRepository) GetCdrVasIN( year int, month int) ([]entity.Cdr, error) {
	var allRecords []entity.Cdr

	tableName := fmt.Sprintf("cdr%04d%02d", year, month)
	var cdrRecords []entity.Cdr

	// Thực hiện truy vấn
	err := r.db.Table(tableName).
			Select("caller", "callee", "time", "duration","minute","cost", "caller_gw").
			Where("callee_object", "DIGITEL_VAS").
			Find(&cdrRecords).Error

	if err != nil {
		// Log lỗi nếu có nhưng không làm gián đoạn toàn bộ quá trình
		log.Printf("Error querying table %s: %v", tableName, err)
		return nil, err
	}

	return allRecords, nil
}