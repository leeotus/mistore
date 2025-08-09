package models

// @todo 用户注册需要引入短信服务,这里没有集成短信服务的功能,留做以后完成
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
