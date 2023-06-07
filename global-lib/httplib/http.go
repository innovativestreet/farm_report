package httplib

import (
	"absolutetech/farm_report/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetResponse(apiURL string) (config.Farm, error) {
	// req, err := http.NewRequest("GET", apiURL, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Authorization", config.AuthKey)
	// req.Header.Add("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()
	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	resp, err := http.Get(apiURL)
	if err != nil {
		return config.Farm{}, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return config.Farm{}, err
	}

	var responseObject config.Farm
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return config.Farm{}, err
	}
	return responseObject, nil
}
