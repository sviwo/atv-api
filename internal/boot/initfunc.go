package boot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sviwo/internal/service"
)

type NoDeferFunc struct {
	F    func(ctx context.Context) error
	Desc string
}

var InitFuncNoDeferListForIotCore = []NoDeferFunc{
	{service.TSLTable().CreateDatabase, "时序数据库创建"},
	{service.TdLogTable().CreateStable, "时序库日志表创建"},
}

func InitTDengineFunc(ctx context.Context, noDeferFuncList []NoDeferFunc) error {
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
