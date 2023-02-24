package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/kitex_gen/comment"
	"tiktok/kitex_gen/comment/commentservice"
	"tiktok/pkg/errno"
	"tiktok/pkg/middleware"
	utils "tiktok/pkg/utils"
	"time"
)

var commentClient commentservice.Client

// Comment RPC 客户端初始化
func InitCommentRpc() {
	Config := utils.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	//utils.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")

	c, err := commentservice.NewClient(
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
	commentClient = c
}

// 评论操作rpc
func Comment(ctx context.Context, req *comment.CommentReq) (*comment.CommentRes, error) {
	res, err := commentClient.CommentAction(ctx, req)
	if res == nil {
		res = new(comment.CommentRes)
		res.BaseResp = new(comment.BaseResp)
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

// 删除评论rcp
func DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (*comment.DeleteCommentRes, error) {
	res, err := commentClient.DeleteComment(ctx, req)
	if res == nil {
		res = new(comment.DeleteCommentRes)
		res.BaseResp = new(comment.BaseResp)
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

// 评论列表rpc
func CommentList(ctx context.Context, req *comment.CommentListReq) (*comment.CommentListRes, error) {
	res, err := commentClient.CommentList(ctx, req)
	if res == nil {
		res = new(comment.CommentListRes)
		res.BaseResp = new(comment.BaseResp)
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
