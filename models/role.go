package models

type Role struct {
	MroID      		int64     	`json:"id"`
	MroCode      	string     	`json:"code"`
	MroName      	string    	`json:"name"`
	MroDescription  string    	`json:"description"`
	MroStatus      	int64     	`json:"status"`
	MroCreatedBy	int64       `json:"created_by"`
	MroUpdatedBy	int64       `json:"updated_by"`
}
