package utils

import (
	"context"
	"errors"
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

func GrpcConn(sName ServiceName) (conn *grpc.ClientConn, err error) {
	ip, err := getServiceIP(sName)
	if err != nil {
		return
	}

	return grpc.Dial(cast.ToString(ip), grpc.WithInsecure())
}

func GrpcConnWithJwt(sName ServiceName, token string) (conn *grpc.ClientConn, err error) {
	//ip := config.GetConfig().Get("backen_service.host")
	//token := config.GetConfig().Get("backen_service.token")

	ip, err := getServiceIP(sName)
	if err != nil {
		return
	}

	auth := TokenAuthentication{
		JwtToken: "bearer " + cast.ToString(token),
	}

	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithInsecure())
	// opts = append(opts, grpc.WithBlock())
	// // 使用自定义认证
	// opts = append(opts, grpc.WithPerRPCCredentials(&auth))

	//conn, err := grpc.Dial("39.100.12.48:5004",  grpc.WithPerRPCCredentials(&opts....))

	return grpc.Dial(cast.ToString(ip), grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
}

func getServiceIP(sName ServiceName) (ip string, err error) {

	switch sName {
	case AuthCenter:
		ip = config.GetConfig().Get("backen_service.AuthCenter.host").(string)
		break
	case Trade:
		ip = config.GetConfig().Get("backen_service.Trade.host").(string)
		break
	default:
		err = errors.New("未找到相关服务配置")
	}
	return
}
