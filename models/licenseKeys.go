package models

type InputCreate struct {
	CountKeys int    `json:"count_keys" binding:"required"`
	TTLCheat  int    `json:"ttl_cheat" binding:"required"`
	Holder    string `json:"holder" binding:"required"`
	Cheat     string `json:"cheat" binding:"required"`
}

type LicenseKeys struct {
	Id             int     `json:"id" db:"id"`
	LicenseKeys    string  `json:"license_key" binding:"required"`
	Cheat          string  `json:"cheat" binding:"required"`
	TTLCheat       int     `json:"ttl_cheat" binding:"required"`
	Holder         string  `json:"holder" binding:"required"`
	Creator        string  `json:"creator" binding:"required"`
	DateCreation   string  `json:"date_creation"`
	DateActivation *string `json:"date_activation"`
	HWID           *string `json:"hwid"`
	HWIDK          *string `json:"hwidk"`
	Banned         int     `json:"banned"`
	IsDeleted      int     `json:"is_deleted"`
}
