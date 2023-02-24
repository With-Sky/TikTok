package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"tiktok/cmd/api/controllor"
	m "tiktok/pkg/middleware"
)

func Router(e *echo.Echo) {

	Router := e.Group("/douyin")
	e.Use(middleware.Logger())
	e.Use(middleware.CSRF())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	commentRouter := Router.Group("/comment")
	commentRouter.GET("/list/", controllor.CommentList)
	commentRouter.Use(m.JWTAuth)
	commentRouter.POST("/action/", controllor.CommentAction)

	favoriteRouter := Router.Group("/favorite")
	favoriteRouter.Use(m.JWTAuth)
	favoriteRouter.POST("/action/", controllor.FavoriteAction)
	favoriteRouter.GET("/list/", controllor.FavoriteList)

	publishRouter := Router.Group("/publish")
	publishRouter.Use(m.JWTAuth)
	publishRouter.GET("/list/", controllor.PublishList)
	publishRouter.POST("/action/", controllor.PublishAction)

	userRouter := Router.Group("/user")
	userRouter.POST("/login/", controllor.Login)
	userRouter.POST("/register/", controllor.Register)
	userRouter.Use(m.JWTAuth)
	userRouter.GET("/", controllor.UserInfo)

	relationRouter := Router.Group("/relation")
	relationRouter.Use(m.JWTAuth)
	relationRouter.POST("/action/", controllor.FollowAction)
	relationRouter.GET("/follow/list/", controllor.FollowList)
	relationRouter.GET("/follower/list/", controllor.FollowerList)
	relationRouter.GET("/friend/list/", controllor.FriendList)

	feedRouter := Router.Group("/feed")
	feedRouter.GET("", controllor.FeedAction)

	messageRouter := Router.Group("/message")
	messageRouter.Use(m.JWTAuth)
	messageRouter.GET("/chat/", controllor.MessageList)
	messageRouter.POST("/action/", controllor.SendMessage)

}
