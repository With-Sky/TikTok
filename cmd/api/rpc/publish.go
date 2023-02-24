package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/kitex_gen/publish"
	"tiktok/kitex_gen/publish/publishservice"
	"tiktok/pkg/errno"
	"tiktok/pkg/middleware"
	utils "tiktok/pkg/utils"
	"time"
)

var publishClient publishservice.Client

// Comment RPC 客户端初始化
func InitPublishRpc() {
	Config := utils.ConfigInit("TIKTOK_PUBLISH", "publishConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")

	c, err := publishservice.NewClient(
		ServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		//client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// 评论操作rpc
func Publish(ctx context.Context, req *publish.PublishReq) (*publish.PublishRes, error) {
	res, err := publishClient.Publish(ctx, req)
	if res == nil {
		res = new(publish.PublishRes)
		res.BaseResp = new(publish.BaseResp)
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
func PublishList(ctx context.Context, req *publish.PublishListReq) (*publish.PublishListRes, error) {
	res, err := publishClient.PublishList(ctx, req)
	if res == nil {
		res = new(publish.PublishListRes)
		res.BaseResp = new(publish.BaseResp)
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
