package models

import "time"

type User struct {
	MuID      		uint64     	`json:"id"`
	MuCode      	string     	`json:"code"`
	MuName      	string    	`json:"name"`
	MuEmail 		string		`json:"email"`
	MuDescription 	string		`json:"description"`
	MuStatus		uint8       `json:"status"`
	MuCreatedBy		uint64      `json:"created_by"`
	MuCreatedAt		time.Time   `json:"created_at"`
	MuUpdatedBy		uint64      `json:"updated_by"`
	MuUpdatedAt		time.Time   `json:"updated_at"`
}

type UserPIN struct {
	LopinID      		uint64     	`json:"id"`
	LopinPIN      		string     	`json:"pin"`
	LopinUserID     	uint64     	`json:"user_id"`
	LopinExpiredDate	time.Time   `json:"updated_at"`
	LopinCreatedBy		uint64      `json:"created_by"`
	LopinCreatedAt		time.Time   `json:"updated_at"`
	LopinStatus			uint8       `json:"status"`
}

type UserStatusInfo struct {
	LusiID      		uint64     	`json:"id"`
	LusiStatusBefore    uint8     	`json:"status_before"`
	LusiStatusCurrent	uint8       `json:"status_current"`
	LusiUserID     		uint64     	`json:"user_id"`
	LusiReason			string      `json:"reason"`
	LusiStatus			string		`json:"status_type"`
	LusiDuration		uint64		`json:"duration"`
	LusiCreatedBy		uint64		`json:"created_by"`
	LusiCreatedAt		time.Time	`json:"created_at"`
}

type UserSecretKey struct {
	LuskID      		uint64     	`json:"id"`
	LuskName    		string     	`json:"name"`
	LuskStatus			string      `json:"status"`
	LuskUserID     		uint64     	`json:"user_id"`
	LuskTypeID			uint8      	`json:"type_id"`
	LuskSource			string		`json:"source"`
	LuskCreatedBy		uint64		`json:"created_by"`
	LuskCreatedAt		time.Time	`json:"created_at"`
}

type UserOtp struct {
	LuoID      			uint64     	`json:"id"`
	LuoOtp    			uint32     	`json:"otp"`
	LuoStatus			uint8       `json:"status"`
	LuoUserID     		uint64     	`json:"user_id"`
	LuoExpiredDate		time.Time   `json:"expired_date"`
	LuoCreatedBy		uint64		`json:"created_by"`
	LuoCreatedAt		time.Time	`json:"created_at"`
}

type UserLogin struct {
	LulID      			uint64     	`json:"id"`
	LulUserID     		uint64     	`json:"user_id"`
	LulIpAddress		string		`json:"ip_address"`
	LulDeviceStatus		uint8   	`json:"device_status"`
	LulCreatedBy		uint64		`json:"created_by"`
	LulCreatedAt		time.Time	`json:"created_at"`
}

type UserToken struct {
	LutID      			uint64     	`json:"id"`
	LutUserID     		uint64     	`json:"user_id"`
	LutToken			string		`json:"token"`
	LutStatus			uint8       `json:"status"`
	LutCreatedBy		uint64		`json:"created_by"`
	LutCreatedAt		time.Time	`json:"created_at"`
}

type UserPassword struct {
	LupID      			uint64     	`json:"id"`
	LupUserID     		uint64     	`json:"user_id"`
	LupExpPasswordLink  string     	`json:"expired_link"`
	LupLinkExpDuration	uint64      `json:"expired_duration"`
	LupUserPassword		string		`json:"password"`
	LupIsExpired    	uint8    	`json:"is_expired"`
	LupStatus			uint8       `json:"status"`
	LupCreatedBy		uint64		`json:"created_by"`
	LupCreatedAt		time.Time	`json:"created_at"`
}

type UserRole struct {
	MurID      		uint64     	`json:"id"`
	MurCode      	string     	`json:"code"`
	MurUserID     	uint64     	`json:"user_id"`
	MurName      	string    	`json:"name"`
	MurStatus		uint8       `json:"status"`
	MurCreatedBy	uint64      `json:"created_by"`
	MurCreatedAt	time.Time   `json:"created_at"`
	MurUpdatedBy	uint64      `json:"updated_by"`
	MurUpdatedAt	time.Time   `json:"updated_at"`
}

type ResetPassword struct {
	UserID 			uint64		`json:"user_id"`
	Password 		string		`json:"password"`
	UpdatedBy		uint64		`json:"updated_by"`
	ExpPasswordLink string		`json:"exp_password_link"`
	ExpDuration     uint64		`json:"exp_duration"`
	IsExpired		uint8		`json:"is_expired"`
	CreatedBy		uint64		`json:"created_by"`
}