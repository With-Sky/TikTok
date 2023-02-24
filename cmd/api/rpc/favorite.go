package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/cmd/api/global"
	"tiktok/kitex_gen/favorite"
	"tiktok/kitex_gen/favorite/favoriteservice"
	"tiktok/pkg/errno"
	"tiktok/pkg/middleware"
	utils "tiktok/pkg/utils"
	"time"
)

var favoriteClient favoriteservice.Client

// Comment RPC 客户端初始化
func InitFavoriteRpc() {
	Config := utils.ConfigInit("TIKTOK_FAVORITE", "favoriteConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")

	c, err := favoriteservice.NewClient(
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
	favoriteClient = c
}

// 评论操作rpc
func Favorite(ctx context.Context, req *favorite.FavoriteReq) (*favorite.FavoriteRes, error) {
	//global.LOG.Info("点赞rpc")
	fmt.Println("点赞")
	res, err := favoriteClient.Favorite(ctx, req)
	if res == nil {
		fmt.Println("点赞服务失败")
		res = new(favorite.FavoriteRes)
		res.BaseResp = new(favorite.BaseResp)
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
func CancelFavorite(ctx context.Context, req *favorite.DeleteFavoriteReq) (*favorite.DeleteFavoriteRes, error) {
	global.LOG.Info("取消点赞rpc")
	fmt.Println("取消点赞rpc")
	res, err := favoriteClient.DeleteFavorite(ctx, req)
	if res == nil {
		res = new(favorite.DeleteFavoriteRes)
		res.BaseResp = new(favorite.BaseResp)
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
func FavoriteList(ctx context.Context, req *favorite.FavoriteListReq) (*favorite.FavoriteListRes, error) {
	global.LOG.Info("点赞列表rpc")
	fmt.Println("点赞列表rpc")
	res, err := favoriteClient.FavoriteList(ctx, req)
	if res == nil {
		res = new(favorite.FavoriteListRes)
		res.BaseResp = new(favorite.BaseResp)
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
