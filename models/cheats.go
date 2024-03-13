package models

type Cheats struct {
	Id                uint16 `json:"id"`
	Name              string `json:"name" binding:"required"`
	Secure            string `json:"secure" binding:"required"`
	IsAllowedGenerate *int   `json:"is_allowed_generate" binding:"required"`
}
