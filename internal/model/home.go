package model

type HomeDataOutput struct {
	TravelKm           int              `json:"travelKm"           description:"行驶公里数"`
	ResidueElectricity int              `json:"residueElectricity" description:"剩余电量"`
	Version            []*VersionOutput `json:"version"            description:"新版本信息"`
}
