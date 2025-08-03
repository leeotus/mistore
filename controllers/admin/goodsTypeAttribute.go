package admin

import (
	"mistore/src/db"
	"mistore/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(c *gin.Context) {

	cateId, err := models.Str2Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	db.MySQLDB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList)
	//获取商品类型属性对应的类型

	goodsType := models.GoodsType{}
	db.MySQLDB.Where("id=?", cateId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})

}
func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	//获取当前商品类型属性对应的类型id

	cateId, err := models.Str2Int(c.Query("cate_id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	db.MySQLDB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"cateId":        cateId,
	})
}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err1 := models.Str2Int(c.PostForm("cate_id"))
	attrType, err2 := models.Str2Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := models.Str2Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?cate_id="+models.Int2Str(cateId))
		return
	}

	if err3 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/add?cate_id="+models.Int2Str(cateId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(models.TimeStamp()),
	}
	err := db.MySQLDB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?cate_id="+models.Int2Str(cateId))
	} else {
		con.Success(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute?id="+models.Int2Str(cateId))
	}

}

func (con GoodsTypeAttributeController) Edit(c *gin.Context) {

	//获取当前要修改数据的id
	id, err := models.Str2Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取当前id对应的商品类型属性
	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	db.MySQLDB.Find(&goodsTypeAttribute)

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	db.MySQLDB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeList":      goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	id, err1 := models.Str2Int(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err2 := models.Str2Int(c.PostForm("cate_id"))
	attrType, err3 := models.Str2Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err4 := models.Str2Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+models.Int2Str(id))
		return
	}
	if err4 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/edit?id="+models.Int2Str(id))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	db.MySQLDB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := db.MySQLDB.Save(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/goodsTypeAttribute/edit?id="+models.Int2Str(id))
		return
	}
	con.Success(c, "需改数据成功", "/admin/goodsTypeAttribute?id="+models.Int2Str(cateId))
}

func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	id, err1 := models.Str2Int(c.Query("id"))
	cateId, err2 := models.Str2Int(c.Query("cate_id"))
	if err1 != nil || err2 != nil {
		con.Error(c, "传入参数错误", "/admin/goodsType")
	} else {
		goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
		db.MySQLDB.Delete(&goodsTypeAttr)
		con.Success(c, "删除数据成功", "/admin/goodsTypeAttribute?id="+models.Int2Str(cateId))
	}
}
