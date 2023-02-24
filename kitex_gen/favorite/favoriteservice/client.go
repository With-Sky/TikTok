// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	favorite "tiktok/kitex_gen/favorite"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Favorite(ctx context.Context, Req *favorite.FavoriteReq, callOptions ...callopt.Option) (r *favorite.FavoriteRes, err error)
	DeleteFavorite(ctx context.Context, Req *favorite.DeleteFavoriteReq, callOptions ...callopt.Option) (r *favorite.DeleteFavoriteRes, err error)
	FavoriteList(ctx context.Context, Req *favorite.FavoriteListReq, callOptions ...callopt.Option) (r *favorite.FavoriteListRes, err error)
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
	return &kFavoriteServiceClient{
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

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) Favorite(ctx context.Context, Req *favorite.FavoriteReq, callOptions ...callopt.Option) (r *favorite.FavoriteRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Favorite(ctx, Req)
}

func (p *kFavoriteServiceClient) DeleteFavorite(ctx context.Context, Req *favorite.DeleteFavoriteReq, callOptions ...callopt.Option) (r *favorite.DeleteFavoriteRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFavorite(ctx, Req)
}

func (p *kFavoriteServiceClient) FavoriteList(ctx context.Context, Req *favorite.FavoriteListReq, callOptions ...callopt.Option) (r *favorite.FavoriteListRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, Req)
}