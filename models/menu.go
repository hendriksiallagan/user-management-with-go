package models

type Menu struct {
	MmID      		int64     	`json:"id"`
	MmCode      	string     	`json:"code"`
	MmName      	string    	`json:"name"`
	MmDescription  	string    	`json:"description"`
	MmUrl  			string    	`json:"url"`
	MmHeaderID  	int64    	`json:"menu_header"`
	MmStatus      	int64     	`json:"status"`
	MmCreatedBy		int64       `json:"created_by"`
	MmUpdatedBy		int64       `json:"updated_by"`
}
