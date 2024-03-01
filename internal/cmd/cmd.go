package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"sviwo/internal/boot"
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
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					service.Middleware().ErrorHandler,
					//service.Middleware().DecodeData,
				)
				//不需要登录的路由组绑定
				group.Bind(
					controller.Common.GetVftCode,
					controller.Common.GetEccPublicKey,
					controller.User.Register,
				)
				//需要登录鉴权的路由组
				group.Group("/api", func(group *ghttp.RouterGroup) {
					err := StartGToken(ctx).Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					//需要登录鉴权的接口放到这里
					group.Bind(
						controller.Home.GetHomeData,
						controller.User.Info,
						controller.User.UpdatePassword,
						controller.User.EditInfo,
						controller.UserAuth.GetUserAuthInfo,
						controller.UserAuth.SubmitUserAuth,
						controller.Common.ImgUpload,
						controller.Version.GetNewVersion,
						controller.TravelRecord.GetTravelRecordList,
						controller.TravelRecord.Delete,
						controller.Car.GetCarList,
						controller.Car.BindingCar,
						controller.Car.DelCar,
						controller.Car.EnabledMobileKey,
						controller.Car.EnabledSpeedLimit,
					)
				})
			})
			boot.Boot(ctx)
			s.Run()
			return nil
		},
	}
)
