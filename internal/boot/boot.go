package boot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/logic/product"
	"sviwo/internal/logic/tdengine"
	"sviwo/internal/network/core/logic/model"
)

/*
显式初始化和隐式初始化如何选择？
在业务场景下，非特殊必要，我们建议大家采用显式初始化的方式，以保证更好的可维护性。
*/

/*
隐式初始化:
隐式初始化一般通过包初始化方法init执行初始化。需要注意的是，如果初始化逻辑存在错误的可能，由于init方法的错误无法被上层捕获，初始化出错时往往直
接终止程序启动。隐式初始化的好处是不需要手动调用初始化方法，对于开发者隐藏了初始化细节，因此开发者没有心智负担。但是缺点也同样如此，开发者不知道
初始化细节，一旦出现错误时，很难快速定位错误原因。因此使用隐式初始化时，往往要求在初始化出错时将详细的错误以及堆栈信息打印出来便于错误定位。
*/
//func init() {
//}

/*
Boot
显式初始化:
显式初始化要求开发在程序启动时，如在main或者boot模块中，调用特定的方法来执行初始化操作。一般来说，基础组件的初始化往往采用隐式初始化多一些，因
为对于使用者来讲并不关心底层基础模块的初始化逻辑，而业务模块的初始化大多数会采用显式初始化。
*/
func Boot(ctx context.Context) {

	deferFuncListIotCore, err := InitSystemDeferFunc(ctx)
	defer func() {
		for _, f := range deferFuncListIotCore {
			if f == nil {
				continue
			}
			if deferErr := f(ctx); deferErr != nil {
				fmt.Printf("defer func error: %s\n", deferErr.Error())
			}
		}
	}()

	err = InitSystem(ctx, InitFuncNoDeferListForIotCore)
	if err != nil {
		fmt.Printf("defer func error: %s\n", err.Error())
		panic(err)
	}
	if err != nil {
		fmt.Printf("defer func error: %s\n", err.Error())
	}

}

var InitFuncNoDeferListForIotCore = []NoDeferFunc{
	{AllSystemInit, "初始化系统变量"},
	{model.InitCoreLogic, "核心处理逻辑"},
	{InitSnowflake, "雪花算法ID"},
	{InitSendEmail, "发送邮件"},
	{InitAmqp, "Amqp"},
	{tdengine.CreateTDDatabase, "时序数据库创建"},
	{tdengine.CreateStable, "时序库日志表创建"},
	{product.DevInitNew().InitProductForTd, "时序库产品表初始化"},
	{product.DevInitNew().InitDeviceForTd, "时序库设备表初始化"},
	{product.DeviceNew().CacheDeviceDetailList, "缓存设备信息"},
}

var initFuncWithDeferList = []DeferFunc{
	{RunQueue, "消息队列"},
}

type NoDeferFunc struct {
	F    func(ctx context.Context) error
	Desc string
}

func InitSystem(ctx context.Context, noDeferFuncList []NoDeferFunc) error {
	for _, funcNode := range noDeferFuncList {
		g.Log().Infof(ctx, "开始初始化%s", funcNode.Desc)
		if err := funcNode.F(ctx); err != nil {
			return fmt.Errorf("初始化%s失败，错误原因是:%w", funcNode.Desc, err)
		} else {
			g.Log().Infof(ctx, "初始化%s成功", funcNode.Desc)
		}
	}
	return nil
}

func InitSystemDeferFunc(ctx context.Context) ([]func(context.Context) error, error) {
	deferFuncList := make([]func(context.Context) error, len(initFuncWithDeferList))
	for index, deferFuncNode := range initFuncWithDeferList {
		g.Log().Infof(ctx, "开始初始化%s", deferFuncNode.Desc)
		if err, deferFunc := deferFuncNode.F(ctx); err != nil {
			return nil, fmt.Errorf("初始化%s失败，错误原因是:%w", deferFuncNode.Desc, err)
		} else {
			deferFuncList[index] = deferFunc
			g.Log().Infof(ctx, "初始化%s成功", deferFuncNode.Desc)
		}
	}
	return deferFuncList, nil

}
