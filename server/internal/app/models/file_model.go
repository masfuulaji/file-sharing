package models

type File struct {
	Id       string `json:"id"`
	FileName string `json:"file_name"`
	UserId   string `json:"user_id"`
	Size     int64  `json:"size"`
}
