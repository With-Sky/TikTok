package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/kitex_gen/message"
	"tiktok/kitex_gen/message/massageservice"
	"tiktok/pkg/middleware"
	"tiktok/pkg/utils"
	"time"
)

var messageClient massageservice.Client

func InitMessageRpc() {
	Config := utils.ConfigInit("TIKTOK_MESSAGE", "messageConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	fmt.Println(EtcdAddress)
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})

	fmt.Println(r)
	if err != nil {
		panic(err)
	}
	ServiceName := Config.Viper.GetString("Server.Name")
	fmt.Println(ServiceName)
	c, err := massageservice.NewClient(
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
	messageClient = c
}

// SendMessage implements the MassageServiceImpl interface.
func SendMessage(ctx context.Context, req *message.MessageActionReq) (resp *message.MessageActionRes, err error) {
	res, _ := messageClient.SendMessage(ctx, req)
	return res, nil
}

// MessageList implements the MassageServiceImpl interface.
func MessageList(ctx context.Context, req *message.MessageChatReq) (*message.MessageChatRes, error) {
	res, _ := messageClient.MessageList(ctx, req)
	fmt.Println(res)
	return res, nil
}
func FriendList(ctx context.Context, req *message.FriendListReq) (*message.FriendListRes, error) {
	fmt.Println("++++++++++++++")
	res, _ := messageClient.FriendList(ctx, req)
	return res, nil
}
