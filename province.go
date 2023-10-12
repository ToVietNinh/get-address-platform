package main

type Province struct {
	ProvinceID      int      `json:"ProvinceID"`
	ProvinceName    string   `json:"ProvinceName"`
	CountryID       int      `json:"CountryID"`
	NameExtension   []string `json:"NameExtension"`
	IsEnable        int      `json:"IsEnable"`
	RegionID        int      `json:"RegionID"`
	RegionCPN       int      `json:"RegionCPN"`
	UpdatedBy       int      `json:"UpdatedBy"`
	CanUpdateCOD    bool     `json:"CanUpdateCOD"`
	Status          int      `json:"Status"`
	UpdatedIP       string   `json:"UpdatedIP"`
	UpdatedEmployee int      `json:"UpdatedEmployee"`
	UpdatedSource   string   `json:"UpdatedSource"`
}

type ResponseForProvince struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Province `json:"data"`
}
