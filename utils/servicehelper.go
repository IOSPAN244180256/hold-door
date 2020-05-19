package utils

import (
	"context"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"hold-door/config"
)

//token认证
type TokenAuthentication struct {
	JwtToken string
}

//组织token认证的metadata信息
func (ta *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": ta.JwtToken,
	}, nil
}

//是否基于TLS认证进行安全传输
func (ta *TokenAuthentication) RequireTransportSecurity() bool {
	return false
}

func GrpcConn() *grpc.ClientConn {
	ip := config.GetConfig().Get("backen_service.host")
	conn, err := grpc.Dial(cast.ToString(ip), grpc.WithInsecure())

	if err != nil {
		return nil
	}

	return conn
}

func GrpcConnWithJwt() *grpc.ClientConn {
	//嵌套类型的解析方式
	//ip2:= config.LoadConfig().Get("element")
	//elementsMap := ip2.([]interface{})
	//for key, value := range elementsMap {
	//	fmt.Print(key)
	//	fmt.Print(value)
	//}

	ip := config.GetConfig().Get("backen_service.host")
	token := config.GetConfig().Get("backen_service.token")

	auth := TokenAuthentication{
		JwtToken: "bearer " + cast.ToString(token),
	}

	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithInsecure())
	// opts = append(opts, grpc.WithBlock())
	// // 使用自定义认证
	// opts = append(opts, grpc.WithPerRPCCredentials(&auth))

	//conn, err := grpc.Dial("39.100.12.48:5004",  grpc.WithPerRPCCredentials(&opts....))

	conn, err := grpc.Dial(cast.ToString(ip), grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))

	if err != nil {
		return nil
	}

	return conn
}
