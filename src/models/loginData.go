package models

import (
	"context"
	"fmt"
	"mistore/src/db"
	"mistore/src/verify"
	"time"
)

// @note 用户登录表单数据
type LoginData struct {
	CapId     string
	VerifyVal string
	Username  string
	Password  string
}

/**
 * @brief 验证验证码输入是否正确
 */
func (data *LoginData) VerifyCaptcha() bool {
	return verify.VerifyCaptcha(data.CapId, data.VerifyVal)
}

/**
 * @brief 登录操作,查看数据库内是否有该用户,以及密码输入是否正确
 * @return bool 成功登录返回true,否则返回false
 * @return *Manager 查询到的用户数据
 */
func (data *LoginData) LogIn() (bool, *Manager) {
	userinfo := []Manager{}

	// 密码需要做加密操作
	pwd := Md5(data.Password)
	res := db.DB.Where("username=? AND password=?", data.Username, pwd).Find(&userinfo)
	if res.RowsAffected <= 0 || len(userinfo) <= 0 {
		return false, nil
	}
	return true, &userinfo[0]
}

/**
 * @brief 保存用户的Session
 * @param ctx 上下文对象
 * @param sessionId 会话id
 * @param sessionVal 保存在Redis里会话的值
 * @param login_time 记录登录时间戳
 * @param super_user 是否是管理员用户
 * @note 用户登录之后用于保存用户的Session
 * @note 后续用户的操作与执行需要判断Session是否过期,过期则需要重新登录,否则操作执行并更新Session
 */
func (data *LoginData) SaveSession(ctx context.Context, sessionId string, sessionVal string, super_user int) error {
	key := fmt.Sprintf("session:%s", sessionId)

	err := db.RedisDB.HSet(ctx, key,
		"session_val", sessionVal,
		"login_time", time.Now().Unix(),
		"super_user", super_user).Err()
	if err != nil {
		return err
	}
	return db.RedisDB.Expire(ctx, key, 1*time.Hour).Err()
}
