package dtos

type JSONResponses struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type JSONSuccessResponses struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   string `json:"data"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type GetAllDataResponse struct {
	Result interface{} `json:"result"`
	GetAllDataRequest
}

type GetAllDataRequest struct {
	Page     int    `query:"page" json:"page"`
	Limit    int    `query:"limit" json:"limit"`
	SortBy   string `query:"sort_by" json:"sort_by"` // DESC / ASC
	Sort     string `query:"sort" json:"sort"`
	Search   string `query:"search" json:"search"`
	SearchBy string `query:"search_by" json:"search_by"`
}
