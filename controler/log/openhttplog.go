package log

import (
	"awesomeProject/protos/logproto"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Queryopenhttplog(ctx *gin.Context) {
	conn := utils.GrpcConn()
	defer conn.Close()
	c := GrpcOpenHttpLogPkg.NewGrpcOpenHttpLogServerClient(conn)
	r, err := c.QueryOpenHttpLog(ctx, &GrpcOpenHttpLogPkg.GrpcOpenHttpLogSel{Page: 1, Pagecount: 1})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//datetime := time.Unix(r.Rows[0].OperationTime.GetSeconds() , 0).Local()
	//
	//fmt.Println(datetime)

	ctx.JSON(200, r.Rows)
}
