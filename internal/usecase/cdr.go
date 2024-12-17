package usecase

import (
	"billing/internal/entity"
	"billing/internal/repository"
	"billing/pkg/helpers"
	"strconv"
)

type CdrUsecase interface{
    Fetch_Cdr_vas(year string, month string, vas string, call_type string) (string,string)
}
type cdrUsecase struct {
	repo repository.CDRRepository
}

func NewCdrUsecase(repo repository.CDRRepository) CdrUsecase {
	return &cdrUsecase{repo: repo}
}

func (u *cdrUsecase) Fetch_Cdr_vas(year string, month string, vas string, call_type string) (string,string) {
    // text := fmt.Sprintf("lấy CDR %s, với calltype = %s tháng %s năm %s", vas, call_type, month, year)
    
    yearInt, _ := strconv.Atoi(year)
    monthInt, _ := strconv.Atoi(month)

    var cdrRecords []entity.Cdr
    var cdrData [][]string
    var err error
    if call_type == "OUT" {
        if vas == "VAS" {
            cdrRecords, err = u.repo.GetCdrVasOUT(yearInt, monthInt)
        } else if vas == "1900" || vas == "1800" {
            cdrRecords, err = u.repo.GetCdr1900OUT(vas, yearInt, monthInt)
        }
        for _, record := range cdrRecords {
            timeStr := record.Time.Format("2006-01-02 15:04:05")
            row := []string{
                record.Caller,    // Caller
                record.Callee,    // Callee
                timeStr,      // Time
                record.Duration, // Duration
                record.Minute, // Minute
                record.Cost, // Minute
                record.Callee_gw,  // Callee Gateway
            }
            cdrData = append(cdrData, row)
        }
    } else if call_type == "IN" {
        if vas == "VAS" {
            cdrRecords, err = u.repo.GetCdrVasIN(yearInt, monthInt)
        } else if vas == "1900" || vas == "1800" {
            cdrRecords, err = u.repo.Getcdr1900IN(vas, yearInt, monthInt)
        }
        for _, record := range cdrRecords {
            timeStr := record.Time.Format("2006-01-02 15:04:05")
            row := []string{
                record.Caller,    // Caller
                record.Callee,    // Callee
                timeStr,      // Time
                record.Duration, // Duration
                record.Minute, // Minute
                record.Cost, // Minute
                record.Caller_gw,  // Callee Gateway
            }
            cdrData = append(cdrData, row)
        }
    }
    filepath,filename := helpers.Export_Cdr_VAS_to_excel(vas,monthInt,yearInt,call_type,cdrData)

    // Kiểm tra lỗi sau khi gọi hàm
    if err != nil {
        return "", "Lỗi khi lấy dữ liệu CDR: " + err.Error()
    }
    
    // Trả về thông điệp sau khi xử lý thành công
    return filepath,filename
}

