package models

import "mistore/src/db"

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

/*
根据商品分类获取推荐商品

	@param {Number} cateId - 分类id
	@param {String} goodsType -  hot  best  new all
	@param {Number} limitNum -  数量
	1  表示顶级分类
		21
		23
		24
*/
func GetGoodsByCategory(cateId int, goodsType string, limitNum int) []Goods {

	//判断cateId 是否是顶级分类
	goodsCate := GoodsCate{Id: cateId}
	db.MySQLDB.Find(&goodsCate)
	var tempSlice []int
	if goodsCate.Pid == 0 { //顶级分类
		//获取顶级分类下面的二级分类
		goodsCateList := []GoodsCate{}
		db.MySQLDB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)

		for i := 0; i < len(goodsCateList); i++ {
			tempSlice = append(tempSlice, goodsCateList[i].Id)
		}

	}
	tempSlice = append(tempSlice, cateId)

	goodsList := []Goods{}
	where := "cate_id in ?"
	switch goodsType {
	case "hot":
		where += " AND is_hot=1"
	case "best":
		where += " AND is_best=1"
	case "new":
		where += " AND is_new=1"
	default:
		break
	}

	db.MySQLDB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Limit(limitNum).Order("sort desc").Find(&goodsList)
	return goodsList
}
