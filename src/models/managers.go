package models

/**
 * @brief 管理员权限的数据表
 */
type Manager struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	RoleId   int    `json:"roleid"`
	AddTime  int    `json:"addtime"`
	IsSuper  int    `json:"super"`
	Role     Role   `gorm:"foreignKey:RoleId"`
}

func (Manager) TableName() string {
	return "manager"
}
