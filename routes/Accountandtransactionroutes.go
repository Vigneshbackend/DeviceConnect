package routes

import (
	"DeviceConnect/controller"
	"DeviceConnect/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"fmt"
)

func HandleAccountRoutes(router *mux.Router) {

	router.HandleFunc("/GetAccountsBalance", func(w http.ResponseWriter, r *http.Request) {
		var out model.SummaryAccountOutput
		var req model.Req
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		fmt.Println("errr",err)
		w.Header().Set("Content-Type", "application/json")

		if err == nil {
			out.Data = controller.GetAllAccount(req.Merchant_id)
			fmt.Println("checking",out.Data.Status)
			if out.Data.Status == "in_progress" {
				out.Code = 500
				out.Error = "please try again in 10 Seconds"
				w.Header().Set("Content-Type", "application/json")

				w.WriteHeader(http.StatusBadRequest)
			} else if out.Data.Status == "complete" {
				out.Code = 200

			}	else if out.Data.Status == "no_data" {
				out.Error = "No data found"
				out.Code = 200
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

			}   else if out.Data.Status != "completed" {
				out.Error = out.Data.Status
				out.Code = 400
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)

			}
			
		} else {

			out.Code = 400
			w.WriteHeader(http.StatusBadRequest)

			out.Error = "checkParameter"
			w.Header().Set("Content-Type", "application/json")

		}
		json.NewEncoder(w).Encode(out)

	}).Methods("POST")

	router.HandleFunc("/GetQRSummary", func(w http.ResponseWriter, r *http.Request) {
		var out model.QrSummaryOutput
		var req model.Req

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		fmt.Print("--->errr", err)
		w.Header().Set("Content-Type", "application/json")

		if err == nil {
			fmt.Println("calledd")

			out.Data = controller.Filter(req.Merchant_id)
			
			
			
			if out.Data.Status == "no_data" {
				out.Error = "No data found"
				out.Code = 200
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

			} else if out.Data.Status == "completed" {
				out.Code = 200
				w.WriteHeader(http.StatusOK)

				out.Error = ""
				w.Header().Set("Content-Type", "application/json")

			} else if out.Data.Status == "retry after 10 second" {
				out.Code = 400
				w.WriteHeader(http.StatusBadRequest)

				out.Error = "retry after 10 second"
				w.Header().Set("Content-Type", "application/json")
			}else if out.Data.Status != "completed" {
				out.Error = out.Data.Status
				out.Code = 400
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)

			}

		} else {

			out.Code = 400
			w.WriteHeader(http.StatusBadRequest)

			out.Error = "checkParameter"
			w.Header().Set("Content-Type", "application/json")

		}
		json.NewEncoder(w).Encode(out)

	}).Methods("POST")

}
