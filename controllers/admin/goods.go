package admin

import (
	"fmt"
	"math"
	"mistore/src/db"
	"mistore/src/models"
	"mistore/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

const GOODS_PAGE_SIZE = 5

func (con GoodsController) Index(c *gin.Context) {
	//当前页数
	page, _ := models.Str2Int(c.Query("page"))
	if page == 0 {
		page = 1
	}

	// 条件
	where := "is_delete=0"
	// 获取keyword
	keyword := c.Query("keyword")
	// fmt.Println("keyword:", keyword)
	suffix := " AND title like '%" + keyword + "%'"
	// fmt.Println("suffix:", suffix)
	if len(keyword) > 0 {
		where += suffix
	}
	// fmt.Println("where:", where)

	//每页查询的数量
	pageSize := GOODS_PAGE_SIZE
	goodsList := []models.Goods{}
	db.MySQLDB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	//获取总数量
	var count int64
	db.MySQLDB.Where(where).Table("goods").Count(&count)

	if len(goodsList) > 0 {
		c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
			"goodsList":  goodsList,
			"totalPages": math.Ceil(float64(count) / float64(pageSize)),
			"page":       page,
			"keyword":    keyword,
		})
	} else {
		if page != 1 {
			c.Redirect(http.StatusFound, "/admin/goods")
		} else {
			c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
				"goodsList":  goodsList,
				"totalPages": math.Ceil(float64(count) / float64(pageSize)),
				"page":       page,
				"keyword":    keyword,
			})
		}
	}
}

func (con GoodsController) Add(c *gin.Context) {
	//获取商品分类
	goodsCateList := []models.GoodsCate{}
	db.MySQLDB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取所有颜色信息
	goodsColorList := []models.GoodsColor{}
	db.MySQLDB.Find(&goodsColorList)

	//获取商品规格包装
	goodsTypeList := []models.GoodsType{}
	db.MySQLDB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

func (con GoodsController) DoAdd(c *gin.Context) {
	// 1. 获取表单传过来的数据,判断
	title := strings.TrimSpace(c.PostForm("title"))
	subTitle := strings.TrimSpace(c.PostForm("sub_title"))
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := models.Str2Int(c.PostForm("cate_id"))
	goodsNumber, _ := models.Str2Int(c.PostForm("goods_number"))

	// 注意小数点:
	marketPrice, _ := models.Str2Float(c.PostForm("market_price"))
	price, _ := models.Str2Float(c.PostForm("price"))

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")

	// 获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")
	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := models.Str2Int(c.PostForm("is_delete"))
	isHot, _ := models.Str2Int(c.PostForm("is_hot"))
	isBest, _ := models.Str2Int(c.PostForm("is_best"))
	isNew, _ := models.Str2Int(c.PostForm("is_new"))
	goodsTypeId, _ := models.Str2Int(c.PostForm("goods_type_id"))
	sort, _ := models.Str2Int(c.PostForm("sort"))
	status, _ := models.Str2Int(c.PostForm("status"))
	addTime := models.TimeStamp()

	// 获取颜色信息,把颜色转化为字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")

	// 上传图片,生成缩略图
	goodsImg, _ := UploadImg(c, "goods_img")

	// 增加商品数据
	goods := models.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    0,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       int(addTime),
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}
	// 上传数据库
	err := db.MySQLDB.Create(&goods).Error
	if err != nil {
		con.Error(c, "增加失败", "/admin/goods/add")
		return
	}

	imgList := c.PostFormArray("goods_image_list")
	// 增加图库信息
	var bgUpload = func(goodsId int, goodsImageList []string) {
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goodsId
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.TimeStamp())
			db.MySQLDB.Create(&goodsImgObj)
		}
	}
	utils.TaskPool.Submit(func() {
		bgUpload(goods.Id, imgList)
	})

	con.Success(c, "增加数据成功", "/admin/goods")
}

func (con GoodsController) Edit(c *gin.Context) {
	// 1、获取要修改的商品数据
	id, err := models.Str2Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/goods")
	}
	goods := models.Goods{Id: id}
	db.MySQLDB.Find(&goods)

	// 2、获取商品分类
	goodsCateList := []models.GoodsCate{}
	db.MySQLDB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	// 3、获取所有颜色 以及选中的颜色
	goodsColorSlice := strings.Split(goods.GoodsColor, ",")
	goodsColorMap := make(map[string]string)
	for _, v := range goodsColorSlice {
		goodsColorMap[v] = v
	}

	goodsColorList := []models.GoodsColor{}
	db.MySQLDB.Find(&goodsColorList)
	fmt.Println("-- goodsColorList:", goodsColorList)
	for i := 0; i < len(goodsColorList); i++ {
		if _, ok := goodsColorMap[models.Int2Str(goodsColorList[i].Id)]; ok {
			goodsColorList[i].Checked = true
		}
	}

	// 4、商品的图库信息
	goodsImageList := []models.GoodsImage{}
	db.MySQLDB.Where("goods_id=?", goods.Id).Find(&goodsImageList)

	// 5、获取商品类型
	goodsTypeList := []models.GoodsType{}
	db.MySQLDB.Find(&goodsTypeList)

	// 6、获取规格信息
	goodsAttr := []models.GoodsAttr{}
	db.MySQLDB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
	goodsAttrStr := ""

	for _, v := range goodsAttr {
		switch v.AttributeType {
		case 1:
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: </span> <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		case 2:
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		default:
			//获取当前类型对应的值
			goodsTypeArttribute := models.GoodsTypeAttribute{Id: v.AttributeId}
			db.MySQLDB.Find(&goodsTypeArttribute)
			attrValueSlice := strings.Split(goodsTypeArttribute.AttrValue, "\n")

			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += `<select name="attr_value_list">`
			for i := 0; i < len(attrValueSlice); i++ {
				if attrValueSlice[i] == v.AttributeValue {
					goodsAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[i], attrValueSlice[i])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[i], attrValueSlice[i])
				}
			}
			goodsAttrStr += `</select>`
			goodsAttrStr += `</li>`

		}
	}

	c.HTML(http.StatusOK, "admin/goods/edit.html", gin.H{
		"goods":          goods,
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
		"goodsAttrStr":   goodsAttrStr,
		"goodsImageList": goodsImageList,
	})
}

func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, err1 := models.Str2Int(c.Query("cateId"))
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err2 := db.MySQLDB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  goodsTypeAttributeList,
		})
	}
}

func (con GoodsController) DoEdit(c *gin.Context) {

	//1、获取表单提交过来的数据
	id, err1 := models.Str2Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/goods")
	}
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := models.Str2Int(c.PostForm("cate_id"))
	goodsNumber, _ := models.Str2Int(c.PostForm("goods_number"))
	//注意小数点
	marketPrice, _ := models.Str2Float(c.PostForm("market_price"))
	price, _ := models.Str2Float(c.PostForm("price"))
	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")
	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := models.Str2Int(c.PostForm("is_delete"))
	isHot, _ := models.Str2Int(c.PostForm("is_hot"))
	isBest, _ := models.Str2Int(c.PostForm("is_best"))
	isNew, _ := models.Str2Int(c.PostForm("is_new"))
	goodsTypeId, _ := models.Str2Int(c.PostForm("goods_type_id"))
	sort, _ := models.Str2Int(c.PostForm("sort"))
	status, _ := models.Str2Int(c.PostForm("status"))

	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")
	//3、修改数据
	goods := models.Goods{Id: id}
	db.MySQLDB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	//4、上传图片   生成缩略图
	goodsImg, err2 := UploadImg(c, "goods_img")
	if err2 == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
	}

	err3 := db.MySQLDB.Save(&goods).Error
	if err3 != nil {
		con.Error(c, "修改失败", "/admin/goods/edit?id="+models.Int2Str(id))
		return
	}

	//5、修改图库 增加图库信息
	bgEditImgs := func(goodsId int, goodsImageList []string) {
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goodsId
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.TimeStamp())
			db.MySQLDB.Create(&goodsImgObj)
		}
	}
	utils.TaskPool.Submit(func() {
		bgEditImgs(goods.Id, c.PostFormArray("goods_image_list"))
	})
	//6、修改规格包装  1、删除当前商品下面的规格包装   2、重新执行增加

	// 6.1删除当前商品下面的规格包装
	goodsAttrObj := models.GoodsAttr{}
	db.MySQLDB.Where("goods_id=?", goods.Id).Delete(&goodsAttrObj)
	//6.2、重新执行增加
	var bgTask = func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := models.Str2Int(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				db.MySQLDB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := models.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i]
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(models.TimeStamp())
				db.MySQLDB.Create(&goodsAttrObj)
			}
		}
	}
	utils.TaskPool.Submit(bgTask)
	con.Success(c, "修改数据成功", "/admin/goods")
}

func (con GoodsController) ImageUpload(c *gin.Context) {
	//上传图片
	imgDir, err := UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		// 开启OSS:
		var ossStatus = 1
		if ossStatus == 1 {
			c.JSON(http.StatusOK, gin.H{
				"link": models.Loader.AliOOSConfig.Domain + imgDir,
			})
		} else {
			//TODO
		}
	}
}

// 修改商品图库关联的颜色
func (con GoodsController) ChangeGoodsImageColor(c *gin.Context) {
	//获取图片id 获取颜色id
	goodsImageId, err1 := models.Str2Int(c.Query("goods_image_id"))
	colorId, err2 := models.Str2Int(c.Query("color_id"))
	goodsImage := models.GoodsImage{Id: goodsImageId}
	db.MySQLDB.Find(&goodsImage)
	goodsImage.ColorId = colorId
	err3 := db.MySQLDB.Save(&goodsImage).Error
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新失败",
			"success": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新成功",
			"success": true,
		})
	}

}

// 删除图库
func (con GoodsController) RemoveGoodsImage(c *gin.Context) {
	//获取图片id
	goodsImageId, err1 := models.Str2Int(c.Query("goods_image_id"))
	goodsImage := models.GoodsImage{Id: goodsImageId}
	err2 := db.MySQLDB.Delete(&goodsImage).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除失败",
			"success": false,
		})
	} else {
		//删除图片
		// os.Remove()
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除成功",
			"success": true,
		})
	}
}

// @brief 删除数据
func (con GoodsController) Delete(c *gin.Context) {
	id, err := models.Str2Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goods")
	} else {
		goods := models.Goods{Id: id}
		db.MySQLDB.Find(&goods)
		goods.IsDelete = 1
		goods.Status = 0
		db.MySQLDB.Save(&goods)

		// 获取上一页
		prevPage := c.Request.Referer()
		// fmt.Println("prevPage:", prevPage)
		if len(prevPage) > 0 {
			con.Success(c, "删除数据成功", prevPage)
		} else {
			con.Success(c, "删除数据成功", "/admin/goods")
		}
	}
}
