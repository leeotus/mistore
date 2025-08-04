package models

type Goods struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	SubTitle      string  `json:"sub_title"`
	GoodsSn       string  `json:"goods_sn"`
	CateId        int     `json:"cate_id"`
	ClickCount    int     `json:"click_count"`
	GoodsNumber   int     `json:"goods_num"`
	Price         float64 `json:"price"`
	MarketPrice   float64 `json:"market_price"`
	RelationGoods string  `json:"relation"`
	GoodsAttr     string  `json:"attr"`
	GoodsVersion  string  `json:"version"`
	GoodsImg      string  `json:"img"`
	GoodsGift     string  `json:"gift"`
	GoodsFitting  string  `json:"fitting"`
	GoodsColor    string  `json:"color"`
	GoodsKeywords string  `json:"keywords"`
	GoodsDesc     string  `json:"desc"`
	GoodsContent  string  `json:"content"`
	IsDelete      int     `json:"is_delete"`
	IsHot         int     `json:"is_hot"`
	IsBest        int     `json:"is_best"`
	IsNew         int     `json:"is_new"`
	GoodsTypeId   int     `json:"type_id"`
	Sort          int     `json:"sort"`
	Status        int     `json:"status"`
	AddTime       int     `json:"add_time"`
}

func (Goods) TableName() string {
	return "goods"
}
