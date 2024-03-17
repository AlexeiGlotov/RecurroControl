package models

import "time"

type InfoKeyDashboard struct {
	CountActive int       `json:"count_active"`
	CountAll    int       `json:"count_all"`
	CountBan    int       `json:"count_ban"`
	Date        time.Time `json:"date" binding:"required"`
}
