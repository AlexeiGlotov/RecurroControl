package models

type Cheats struct {
	Id                uint16 `json:"id"`
	Name              string `json:"name"`
	Secure            string `json:"secure"`
	IsAllowedGenerate int    `json:"isAllowedGenerate"`
}
