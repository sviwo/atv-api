package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
	"os"
	"sviwo/internal/controller"
	"sviwo/internal/service"
	"syscall"
)

func RunServer(ctx context.Context, stopSignal chan os.Signal) {
	var s = g.Server()

	// 错误状态码接管
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Writeln("404 - 没有找到…")
	})

	s.BindStatusHandler(403, func(r *ghttp.Request) {
		r.Response.Writeln("403 - 拒绝显示")
	})

	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().CORSHandler,
			service.Middleware().CtxHandler,
			service.Middleware().I18NHandler,
			service.Middleware().ResponseHandler,
			service.Middleware().ErrorHandler,
			//service.Middleware().DecodeData,
		)
		//不需要登录的路由组绑定
		group.Bind(
			controller.Common.GetVftCode,
			controller.Common.GetEccPublicKey,
			controller.User.Register,
			controller.User.UpdatePassword,
			controller.DeviceProperty, // 设备属性设置
		)
		//需要登录鉴权的路由组
		group.Group("/api", func(group *ghttp.RouterGroup) {
			err := StartGToken(ctx).Middleware(ctx, group)
			if err != nil {
				panic(err)
			}
			//需要登录鉴权的接口放到这里
			group.Bind(
				controller.Common.ImgUpload,
				controller.User.Info,
				controller.User.EditInfo,
				controller.Device.GetDeviceSecret,
				controller.Home,
				controller.UserAuth,
				controller.Version,
				controller.TravelRecord,
				controller.Car,
				controller.AppText,
				controller.AppVideos,
			)
		})
	})

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic 产生，错误:", err)
			}
		}()
		// https
		https := g.Cfg().MustGet(ctx, "server.https").Bool()
		if https {
			certFile := g.Cfg().MustGet(ctx, "server.httpsCertFile").String()
			keyFile := g.Cfg().MustGet(ctx, "server.httpsKeyFile").String()
			s.EnableHTTPS(certFile, keyFile)
		}

		go s.Run()
		select {
		case <-ctx.Done():
		}
		stopSignal <- syscall.SIGQUIT
	}()

}
