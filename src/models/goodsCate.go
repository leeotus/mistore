package models

// @note 最后一个字段GoodsCateItems用来实现分类的层级结构,即一个分类可以有子分类
// @note Pid字段表示父分类的ID, 而GoodsCateIterms用来存储当前分类的所有子分类
type GoodsCate struct {
	Id             int         `json:"id"`
	Title          string      `json:"title"`
	CateImg        string      `json:"cate_img"`
	Link           string      `json:"link"`
	Template       string      `json:"template"`
	Pid            int         `json:"pid"`
	SubTitle       string      `json:"sub_title"`
	Keywords       string      `json:"keywords"`
	Description    string      `json:"desc"`
	Sort           int         `json:"sort"`
	Status         int         `json:"status"`
	AddTime        int         `json:"add_time"`
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:Id"`
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
