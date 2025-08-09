package models

type User struct {
	Id       int    `json:"id"`
	Phone    string `json:"phone"`
	Password string `json:"pwd"`
	AddTime  int    `json:"add_time"`
	LastIp   string `json:"last_ip"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

func (User) TableName() string {
	return "user"
}
