package model

import (
	"github.com/jinzhu/gorm"
)

type Accounts struct {
	gorm.Model
	CustomerID          string `gorm:"column:customer_id"`
	Merchantdetails     Merchantdetails
	Company             string
	Last_used_date      string
	Number              string
	Latest_balance_date string
	Latest_bill         float64
	Account_hash        string
	Latest_balance      float64
	Latest_bill_date    string
	Is_primary          bool
	Limit               float64
	Active_months_list  string
	Type                string
}
