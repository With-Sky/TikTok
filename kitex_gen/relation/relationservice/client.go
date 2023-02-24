// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "tiktok/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Follow(ctx context.Context, Req *relation.FollowReq, callOptions ...callopt.Option) (r *relation.FollowRes, err error)
	CancelFollow(ctx context.Context, Req *relation.CancelFollowReq, callOptions ...callopt.Option) (r *relation.CancelFollowRes, err error)
	FollowList(ctx context.Context, Req *relation.FollowListReq, callOptions ...callopt.Option) (r *relation.FollowListRes, err error)
	FollowerList(ctx context.Context, Req *relation.FollowerListReq, callOptions ...callopt.Option) (r *relation.FollowerListRes, err error)
	FriendList(ctx context.Context, Req *relation.FriendListReq, callOptions ...callopt.Option) (r *relation.FriendListRes, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) Follow(ctx context.Context, Req *relation.FollowReq, callOptions ...callopt.Option) (r *relation.FollowRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, Req)
}

func (p *kRelationServiceClient) CancelFollow(ctx context.Context, Req *relation.CancelFollowReq, callOptions ...callopt.Option) (r *relation.CancelFollowRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelFollow(ctx, Req)
}

func (p *kRelationServiceClient) FollowList(ctx context.Context, Req *relation.FollowListReq, callOptions ...callopt.Option) (r *relation.FollowListRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, Req)
}

func (p *kRelationServiceClient) FollowerList(ctx context.Context, Req *relation.FollowerListReq, callOptions ...callopt.Option) (r *relation.FollowerListRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, Req)
}

func (p *kRelationServiceClient) FriendList(ctx context.Context, Req *relation.FriendListReq, callOptions ...callopt.Option) (r *relation.FriendListRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, Req)
}
