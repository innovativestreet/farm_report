package genratereport

import (
	"absolutetech/farm_report/config"
	"absolutetech/farm_report/global-lib/envutils"
	"absolutetech/farm_report/global-lib/httplib"
	"fmt"
)

func GenerateReport(farm_id int64, environment envutils.Env) error {

	apiURL := config.FarmAPIUrlPROD + fmt.Sprintf("?q=id:%d&wt=json&indent=true", farm_id)
	if environment == "testing" {
		apiURL = config.FarmAPIUrlUAT + fmt.Sprintf("?q=id:%d&wt=json&indent=true", farm_id)
	}

	responseObject, err := httplib.GetResponse(apiURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("API Response as struct %+v\n", responseObject)
	return nil
}
