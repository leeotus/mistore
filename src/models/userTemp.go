package models

type UserTemp struct {
	Id        int    `json:"id"`
	Ip        string `json:"ip"`
	Phone     string `json:"phone"`
	SendCount int    `json:"send_count"`
	AddDay    string `json:"add_day"`
	AddTime   int    `json:"add_time"`
	Sign      string `json:"sign"`
}

func (UserTemp) TableName() string {
	return "user_temp"
}
