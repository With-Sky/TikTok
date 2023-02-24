package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/pkg/errno"
	"tiktok/pkg/middleware"
	utils "tiktok/pkg/utils"

	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/kitex_gen/relation"
	"tiktok/kitex_gen/relation/relationservice"
	"time"
)

var relationClient relationservice.Client

// Comment RPC 客户端初始化
func InitRelationRpc() {
	Config := utils.ConfigInit("TIKTOK_RELATION", "relationConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")

	c, err := relationservice.NewClient(
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
	relationClient = c
}

// 评论操作rpc
func Follow(ctx context.Context, req *relation.FollowReq) (*relation.FollowRes, error) {
	res, err := relationClient.Follow(ctx, req)
	if res == nil {
		res = new(relation.FollowRes)
		res.BaseResp = new(relation.BaseResp)
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

func CancelFollow(ctx context.Context, req *relation.CancelFollowReq) (*relation.CancelFollowRes, error) {
	res, err := relationClient.CancelFollow(ctx, req)
	if res == nil {
		res = new(relation.CancelFollowRes)
		res.BaseResp = new(relation.BaseResp)
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

// 评论操作rpc
func FollowList(ctx context.Context, req *relation.FollowListReq) (*relation.FollowListRes, error) {
	res, err := relationClient.FollowList(ctx, req)
	if res == nil {
		res = new(relation.FollowListRes)
		res.BaseResp = new(relation.BaseResp)
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

// 评论操作rpc
func FollowerList(ctx context.Context, req *relation.FollowerListReq) (*relation.FollowerListRes, error) {
	res, err := relationClient.FollowerList(ctx, req)
	if res == nil {
		res = new(relation.FollowerListRes)
		res.BaseResp = new(relation.BaseResp)
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

//func FriendList(ctx context.Context, req *relation.FriendListReq) (*relation.FriendListRes, error) {
//	res, err := relationClient.FriendList(ctx, req)
//	if res == nil {
//		res = new(relation.FriendListRes)
//		res.BaseResp = new(relation.BaseResp)
//		res.BaseResp.StatusCode = errno.ServiceErrCode
//		res.BaseResp.StatusMessage = "服务请求失败"
//		return res, err
//	}
//
//	if err != nil {
//		return res, err
//	}
//	if res.BaseResp.StatusCode != 0 {
//		return res, errno.NewErrNo(int64(res.BaseResp.StatusCode), res.BaseResp.StatusMessage)
//	}
//	return res, nil
//}
