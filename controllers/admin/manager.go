package admin

import (
	"encoding/json"
	"mistore/src/db"
	"mistore/src/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @brief 管理员登录完成后的界面
 * @author leeotus (leeotus@163.com)
 */

type ManagerController struct {
	BaseController
}

// @brief 管理员主界面
func (ctl ManagerController) Index(ctx *gin.Context) {
	managerList := []models.Manager{}
	db.MySQLDB.Preload("Role").Find(&managerList)
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{"managerList": managerList})
}

// @brief 管理员的界面视图
func (ctl ManagerController) Add(ctx *gin.Context) {
	const cacheKey = "role_list"
	var roleList []models.Role

	// 查看是否在Redish有缓存
	val, err := db.RedisDB.Get(ctx.Request.Context(), cacheKey).Result()
	if err == nil && val != "" {
		// 缓存命中
		json.Unmarshal([]byte(val), &roleList)
	} else {
		// 缓存没有命中,前往MySQL里数据
		db.MySQLDB.Find(&roleList)
		if data, err := json.Marshal(roleList); err == nil {
			db.RedisDB.Set(ctx.Request.Context(), cacheKey, data, 3*time.Minute)
		}
	}

	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

// @brief 增加管理员操作
func (ctl ManagerController) DoAdd(ctx *gin.Context) {
	roleId, err := models.Str2Int(ctx.PostForm("role_id"))
	if err != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/manager/add")
		return
	}

	username := strings.TrimSpace(ctx.PostForm("username"))
	password := strings.TrimSpace(ctx.PostForm("password"))
	mobile := strings.TrimSpace(ctx.PostForm("mobile"))
	email := strings.TrimSpace(ctx.PostForm("email"))

	if len(username) < 2 || len(password) < 6 {
		ctl.Error(ctx, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}

	var m models.Manager
	result := db.MySQLDB.Where("username = ?", username).First(&m)
	if result.RowsAffected > 0 {
		// 管理员已存在，给出提示
		ctl.Error(ctx, "管理员已存在", "/admin/manager/add")
		return
	}

	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		RoleId:   roleId,
		AddTime:  int(models.TimeStamp()),
		Status:   0,
		IsSuper:  1,
	}

	err = db.MySQLDB.Create(&manager).Error
	if err != nil {
		ctl.Error(ctx, "增加管理员失败", "/admin/manager/add")
	}
	ctl.Success(ctx, "增加管理员成功", "/admin/manager")
}

// @brief 编辑管理员操作
func (ctl ManagerController) Edit(ctx *gin.Context) {
	//获取管理员
	id, err := models.Str2Int(ctx.Query("id"))
	if err != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	db.MySQLDB.Find(&manager)

	//获取所有的角色
	roleList := []models.Role{}
	db.MySQLDB.Find(&roleList)

	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

// @brief 删除管理员操作
func (ctl ManagerController) Delete(ctx *gin.Context) {
	id, err := models.Str2Int(ctx.Query("id"))
	if err != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/manager")
	} else {
		manager := models.Manager{Id: id}
		db.MySQLDB.Delete(&manager)
		ctl.Success(ctx, "删除数据成功", "/admin/manager")
	}
}
