package middleware

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SID = "SID"

type User struct {
	ID   int
	Name string
}

func SetUserSession(c *gin.Context, u *User, s sessions.Options) {
	session := sessions.Default(c)
	session.Options(s)
	juser, _ := json.Marshal(&u)
	session.Set("user", juser)
	session.Save()
}

func getUserSession(c *gin.Context, u *User) {
	session := sessions.Default(c)
	juser, _ := session.Get("user").([]byte)
	json.Unmarshal(juser, &u)
}
func DelUserSession(c *gin.Context, s *sessions.Options) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	if cstr, _ := c.Cookie(SID); cstr != "" {
		c.SetCookie(SID, "", -60, s.Path, s.Domain, s.Secure, s.HttpOnly)
	}
}

func AuthUser(c *gin.Context) {
	var u = User{}
	getUserSession(c, &u)
	// fmt.Printf("sNew var =%#v\n", u)
	if u.ID > 0 {
		c.Set("u", u)
		c.Next()
	}
	c.Redirect(302, "/user/")
}
