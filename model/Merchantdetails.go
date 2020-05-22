package model

import "github.com/jinzhu/gorm"

type Merchantdetails struct {
	gorm.Model
	Status        string
	Daterequested string
	Dateprocessed string
	Requestid     string
	Message       string
	Customerid    string
	Accounts      []Accounts `gorm:"ForeignKey:CustomerID"`
}
