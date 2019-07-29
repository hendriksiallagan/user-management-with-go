package models

type GetResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
	Data		interface{}  	`json:"data"`
}

type GetDetailResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
	Data		interface{}     `json:"data"`
}

type DataResponse struct {
	Page        int           `json:"page"`
	TotalData   int           `json:"total_data"`
	Data		interface{}   `json:"data"`
}

type Response struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
}