package models

type Role struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      int    `json:"status"`
	AddTime     int    `json:"addtime"`
}

// @brief 表名
func (Role) TableName() string {
	return "role"
}
