package todo

import "database/sql"

type StLicenseKeys struct {
	Id           uint16         `json:"id"`
	License_key  string         `json:"license_key"`
	Cheat        string         `json:"cheat"`
	Time         string         `json:"time"`
	Seller       string         `json:"seller"`
	Status       int            `json:"status"`
	Banned       int            `json:"banned"`
	Purchased_id string         `json:"purchased_id"`
	Hwid         string         `json:"hwid"`
	HwidK        string         `json:"hwidK"`
	End          string         `json:"end"`
	Createdata   string         `json:"createdata"`
	Activedata   sql.NullString `json:"activedata"`
	Owner        sql.NullString `json:"owner"`
}

type User struct {
	Id           int            `json:"id" db:"id"`
	Login        string         `json:"login" binding:"required"`
	Password     string         `json:"password" binding:"required"`
	Access_key   string         `json:"Access_key"`
	Job          string         `json:"job"`
	Status       int            `json:"status"`
	Banned       int            `json:"banned"`
	Cankeygen    int            `json:"cankeygen"`
	Keygen       int            `json:"keygen"`
	Keyactivated int            `json:"keyactivated"`
	Owner        sql.NullString `json:"owner"`
}

type StCheats struct {
	Id       uint16         `json:"id"`
	Name     string         `json:"name"`
	Secure   string         `json:"secure"`
	Cangen   int            `json:"cangen"`
	Procname sql.NullString `json:"procname"`
	X64      sql.NullString `json:"x64"`
	Path     sql.NullString `json:"path"`
	Path2    sql.NullString `json:"path2"`
	Dll      sql.NullString `json:"dll"`
	Dlltest  sql.NullString `json:"dll_test"`
}
