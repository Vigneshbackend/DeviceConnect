package model
import(
	"github.com/jinzhu/gorm"
)



type Transactiondetails struct {
	gorm.Model
	CustomerID          string `gorm:"column:customer_id"`
	Merchantdetails     Merchantdetails
	merchant    string
	transaction_hash   string
	time  string
	timeinmilis int64
	servicename string
	account_hash string
	category string
	amount float64
	tag string
	account_number string
	related_transaction_hash string
	Type string
	channel string
	
}

