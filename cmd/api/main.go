package main

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"tiktok/cmd/api/global"
	"tiktok/cmd/api/router"
	"tiktok/cmd/api/rpc"
	"tiktok/pkg/tracer"
	utils "tiktok/pkg/utils"
)

var (
	ServiceName = global.Config.Viper.GetString("Server.Name")
	Logger      *zap.Logger
)

func init() {
	global.LOG = utils.Zap(global.Config)
	rpc.InitRpc()
	tracer.InitJaeger(ServiceName)
}

func main() {
	//EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	//fmt.Println(EtcdAddress)
	//r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	//if err != nil {
	//	global.LOG.Error("启动失败")
	//	//panic(err)
	//}
	//ServiceAddr := fmt.Sprintf("%s:%d", Config.Viper.GetString("Service.Address"), Config.Viper.GetInt("Service.Port"))
	//addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//
	//server.WithRegistry(r) // registry
	//server.WithServiceAddr(addr)
	//server.WithSuite(trace.NewDefaultServerSuite())
	global.LOG.Error("+++++++++++++++++++++++")
	e := echo.New()
	router.Router(e)
	if err := e.Start(":8088"); err != nil {
		//logger.GetZapLogger().Errorf(err.Error())
	}
}

//package main

//func main() {
//	client, err := userservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		req := &user.LoginReq{
//			Username: "uuuu",
//			Password: "mjjj",
//		}
//		resp, err := client.Login(context.Background(), req)
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Println(resp)
//		time.Sleep(time.Second)
//	}
//}
