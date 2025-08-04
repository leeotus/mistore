package models

type GoodsColor struct {
	Id         int    `json:"id"`
	ColorName  string `json:"color_name"`
	ColorValue string `json:"color_value"`
	Status     int    `json:"status"`
	Checked    bool   `gorm:"-"` // 忽略本字段
}

func (GoodsColor) TableName() string {
	return "goods_color"
}
