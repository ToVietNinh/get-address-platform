package main

type District struct {
	DistrictID      int      `json:"DistrictID"`
	ProvinceID      int      `json:"ProvinceID"`
	DistrictName    string   `json:"DistrictName"`
	Code            string   `json:"Code"`
	Type            int      `json:"Type"`
	SupportType     int      `json:"SupportType"`
	NameExtension   []string `json:"NameExtension"`
	IsEnable        int      `json:"IsEnable"`
	UpdatedBy       int      `json:"UpdatedBy"`
	CreatedAt       string   `json:"CreatedAt"`
	UpdatedAt       string   `json:"UpdatedAt"`
	CanUpdateCOD    bool     `json:"CanUpdateCOD"`
	Status          int      `json:"Status"`
	PickType        int      `json:"PickType"`
	DeliverType     int      `json:"DeliverType"`
	ReasonCode      string   `json:"ReasonCode"`
	ReasonMessage   string   `json:"ReasonMessage"`
	OnDates         []string `json:"OnDates"`
	UpdatedEmployee int      `json:"UpdatedEmployee"`
	UpdatedIP       string   `json:"UpdatedIP"`
	UpdatedSource   string   `json:"UpdatedSource"`
	UpdatedDate     string   `json:"UpdatedDate"`
}

type ResponseForDistrict struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []District `json:"data"`
}
