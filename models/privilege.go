package models

type Privilege struct {
	MrprID      		int     	`json:"id"`
	MrprCode      		string     	`json:"code"`
	MrprName      		string    	`json:"name"`
	MrprTypeID  		int    		`json:"type_id"`
	MrprRoleID  		int    		`json:"role_id"`
	MrprMenuID  		int    		`json:"menu_id"`
	MrprApiID      		int     	`json:"api_id"`
	MrprPageID			int       	`json:"page_id"`
	MrprElementID		int       	`json:"element_id"`
	MrprActionElementID string     	`json:"action_element_id"`
	MrprStatus			int       	`json:"status"`
	MrprCreatedBy		int       	`json:"created_by"`
	MrprUpdatedBy		int       	`json:"updated_by"`
}
