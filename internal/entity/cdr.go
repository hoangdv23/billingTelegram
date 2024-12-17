package entity

import "time"

type Cdr struct {
    Caller  string `gorm:"type:varchar(100)"`
    Callee string `gorm:"type:varchar(100)"`
    Time time.Time `gorm:"type:varchar(100)"`
    Minute string `gorm:"type:varchar(100)"`
    Duration string `gorm:"type:varchar(100)"`
    Cost string `gorm:"type:varchar(100)"`
    Caller_gw string `gorm:"type:varchar(100)"`
    Caller_object string `gorm:"type:varchar(100)"`
	Callee_gw string `gorm:"type:varchar(100)"`
    Callee_object string `gorm:"type:varchar(100)"`
}


