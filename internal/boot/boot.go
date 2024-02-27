package boot

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/gomail.v2"
	"sviwo/internal/service"
)

/*
显式初始化和隐式初始化如何选择？
在业务场景下，非特殊必要，我们建议大家采用显式初始化的方式，以保证更好的可维护性。
*/

var (
	GID       *snowflake.Node
	NewDialer *gomail.Dialer
)

/*
隐式初始化:
隐式初始化一般通过包初始化方法init执行初始化。需要注意的是，如果初始化逻辑存在错误的可能，由于init方法的错误无法被上层捕获，初始化出错时往往直
接终止程序启动。隐式初始化的好处是不需要手动调用初始化方法，对于开发者隐藏了初始化细节，因此开发者没有心智负担。但是缺点也同样如此，开发者不知道
初始化细节，一旦出现错误时，很难快速定位错误原因。因此使用隐式初始化时，往往要求在初始化出错时将详细的错误以及堆栈信息打印出来便于错误定位。
*/
func init() {
}

/*
Boot
显式初始化:
显式初始化要求开发在程序启动时，如在main或者boot模块中，调用特定的方法来执行初始化操作。一般来说，基础组件的初始化往往采用隐式初始化多一些，因
为对于使用者来讲并不关心底层基础模块的初始化逻辑，而业务模块的初始化大多数会采用显式初始化。
*/
func Boot(ctx context.Context) {
	initSnowflake()
	initTDengine(ctx)
	initSendEmail(ctx)
}

func initTDengine(ctx context.Context) {
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

	// 启动失败的话请注释掉
	//if err := network.ReloadNetwork(context.Background()); err != nil {
	//	g.Log().Errorf(ctx, "载入网络错误,错误原因:%+#v", err)
	//}
}

func initSnowflake() {
	Node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	GID = Node
}

func initSendEmail(ctx context.Context) {
	NewDialer = gomail.NewDialer(
		g.Cfg().MustGet(ctx, "email.host").String(),
		g.Cfg().MustGet(ctx, "email.port").Int(),
		g.Cfg().MustGet(ctx, "email.username").String(),
		g.Cfg().MustGet(ctx, "email.password").String(),
	)
}
