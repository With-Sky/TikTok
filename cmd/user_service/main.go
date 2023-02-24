// package main
//
// import (
//
//	"fmt"
//	"github.com/cloudwego/kitex/pkg/klog"
//	"github.com/cloudwego/kitex/pkg/limit"
//	//"github.com/cloudwego/kitex/pkg/remote/bound"
//	"github.com/cloudwego/kitex/pkg/rpcinfo"
//	"github.com/cloudwego/kitex/server"
//	etcd "github.com/kitex-contrib/registry-etcd"
//	trace "github.com/kitex-contrib/tracer-opentracing"
//	"log"
//	"net"
//	"tiktok/cmd/user_service/global"
//	user "tiktok/kitex_gen/user/userservice"
//	"tiktok/pkg/middleware"
//	"tiktok/pkg/tracer"
//	utils "tiktok/pkg/utils"
//
// )
//
// var (
//
//	ServiceName = global.Config.Viper.GetString("Server.Name")
//
// )
//
//	func init() {
//		tracer.InitJaeger(ServiceName)
//		global.LOG = utils.Zap(global.Config)
//		global.Redis = utils.Redis(global.Config, global.LOG)
//	}
//
// func main() {
//
//		//Config := utils.ConfigInit("TIKTOK_FEED", "feedConfig")
//		EtcdAddress := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Etcd.Address"), global.Config.Viper.GetInt("Etcd.Port"))
//		r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
//		if err != nil {
//			panic(err)
//		}
//		ServiceAddr := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Service.Address"), global.Config.Viper.GetInt("Service.Port"))
//		addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
//		if err != nil {
//			klog.Fatal(err)
//		}
//		svr := user.NewServer(new(UserServiceImpl),
//			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}), // server name
//			server.WithMiddleware(middleware.CommonMiddleware),                               // middleWare
//			server.WithMiddleware(middleware.ServerMiddleware),
//			server.WithServiceAddr(addr),                                       // address
//			server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
//			server.WithMuxTransport(),                                          // Multiplex
//			server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
//			//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
//			server.WithRegistry(r), // registry
//		)
//		err = svr.Run()
//
//		if err != nil {
//			log.Println(err.Error())
//		}
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"github.com/cloudwego/kitex/pkg/klog"
//	"github.com/cloudwego/kitex/pkg/limit"
//	"github.com/cloudwego/kitex/pkg/rpcinfo"
//	"github.com/cloudwego/kitex/server"
//	etcd "github.com/kitex-contrib/registry-etcd"
//	trace "github.com/kitex-contrib/tracer-opentracing"
//	"log"
//	"net"
//	"tiktok/cmd/comment_service/dal"
//	"tiktok/cmd/comment_service/global"
//		user "tiktok/kitex_gen/user/userservice"
//
//	"tiktok/pkg/middleware"
//	"tiktok/pkg/tracer"
//	"tiktok/pkg/utils"
//
// )
//
// var (
//
//	ServiceName = global.Config.Viper.GetString("Server.Name")
//
// )
//
//	func init() {
//		tracer.InitJaeger(ServiceName)
//		global.LOG = utils.Zap(global.Config)
//		dal.Init()
//	}
//
//	func main() {
//		EtcdAddress := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Etcd.Address"), global.Config.Viper.GetInt("Etcd.Port"))
//		fmt.Println(EtcdAddress)
//		r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
//		fmt.Println(err)
//		if err != nil {
//			panic(err)
//		}
//
//		ServiceAddr := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Server.Address"), global.Config.Viper.GetInt("Server.Port"))
//		addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
//		if err != nil {
//			klog.Fatal(err)
//		}
//		svr := user.NewServer(new(UserServiceImpl),
//			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}), // server name
//			server.WithMiddleware(middleware.CommonMiddleware),                               // middleWare
//			server.WithMiddleware(middleware.ServerMiddleware),
//			server.WithServiceAddr(addr),                                       // address
//			server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
//			server.WithMuxTransport(),                                          // Multiplex
//			server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
//			//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
//			server.WithRegistry(r), // registry
//		)
//		err = svr.Run()
//
//		if err != nil {
//			log.Println(err.Error())
//		}
//
// }
package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"tiktok/cmd/user_service/dal/db"
	"tiktok/pkg/middleware"

	"log"

	"tiktok/cmd/user_service/global"
	user "tiktok/kitex_gen/user/userservice"
	"tiktok/pkg/tracer"
	utils "tiktok/pkg/utils"
)

var (
	ServiceName = global.Config.Viper.GetString("Server.Name")
)

func init() {
	tracer.InitJaeger(ServiceName)
	fmt.Println("+++++++++++")
	global.LOG = utils.Zap(global.Config)
	fmt.Println("++++++++++++")
	db.Init()
	//global.Redis = utils.Redis(global.Config, global.LOG)
}
func main() {
	//Config := utils.ConfigInit("TIKTOK_FEED", "feedConfig")
	EtcdAddress := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Etcd.Address"), global.Config.Viper.GetInt("Etcd.Port"))
	r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceAddr := fmt.Sprintf("%s:%d", global.Config.Viper.GetString("Server.Address"), global.Config.Viper.GetInt("Server.Port"))
	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                               // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		//server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		////server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
