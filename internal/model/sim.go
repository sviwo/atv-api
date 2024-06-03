package model

type SimDataTrafficOutput struct {
	TotalDataTraffic   string `json:"totalDataTraffic"          dc:""`
	ConsumeDataTraffic string `json:"consumeDataTraffic"        dc:""`
	SurplusDataTraffic string `json:"surplusDataTraffic"        dc:""`
}
