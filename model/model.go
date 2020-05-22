package model

type AccountsAndTransactional struct {
	account     interface{}
	transaction interface{}
}

type Output struct {
	Code  int64
	Data  interface{}
	Error string
}
type Transactionresponse struct {
	data []string
}

type Req struct {
	Merchant_id string`json:"merchant_id"`
}

type Reqbody struct {
	Customer_id string `json:"customer_id"`
	Version     int    `json:"version"`
	Salt        string `json:"salt"`
	Num_records string `json:"num_records"`
	Start_time  string `json:"start_time"`
	End_time    string `json:"end_time"`
}

type DateFilter struct {
	Customer_id string
	Version     int
	Salt        string
	Start_date  string
	End_date    string
}

type Transactionobject struct {
	Status         string
	date_requested string
	date_processed string
	request_id     string
	message        string
	customer_id    string
	Data           []Data
}
type Data struct {
	Merchant                 string
	Transaction_hash         string
	Time                     string
	Timeinmilis              int
	Servicename              string
	Account_hash             string
	Category                 string
	Amount                   float64
	Tag                      string
	Account_number           string
	Related_transaction_hash string
	Type                     string
	Channel                  string
}

type Error struct {
	Code int
	Msg  string
}
type Transactionoutput struct {
	Code int64
	// Data  interface{}
	Error             string
	TotalTransaction  int
	TransactionAmount float64
}

type Datainsert struct {
	CustomerID               string
	Merchant                 string
	Transaction_hash         string
	Time                     string
	Timeinmilis              int
	Servicename              string
	Account_hash             string
	Category                 string
	Amount                   float64
	Tag                      string
	Account_number           string
	Related_transaction_hash string
	Type                     string
	Channel                  string
}

type RecentlyAccessed struct {
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
	Active_months_list  []string
	Type                string
}



type AccountList struct{

	List []RecentlyAccessed

}



type AccoutResponse struct{
	Status        string
	Daterequested string
	Dateprocessed string
	Requestid     string
	Message       string
	Customerid    string
	Data []AccountList
}









type AccountOutput struct{
	Customerid string `json:"customer_id"`
	Date_processed string
	Data []RecentlyAccessed 
	Date_requested string
	Message string
	Requestid string `json:"request_id"`
	Status string 
}


type SummaryAccount struct{


	Customerid string `json:"customer_id"`
	Date_processed string
	Data []SummaryAccountData
	Date_requested string
	Message string
	Requestid string `json:"request_id"`
	Status string 


}



type SummaryAccountData struct{




	Company             string
	Number              string
	Latest_balance_date string
	Latest_balance      float64
	Type                string
}

type SummaryAccountOutput struct{
	Code int
	Data SummaryAccount
	Error string
}



type QrSummaryOutput struct{
Code int
Error string
Data QrSummaryData 


}


type QrSummaryData struct{
	TotalTransaction int
	Status string
    TransactionAmount float64
}

