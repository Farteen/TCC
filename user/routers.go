package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRouterRegister(r *gin.RouterGroup) {
	r.POST("/register", UserReigsterCreate)
}
/// 注册信息包含用户的头像,用户名,手机号,注册时间
func UserReigsterCreate(c *gin.Context) {
	// userName := c.Param("userName")
	// userAvatar := c.Param("userAvatar")
	// userTel := c.Param("userTel")
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}