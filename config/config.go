package config

type Farm struct {
	Response struct {
		NumFound int `json:"numFound,omitempty"`
		Start    int `json:"start,omitempty"`
		Docs     []struct {
			BoundaryCoord             string  `json:"boundary_coord,omitempty"`
			BatchCount                int     `json:"batch_count,omitempty"`
			StateName                 string  `json:"state_name,omitempty"`
			RemainingAcerage          float64 `json:"remaining_acerage,omitempty"`
			ID                        string  `json:"id,omitempty"`
			ExpectedProduce           float64 `json:"expected_produce,omitempty"`
			Pincode                   string  `json:"pincode,omitempty"`
			LgdStateID                int     `json:"lgd_state_id,omitempty"`
			Lng                       float64 `json:"lng,omitempty"`
			UserID                    int     `json:"user_id,omitempty"`
			VarietyName               string  `json:"variety_name,omitempty"`
			BatchID                   int     `json:"batch_id,omitempty"`
			SubDistrictName           string  `json:"sub_district_name,omitempty"`
			FarmUpdatedAt             string  `json:"farm_updated_at,omitempty"`
			FarmScore                 int     `json:"farm_score,omitempty"`
			IsBatchActive             bool    `json:"is_batch_active,omitempty"`
			IsActive                  bool    `json:"is_active,omitempty"`
			TotalAcerage              float64 `json:"total_acerage,omitempty"`
			BatchName                 string  `json:"batch_name,omitempty"`
			Stage                     string  `json:"stage,omitempty"`
			VillageName               string  `json:"village_name,omitempty"`
			FarmName                  string  `json:"farm_name,omitempty"`
			BatchHealth               int     `json:"batch_health,omitempty"`
			CommodityImage            string  `json:"commodity_image,omitempty"`
			BatchStatus               string  `json:"batch_status,omitempty"`
			InsuranceStatus           int     `json:"insurance_status,omitempty"`
			FranchiseType             string  `json:"franchise_type,omitempty"`
			FarmStatus                string  `json:"farm_status,omitempty"`
			FarmerID                  int     `json:"farmer_id,omitempty"`
			SopAdherence              float64 `json:"sop_adherence,omitempty"`
			Lat                       float64 `json:"lat,omitempty"`
			Group                     int     `json:"group,omitempty"`
			CommodityName             string  `json:"commodity_name,omitempty"`
			Acerage                   float64 `json:"acerage,omitempty"`
			UniqueKey                 string  `json:"unique_key,omitempty"`
			ActualProduce             float64 `json:"actual_produce,omitempty"`
			SubDistrictCode           int     `json:"sub_district_code,omitempty"`
			DistrictCode              int     `json:"district_code,omitempty"`
			ActualYieldPerAcre        float64 `json:"actual_yield_per_acre,omitempty"`
			CommodityVarietyID        int     `json:"commodity_variety_id,omitempty"`
			BatchUpdatedAt            string  `json:"batch_updated_at,omitempty"`
			DistrictName              string  `json:"district_name,omitempty"`
			VillageCode               int     `json:"village_code,omitempty"`
			Gh                        int     `json:"gh,omitempty"`
			Mobile                    string  `json:"mobile,omitempty"`
			IrrigationType            string  `json:"irrigation_type,omitempty"`
			CommodityID               int     `json:"commodity_id,omitempty"`
			Location                  string  `json:"location,omitempty"`
			Username                  string  `json:"username,omitempty"`
			Version                   int64   `json:"_version_,omitempty"`
			SrAssigneeName            string  `json:"sr_assignee_name,omitempty"`
			ExpectedDeliveryDate      string  `json:"expected_delivery_date,omitempty"`
			Locality                  string  `json:"locality,omitempty"`
			AssigneeName              string  `json:"assignee_name,omitempty"`
			SubfarmerMobileNo         string  `json:"subfarmer_mobile_no,omitempty"`
			AssigneeMobile            string  `json:"assignee_mobile,omitempty"`
			CommodityImage96PxIcon    string  `json:"commodity_image_96px_icon,omitempty"`
			SoilReportCount           int     `json:"soil_report_count,omitempty"`
			ExpectedYieldDeliveryDate string  `json:"expected_yield_delivery_date,omitempty"`
			SubfarmerName             string  `json:"subfarmer_name,omitempty"`
			SrAssigneeMobile          string  `json:"sr_assignee_mobile,omitempty"`
		} `json:"docs,omitempty"`
	} `json:"response,omitempty"`
}

var SurveyDictionary = map[int64]int64{
	115680: 186,
	115704: 68,
}

var FarmAPIUrlUAT = "http://uat.itrade.ag:8983/solr/farm_data/select"
var FarmAPIUrlPROD = "http://itrade.ag:8983/solr/farm_data/select"

var FarmWeatherAPIUrl = "https://whapiuat.igrow.ag/weather/earthengine/getSoilParams?location={}&deviceId={}&farm_id={}"
