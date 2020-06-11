package utils

import (
	"errors"
	"golang.org/x/net/context"
	"hold-door/config"
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
		token = val.token
		return
	}

	appSercret := config.GetConfig().Get("app_secret").(string)
	var serviceSecret string
	switch sName {
	case AuthCenter:
		serviceSecret = config.GetConfig().Get("service_secret.AuthCenter").(string)
		break
	case Trade:
		serviceSecret = config.GetConfig().Get("service_secret.Trade").(string)
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
	conn := GrpcConn()
	defer conn.Close()
	c := service_authrization.NewServiceAuthrizationGrpcServerClient(conn)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second)))
	authInfo := &service_authrization.GrpcRequestToken{AppSecret: appSercret, ServiceSecret: serviceSercret}
	defer cancel()

	r, err := c.RequestToken(ctx, authInfo)
	if err != nil {
		return token, err
	}

	if r.Result.Code != 1 {
		return token, errors.New(r.Result.Error_Message)
	}

	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	expire, _ := time.ParseInLocation(timeTemplate1, r.Data.Expires, time.Local)
	token = tokenInfo{token: r.Data.Token, expires: expire}

	return token, nil
}
