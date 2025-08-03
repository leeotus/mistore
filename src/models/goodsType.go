package models

type GoodsType struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Status      int    `json:"status"`
	AddTime     int    `json:"add_time"`
}

func (GoodsType) TableName() string {
	return "goods_type"
}
