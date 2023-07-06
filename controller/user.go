package controller

import (
	"kvm_backup/dao"
	"kvm_backup/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserLoginPage(c *gin.Context) {
	loginOut(c)
	c.HTML(http.StatusOK, "user.tpl", gin.H{})
}

func UserLoginPost(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	if name != "admin" && pwd != "123456" {
		c.HTML(http.StatusOK, "tip.tpl", gin.H{
			"tip": "账号密码错误",
		})
		return
	}
	middleware.SetUserSession(c,
		&middleware.User{
			ID:   10000,
			Name: name,
		},
		sessions.Options{
			Path:     "/",
			Domain:   dao.S("session_domain"),
			MaxAge:   600,   //存储有效时间，store引擎的生命周期也与其同步reddis 可通过ttl key查看
			Secure:   false, //Secure=true，那么这个cookie只能用https协议发送给服务器，要求协议用https
			HttpOnly: true,  //设置HttpOnly=true的cookie不能被js获取到
			SameSite: http.SameSiteDefaultMode,
		})
	c.Redirect(302, "/kvm_backup/task/")
}

func UserLogOut(c *gin.Context) {
	loginOut(c)
	c.Redirect(302, "/user/")
}
func loginOut(c *gin.Context) {
	middleware.DelUserSession(c,
		&sessions.Options{
			Path:     "/",
			Domain:   dao.S("session_domain"),
			Secure:   false,
			HttpOnly: true,
		})
}
