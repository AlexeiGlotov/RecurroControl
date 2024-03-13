package models

const (
	Admin        = "admin"
	Distributors = "distributors"
	Reseller     = "reseller"
	Salesman     = "salesman"
)

type User struct {
	Id    int    `json:"id" db:"id"`
	Login string `json:"login" binding:"required"`
	//Password      string `json:"password_hash" binding:"required"`
	Role          string `json:"role"`
	KeysGenerated int    `json:"keys_generated"`
	KeysActivated int    `json:"keys_activated"`
	Banned        int    `json:"banned"`
	Owner         string `json:"owner" binding:"required"`
	IsDeleted     int    `json:"is_deleted"`
}

type SignUpInput struct {
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required"`
	Access_Key string `json:"access_key" binding:"required"`
	Owner      string `json:"owner" binding:""`
	Role       string `json:"role" db:"role"`
}

type AccessKey struct {
	Id        int     `json:"id" db:"id"`
	AccessKey string  `json:"access_key" db:"access_key"`
	Owner     string  `json:"owner" db:"owner"`
	IsLogin   *string `json:"is_login" db:"is_login"`
	Role      string  `json:"role" db:"role" binding:"required"`
}
