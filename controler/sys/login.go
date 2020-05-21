package sys

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hold-door/models"
)

func Login(ctx *gin.Context) {
	var re models.ReturnModel

	username := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	if username != "panyuqing" {
		re.Code = 2
		re.Message = "账号信息错误"
		ctx.JSON(200, re)
		return
	}
	if pwd != "pyq.123987" {
		re.Code = 2
		re.Message = "口令错误"
		ctx.JSON(200, re)
		return
	}

	//user := models.User{ UserID: 110120, UerName: loignName}

	session := sessions.Default(ctx)
	session.Set("user", username)
	session.Save()

	re.Code = 1
	re.Data = "13501161788"
	ctx.JSON(200, re)
	return

}
