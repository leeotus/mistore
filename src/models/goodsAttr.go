package models

type GoodsAttr struct {
	Id              int    `json:"id"`
	GoodsId         int    `json:"goods_id"`
	AttributeCateId int    `json:"attr_cate_id"`
	AttributeId     int    `json:"attr_id"`
	AttributeTitle  string `json:"attr_title"`
	AttributeType   int    `json:"attr_type"`
	AttributeValue  string `json:"attr_value"`
	Sort            int    `json:"sort"`
	AddTime         int    `json:"add_time"`
	Status          int    `json:"status"`
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
