package utils

import (
	"errors"
	"golang.org/x/net/context"
	"hold-door/config"
	"hold-door/models"
	service_authrization "hold-door/protos/authtoken"
	"time"
)

type ServiceName string

const (
	AuthCenter ServiceName = "AuthCenter"
	Trade      ServiceName = "Trade"
)

type tokenInfo struct {
	token   string
	expires time.Time
}

var tokenMap map[ServiceName]tokenInfo

func (sName ServiceName) MatchToken() (token string) {
	if tokenMap == nil {
		tokenMap = make(map[ServiceName]tokenInfo)
	}

	if val, ok := tokenMap[sName]; ok {
		if val.expires.Add(-60 * time.Second).After(time.Now()) {
			//校验token是否过期
			token = val.token
			return
		}
	}

	appSercret := config.GetConfig().Get("app_secret").(string)
	var serviceSecret string
	switch sName {
	case AuthCenter:
		serviceSecret = config.GetConfig().Get("backen_service.AuthCenter.service_secret").(string)
		break
	case Trade:
		serviceSecret = config.GetConfig().Get("backen_service.Trade.service_secret").(string)
		break
	default:
		return
	}

	tokenInfo, err := getServiceToken(appSercret, serviceSecret)
	if err != nil {
		return
	}

	tokenMap[sName] = tokenInfo
	token = tokenInfo.token
	return
}

func getServiceToken(appSercret string, serviceSercret string) (tokenInfo, error) {
	var token tokenInfo
	conn, err := GrpcConn(AuthCenter) //请求授权中心获取token
	defer conn.Close()
	if err != nil {

		return token, err
	}

	client := service_authrization.NewServiceAuthrizationGrpcServerClient(conn)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
	authInfo := &service_authrization.GrpcRequestToken{AppSecret: appSercret, ServiceSecret: serviceSercret}
	defer cancel()

	r, err2 := client.RequestToken(ctx, authInfo)
	if err2 != nil {
		return token, err2
	}

	if r.Result.Code != 1 {
		return token, errors.New(r.Result.Error_Message)
	}

	expire, s := time.ParseInLocation(models.DatetimeTemplate, r.Data.Expires, time.Local)
	if s != nil {
		return token, s
	}

	token = tokenInfo{token: r.Data.Token, expires: expire}
	return token, nil
}
