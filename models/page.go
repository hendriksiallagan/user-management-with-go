package models

type Page struct {
	MpID      		int64     	`json:"id"`
	MpCode      	string     	`json:"code"`
	MpName      	string    	`json:"name"`
	MpDescription   string    	`json:"description"`
	MpFilepath 		string		`json:"file_path"`
	MpStatus      	int64     	`json:"status"`
	MpCreatedBy		int64       `json:"created_by"`
	MpUpdatedBy		int64       `json:"updated_by"`
}
