package reporter

import (
	"context"
	"fmt"
	"sviwo/pkg/iotModel/topicModel"
)

func Init() (err error) {
	return nil
}

// ReportProperty 属性上报
func ReportProperty(ctx context.Context, data topicModel.TopicHandlerData) error {
	//对存在转义字符的进行全量替换
	fmt.Println("处理事件")
	return nil

}
