package api

import (
	"encoding/json"
	"kvm_backup/dao"
	"kvm_backup/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const SID = "SID"

// 设置session存储模型
func GetSessionStore() redis.Store {
	store, err := redis.NewStore(10, "tcp", dao.S("session_redis_host"), "", []byte(dao.S("session_secret")))
	if err != nil {
		panic("redis connect err")
	}
	return store
}

func SetUserSession(c *gin.Context, u *model.AuthUser, s sessions.Options) {
	session := sessions.Default(c)
	session.Options(s)
	juser, _ := json.Marshal(u)
	session.Set("user", juser)
	session.Save()
}

func GetUserSession(c *gin.Context, u *model.AuthUser) {
	session := sessions.Default(c)
	juser, _ := session.Get("user").([]byte)
	json.Unmarshal(juser, &u)
}
func DelUserSession(c *gin.Context, s *sessions.Options) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	if _, err := c.Cookie(SID); err == nil {
		c.SetCookie(SID, "", -60, s.Path, s.Domain, s.Secure, s.HttpOnly)
	}
}
