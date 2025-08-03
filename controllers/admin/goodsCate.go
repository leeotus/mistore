package admin

import (
	"mistore/src/db"
	"mistore/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsCateController struct {
	BaseController
}

func (ctl GoodsCateController) Index(ctx *gin.Context) {
	goodsCateList := []models.GoodsCate{}

	db.MySQLDB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	ctx.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (ctl GoodsCateController) Add(ctx *gin.Context) {
	// 获取顶级分类
	goodsCateList := []models.GoodsCate{}
	db.MySQLDB.Where("pid=0").Find(&goodsCateList)
	ctx.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (ctl GoodsCateController) DoAdd(ctx *gin.Context) {
	// 获取表单传过来的数据
	title := ctx.PostForm("title")
	pid, err1 := models.Str2Int(ctx.PostForm("pid"))
	link := ctx.PostForm("link")
	template := ctx.PostForm("template")
	subTitle := ctx.PostForm("sub_title")
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	sort, err2 := models.Str2Int(ctx.PostForm("sort"))
	status, err3 := models.Str2Int(ctx.PostForm("status"))

	if err1 != nil || err3 != nil {
		ctl.Error(ctx, "传入参数类型不正确", "/admin/goodsCate/add")
		return
	}
	if err2 != nil {
		ctl.Error(ctx, "排序值必须是整数!", "/admin/goodsCate/add")
		return
	}

	cateImgDir, _ := UploadImg(ctx, "cate_img")
	goods_cate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		Link:        link,
		Template:    template,
		SubTitle:    subTitle,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(models.TimeStamp()),
	}

	// 写入数据库
	err := db.MySQLDB.Create(&goods_cate).Error
	if err != nil {
		ctl.Error(ctx, "增加数据失败", "/admin/goodsCate/add")
		return
	}
	ctl.Success(ctx, "增加数据成功", "/admin/goodsCate")
}

func (ctl GoodsCateController) Edit(ctx *gin.Context) {
	// 获取要修改的数据
	id, err := models.Str2Int(ctx.Query("id"))
	if err != nil {
		ctl.Error(ctx, "传入参数错误", "/admin/goodsCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	db.MySQLDB.Find(&goodsCate)

	// 获取顶级分类
	goodsCateList := []models.GoodsCate{}
	db.MySQLDB.Where("pid=0").Find(&goodsCateList)
	ctx.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

func (ctl GoodsCateController) DoEdit(ctx *gin.Context) {
	id, err1 := models.Str2Int(ctx.PostForm("id"))
	title := ctx.PostForm("title")
	pid, err2 := models.Str2Int(ctx.PostForm("pid"))
	link := ctx.PostForm("link")
	template := ctx.PostForm("template")
	subTitle := ctx.PostForm("sub_title")
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	sort, err3 := models.Str2Int(ctx.PostForm("sort"))
	status, err4 := models.Str2Int(ctx.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		ctl.Error(ctx, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err3 != nil {
		ctl.Error(ctx, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := UploadImg(ctx, "cate_img")

	goodsCate := models.GoodsCate{Id: id}
	db.MySQLDB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}
	err := db.MySQLDB.Save(&goodsCate).Error
	if err != nil {
		ctl.Error(ctx, "修改失败", "/admin/goodsCate/edit?id="+models.Int2Str(id))
		return
	}
	ctl.Success(ctx, "修改成功", "/admin/goodsCate")
}

func (ctl GoodsCateController) Delete(ctx *gin.Context) {
	id, err := models.Str2Int(ctx.Query("id"))
	if err != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/goodsCate")
	} else {
		//获取我们要删除的数据
		goodsCate := models.GoodsCate{Id: id}
		db.MySQLDB.Find(&goodsCate)
		if goodsCate.Pid == 0 { //顶级分类
			goodsCateList := []models.GoodsCate{}
			db.MySQLDB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)
			if len(goodsCateList) > 0 {
				ctl.Error(ctx, "当前分类下面子分类，请删除子分类作以后再来删除这个数据", "/admin/goodsCate")
			} else {
				db.MySQLDB.Delete(&goodsCate)
				ctl.Success(ctx, "删除数据成功", "/admin/goodsCate")
			}
		} else { //操作 或者菜单
			db.MySQLDB.Delete(&goodsCate)
			ctl.Success(ctx, "删除数据成功", "/admin/goodsCate")
		}

	}
}
