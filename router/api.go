package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/controller"
	"github.com/taoshihan1991/imaptool/middleware"
)

func InitApiRouter(engine *gin.Engine) {
	//首页
	engine.GET("/", controller.Index)
	engine.POST("/check", controller.LoginCheckPass)
	engine.POST("/check_auth", middleware.JwtApiMiddleware, controller.MainCheckAuth)
	//前后聊天
	engine.GET("/chat_server", middleware.Ipblack, controller.NewChatServer)
	//获取消息
	engine.GET("/messages", controller.GetVisitorMessage)
	engine.GET("/message_notice", controller.SendVisitorNotice)
	//发送单条消息
	engine.POST("/message", middleware.Ipblack, controller.SendMessage)
	//发送关闭消息
	engine.GET("/message_close", controller.SendCloseMessage)
	//上传文件
	engine.POST("/uploadimg", middleware.Ipblack, controller.UploadImg)
	//获取未读消息数
	engine.GET("/message_status", controller.GetVisitorMessage)
	//设置消息已读
	engine.POST("/message_status", controller.GetVisitorMessage)

	//获取客服信息
	engine.GET("/kefuinfo", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetKefuInfo)
	engine.GET("/kefuinfo_setting", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetKefuInfoSetting)
	engine.POST("/kefuinfo", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostKefuInfo)
	engine.DELETE("/kefuinfo", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.DeleteKefuInfo)
	engine.GET("/kefulist", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetKefuList)
	//角色列表
	engine.GET("/roles", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetRoleList)
	engine.POST("/role", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostRole)
	//邮件夹列表
	engine.GET("/folders", controller.GetFolders)

	engine.GET("/mysql", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.MysqlGetConf)
	engine.POST("/mysql", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.MysqlSetConf)
	engine.GET("/visitors_online", controller.GetVisitorOnlines)
	engine.GET("/clear_online_tcp", controller.DeleteOnlineTcp)
	engine.POST("/visitor_login", middleware.Ipblack, controller.PostVisitorLogin)
	engine.POST("/visitor", controller.PostVisitor)
	engine.GET("/visitor", middleware.JwtApiMiddleware, controller.GetVisitor)
	engine.GET("/visitors", middleware.JwtApiMiddleware, controller.GetVisitors)
	engine.GET("/statistics", middleware.JwtApiMiddleware, controller.GetStatistics)
	//前台接口
	engine.GET("/notice", middleware.SetLanguage, controller.GetNotice)
	engine.POST("/notice", middleware.JwtApiMiddleware, controller.PostNotice)
	engine.DELETE("/notice", middleware.JwtApiMiddleware, controller.DelNotice)
	engine.POST("/notice_save", middleware.JwtApiMiddleware, controller.PostNoticeSave)
	engine.GET("/notices", middleware.JwtApiMiddleware, controller.GetNotices)
	engine.POST("/ipblack", middleware.JwtApiMiddleware, controller.PostIpblack)
	engine.DELETE("/ipblack", middleware.JwtApiMiddleware, controller.DelIpblack)
	engine.GET("/ipblacks_all", middleware.JwtApiMiddleware, controller.GetIpblacks)
	engine.GET("/configs", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.GetConfigs)
	engine.POST("/config", middleware.JwtApiMiddleware, middleware.RbacAuth, controller.PostConfig)
	//微信接口
	engine.GET("/micro_program", controller.GetCheckWeixinSign)
}
