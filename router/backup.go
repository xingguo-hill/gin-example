package router

import (
	"html/template"
	"kvm_backup/common"
	"kvm_backup/controller"
	"kvm_backup/dao"
	"kvm_backup/middleware"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	//设置gin日志输出模式
	if dao.S("log_type") == "file" {
		file, _ := os.Create(dao.S("log_file"))
		gin.DefaultWriter = file
	}

	//设置服务运行模式
	if dao.S("env") == "product" {
		gin.SetMode(gin.ReleaseMode)
	}
	//设置session存储模型
	store, err := redis.NewStore(10, "tcp", dao.S("session_redis_host"), "", []byte(dao.S("session_secret")))
	if err != nil {
		panic("redis connect err")
	}
	r.Use(sessions.Sessions(middleware.SID, store))

	//模版自定义函数
	r.SetFuncMap(template.FuncMap{
		"formatDate": common.FormatDateYs,
	})
}
func RouterInfo() {
	r.LoadHTMLGlob("tpl/*")
	r.Static("/css", "./static/css")
	// Simple group: v1
	v1 := r.Group("/kvm_backup/", middleware.AuthUser)
	{
		//task
		v1.GET("/task/:id", controller.TaskGetIndex)
		v1.GET("/task/", controller.TaskGetIndex)
		v1.POST("/task/", controller.TaskPostIndex)
		//log
		v1.Any("/log/", controller.LogGetIndex)
		v1.GET("/log/:bid", controller.LogGetIndex)
	}
	v1 = r.Group("/user/")
	{
		v1.GET("/", controller.UserLoginPage)
		v1.POST("/login/", controller.UserLoginPost)
		v1.GET("/logout/", controller.UserLogOut)
	}
}
func Run() {
	/**
	endless 在更新可执行文件的时候使用
	进行平滑重启 kill -1 pid
	强杀 kill -9 pid
	*/
	if dao.S("env") == "product" {
		endless.ListenAndServe("0.0.0.0:80", r)
	} else {
		r.Run("0.0.0.0:81")
	}

}
