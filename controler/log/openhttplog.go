package log

import (
	"github.com/gin-gonic/gin"
	GrpcOpenHttpLogPkg "hold-door/protos/logproto"
	"hold-door/utils"
)

func Queryopenhttplog(ctx *gin.Context) {
	conn, err := utils.GrpcConn(utils.Trade)
	defer conn.Close()
	c := GrpcOpenHttpLogPkg.NewGrpcOpenHttpLogServerClient(conn)
	r, err := c.QueryOpenHttpLog(ctx, &GrpcOpenHttpLogPkg.GrpcOpenHttpLogSel{Page: 1, Pagecount: 1})
	if err != nil {

		return
	}

	ctx.JSON(200, r.Rows)
}
