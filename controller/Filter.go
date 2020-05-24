package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/jinzhu/gorm"

	"DeviceConnect/model"
	"DeviceConnect/repo"

	"strconv"
	"time"
)

var start_time string
var end_time string
var Transactionamount float64

type Result struct {
	Data []Resultdata
}
type FilterController struct {
	DB *gorm.DB
}

type Resultdata struct {
	Merchant                 string
	Transaction_hash         string
	Time                     string
	Timeinmilis              int
	Servicename              string
	Account_hash             string
	Category                 string
	Amount                   int
	Tag                      string
	Account_number           string
	Related_transaction_hash string
	Type                     string
	Channel                  string
}

func Timeconversion() {
	// layout := "01/02/2006 3:04:05 PM"
	// refer := retime
	// fmt.Println("---->print", rettime)
	// t, err := time.Parse(layout, rettime)
	// if err != nil {
	// }
	// unix := t.UnixNano() / 1000000

	// if name == "start_time" {
	// 	start_time = strconv.FormatInt(unix, 10)
	// } else {
	// 	end_time = strconv.FormatInt(unix, 10)

	// }

	now := time.Now()
	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000
	end_time = strconv.FormatInt(umillisec, 10)

	now_string := now.Format("2006-02-01")
	midnight_time, _ := time.Parse("2006-02-01", now_string)

	unixNano = midnight_time.UnixNano()
	umillisec = unixNano / 1000000
	start_time = strconv.FormatInt(umillisec, 10)

	fmt.Print("-------------------->midnight", midnight_time)

	// end_time = time.ParseD(time.Now().Year(), time.Now().Month(), time.Now().Day(),time.Now().Hour()-time.Now().Hour(), time.Now().Minute()-time.Now().Minute(), time.Now().Second()-time.Now().Second(), 0, time.UTC)

}

func Filter(name string) model.QrSummaryData {

	var out model.QrSummaryData

	if name != "" {

		Init()
		var refer model.Reqbody
		var transaction model.Transactionobject

		Timeconversion()
		refer.Customer_id = name
		refer.Version = version
		refer.Salt = repo.GetSaltForCustomer(name)
		refer.Num_records = ""
		refer.Start_time = end_time
		refer.End_time = start_time

		b, err := json.Marshal(refer)
		if err == nil {
			client := &http.Client{}

			apiURL := "https://insights.finbox.in/staging/transactions?num_records=1000&start_time=" + refer.Start_time + "&end_time=" + refer.End_time
			u, _ := url.ParseRequestURI(apiURL)
			urlStr := u.String()
			req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(b))
			req.Header.Add("X-api-key", xapikey)
			req.Header.Add("Content-Type", `application/json`)
			fmt.Print("---->req", req)
			resp, err := client.Do(req)
			defer resp.Body.Close()

			if err != nil {
				print("-------------->", err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			// err = json.Unmarshal([]byte(body), &out.Data)

			err = json.Unmarshal([]byte(body), &transaction)
			if err != nil {
				fmt.Print("___________________>", err)
			}

			if transaction.Status != "in_progress" {

				var test []model.Data
				for _, b := range transaction.Data {
					fmt.Println("------------------------------>tr", b)
					if b.Type == "debit" && b.Channel == "upi" {
						Transactionamount = Transactionamount + b.Amount
						test = append(test, b)
					}

				}

				out.TotalTransaction = len(test)
				out.TransactionAmount = Transactionamount
				if err != nil {
					print(err)
					out.Status = "fetch failed"
				} else {
					out.Status = "completed"
				}
			} else {
				out.Status = "retry after 10 second"
			}
		}

	} else {
		out.Status = "in_progress"

	}
	return out
}
