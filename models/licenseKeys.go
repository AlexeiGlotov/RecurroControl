package models

type LicenseKeys struct {
	Id           int    `json:"id" db:"id"`
	LicenseKeys  string `json:"license_key" binding:"required"`
	TTLCheat     int    `json:"ttl_cheat" binding:"required"`
	Holder       string `json:"holder" binding:"required"`
	Creator      string `json:"creator" binding:"required"`
	DateCreation string `json:"date_creation"`
}
