package trade

import (
	"awesomeProject/protos/tradeproto"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func QueryTradeDetail(ctx *gin.Context) {
	conn := utils.GrpcConnWithJwt()
	defer conn.Close()
	c := GrpcTradeDetail.NewGrpcTradeDetailServerClient(conn)
	r, err := c.QueryTradeDetail(ctx, &GrpcTradeDetail.GrpcTradeDetailSel{Page: 1, Pagecount: 1})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(200, r.Rows)
}
