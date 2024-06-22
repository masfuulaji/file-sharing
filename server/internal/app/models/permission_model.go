package models

type Permission struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	FileId string `json:"file_id"`
}
