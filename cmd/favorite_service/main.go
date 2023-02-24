package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"tiktok/cmd/favorite_service/dal"
	"tiktok/cmd/favorite_service/global"
	favorite "tiktok/kitex_gen/favorite/favoriteservice"
	"tiktok/pkg/middleware"
	"tiktok/pkg/tracer"
	"tiktok/pkg/utils"
)

var (
	ServiceName = global.Config.Viper.GetString("Server.Name")
)

func init() {
	tracer.InitJaeger(ServiceName)
	global.LOG = utils.Zap(global.Config)
	dal.Init()
	fmt.Println("数据库初始化成功")
}
func main() {

	//Config := utils.ConfigInit("TIKTOK_FEED", "feedConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Etcd.Address"), global.Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceAddr := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Service.Address"), global.Config.Viper.GetInt("Service.Port"))
	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}
	global.LOG.Info("test")
	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                               // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		//server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
