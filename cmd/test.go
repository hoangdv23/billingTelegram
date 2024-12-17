package cmd

import (
	"billing/internal/databases"
	// "billing/internal/repository"
	// "billing/internal/usecase"
	"fmt"
)
func Test(){
	db136 := databases.GetDB136()
	fmt.Println(db136)
	// cdrRepo := repository.NewCDRRepository(db136)
	// cdrUsecase := usecase.NewCdrUsecase(cdrRepo)
	// year := 2024
	// month := 11
	// telco := "Viettel"
	// vas := "1900";
	// callType := "IN"
	// records, _ := cdrUsecase.Fetch_Cdr1900OUT(year, month,vas,callType, telco)

	// fmt.Println(records)
}