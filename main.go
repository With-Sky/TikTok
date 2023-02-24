package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"tiktok/cmd/api/controllor"
	"tiktok/kitex_gen/user"
)

//var (
//	Config      = utils.ConfigInit("TIKTOK_COMMENT", "commentConfig")
//	ServiceName = Config.Viper.GetString("ServiceName")
//)
//
//func init() {
//	tracer.InitJaeger(ServiceName)
//}

type RegReq struct {
	UserName string `json:"username" validate:"required"`
	PassWord string `json:"password" validate:"required"`
}

func main() {
	//go service.RunMessageServer()
	//
	//r := gin.Default()
	//	var R=gin.New()
	var ()
	e := echo.New()

	e.POST("/douyin/user/register/", func(con echo.Context) error {
		fmt.Println("+")
		var registerRequestData RegReq
		//registerRequestData.UserName = con.QueryParam("username")
		//registerRequestData.PassWord = con.QueryParam("password")
		//global.LOG.Info("请求注册")
		if err := con.Bind(&registerRequestData); err != nil {
			controllor.FailWithMessage("获取请求失败", con)
			println(err)
			return err
		}
		return con.JSON(http.StatusNotExtended, user.RegisterRes{
			BaseResp: &user.BaseResp{
				StatusCode:    666,
				StatusMessage: "GGGG",
				ServiceTime:   0,
			},
			UserId: 0,
			Token:  "dd",
		})
	})
	e.Start(":8080")
	//EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	//r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	//if err != nil {
	//	panic(err)
	//}
	//ServiceAddr := fmt.Sprintf("%s:%d", Config.Viper.GetString("Service.Address"), Config.Viper.GetInt("Service.Port"))
	//addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//svr := comment.NewServer(new(CommentServiceImpl))
	////	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}), // server name
	////	server.WithMiddleware(middleware.CommonMiddleware),                               // middleWare
	////	server.WithMiddleware(middleware.ServerMiddleware),
	////	server.WithServiceAddr(addr),                                       // address
	////	server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
	////	server.WithMuxTransport(),                                          // Multiplex
	////	server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
	////	//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
	////	server.WithRegistry(r), // registry
	////)
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//r.Run()

}

//func  Login(c *gin.Context) {
//	fmt.Println("+++++")
//c.JSON(http.StatusOK, user.RegisterRes{
//		BaseResp: &user.BaseResp{
//			StatusCode:    200,
//			StatusMessage: "iii",
//			ServiceTime:   0,
//		},
//		UserId:   11,
//		Token:    "fdfdfdfh",
//	})
//
////fmt.Println(a.String())
//}
