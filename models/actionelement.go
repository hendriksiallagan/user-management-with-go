package models

type Actionelement struct {
	MaeID      		int64     	`json:"id"`
	MaeCode      	string     	`json:"code"`
	MaeName      	string    	`json:"name"`
	MaeDescription  string    	`json:"description"`
	MaeScript  		string    	`json:"script"`
	MaeStatus      	int64     	`json:"status"`
	MaeCreatedBy	int64       `json:"created_by"`
	MaeUpdatedBy	int64       `json:"updated_by"`
}
