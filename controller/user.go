package controller

import (
	"kvm_backup/api"
	"kvm_backup/dao"
	"kvm_backup/model"
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
	Options := sessions.Options{
		Path:     "/",
		Domain:   dao.S("cookie_domain"),
		MaxAge:   600,   //存储有效时间，store引擎的生命周期也与其同步reddis 可通过ttl key查看
		Secure:   false, //Secure=true，那么这个cookie只能用https协议发送给服务器，要求协议用https
		HttpOnly: true,  //设置HttpOnly=true的cookie不能被js获取到
		SameSite: http.SameSiteDefaultMode,
	}
	if dao.S("auth_type") == "jwt" {
		api.SetUserToken(c, &model.AuthUser{ID: 10000, Name: name}, Options)
	} else {
		api.SetUserSession(c, &model.AuthUser{ID: 10000, Name: name}, Options)
	}
	c.Redirect(302, "/kvm_backup/task/")
}

func UserLogOut(c *gin.Context) {
	loginOut(c)
	c.Redirect(302, "/user/")
}
func loginOut(c *gin.Context) {
	if dao.S("auth_type") == "jwt" {
		api.DelUserToken(c,
			&sessions.Options{
				Path:     "/",
				Domain:   dao.S("cookie_domain"),
				Secure:   false,
				HttpOnly: true,
			})
	} else {
		api.DelUserSession(c,
			&sessions.Options{
				Path:     "/",
				Domain:   dao.S("cookie_domain"),
				Secure:   false,
				HttpOnly: true,
			})
	}

}
func AuthUser(c *gin.Context) {
	var u = model.AuthUser{}
	if dao.S("auth_type") == "jwt" {
		api.GetUserToken(c, &u)
	} else {
		api.GetUserSession(c, &u)
	}
	if u.ID > 0 {
		c.Set("u", u)
		c.Next()
	}
	c.Redirect(302, "/user/")
}
