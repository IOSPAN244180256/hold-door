package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"hold-door/config"
	"net/http"
)

func RedisSessionStore() redis.Store {
	address := config.GetConfig().Get("redis_setting.host")
	token := config.GetConfig().Get("redis_setting.token")
	secret := config.GetConfig().Get("redis_setting.secret")
	db := config.GetConfig().Get("redis_setting.db")
	session_expire := config.GetConfig().Get("session_expire")

	store, err := redis.NewStoreWithDB(2, "tcp", address.(string), token.(string), db.(string), []byte(secret.(string)))
	if err != nil {
		panic(err)
	}

	store.Options(sessions.Options{
		MaxAge: session_expire.(int),
		Path:   "/",
	})
	return store
}

func ValidataAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			if c.FullPath() == "/sys/login" {
				c.Next()
				return
			}
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		} else {
			//刷新session
			session.Set("user", user)
			session.Save()
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		}

	}
}
