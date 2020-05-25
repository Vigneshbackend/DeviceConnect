package controller

import (
	"DeviceConnect/model"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var date string
var Accountnumber int
var year float64
var finalaccount int

func FindRecentlyAccessed(array []model.RecentlyAccessed) []model.SummaryAccountData {

	var Accounts model.AccountList
	var test []model.RecentlyAccessed

	Accounts.List = array

	refer := Accounts.List[0].Active_months_list[len(Accounts.List[0].Active_months_list)-1]

	split := strings.Split(refer, "-")

	yearstring := strings.Join(split, "")

	year, err := strconv.Atoi(yearstring)
	if err != nil {
		fmt.Println("----->err", err)
	}

	var monthcount [12]int

	if len(Accounts.List) > 3 {
		for n := 1; n <= 6; n++ {

			for i := 0; i < len(Accounts.List); i++ {

				for j := ((len(Accounts.List[i].Active_months_list)) - 1); j >= 0; j-- {

					currentyear := strings.ReplaceAll(Accounts.List[i].Active_months_list[j], "-", "")
					if yearstring == currentyear {
						monthcount[i] = monthcount[i] + 1
					}

				}

			}

			floatyear := float64(year + 1)
			if math.Remainder(floatyear, 1000) == 0 {

				currentstring := strings.Split(refer, "-")
				currentstring[1] = "12"
				referenceyear, err := strconv.Atoi(currentstring[0])
				if err != nil {
					fmt.Print("------>err", err)
				}
				referenceyear = referenceyear - 1
				currentstring[0] = strconv.Itoa(referenceyear)

			}
			split := strings.Split(refer, "-")
			referencemonth, err := strconv.Atoi(split[1])
			if err != nil {

				fmt.Print("--->err", err)
			}
			referencemonth = referencemonth - 1
			if referencemonth <= 9 {

				split[0] = split[0] + "0"
			}

			rfm := strconv.Itoa(referencemonth)
			split[1] = rfm

			yearstring = strings.Join(split, "")

		}

		// 	}

		// }

		a := monthcount[:]

		sort.Ints(a)

		highestcount := a[11]
		referint := 0
		for n := ((len(monthcount)) - 1); n >= 0; n-- {
			if monthcount[n] == highestcount {
				finalaccount = referint

				break

			}
			referint = referint + 1
		}

		reference := 0
		flag := 0
		for n := ((len(monthcount)) - 1); n >= 0; n-- {
			for i := ((len(a)) - 1); i >= len(a)-3; i-- {

				if a[i] == monthcount[n] && flag == 0 {

					test = append(test, Accounts.List[reference])
					flag = 1

				}

			}
			if len(test) == 3 {
				break
			}
			reference = reference + 1
			flag = 0
		}

		fmt.Println("responseee", monthcount, test[0])

		var responserefer []model.SummaryAccountData

		fmt.Println("heeeeeeeeelo", test[0].Company, len(test))

		var ref model.SummaryAccountData
		for n := 0; n < len(test); n++ {
			if(test[n].Latest_balance != 0 && test[n].Latest_balance_date != ""){

				if(test[n].Number!=""){
					ref.Company = test[n].Company + "-"+test[n].Number

				}else{
					ref.Company = test[n].Company 

				}

				ref.Latest_balance = test[n].Latest_balance
				ref.Latest_balance_date = test[n].Latest_balance_date
				ref.Number = test[n].Number
				ref.Type = test[n].Type
				responserefer = append(responserefer, ref)
			}
		}

		return responserefer

	} else {

		test = Accounts.List
		var ref model.SummaryAccountData

		var responserefer []model.SummaryAccountData
		for n := 0; n < len(test); n++ {
			if(test[n].Latest_balance != 0 && test[n].Latest_balance_date != ""){
				if(test[n].Number!=""){
					ref.Company = test[n].Company + "-"+test[n].Number

				}else{
					ref.Company = test[n].Company 

				}
				ref.Latest_balance = test[n].Latest_balance
				ref.Latest_balance_date = test[n].Latest_balance_date
				ref.Number = test[n].Number
				ref.Type = test[n].Type
				responserefer = append(responserefer, ref)

			}
		}
		return responserefer
	}
}
