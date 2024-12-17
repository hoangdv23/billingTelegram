package helpers

import (
	"fmt"
	"log"
	"github.com/xuri/excelize/v2"
)

func Export_Cdr_VAS_to_excel(vas string,month int, year int, call_type string,data [][]string) (string, string){
	var filename string
	f := excelize.NewFile()
	sheetName := "Sheet1"
	
	index,_ := f.NewSheet(sheetName)

	f.SetCellValue(sheetName, "A1", "Caller")
	f.SetCellValue(sheetName, "B1", "Callee")
	f.SetCellValue(sheetName, "C1", "Time")
	f.SetCellValue(sheetName, "D1", "Duration")
	f.SetCellValue(sheetName, "E1", "Minute")
	f.SetCellValue(sheetName, "F1", "Cost")
	if call_type == "OUT" {
		filename = fmt.Sprintf("Cdr_Digitel_call_OUT_TELCO_%s_%d_%d.xlsx",vas,month,year)
		f.SetCellValue(sheetName, "G1", "Callee GateWay")
	}else if call_type == "IN" {
		filename = fmt.Sprintf("Cdr_%s_call_IN_Digitel_%d_%d.xlsx",vas,month,year)
		f.SetCellValue(sheetName, "G1", "Caller GateWay")
	}
	for i, row := range data {
		for j, cell := range row {
			cellName := fmt.Sprintf("%c%d", 'A'+j, i+2) // Bắt đầu từ hàng 2
			f.SetCellValue(sheetName, cellName, cell)
		}
	}
	
	f.SetActiveSheet(index)
	filepath := "/root/billingTeleBot/assets/excel/" + filename;

	if err := f.SaveAs(filepath); err != nil {
		log.Fatalf("Lỗi khi lưu file Excel: %v", err)
	} else {
		fmt.Println("Done, được lưu tại "+ filepath)
	}
	return filepath,filename
}