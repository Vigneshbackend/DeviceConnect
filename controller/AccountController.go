package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"

	"DeviceConnect/model"
	"DeviceConnect/repo"

	"github.com/joho/godotenv"

	"fmt"
	"strconv"
)

var xapikey string
var version int
var salt string

type AccountController struct {
	DB *gorm.DB
}

//init func get default configuration values
func Init() {
	err := godotenv.Load(".env")
	fmt.Println(err)

	if err != nil {
		fmt.Println("Cannot find .env file")
	}
	xapikey = os.Getenv("SERVER_API_KEY")
	version, _ = strconv.Atoi(os.Getenv("DC_PREDICTORS_VERSION"))
	salt = os.Getenv("Salt")
}

func GetAllAccount(name string) model.SummaryAccount {

	var out model.SummaryAccount
	if name != "" {
		Init()
		var refer model.Reqbody
		refer.Customer_id = name
		refer.Version = version
		refer.Salt = repo.GetSaltForCustomer(name)
		refer.Num_records = "1"

		b, err := json.Marshal(refer)
		fmt.Println("--------->b", err)

		if err == nil {
			client := &http.Client{}

			apiUrl := "https://insights.finbox.in/staging/accounts"
			u, _ := url.ParseRequestURI(apiUrl)
			urlStr := u.String()
			req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(b))
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("X-api-key", xapikey)
			req.Header.Set("Content-Type", `application/json`)

			resp, err := client.Do(req)

			if err != nil {
				fmt.Println("-------------->", err)
			}
			defer resp.Body.Close()
			var response model.AccountOutput
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println("bodyy",body)
			// err = json.Unmarshal([]byte(body), &out.Data)

			if err != nil {
				fmt.Println(err)
				// out.Error = "failed"
				// out.Code = 500
			} else {
				var check model.NoData
				err = json.Unmarshal([]byte(body), &check)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("responseee",check)
				fmt.Println("verify",check.Status)
				

					err = json.Unmarshal([]byte(body), &response)
					fmt.Println("responseee",response)
					if err != nil {
						fmt.Println(err)
					}
					var finalizesummary model.SummaryAccount
					finalizesummary.Customerid = response.Customerid
					finalizesummary.Date_processed = response.Date_processed
					finalizesummary.Date_requested = response.Date_requested
					finalizesummary.Message = response.Message
					finalizesummary.Requestid = response.Requestid
					finalizesummary.Status = response.Status

					if response.Status != "in_progress" && response.Status!="no_data" {
						finalizesummary.Data = FindRecentlyAccessed(response.Data)

					}
					out = finalizesummary
					if err != nil {
						fmt.Println(err)
					}
					
			}
		}
	} else {

	}
	return out
}
