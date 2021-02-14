package types

type KKM struct {
	ID                 int64  `json:"id"`
	ClientName         string `json:"clientName"`
	ModelName          string `json:"modelName"`
	SerialNumber       string `json:"serialNumber"`
	RegisterDate       string `json:"registerDate"`
	OFD                string `json:"ofd"`
	IsExcise           bool   `json:"isExcise"`
	SystemNO           string `json:"systemNo"`
	Type               string `json:"type"`
	FN                 int64  `json:"fn"`
	Address            string `json:"address"`
	EndDateOFD         string `json:"endDateOfd"`
	EndDateFN          string `json:"endDateFn"`
	InspectionDayCount int64  `json:"inspectionDayCount"` //ТО
	Comment            string `json:"comment"`
	ClientId           int64  `json:"clientId"`
	ModelId            int64  `json:"modelId"`
}

type Client struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ModelKKM struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
