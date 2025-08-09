package api

import (
	"mistore/src/db"
	"mistore/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @note 可以考虑写入ini文件中,之后从文件中读取数据
const DEFAULT_EXPIRATION = 60 * 60 // 单位:s

type BaseController struct{}

func (con BaseController) Render(c *gin.Context, tpl string, data map[string]interface{}) {

	//1、获取顶部导航
	topNavList := []models.Nav{}
	if hasTopNavList := db.RedisCache.GetWithCache(c.Request.Context(), "topNavList", &topNavList); !hasTopNavList {
		db.MySQLDB.Where("status=1 AND position=1").Find(&topNavList)
		db.RedisCache.SetWithContext(c.Request.Context(), "topNavList", topNavList, DEFAULT_EXPIRATION)
	}

	//2、获取分类的数据
	goodsCateList := []models.GoodsCate{}

	if hasGoodsCateList := db.RedisCache.GetWithCache(c.Request.Context(), "goodsCateList", &goodsCateList); !hasGoodsCateList {
		//https://gorm.io/zh_CN/docs/preload.html
		db.MySQLDB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		db.RedisCache.SetWithContext(c.Request.Context(), "goodsCateList", goodsCateList, DEFAULT_EXPIRATION)
	}

	//3、获取中间导航
	middleNavList := []models.Nav{}
	if hasMiddleNavList := db.RedisCache.GetWithCache(c.Request.Context(), "middleNavList", &middleNavList); !hasMiddleNavList {
		db.MySQLDB.Where("status=1 AND position=2").Find(&middleNavList)
		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
			relationIds := strings.Split(relation, ",")
			goodsList := []models.Goods{}
			db.MySQLDB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		db.RedisCache.SetWithContext(c.Request.Context(), "middleNavList", middleNavList, DEFAULT_EXPIRATION)
	}

	renderData := gin.H{
		"topNavList":    topNavList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
	}

	for key, v := range data {
		renderData[key] = v
	}

	c.HTML(http.StatusOK, tpl, renderData)

}
