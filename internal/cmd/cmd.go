package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"sviwo/internal/consts"
	"sviwo/internal/controller"
	"sviwo/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			gToken := StartGToken(ctx)
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					service.Middleware().ErrorHandler,
				)
				//不需要登录的路由组绑定
				group.Bind(
					controller.Common.GetVftCode,
					controller.User.Register,
				)
				//需要登录鉴权的路由组
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Middleware(
						service.Middleware().DecodeData,
					)
					//需要登录鉴权的接口放到这里
					group.Bind(
						controller.User.Info,
						controller.User.UpdatePassword,
						controller.User.EditInfo,
						controller.Common.ImgUpload,
						controller.Version.GetNewVersion,
					)
				})
			})

			//初始化其他组件
			initTool(ctx)

			s.Run()
			return nil
		},
	}
)

func initTool(ctx context.Context) {

	// TDengine 初始化
	if err := service.TSLTable().CreateDatabase(ctx); err != nil {
		g.Log().Fatal(ctx, "TDengine 数据库创建失败：", err)
	}
	if err := service.TdLogTable().CreateStable(ctx); err != nil {
		g.Log().Fatal(ctx, "TDengine 日志超级表创建失败：", err)
	}

	//mqtt启动
	//if err := mqtt.InitSystemMqtt(); err != nil {
	//	g.Log().Errorf(ctx, "MQTT 初始化mqtt客户端失败,失败原因:%+#v", err)
	//}
	//defer mqtt.Close()
}
