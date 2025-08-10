package api

import (
	"mistore/src/db"
	"mistore/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hunterhug/go_image"
	"gorm.io/gorm"
)

type DefaultController struct {
	BaseController
}

// @todo: 首页里的各种数据保存到Redis中
func (con DefaultController) Index(c *gin.Context) {
	// 1.获取顶部导航
	topNavList := []models.Nav{}
	err := db.RedisCache.GetMultiLevel(c.Request.Context(), "topNavList", &topNavList)
	if err != nil {
		// L1+L2都没有命中,查询数据库
		db.MySQLDB.Where("status=1 AND position=1").Find(&topNavList)
		db.RedisCache.SetMultiLevel(c.Request.Context(), "topNavList", topNavList, db.DEFAULT_REDIS_CACHE_EXPIRATION, db.DEFAULT_LOCAL_CACHE_EXPIRATION)
	}

	// 2.获取轮播图数据
	focusList := []models.Focus{}
	err = db.RedisCache.GetMultiLevel(c.Request.Context(), "focusList", &focusList)
	if err != nil {
		db.MySQLDB.Where("status=1 AND focus_type=1").Find(&focusList)
		db.RedisCache.SetMultiLevel(c.Request.Context(), "focusList", focusList, db.DEFAULT_REDIS_CACHE_EXPIRATION, db.DEFAULT_LOCAL_CACHE_EXPIRATION)
	}

	// 3.获取分类的数据
	goodsCateList := []models.GoodsCate{}
	err = db.RedisCache.GetMultiLevel(c.Request.Context(), "goodsCateList", &goodsCateList)
	if err != nil {
		db.MySQLDB.Where("pid=0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)
		db.RedisCache.SetMultiLevel(c.Request.Context(), "goodsCateList", goodsCateList, db.DEFAULT_REDIS_CACHE_EXPIRATION, db.DEFAULT_LOCAL_CACHE_EXPIRATION)
	}

	// 4.获取中间导航
	middleNavList := []models.Nav{}
	err = db.RedisCache.GetMultiLevel(c.Request.Context(), "middleNavList", &middleNavList)
	if err != nil {
		db.MySQLDB.Where("status=1 AND position=2").Find(&middleNavList)

		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",")
			relationIds := strings.Split(relation, ",")
			goodsList := []models.Goods{}
			db.MySQLDB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		db.RedisCache.SetMultiLevel(c.Request.Context(), "middleNavList", middleNavList, db.DEFAULT_REDIS_CACHE_EXPIRATION, db.DEFAULT_LOCAL_CACHE_EXPIRATION)
	}

	// 手机
	phoneList := []models.Goods{}
	err = db.RedisCache.GetMultiLevel(c.Request.Context(), "phoneList", &phoneList)
	if err != nil {
		phoneList = models.GetGoodsByCategory(23, "best", 8)
		db.RedisCache.SetMultiLevel(c.Request.Context(), "phoneList", phoneList, db.DEFAULT_REDIS_CACHE_EXPIRATION, db.DEFAULT_LOCAL_CACHE_EXPIRATION)
	}
	// 配件
	otherList := []models.Goods{}
	err = db.RedisCache.GetMultiLevel(c.Request.Context(), "otherList", &otherList)
	if err != nil {
		otherList = models.GetGoodsByCategory(9, "all", 1)
		db.RedisCache.SetMultiLevel(c.Request.Context(), "otherList", otherList, db.DEFAULT_REDIS_CACHE_EXPIRATION, db.DEFAULT_LOCAL_CACHE_EXPIRATION)
	}

	c.HTML(http.StatusOK, "api/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		"otherList":     otherList,
	})
}

func (con DefaultController) Thumbnail1(c *gin.Context) {
	//按宽度进行比例缩放，输入输出都是文件
	//filename string, savepath string, width int
	filename := "static/upload/0.png"
	savepath := "static/upload/0_600.png"
	err := go_image.ScaleF2F(filename, savepath, 600)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail1 成功")
}
