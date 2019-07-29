package models

type Element struct {
	MeID      		int64     	`json:"id"`
	MeCode      	string     	`json:"code"`
	MeName      	string    	`json:"name"`
	MeDescription  	string    	`json:"description"`
	MePageID  		int64    	`json:"page"`
	MeXpath  		string    	`json:"xpath"`
	MeActionElement string   	`json:"action_element"`
	MeStatus      	int64     	`json:"status"`
	MeCreatedBy		int64       `json:"created_by"`
	MeUpdatedBy		int64       `json:"updated_by"`
}
