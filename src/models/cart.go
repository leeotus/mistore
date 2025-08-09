package models

type Cart struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Price        float64 `json:"price"`
	GoodsVersion string  `json:"goods_version"`
	Uid          int     `json:"uid"`
	Num          int     `json:"num"`
	GoodsGift    string  `json:"goods_gift"`
	GoodsFitting string  `json:"goods_fitting"`
	GoodsColor   string  `json:"goods_color"`
	GoodsImg     string  `json:"goods_img"`
	GoodsAttr    string  `json:"goods_attr"`
	Checked      bool    `json:"checked"`
}

// @brief 判断购物车里面有没有当前数据
func HasCartData(cartList []Cart, currentData Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
			return true
		}
	}
	return false
}
