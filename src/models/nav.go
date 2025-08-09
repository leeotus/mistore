package models

type Nav struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Link       string  `json:"link"`
	Position   int     `json:"position"`
	IsOpennew  int     `json:"is_opennew"`
	Relation   string  `json:"relation"`
	Sort       int     `json:"sort"`
	Status     int     `json:"status"`
	AddTime    int     `json:"add_time"`
	GoodsItems []Goods `gorm:"-"`
}

func (Nav) TableName() string {
	return "nav"
}
