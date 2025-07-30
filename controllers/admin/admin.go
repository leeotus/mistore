package admin

import (
	"mistore/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	controllers.BaseController
}

// TODO: 管理员主页
func (AdminController) AdminIndex(c *gin.Context) {
	c.String(http.StatusOK, "Admin Page")
}
