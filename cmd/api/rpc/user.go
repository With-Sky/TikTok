package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/pkg/errno"
	"tiktok/pkg/middleware"
	utils "tiktok/pkg/utils"
	"time"
)

var userClient userservice.Client

// Comment RPC 客户端初始化
func InitUserRpc() {
	Config := utils.ConfigInit("TIKTOK_USER", "userConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	fmt.Println(EtcdAddress)
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})

	fmt.Println(r)
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")
	fmt.Println(ServiceName)
	c, err := userservice.NewClient(
		ServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		////client.WithMuxConnection(1), // mux
		client.WithRPCTimeout(300*time.Second),            // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
		//// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	//	client, err := userservice.NewClient(ServiceName, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	userClient = c
}

// 评论操作rpc
func Login(ctx context.Context, req *user.LoginReq) (*user.LoginRes, error) {
	res, err := userClient.Login(ctx, req)
	if res == nil {
		res = new(user.LoginRes)
		res.BaseResp = new(user.BaseResp)
		res.BaseResp.StatusCode = errno.ServiceErrCode
		res.BaseResp.StatusMessage = "服务请求失败"
		return res, err
	}
	if err != nil {
		return res, err
	}
	if res.BaseResp.StatusCode != 0 {
		return res, errno.NewErrNo(int64(res.BaseResp.StatusCode), res.BaseResp.StatusMessage)
	}
	return res, nil
}
func Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterRes, error) {
	res, err := userClient.Register(ctx, req)
	if res == nil {
		fmt.Println("失败1")
		res = new(user.RegisterRes)
		res.BaseResp = new(user.BaseResp)
		res.BaseResp.StatusCode = errno.ServiceErrCode
		res.BaseResp.StatusMessage = "服务请求失败"
		return res, err
	}
	if err != nil {
		fmt.Println("失败3")
		return res, err
	}
	if res.BaseResp.StatusCode != 0 {
		fmt.Println("失败2")
		return res, errno.NewErrNo(int64(res.BaseResp.StatusCode), res.BaseResp.StatusMessage)
	}
	fmt.Println("成功")
	return res, nil
}
func UserInfo(ctx context.Context, req *user.UserInfoReq) (*user.UserInfoRes, error) {
	res := new(user.UserInfoRes)
	var err error
	res, err = userClient.UserInfo(ctx, req)
	if res == nil {
		res = new(user.UserInfoRes)
		res.BaseResp = new(user.BaseResp)
		res.BaseResp.StatusCode = errno.ServiceErrCode
		res.BaseResp.StatusMessage = "服务请求失败"
		return res, err
	}
	if err != nil {
		return res, err
	}
	if res.BaseResp.StatusCode != 0 {
		return res, errno.NewErrNo(int64(res.BaseResp.StatusCode), res.BaseResp.StatusMessage)
	}
	fmt.Println("res")
	return res, nil
}
