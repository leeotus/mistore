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
 * @brief 轮播图实现
 * @author leeotus (leeotus@163.com)
 */

type RoleController struct {
	BaseController
}

/**
 * @brief 获取角色列表并显示
 * @note 这里role角色表的数据比较少,所以不用分页查询,不然需要做分页查询,每次只查询一页的数据保存到redis
 * @todo 查询过程可以封装成为一个任务提交到线程池里
 */
func (ctl RoleController) Index(ctx *gin.Context) {
	const cacheKey = "role_list"
	var roleList []models.Role

	// 1. 先查 Redis
	val, err := db.RedisDB.Get(ctx.Request.Context(), cacheKey).Result()
	if err == nil && val != "" {
		// 命中缓存
		// fmt.Println("查询角色命中缓存")
		json.Unmarshal([]byte(val), &roleList) // 反序列化
	} else {
		// 未命中缓存，查数据库
		// fmt.Println("查询角色未命中缓存")
		db.MySQLDB.Find(&roleList)
		// 写入 Redis，缓存5分钟
		if data, err := json.Marshal(roleList); err == nil {
			db.RedisDB.Set(ctx.Request.Context(), cacheKey, data, 3*time.Minute)
		}
	}

	ctx.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

/**
 * @brief 增加角色的页面
 */
func (ctl RoleController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

/**
 * @brief 执行增加操作
 */
func (ctl RoleController) DoAdd(ctx *gin.Context) {
	title := strings.TrimSpace(ctx.PostForm("title"))
	desc := strings.TrimSpace(ctx.PostForm("description"))

	if title == "" {
		ctl.Error(ctx, "角色标题不能为空", "/admin/role/add")
		return
	}
	// 需将收集到的信息存储到数据库中
	role := models.Role{}
	role.Title = title
	role.Description = desc
	role.Status = 1
	role.AddTime = int(models.TimeStamp())

	err := db.MySQLDB.Create(&role).Error
	if err != nil {
		// 增加角色的操作失败了,需要返回到原本的增加角色的界面
		// @note: 一般很少发生错误
		ctl.Error(ctx, "增加角色失败 请稍后重试!", "/admin/role/add")
		return
	}

	// 新增后更新 Redis 缓存
	const cacheKey = "role_list"
	var roleList []models.Role
	db.MySQLDB.Find(&roleList)
	if data, err := json.Marshal(roleList); err == nil {
		db.RedisDB.Set(ctx.Request.Context(), cacheKey, data, 3*time.Minute)
	}

	ctl.Success(ctx, "增加角色成功", "/admin/role")
}

/**
 * @brief 修改角色的数据
 */
func (ctl RoleController) Edit(ctx *gin.Context) {
	id, err := models.Str2Int(ctx.Query("id"))
	if err != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		db.MySQLDB.Find(&role)
		ctx.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
}

/**
 * @brief 修改数据并提交
 */
func (ctl RoleController) DoEdit(ctx *gin.Context) {
	id, err1 := models.Str2Int(ctx.PostForm("id"))
	if err1 != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/role")
		return
	}
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")

	if title == "" {
		ctl.Error(ctx, "角色的标题不能为空", "/admin/role/edit")
		return
	}

	role := models.Role{Id: id}
	db.MySQLDB.Find(&role)
	role.Title = title
	role.Description = description

	err2 := db.MySQLDB.Save(&role).Error
	if err2 != nil {
		ctl.Error(ctx, "修改数据失败", "/admin/role/edit?id="+models.Int2Str(id))
		return
	}

	// @note: 更新Redis缓存,(主动失效Active Invalidation)
	const cacheKey = "role_list"
	var roleList []models.Role
	db.MySQLDB.Find(&roleList)
	if data, err := json.Marshal(roleList); err == nil {
		db.RedisDB.Set(ctx.Request.Context(), cacheKey, data, 3*time.Minute)
	}

	ctl.Success(ctx, "修改数据成功", "/admin/role/edit?id="+models.Int2Str(id))
}

func (ctl RoleController) Delete(ctx *gin.Context) {
	id, err := models.Str2Int(ctx.Query("id"))
	if err != nil {
		ctl.Error(ctx, "传入数据错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	db.MySQLDB.Delete(&role)

	// 删除后更新 Redis 缓存
	const cacheKey = "role_list"
	var roleList []models.Role
	db.MySQLDB.Find(&roleList)
	if data, err := json.Marshal(roleList); err == nil {
		db.RedisDB.Set(ctx.Request.Context(), cacheKey, data, 3*time.Minute)
	}

	ctl.Success(ctx, "删除数据成功", "/admin/role")
}
