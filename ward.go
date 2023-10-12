package main

type Ward struct {
	WardCode      string   `json:"WardCode"`
	DistrictID    int      `json:"DistrictID"`
	WardName      string   `json:"WardName"`
	NameExtension []string `json:"NameExtension"`
	IsEnable      int      `json:"IsEnable"`
	CanUpdateCOD  bool     `json:"CanUpdateCOD"`
	UpdatedBy     int      `json:"UpdatedBy"`
	CreatedAt     string   `json:"CreatedAt"`
	UpdatedAt     string   `json:"UpdatedAt"`
	SupportType   int      `json:"SupportType"`
	PickType      int      `json:"PickType"`
	DeliverType   int      `json:"DeliverType"`
	Status        int      `json:"Status"`
	ReasonCode    string   `json:"ReasonCode"`
	ReasonMessage string   `json:"ReasonMessage"`
}

type ResponseForWard struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Ward `json:"data"`
}
